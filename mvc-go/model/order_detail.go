package model

type OrderDetail struct {
	Id             int     `gorm:"primaryKey"`
	PrecioUnitario float32 `gorm:"type:decimal(60,4);not null"`
	Cantidad       float32 `gorm:"type:decimal(60,4);not null"`
	Total          float32 `gorm:"type:decimal(60,4);not null"`
	Nombre         string  `gorm:"type:varchar(550);not null"`
	IdProduct      int     `gorm:"type:int(150);not null"`
	IdOrder        int     `gorm:"type:int(150);not null"`
}

type OrderDetails []OrderDetail
