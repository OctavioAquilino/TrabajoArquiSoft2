package orderDetail

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertOrderDetail(orderDetail model.OrderDetail) model.OrderDetail {
	result := Db.Create(&orderDetail)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("OrderDetail Created: ", orderDetail.Id)
	result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.IdProduct).Update("stock", gorm.Expr("stock - ? ", orderDetail.Cantidad))

	if result1.Error != nil {
		log.Error("Pdto no encontrado")
	}

	return orderDetail
}

func GetOrderDetailByIdOrder(idOrder int) model.OrderDetails {
	var ordersDetail model.OrderDetails

	Db.Where("id_order = ?", idOrder).Find(&ordersDetail)
	log.Debug("OrderDetail: ", ordersDetail)

	return ordersDetail
}

func InsertOrdersDetail(ordersDetail model.OrderDetails) model.OrderDetails {

	for _, orderDetail := range ordersDetail {

		result := Db.Create(&orderDetail)

		if result.Error != nil {
			log.Error("Error al crear")
		}
		log.Debug("OrderDetail Created: ", orderDetail.Id)
		result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.IdProduct).Update("stock", gorm.Expr("stock - ? ", orderDetail.Cantidad))

		if result1.Error != nil {
			log.Error("Pdto no encontrado")
		}
	}

	return ordersDetail
}
