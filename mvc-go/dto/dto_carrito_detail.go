package dto

type CarritoDetailDto struct {
	IdCarritoDetail int        `json:"id_carrito_detail"`
	PrecioUnitario  int        `json:"precio_unitario"`
	Cantidad        int        `json:"cantidad"`
	Producto        ProductDto `json:"Product"`
}

type CarritoDetailsDto []CarritoDetailDto
