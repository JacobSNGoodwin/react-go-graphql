package gql

import "github.com/graphql-go/graphql"

// avoid declaration cycles
func init() {
	productType.AddFieldConfig("categories", &graphql.Field{
		Type:        graphql.NewList(categoryType),
		Description: "Holds a list of categories pertaining to a product",
	})
	categoryType.AddFieldConfig("products", &graphql.Field{
		Type:        graphql.NewList(productType),
		Description: "Holds a list of products pertaining to a category",
	})
}
