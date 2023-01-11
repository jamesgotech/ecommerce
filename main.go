package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gotech/ecommerce/controllers"
	"github.com/gotech/ecommerce/database"
	"github.com/gotech/ecommerce/middleware"
	"github.com/gotech/ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	//other routes apart from user
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.Instantbuy())

	log.Fatal(router.Run(":" + port))
}
