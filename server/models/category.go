package models

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/database"
	"github.com/maxbrain0/react-go-graphql/server/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
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

// GetByID gets category from database based on the its id
func (c *Category) GetByID(p graphql.ResolveParams) error {
	db := database.Conn

	ctxLogger.WithField("id", p.Args["id"].(string)).Infoln("GetByID Categories")
	// Find by uuid or email, which should both be unique
	if err := db.
		Where("id = ?", uuid.FromStringOrNil(p.Args["id"].(string))).
		Find(&c).Error; err != nil {
		ctxLogger.WithError(err).Debugln("DB Error finding category by ID")
		return errors.NewInternal("Error finding category", nil)
	}

	return nil
}

// Create adds a new Category to the database
// If it fails, returns a Failed to create error
func (c *Category) Create(p graphql.ResolveParams) error {
	db := database.Conn
	ctxUser := p.Context.Value(ContextKeyUser).(User)

	if !hasRole(ctxUser.Roles, "admin") && !hasRole(ctxUser.Roles, "editor") {
		return errors.NewForbidden("Not authorized", nil)
	}

	ctxLogger.WithFields(logrus.Fields{
		"Title": c.Title,
	}).Infoln("Creating category")

	if err := db.Create(&c).Model(&c).Error; err != nil {
		ctxLogger.WithError(err).Debugln("DB Error creating category")
		return errors.NewInternal("Error creating category", nil)
	}

	return nil
}
