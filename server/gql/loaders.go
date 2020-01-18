package gql

import (
	"context"
	"time"

	"github.com/maxbrain0/react-go-graphql/server/database"
	"github.com/maxbrain0/react-go-graphql/server/models"
	uuid "github.com/satori/go.uuid"
)

func newProductCategoriesLoader(ctx context.Context) *ProductCategoriesLoader {
	var db = database.Conn
	return &ProductCategoriesLoader{
		wait:     10 * time.Millisecond,
		maxBatch: 100,
		fetch: func(ids []uuid.UUID) ([][]models.Category, []error) {
			categories := make([][]models.Category, len(ids))
			errors := make([]error, len(ids))

			for i, id := range ids {
				var product models.Product
				var productCategories models.Categories
				product.ID = id

				if e := db.Model(&product).Association("Categories").Find(&productCategories).Error; e != nil {
					errors[i] = e
				}

				categories[i] = productCategories
			}
			return categories, errors
		},
	}
}

func newCategoryProductsLoader(ctx context.Context) *CategoryProductsLoader {
	var db = database.Conn
	return &CategoryProductsLoader{
		wait:     10 * time.Millisecond,
		maxBatch: 100,
		fetch: func(ids []uuid.UUID) ([][]models.Product, []error) {
			products := make([][]models.Product, len(ids))
			errors := make([]error, len(ids))

			for i, id := range ids {
				var category models.Category
				var categoryProducts models.Products
				category.ID = id

				if e := db.Model(&category).Association("Products").Find(&categoryProducts).Error; e != nil {
					errors[i] = e
				}

				products[i] = categoryProducts
			}
			return products, errors
		},
	}
}
