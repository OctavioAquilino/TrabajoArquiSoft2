package product

//ORM tradutcotr
import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProductById(id int) model.Product {
	var product model.Product

	Db.Where("id = ?", id).First(&product) //traduccion y seteo en product
	log.Debug("Product: ", product)

	return product
}

func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products) //busca y guarda todo en products

	log.Debug("Products: ", products)

	return products
}

func InsertProduct(product model.Product) model.Product {
	result := Db.Create(&product)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Product Created: ", product.Id)
	return product
}

func GetProductByProductName(productName string) model.Product {
	var product model.Product

	Db.Where("product_name = ?", productName).First(&product) //traduccion y seteo en product
	//	log.Debug("Product: ", product)

	return product
}