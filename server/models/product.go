package models

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/database"
	"github.com/maxbrain0/react-go-graphql/server/errors"
	uuid "github.com/satori/go.uuid"
)

// Product holds information about a product and its price
type Product struct {
	Base
	Name        string      `json:"name" gorm:"type:varchar(100);not null"`
	Description string      `json:"description" gorm:"type:text;not null"`
	Price       int         `json:"price" gorm:"type:integer;not null"`
	ImageURI    string      `json:"imageUri" gorm:"type:text"`
	Location    string      `json:"location" grom:"type:varchar(6);not null"`
	Categories  []*Category `json:"categories" gorm:"many2many:product_categories;PRELOAD:false"`
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
			Find(&pr); result.Error != nil {
		return errors.NewInternal("Error fetching products", result.Error)
	}

	return nil
}

// GetByID gets product from database based on the its id
func (pr *Product) GetByID(p graphql.ResolveParams) error {
	db := database.Conn
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") && !hasRole(ctxUser.Roles, "editor") {
		return errors.NewForbidden("Not authorized", nil)
	}

	ctxLogger.WithField("id", p.Args["id"].(string)).Infoln("GetByID Products")
	// Find by uuid or email, which should both be unique
	if err := db.
		Where("id = ?", uuid.FromStringOrNil(p.Args["id"].(string))).
		Find(&pr).Error; err != nil {
		ctxLogger.WithError(err).Debugln("DB Error finding user by ID")
		return errors.NewInternal("Error finding user", nil)
	}

	return nil
}
