package gql

import (
	"github.com/graphql-go/graphql"
)

// RootQuery holds all of the main queries for our GQL schema. It is exported as the schema is
// currently configured in main
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "A list of all users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		},
	},
})
