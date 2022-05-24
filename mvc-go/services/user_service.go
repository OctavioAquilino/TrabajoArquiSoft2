package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	userCliente "mvc-go/clients/user" //DAO
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	"github.com/dgrijalva/jwt-go"
)

type userService struct{}

type userServiceInterface interface {
	//siempre devuelve dto o error
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
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
	userDto.Id = user.Id
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.Name
		userDto.Id = user.Id

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName
	user.Password = userDto.Password

	user = userCliente.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}

var jwtKey = []byte("secret_key")

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError) {
	//var user model.User
	var user model.User = userCliente.GetUserByUserName(loginDto.UserName) //objeto de la DB, a traves del DAO
	//var userDto dto.UserDto
	var tokenDto dto.TokenDto

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("user not found")
	}

	if user.Password == loginDto.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
	}
	return tokenDto, nil
}
