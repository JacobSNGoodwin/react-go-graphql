package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/models"
)

// userType holds information for users
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "A user with its accompanying properties",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// need to resolve uuid to string
				if user, ok := p.Source.(models.User); ok {
					return user.ID.String(), nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Holds the user's unique email address",
		},
		"imageUri": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Holds the user's image Uri, if any",
		},
	},
})

// userType holds information for users
var fbLoginType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "FBLoginType",
	Description: "Fields required for fbLoginWithToken",
	Fields: graphql.InputObjectConfigFieldMap{
		"token": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The access token as a string provided by client-side Login with Facebook",
		},
		"name": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The full name as provided in the name field by client-side Login with Facebook",
		},
		"email": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The full name as provided in the name field by client-side Login with Facebook",
		},
		"imageUri": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "The image URI provided by client-side Login with Facebook",
		},
	},
})
