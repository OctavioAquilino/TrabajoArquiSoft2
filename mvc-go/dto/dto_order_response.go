package dto

import (
	"time"
)

type OrderResponseDto struct {
	Id         int     `json:"id_order"`
	MontoFinal float32 `json:"monto_final"`
	Fecha      time.Time
}

type OrdersResponseDto []OrderResponseDto
