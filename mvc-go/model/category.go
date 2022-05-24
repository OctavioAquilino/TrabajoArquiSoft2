package model

type Category struct {
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"varchar(150);not null"`
	Description string `gorm:"varchar(350);not null"`
}

type Categories []Category
