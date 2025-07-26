package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Settings struct {
	Theme          string `json:"theme"`
	LogoPath       string `json:"logo"`
	BackgroundPath string `json:"backgroundImage"`
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 100 << 20

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// Existing upload endpoint
	router.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no file"})
			return
		}
		defer file.Close()

		id := uuid.New().String()
		outPath := "./uploads/" + id + "_" + header.Filename
		out, err := os.Create(outPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save file"})
			return
		}
		defer out.Close()

		io.Copy(out, file)
		c.JSON(200, gin.H{"download_url": "/download/" + id})
	})

	// Existing download endpoint
	router.GET("/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		files, _ := os.ReadDir("./uploads")
		for _, f := range files {
			if f.Name()[:len(id)] == id {
				c.File("./uploads/" + f.Name())
				return
			}
		}
		c.JSON(404, gin.H{"error": "not found"})
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
			logoDir := "./logos"
			if err := os.MkdirAll(logoDir, 0755); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create logo directory"})
				return
			}
			logoID := uuid.New().String()
			logoExt := filepath.Ext(c.PostForm("logo"))
			if logoExt == "" {
				logoExt = ".png"
			}
			logoPath := filepath.Join(logoDir, logoID+logoExt)
			out, err := os.Create(logoPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save logo"})
				return
			}
			defer out.Close()
			io.Copy(out, logo)
			settings.LogoPath = "/logos/" + logoID + logoExt
			fmt.Println("Saved logo to:", settings.LogoPath)
		}

		// Handle background image upload
		background, _, err := c.Request.FormFile("backgroundImage")
		if err == nil {
			defer background.Close()
			backgroundDir := "./backgrounds"
			if err := os.MkdirAll(backgroundDir, 0755); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create background directory"})
				return
			}
			backgroundID := uuid.New().String()
			backgroundExt := filepath.Ext(c.PostForm("backgroundImage"))
			if backgroundExt == "" {
				backgroundExt = ".jpg"
			}
			backgroundPath := filepath.Join(backgroundDir, backgroundID+backgroundExt)
			out, err := os.Create(backgroundPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save background image"})
				return
			}
			defer out.Close()
			io.Copy(out, background)
			settings.BackgroundPath = "/backgrounds/" + backgroundID + backgroundExt
			fmt.Println("Saved background to:", settings.BackgroundPath)
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

	// Serve static files for logos and backgrounds
	router.Static("/logos", "./logos")
	router.Static("/backgrounds", "./backgrounds")

	// Ensure directories exist
	os.MkdirAll("./uploads", os.ModePerm)
	os.MkdirAll("./logos", os.ModePerm)
	os.MkdirAll("./backgrounds", os.ModePerm)

	router.Run(":8080")
}
