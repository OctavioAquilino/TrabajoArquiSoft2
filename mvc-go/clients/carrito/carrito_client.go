package carrito

//ORM tradutcotr
import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCarritoById(id int) model.Carrito {
	var carrito model.Carrito

	Db.Where("id = ?", id).First(&carrito) //traduccion y seteo en carrito
	log.Debug("Carrito: ", carrito)

	return carrito
}

func GetCarritos() model.Carritos {
	var carritos model.Carritos
	Db.Find(&carritos) //busca y guarda todo en carritos

	log.Debug("Carritos: ", carritos)

	return carritos
}

func InsertCarrito(carrito model.Carrito) model.Carrito {
	result := Db.Create(&carrito)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Carrito Created: ", carrito.Id)
	return carrito
}

func GetCarritoByCarritoName(carritoName string) model.Carrito {
	var carrito model.Carrito

	Db.Where("carrito_name = ?", carritoName).First(&carrito) //traduccion y seteo en carrito
	//	log.Debug("Carrito: ", carrito)

	return carrito
}
