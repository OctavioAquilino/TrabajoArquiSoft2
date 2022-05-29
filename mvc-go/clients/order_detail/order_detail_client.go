package orderDetail

//ORM tradutcotr
import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetOrderDetailById(id int) model.OrderDetail {
	var orderDetail model.OrderDetail

	Db.Where("id = ?", id).First(&orderDetail) //traduccion y seteo en orderDetail
	log.Debug("OrderDetail: ", orderDetail)

	return orderDetail
}

func GetOrderDetails() model.OrderDetails {
	var orderDetails model.OrderDetails
	Db.Find(&orderDetails) //busca y guarda todo en orderDetails

	log.Debug("OrderDetails: ", orderDetails)

	return orderDetails
}

func InsertOrderDetail(orderDetail model.OrderDetail) model.OrderDetail {
	result := Db.Create(&orderDetail)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("OrderDetail Created: ", orderDetail.Id)
	return orderDetail
}
