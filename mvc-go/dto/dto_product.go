package dto

type ProductDto struct {
	Name string `json:"name"`
	//UniversalCode string      `json:"universal_code"`
	Picture     string  `json:"picture_url"`
	Price       float32 `json:"base_price"`
	Id          int     `json:"id"`
	IdCategory  int     `json:"id_category"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

type ProductsDto []ProductDto

//cambio de prueba
