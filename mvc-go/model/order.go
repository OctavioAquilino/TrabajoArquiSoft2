package model

type Order struct {
	Id          int     `gorm:"primaryKey"`
	MontoFinal  float32 `gorm:"type:float(150);not null"`
	Fecha       string  `gorm:"type:varchar(350);not null"`
	Estado      string  `gorm:"type:varchar(350);not null"`
	Usuario     Users
	OrderDetail OrderDetails
}

type Orders []Order
