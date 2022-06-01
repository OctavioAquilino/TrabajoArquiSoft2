package model

type Order struct {
	Id         int     `gorm:"primaryKey"`
	MontoFinal float32 `gorm:"type:decimal(60,4);not null"`
	Fecha      string  `gorm:"type:varchar(350);not null"`
	//Estado     string  `gorm:"type:varchar(350);not null"`
	IdUser int `gorm:"type:integer;not null"`
}

type Orders []Order
