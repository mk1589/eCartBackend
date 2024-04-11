package routes

import (
	"github.com/gin-gonic/gin"
)

func UserReoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup")
	incomingRoutes.POST("/users/login")
	incomingRoutes.POST("/users/addproduct")
	incomingRoutes.GET("/users/productivew")
	incomingRoutes.GET("/users/search")

}
