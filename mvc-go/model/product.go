package model

type Product struct {
	Id            int     `gorm:"primaryKey"`
	Name          string  `gorm:"type:varchar(350);not null"`
	Price         float32 `gorm:"type:float(150);not null"`
	UniversalCode string  `gorm:"type:varchar(350);not null"`
	Picture       string  `gorm:"type:varchar(350);not null"`
	Categoria     Categories
}

type Products []Product