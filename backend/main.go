package main

import (
	"archive/zip"
	"bufio"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtSecret []byte

// Load environment variables from .env file
func loadEnvFile() {
	envFile := ".env"
	// Try to find .env file in current directory or parent directory
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		envFile = "../.env"
		if _, err := os.Stat(envFile); os.IsNotExist(err) {
			log.Println("No .env file found, using system environment variables only")
			return
		}
	}

	file, err := os.Open(envFile)
	if err != nil {
		log.Printf("Warning: Could not open .env file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE format
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Only set if not already set in environment
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Warning: Error reading .env file: %v", err)
	} else {
		log.Println("Loaded environment variables from .env file")
	}
}

// Initialize JWT secret from environment variable
func initJWTSecret() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	if len(secret) < 32 {
		log.Fatal("JWT_SECRET must be at least 32 characters long")
	}
	jwtSecret = []byte(secret)
}

type Settings struct {
	ID                    int    `json:"id" db:"id"`
	Theme                 string `json:"theme" db:"theme"`
	LogoPath              string `json:"logo" db:"logo_path"`
	BackgroundPath        string `json:"backgroundImage" db:"background_path"`
	NavbarTitle           string `json:"navbarTitle" db:"navbar_title"`
	MaxUploadSize         int64  `json:"maxUploadSize" db:"max_upload_size"`                 // In bytes
	UploadBoxTransparency int    `json:"uploadBoxTransparency" db:"upload_box_transparency"` // Transparency percentage (0-100)
	BlurIntensity         int    `json:"blurIntensity" db:"blur_intensity"`                  // Blur intensity (0-24)
	MaxValidity           string `json:"maxValidity" db:"max_validity"`                      // "7days", "1month", "6months", "1year", "never"
	AllowRegistration     bool   `json:"allowRegistration" db:"allow_registration"`          // Allow user registration
}

