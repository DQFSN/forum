package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, status := c.Request.BasicAuth()
		errMsg := "The authentication is failed"
		if !status {
			c.AbortWithStatusJSON(401, gin.H{"error": errMsg})
			return
		}
		if !validateUser(username, password) {
			c.AbortWithStatusJSON(401, gin.H{"error": errMsg})
			return
		}
		c.Next()
	}
}

func validateUser(username string, password string) bool {
	// 认证处理逻辑
	log.Printf(username)
	log.Println(password)
	return true
}

func main() {
	api := gin.Default()

	api.Use(AuthMiddleware())
	api.Run()
}