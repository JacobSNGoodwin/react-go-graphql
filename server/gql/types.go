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
		"roles": &graphql.Field{
			Type:        graphql.NewList(roleEnum),
			Description: "Holds a list of roles for the user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				roles := []string{}
				if user, ok := p.Source.(models.User); ok {
					for _, role := range user.Roles {
						roles = append(roles, role.Name)
					}
					return roles, nil
				}
				return nil, nil
			},
		},
	},
})

var userInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "Name of the user who you want to provide access to",
			},
			"email": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "The email address of the user. User must use this email on FB or Google",
			},
			"imageUri": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "Holds the user's image Uri, if any",
			},
			"roles": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewList(roleEnum),
				Description: "An array of roles to assign to the user",
			},
		},
	},
)

var roleEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Role",
	Description: "Holds the roles available for this API",
	Values: graphql.EnumValueConfigMap{
		"Admin": &graphql.EnumValueConfig{
			Value: "admin",
		},
		"Editor": &graphql.EnumValueConfig{
			Value: "editor",
		},
	},
})
