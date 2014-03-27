package querydsl

import (
    "encoding/json"
)


type FilteredQuery struct{
    Query interface{} `json:"query"`
    // Filter can be
    // bool
    // and
    // or 
    // query
    Filter map[string]interface{} `json:"filter"`
}

type RangeFilter struct{
    Range map[string]Range `json:"range"`
}

type Range struct{
    Field string `json:"-"`
    Gte interface{} `json:"gte,omitempty"`
    Lte  interface{} `json:"lte,omitempty"`  
}

func (r *Range) MarshalJSON() ([]byte, error){
    return json.Marshal(map[string]interface{}{
        "range": map[string]interface{}{
            r.Field: *r },
        })
}



func (fq *FilteredQuery) MarshalJSON() ([]byte, error){
    // If filter is empty initialize to 
    if fq.Filter == nil{
        fq.Filter = make(map[string]interface{})
    }
    return json.Marshal(map[string]interface{}{
           "filtered": *fq})


}


func (fq *FilteredQuery) And (f interface{}){
    if fq.Filter == nil{
        fq.Filter = make(map[string]interface{})
    }
    if fq.Filter["and"] == nil {
        var andFilter []interface{}
        fq.Filter["and"] = andFilter
    }
    fq.Filter["and"] = append(fq.Filter["and"].([]interface{}), f)
}

func (fq *FilteredQuery) Range(from interface{}, to interface{}){
    
}
