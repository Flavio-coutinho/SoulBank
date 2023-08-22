package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Auth 
	router.GET("/login", users.login)
}