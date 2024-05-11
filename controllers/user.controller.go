package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/aliftech/go-jwt/models"
	"github.com/aliftech/go-jwt/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get email and password of req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to read request body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to hash password",
		})

		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := utils.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to create new user",
		})

		return
	}

	// Show response
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "New user have been created",
	})
}

func Login(c *gin.Context) {
	// Get email and password od req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to read request body",
		})

		return
	}

	// lookup the requested user
	var user models.User
	utils.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid email or password",
		})

		return
	}

	// Compare sent in pass with saved user hased pass
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid email or password",
		})

		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID,
		"expired": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 days
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("jwt.secret")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to create token",
		})

		return
	}

	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Login Success.",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"error":    false,
		"is_login": true,
		"user":     user,
	})
}
