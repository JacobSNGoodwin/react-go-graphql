package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/middleware"
	"github.com/maxbrain0/react-go-graphql/server/models"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func user(p graphql.ResolveParams) (interface{}, error) {
	db := middleware.GetDB(p.Context)
	var user models.User

	// Find by uuid or email, which should both be unique
	if result := db.
		Preload("Roles").
		Where("id = ?", uuid.FromStringOrNil(p.Args["id"].(string))).
		Find(&user); result.Error != nil {
		return nil, nil
	}

	ctxLogger.WithFields(logrus.Fields{
		"ID":    user.ID,
		"Email": user.Email,
		"Name":  user.Name,
		"Roles": user.Roles,
	}).Debugln("Found user by id")

	for _, role := range user.Roles {
		ctxLogger.WithFields(logrus.Fields{
			"UserID":   user.ID,
			"RoleID":   role.ID,
			"RoleName": role.Name,
		}).Debugln("User Roles")
	}

	return user, nil
}
