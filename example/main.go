package main

import (
	"encoding/json"
	"esquery"
	"fmt"
)

type Person struct {
	Name   string
	Age    string
	Gender string
}

func main() {

	esClient := esquery.ESClient{Url: "http://localhost:9200", IndexName: "person"}

	esClient.Init()

	IndexPerson(&esClient)

	SearchPersonByName(&esClient)

	SearchPersonByGender(&esClient)

}

func SearchPersonByName(esClient *esquery.ESClient) {
	query := esquery.NewQuery()

	term := query.Term("Name", "test")
	query.AddPart(term)

	response, _ := esClient.Search(query)

	fmt.Println("search result count person by name :> ")
	fmt.Println(response.Hits.Total)

}

func SearchPersonByGender(esClient *esquery.ESClient) {
	query := esquery.NewQuery()

	term := query.Term("Gender", "male")
	query.AddPart(term)

	response, _ := esClient.Search(query)

	fmt.Println("search   result count person by gender :> ")
	fmt.Println(response.Hits.Total)

}

func IndexPerson(esClient *esquery.ESClient) {
	//index person 1
	person1 := &Person{Name: "test", Age: "20", Gender: "male"}

	b, err := json.Marshal(person1)
	if err != nil {
		fmt.Println(err)
	}
	esClient.Index(string(b), "true")
}
