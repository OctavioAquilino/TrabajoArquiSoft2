package orderDetailController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderDetailById(c *gin.Context) {
	log.Debug("OrderDetail id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var orderDetailDto dto.OrderDetailDto

	orderDetailDto, err := service.OrderDetailService.GetOrderDetailById(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDetailDto)
}

func GetOrderDetails(c *gin.Context) {
	var orderDetailsDto dto.OrderDetailsDto
	orderDetailsDto, err := service.OrderDetailService.GetOrderDetails() //delega

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, orderDetailsDto)
}

func OrderDetailInsert(c *gin.Context) {
	var orderDetailDto dto.OrderDetailDto
	err := c.BindJSON(&orderDetailDto) //marshall, convierte de json a orderDetailDto

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDetailDto, er := service.OrderDetailService.InsertOrderDetail(orderDetailDto) //delega, simpre envia dto
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, orderDetailDto)
}

func GetOrderDetailByIdOrder(c *gin.Context) {
	log.Debug("OrderDetail id to load: " + c.Param("idOrder"))

	idOrder, _ := strconv.Atoi(c.Param("idOrder")) //se pasa el id de array a int
	var orderDetailResDto dto.OrderDetailsResDto

	orderDetailResDto, err := service.OrderDetailService.GetOrderDetailByIdOrder(idOrder) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDetailResDto)
}
