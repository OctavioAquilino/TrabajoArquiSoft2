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

func GetOrderDetailByIdOrder(c *gin.Context) {
	log.Debug("OrderDetail id to load: " + c.Param("idOrder"))

	idOrder, _ := strconv.Atoi(c.Param("idOrder")) //se pasa el id de array a int
	var orderDetailResDto dto.OrderDetailsDto

	orderDetailResDto, err := service.OrderDetailService.GetOrderDetailByIdOrder(idOrder) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDetailResDto)
}
