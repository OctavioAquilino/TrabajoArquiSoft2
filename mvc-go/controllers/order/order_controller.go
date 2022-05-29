package orderController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderById(c *gin.Context) {
	log.Debug("Order id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var orderDto dto.OrderDto

	orderDto, err := service.OrderService.GetOrderById(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDto)
}

func GetOrders(c *gin.Context) {
	var ordersDto dto.OrdersDto
	ordersDto, err := service.OrderService.GetOrders() //delega

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}

func OrderInsert(c *gin.Context) {
	var orderDto dto.OrderDto
	err := c.BindJSON(&orderDto) //marshall, convierte de json a orderDto

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDto, er := service.OrderService.InsertOrder(orderDto) //delega, simpre envia dto
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, orderDto)
}
