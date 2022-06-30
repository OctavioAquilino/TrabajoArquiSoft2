package order

//ORM tradutcotr
import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertOrder(order model.Order) model.Order {
	result := Db.Create(&order)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Order Created: ", order.Id)
	return order
}

//busuqueda por idUser
func GetOrdersByIdUser(idUser int) model.Orders {
	var orders model.Orders

	log.Debug("idUser: ", idUser)
	Db.Where("id_user = ?", idUser).Find(&orders)
	log.Debug("Order: ", orders)

	return orders
}

func UpdateMontoFinal(monto float32, idOrder int) float32 {
	result := Db.Model(&model.Order{}).Where("id = ?", idOrder).Update("monto_final", monto)

	if result.Error != nil {
		log.Error("Order no encontrada")
	}
	return monto
}
