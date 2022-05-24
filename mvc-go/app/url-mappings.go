package app

import (
	productController "mvc-go/controllers/product"
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)
	router.POST("/login", userController.UserLogin)

	//Product Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)
	router.POST("/product", productController.ProductInsert)

	log.Info("Finishing mappings configurations")
}
