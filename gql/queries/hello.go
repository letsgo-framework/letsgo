package queries

import "github.com/graphql-go/graphql"

// Hello world grqphql query resolver
var Hello = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "letsGo", nil
	},
}
