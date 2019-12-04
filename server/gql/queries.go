package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/maxbrain0/react-go-graphql/server/models"
	uuid "github.com/satori/go.uuid"
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
				db, _ := GetDB(p.Context)
				var users []models.User
				if result :=
					db.
						Order("email").
						Limit(p.Args["limit"].(int)).
						Offset(p.Args["offset"].(int)).
						Find(&users); result.Error != nil {
					return nil, nil
				}

				return users, nil
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db, _ := GetDB(p.Context)
				var user models.User

				// Find by uuid or email, which should both be unique
				if result := db.
					Where("id = ?", uuid.FromStringOrNil(p.Args["id"].(string))).
					Find(&user); result.Error != nil {
					return nil, nil
				}

				ctxLogger.WithFields(logrus.Fields{
					"ID":    user.ID,
					"Email": user.Email,
					"Name":  user.Name,
				}).Debug("Found user by id or email")

				return user, nil
			},
		},
	},
})
