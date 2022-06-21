package userController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var userDto dto.UserDto

	userDto, err := service.UserService.GetUserById(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

//PARTE DE AUTENTICACION JWT

var jwtKey = []byte("secret_key")

func UserLogin(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto)
	log.Debug(loginDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenDto, er := service.UserService.LoginUser(loginDto) //delega, simpre envia dto

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	/*if tokenDto.Token == "" {
		var errorDto dto.ErrorDto
		errorDto.Error = "invalid login"
		errorDto.Code = "401"
		c.JSON(http.StatusUnauthorized, errorDto)
		return
	}*/
	tkn, err := jwt.Parse(tokenDto.Token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, "Invalid Signature")
			return
		}
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	} ///esto deberia estar en el service

	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, "Invalid token")
		return
	} // ver q onda

	c.JSON(http.StatusCreated, tokenDto)

}
