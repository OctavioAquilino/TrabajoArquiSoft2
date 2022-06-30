package model

type Product struct {
	Id          int     `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(350);not null"`
	Price       float32 `gorm:"type:decimal(60,4);not null"`
	Picture     string  `gorm:"type:varchar(350);not null"`
	Stock       int     `gorm:"type:int(150);not null"`
	IdCategory  int     `gorm:"type:int(150);not null"`
	Description string  `gorm:"type:varchar(350);not null"`
}

type Products []Product
