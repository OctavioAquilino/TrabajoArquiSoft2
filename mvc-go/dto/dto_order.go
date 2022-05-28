package dto

type OrderDto struct {
	IdOrder     int            `json:"id_order"`
	MontoFinal  float32        `json:"monto_final"`
	Fecha       string         `json:"fecha"`
	Estado      string         `json:"estado"`
	Usuario     UserDto        `json:"Usuario"`
	OrderDetail OrderDetailDto `json:"OrderDetalle"`
}

type OrdersDto []OrderDto
