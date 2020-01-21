package models

import "github.com/sirupsen/logrus"

import "github.com/maxbrain0/react-go-graphql/server/database"

// Init assures tables for provided models are available and initialized a couple of users and roles
func Init() {
	db := database.Conn

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Category{})

	RoleMap = make(map[string]*Role)
	// create map of roles (for more easily getting role ref to use with GORM in gql mutations/queries)
	RoleMap["admin"] = &Role{
		Name: "admin",
	}
	RoleMap["editor"] = &Role{
		Name: "editor",
	}

	// Create Admin and Editor Roles
	// Can iterate over array or map if we need many roles in the future
	admin, _ := RoleMap["admin"]
	db.Where(*admin).FirstOrCreate(admin)
	ctxLogger.WithFields(logrus.Fields{
		"id":        admin.ID,
		"Name":      admin.Name,
		"UpdatedAt": admin.UpdatedAt,
	}).Debugln("Created or found role")

	editor, _ := RoleMap["editor"]
	db.Where(*editor).FirstOrCreate(editor)
	ctxLogger.WithFields(logrus.Fields{
		"id":        editor.ID,
		"Name":      editor.Name,
		"UpdatedAt": editor.UpdatedAt,
	}).Debugln("Created or found role")

	// Create users and append roles
	var user1 User
	db.FirstOrCreate(&user1, User{
		Name:     "Jacob",
		Email:    "jacob.goodwin@gmail.com",
		ImageURI: "https://lh3.googleusercontent.com/a-/AAuE7mCsAHdorySC7ttxiSQOx7xtcUHhMwX6LlJwDT65LsE=s96-c",
	}).Model(&user1).Association("Roles").Append([]Role{*RoleMap["admin"], *RoleMap["editor"]})

	ctxLogger.WithFields(logrus.Fields{
		"id":        user1.ID,
		"Name":      user1.Name,
		"UpdatedAt": user1.UpdatedAt,
		"Roles":     user1.Roles,
	}).Debugln("Created or found user")

	// Create sample categories and products
	apparel := &Category{
		Title:       "apparel",
		Description: "Clothing, textiles, and all things wearable",
	}
	footwear := &Category{
		Title:       "footwear",
		Description: "Things you use to cover ya dadgummed feet",
	}

	db.Where(*apparel).FirstOrCreate(apparel)
	ctxLogger.WithFields(logrus.Fields{
		"id":        apparel.ID,
		"Title":     apparel.Title,
		"UpdatedAt": apparel.UpdatedAt,
	}).Debugln("Created or found category")

	db.Where(*footwear).FirstOrCreate(footwear)
	ctxLogger.WithFields(logrus.Fields{
		"id":        footwear.ID,
		"Title":     footwear.Title,
		"UpdatedAt": footwear.UpdatedAt,
	}).Debugln("Created or found category")

	// Create products and append categories
	var product1 Product
	db.FirstOrCreate(&product1, Product{
		Name:        "Swoosh125",
		Description: "Totally Nike knock-offs",
		Price:       5999,
		ImageURI:    "https://skitterphoto.com/photos/skitterphoto-1480-default.jpg",
		Location:    "A15",
	}).Model(&product1).Association("Categories").Append([]Category{*apparel, *footwear})

	ctxLogger.WithFields(logrus.Fields{
		"id":         product1.ID,
		"Name":       product1.Name,
		"UpdatedAt":  product1.UpdatedAt,
		"Categories": product1.Categories,
	}).Debugln("Created or found product")

	// Create products and append categories
	var product2 Product
	db.FirstOrCreate(&product2, Product{
		Name:        "Bidniz Pants",
		Description: "Taking down the SV boys club kinda pants, and offending someone with this description!",
		Price:       8599,
		ImageURI:    "https://p1.pxfuel.com/preview/879/790/566/fashion-textile-clothes-daily-current-season-white-royalty-free-thumbnail.jpg",
		Location:    "Y15",
	}).Model(&product2).Association("Categories").Append([]Category{*apparel})

	ctxLogger.WithFields(logrus.Fields{
		"id":         product2.ID,
		"Name":       product2.Name,
		"UpdatedAt":  product2.UpdatedAt,
		"Categories": product2.Categories,
	}).Debugln("Created or found product")
}
