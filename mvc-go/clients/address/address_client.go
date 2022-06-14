package address

//ORM tradutcotr
import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetAddressesByIdUser(idUser int) model.Addresses {
	var addresses model.Addresses

	Db.Where("id_user = ?", idUser).Find(&addresses)

	log.Debug("Addresses: ", addresses)

	return addresses
}
