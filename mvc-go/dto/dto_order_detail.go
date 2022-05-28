package dto

type OrderDetailDto struct {
	IdOrderDeatail int        `json:"id_order_detail"`
	PrecioUnitario float32    `json:"precio_unitario"`
	Cantidad       int        `json:"cantidad"`
	Total          float32    `json:"total"`
	Detalle        string     `json:"detalle"`
	Producto       ProductDto `json:"Product"`
}

type OrderDetailsDto []OrderDetailDto
