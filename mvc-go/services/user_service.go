package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	userCliente "mvc-go/clients/user" //DAO
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	"crypto/md5"

	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type userService struct{}

type userServiceInterface interface {
	//siempre devuelve dto o error
	GetUserById(id int) (dto.UserDto, e.ApiError)
	LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id) //objeto de la DB, a traves del DAO
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Phone = user.Phone
	userDto.Id = user.Id
	return userDto, nil
}

//LOGIN
var jwtKey = []byte("secret_key")

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError) {

	log.Debug(loginDto)
	var user model.User = userCliente.GetUserByUserName(loginDto.UserName)

	var tokenDto dto.TokenDto

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("user not found")
	}

	var pswMd5 = md5.Sum([]byte(loginDto.Password))
	pswMd5String := hex.EncodeToString(pswMd5[:])

	if pswMd5String == user.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id,
		})
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.IdUser = user.Id

		return tokenDto, nil
	} else {
		return tokenDto, e.NewBadRequestApiError("contraseña incorrecta")
	}

}
