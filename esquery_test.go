package esquery

import (
	"testing"
)

func TestTermClause(t *testing.T) {
	query := NewQuery()

	//term query in a specific field
	term := query.Term("testField1", "testVal")

	//Add the 'term' to the original query
	query.AddPart(term)

	//final query
	got := query.query()
	expected := `{ "query" :  { "term" :  { "testField1" : "testVal" }  }  } `

	if got != expected {
		t.Error("expected :" + expected + " got :" + got)
	}
}

func TestMatchClause(t *testing.T) {
	query := NewQuery()

	//match query in a specific field
	match := query.Match("testField1", "testVal")

	//Add the 'match' to the original query
	query.AddPart(match)

	//final query
	got := query.query()
	expected := `{ "query" :  { "match" :  { "testField1" : "testVal" }  }  } `

	if got != expected {
		t.Error("expected :" + expected + " got :" + got)
	}
}

func TestMustClause(t *testing.T) {
	query := NewQuery()
	//term query in testField1
	term1 := query.Term("testField1", "testVal1")

	//term query in testField1
	term2 := query.Term("testField2", "testVal2")

	//create must descriptor with two terms
	must := query.Must([]*QueryPart{term1, term2})

	//Add 'must' part to the original query
	query.AddPart(must)

	//final query
	got := query.query()

	expected := `{ "query" :  { "must" :  [  { "term" :  { "testField1" : "testVal1" }  } , { "term" :  { "testField2" : "testVal2" }  }  ]  }  } `

	if got != expected {
		t.Error("expected :" + expected + " got :" + got)
	}
}

func TestShouldClause(t *testing.T) {
	query := NewQuery()
	//term query in testField1
	term1 := query.Term("testField1", "testVal1")

	//term query in testField1
	term2 := query.Term("testField2", "testVal2")

	//create 'should' descriptor with two terms
	should := query.Should([]*QueryPart{term1, term2})

	//Add 'should' part to the main query
	query.AddPart(should)

	got := query.query()

	expected := `{ "query" :  { "should" :  [  { "term" :  { "testField1" : "testVal1" }  } , { "term" :  { "testField2" : "testVal2" }  }  ]  }  } `

	if got != expected {
		t.Error("expected :" + expected + " got :" + got)
	}
}

func TestBoolClause(t *testing.T) {
	query := NewQuery()
	//term query in testField1
	term1 := query.Term("testField1", "testVal1")

	//term query in testField1
	term2 := query.Term("testField2", "testVal2")

	//create 'should' descriptor with two terms
	should := query.Should([]*QueryPart{term1, term2})

	//create 'bool' descriptor with the 'should' clause
	bool := query.Bool([]*QueryPart{should})

	//Add the 'bool' to the main query
	query.AddPart(bool)

	//final query
	got := query.query()

	expected := `{ "query" :  { "bool" :  { "should" :  [  { "term" :  { "testField1" : "testVal1" }  } , { "term" :  { "testField2" : "testVal2" }  }  ]  }  }  } `

	if got != expected {
		t.Error("expected :" + expected + " got :" + got)
	}
}

func TestBoolWithMustShouldMustNotClauses(t *testing.T) {
	query := NewQuery()
	//term query in testField1
	term1 := query.Term("testField1", "testVal1")

	term2 := query.Term("testField2", "testVal2")

	//create must, must_not and should nested queries
	must := query.Must([]*QueryPart{term1})

	mustNot := query.MustNot([]*QueryPart{term2})

	should := query.Should([]*QueryPart{term1, term2})

	//create 'bool' descriptor with the 'should' clause
	bool := query.Bool([]*QueryPart{must, mustNot, should})

	//Add the 'bool' to the main query
	query.AddPart(bool)

	//final query
	got := query.query()

	expected := `{ "query" :  { "bool" :  { "must" :  { "term" :  { "testField1" : "testVal1" }  } ,"must_not" :  { "term" :  { "testField2" : "testVal2" }  } ,"should" :  [  { "term" :  { "testField1" : "testVal1" }  } , { "term" :  { "testField2" : "testVal2" }  }  ]  }  }  } `

	if got != expected {
		t.Error("expected :" + expected + " got :" + got)
	}
}

func TestFilterClause(t *testing.T) {
	query := NewQuery()
	//term query in testField1
	term1 := query.Term("testField1", "testVal1")

	//term query in testField1
	term2 := query.Term("testField2", "testVal2")

	//create 'filter' descriptor with the 'should' clause
	filter := query.Filter([]*QueryPart{term1, term2})

	//Add the 'filter' to the main query
	query.AddPart(filter)

	//final query
	got := query.query()

	expected := `{ "query" :  { "filter" :  [  { "term" :  { "testField1" : "testVal1" }  } , { "term" :  { "testField2" : "testVal2" }  }  ]  }  } `

	if got != expected {
		t.Error("expected :" + expected + " got :" + got)
	}
}

func TestBoolWithFilterClause(t *testing.T) {
	query := NewQuery()
	//term query in testField1
	term1 := query.Term("testField1", "testVal1")

	//term query in testField1
	term2 := query.Term("testField2", "testVal2")

	//create 'filter' descriptor
	filter := query.Filter([]*QueryPart{term1, term2})

	//create 'bool' clause with 'must' and 'filter' having two terms
	bool := query.Bool([]*QueryPart{query.Must([]*QueryPart{term1, term2}), filter})

	query.AddPart(bool)

	//final query
	got := query.query()

	expected := `{ "query" :  { "bool" :  { "must" :  [  { "term" :  { "testField1" : "testVal1" }  } , { "term" :  { "testField2" : "testVal2" }  }  ] ,"filter" :  [  { "term" :  { "testField1" : "testVal1" }  } , { "term" :  { "testField2" : "testVal2" }  }  ]  }  }  } `

	if got != expected {
		t.Error("expected :" + expected + " got :" + got)
	}
}
