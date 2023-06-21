package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getRoot)

	router.Run("localhost:8080")
}

func getRoot(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "Hello Golang"},
	)
}
