package services

import (
	"GarageSaleAPI/domain/sale"
	"GarageSaleAPI/infrastructure/persistence/memory"
	"GarageSaleAPI/interfaces/dto"

	"github.com/google/uuid"
)

var saleRepository sale.SaleRepository = new(memory.InMemorySaleRepository)

func AddSale(saleDTO dto.SaleDTO) error {
	saleId := uuid.NewString()
	sale := sale.CreateSale(saleId, saleDTO.Name, saleDTO.Address)

	saleRepository.AddSale(sale)

	return nil
}
