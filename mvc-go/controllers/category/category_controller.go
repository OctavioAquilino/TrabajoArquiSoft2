package categoryController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	_ "strconv"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

func GetCategories(c *gin.Context) {
	var categoriesDto dto.CategoriesDto
	categoriesDto, err := service.CategoryService.GetCategories() //delega

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, categoriesDto)
}
