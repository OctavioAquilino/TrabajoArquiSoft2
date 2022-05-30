package dto

type OrderDetailDto struct {
	Id             int     `json:"id"`
	PrecioUnitario float32 `json:"precio_unitario"`
	Cantidad       int     `json:"cantidad"`
	Total          float32 `json:"total"`
	Detalle        string  `json:"detalle"`
	IdProducto     int     `json:"id_product"`
	IdOrder        int     `json:"id_order"`
}

type OrderDetailsDto []OrderDetailDto
