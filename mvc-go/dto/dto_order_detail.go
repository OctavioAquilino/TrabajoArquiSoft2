package dto

type OrderDetailDto struct {
	Id             int     `json:"id"`
	PrecioUnitario float32 `json:"precio_unitario"`
	Cantidad       float32 `json:"cantidad"`
	Total          float32 `json:"total"`
	Nombre         string  `json:"nombre"`
	IdProducto     int     `json:"id_product"`
	IdOrder        int     `json:"id_order"`
}

type OrderDetailsDto []OrderDetailDto
