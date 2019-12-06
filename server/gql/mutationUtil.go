package gql

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

var fbClient = &http.Client{
	Timeout: time.Second * 5,
}

// FBVerificationResponse used for getting json data response for validating respons
type FBVerificationResponse struct {
	Data FBVerificationData `json:"data"`
}

// FBVerificationData holds data used in verifying token
type FBVerificationData struct {
	AppID               string   `json:"app_id"`
	UserID              string   `json:"user_id"`
	Type                string   `json:"type"`
	Application         string   `json:"application"`
	IsValid             bool     `json:"is_valid"`
	DataAccessExpiresAt int      `json:"data_access_expires_at"`
	ExpiresAt           int      `json:"expires_at"`
	Scopes              []string `json:""`
}

// googleLoginWithToken is a helper function to verify the validity of the id_token provided by Google
func googleLoginWithToken(p graphql.ResolveParams) (interface{}, error) {
	auth, _ := GetAuth(p.Context)
	rawToken := p.Args["idToken"].(string)

	idToken, err := auth.GoogleVerifier.Verify(p.Context, rawToken)

	if err != nil {
		return false, err
	}

	var claims struct {
		Email      string `json:"email"`
		Verified   bool   `json:"email_verified"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
		Picture    string `json:"picture"`
	}

	if err := idToken.Claims(&claims); err != nil {
		return nil, err
	}

	ctxLogger.WithFields(logrus.Fields{
		"Email":      claims.Email,
		"Verified":   claims.Verified,
		"GivenName":  claims.GivenName,
		"FamilyName": claims.FamilyName,
		"Picture":    claims.Picture,
	}).Debugln("Successfully verified Google ID Token")

	return true, nil
}

func fbLoginWithToken(p graphql.ResolveParams) (interface{}, error) {
	auth, _ := GetAuth(p.Context)
	inputToken := p.Args["accessToken"].(string)
	appToken := auth.FBAccessToken

	// verify Facebook user at prescribed endpoint
	fbUserReqURL := fmt.Sprintf("https://graph.facebook.com/debug_token?input_token=%s&access_token=%s",
		inputToken,
		appToken,
	)

	ctxLogger.WithField("RequestURL", fbUserReqURL).Debug("Checking FB token validity")

	resp, err := fbClient.Get(fbUserReqURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fbData := new(FBVerificationResponse)
	json.NewDecoder(resp.Body).Decode(&fbData)

	ctxLogger.WithField("IsValid", fbData.Data.IsValid).Debug("FB verification response reveived")

	return fbData.Data.IsValid, nil
}
