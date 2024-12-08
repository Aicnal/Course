package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"courses": []string{"数学", "物理", "化学"},
	})
}
