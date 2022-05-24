package dto

type CategoryDto struct {
	IdCtegory   int    `json:"id_category"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type CategoriesDto []CategoryDto
