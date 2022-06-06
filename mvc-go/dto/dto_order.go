package dto

import (
	"time"
)

type OrderDto struct {
	Id         int       `json:"id_order"`
	MontoFinal float32   `json:"monto_final"`
	Fecha      time.Time `json:"fecha"`
	//Estado     string  `json:"estado"`
	IdUsuario    int             `json:"id_user"`
	OrdersDetail OrderDetailsDto `json:"OrdersDetalle"`
}

type OrdersDto []OrderDto
