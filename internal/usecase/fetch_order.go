package usecase

import (
	"github.com/booscaaa/desafio-clean-arch-go-expert-pos/internal/entity"
)

type FetchOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFetchOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *FetchOrderUseCase {
	return &FetchOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *FetchOrderUseCase) Execute() ([]entity.Order, error) {
	orders, err := c.OrderRepository.Fetch()

	if err != nil {
		return nil, err
	}

	return orders, nil
}
