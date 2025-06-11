package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseService struct {
	client *auth.Client
}

func NewFirebaseService(service_path string) *FirebaseService {
	opt := option.WithCredentialsFile(service_path)
	app, _ := firebase.NewApp(context.Background(), nil, opt)
	client, _ := app.Auth(context.Background())
	return &FirebaseService{client: client}
}

func (f *FirebaseService) VerifyByID(ctx context.Context, idToken string) (*auth.Token, error) {
	return f.client.VerifyIDToken(ctx, idToken)
}
