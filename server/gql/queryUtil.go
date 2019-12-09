package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/middleware"
	"github.com/maxbrain0/react-go-graphql/server/models"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func users(p graphql.ResolveParams) (interface{}, error) {
	db := middleware.GetDB(p.Context)
	var users []models.User
	if result :=
		db.
			Order("email").
			Limit(p.Args["limit"].(int)).
			Offset(p.Args["offset"].(int)).
			Find(&users); result.Error != nil {
		return nil, nil
	}

	return users, nil
}

func user(p graphql.ResolveParams) (interface{}, error) {
	db := middleware.GetDB(p.Context)
	var user models.User

	// Find by uuid or email, which should both be unique
	if result := db.
		Where("id = ?", uuid.FromStringOrNil(p.Args["id"].(string))).
		Find(&user); result.Error != nil {
		return nil, nil
	}

	ctxLogger.WithFields(logrus.Fields{
		"ID":    user.ID,
		"Email": user.Email,
		"Name":  user.Name,
	}).Debug("Found user by id or email")

	return user, nil
}
