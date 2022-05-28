package model

type OrderDetail struct {
	Id             int    `gorm:"primaryKey"`
	PrecioUnitario float32    `gorm:"type:float32(150);not null"`
	Cantidad       int    `gorm:"type:int(150);not null"`
	Total          float32    `gorm:"type:float(150);not null"`
	Detalle        string `gorm:"type:varchar(250);not null"`
	Producto       Products
}

type OrderDetails []OrderDetail
