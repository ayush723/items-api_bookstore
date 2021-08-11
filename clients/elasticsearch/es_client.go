package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/ayush723/users-api_bookstore/logger"

	"github.com/olivere/elastic"
)

var (
	//making client available through interface for mockup during tests
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetRetrier(NewCustomRetrier()),
		// elastic.SetGzip(true),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
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

func (c *esClient)setClient(client *elastic.Client){
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err :=  c.client.Index().
	Index("items").
	BodyJson(doc).
	Do(ctx)

	if err != nil{
		logger.Error(fmt.Sprintf("error when trying to index document in index %s",index), err)
		return nil, err
	}
	return result, nil
}
