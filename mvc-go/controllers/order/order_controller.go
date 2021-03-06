package orderController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func OrderInsert(c *gin.Context) {
	var orderDto dto.OrderDto
	err := c.BindJSON(&orderDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDto, er := service.OrderService.InsertOrder(orderDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, orderDto)
}

//Buscar orden por token

func GetOrdersByIdUser(c *gin.Context) {

	var ordersDto dto.OrdersDto
	token := c.Param("idUser")
	ordersDto, err := service.OrderService.GetOrdersByIdUser(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersDto)
}
