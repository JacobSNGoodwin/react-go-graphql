package models

import (
	"time"

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
	// var db = database.Conn
	return &CategoriesLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(ids []uuid.UUID) ([][]Category, []error) {
			categories := make([][]Category, len(ids))
			errors := make([]error, len(ids))

			ctxLogger.WithField("ids", ids).Infoln("Product Categories Fetch")

			// var data ProductCategories
			// db.
			// 	Raw("SELECT * FROM categories JOIN product_categories ON product_categories.category_id = id WHERE product_id IN (?)", ids).
			// 	Scan(&data)

			// for _, row := range data {
			// 	productID := row.ProductID
			// }

			return categories, errors
		},
	}
}

// NewCategoryProductsLoader returns a Products loader that access products beloning to a category
func NewCategoryProductsLoader() *ProductsLoader {
	// var db = database.Conn
	return &ProductsLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(ids []uuid.UUID) ([][]Product, []error) {
			products := make([][]Product, len(ids))
			errors := make([]error, len(ids))

			ctxLogger.WithField("ids", ids).Infoln("Category Products Fetch")

			// var data CategoryProducts

			// db.
			// 	Raw("SELECT * FROM products JOIN product_categories ON product_categories.product_id = products.id WHERE category_id IN (?)", ids).
			// 	Scan(&data)

			// for _, product := range data {
			// 	ctxLogger.WithFields(logrus.Fields{
			// 		"CategoryID": product.CategoryID,
			// 		"ProductID":  product.ID,
			// 	}).Infoln("Product Row")
			// }

			return products, errors
		},
	}
}
