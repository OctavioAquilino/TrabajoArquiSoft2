package dto

type AddressDto struct {
	Id           int     `json:"id_address"`
	Street       string  `json:"calle"`
	Number       int     `json:"numero"`
	Neighborhood string  `json:"barrio"`
	City         string  `json:"ciudad"`
	Users        UserDto `json:"Usuario"`
}

type AddressesDto []AddressDto
