package carritoController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCarritoById(c *gin.Context) {
	log.Debug("Carrito id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var carritoDto dto.CarritoDto

	carritoDto, err := service.CarritoService.GetCarritoById(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, carritoDto)
}

func GetCarritos(c *gin.Context) {
	var carritosDto dto.CarritosDto
	carritosDto, err := service.CarritoService.GetCarritos() //delega

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, carritosDto)
}

func CarritoInsert(c *gin.Context) {
	var carritoDto dto.CarritoDto
	err := c.BindJSON(&carritoDto) //marshall, convierte de json a carritoDto

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	carritoDto, er := service.CarritoService.InsertCarrito(carritoDto) //delega, simpre envia dto
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, carritoDto)
}
