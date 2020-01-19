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
				if err := users.GetAll(p); err != nil {
					return nil, err
				}
				return users, nil
			},
		},
		"user": &graphql.Field{
			Type:        userType,
			Description: "Gets a single user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "A v4 uuid casted as a string",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				if err := user.GetByID(p); err != nil {
					return nil, err
				}
				return user, nil
			},
		},
		"me": &graphql.Field{
			Type:        userType,
			Description: "Returns user for currently logged-in user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				if err := user.GetCurrent(p); err != nil {
					return nil, err
				}
				return user, nil
			},
		},
		"products": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "A list of products",
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
				var products models.Products
				if err := products.GetAll(p); err != nil {
					return nil, err
				}
				return products, nil
			},
		},
		"categories": &graphql.Field{
			Type:        graphql.NewList(categoryType),
			Description: "A list of categories",
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
				var categories models.Categories
				if err := categories.GetAll(p); err != nil {
					return nil, err
				}
				return categories, nil
			},
		},
		"product": &graphql.Field{
			Type:        productType,
			Description: "Gets a single product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "A v4 uuid casted as a string",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var product models.Product
				if err := product.GetByID(p); err != nil {
					return nil, err
				}
				return product, nil
			},
		},
	},
})
