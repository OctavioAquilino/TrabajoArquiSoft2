package dto

type CarritoDto struct {
	IdCarrito  int         `json:"id_carrito"`
	MontoFinal int         `json:"monto_final"`
	Productos  ProductsDto `json:"Productos"`
}

type CarritosDto []CarritoDto
