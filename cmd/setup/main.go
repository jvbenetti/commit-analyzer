package main

import (
	"log"
	"net/http"
	"os"
	"setup/internal/provider"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Defined router

	r.GET("/metrics/:user/:repo", GetCommits) // Defined endpoint

	err := r.Run() // Run var
	if err != nil {
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
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
