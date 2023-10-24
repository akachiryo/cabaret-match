package controllers

import (
	"api/models"
	"api/utils"
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

	// 追加部分
	token, err := utils.GenerateToken(host.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign up",
		})
		return
	}

		// Cookieにトークンをセット
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"data": hostResponse,
	})
}

func GetHost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"成功": "成功"})
}
