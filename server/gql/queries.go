package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/data"
)

// GetRootQuery returns the root query with the datasource plugged into it
func GetRootQuery(ds *data.Data) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "A list of all users",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return ds.Users, nil
				},
			},
		},
	})
}
