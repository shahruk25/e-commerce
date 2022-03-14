package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shahruk25/e-commerce/controller"
	"github.com/shahruk25/e-commerce/middleware"
	"github.com/shahruk25/e-commerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controller.NewApp()
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	routes.UserRoutes(ginEngine)
	ginEngine.Use(middleware.Authentication())

	ginEngine.GET("/addToCart", app.addToCart())
	ginEngine.GET("/removeItem", app.removeItem())
	ginEngine.GET("/cartCheckout", app.cartCheckout())
	ginEngine.GET("/instantBuy", app.instantBuy())

	log.Fatal(ginEngine.Run(":" + port))
}
