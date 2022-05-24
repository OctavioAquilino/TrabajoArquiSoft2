package dto

//deberia ser mucho mas compleja, con mas datos
//No necesariamente un dto por objeto puede haber mas
type UserDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Id       int    `json:"id"`
}

type UsersDto []UserDto
