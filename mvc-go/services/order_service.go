package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	orderCliente "mvc-go/clients/order" //DAO
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	//"mvc-go/utils/modeladto"
)

type orderService struct{}

type orderServiceInterface interface {
	//siempre devuelve dto o error
	GetOrderById(id int) (dto.OrderDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) GetOrderById(id int) (dto.OrderDto, e.ApiError) {

	var order model.Order = orderCliente.GetOrderById(id) //objeto de la DB, a traves del DAO
	var orderDto dto.OrderDto

	if order.Id == 0 {
		return orderDto, e.NewBadRequestApiError("order not found")
	}
	orderDto.IdOrder = order.Id
	orderDto.Estado = order.Estado
	orderDto.Fecha = order.Fecha
	orderDto.MontoFinal = order.MontoFinal
	//orderDto.OrderDetail = order.OrderDetail --------- creo que no va
	//orderDto.Usuario = order.Usuario
	orderDto.Usuario.Id = order.Usuario.Id
	orderDto.Usuario.LastName = order.Usuario.LastName
	orderDto.Usuario.Name = order.Usuario.Name
	orderDto.Usuario.Password = order.Usuario.Password
	orderDto.Usuario.UserName = order.Usuario.UserName

	return orderDto, nil
}

func (s *orderService) GetOrders() (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderCliente.GetOrders()
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.IdOrder = order.Id
		orderDto.Estado = order.Estado
		orderDto.Fecha = order.Fecha
		orderDto.MontoFinal = order.MontoFinal
		//orderDto.OrderDetail = order.OrderDetail ---------
		//orderDto.Usuario = order.Usuario -----------------manejar

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderService) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError) {

	var order model.Order

	order.Estado = orderDto.Estado
	order.Fecha = orderDto.Fecha
	order.MontoFinal = orderDto.MontoFinal
	//order.OrderDetail = orderDto.OrderDetail  --------- creo que no va
	// order.Usuario = orderDto.Usuario -----------------manejar

	order = orderCliente.InsertOrder(order)

	order.Id = orderDto.IdOrder

	return orderDto, nil
}
