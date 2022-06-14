package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	addressCliente "mvc-go/clients/address" //DAO
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type addressService struct{}

type addressServiceInterface interface {
	//siempre devuelve dto o error
	GetAddressById(id int) (dto.AddressDto, e.ApiError)
	GetAddresses() (dto.AddressesDto, e.ApiError)
	GetAddressesByIdUser(idUser int) (dto.AddressesDto, e.ApiError)
}

var (
	AddressService addressServiceInterface
)

func init() {
	AddressService = &addressService{}
}

func (s *addressService) GetAddressById(id int) (dto.AddressDto, e.ApiError) {

	var address model.Address = addressCliente.GetAddressById(id) //objeto de la DB, a traves del DAO
	var addressDto dto.AddressDto

	if address.Id == 0 {
		return addressDto, e.NewBadRequestApiError("address not found")
	}
	addressDto.Id = address.Id
	addressDto.Number = address.Number
	addressDto.Neighborhood = address.Neighborhood
	addressDto.Street = address.Street
	addressDto.City = address.City
	addressDto.IdUser = address.IdUser
	return addressDto, nil
}

func (s *addressService) GetAddresses() (dto.AddressesDto, e.ApiError) {

	var addresses model.Addresses = addressCliente.GetAddresses()
	var addressesDto dto.AddressesDto

	for _, address := range addresses {
		var addressDto dto.AddressDto
		addressDto.Id = address.Id
		addressDto.Number = address.Number
		addressDto.Neighborhood = address.Neighborhood
		addressDto.Street = address.Street
		addressDto.City = address.City
		addressDto.IdUser = address.IdUser
		addressesDto = append(addressesDto, addressDto)
	}

	return addressesDto, nil
}

func (s *addressService) GetAddressesByIdUser(idUser int) (dto.AddressesDto, e.ApiError) {

	var addresses model.Addresses = addressCliente.GetAddressesByIdUser(idUser) //objeto de la DB, a traves del DAO
	var addressesDto dto.AddressesDto
	/*
		if addresses.Id == 0 {
			return addressesDto, e.NewBadRequestApiError("addresses not found")
		}*/
	for _, address := range addresses {
		var addressDto dto.AddressDto
		addressDto.Id = address.Id
		addressDto.Number = address.Number
		addressDto.Neighborhood = address.Neighborhood
		addressDto.Street = address.Street
		addressDto.City = address.City
		addressDto.IdUser = address.IdUser
		addressesDto = append(addressesDto, addressDto)
	}
	return addressesDto, nil

}
