package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	productCliente "mvc-go/clients/product" //DAO
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type productService struct{}

type productServiceInterface interface {
	//siempre devuelve dto o error
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	GetProductsByIdCategory(idCategory int) (dto.ProductsDto, e.ApiError)
	GetProductsByText(texto string) (dto.ProductsDto, e.ApiError)
	GetRandomProducts(cantidad int) (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id) //objeto de la DB, a traves del DAO
	var productDto dto.ProductDto

	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}
	productDto.Name = product.Name
	productDto.Price = product.Price
	productDto.Id = product.Id
	productDto.Description = product.Description
	productDto.Stock = product.Stock
	productDto.IdCategory = product.IdCategory
	return productDto, nil
}

func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Id = product.Id
		productDto.Description = product.Description
		productDto.Stock = product.Stock
		productDto.IdCategory = product.IdCategory
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

//filtro por category

func (s *productService) GetProductsByIdCategory(idCategory int) (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProductsByIdCategory(idCategory) //objeto de la DB, a traves del DAO
	var productsDto dto.ProductsDto
	/*
		if product.Id == 0 {
			return productDto, e.NewBadRequestApiError("product not found")
		}*/
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Id = product.Id
		productDto.Description = product.Description
		productDto.Stock = product.Stock
		productDto.IdCategory = product.IdCategory
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}

//filtro por texto

func (s *productService) GetProductsByText(texto string) (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProductsByText(texto) //objeto de la DB, a traves del DAO
	var productsDto dto.ProductsDto
	/*
		if product.Id == 0 {
			return productDto, e.NewBadRequestApiError("product not found")
		}*/
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Id = product.Id
		productDto.Description = product.Description
		productDto.Stock = product.Stock
		productDto.IdCategory = product.IdCategory
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}

//random products
func (s *productService) GetRandomProducts(cantidad int) (dto.ProductsDto, e.ApiError) {
	var products model.Products = productCliente.GetRandomProducts(cantidad) //objeto de la DB, a traves del DAO
	var productsDto dto.ProductsDto
	/*
		if product.Id == 0 {
			return productDto, e.NewBadRequestApiError("product not found")
		}*/
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Id = product.Id
		productDto.Description = product.Description
		productDto.Stock = product.Stock
		productDto.IdCategory = product.IdCategory
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}
