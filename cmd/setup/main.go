package main

import (
	"log"
	"net/http"
	"os"
	"setup/internal/provider"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()                        // Defined router
	r.GET("/metrics/:user/:repo", GetCommits) // Defined endpoint
	return r
}

func main() {
	r := SetupRouter()
	port := os.Getenv("PORT") // Var to port with conditional
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port) // Logs before the port

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func GetCommits(c *gin.Context) {
	user := c.Param("user")
	repo := c.Param("repo")
	commits, err := provider.SearchCommits(user, repo)
	if err != nil { // Return error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, commits)
}
