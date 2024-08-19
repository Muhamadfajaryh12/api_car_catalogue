package categoryController

import (
	"go/api_catalogue/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var category []model.Category

	model.DB.Find(&category)
	c.JSON(http.StatusOK,gin.H{"category":category})
}

func Show(c *gin.Context){
	var category model.Category
	id:= c.Param("id")

	if err := model.DB.First(&category,id).Error; err!= nil{
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound,gin.H{"message":"Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		}
	}

	c.JSON(http.StatusOK,gin.H{"category":category})
}


func Create(c *gin.Context){
	var category model.Category

	if err := c.ShouldBindJSON(&category); err!= nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}

	model.DB.Create(&category)
	c.JSON(http.StatusOK,gin.H{"category":category})

}

func Update(c *gin.Context){
	var category model.Category
	id:= c.Param("id")

	if err := c.ShouldBindJSON(&category); err!= nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}

	if model.DB.Model(&category).Where("id = ?",id).Updates(&category).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message": "Tidak dapat mengupdate category"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":category})
}

func Delete(c *gin.Context) {
	var category model.Category
	id := c.Param("id")

	if err := model.DB.Delete(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category berhasil dihapus"})
}