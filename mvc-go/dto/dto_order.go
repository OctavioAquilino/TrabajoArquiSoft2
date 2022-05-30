package dto

type OrderDto struct {
	Id         int     `json:"id_order"`
	MontoFinal float32 `json:"monto_final"`
	Fecha      string  `json:"fecha"`
	//Estado     string  `json:"estado"`
	IdUsuario int `json:"id_user"`
	//OrderDetail OrderDetailDto `json:"OrderDetalle"`
}

type OrdersDto []OrderDto
