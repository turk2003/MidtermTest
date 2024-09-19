package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) Login(ctx *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if credentials.Username == "admin" && credentials.Password == "secret" {
		// สร้าง JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": credentials.Username,
			"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token หมดอายุใน 1 ชั่วโมง
		})

		// เซ็นด้วย secret key
		tokenString, err := token.SignedString([]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiYXVkIjoiU3VwZXIgQWRtaW4iLCJleHAiOjE"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		// ตั้งค่า Cookie
		ctx.SetCookie("token", tokenString, 3600, "/", "", false, true)

		ctx.JSON(http.StatusOK, gin.H{"message": "Login Succeed",
			"token": tokenString})
		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
