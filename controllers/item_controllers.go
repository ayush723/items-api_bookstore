package controllers

import (
	"fmt"
	"net/http"

	"github.com/ayush723/items-api_bookstore/domain/items"
	"github.com/ayush723/items-api_bookstore/services"

	"github.com/ayush723/oauth-go_bookstore/oauth"
)

var (
	ItemController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (s *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: return error to the user
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//todo:retunr error json to the user
	}

	fmt.Println(result)

	//todo:return created item as json with http status 201- created
}

func (s *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
