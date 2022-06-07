package app

import (
	addressController "mvc-go/controllers/address"
	categoryController "mvc-go/controllers/category"
	orderController "mvc-go/controllers/order"
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
	router.POST("/login", userController.UserLogin)
	//router.OPTIONS("/login", userController.UserLogin)

	//Product Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)

	//filtros
	router.GET("/productCategory/:idCategory", productController.GetProductsByIdCategory)

	router.GET("/productText", productController.GetProductsByText)

	//Address Mapping
	router.GET("/address/:id", addressController.GetAddressById)
	router.GET("/address", addressController.GetAddresses)

	//Category Mapping
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/category", categoryController.GetCategories)

	//OrderDetail Mapping
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById)
	router.GET("/orderDetail", orderDetailController.GetOrderDetails)
	router.POST("/orderDetail", orderDetailController.OrderDetailInsert)

	router.GET("/orderDetailOrder/:idOrder", orderDetailController.GetOrderDetailByIdOrder)

	//Order Mapping
	router.GET("/order/:id", orderController.GetOrderById)
	router.GET("/order", orderController.GetOrders)
	router.POST("/order", orderController.OrderInsert)

	router.GET("/orderUser/:idUser", orderController.GetOrdersByIdUser)

	log.Info("Finishing mappings configurations")
}
