package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"verralive/models"
)

type EmailInput struct {
	Email string `json:"email" binding:"required"`
}

func CreateEmail(c *gin.Context) {
	var input EmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e, _ := models.GetEmailByEmail(input.Email)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{"email": "Email already exists."})
		return
	}

	email := models.Email{
		Email:     input.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := email.CreateEmail()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": email})
}
