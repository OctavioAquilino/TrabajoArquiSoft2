package dto

type CategoryDto struct {
	IdCategory int    `json:"id_category"`
	Nombre     string `json:"nombre"`
}

type CategoriesDto []CategoryDto