type UploadRequest struct {
	Email    string `json:"email"`
	Validity string `json:"validity"` // "7days", "1month", "6months", "1year", "never"
}

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password_hash"`
	IsAdmin   bool      `json:"is_admin" db:"is_admin"`
	Avatar    string    `json:"avatar" db:"avatar"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type DeletionLog struct {
	ID             int        `json:"id" db:"id"`
	UserID         int        `json:"user_id" db:"user_id"`
	Username       string     `json:"username" db:"username"`
	UploadID       string     `json:"upload_id" db:"upload_id"`
	Files          string     `json:"files" db:"files"`
	TotalSize      int64      `json:"total_size" db:"total_size"`
	Email          string     `json:"email" db:"email"`
	DownloadURL    string     `json:"download_url" db:"download_url"`
	UploadedAt     time.Time  `json:"uploaded_at" db:"uploaded_at"`
	DeletedAt      time.Time  `json:"deleted_at" db:"deleted_at"`
	ExpiresAt      *time.Time `json:"expires_at" db:"expires_at"`
	IsReverse      bool       `json:"is_reverse" db:"is_reverse"`
	ReverseToken   string     `json:"reverse_token" db:"reverse_token"`
	DeletionReason string     `json:"deletion_reason" db:"deletion_reason"`
}

type Upload struct {
	ID           int        `json:"id" db:"id"`
	UserID       int        `json:"user_id" db:"user_id"`
	UploadID     string     `json:"upload_id" db:"upload_id"`
	Files        string     `json:"files" db:"files"`
	TotalSize    int64      `json:"total_size" db:"total_size"`
	Email        string     `json:"email" db:"email"`
	DownloadURL  string     `json:"download_url" db:"download_url"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	ExpiresAt    *time.Time `json:"expires_at" db:"expires_at"`
	IsAvailable  bool       `json:"is_available" db:"is_available"`   // User can toggle availability
	IsReverse    bool       `json:"is_reverse" db:"is_reverse"`       // Uploaded via reverse share token
	ReverseToken string     `json:"reverse_token" db:"reverse_token"` // Token used for reverse upload
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ReverseShareToken struct {
	ID        int        `json:"id" db:"id"`
	UserID    int        `json:"user_id" db:"user_id"`
	Token     string     `json:"token" db:"token"`
	Name      string     `json:"name" db:"name"`
	UsedCount int        `json:"used_count" db:"used_count"`
	MaxUses   int        `json:"max_uses" db:"max_uses"` // -1 for unlimited
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	ExpiresAt *time.Time `json:"expires_at" db:"expires_at"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func initDB() {
	// Get database configuration from environment variables
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "5432")
	user := getEnvOrDefault("DB_USER", "user")
	password := getEnvOrDefault("DB_PASSWORD", "pass")
	dbname := getEnvOrDefault("DB_NAME", "filesharing")
	sslmode := getEnvOrDefault("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	createTables()
	fmt.Println("Database connected successfully")
}

// Helper function to get environment variable with default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// File type validation
func isAllowedFileType(filename, contentType string) bool {
	// Get allowed file types from environment variable
	allowedTypes := getEnvOrDefault("ALLOWED_FILE_TYPES", "image/jpeg,image/png,image/gif,image/webp,application/pdf,text/plain,application/zip,application/x-zip-compressed,video/mp4,video/avi,audio/mpeg,audio/wav")
	allowedTypesSlice := strings.Split(allowedTypes, ",")

	// Check MIME type
	for _, allowedType := range allowedTypesSlice {
		if strings.TrimSpace(allowedType) == contentType {
			return true
		}
	}

	// Also check file extension as backup
	ext := strings.ToLower(filepath.Ext(filename))
	allowedExtensions := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
		".pdf": true, ".txt": true, ".zip": true,
		".mp4": true, ".avi": true, ".mov": true,
		".mp3": true, ".wav": true, ".ogg": true,
		".doc": true, ".docx": true, ".xls": true, ".xlsx": true, ".ppt": true, ".pptx": true,
	}

	return allowedExtensions[ext]
}

// Sanitize filename to prevent directory traversal
func sanitizeFilename(filename string) string {
	// Remove any path separators and only keep the base filename
	filename = filepath.Base(filename)
	// Remove any potentially dangerous characters
	filename = strings.ReplaceAll(filename, "..", "")
	filename = strings.ReplaceAll(filename, "/", "")
	filename = strings.ReplaceAll(filename, "\\", "")
	return filename
}

// Simple rate limiting middleware
var rateLimitMap = make(map[string][]time.Time)
var rateLimitMutex sync.RWMutex

func rateLimitMiddleware(requestsPerMinute int) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		now := time.Now()

		rateLimitMutex.Lock()
		defer rateLimitMutex.Unlock()

		// Clean old entries
		var validRequests []time.Time
		if requests, exists := rateLimitMap[clientIP]; exists {
			for _, requestTime := range requests {
				if now.Sub(requestTime) < time.Minute {
					validRequests = append(validRequests, requestTime)
				}
			}
		}

		// Check if limit exceeded
		if len(validRequests) >= requestsPerMinute {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		// Add current request
		validRequests = append(validRequests, now)
		rateLimitMap[clientIP] = validRequests

		c.Next()
	}
}

func createTables() {
	// Users table with admin field
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		is_admin BOOLEAN DEFAULT FALSE,
		is_blocked BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Settings table
	settingsTable := `
	CREATE TABLE IF NOT EXISTS settings (
		id SERIAL PRIMARY KEY,
		theme VARCHAR(50) DEFAULT 'light',
		logo_path VARCHAR(500),
		background_path VARCHAR(500),
		navbar_title VARCHAR(255) DEFAULT 'PInGO Share',
		max_upload_size BIGINT DEFAULT 104857600,
		upload_box_transparency INTEGER DEFAULT 0,
		blur_intensity INTEGER DEFAULT 0,
		max_validity VARCHAR(20) DEFAULT '7days',
		allow_registration BOOLEAN DEFAULT TRUE,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Uploads table with availability and reverse share support
	uploadsTable := `
	CREATE TABLE IF NOT EXISTS uploads (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		upload_id VARCHAR(255) UNIQUE NOT NULL,
		files TEXT NOT NULL,
		total_size BIGINT NOT NULL,
		email VARCHAR(255),
		download_url VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMP NULL,
		is_available BOOLEAN DEFAULT TRUE,
		is_reverse BOOLEAN DEFAULT FALSE,
		reverse_token VARCHAR(255)
	);`

	// Reverse share tokens table
	reverseTokensTable := `
	CREATE TABLE IF NOT EXISTS reverse_share_tokens (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		token VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		used_count INTEGER DEFAULT 0,
		max_uses INTEGER DEFAULT -1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMP NULL
	);`

	// Deletion logs table to track all deleted files
	deletionLogsTable := `
	CREATE TABLE IF NOT EXISTS deletion_logs (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
		username VARCHAR(255) NOT NULL,
		upload_id VARCHAR(255) NOT NULL,
		files TEXT NOT NULL,
		total_size BIGINT NOT NULL,
		email VARCHAR(255),
		download_url VARCHAR(255) NOT NULL,
		uploaded_at TIMESTAMP NOT NULL,
		deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMP NULL,
		is_reverse BOOLEAN DEFAULT FALSE,
		reverse_token VARCHAR(255),
		deletion_reason VARCHAR(500) DEFAULT 'User deleted'
	);`

	// Add is_admin column to existing users if it doesn't exist
	alterUsersTable := `
	ALTER TABLE users ADD COLUMN IF NOT EXISTS is_admin BOOLEAN DEFAULT FALSE;
	ALTER TABLE users ADD COLUMN IF NOT EXISTS is_blocked BOOLEAN DEFAULT FALSE;
	ALTER TABLE users ADD COLUMN IF NOT EXISTS avatar VARCHAR(500);`

	// Alter uploads table to add new columns
	alterUploadsTable := `
	ALTER TABLE uploads ALTER COLUMN expires_at DROP NOT NULL;
	ALTER TABLE uploads ADD COLUMN IF NOT EXISTS is_available BOOLEAN DEFAULT TRUE;
	ALTER TABLE uploads ADD COLUMN IF NOT EXISTS is_reverse BOOLEAN DEFAULT FALSE;
	ALTER TABLE uploads ADD COLUMN IF NOT EXISTS reverse_token VARCHAR(255);`

	// Alter settings table to add allow_registration
	alterSettingsTable := `
	ALTER TABLE settings ADD COLUMN IF NOT EXISTS allow_registration BOOLEAN DEFAULT TRUE;
	ALTER TABLE settings ADD COLUMN IF NOT EXISTS navbar_title VARCHAR(255) DEFAULT 'PInGO Share';`

	// Create indexes
	indexes := `
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	CREATE INDEX IF NOT EXISTS idx_uploads_user_id ON uploads(user_id);
	CREATE INDEX IF NOT EXISTS idx_uploads_upload_id ON uploads(upload_id);
	CREATE INDEX IF NOT EXISTS idx_uploads_expires_at ON uploads(expires_at);
	CREATE INDEX IF NOT EXISTS idx_reverse_tokens_token ON reverse_share_tokens(token);
	CREATE INDEX IF NOT EXISTS idx_reverse_tokens_user_id ON reverse_share_tokens(user_id);
	`

	// Insert default settings if none exist
	defaultSettings := `
	INSERT INTO settings (theme, max_upload_size, upload_box_transparency, blur_intensity, max_validity, allow_registration)
	SELECT 'light', 104857600, 0, 0, '7days', TRUE
	WHERE NOT EXISTS (SELECT 1 FROM settings);`

	tables := []string{usersTable, settingsTable, uploadsTable, reverseTokensTable, deletionLogsTable, alterUsersTable, alterUploadsTable, alterSettingsTable, indexes, defaultSettings}
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			fmt.Printf("Error executing SQL: %v\n", err)
		}
	}
	fmt.Println("Database tables created successfully")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(userID int) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func validateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			// Try to get token from cookie
			if cookie, err := c.Cookie("auth_token"); err == nil {
				tokenString = cookie
			}
		} else {
			// Remove "Bearer " prefix if present
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
			c.Abort()
			return
		}

		claims, err := validateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

func adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		var isAdmin bool
		err := db.QueryRow("SELECT is_admin FROM users WHERE id = $1", userID).Scan(&isAdmin)
		if err != nil || !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func blockCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		var isBlocked bool
		err := db.QueryRow("SELECT COALESCE(is_blocked, false) FROM users WHERE id = $1", userID).Scan(&isBlocked)
		if err != nil {
			// If column doesn't exist, assume user is not blocked
			if strings.Contains(err.Error(), "column") && strings.Contains(err.Error(), "is_blocked") {
				c.Next()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user status"})
			c.Abort()
			return
		}

		if isBlocked {
			c.JSON(http.StatusForbidden, gin.H{"error": "Account blocked - uploads are not allowed"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func calculateExpiryTime(validity string) *time.Time {
	now := time.Now()
	var expiry time.Time

	switch validity {
	case "7days":
		expiry = now.Add(7 * 24 * time.Hour)
	case "1month":
		expiry = now.AddDate(0, 1, 0)
	case "6months":
		expiry = now.AddDate(0, 6, 0)
	case "1year":
		expiry = now.AddDate(1, 0, 0)
	case "never":
		return nil // NULL for never expires
	default:
		expiry = now.Add(7 * 24 * time.Hour) // Default to 7 days
	}

	return &expiry
}

func isValidityAllowed(requestedValidity, maxValidity string) bool {
	validityOrder := map[string]int{
		"7days":   1,
		"1month":  2,
		"6months": 3,
		"1year":   4,
		"never":   5,
	}

	requestedLevel, exists1 := validityOrder[requestedValidity]
	maxLevel, exists2 := validityOrder[maxValidity]

	if !exists1 || !exists2 {
		return false
	}

	return requestedLevel <= maxLevel
}

func getSettings() (Settings, error) {
	var settings Settings
	err := db.QueryRow(`
		SELECT id, theme, COALESCE(logo_path, ''), COALESCE(background_path, ''), 
		       COALESCE(navbar_title, 'PInGO Share'), max_upload_size, upload_box_transparency, blur_intensity, max_validity, COALESCE(allow_registration, TRUE)
		FROM settings ORDER BY id LIMIT 1
	`).Scan(&settings.ID, &settings.Theme, &settings.LogoPath, &settings.BackgroundPath,
		&settings.NavbarTitle, &settings.MaxUploadSize, &settings.UploadBoxTransparency, &settings.BlurIntensity, &settings.MaxValidity, &settings.AllowRegistration)

	return settings, err
}

func main() {
	// Load environment variables from .env file first
	loadEnvFile()

	// Initialize JWT secret
	initJWTSecret()

	initDB()
	defer db.Close()

	router := gin.Default()
	router.MaxMultipartMemory = 100 << 20

	// Add security headers middleware
	router.Use(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Next()
	})

	// Enable CORS with environment-based configuration
	config := cors.DefaultConfig()
	allowedOrigins := getEnvOrDefault("ALLOWED_ORIGINS", "http://localhost:5173,http://localhost:5174,http://localhost:5175")
	config.AllowOrigins = strings.Split(allowedOrigins, ",")
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Auth routes with rate limiting
	router.POST("/register", rateLimitMiddleware(5), func(c *gin.Context) {
		// Check if registration is allowed
		settings, err := getSettings()
		if err == nil && !settings.AllowRegistration {
			c.JSON(http.StatusForbidden, gin.H{"error": "User registration is currently disabled"})
			return
		}

		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if user already exists
		var existingID int
		dbErr := db.QueryRow("SELECT id FROM users WHERE email = $1 OR username = $2", req.Email, req.Username).Scan(&existingID)
		if dbErr == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		// Check if this will be the first user (should be admin)
		var userCount int
		err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&userCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user count"})
			return
		}
		isFirstUser := userCount == 0

		// Hash password
		hashedPassword, err := hashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// Create user (first user is automatically admin)
		var userID int
		err = db.QueryRow("INSERT INTO users (username, email, password_hash, is_admin) VALUES ($1, $2, $3, $4) RETURNING id",
			req.Username, req.Email, hashedPassword, isFirstUser).Scan(&userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Generate JWT
		token, err := generateJWT(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Set cookie
		c.SetCookie("auth_token", token, 86400, "/", "", false, true)

		message := "User created successfully"
		if isFirstUser {
			message = "First user created successfully with admin privileges"
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": message,
			"token":   token,
			"user": gin.H{
				"id":       userID,
				"username": req.Username,
				"email":    req.Email,
				"is_admin": isFirstUser,
			},
		})
	})

	router.POST("/login", rateLimitMiddleware(10), func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get user from database
		var user User
		err := db.QueryRow("SELECT id, username, email, password_hash, is_admin, COALESCE(avatar, '') as avatar FROM users WHERE email = $1", req.Email).
			Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsAdmin, &user.Avatar)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Check password
		if !checkPasswordHash(req.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT
		token, err := generateJWT(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Set cookie
		c.SetCookie("auth_token", token, 86400, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
				"is_admin": user.IsAdmin,
				"avatar":   user.Avatar,
			},
		})
	})

	router.POST("/logout", func(c *gin.Context) {
		c.SetCookie("auth_token", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	})

	router.GET("/me", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		var user User
		err := db.QueryRow("SELECT id, username, email, is_admin, COALESCE(avatar, '') as avatar, created_at FROM users WHERE id = $1", userID).
			Scan(&user.ID, &user.Username, &user.Email, &user.IsAdmin, &user.Avatar, &user.CreatedAt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"id":         user.ID,
				"username":   user.Username,
				"email":      user.Email,
				"is_admin":   user.IsAdmin,
				"avatar":     user.Avatar,
				"created_at": user.CreatedAt,
			},
		})
	})

	// Avatar upload endpoint
	router.POST("/avatar", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		// Get current user info to get username and old avatar
		var user User
		err := db.QueryRow("SELECT id, username, email, is_admin, COALESCE(avatar, '') as avatar, created_at FROM users WHERE id = $1", userID).
			Scan(&user.ID, &user.Username, &user.Email, &user.IsAdmin, &user.Avatar, &user.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
			return
		}

		// Parse multipart form
		file, header, err := c.Request.FormFile("avatar")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No avatar file provided"})
			return
		}
		defer file.Close()

		// Validate file type (only PNG, JPG, JPEG, and GIF)
		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/jpg":  true,
			"image/png":  true,
			"image/gif":  true,
		}

		contentType := header.Header.Get("Content-Type")
		if !allowedTypes[contentType] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only PNG, JPG, JPEG, and GIF are allowed"})
			return
		}

		// Validate file size (max 5MB)
		if header.Size > 5*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File size too large. Maximum 5MB allowed"})
			return
		}

		// Create avatars directory if it doesn't exist
		avatarsDir := "./avatars"
		if _, err := os.Stat(avatarsDir); os.IsNotExist(err) {
			os.MkdirAll(avatarsDir, 0755)
		}

		// Read file content for checksum
		fileContent, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}

		// Generate checksum
		hash := sha256.Sum256(fileContent)
		hashString := hex.EncodeToString(hash[:])

		// Get file extension
		ext := filepath.Ext(header.Filename)
		if ext == "" {
			// Determine extension from content type
			switch contentType {
			case "image/jpeg":
				ext = ".jpg"
			case "image/png":
				ext = ".png"
			case "image/gif":
				ext = ".gif"
			default:
				ext = ".jpg"
			}
		}

		// Create new filename: username + checksum + extension
		filename := user.Username + "$" + hashString + ext
		avatarPath := "/avatars/" + filename
		fullPath := "./avatars/" + filename

		// Delete old avatar if it exists and is different from new one
		if user.Avatar != "" && user.Avatar != avatarPath {
			oldAvatarPath := "." + user.Avatar
			if _, err := os.Stat(oldAvatarPath); err == nil {
				os.Remove(oldAvatarPath)
				fmt.Printf("Deleted old avatar: %s\n", oldAvatarPath)
			}
		}

		// Save new file
		err = os.WriteFile(fullPath, fileContent, 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar"})
			return
		}

		// Update user avatar in database
		_, err = db.Exec("UPDATE users SET avatar = $1 WHERE id = $2", avatarPath, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update avatar"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Avatar uploaded successfully",
			"avatar":  avatarPath,
		})
	})

	// Protected upload endpoint
	router.POST("/upload", rateLimitMiddleware(20), authMiddleware(), blockCheckMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		// Load current settings from database
		settings, err := getSettings()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load settings"})
			return
		}

		maxSize := settings.MaxUploadSize
		if maxSize == 0 {
			maxSize = 100 << 20 // Default to 100MB if not set
		}

		// Parse multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form"})
			return
		}

		files := form.File["files"]
		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no files"})
			return
		}

		// Check total size and validate file types
		var totalSize int64
		for _, file := range files {
			totalSize += file.Size

			// Validate file type by checking both extension and MIME type
			if !isAllowedFileType(file.Filename, file.Header.Get("Content-Type")) {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("File type not allowed: %s", file.Filename)})
				return
			}

			// Check individual file size (prevent single huge files)
			if file.Size > maxSize {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("File %s is too large (%d bytes)", file.Filename, file.Size)})
				return
			}
		}
		if totalSize > maxSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Total file size (%d bytes) exceeds maximum allowed size (%d bytes)", totalSize, maxSize)})
			return
		}

		// Get validity from form
		validity := c.PostForm("validity")
		if validity == "" {
			validity = "7days" // Default
		}

		// Check if requested validity is allowed
		if !isValidityAllowed(validity, settings.MaxValidity) {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Requested validity '%s' exceeds maximum allowed '%s'", validity, settings.MaxValidity)})
			return
		}

		// Generate a single ID for all files
		id := uuid.New().String()
		var uploadedFiles []string

		// Save each file
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not open file"})
				return
			}
			defer file.Close()

			// Sanitize filename to prevent directory traversal
			sanitizedFilename := sanitizeFilename(fileHeader.Filename)
			if sanitizedFilename == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filename"})
				return
			}

			// Ensure uploads directory exists and is secure
			uploadsDir := "./uploads"
			if err := os.MkdirAll(uploadsDir, 0755); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create uploads directory"})
				return
			}

			outPath := filepath.Join(uploadsDir, id+"_"+sanitizedFilename)
			out, err := os.Create(outPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save file"})
				return
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not copy file"})
				return
			}

			uploadedFiles = append(uploadedFiles, fileHeader.Filename)
		}

		// Get email from form
		email := c.PostForm("email")

		// Calculate expiry time
		expiresAt := calculateExpiryTime(validity)

		// Save upload to database
		filesJSON, _ := json.Marshal(uploadedFiles)
		downloadURL := "/download/" + id

		_, err = db.Exec(`
			INSERT INTO uploads (user_id, upload_id, files, total_size, email, download_url, expires_at, is_available, is_reverse)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`, userID, id, string(filesJSON), totalSize, email, downloadURL, expiresAt, true, false)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save upload record"})
			return
		}

		c.JSON(200, gin.H{
			"download_url": downloadURL,
			"files":        uploadedFiles,
			"count":        len(uploadedFiles),
			"expires_at":   expiresAt,
		})
	})

	// Get user's upload history
	router.GET("/uploads", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		rows, err := db.Query(`
			SELECT id, upload_id, files, total_size, email, download_url, created_at, expires_at, 
			       COALESCE(is_available, TRUE), COALESCE(is_reverse, FALSE), COALESCE(reverse_token, '')
			FROM uploads 
			WHERE user_id = $1 
			ORDER BY created_at DESC
		`, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch uploads"})
			return
		}
		defer rows.Close()

		var uploads []Upload
		for rows.Next() {
			var upload Upload
			err := rows.Scan(
				&upload.ID, &upload.UploadID, &upload.Files, &upload.TotalSize,
				&upload.Email, &upload.DownloadURL, &upload.CreatedAt, &upload.ExpiresAt,
				&upload.IsAvailable, &upload.IsReverse, &upload.ReverseToken,
			)
			if err != nil {
				continue
			}
			upload.UserID = userID
			uploads = append(uploads, upload)
		}

		c.JSON(http.StatusOK, gin.H{"uploads": uploads})
	})

	// Delete upload
	router.DELETE("/uploads/:id", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")
		uploadID := c.Param("id")

		// Get upload details to log before deletion
		var upload Upload
		var username string
		err := db.QueryRow(`
			SELECT u.upload_id, u.files, u.total_size, u.email, u.download_url, 
			       u.created_at, u.expires_at, u.is_reverse, u.reverse_token, users.username
			FROM uploads u
			JOIN users ON u.user_id = users.id
			WHERE u.user_id = $1 AND u.upload_id = $2
		`, userID, uploadID).Scan(
			&upload.UploadID, &upload.Files, &upload.TotalSize, &upload.Email,
			&upload.DownloadURL, &upload.CreatedAt, &upload.ExpiresAt,
			&upload.IsReverse, &upload.ReverseToken, &username)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Upload not found"})
			return
		}

		// Log deletion before deleting
		_, err = db.Exec(`
			INSERT INTO deletion_logs (
				user_id, username, upload_id, files, total_size, email, 
				download_url, uploaded_at, expires_at, is_reverse, reverse_token, deletion_reason
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		`, userID, username, upload.UploadID, upload.Files, upload.TotalSize,
			upload.Email, upload.DownloadURL, upload.CreatedAt, upload.ExpiresAt,
			upload.IsReverse, upload.ReverseToken, "User deleted")

		if err != nil {
			fmt.Printf("Failed to log deletion: %v\n", err)
			// Continue with deletion even if logging fails
		}

		// Parse files and delete them
		var files []string
		json.Unmarshal([]byte(upload.Files), &files)

		for _, filename := range files {
			filePath := "./uploads/" + upload.UploadID + "_" + filename
			os.Remove(filePath)
		}

		// Delete from database
		_, err = db.Exec("DELETE FROM uploads WHERE user_id = $1 AND upload_id = $2", userID, uploadID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete upload"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Upload deleted successfully"})
	})

	// Toggle upload availability
	router.PUT("/uploads/:id/availability", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")
		uploadID := c.Param("id")

		type AvailabilityRequest struct {
			IsAvailable bool `json:"is_available"`
		}

		var req AvailabilityRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update availability
		_, err := db.Exec("UPDATE uploads SET is_available = $1 WHERE user_id = $2 AND upload_id = $3",
			req.IsAvailable, userID, uploadID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update availability"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Upload availability updated successfully"})
	})

	// Update upload expiration date
	router.PUT("/uploads/:id/expiration", authMiddleware(), blockCheckMiddleware(), func(c *gin.Context) {
		uploadID := c.Param("id")
		userID := c.GetInt("user_id")

		type ExpirationRequest struct {
			Validity string `json:"validity" binding:"required"` // "7days", "1month", "6months", "1year", "never"
		}

		var req ExpirationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if user owns the upload
		var ownerID int
		err := db.QueryRow("SELECT user_id FROM uploads WHERE upload_id = $1", uploadID).Scan(&ownerID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Upload not found"})
			return
		}

		if ownerID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own uploads"})
			return
		}

		// Calculate new expiry time
		expiryTime := calculateExpiryTime(req.Validity)

		// Update the upload
		_, err = db.Exec("UPDATE uploads SET expires_at = $1 WHERE upload_id = $2", expiryTime, uploadID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expiration"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Upload expiration updated successfully"})
	})

	// Create reverse share token
	router.POST("/reverse-tokens", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		type TokenRequest struct {
			Name      string `json:"name" binding:"required"`
			MaxUses   int    `json:"max_uses"`   // -1 for unlimited
			ExpiresIn string `json:"expires_in"` // "7days", "1month", etc. or empty for no expiry
		}

		var req TokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate token
		token := uuid.New().String()

		// Calculate expiry
		var expiresAt *time.Time
		if req.ExpiresIn != "" {
			expiresAt = calculateExpiryTime(req.ExpiresIn)
		}

		maxUses := req.MaxUses
		if maxUses == 0 {
			maxUses = -1 // Default to unlimited
		}

		// Save token to database
		var tokenID int
		err := db.QueryRow(`
			INSERT INTO reverse_share_tokens (user_id, token, name, max_uses, expires_at)
			VALUES ($1, $2, $3, $4, $5) RETURNING id
		`, userID, token, req.Name, maxUses, expiresAt).Scan(&tokenID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":         tokenID,
			"token":      token,
			"name":       req.Name,
			"max_uses":   maxUses,
			"used_count": 0,
			"expires_at": expiresAt,
		})
	})

	// Get user's reverse share tokens
	router.GET("/reverse-tokens", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		rows, err := db.Query(`
			SELECT id, token, name, used_count, max_uses, created_at, expires_at
			FROM reverse_share_tokens 
			WHERE user_id = $1 
			ORDER BY created_at DESC
		`, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tokens"})
			return
		}
		defer rows.Close()

		var tokens []ReverseShareToken
		for rows.Next() {
			var token ReverseShareToken
			err := rows.Scan(
				&token.ID, &token.Token, &token.Name, &token.UsedCount,
				&token.MaxUses, &token.CreatedAt, &token.ExpiresAt,
			)
			if err != nil {
				continue
			}
			token.UserID = userID
			tokens = append(tokens, token)
		}

		c.JSON(http.StatusOK, gin.H{"tokens": tokens})
	})

	// Delete reverse share token
	router.DELETE("/reverse-tokens/:id", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")
		tokenID := c.Param("id")

		_, err := db.Exec("DELETE FROM reverse_share_tokens WHERE user_id = $1 AND id = $2", userID, tokenID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Token deleted successfully"})
	})

	// Upload via reverse share token (public endpoint)
	router.POST("/reverse-upload/:token", func(c *gin.Context) {
		token := c.Param("token")

		// Validate token
		var tokenData ReverseShareToken
		err := db.QueryRow(`
			SELECT id, user_id, name, used_count, max_uses, expires_at
			FROM reverse_share_tokens 
			WHERE token = $1
		`, token).Scan(&tokenData.ID, &tokenData.UserID, &tokenData.Name,
			&tokenData.UsedCount, &tokenData.MaxUses, &tokenData.ExpiresAt)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid token"})
			return
		}

		// Check if token is expired
		if tokenData.ExpiresAt != nil && tokenData.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Token has expired"})
			return
		}

		// Check if token has reached max uses
		if tokenData.MaxUses != -1 && tokenData.UsedCount >= tokenData.MaxUses {
			c.JSON(http.StatusForbidden, gin.H{"error": "Token has reached maximum uses"})
			return
		}

		// Load current settings from database for upload limits
		settings, err := getSettings()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load settings"})
			return
		}

		maxSize := settings.MaxUploadSize
		if maxSize == 0 {
			maxSize = 100 << 20 // Default to 100MB if not set
		}

		// Parse multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form"})
			return
		}

		files := form.File["files"]
		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no files"})
			return
		}

		// Check total size of all files
		var totalSize int64
		for _, file := range files {
			totalSize += file.Size
		}
		if totalSize > maxSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Total file size (%d bytes) exceeds maximum allowed size (%d bytes)", totalSize, maxSize)})
			return
		}

		// Generate a single ID for all files
		id := uuid.New().String()
		var uploadedFiles []string

		// Save each file
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not open file"})
				return
			}
			defer file.Close()

			outPath := "./uploads/" + id + "_" + fileHeader.Filename
			out, err := os.Create(outPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save file"})
				return
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not copy file"})
				return
			}

			uploadedFiles = append(uploadedFiles, fileHeader.Filename)
		}

		// Get email from form
		email := c.PostForm("email")

		// Get validity from form or use default
		validity := c.PostForm("validity")
		if validity == "" {
			validity = "7days" // Default
		}

		// Check if requested validity is allowed
		if !isValidityAllowed(validity, settings.MaxValidity) {
			validity = settings.MaxValidity // Use max allowed if requested exceeds limit
		}

		// Calculate expiry time
		expiresAt := calculateExpiryTime(validity)

		// Save upload to database (reverse upload)
		filesJSON, _ := json.Marshal(uploadedFiles)
		downloadURL := "/download/" + id

		_, err = db.Exec(`
			INSERT INTO uploads (user_id, upload_id, files, total_size, email, download_url, expires_at, is_available, is_reverse, reverse_token)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		`, tokenData.UserID, id, string(filesJSON), totalSize, email, downloadURL, expiresAt, true, true, token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save upload record"})
			return
		}

		// Update token usage count
		_, err = db.Exec("UPDATE reverse_share_tokens SET used_count = used_count + 1 WHERE id = $1", tokenData.ID)
		if err != nil {
			// Log error but don't fail the upload
			fmt.Printf("Failed to update token usage count: %v\n", err)
		}

		c.JSON(200, gin.H{
			"message":      "Upload successful",
			"download_url": downloadURL,
			"files":        uploadedFiles,
			"count":        len(uploadedFiles),
			"expires_at":   expiresAt,
		})
	})

	// Download endpoint - handles single files directly, multiple files as ZIP
	router.GET("/download/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Check if upload exists and is available
		var isAvailable bool
		var userID int
		err := db.QueryRow("SELECT COALESCE(is_available, TRUE), user_id FROM uploads WHERE upload_id = $1", id).Scan(&isAvailable, &userID)
		if err != nil {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}

		// Check if file is available or if the user who uploaded it is requesting it
		currentUserID := -1
		var tokenString string

		// Try Authorization header first
		if authHeader := c.GetHeader("Authorization"); authHeader != "" {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else if cookie, err := c.Cookie("auth_token"); err == nil {
			// Fallback to cookie if no Authorization header
			tokenString = cookie
		}

		// Validate token if present
		if tokenString != "" {
			if claims, err := validateJWT(tokenString); err == nil {
				currentUserID = claims.UserID
			}
		}

		// If file is not available and requester is not the owner, return expired message
		if !isAvailable && currentUserID != userID {
			c.JSON(410, gin.H{"error": "This file has expired or is no longer available"})
			return
		}

		files, _ := os.ReadDir("./uploads")

		// Find all files with this ID
		var matchingFiles []string
		for _, f := range files {
			if strings.HasPrefix(f.Name(), id+"_") {
				matchingFiles = append(matchingFiles, f.Name())
			}
		}

		if len(matchingFiles) == 0 {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}

		// If only one file, serve it directly
		if len(matchingFiles) == 1 {
			c.File("./uploads/" + matchingFiles[0])
			return
		}

		// Multiple files - create and serve ZIP
		c.Header("Content-Type", "application/zip")
		c.Header("Content-Disposition", "attachment; filename=\"files.zip\"")

		zipWriter := zip.NewWriter(c.Writer)
		defer zipWriter.Close()

		for _, fileName := range matchingFiles {
			// Extract original filename (remove ID prefix)
			originalName := strings.TrimPrefix(fileName, id+"_")

			file, err := os.Open("./uploads/" + fileName)
			if err != nil {
				continue
			}
			defer file.Close()

			zipFile, err := zipWriter.Create(originalName)
			if err != nil {
				continue
			}

			_, err = io.Copy(zipFile, file)
			if err != nil {
				continue
			}
		}
	})

	// Files metadata endpoint
	router.GET("/files/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Check if upload exists and is available
		var isAvailable bool
		var uploaderUserID int
		err := db.QueryRow("SELECT COALESCE(is_available, TRUE), user_id FROM uploads WHERE upload_id = $1", id).Scan(&isAvailable, &uploaderUserID)
		if err != nil {
			c.JSON(404, gin.H{"error": "files not found"})
			return
		}

		// Check if file is available or if the user who uploaded it is requesting it
		currentUserID := -1
		if tokenString := c.GetHeader("Authorization"); tokenString != "" {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			if claims, err := validateJWT(tokenString); err == nil {
				currentUserID = claims.UserID
			}
		} else if cookie, err := c.Cookie("auth_token"); err == nil {
			if claims, err := validateJWT(cookie); err == nil {
				currentUserID = claims.UserID
			}
		}

		// If file is not available and requester is not the owner, return expired message
		if !isAvailable && currentUserID != uploaderUserID {
			c.JSON(410, gin.H{"error": "This file has expired or is no longer available"})
			return
		}

		// Get uploader information from database
		var uploaderInfo struct {
			Username  string     `json:"username"`
			Avatar    string     `json:"avatar"`
			Email     string     `json:"email"`
			ExpiresAt *time.Time `json:"expirationDate"`
		}

		err = db.QueryRow(`
			SELECT u.username, COALESCE(u.avatar, '') as avatar, up.email, up.expires_at
			FROM uploads up 
			JOIN users u ON up.user_id = u.id 
			WHERE up.upload_id = $1
		`, id).Scan(&uploaderInfo.Username, &uploaderInfo.Avatar, &uploaderInfo.Email, &uploaderInfo.ExpiresAt)

		if err != nil {
			// If no uploader found, continue without uploader info
			uploaderInfo.Username = "Unknown"
			uploaderInfo.Avatar = ""
			uploaderInfo.Email = ""
			uploaderInfo.ExpiresAt = nil
		}

		files, _ := os.ReadDir("./uploads")

		// Find all files with this ID
		var fileInfos []map[string]interface{}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), id+"_") {
				// Extract original filename (remove ID prefix)
				originalName := strings.TrimPrefix(f.Name(), id+"_")

				// Get file info
				fileInfo, err := f.Info()
				if err != nil {
					continue
				}

				fileInfos = append(fileInfos, map[string]interface{}{
					"name": originalName,
					"size": fileInfo.Size(),
					"url":  "/download/" + id, // Use controlled download endpoint instead of direct file access
				})
			}
		}

		if len(fileInfos) == 0 {
			c.JSON(404, gin.H{"error": "files not found"})
			return
		}

		c.JSON(200, gin.H{
			"files":    fileInfos,
			"uploader": uploaderInfo,
		})
	})

	// Admin-only settings endpoints
	router.POST("/admin/settings", authMiddleware(), adminMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		// Get username for file naming
		var username string
		err := db.QueryRow("SELECT username FROM users WHERE id = $1", userID).Scan(&username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
			return
		}

		// Load existing settings
		settings, err := getSettings()
		if err != nil {
			// If no settings exist, create default
			settings = Settings{
				Theme:                 "light",
				MaxUploadSize:         104857600,
				UploadBoxTransparency: 0,
				BlurIntensity:         0,
				MaxValidity:           "7days",
				AllowRegistration:     true,
			}
		}

		theme := c.PostForm("theme")
		if theme != "" {
			settings.Theme = theme
		}

		maxValidity := c.PostForm("maxValidity")
		if maxValidity != "" {
			validOptions := []string{"7days", "1month", "6months", "1year", "never"}
			valid := false
			for _, option := range validOptions {
				if maxValidity == option {
					valid = true
					break
				}
			}
			if !valid {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maximum validity option"})
				return
			}
			settings.MaxValidity = maxValidity
		}

		// Handle navbar title
		navbarTitle := c.PostForm("navbarTitle")
		if navbarTitle != "" {
			settings.NavbarTitle = navbarTitle
		}

		// Handle logo upload
		logo, _, err := c.Request.FormFile("logo")
		if err == nil {
			defer logo.Close()

			// Delete old logo if it exists
			if settings.LogoPath != "" {
				oldLogoPath := "." + settings.LogoPath
				if _, err := os.Stat(oldLogoPath); err == nil {
					os.Remove(oldLogoPath)
					fmt.Printf("Deleted old logo: %s\n", oldLogoPath)
				}
			}

			logoDir := "./logos"
			if err := os.MkdirAll(logoDir, 0755); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create logo directory"})
				return
			}

			// Read file content to generate hash
			logoContent, err := io.ReadAll(logo)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read logo file"})
				return
			}

			// Generate SHA256 hash
			hash := sha256.Sum256(logoContent)
			hashString := hex.EncodeToString(hash[:])

			logoExt := filepath.Ext(c.PostForm("logo"))
			if logoExt == "" {
				logoExt = ".png"
			}

			// Create filename: username + $ + checksum + extension
			filename := username + "$" + hashString + logoExt
			logoPath := filepath.Join(logoDir, filename)
			out, err := os.Create(logoPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save logo"})
				return
			}
			defer out.Close()

			_, err = out.Write(logoContent)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write logo file"})
				return
			}

			settings.LogoPath = "/logos/" + filename
		}

		// Handle background image upload
		background, _, err := c.Request.FormFile("backgroundImage")
		if err == nil {
			defer background.Close()

			// Delete old background if it exists
			if settings.BackgroundPath != "" {
				oldBackgroundPath := "." + settings.BackgroundPath
				if _, err := os.Stat(oldBackgroundPath); err == nil {
					os.Remove(oldBackgroundPath)
				}
			}

			backgroundDir := "./backgrounds"
			if err := os.MkdirAll(backgroundDir, 0755); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create background directory"})
				return
			}

			backgroundContent, err := io.ReadAll(background)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read background file"})
				return
			}

			hash := sha256.Sum256(backgroundContent)
			hashString := hex.EncodeToString(hash[:])

			backgroundExt := filepath.Ext(c.PostForm("backgroundImage"))
			if backgroundExt == "" {
				backgroundExt = ".jpg"
			}
			backgroundPath := filepath.Join(backgroundDir, hashString+backgroundExt)
			out, err := os.Create(backgroundPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save background image"})
				return
			}
			defer out.Close()

			_, err = out.Write(backgroundContent)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write background file"})
				return
			}

			settings.BackgroundPath = "/backgrounds/" + hashString + backgroundExt
		}

		// Handle max upload size
		if maxSizeStr := c.PostForm("maxUploadSize"); maxSizeStr != "" {
			maxSize, err := parseSize(maxSizeStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid max upload size format"})
				return
			}
			settings.MaxUploadSize = maxSize
		}

		// Handle upload box transparency
		if transparencyStr := c.PostForm("uploadBoxTransparency"); transparencyStr != "" {
			transparency, err := strconv.Atoi(transparencyStr)
			if err != nil || transparency < 0 || transparency > 100 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transparency value (must be 0-100)"})
				return
			}
			settings.UploadBoxTransparency = transparency
		}

		// Handle blur intensity
		if blurStr := c.PostForm("blurIntensity"); blurStr != "" {
			blur, err := strconv.Atoi(blurStr)
			if err != nil || blur < 0 || blur > 24 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid blur intensity value (must be 0-24)"})
				return
			}
			settings.BlurIntensity = blur
		}

		// Handle allow registration
		if allowRegStr := c.PostForm("allowRegistration"); allowRegStr != "" {
			allowReg, err := strconv.ParseBool(allowRegStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid allow registration value (must be true or false)"})
				return
			}
			settings.AllowRegistration = allowReg
		}

		// Save settings to database
		if settings.ID == 0 {
			// Insert new settings
			err = db.QueryRow(`
				INSERT INTO settings (theme, logo_path, background_path, navbar_title, max_upload_size, upload_box_transparency, blur_intensity, max_validity, allow_registration)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
			`, settings.Theme, settings.LogoPath, settings.BackgroundPath, settings.NavbarTitle, settings.MaxUploadSize,
				settings.UploadBoxTransparency, settings.BlurIntensity, settings.MaxValidity, settings.AllowRegistration).Scan(&settings.ID)
		} else {
			// Update existing settings
			_, err = db.Exec(`
				UPDATE settings SET theme = $1, logo_path = $2, background_path = $3, navbar_title = $4, max_upload_size = $5,
				upload_box_transparency = $6, blur_intensity = $7, max_validity = $8, allow_registration = $9, updated_at = CURRENT_TIMESTAMP
				WHERE id = $10
			`, settings.Theme, settings.LogoPath, settings.BackgroundPath, settings.NavbarTitle, settings.MaxUploadSize,
				settings.UploadBoxTransparency, settings.BlurIntensity, settings.MaxValidity, settings.AllowRegistration, settings.ID)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save settings"})
			return
		}

		c.JSON(http.StatusOK, settings)
	})

	// Admin statistics endpoint
	router.GET("/admin/stats", authMiddleware(), adminMiddleware(), func(c *gin.Context) {
		type AdminStats struct {
			TotalUsers   int   `json:"totalUsers"`
			TotalUploads int   `json:"totalUploads"`
			StorageUsed  int64 `json:"storageUsed"`
		}

		var stats AdminStats

		// Get total users count
		err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&stats.TotalUsers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user count"})
			return
		}

		// Get total uploads count
		err = db.QueryRow("SELECT COUNT(*) FROM uploads").Scan(&stats.TotalUploads)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get upload count"})
			return
		}

		// Calculate total storage used by summing total_size from uploads
		var totalSize sql.NullInt64
		err = db.QueryRow(`
			SELECT COALESCE(SUM(total_size), 0) as total_size 
			FROM uploads
		`).Scan(&totalSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate storage usage"})
			return
		}

		if totalSize.Valid {
			stats.StorageUsed = totalSize.Int64
		}

		c.JSON(http.StatusOK, stats)
	})

	// Admin users list endpoint
	router.GET("/admin/users", authMiddleware(), adminMiddleware(), func(c *gin.Context) {
		type AdminUser struct {
			ID           int        `json:"id"`
			Username     string     `json:"username"`
			Email        string     `json:"email"`
			Avatar       *string    `json:"avatar"`
			IsAdmin      bool       `json:"isAdmin"`
			IsBlocked    bool       `json:"isBlocked"`
			UploadCount  int        `json:"uploadCount"`
			StorageUsed  int64      `json:"storageUsed"`
			CreatedAt    time.Time  `json:"createdAt"`
			LastActivity *time.Time `json:"lastActivity"`
		}

		rows, err := db.Query(`
			SELECT 
				u.id, u.username, u.email, COALESCE(u.avatar, ''), u.is_admin, 
				COALESCE(u.is_blocked, false) as is_blocked,
				u.created_at,
				COUNT(CASE WHEN up.id IS NOT NULL THEN 1 END) as upload_count,
				COALESCE(SUM(up.total_size), 0) as storage_used,
				MAX(up.created_at) as last_activity
			FROM users u
			LEFT JOIN uploads up ON u.id = up.user_id
			GROUP BY u.id, u.username, u.email, u.avatar, u.is_admin, u.is_blocked, u.created_at
			ORDER BY u.created_at DESC
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}
		defer rows.Close()

		var users []AdminUser
		for rows.Next() {
			var user AdminUser
			var lastActivity sql.NullTime
			var avatar sql.NullString

			err := rows.Scan(
				&user.ID, &user.Username, &user.Email, &avatar, &user.IsAdmin, &user.IsBlocked,
				&user.CreatedAt, &user.UploadCount, &user.StorageUsed, &lastActivity,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user data"})
				return
			}

			if avatar.Valid && avatar.String != "" {
				user.Avatar = &avatar.String
			}

			if lastActivity.Valid {
				user.LastActivity = &lastActivity.Time
			}

			users = append(users, user)
		}

		c.JSON(http.StatusOK, users)
	})

	// Admin user block/unblock endpoint
	router.POST("/admin/users/:id/block", authMiddleware(), adminMiddleware(), func(c *gin.Context) {
		userID := c.Param("id")

		type BlockRequest struct {
			Blocked bool `json:"blocked"`
		}

		var req BlockRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Check if target user is an admin
		var isTargetAdmin bool
		err := db.QueryRow("SELECT COALESCE(is_admin, false) FROM users WHERE id = $1", userID).Scan(&isTargetAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user status"})
			return
		}

		// Prevent blocking admins
		if isTargetAdmin && req.Blocked {
			c.JSON(http.StatusForbidden, gin.H{"error": "Cannot block admin users"})
			return
		}

		// Add is_blocked column if it doesn't exist
		_, err = db.Exec("ALTER TABLE users ADD COLUMN IF NOT EXISTS is_blocked BOOLEAN DEFAULT FALSE")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update database schema"})
			return
		}

		_, err = db.Exec("UPDATE users SET is_blocked = $1 WHERE id = $2", req.Blocked, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
			return
		}

		action := "unblocked"
		if req.Blocked {
			action = "blocked"
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s successfully", action)})
	})

	// Admin quick settings update endpoint
	router.POST("/admin/quick-settings", authMiddleware(), adminMiddleware(), func(c *gin.Context) {
		type QuickSettingRequest struct {
			Setting string      `json:"setting"`
			Value   interface{} `json:"value"`
		}

		var req QuickSettingRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Load current settings
		settings, err := getSettings()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load settings"})
			return
		}

		// Update the specific setting
		switch req.Setting {
		case "allowRegistration":
			if boolVal, ok := req.Value.(bool); ok {
				settings.AllowRegistration = boolVal
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for allowRegistration"})
				return
			}
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown setting"})
			return
		}

		// Save updated settings
		if settings.ID == 0 {
			err = db.QueryRow(`
				INSERT INTO settings (theme, logo_path, background_path, max_upload_size, upload_box_transparency, blur_intensity, max_validity, allow_registration)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING id
			`, settings.Theme, settings.LogoPath, settings.BackgroundPath, settings.MaxUploadSize,
				settings.UploadBoxTransparency, settings.BlurIntensity, settings.MaxValidity, settings.AllowRegistration).Scan(&settings.ID)
		} else {
			_, err = db.Exec(`
				UPDATE settings SET theme = $1, logo_path = $2, background_path = $3, max_upload_size = $4,
				upload_box_transparency = $5, blur_intensity = $6, max_validity = $7, allow_registration = $8
				WHERE id = $9
			`, settings.Theme, settings.LogoPath, settings.BackgroundPath, settings.MaxUploadSize,
				settings.UploadBoxTransparency, settings.BlurIntensity, settings.MaxValidity, settings.AllowRegistration, settings.ID)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update setting"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Setting updated successfully"})
	})

	// Temporary endpoint to make first user admin (remove in production)
	router.POST("/promote-first-admin", func(c *gin.Context) {
		_, err := db.Exec("UPDATE users SET is_admin = TRUE WHERE id = (SELECT MIN(id) FROM users)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to promote user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "First user promoted to admin"})
	})

	// Settings get endpoint (public)
	router.GET("/settings", func(c *gin.Context) {
		settings, err := getSettings()
		if err != nil {
			// Return default settings if none exist
			settings = Settings{
				Theme:                 "light",
				NavbarTitle:           "PInGO Share",
				MaxUploadSize:         104857600,
				UploadBoxTransparency: 0,
				BlurIntensity:         0,
				MaxValidity:           "7days",
				AllowRegistration:     true,
			}
		}
		c.JSON(http.StatusOK, settings)
	})

	// Cleanup expired uploads (run periodically)
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			// Delete expired uploads from database and files (excluding never-expiring uploads)
			rows, err := db.Query("SELECT upload_id, files FROM uploads WHERE expires_at IS NOT NULL AND expires_at < NOW()")
			if err != nil {
				continue
			}

			var expiredUploads []Upload
			for rows.Next() {
				var upload Upload
				rows.Scan(&upload.UploadID, &upload.Files)
				expiredUploads = append(expiredUploads, upload)
			}
			rows.Close()

			for _, upload := range expiredUploads {
				// Parse files and delete them
				var files []string
				json.Unmarshal([]byte(upload.Files), &files)

				for _, filename := range files {
					filePath := "./uploads/" + upload.UploadID + "_" + filename
					os.Remove(filePath)
				}
			}

			// Delete expired records from database
			db.Exec("DELETE FROM uploads WHERE expires_at IS NOT NULL AND expires_at < NOW()")
		}
	}()

	// Serve static files for logos, backgrounds, and avatars only
	// NOTE: uploads directory is NOT served statically for security - access is controlled through /download/:id endpoint
	router.Static("/logos", "./logos")
	router.Static("/backgrounds", "./backgrounds")
	router.Static("/avatars", "./avatars")

	// Ensure directories exist
	os.MkdirAll("./uploads", os.ModePerm)
	os.MkdirAll("./logos", os.ModePerm)
	os.MkdirAll("./backgrounds", os.ModePerm)
	os.MkdirAll("./avatars", os.ModePerm)

	// Start server on configurable port
	port := getEnvOrDefault("SERVER_PORT", "8080")
	log.Printf("Starting server on port %s", port)
	router.Run(":" + port)
}

// Helper function to parse size (e.g., "10MB", "500KB", "1GB")
func parseSize(sizeStr string) (int64, error) {
	var size int64
	var multiplier int64 = 1

	// Extract number and unit
	num := ""
	for _, r := range sizeStr {
		if r >= '0' && r <= '9' {
			num += string(r)
		} else {
			break
		}
	}
	unit := strings.ToUpper(strings.TrimPrefix(sizeStr, num))

	if num == "" {
		return 0, fmt.Errorf("invalid size format")
	}

	size, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0, err
	}

	switch unit {
	case "KB":
		multiplier = 1024
	case "MB":
		multiplier = 1024 * 1024
	case "GB":
		multiplier = 1024 * 1024 * 1024
	case "":
		// No unit, assume bytes
	default:
		return 0, fmt.Errorf("unsupported unit: %s", unit)
	}

	return size * multiplier, nil
}
