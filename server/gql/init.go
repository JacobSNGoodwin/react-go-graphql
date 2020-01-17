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
			// needed to handle categories as a pointer (used for GORM reasons)
			var categories = models.Categories{}
			if product, ok := p.Source.(models.Product); ok {
				for _, category := range product.Categories {
					categories = append(categories, *category)
				}
				return categories, nil
			}
			return nil, nil
		},
	})
	categoryType.AddFieldConfig("products", &graphql.Field{
		Type:        graphql.NewList(productType),
		Description: "Holds a list of products pertaining to a category",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// needed to handle categories as a pointer (used for GORM reasons)
			var products = models.Products{}
			if category, ok := p.Source.(models.Category); ok {
				for _, product := range category.Products {
					products = append(products, *product)
				}
				return products, nil
			}
			return nil, nil
		},
	})
}
