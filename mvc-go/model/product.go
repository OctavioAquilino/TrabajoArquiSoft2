package model

type Product struct {
	Id    int     `gorm:"primaryKey"`
	Name  string  `gorm:"type:varchar(350);not null"`
	Price float32 `gorm:"type:decimal(60,4);not null"`
	//UniversalCode string  `gorm:"type:varchar(350);not null"`
	Picture     string `gorm:"type:varchar(350);not null"`
	IdCategory  int    `gorm:"type:integer;not null"` //chequear lo de integer
	Description string `gorm:"type:varchar(350);not null"`
}

type Products []Product
