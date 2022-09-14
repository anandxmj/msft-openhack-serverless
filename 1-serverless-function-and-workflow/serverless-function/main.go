package main

import (
	"log"
	"net/http"
	"os"
	"serverless-function/product"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8080"
	if portParam, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + portParam
	}
	router := gin.Default()
	router.GET("/api/product", HandleProduct)
	log.Println("serverless-function WebServer Starting on Port ", port)
	log.Fatal(router.Run(port))
}

func HandleProduct(c *gin.Context) {
	result := product.GetById(c.Query("productId"))
	c.String(http.StatusOK, result)
}
