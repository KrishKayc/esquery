package esquery

import (
	"strings"
)

//Represents an 'Elastic Search Query'
type Query struct {
	Parts []*QueryPart
}

//QueryPart consists of a descriptor 'eg: bool, term, must' and val
type QueryPart struct {
	descriptor    string
	value         string
	isArray       bool
	nestedQueries []*QueryPart
}

func NewQuery() *Query {
	return &Query{Parts: make([]*QueryPart, 0)}
}

func (q *Query) query() string {
	sb := strings.Builder{}
	sb.WriteString(`{ "query" : `)
	sb.WriteString(" { ")
	sb.WriteString(getQuery(q.Parts, false))
	sb.WriteString(" } ")
	sb.WriteString(" } ")
	return sb.String()
}

func getQuery(parts []*QueryPart, isParentArray bool) string {
	sb := strings.Builder{}
	totalParts := len(parts)

	for ind, part := range parts {

		if isParentArray {
			sb.WriteString(" { ")
		}

		sb.WriteString(`"` + part.descriptor + `"` + " : ")

		if part.value != "" {
			sb.WriteString(`"` + part.value + `"`)
		}
		if part.nestedQueries != nil {

			if part.isArray {
				sb.WriteString(" [ ")
			} else {
				sb.WriteString(" { ")
			}
			sb.WriteString(getQuery(part.nestedQueries, part.isArray))

			if part.isArray {
				sb.WriteString(" ] ")
			} else {
				sb.WriteString(" } ")
			}

		}

		if isParentArray {
			sb.WriteString(" } ")
		}

		if ind != (totalParts - 1) {
			sb.WriteString(",")
		}
	}

	return sb.String()
}

func (q *Query) AddPart(part *QueryPart) {
	q.Parts = append(q.Parts, part)
}

func (q *Query) AddParts(parts []*QueryPart) {
	q.Parts = append(q.Parts, parts...)
}

//Represents a 'bool' query in elastic search
func (q *Query) Bool(nestedQueries []*QueryPart) *QueryPart {
	bool := &QueryPart{descriptor: "bool"}
	bool.nestedQueries = nestedQueries
	return bool
}

//Represents 'should' descriptor of elastic search
func (q *Query) Should(nestedQueries []*QueryPart) *QueryPart {
	should := &QueryPart{descriptor: "should", isArray: len(nestedQueries) > 1}
	should.nestedQueries = nestedQueries
	return should
}

//Represents 'must' descriptor of elastic search
func (q *Query) Must(nestedQueries []*QueryPart) *QueryPart {
	must := &QueryPart{descriptor: "must", isArray: len(nestedQueries) > 1}
	must.nestedQueries = nestedQueries
	return must
}

//Represents 'must_not' descriptor of elastic search
func (q *Query) MustNot(nestedQueries []*QueryPart) *QueryPart {
	mustNot := &QueryPart{descriptor: "must_not", isArray: len(nestedQueries) > 1}
	mustNot.nestedQueries = nestedQueries
	return mustNot
}

//Represents 'filter' descriptor of elastic search
func (q *Query) Filter(nestedQueries []*QueryPart) *QueryPart {
	filter := &QueryPart{descriptor: "filter", isArray: len(nestedQueries) > 1}
	filter.nestedQueries = nestedQueries
	return filter
}

//Represent 'term' Elastic search query with field name and fieldValue
func (q *Query) Term(field string, value string) *QueryPart {
	term := &QueryPart{descriptor: "term"}
	term.SetNestedQueryPart(field, value)
	return term
}

//Represent 'match' Elastic search query with field name and fieldValue
func (q *Query) Match(field string, value string) *QueryPart {
	match := &QueryPart{descriptor: "match"}
	match.SetNestedQueryPart(field, value)
	return match
}

func (parent *QueryPart) SetNestedQueryPart(descriptor string, value string) {
	parent.nestedQueries = make([]*QueryPart, 0)
	parent.nestedQueries = append(parent.nestedQueries, getQueryPart(descriptor, value))
}

func getQueryPart(descriptor string, value string) *QueryPart {
	return &QueryPart{descriptor: descriptor, value: value}
}
