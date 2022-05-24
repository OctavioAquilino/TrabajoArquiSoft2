package productController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var productDto dto.ProductDto

	productDto, err := service.ProductService.GetProductById(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts() //delega

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func ProductInsert(c *gin.Context) {
	var productDto dto.ProductDto
	err := c.BindJSON(&productDto) //marshall, convierte de json a productDto

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productDto, er := service.ProductService.InsertProduct(productDto) //delega, simpre envia dto
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, productDto)
}
