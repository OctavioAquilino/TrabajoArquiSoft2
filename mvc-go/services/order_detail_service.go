package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	orderDetailCliente "mvc-go/clients/order_detail" //DAO
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	//siempre devuelve dto o error
	GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError)
	GetOrderDetails() (dto.OrderDetailsDto, e.ApiError)
	InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, e.ApiError)
	GetOrderDetailByIdOrder(idOrder int) (dto.OrderDetailsResDto, e.ApiError)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

func (s *orderDetailService) GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail = orderDetailCliente.GetOrderDetailById(id) //objeto de la DB, a traves del DAO
	var orderDetailDto dto.OrderDetailDto

	if orderDetail.Id == 0 {
		return orderDetailDto, e.NewBadRequestApiError("orderDetail not found")
	}
	orderDetailDto.Id = orderDetail.Id
	orderDetailDto.Nombre = orderDetail.Nombre
	orderDetailDto.Cantidad = orderDetail.Cantidad
	orderDetailDto.PrecioUnitario = orderDetail.PrecioUnitario
	orderDetailDto.Total = orderDetail.Total
	orderDetailDto.IdOrder = orderDetail.IdOrder
	//orderDetailDto.IdProducto = orderDetail.IdProduct
	//orderDetailDto.Producto = orderDetail.Producto ----------------Manejar
	return orderDetailDto, nil
}

func (s *orderDetailService) GetOrderDetails() (dto.OrderDetailsDto, e.ApiError) {

	var orderDetails model.OrderDetails = orderDetailCliente.GetOrderDetails()
	var orderDetailsDto dto.OrderDetailsDto

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Nombre = orderDetail.Nombre
		orderDetailDto.Cantidad = orderDetail.Cantidad
		orderDetailDto.PrecioUnitario = orderDetail.PrecioUnitario
		orderDetailDto.Total = orderDetail.Total
		orderDetailDto.IdOrder = orderDetail.IdOrder
		//orderDetailDto.IdProducto = orderDetail.IdProduct
		//orderDetailDto.Producto = orderDetail.Producto ----------------Manejar

		orderDetailsDto = append(orderDetailsDto, orderDetailDto)
	}

	return orderDetailsDto, nil
}

func (s *orderDetailService) InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail

	orderDetail.Nombre = orderDetailDto.Nombre
	orderDetail.Cantidad = orderDetailDto.Cantidad
	orderDetail.PrecioUnitario = orderDetailDto.PrecioUnitario
	orderDetail.Total = orderDetailDto.PrecioUnitario * orderDetailDto.Cantidad
	orderDetail.IdOrder = orderDetailDto.IdOrder
	//orderDetail.IdProduct = orderDetailDto.IdProducto

	orderDetail = orderDetailCliente.InsertOrderDetail(orderDetail)

	orderDetailDto.Id = orderDetail.Id
	orderDetailDto.Total = orderDetail.Total

	return orderDetailDto, nil
}

func (s *orderDetailService) GetOrderDetailByIdOrder(idOrder int) (dto.OrderDetailsResDto, e.ApiError) {

	var ordersDetail model.OrderDetails = orderDetailCliente.GetOrderDetailByIdOrder(idOrder) //objeto de la DB, a traves del DAO
	var ordersDetailResDto dto.OrderDetailsResDto
	/*
		if orderDetailRes.Id == 0 {
			return orderDetailDto, e.NewBadRequestApiError("orderDetail not found")
		}*/
	for _, orderDetailRes := range ordersDetail {
		var orderDetailResDto dto.OrderDetailResDto
		orderDetailResDto.Id = orderDetailRes.Id
		orderDetailResDto.Nombre = orderDetailRes.Nombre
		orderDetailResDto.Cantidad = orderDetailRes.Cantidad
		orderDetailResDto.PrecioUnitario = orderDetailRes.PrecioUnitario
		orderDetailResDto.Total = orderDetailRes.Total
		//orderDetailResDto.IdProducto = orderDetailRes.IdProduct

		ordersDetailResDto = append(ordersDetailResDto, orderDetailResDto)
	}
	return ordersDetailResDto, nil
}
