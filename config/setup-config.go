package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "891267034828-iaqbrrn62te9ps06esnl8j7fnqk2uok2.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:5000/api/auth/google-callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
