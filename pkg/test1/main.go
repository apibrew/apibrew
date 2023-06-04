package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "4219747860-ijobcugr71kn8kd0uq2vmiq324uglmc3.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: "http://localhost:9009/oauth2/callback",
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.

	// http://localhost:9009/oauth2/callback?state=state&code=4%2F0AbUR2VNd4NrTNphLSnJihaj-M5vGrnIYvUN-5OPDFFt7H4Z1PDODfsakdPbBgHX3ldUvaA&scope=email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+openid&authuser=0&prompt=none

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
	log.Println(resp, err)

	data, err := io.ReadAll(resp.Body)

	log.Println(string(data), err)
}
