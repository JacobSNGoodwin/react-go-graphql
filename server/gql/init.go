package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/models"
)

func init() {
	productType.AddFieldConfig("categories", &graphql.Field{
		Type:        graphql.NewList(categoryType),
		Description: "Holds a list of categories pertaining to a product",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get associated categories for given product
			if product, ok := p.Source.(models.Product); ok {
				thunk := p.Context.Value(models.ContextKeyLoaders).(models.Loaders).
					ProductCategoriesLoader.LoadThunk(product.ID)
				return func() (interface{}, error) {
					return thunk()
				}, nil
			}
			return nil, nil
		},
	})
	categoryType.AddFieldConfig("products", &graphql.Field{
		Type:        graphql.NewList(productType),
		Description: "Holds a list of products pertaining to a category",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if category, ok := p.Source.(models.Category); ok {
				thunk := p.Context.Value(models.ContextKeyLoaders).(models.Loaders).
					CategoryProductsLoader.LoadThunk(category.ID)
				return func() (interface{}, error) {
					return thunk()
				}, nil
			}
			return nil, nil
		},
	})
}
