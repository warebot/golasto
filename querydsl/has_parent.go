package querydsl

import (
    "encoding/json"
)

type HasParent struct{
    Query_ interface{} `json:"query"`
    ParentType string `json:"parent_type"`
}

func HasParentQueryBuilder() *HasParent{
    return new(HasParent)
}


func (hp *HasParent) MarshalJSON() ([]byte, error){
    return json.Marshal(map[string]interface{}{
           "has_parent": *hp})
}

func (hp *HasParent) Query(query interface{}) *HasParent{
    hp.Query_ = query
    return hp
}

func (hp *HasParent) Parent(parentType string) *HasParent{
    hp.ParentType = parentType
    return hp
}
