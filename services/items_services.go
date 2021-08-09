package services

import (
	"github.com/ayush723/items-api_bookstore/domain/items"

	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

var(
	//interface instance of struct itemsservice. in order to mock the functions
	ItemsService itemsServiceInterface = &itemsService{}
)


type itemsServiceInterface interface{
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)

}

type itemsService struct{}

func(s *itemsService) Create(items.Item)(*items.Item, rest_errors.RestErr){
	return nil, nil
}

func(s *itemsService) Get(string) (*items.Item, rest_errors.RestErr){
	return nil, nil
}