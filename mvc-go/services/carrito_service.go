package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	carritoCliente "mvc-go/clients/carrito" //DAO
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type carritoService struct{}

type carritoServiceInterface interface {
	//siempre devuelve dto o error
	GetCarritoById(id int) (dto.CarritoDto, e.ApiError)
	GetCarritos() (dto.CarritosDto, e.ApiError)
	InsertCarrito(carritoDto dto.CarritoDto) (dto.CarritoDto, e.ApiError)
}

var (
	CarritoService carritoServiceInterface
)

func init() {
	CarritoService = &carritoService{}
}

func (s *carritoService) GetCarritoById(id int) (dto.CarritoDto, e.ApiError) {

	var carrito model.Carrito = carritoCliente.GetCarritoById(id) //objeto de la DB, a traves del DAO
	var carritoDto dto.CarritoDto

	if carrito.Id == 0 {
		return carritoDto, e.NewBadRequestApiError("carrito not found")
	}
	//carritoDto.IdCarrito = carrito.Id
	carritoDto.MontoFinal = carrito.MontoFinal
	return carritoDto, nil
}

func (s *carritoService) GetCarritos() (dto.CarritosDto, e.ApiError) {

	var carritos model.Carritos = carritoCliente.GetCarritos()
	var carritosDto dto.CarritosDto

	for _, carrito := range carritos {
		var carritoDto dto.CarritoDto
		carritoDto.MontoFinal = carrito.MontoFinal
		//carritoDto.Productos = carrito.Productos
		//manejar con un for
		carritoDto.IdCarrito = carrito.Id

		carritosDto = append(carritosDto, carritoDto)
	}

	return carritosDto, nil
}

func (s *carritoService) InsertCarrito(carritoDto dto.CarritoDto) (dto.CarritoDto, e.ApiError) {

	var carrito model.Carrito

	carrito.MontoFinal = carritoDto.MontoFinal
	//carrito.Productos = carritoDto.Productos

	carrito = carritoCliente.InsertCarrito(carrito)

	carritoDto.IdCarrito = carrito.Id

	return carritoDto, nil
}
