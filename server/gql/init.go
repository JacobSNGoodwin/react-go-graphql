package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/models"
)

// avoid declaration cycles
func init() {
	productType.AddFieldConfig("categories", &graphql.Field{
		Type:        graphql.NewList(categoryType),
		Description: "Holds a list of categories pertaining to a product",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get associated categories for given product
			if product, ok := p.Source.(models.Product); ok {
				return newProductCategoriesLoader(p.Context).Load(product.ID)
			}
			return nil, nil
		},
	})
	categoryType.AddFieldConfig("products", &graphql.Field{
		Type:        graphql.NewList(productType),
		Description: "Holds a list of products pertaining to a category",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get accosciated products for given category
			if category, ok := p.Source.(models.Category); ok {
				return newCategoryProductsLoader(p.Context).Load(category.ID)
			}
			return nil, nil
		},
	})
}
