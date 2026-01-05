package main

import (
	"log"
	"net/http"
	"os"
	"setup/internal/provider"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/metrics/:user/:repo", GetCommits)

	err := r.Run()
	if err != nil {
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
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
