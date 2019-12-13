package gql

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/auth"
	"github.com/maxbrain0/react-go-graphql/server/models"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var fbClient = &http.Client{
	Timeout: time.Second * 5,
}

// FBVerificationResponse used for getting json data response for validating respons
type FBVerificationResponse struct {
	Data struct {
		IsValid             bool   `json:"is_valid"`
		AppID               string `json:"app_id"`
		UserID              string `json:"user_id"`
		DataAccessExpiresAt int    `json:"data_access_expires_at"`
		ExpiresAt           int    `json:"expires_at"`
	} `json:"data"`
}

// FBUserResponse holds profile information used for creating initial user on our site
type FBUserResponse struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
}

// GoogleIDClaims holds data from Google ID token
type GoogleIDClaims struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// googleLoginWithToken is a helper function to verify the validity of the id_token provided by Google
func googleLoginWithToken(p graphql.ResolveParams) (interface{}, error) {
	rawToken := p.Args["idToken"].(string)

	idToken, err := auth.GoogleVerifier.Verify(p.Context, rawToken)

	if err != nil {
		return false, err
	}

	var claims GoogleIDClaims

	if err := idToken.Claims(&claims); err != nil {
		return nil, err
	}

	ctxLogger.WithFields(logrus.Fields{
		"Email":   claims.Email,
		"Name":    claims.Name,
		"Picture": claims.Picture,
	}).Debugln("Successfully verified Google ID Token")

	user := models.User{
		Email:    claims.Email,
		Name:     claims.Name,
		ImageURI: claims.Picture,
	}

	loginErr := user.LoginOrCreate(p)

	if loginErr != nil {
		ctxLogger.WithFields(logrus.Fields{
			"Email":   user.Email,
			"Message": loginErr,
		}).Warn("Unable to login user")
		return err, nil
	}

	return user, nil
}

// fbLoginWithToken is a helper function to verify the validity of the access token provided by FB
// this token is not the same as the ID token. We also verify this token with FB via and http req
func fbLoginWithToken(p graphql.ResolveParams) (interface{}, error) {
	userToken := p.Args["accessToken"].(string)
	appToken := auth.FBAccessToken

	// verify Facebook user at prescribed endpoint
	fbTokenReqURL := fmt.Sprintf("https://graph.facebook.com/debug_token?input_token=%s&access_token=%s",
		userToken,
		appToken,
	)

	respToken, err := fbClient.Get(fbTokenReqURL)

	if err != nil {
		return nil, err
	}

	var fbTokenData FBVerificationResponse

	json.NewDecoder(respToken.Body).Decode(&fbTokenData)

	ctxLogger.WithFields(logrus.Fields{
		"IsValid": fbTokenData.Data.IsValid,
	}).Debugln("Successfully verified FB access token validity")

	// make sure token is Valid
	if !fbTokenData.Data.IsValid || (os.Getenv("FACEBOOK_CLIENT_ID") != fbTokenData.Data.AppID) {
		return nil, fmt.Errorf("Facebook access token is invalid for this application")
	}

	respToken.Body.Close()

	// verify the user
	fbUserReqURL := fmt.Sprintf("https://graph.facebook.com/v5.0/me?fields=name,email,picture{url}&access_token=%v",
		userToken,
	)

	respUser, err := fbClient.Get(fbUserReqURL)
	if err != nil {
		return nil, err
	}

	defer respUser.Body.Close()

	var fbUserData FBUserResponse
	json.NewDecoder(respUser.Body).Decode(&fbUserData)

	user := models.User{
		Email:    fbUserData.Email,
		Name:     fbUserData.Name,
		ImageURI: fbUserData.Picture.Data.URL,
	}

	// create jwt and send cookie
	loginErr := user.LoginOrCreate(p)

	if loginErr != nil {
		ctxLogger.WithFields(logrus.Fields{
			"Email":   user.Email,
			"Message": loginErr,
		}).Warn("Unable to login user")
		return err, nil
	}


	return user, nil
}

func createUser(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	user := p.Args["user"].(map[string]interface{})
	inputRoles := user["roles"].([]string)

	modelRoles := 

	return models.User{
		Base: models.Base{
			ID: uuid.FromStringOrNil(id),
		},
		Name: user["name"].(string),
	}, nil
}
