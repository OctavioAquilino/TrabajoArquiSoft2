package addressController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAddressById(c *gin.Context) {
	log.Debug("Address id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var addressDto dto.AddressDto

	addressDto, err := service.AddressService.GetAddressById(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, addressDto)
}

func GetAddresses(c *gin.Context) {
	var addressesDto dto.AddressesDto
	addressesDto, err := service.AddressService.GetAddresses() //delega

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, addressesDto)
}

func AddressInsert(c *gin.Context) {
	var addressDto dto.AddressDto
	err := c.BindJSON(&addressDto) //marshall, convierte de json a addressDto

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	addressDto, er := service.AddressService.InsertAddress(addressDto) //delega, simpre envia dto
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, addressDto)
}
