package controllers

import (
	"api/models"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
    Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func LoginHost(c *gin.Context) {
    var input LoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    host, err := models.HostAuthenticated(input.Email, input.Password)

    tokenStr, err := utils.GenerateToken(host.ID)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.SetCookie("token", tokenStr, 3600, "/", "localhost", false, true)

    c.JSON(http.StatusOK, gin.H{
        "data": host,
    })
}
