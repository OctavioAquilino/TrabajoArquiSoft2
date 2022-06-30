package productController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var productDto dto.ProductDto

	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func GetProductsByIdCategory(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("idCategory"))

	idCategory, _ := strconv.Atoi(c.Param("idCategory"))
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByIdCategory(idCategory)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

//Productos por TEXTO
func GetProductsByText(c *gin.Context) {

	var text string = c.Param("texto")
	log.Debug("texto:", text)
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByText(text)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

//Productos aleatorios
func GetRandomProducts(c *gin.Context) {
	cantidad, _ := strconv.Atoi(c.Param("cantidad"))
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetRandomProducts(cantidad)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}
