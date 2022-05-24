package model

type Address struct {
	Id           int    `gorm:"primaryKey"`
	City         string `gorm:"type:varchar(350);not null"`
	Neighborhood string `gorm:"type:varchar(250);not null"`
	Street       string `gorm:"type:varchar(150);not null"`
	Number       int    `gorm:"type:int(150);not null"`
	User         User
}

type Addresses []Address
