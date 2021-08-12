package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/ayush723/utils-go_bookstore/logger"
	"github.com/olivere/elastic"
)

var (
	//making client available through interface for mockup during tests
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetRetrier(NewCustomRetrier()),
		// elastic.SetGzip(true),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		// elastic.SetHeaders(http.Header{
		//   "X-Caller-Id": []string{"..."},
		// }),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)

	//Create index if it does not exists

}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get document by id %s", id), err)
		return nil, err
	}
	return result, nil
}
//todo:
func (c *esClient) Search(index string){
	ctx := context.Background()

	// elastic.NewMatchQuery()
	result, err := c.client.Search(index).Query()
}
func (c *esClient) Add(){}
func (c *esClient) Update(){}
func (c *esClient) Delete(){}