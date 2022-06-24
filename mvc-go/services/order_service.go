package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	orderCliente "mvc-go/clients/order" //DAO
	orderDetailCliente "mvc-go/clients/order_detail"
	productCliente "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	//"mvc-go/utils/modeladto"
)

type orderService struct{}

type orderServiceInterface interface {
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
	GetOrdersByIdUser(token string) (dto.OrdersDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
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

func (s *orderService) GetOrdersByIdUser(token string) (dto.OrdersDto, e.ApiError) {
	var idUser float64
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, e.NewUnauthorizedApiError("Unauthorized")
		}
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}

	if !tkn.Valid {
		return nil, e.NewUnauthorizedApiError("Unauthorized")

	}
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {

		idUser = (claims["id_user"].(float64))

	} else {
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}
	var IdUserX int = int(idUser)
	var orders model.Orders = orderCliente.GetOrdersByIdUser(IdUserX) //objeto de la DB, a traves del DAO
	var ordersDto dto.OrdersDto

	if len(orders) == 0 {
		return ordersDto, e.NewBadRequestApiError("order not found")
	}

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
