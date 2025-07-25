package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 100 << 20

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

	os.MkdirAll("./uploads", os.ModePerm)
	router.Run(":8080")
}
