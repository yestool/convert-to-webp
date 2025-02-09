package main

import (
	"github.com/h2non/bimg"
	"io"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 设置文件上传的路由
	r.POST("/convert-to-webp", func(c *gin.Context) {
		// 从请求中获取上传的文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}

		// 打开上传的文件
		uploadedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer uploadedFile.Close()

		// 读取文件内容到字节数组
		fileBytes, err := io.ReadAll(uploadedFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}

		// 使用bimg将图片转换为WebP格式
		convertedImage, err := bimg.NewImage(fileBytes).Convert(bimg.WEBP)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert image to WebP"})
			return
		}

		// 设置响应头，返回WebP格式的图片
		c.Header("Content-Type", "image/webp")
		c.Header("Content-Disposition", "attachment; filename=converted.webp")
		c.Data(http.StatusOK, "image/webp", convertedImage)
	})

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
