package views

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
}

func Error(c *gin.Context, data any) {
	c.JSON(400, gin.H{
		"status": "error",
		"data":   data,
	})
}
