package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	_ "fmt"
	orderCliente "mvc-go/clients/order" //DAO
	orderDetailCliente "mvc-go/clients/order_detail"
	productCliente "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
	//"mvc-go/utils/modeladto"
)

type orderService struct{}

type orderServiceInterface interface {
	//siempre devuelve dto o error
	GetOrderById(id int) (dto.OrderDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
	GetOrdersByIdUser(idUser int) (dto.OrdersDto, e.ApiError)
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
	orderDto.Id = order.Id
	orderDto.Fecha = order.Fecha
	orderDto.MontoFinal = order.MontoFinal
	orderDto.IdUsuario = order.IdUser

	return orderDto, nil
}

func (s *orderService) GetOrders() (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderCliente.GetOrders()
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id = order.Id

		orderDto.Fecha = order.Fecha
		orderDto.MontoFinal = order.MontoFinal
		orderDto.IdUsuario = order.IdUser

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderService) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError) {

	var order model.Order

	order.Fecha = time.Now()

	order.IdUser = orderDto.IdUsuario

	order = orderCliente.InsertOrder(order)

	orderDto.Id = order.Id

	var ordersDetail model.OrderDetails
	var montofinal float32
	for _, orderDetailDto := range orderDto.OrdersDetail {
		var orderDetail model.OrderDetail

		orderDetail.IdProduct = orderDetailDto.IdProducto

		var product model.Product = productCliente.GetProductById(orderDetail.IdProduct)
		orderDetail.Nombre = product.Name
		orderDetail.PrecioUnitario = product.Price

		orderDetail.Cantidad = orderDetailDto.Cantidad
		orderDetail.Total = orderDetail.PrecioUnitario * orderDetail.Cantidad

		montofinal += orderDetail.Total

		orderDetail.IdOrder = orderDto.Id

		ordersDetail = append(ordersDetail, orderDetail)
	}

	order.MontoFinal = orderCliente.UpdateMontoFinal(montofinal, order.Id)

	ordersDetail = orderDetailCliente.InsertOrdersDetail(ordersDetail)
	// cargado de la order response

	var orderResponseDto dto.OrderDto

	orderResponseDto.Fecha = order.Fecha
	orderResponseDto.Id = order.Id
	orderResponseDto.IdUsuario = order.IdUser
	orderResponseDto.MontoFinal = order.MontoFinal

	for _, orderDetail := range ordersDetail {
		var orderDetailDto dto.OrderDetailDto

		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Cantidad = orderDetail.Cantidad
		orderDetailDto.IdProducto = orderDetail.IdProduct
		orderDetailDto.PrecioUnitario = orderDetail.PrecioUnitario
		orderDetailDto.Total = orderDetail.Total
		orderDetailDto.IdOrder = orderDetail.IdOrder
		orderDetailDto.Nombre = orderDetail.Nombre

		orderResponseDto.OrdersDetail = append(orderResponseDto.OrdersDetail, orderDetailDto)
	}

	return orderResponseDto, nil
}

//Buscar orden por IDuser

func (s *orderService) GetOrdersByIdUser(idUser int) (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderCliente.GetOrdersByIdUser(idUser) //objeto de la DB, a traves del DAO
	var ordersDto dto.OrdersDto

	/*
		if (size(orders) == 0) {
			return orderResponseDto, e.NewBadRequestApiError("order not found")
		}*/

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id = order.Id
		orderDto.Fecha = order.Fecha
		orderDto.MontoFinal = order.MontoFinal

		var ordersDetail model.OrderDetails = orderDetailCliente.GetOrderDetailByIdOrder(order.Id)
		for _, orderDetail := range ordersDetail {
			var orderDetailDto dto.OrderDetailDto

			orderDetailDto.Id = orderDetail.Id
			orderDetailDto.Cantidad = orderDetail.Cantidad
			orderDetailDto.IdProducto = orderDetail.IdProduct
			orderDetailDto.PrecioUnitario = orderDetail.PrecioUnitario
			orderDetailDto.Total = orderDetail.Total
			orderDetailDto.IdOrder = orderDetail.IdOrder
			orderDetailDto.Nombre = orderDetail.Nombre

			orderDto.OrdersDetail = append(orderDto.OrdersDetail, orderDetailDto)
		}
		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}
