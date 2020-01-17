package models

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/database"
)

// Category holds the title and description of a product
type Category struct {
	Base
	Title       string     `json:"name" gorm:"type:varchar(25);not null"`
	Description string     `json:"description" gorm:"type:varchar(100);not null"`
	Products    []*Product `json:"products" gorm:"many2many:product_categories;PRELOAD:false"`
}

// Categories holds an array of Category
type Categories []Category

// GetAll returns a list of all products
func (c *Categories) GetAll(p graphql.ResolveParams) error {
	db := database.Conn

	ctxLogger.Infoln("GetAll Categories")
	if result :=
		db.
			Order("title").
			Limit(p.Args["limit"].(int)).
			Offset(p.Args["offset"].(int)).
			Find(&c); result.Error != nil {
		return result.Error
	}

	return nil
}
