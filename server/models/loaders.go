package models

import (
	"time"

	"github.com/maxbrain0/react-go-graphql/server/database"
	"github.com/maxbrain0/react-go-graphql/server/errors"
	uuid "github.com/satori/go.uuid"
)

// Loaders used to store loaders for access on the middleware context
type Loaders struct {
	ProductCategoriesLoader *CategoriesLoader
	CategoryProductsLoader  *ProductsLoader
}

// ProductCategory used to scan data from a single query using a dataloader
type ProductCategory struct {
	ProductID uuid.UUID
	Category
}

// ProductCategories is a slice of ProductCategory
type ProductCategories []ProductCategory

// CategoryProduct used to scan data from a single query using a dataloader
type CategoryProduct struct {
	CategoryID uuid.UUID
	Product
}

// CategoryProducts is a slice of CategoryProduct
type CategoryProducts []CategoryProduct

// NewProductCategoriesLoader returns a Categories loader that access categories beloning to a product
func NewProductCategoriesLoader() *CategoriesLoader {
	var db = database.Conn
	return &CategoriesLoader{
		wait:     5 * time.Millisecond,
		maxBatch: 100,
		fetch: func(ids []uuid.UUID) ([][]Category, []error) {
			output := make([][]Category, len(ids))
			outputErrors := make([]error, len(ids))

			rows, e := db.
				Raw("SELECT * FROM categories JOIN product_categories ON product_categories.category_id = id WHERE product_id IN (?)", ids).
				Rows()

			// also can add map of product ID to errors
			productCategories := make(map[uuid.UUID][]Category, len(ids))
			defer rows.Close()
			for rows.Next() {
				productCategory := ProductCategory{}
				db.ScanRows(rows, &productCategory)
				category := productCategory.Category

				productCategories[productCategory.ProductID] = append(productCategories[productCategory.ProductID], category)
			}

			for i, id := range ids {
				if e != nil {
					outputErrors[i] = errors.NewInternal("Failed to load categories for products", nil)
				}
				outputID, ok := productCategories[id]
				if !ok {
					output[i] = []Category{}
				} else {
					output[i] = outputID
				}
			}

			return output, outputErrors
		},
	}
}

// NewCategoryProductsLoader returns a Products loader that access products beloning to a category
func NewCategoryProductsLoader() *ProductsLoader {
	var db = database.Conn
	return &ProductsLoader{
		wait:     50 * time.Millisecond,
		maxBatch: 100,
		fetch: func(ids []uuid.UUID) ([][]Product, []error) {
			output := make([][]Product, len(ids))
			outputErrors := make([]error, len(ids))

			rows, e := db.
				Raw("SELECT * FROM products JOIN product_categories ON product_categories.product_id = id WHERE category_id IN (?)", ids).
				Rows()

			// also can add map of product ID to errors
			categoryProducts := make(map[uuid.UUID][]Product, len(ids))
			defer rows.Close()
			for rows.Next() {
				categoryProduct := CategoryProduct{}
				db.ScanRows(rows, &categoryProduct)
				product := categoryProduct.Product

				categoryProducts[categoryProduct.CategoryID] = append(categoryProducts[categoryProduct.CategoryID], product)
			}

			for i, id := range ids {
				if e != nil {
					outputErrors[i] = errors.NewInternal("Failed to load products for categories", nil)
				}
				outputID, ok := categoryProducts[id]
				if !ok {
					output[i] = []Product{}
				} else {
					output[i] = outputID
				}
			}

			return output, outputErrors
		},
	}
}
