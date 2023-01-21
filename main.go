package main

import (
	"log"
	"products/controller"
	"products/database"
	"products/migrator"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := database.Connection()
	migrator.Migrate(db)
	defer db.Close()

	router := gin.Default()
	router.POST("/products", controller.PostProduct)
	router.GET("/products", controller.GetProducts)
	router.GET("/product/:id", controller.GetProductById)
	router.DELETE("/product/:id", controller.DeleteProduct)
	router.PATCH("/product/:id", controller.UpdateProduct)
	log.Fatal(router.Run(":8080"))

}
