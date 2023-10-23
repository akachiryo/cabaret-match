package controllers

import (
	"net/http"

	"api/models"

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

    tokenStr, host, err := models.GenerateToken(input.Email, input.Password)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.SetCookie("token", tokenStr, 3600, "/", "localhost", false, true)

    c.JSON(http.StatusOK, gin.H{
        "data": host,
    })
}
