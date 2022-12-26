package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:5000/api/auth/google-callback",
		Scopes: []string{
			os.Getenv("USERINFO_PROFILE"),
			os.Getenv("USERINFO_EMAIL"),
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
