package db

import (
	//product "mvc-go/clients/product"

	categoryClient "mvc-go/clients/category"
	orderClient "mvc-go/clients/order"
	orderDetailClient "mvc-go/clients/order_detail"
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
	DBName := "trabajoarqui"
	DBUser := "root"
	DBPass := ""
	DBHost := "database"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True") //ver puerto

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	userClient.Db = db
	productClient.Db = db
	categoryClient.Db = db
	orderDetailClient.Db = db
	orderClient.Db = db
}

func StartDbEngine() {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.OrderDetail{})
	db.AutoMigrate(&model.Order{})
	log.Info("Finishing Migration Database Tables")
}
