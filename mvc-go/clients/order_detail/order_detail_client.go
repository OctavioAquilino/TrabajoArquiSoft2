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
	//result1 := Db.UpdateColumn()
	log.Debug("OrderDetail Created: ", orderDetail.Id)
	result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.IdProduct).Update("stock", gorm.Expr("stock - ? ", orderDetail.Cantidad))

	if result1.Error != nil {
		//TODO Manage Errors
		log.Error("Pdto no encontrado")
	}

	return orderDetail
}

func GetOrderDetailByIdOrder(idOrder int) model.OrderDetails {
	var ordersDetail model.OrderDetails

	Db.Where("id_order = ?", idOrder).Find(&ordersDetail) //traduccion y seteo en orderDetail
	log.Debug("OrderDetail: ", ordersDetail)

	return ordersDetail
}

func InsertOrdersDetail(ordersDetail model.OrderDetails) model.OrderDetails {

	for _, orderDetail := range ordersDetail {

		result := Db.Create(&orderDetail)

		if result.Error != nil {
			//TODO Manage Errors
			log.Error("")
		}
		//result1 := Db.UpdateColumn()
		log.Debug("OrderDetail Created: ", orderDetail.Id)
		result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.IdProduct).Update("stock", gorm.Expr("stock - ? ", orderDetail.Cantidad))

		if result1.Error != nil {
			//TODO Manage Errors
			log.Error("Pdto no encontrado")
		}

	}
	for _, orderDetail := range ordersDetail {
		log.Debug("Id de la order detail: ", orderDetail.Id)
	}
	return ordersDetail
}
