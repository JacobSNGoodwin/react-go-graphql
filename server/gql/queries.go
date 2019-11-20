package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/logger"
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
				header, _ := GetHeader(p.Context)

					ctxLogger.WithFields(logrus.Fields{
					"user": header.Get("user"),
					"role": header.Get("role"),
					).Debugln("In users query")
				return nil, nil
			},
			,
		"user": &graphql.Field{
			Type:        userType,
			Description: "Gets a single user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
					,
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, nil
					
				return nil,nil
				, 
		},
	})
},
