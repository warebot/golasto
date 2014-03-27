package querydsl

import (
    "encoding/json"
)

type QueryString struct{
    Qry string `json:"query"`
}


func QueryStringBuilder() *QueryString{
    return new(QueryString)

}


func (qs *QueryString) MarshalJSON() ([]byte, error){
    return json.Marshal(map[string]interface{}{
           "query_string": *qs})
}


func (qs *QueryString) Query(query string) *QueryString{
    qs.Qry = query
    return qs
}

