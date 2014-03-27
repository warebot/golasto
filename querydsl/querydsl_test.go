package querydsl

import ("testing"
        "encoding/json"
)
func TestFilteredQuery(t *testing.T){

    queryTst := `{"filtered":{"query":{"query_string":{"query":"something"}},"filter":{"and":[{"range":{"detecttime":{"gte":1232131}}},{"range":{"detecttime":{"lte":1232131}}}]}}}` 
    var queryString = QueryStringBuilder().Query("something")
    rangeF := Range{Field: "detecttime", Gte: 1232131}
    rangeT := Range{Field: "detecttime", Lte: 1232131}
    var filteredQuery = FilteredQueryBuilder().And(&rangeF).And(&rangeT)
    filteredQuery.Query(queryString)
    j, err := json.Marshal(filteredQuery)
    if err != nil {
        t.Fatalf(err.Error())
    }
    if queryTst != string(j){
        t.Fatalf("NOT EQUAL")
    }

}

func TestHasParentQuery(t *testing.T){
    queryTst := `{"has_parent":{"query":{"term":{"address":"paypal.com"}},"parent_type":"threat"}}`
    termQuery := Term{"address", "paypal.com"}
    var hasParentQuery = HasParentQueryBuilder().Parent("threat").Query(&termQuery)
    j, err := json.Marshal(&hasParentQuery)
    if err != nil {
        t.Fatalf(err.Error())
    }
    if queryTst != string(j){
        t.Fatalf("NOT EQUAL")
    }

}

func TestQueryString(t *testing.T){
    queryTst := `{"query_string":{"query":"paypal"}}`
    var queryString = QueryStringBuilder().Query("paypal")
    j, err := json.Marshal(queryString)
    if err != nil {
        t.Fatalf(err.Error())
    }
    if queryTst != string(j){
        t.Fatalf("NOT EQUAL")
    }



}

