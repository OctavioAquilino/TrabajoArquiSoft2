package dto

type AddressDto struct {
	Id     int    `json:"id_address"`
	Calle  string `json:"calle"`
	Numero int    `json:"numero"`
	Barrio string `json:"barrio"`
	Ciudad string `json:"ciudad"`
}

type AddressesDto []AddressDto
