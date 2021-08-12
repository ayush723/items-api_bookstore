package items

import (
	"errors"

	"github.com/ayush723/items-api_bookstore/clients/elasticsearch"

	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}