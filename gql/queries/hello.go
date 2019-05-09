package queries

import "github.com/graphql-go/graphql"

var Hello = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "letsGo", nil
	},
}
