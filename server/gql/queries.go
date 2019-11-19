package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/data"
	"github.com/maxbrain0/react-go-graphql/server/logger"
	"github.com/sirupsen/logrus"
)

var ctxLogger = logger.CtxLogger

// GetRootQuery returns the root query with the datasource plugged into it
func GetRootQuery(ds *data.Data) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "A list of all users",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					header, _ := GetHeader(p.Context)

					ctxLogger.WithFields(logrus.Fields{
						"user": header.Get("user"),
						"role": header.Get("role"),
					}).Debugln("In users query")
					return ds.GetUsers(), nil
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
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, nil
					}
					return ds.GetUserByID(id)
				},
			},
		},
	})
}
