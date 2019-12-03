package config

import (
	"os"

	"golang.org/x/oauth2"
)

// Auth holds a map of oauth2 configurations that are initialized with environment variables
// in the LoadConfig method
type Auth struct {
	Configs map[string]*oauth2.Config
}

// LoadConfigs sets up a map of Configs in the Auth struct
func (a *Auth) LoadConfigs() {
	// initialize map
	a.Configs = map[string]*oauth2.Config{}

	a.Configs["google"] = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	a.Configs["facebook"] = &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		Scopes:       []string{"email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/v5.0/dialog/oauth",
			TokenURL: "https://graph.facebook.com/v5.0/oauth/access_token",
		},
	}
}
