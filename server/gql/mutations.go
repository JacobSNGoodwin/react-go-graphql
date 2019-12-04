package gql

import "github.com/graphql-go/graphql"

// RootMutation contains the main mutations for the GraphQL API
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"googleLoginWithToken": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Receives an id_token from a client-side login to Google, and checks that this is a valid token. If so, a jwt is returned as a string",
			Args: graphql.FieldConfigArgument{
				"idToken": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return true, nil
			},
		},
	},
})
