package controllers

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"net/http"

	"github.com/ayush723/items-api_bookstore/utils/http_utils"
	"github.com/gorilla/mux"

	"github.com/ayush723/items-api_bookstore/domain/items"
	"github.com/ayush723/items-api_bookstore/domain/queries"
	"github.com/ayush723/items-api_bookstore/services"
	"github.com/ayush723/utils-go_bookstore/rest_errors"

	"github.com/ayush723/oauth-go_bookstore/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (s *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// http_utils.RespondError(w, *err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("unable to retrieve user information from given access_token")
		http_utils.RespondError(w, respErr)
		return
	}

	var itemRequest items.Item
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resperr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, resperr)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		resperr := rest_errors.NewBadRequestError("invalid items json")
		http_utils.RespondError(w, resperr)
		return
	}

	itemRequest.Seller = sellerId
	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (s *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondJson(w, http.StatusOK, item)
	}
}

func (s *itemsController) Search(w http.ResponseWriter, r *http.Request){
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil{
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()
	//the incoming request should have format of EsQuery type
	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil{
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	items, searchErr := services.ItemsService.Search(query)
	if  searchErr != nil{
		http_utils.RespondError(w, searchErr)
		return
	}
	http_utils.RespondJson(w,http.StatusOK,items)
	
	
}