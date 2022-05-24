package model

type OrderDetail struct {
	Id             int `gorm:"primaryKey"`
	PrecioUnitario int `gorm:"type:int(150);not null"`
	Cantidad       int `gorm:"type:int(150);not null"`
	Total          int `gorm:"type:int(150);not null"`
}

type OrderDetails []OrderDetail
