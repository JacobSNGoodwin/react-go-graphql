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
