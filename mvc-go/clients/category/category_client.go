package category

//ORM tradutcotr
import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCategoryById(id int) model.Category {
	var category model.Category

	Db.Where("id = ?", id).First(&category) //traduccion y seteo en category
	log.Debug("Category: ", category)

	return category
}

func GetCategories() model.Categories {
	var categories model.Categories
	Db.Find(&categories) //busca y guarda todo en categories

	log.Debug("Categories: ", categories)

	return categories
}
