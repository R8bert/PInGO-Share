package main

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Settings struct {
	Theme                 string `json:"theme"`
	LogoPath              string `json:"logo"`
	BackgroundPath        string `json:"backgroundImage"`
	MaxUploadSize         int64  `json:"maxUploadSize"`         // In bytes
	UploadBoxTransparency int    `json:"uploadBoxTransparency"` // Transparency percentage (0-100)
	BlurIntensity         int    `json:"blurIntensity"`         // Blur intensity (0-24)
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 100 << 20

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5174", "http://localhost:5175"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// Upload endpoint with support for multiple files
	router.POST("/upload", func(c *gin.Context) {
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

		c.JSON(200, gin.H{
			"download_url": "/download/" + id,
			"files":        uploadedFiles,
			"count":        len(uploadedFiles),
		})
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

	// New settings save endpoint
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

	// New settings get endpoint
	router.GET("/settings", func(c *gin.Context) {
		var settings Settings
		if data, err := os.ReadFile("settings.json"); err == nil {
			json.Unmarshal(data, &settings)
			fmt.Println("Loaded settings from settings.json:", string(data))
		}
		c.JSON(http.StatusOK, settings)
	})

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
