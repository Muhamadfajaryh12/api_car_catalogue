package productController

import (
	"fmt"
	"go/api_catalogue/model"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var products []model.Product

	model.DB.Preload("Category").Find(&products)
	c.JSON(http.StatusOK,gin.H{"products":products})
}

func Show(c *gin.Context){
	var product model.Product
	id:= c.Param("id")

	if err := model.DB.First(&product,id).Error; err!= nil{
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound,gin.H{"message":"Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		}
	}

	c.JSON(http.StatusOK,gin.H{"product":product})
}


func Create(c *gin.Context) {
    var product model.Product

    file, header, err := c.Request.FormFile("gambar_product")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file: " + err.Error()})
        return
    }
    defer file.Close()

    fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(header.Filename))
    filePath := filepath.Join("uploads", fileName)

    if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory: " + err.Error()})
        return
    }

    dst, err := os.Create(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file: " + err.Error()})
        return
    }
    defer dst.Close()

    if _, err := io.Copy(dst, file); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + err.Error()})
        return
    }

    if err := c.ShouldBind(&product); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON: " + err.Error()})
        return
    }

    product.GambarProduct = filePath

    if err := model.DB.Create(&product).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": product})
        return
    }

    // Memuat relasi Category untuk produk yang baru dibuat
	var savedProduct model.Product
	if err := model.DB.Preload("Category").First(&savedProduct, product.ID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": savedProduct})
		return
	}
	

    c.JSON(http.StatusOK, gin.H{"product": savedProduct})
}

func Update(c *gin.Context){
	var product model.Product
	id:= c.Param("id")

	if err := c.ShouldBindJSON(&product); err!= nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}

	if model.DB.Model(&product).Where("id = ?",id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message": "Tidak dapat mengupdate Product"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":product})
}

func Delete(c *gin.Context) {
	var product model.Product
	id := c.Param("id")

	if err := model.DB.Delete(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product berhasil dihapus"})
}