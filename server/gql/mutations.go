package gql

import (
	"github.com/graphql-go/graphql"
)

// RootMutation contains the main mutations for the GraphQL API
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"googleLoginWithToken": &graphql.Field{
			Type:        userType,
			Description: "Receives an id_token from a client-side login to Google.",
			Args: graphql.FieldConfigArgument{
				"idToken": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: googleLoginWithToken,
		},
		"fbLoginWithToken": &graphql.Field{
			Type:        userType,
			Description: "Receives an access_token from a client-side login to Facebook, and checks with FB that this is a valid token.",
			Args: graphql.FieldConfigArgument{
				"accessToken": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: fbLoginWithToken,
		},
		"createUser": &graphql.Field{
			Type:        userType,
			Description: "Allows admins to create users",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type:        userCreateType,
					Description: "The data payload for the user to add",
				},
			},
			Resolve: createUser,
		},
		"editUser": &graphql.Field{
			Type:        userType,
			Description: "Allows admins to create users",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type:        userEditType,
					Description: "The data payload for the user to edit.",
				},
			},
			Resolve: editUser,
		},
		"deleteUser": &graphql.Field{
			Type:        graphql.String,
			Description: "Allows admins to delete a user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Deletes the user with the given id and returns the id as confirmation",
				},
			},
			Resolve: deleteUser,
		},
		"createCategory": &graphql.Field{
			Type:        categoryType,
			Description: "Allows admins and editors to create categories",
			Args: graphql.FieldConfigArgument{
				"category": &graphql.ArgumentConfig{
					Type:        categoryCreateType,
					Description: "The data payload for the category to add",
				},
			},
			Resolve: createCategory,
		},
	},
})
