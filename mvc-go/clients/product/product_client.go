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

	Db.Where("id = ?", id).First(&product)
	log.Debug("Product: ", product)

	return product
}

func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products)

	log.Debug("Products: ", products)

	return products
}

func GetProductsByIdCategory(idCategory int) model.Products {
	var products model.Products

	Db.Where("id_category = ?", idCategory).Find(&products)
	log.Debug("Product: ", products)

	return products
}

func GetProductsByText(texto string) model.Products {
	var products model.Products

	Db.Where("name LIKE ?", "%"+texto+"%").Find(&products)
	log.Debug("Product: ", products)

	return products
}

func GetRandomProducts(cantidad int) model.Products {
	var products model.Products
	Db.Order("RAND()").Limit(cantidad).Find(&products)

	log.Debug("Products: ", products)

	return products

}
