package addressController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

func GetAddressByIdUser(c *gin.Context) {

	idUser, _ := strconv.Atoi(c.Param("idUser"))
	var addressDto dto.AddressesDto

	addressDto, err := service.AddressService.GetAddressesByIdUser(idUser)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, addressDto)
}
