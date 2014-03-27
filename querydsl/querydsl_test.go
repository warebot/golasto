package querydsl

import ("testing"
        "fmt"
        "encoding/json"
)

func TestFilteredQuery(t *testing.T){
    filteredQuery := FilteredQuery{} 
    queryString := QueryString{"*"}
    filteredQuery.Query = &queryString
    rangeF := Range{Field: "detecttime", Gte: 1232131}
    rangeT := Range{Field: "detecttime", Lte: 1232131}
    filteredQuery.And(&rangeF)
    filteredQuery.And(&rangeT)
    j, err := json.Marshal(&filteredQuery)
    if err != nil {
        t.Fatalf(err.Error())
    }
    fmt.Println(string(j))


}

func TestHasParentQuery(t *testing.T){
    hasParentQuery := HasParent{}
    hasParentQuery.ParentType = "threat"
    //queryString := QueryString{"*"}
    termQuery := Term{"address", "paypal.com"}
    hasParentQuery.Query = &termQuery
    j, err := json.Marshal(&hasParentQuery)
    if err != nil {
        t.Fatalf(err.Error())
    }
    fmt.Println(string(j))


}
