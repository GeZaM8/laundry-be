package controllers

import (
	"net/http"

	"github.com/GeZaM8/laundry-be/auth"
	"github.com/GeZaM8/laundry-be/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (ac *AuthController) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, model.Response{
            Status:  false,
            Message: "Terjadi Kesalahan",
        })
        return
    }

    var user model.User
    if err := ac.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, model.Response{
            Status:  false,
            Message: "Email atau password salah!",
        })
        return
    }

    if user.Password != req.Password {
        c.JSON(http.StatusUnauthorized, model.Response{
            Status:  false,
            Message: "Email atau password salah",
        })
        return
    }

    token, _ := auth.GenerateToken(uint(user.ID))

    c.JSON(http.StatusOK, model.Response{
        Status:  true,
        Data: gin.H{
            "token": token,
            "user": gin.H{
                "id":    user.ID,
                "email": user.Email,
            },
        },
    })
}

