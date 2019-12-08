package gql

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/maxbrain0/react-go-graphql/server/middleware"
	"github.com/maxbrain0/react-go-graphql/server/models"
	"github.com/sirupsen/logrus"
)

var fbClient = &http.Client{
	Timeout: time.Second * 5,
}

// FBVerificationResponse used for getting json data response for validating respons
type FBVerificationResponse struct {
	Data struct {
		IsValid             bool `json:"is_valid"`
		DataAccessExpiresAt int  `json:"data_access_expires_at"`
		ExpiresAt           int  `json:"expires_at"`
	} `json:"data"`
}

// GoogleIDClaims holds data from Google ID token
type GoogleIDClaims struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// googleLoginWithToken is a helper function to verify the validity of the id_token provided by Google
func googleLoginWithToken(p graphql.ResolveParams) (interface{}, error) {
	auth := middleware.GetAuth(p.Context)
	rawToken := p.Args["idToken"].(string)

	idToken, err := auth.GoogleVerifier.Verify(p.Context, rawToken)

	if err != nil {
		return false, err
	}

	var claims = new(GoogleIDClaims)

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
		ctxLogger.WithField("Email", user.Email).Warn("Unable to login user")
		return err, nil
	}

	return user, nil
}

// fbLoginWithToken is a helper function to verify the validity of the access token provided by FB
// this token is not the same as the ID token. We also verify this token with FB via and http req
//Therefore, we receive email, name, and picture as parameters to this mutation
func fbLoginWithToken(p graphql.ResolveParams) (interface{}, error) {
	auth := middleware.GetAuth(p.Context)
	inputData := p.Args["fbLoginData"].(map[string]interface{})
	inputToken := inputData["token"].(string)
	email := inputData["email"].(string)
	name := inputData["name"].(string)
	imageURI := inputData["imageUri"].(string)

	appToken := auth.FBAccessToken

	// ctxLogger.WithField("Token", inputToken).Debugln("Input token received as argument")

	// verify Facebook user at prescribed endpoint
	fbUserReqURL := fmt.Sprintf("https://graph.facebook.com/debug_token?input_token=%s&access_token=%s",
		inputToken,
		appToken,
	)

	resp, err := fbClient.Get(fbUserReqURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fbData := new(FBVerificationResponse)

	json.NewDecoder(resp.Body).Decode(&fbData)

	ctxLogger.WithFields(logrus.Fields{
		"IsValid": fbData.Data.IsValid,
	}).Debugln("Successfully verified FB access token validity")

	if !fbData.Data.IsValid {
		return nil, fmt.Errorf("Facebook access token is invalid")
	}

	user := models.User{
		Email:    email,
		Name:     name,
		ImageURI: imageURI,
	}

	loginErr := user.LoginOrCreate(p)

	if loginErr != nil {
		ctxLogger.WithField("Email", user.Email).Warn("Unable to login user")
		return err, nil
	}

	return user, nil
}
