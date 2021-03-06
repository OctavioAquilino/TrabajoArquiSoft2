package app

import (
	categoryController "mvc-go/controllers/category"
	orderController "mvc-go/controllers/order"
	orderDetailController "mvc-go/controllers/order_detail"
	productController "mvc-go/controllers/product"
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

//para el final dockerizar hacer git pull y docker compose para levantar todo de una
//y testing, md5 para la contraseña no tnenerlo en texto plano, hasheo
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

	router.GET("/productText/:texto", productController.GetProductsByText)

	//Category Mapping
	router.GET("/category", categoryController.GetCategories)

	//OrderDetail Mapping
	router.GET("/orderDetailOrder/:idOrder", orderDetailController.GetOrderDetailByIdOrder)

	//Order Mapping
	router.POST("/order", orderController.OrderInsert)

	router.GET("/orderUser/:idUser", orderController.GetOrdersByIdUser)

	log.Info("Finishing mappings configurations")
}
