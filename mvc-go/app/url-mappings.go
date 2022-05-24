package app

import (
	addressController "mvc-go/controllers/address"
	carritoController "mvc-go/controllers/carrito"
	productController "mvc-go/controllers/product"
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.OPTIONS("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)
	router.POST("/login", userController.UserLogin)

	//Product Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)
	router.POST("/product", productController.ProductInsert)

	//Address Mapping
	router.GET("/address/:id", addressController.GetAddressById)
	router.GET("/address", addressController.GetAddresses)
	router.POST("/address", addressController.AddressInsert)

	//Carrito Mapping
	router.GET("/carrito/:id", carritoController.GetCarritoById)
	router.GET("/carrito", carritoController.GetCarritos)
	router.POST("/carrito", carritoController.CarritoInsert)

	log.Info("Finishing mappings configurations")
}
