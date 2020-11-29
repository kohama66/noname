package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var authClient *auth.Client

func init() {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	authClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// VerifyIDToken tokenIDを確認
func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return authClient.VerifyIDToken(ctx, idToken)
}
