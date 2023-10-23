package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequestParams struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterHost(c *gin.Context) {
	var input RegisterRequestParams

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	host := models.Host{Email: input.Email, Password: input.Password}

	hostResponse, err := host.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hostResponse,
	})
}
