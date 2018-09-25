package controller

import (
	// . "fmt"
	"github.com/gin-gonic/gin"
	"myteam/model"
	"myteam/types"
	jwt "github.com/dgrijalva/jwt-go"
)

func AuthSignin(c *gin.Context) {
	var data types.User
	c.BindJSON(&data)

	user, err := model.GetUserAuth(data.Username, data.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Auth failed",
			"error": err.Error(),
		})
		return
	}

	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"user": user,
	})
	token, err := sign.SignedString([]byte("rahasiadong"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Signing in failed",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Signing in Success",
		"token": token,
	})
}