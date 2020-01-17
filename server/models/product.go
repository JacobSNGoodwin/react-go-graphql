package models

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/database"
)

// Product holds information about a product and its price
type Product struct {
	Base
	Name        string      `json:"name" gorm:"type:varchar(100);not null"`
	Description string      `json:"description" gorm:"type:text;not null"`
	Price       int         `json:"price" gorm:"type:integer;not null"`
	ImageURI    string      `json:"imageUri" gorm:"type:text"`
	Location    string      `json:"location" grom:"type:varchar(6);not null"`
	Categories  []*Category `json:"categories" gorm:"many2many:product_categories"`
}

// Products holds an array of Product
type Products []Product

// GetAll returns a list of all products
func (pr *Products) GetAll(p graphql.ResolveParams) error {
	db := database.Conn

	ctxLogger.Infoln("GetAll Products")
	if result :=
		db.
			Order("name").
			Limit(p.Args["limit"].(int)).
			Offset(p.Args["offset"].(int)).
			Preload("Categories").
			Find(&pr); result.Error != nil {
		return result.Error
	}

	return nil
}
