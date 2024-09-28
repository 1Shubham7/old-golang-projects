package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	router := gin.Default()
	router.GET("hi", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
// you can also write 200 instead of http.StatusOK
			"message": "I Love GeeksforGeeks!",
		})
	})
	router.Run()
}