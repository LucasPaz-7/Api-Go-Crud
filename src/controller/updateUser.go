package controller

import "github.com/gin-gonic/gin"

func UpdateUser(p *gin.Context) {
	
	p.JSON(200, gin.H{
		"message": "UpdateUser",
	})
}