package model

type Order struct {
	Id          int    `gorm:"primaryKey"`
	MontoFinal  int    `gorm:"type:int(150);not null"`
	Fecha       string `gorm:"type:varchar(350);not null"`
	Estado      string `gorm:"type:varchar(350);not null"`
	Usuario     Users
	OrderDetail OrderDetails
}

type Orders []Order
