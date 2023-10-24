package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Shiirookami/restapi_gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{("products"): products})
}
func Show(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")

	if err := models.DB.First(&Product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": Product})
}
func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Product Updated"})
}
func Delete(c *gin.Context) {
	var product models.Product

	var number struct {
		Id json.Number
	}
	if err := c.ShouldBindJSON(&number); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := number.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Product Has Deleted"})
}
