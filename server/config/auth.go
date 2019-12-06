package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/maxbrain0/react-go-graphql/server/logger"
)

var ctxLogger = logger.CtxLogger
var fbClient = &http.Client{
	Timeout: time.Second * 5,
}

// FBResponse holds the response to request for app access token
type FBResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// Auth holds a map of oauth2 configurations that are initialized with environment variables
// in the LoadConfig method
type Auth struct {
	GoogleVerifier *oidc.IDTokenVerifier
	FBAccessToken  string
}

// Load sets up necessary server-side auth verifications
// for 3rd party providers like Google and Facebook
func (a *Auth) Load() {
	// Get a verifier for verifying Google ID tokens
	prodiver, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")

	if err != nil {
		ctxLogger.Fatalf("Unable to create google oid provider: %v", err.Error())
	}

	a.GoogleVerifier = prodiver.Verifier(&oidc.Config{
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
	})

	// Get a token which allows us to verify facebook access tokens later on
	fbTokenURL := fmt.Sprintf("https://graph.facebook.com/oauth/access_token?client_id=%s&client_secret=%s&grant_type=client_credentials",
		os.Getenv("FACEBOOK_CLIENT_ID"),
		os.Getenv("FACEBOOK_CLIENT_SECRET"),
	)

	fbRes, err := fbClient.Get(fbTokenURL)

	if err != nil {
		ctxLogger.Fatalf("Unable to retrieve faccebook app access token: %v", err.Error())
	}

	defer fbRes.Body.Close()

	decodedFBRes := new(FBResponse)
	json.NewDecoder(fbRes.Body).Decode(&decodedFBRes)

	a.FBAccessToken = decodedFBRes.AccessToken
}
