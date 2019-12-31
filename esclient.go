package esquery

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

var ElasticClient *ESClient

type ESClient struct {
	Client    *elasticsearch.Client
	Url       string
	IndexName string
}

type Response struct {
	Took int
	Hits struct {
		Total int
		Hits  []struct {
			ID         string          `json:"_id"`
			Source     json.RawMessage `json:"_source"`
			Highlights json.RawMessage `json:"highlight"`
			Sort       []interface{}   `json:"sort"`
		}
	}
}

func (esClient *ESClient) Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{esClient.Url},
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	esClient.Client = es
	ElasticClient = esClient

}

func (esClient *ESClient) Index(body string, refersh string) error {
	req := esapi.IndexRequest{
		Index:   esClient.IndexName,
		Body:    strings.NewReader(body),
		Refresh: refersh,
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), esClient.Client)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func (esClient *ESClient) Search(query *Query) (*Response, error) {

	client := esClient.Client
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(esClient.IndexName),
		client.Search.WithBody(strings.NewReader(query.query())),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r Response
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil

}

func (esClient *ESClient) GetQuery(query *Query) string {

	return query.query()

}
