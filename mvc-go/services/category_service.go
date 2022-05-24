package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	categoryCliente "mvc-go/clients/category" //DAO
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type categoryService struct{}

type categoryServiceInterface interface {
	//siempre devuelve dto o error
	GetCategoryById(id int) (dto.CategoryDto, e.ApiError)
	GetCategories() (dto.CategoriesDto, e.ApiError)
	InsertCategory(categoryDto dto.CategoryDto) (dto.CategoryDto, e.ApiError)
}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}

func (s *categoryService) GetCategoryById(id int) (dto.CategoryDto, e.ApiError) {

	var category model.Category = categoryCliente.GetCategoryById(id) //objeto de la DB, a traves del DAO
	var categoryDto dto.CategoryDto

	if category.Id == 0 {
		return categoryDto, e.NewBadRequestApiError("category not found")
	}
	categoryDto.Descripcion = category.Description
	//	categoryDto.IdCategory = category.Id
	categoryDto.Nombre = category.Name
	return categoryDto, nil
}

func (s *categoryService) GetCategories() (dto.CategoriesDto, e.ApiError) {

	var categories model.Categories = categoryCliente.GetCategories()
	var categoriesDto dto.CategoriesDto

	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Descripcion = category.Description
		categoryDto.IdCategory = category.Id
		categoryDto.Nombre = category.Name

		categoriesDto = append(categoriesDto, categoryDto)
	}

	return categoriesDto, nil
}

func (s *categoryService) InsertCategory(categoryDto dto.CategoryDto) (dto.CategoryDto, e.ApiError) {

	var category model.Category

	category.Name = categoryDto.Nombre
	category.Description = categoryDto.Descripcion

	category = categoryCliente.InsertCategory(category)

	categoryDto.IdCategory = category.Id

	return categoryDto, nil
}
