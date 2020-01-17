package models

// Product holds information about a product and its price
type Product struct {
	Base
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Price       int    `json:"price" gorm:"type:integer;not null"`
	ImageURI    string `json:"imageUri" gorm:"type:text"`
}
