package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aliftech/go-jwt/models"
	"github.com/aliftech/go-jwt/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateAuth(c *gin.Context) {
	// Get the cookie of request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Unauthorized.",
		})
	}

	// Decode or validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("jwt.secret")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the expiration
		if float64(time.Now().Unix()) > claims["expired"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Unauthorized.",
			})
		}

		// Find user with token sub
		var user models.User
		utils.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Unauthorized.",
			})
		}

		// Attached to request
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Unauthorized.",
		})
	}

}
