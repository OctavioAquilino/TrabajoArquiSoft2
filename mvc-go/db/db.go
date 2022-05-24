package db

import (
	//product "mvc-go/clients/product"
	addressClient "mvc-go/clients/address"
	productClient "mvc-go/clients/product"
	userClient "mvc-go/clients/user"
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "ejemplo"
	DBUser := "root"
	DBPass := "root"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "127.0.0.1" //127001
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True") //ver puerto

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	productClient.Db = db
	addressClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{}) //se debe agreagar por cada carpeta del model
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Address{})
	log.Info("Finishing Migration Database Tables")
}
