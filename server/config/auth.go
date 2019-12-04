package config

import (
	"context"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/maxbrain0/react-go-graphql/server/logger"
)

var ctxLogger = logger.CtxLogger

// Auth holds a map of oauth2 configurations that are initialized with environment variables
// in the LoadConfig method
type Auth struct {
	Verifier *oidc.IDTokenVerifier
}

// LoadConfigs sets up a map of Configs in the Auth struct
func (a *Auth) LoadConfigs() {
	prodiver, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")

	if err != nil {
		ctxLogger.Fatalf("Unable to create google oid provider: %v", err.Error())
	}

	a.Verifier = prodiver.Verifier(&oidc.Config{
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
	})

}
