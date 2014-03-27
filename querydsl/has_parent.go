package querydsl

import (
    "encoding/json"
)

type HasParent struct{
    Query interface{} `json:"query"`
    ParentType string `json:"parent_type"`
}


func (hp *HasParent) MarshalJSON() ([]byte, error){
    return json.Marshal(map[string]interface{}{
           "has_parent": *hp})
}


