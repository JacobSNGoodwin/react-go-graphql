package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/data"
)

// UserType holds information for users
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "A user with id and name",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(data.User); ok {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(data.User); ok {
					return user.Name, nil
				}
				return nil, nil
			},
		},
	},
})
