package dto

type AddressDto struct {
	Id           int    `json:"id_address"`
	Street       string `json:"calle"`
	Number       int    `json:"numero"`
	Neighborhood string `json:"barrio"`
	City         string `json:"ciudad"`
	IdUser       int    `json:"id_user"`
}

type AddressesDto []AddressDto

//..
