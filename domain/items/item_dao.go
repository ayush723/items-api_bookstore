package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ayush723/items-api_bookstore/clients/elasticsearch"
	"github.com/ayush723/items-api_bookstore/domain/queries"

	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

const (
	indexItem = "items"
	typeItem  = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	//so cpying the parameter id to reassign again at last
	itemId := i.Id
	//doesnot export id as json 
	result, err := elasticsearch.Client.Get(indexItem, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404"){
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))

	}
	if !result.Found {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil{
		return rest_errors.NewInternalServerError("error when trying to parse database response",errors.New("database error"))
	}
	if err := json.Unmarshal(bytes,i); err != nil{
		return rest_errors.NewInternalServerError("error when trying to parse database response",errors.New("database error"))
	}
	
	//reassigning id from parameter
	i.Id = itemId
	return nil
}


func (i *Item) Search(query queries.EsQuery) ([]Item ,rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItem, query.Build())
	if err != nil{
		return nil , rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits{
		bytes,_ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil{
		return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index]  = item
	}
	
	if len(items) == 0{
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")

	}

	return items, nil
}