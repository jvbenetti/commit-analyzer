package setup

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"setup/internal/provider"
)

func GetCommits(c *gin.Context) {
	user := c.Param("user")
	repo := c.Param("repo")
	commits, err := provider.SearchCommits(user, repo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

}
