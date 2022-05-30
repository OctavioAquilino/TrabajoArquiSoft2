package order

//ORM tradutcotr
import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetOrderById(id int) model.Order {
	var order model.Order

	Db.Where("id = ?", id).First(&order) //traduccion y seteo en order
	log.Debug("Order: ", order)

	return order
}

func GetOrders() model.Orders {
	var orders model.Orders
	Db.Find(&orders) //busca y guarda todo en orders

	log.Debug("Orders: ", orders)

	return orders
}

func InsertOrder(order model.Order) model.Order {
	result := Db.Create(&order)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Order Created: ", order.Id)
	return order
}
