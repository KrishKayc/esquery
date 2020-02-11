# esquery
Simple go library to index and search from elastic search with ease

## Usage

### Initialize client

```
esClient := esquery.ESClient{Url: "http://localhost:9200", IndexName: "person"}

esClient.Init()
```

### Index documents

```
person1 := &Person{Name: "test", Age: "20", Gender: "male"}

b, err := json.Marshal(person1)

esClient.Index(string(b), "true")
```

## Search Easily with 'Built-In' functions

### term

```
query := esquery.NewQuery()

term := query.Term("Name", "test")

query.AddPart(term)

response, _ := esClient.Search(query)
```

### match

```
query := esquery.NewQuery()

match := query.Match("Name", "test")

query.AddPart(match)

response, _ := esClient.Search(query)
```

## Supported ES descriptors

Term, Match, Bool, Must, Should, Filter.. yet to come.
