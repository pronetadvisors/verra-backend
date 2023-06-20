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
		Email:      input.Email,
		TimeViewed: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	_, err := email.CreateEmail()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": email})
}

func GetEmails(c *gin.Context) {
	var email models.Email

	emails, err := email.GetEmails()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"emails": emails})
}

func IncrementTime(c *gin.Context) {
	var email models.Email

	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e, err := models.GetEmailByEmail(email.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e.TimeViewed += 5
	_, err = e.UpdateEmail()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": email})
}
