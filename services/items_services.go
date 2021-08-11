package services

import (
	"github.com/ayush723/items-api_bookstore/domain/items"
	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

var (
	//interface instance of struct itemsservice. in order to mock the functions
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) ( *items.Item, *rest_errors.RestErr)
	Get(string) *items.Item
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item)( *items.Item, *rest_errors.RestErr) {
	if err :=  item.Save(); err != nil{
		return nil,err
	}
	return &item, nil

}

func (s *itemsService) Get(string) *items.Item {
	return nil
}
