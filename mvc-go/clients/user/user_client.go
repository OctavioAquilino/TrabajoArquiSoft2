package user

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUserByUserName(userName string) model.User {
	var user model.User

	Db.Where("user_name = ?", userName).First(&user)
	log.Debug("User: ", user)

	return user
}
