package dto

type OrderDetailDto struct {
	IdOrderDeatail int        `json:"id_order_detail"`
	PrecioUnitario int        `json:"precio_unitario"`
	Cantidad       int        `json:"cantidad"`
	Total          int        `json:"total"`
	Detalle        string     `json:"detalle"`
	Producto       ProductDto `json:"Product"`
}

type OrderDetailsDto []OrderDetailDto
