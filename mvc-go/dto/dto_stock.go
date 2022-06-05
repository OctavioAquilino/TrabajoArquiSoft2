package dto

type StockDto struct {
	IdProduct int `json:"id_product"`
	Stock     int `json:"stock"`
}

type StocksDto []StockDto
