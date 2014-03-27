package querydsl

import (
    "encoding/json"
)

type Query struct{
}


func (q *Query) MarshalJSON() ([]byte, error){
    return json.Marshal(*q)
}



type Term struct {
    Field string
    Value string
}

func (t *Term) MarshalJSON() ([]byte, error){
    term := make(map[string]interface{})
    term[t.Field] = t.Value
    return json.Marshal(map[string]interface{}{"term": term})
}



