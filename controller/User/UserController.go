package userController

import (
	"go/api_catalogue/model"
	"go/api_catalogue/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context){
	var user model.User

	if err := c.ShouldBind(&user); err!= nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	user.Password = string(hash)
	model.DB.Create(&user)
	c.JSON(http.StatusOK,gin.H{"Message":"Berhasil membuat akun"})
}

func Login (c *gin.Context){
	var user model.User
	var input model.User

	if err := c.ShouldBind(&input); err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	if err := model.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful","token":token})

}
