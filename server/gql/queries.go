package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/maxbrain0/react-go-graphql/server/models"
	"github.com/sirupsen/logrus"
)

var ctxLogger = logger.CtxLogger

// RootQuery contains the main query for the GQL api
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "A list of all users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db, _ := GetDB(p.Context)
				var users []models.User
				db.Limit(1).Find(&users)

				ctxLogger.WithFields(logrus.Fields{
					"Id":    users[0].ID,
					"Name":  users[0].Name,
					"Email": users[0].Email,
				}).Debug("Users Found")

				return users, nil
			},
		},
		"user": &graphql.Field{
			Type:        userType,
			Description: "Gets a single user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, ok := p.Args["id"].(int)
				if !ok {
					return nil, nil
				}
				return nil, nil
			},
		},
	},
})
