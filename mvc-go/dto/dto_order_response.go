package dto

type OrderResponseDto struct {
	Id         int     `json:"id_order"`
	MontoFinal float32 `json:"monto_final"`
	Fecha      string  `json:"fecha"`
}

type OrdersResponseDto []OrderResponseDto
