package app

import (
	addressController "mvc-go/controllers/address"
	categoryController "mvc-go/controllers/category"
	orderDetailController "mvc-go/controllers/order_detail"
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

	//Category Mapping
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/category", categoryController.GetCategories)
	router.POST("/category", categoryController.CategoryInsert)

	//OrderDetail Mapping
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById)
	router.GET("/orderDetail", orderDetailController.GetOrderDetails)
	router.POST("/orderDetail", orderDetailController.OrderDetailInsert)

	log.Info("Finishing mappings configurations")
}
