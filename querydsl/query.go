package querydsl

import (
    "encoding/json"
)

type Query struct{
    Query interface{} `json:"query"`
    Size_ int `json:"size,omitempty"`
    From_ int `json:"from,omitempty"`
}


func QueryBuilder() *Query{
    return new (Query)
}

// SetQuery sets the query param of the elasticsearch query
func (q *Query) SetQuery(query interface{}) *Query{
    q.Query = query
    return q
}

func (q *Query) Size(size int) *Query{
    q.Size_ = size
    return q
}

func (q *Query) From(from int) *Query{
    q.From_ = from
    return q
}


func (q *Query) MarshalJSON() ([]byte, error){
    return json.Marshal(*q)
}


type Term struct {
    Field_ string
    Value_ string
}

func TermQuery() *Term{
    return new(Term)
}

func (t* Term) Field(field string) *Term{
    t.Field_ = field
    return t
}

func (t* Term) Value(value string) *Term{
    t.Value_ = value
    return t
}
func (t *Term) MarshalJSON() ([]byte, error){
    term := make(map[string]interface{})
    term[t.Field_] = t.Value_
    return json.Marshal(map[string]interface{}{"term": term})
}

type WildCard struct {
    Field_ string
    Value_ string
}

func WildCardBuilder() *WildCard{
    return new(WildCard)
}

func (w *WildCard) Field(field string) *WildCard{
    w.Field_ = field
    return w
}

func (w *WildCard) Value(value string) *WildCard{
    w.Value_ = value
    return w
}

func (w *WildCard) MarshalJSON() ([]byte, error){
    wildCard := make(map[string]interface{})
    wildCard[w.Field_] = w.Value_
    return json.Marshal(map[string]interface{}{"wildcard": wildCard})
}
