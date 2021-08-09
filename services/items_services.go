package services

import (
	"github.com/ayush723/items-api_bookstore/domain/items"
)

var (
	//interface instance of struct itemsservice. in order to mock the functions
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item)
	Get(string) (*items.Item)
}

type itemsService struct{}

func (s *itemsService) Create(items.Item) (*items.Item ) {
	return nil
}

func (s *itemsService) Get(string) (*items.Item) {
	return nil
}
