package gql

import "github.com/graphql-go/graphql"

// RootMutation contains the main mutations for the GraphQL API
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"googleLoginWithToken": &graphql.Field{
			Type:        graphql.String,
			Description: "Receives an id_token from a client-side login to Google, and checks that this is a valid token. If so, a jwt is returned as a string",
			Args: graphql.FieldConfigArgument{
				"idToken": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				auth, _ := GetAuth(p.Context)
				rawToken := p.Args["idToken"].(string)

				idToken, err := auth.Verifier.Verify(p.Context, rawToken)

				if err != nil {
					return nil, err
				}

				var claims struct {
					Email    string `json:"email"`
					Verified bool   `json:"email_verified"`
				}

				if err := idToken.Claims(&claims); err != nil {
					return nil, err
				}

				return claims.Email, nil
			},
		},
	},
})
