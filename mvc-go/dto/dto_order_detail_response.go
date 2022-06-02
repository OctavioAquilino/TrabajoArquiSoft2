package dto

type OrderDetailResDto struct {
	Id             int     `json:"id"`
	PrecioUnitario float32 `json:"precio_unitario"`
	Cantidad       float32 `json:"cantidad"`
	Total          float32 `json:"total"`
	Detalle        string  `json:"detalle"`
	IdProducto     int     `json:"id_product"`
}

type OrderDetailsResDto []OrderDetailResDto
