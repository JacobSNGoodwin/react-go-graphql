package gql

import (
	"github.com/graphql-go/graphql"
)

// RootMutation contains the main mutations for the GraphQL API
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"googleLoginWithToken": &graphql.Field{
			Type:        graphql.String,
			Description: "Receives an id_token from a client-side login to Google. The mutation check",
			Args: graphql.FieldConfigArgument{
				"idToken": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: googleLoginWithToken,
		},
		"fbLoginWithToken": &graphql.Field{
			Type:        graphql.String,
			Description: "Receives an access_token from a client-side login to Facebook, and checks with FB that this is a valid token. If so, a jwt is returned as a string",
			Args: graphql.FieldConfigArgument{
				"fbLoginData": &graphql.ArgumentConfig{
					Type: fbLoginType,
				},
			},
			Resolve: fbLoginWithToken,
		},
	},
})
