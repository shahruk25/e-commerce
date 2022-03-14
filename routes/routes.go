package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shahruk25/e-commerce/controller"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.Signup())
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.POST("/admin/addProduct", controller.ProductViwerAdmin())
	incomingRoutes.GET("/user/productView", controller.SearchProduct())
	incomingRoutes.GET("/users/search", controller.SearchProductByQuery())
}
