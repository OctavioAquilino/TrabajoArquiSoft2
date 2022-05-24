package model

type CarritoDetail struct {
	Id             int `gorm:"primaryKey"`
	PrecioUnitario int `gorm:"type:int(150);not null"`
}

type CarritoDetails []CarritoDetail
