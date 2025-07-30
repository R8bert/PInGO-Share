package main

import (
	"archive/zip"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtSecret = []byte("your-super-secret-jwt-key-change-in-production")

type Settings struct {
	Theme                 string `json:"theme"`
	LogoPath              string `json:"logo"`
	BackgroundPath        string `json:"backgroundImage"`
	MaxUploadSize         int64  `json:"maxUploadSize"`         // In bytes
	UploadBoxTransparency int    `json:"uploadBoxTransparency"` // Transparency percentage (0-100)
	BlurIntensity         int    `json:"blurIntensity"`         // Blur intensity (0-24)
}

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Upload struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	UploadID    string    `json:"upload_id" db:"upload_id"`
	Files       string    `json:"files" db:"files"`
	TotalSize   int64     `json:"total_size" db:"total_size"`
	Email       string    `json:"email" db:"email"`
	DownloadURL string    `json:"download_url" db:"download_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ExpiresAt   time.Time `json:"expires_at" db:"expires_at"`
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

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func initDB() {
	connStr := "host=192.168.31.180 port=5555 user=pingo password=pingo_filetrasnfer dbname=pingo_db sslmode=disable"
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

func createTables() {
	// Users table
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Uploads table
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
		expires_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP + INTERVAL '7 days'
	);`

	// Create indexes
	indexes := `
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	CREATE INDEX IF NOT EXISTS idx_uploads_user_id ON uploads(user_id);
	CREATE INDEX IF NOT EXISTS idx_uploads_upload_id ON uploads(upload_id);
	CREATE INDEX IF NOT EXISTS idx_uploads_expires_at ON uploads(expires_at);
	`

	tables := []string{usersTable, uploadsTable, indexes}
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			panic(fmt.Sprintf("Failed to create table: %v", err))
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
			if strings.HasPrefix(tokenString, "Bearer ") {
				tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			}
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

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()
	router.MaxMultipartMemory = 100 << 20

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5174", "http://localhost:5175"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Auth routes
	router.POST("/register", func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if user already exists
		var existingID int
		err := db.QueryRow("SELECT id FROM users WHERE email = $1 OR username = $2", req.Email, req.Username).Scan(&existingID)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		// Hash password
		hashedPassword, err := hashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// Create user
		var userID int
		err = db.QueryRow("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
			req.Username, req.Email, hashedPassword).Scan(&userID)
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

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"token":   token,
			"user": gin.H{
				"id":       userID,
				"username": req.Username,
				"email":    req.Email,
			},
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get user from database
		var user User
		err := db.QueryRow("SELECT id, username, email, password_hash FROM users WHERE email = $1", req.Email).
			Scan(&user.ID, &user.Username, &user.Email, &user.Password)
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
		err := db.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = $1", userID).
			Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"id":         user.ID,
				"username":   user.Username,
				"email":      user.Email,
				"created_at": user.CreatedAt,
			},
		})
	})

	// Protected upload endpoint
	router.POST("/upload", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		// Load current max upload size
		var settings Settings
		if data, err := os.ReadFile("settings.json"); err == nil {
			json.Unmarshal(data, &settings)
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

		// Save upload to database
		filesJSON, _ := json.Marshal(uploadedFiles)
		downloadURL := "/download/" + id
		expiresAt := time.Now().Add(7 * 24 * time.Hour) // 7 days expiry

		_, err = db.Exec(`
			INSERT INTO uploads (user_id, upload_id, files, total_size, email, download_url, expires_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, userID, id, string(filesJSON), totalSize, email, downloadURL, expiresAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save upload record"})
			return
		}

		c.JSON(200, gin.H{
			"download_url": downloadURL,
			"files":        uploadedFiles,
			"count":        len(uploadedFiles),
		})
	})

	// Get user's upload history
	router.GET("/uploads", authMiddleware(), func(c *gin.Context) {
		userID := c.GetInt("user_id")

		rows, err := db.Query(`
			SELECT id, upload_id, files, total_size, email, download_url, created_at, expires_at
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

		// Get upload details to delete files
		var upload Upload
		err := db.QueryRow(`
			SELECT upload_id, files FROM uploads 
			WHERE user_id = $1 AND upload_id = $2
		`, userID, uploadID).Scan(&upload.UploadID, &upload.Files)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Upload not found"})
			return
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

	// Download endpoint - handles single files directly, multiple files as ZIP
	router.GET("/download/:id", func(c *gin.Context) {
		id := c.Param("id")
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
					"url":  "/uploads/" + f.Name(),
				})
			}
		}

		if len(fileInfos) == 0 {
			c.JSON(404, gin.H{"error": "files not found"})
			return
		}

		c.JSON(200, gin.H{"files": fileInfos})
	})

	// Settings save endpoint
	router.POST("/settings/save", func(c *gin.Context) {
		// Load existing settings
		var existingSettings Settings
		if data, err := os.ReadFile("settings.json"); err == nil {
			json.Unmarshal(data, &existingSettings)
		}

		// Use existing settings as base
		settings := existingSettings

		theme := c.PostForm("theme")
		if theme != "" {
			settings.Theme = theme
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
					fmt.Println("Deleted old logo:", oldLogoPath)
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
			fmt.Println("Logo hash sum:", hashString)

			logoExt := filepath.Ext(c.PostForm("logo"))
			if logoExt == "" {
				logoExt = ".png"
			}
			logoPath := filepath.Join(logoDir, hashString+logoExt)
			out, err := os.Create(logoPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save logo"})
				return
			}
			defer out.Close()

			// Write the content we already read
			_, err = out.Write(logoContent)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write logo file"})
				return
			}

			settings.LogoPath = "/logos/" + hashString + logoExt
			fmt.Println("Saved logo to:", settings.LogoPath)
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
					fmt.Println("Deleted old background:", oldBackgroundPath)
				}
			}

			backgroundDir := "./backgrounds"
			if err := os.MkdirAll(backgroundDir, 0755); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create background directory"})
				return
			}

			// Read file content to generate hash
			backgroundContent, err := io.ReadAll(background)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read background file"})
				return
			}

			// Generate SHA256 hash
			hash := sha256.Sum256(backgroundContent)
			hashString := hex.EncodeToString(hash[:])
			fmt.Println("Background hash sum:", hashString)

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

			// Write the content we already read
			_, err = out.Write(backgroundContent)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write background file"})
				return
			}

			settings.BackgroundPath = "/backgrounds/" + hashString + backgroundExt
			fmt.Println("Saved background to:", settings.BackgroundPath)
		}

		// Handle max upload size
		if maxSizeStr := c.PostForm("maxUploadSize"); maxSizeStr != "" {
			maxSize, err := parseSize(maxSizeStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid max upload size format"})
				return
			}
			settings.MaxUploadSize = maxSize
			fmt.Println("Set max upload size to:", maxSize, "bytes")
		}

		// Handle upload box transparency
		if transparencyStr := c.PostForm("uploadBoxTransparency"); transparencyStr != "" {
			transparency, err := strconv.Atoi(transparencyStr)
			if err != nil || transparency < 0 || transparency > 100 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transparency value (must be 0-100)"})
				return
			}
			settings.UploadBoxTransparency = transparency
			fmt.Println("Set upload box transparency to:", transparency, "%")
		}

		// Handle blur intensity
		if blurStr := c.PostForm("blurIntensity"); blurStr != "" {
			blur, err := strconv.Atoi(blurStr)
			if err != nil || blur < 0 || blur > 24 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid blur intensity value (must be 0-24)"})
				return
			}
			settings.BlurIntensity = blur
			fmt.Println("Set blur intensity to:", blur)
		}

		// Save settings to a file
		data, err := json.Marshal(settings)
		if err == nil {
			os.WriteFile("settings.json", data, 0644)
			fmt.Println("Saved settings to settings.json:", string(data))
		}

		c.JSON(http.StatusOK, settings)
	})

	// Settings get endpoint
	router.GET("/settings", func(c *gin.Context) {
		var settings Settings
		if data, err := os.ReadFile("settings.json"); err == nil {
			json.Unmarshal(data, &settings)
			fmt.Println("Loaded settings from settings.json:", string(data))
		}
		c.JSON(http.StatusOK, settings)
	})

	// Cleanup expired uploads (run periodically)
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			// Delete expired uploads from database and files
			rows, err := db.Query("SELECT upload_id, files FROM uploads WHERE expires_at < NOW()")
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
			db.Exec("DELETE FROM uploads WHERE expires_at < NOW()")
		}
	}()

	// Serve static files for logos, backgrounds, and uploads
	router.Static("/logos", "./logos")
	router.Static("/backgrounds", "./backgrounds")
	router.Static("/uploads", "./uploads")

	// Ensure directories exist
	os.MkdirAll("./uploads", os.ModePerm)
	os.MkdirAll("./logos", os.ModePerm)
	os.MkdirAll("./backgrounds", os.ModePerm)

	router.Run(":8080")
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
