package models

// Category holds the title and description of a product
type Category struct {
	Base
	Title       string     `json:"name" gorm:"type:varchar(25);not null"`
	Description string     `json:"description" gorm:"type:varchar(100);not null"`
	Products    []*Product `json:"products" gorm:"many2many:product_categories"`
}
