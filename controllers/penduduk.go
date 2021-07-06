package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kholid74/golang/jwt-go/database"
	"github.com/kholid74/golang/jwt-go/models"
	"gorm.io/gorm"
)

func Penduduk(c *gin.Context) {
	var penduduk []models.Penduduk

	result := database.GlobalDB.Find(&penduduk)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "data not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get data penduduk",
		})
		c.Abort()
		return
	}

	c.JSON(200, penduduk)

	return
}
