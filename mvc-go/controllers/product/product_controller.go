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

//filtros por categoria
func GetProductsByIdCategory(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("idCategory"))

	idCategory, _ := strconv.Atoi(c.Param("idCategory")) //se pasa el id de array a int
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByIdCategory(idCategory) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

//Productos por TEXTO
func GetProductsByText(c *gin.Context) {

	/*
		log.Debug("Text looking: " + c.Param("texto"))
		var texto string = c.Param("texto")
	*/
	//var textoDto dto.TextoDto
	//var texto string

	//err1 := c.BindJSON(&textoDto)

	//var texto string = textoDto.Texto
	var text string = c.Param("texto")
	log.Debug("texto:", text)
	/*if err1 != nil {
		log.Error(err1.Error())
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}
	*/
	/*var texto string
	for _, letra := range text {

		//l := strconv.QuoteRune(letra)
		l, _ := strconv.Unquote(strconv.QuoteRune(letra))
		if letra == '+' {
			l = " "
		}
		texto = texto + l
	}*/
	//manejar errores
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByText(text) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

//Productos aleatorios
func GetRandomProducts(c *gin.Context) {
	//	log.Debug("Product id to load: " + c.Param("idCategory"))

	cantidad, _ := strconv.Atoi(c.Param("cantidad")) //se pasa el id de array a int
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetRandomProducts(cantidad) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}
