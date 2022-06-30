package category

//ORM tradutcotr
import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCategories() model.Categories {
	var categories model.Categories
	Db.Find(&categories)

	log.Debug("Categories: ", categories)

	return categories
}
