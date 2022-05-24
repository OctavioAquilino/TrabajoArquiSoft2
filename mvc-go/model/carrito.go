package model

type Carrito struct {
	Id            int `gorm:"primaryKey"`
	MontoFinal    int `gorm:"type:int(150);not null"`
	Productos     Products
	CarritoDetail CarritoDetails
}

type Carritos []Carrito
