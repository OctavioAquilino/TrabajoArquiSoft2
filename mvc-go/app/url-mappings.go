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
	router.POST("/login", userController.UserLogin)

	//Product Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)
	router.GET("/productRandom/:cantidad", productController.GetRandomProducts)

	//filtros
	router.GET("/productCategory/:idCategory", productController.GetProductsByIdCategory)

	router.POST("/productText", productController.GetProductsByText)

	//Address Mapping
	router.GET("/addressUser/:idUser", addressController.GetAddressByIdUser)

	//Category Mapping
	router.GET("/category", categoryController.GetCategories)

	//OrderDetail Mapping
	router.POST("/orderDetail", orderDetailController.OrderDetailInsert)

	router.GET("/orderDetailOrder/:idOrder", orderDetailController.GetOrderDetailByIdOrder)

	//Order Mapping
	router.POST("/order", orderController.OrderInsert)

	router.GET("/orderUser/:idUser", orderController.GetOrdersByIdUser)

	log.Info("Finishing mappings configurations")
}
