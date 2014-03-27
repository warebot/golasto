package querydsl

import (
    "encoding/json"
)

type QueryString struct{
    Query string `json:"query"`
}


func (qs *QueryString) MarshalJSON() ([]byte, error){
    return json.Marshal(map[string]interface{}{
           "query_string": *qs})


}


