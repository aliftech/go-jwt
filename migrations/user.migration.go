package migrations

import (
	"github.com/aliftech/go-jwt/models"
	"github.com/aliftech/go-jwt/utils"
)

func UserMigrate() {
	utils.DB.AutoMigrate(&models.User{})
}
