package model

type OrderDetail struct {
	Id             int    `gorm:"primaryKey"`
	PrecioUnitario int    `gorm:"type:int(150);not null"`
	Cantidad       int    `gorm:"type:int(150);not null"`
	Total          int    `gorm:"type:int(150);not null"`
	Detalle        string `gorm:"type:varchar(250);not null"`
	Producto       Products
}

type OrderDetails []OrderDetail
