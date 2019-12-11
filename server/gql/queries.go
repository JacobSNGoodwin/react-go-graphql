package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/maxbrain0/react-go-graphql/server/models"
)

var ctxLogger = logger.CtxLogger

// RootQuery contains the main query for the GQL api
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "A list of all users",
			Args: graphql.FieldConfigArgument{
				"limit": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 5,
				},
				"offset": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 0,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var users models.Users
				err := users.GetAll(p)
				return users, err
			},
		},
		"user": &graphql.Field{
			Type:        userType,
			Description: "Gets a single user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:         graphql.String,
					Description:  "A v4 uuid casted as a string",
					DefaultValue: "",
				},
			},
			Resolve: user,
		},
	},
})
