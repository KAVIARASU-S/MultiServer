package routes

import (
	"client/controllers"


	"github.com/gin-gonic/gin"
)


func Routes(router *gin.Engine) {
	router.GET("/sample",controllers.DisplayCompany)
	router.POST("/sample",controllers.InsertCompany)
}
