package gql

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/letsgo-framework/letsgo/gql/queries"
	"log"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"hello": queries.Hello,
	},
})

func InitGraphql(r *gin.Engine)  {

	schemaConfig := graphql.SchemaConfig{Query: RootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	g := r.Group("/graphql")
	{
		g.POST("/", func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		})
		g.GET("/", func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		})
	}

}