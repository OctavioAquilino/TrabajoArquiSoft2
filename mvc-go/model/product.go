package model

type Product struct {
	Id    int     `gorm:"primaryKey"`
	Name  string  `gorm:"type:varchar(350);not null"`
	Price float32 `gorm:"type:float;not null"`
	//agregar mas
}

type Products []Product
