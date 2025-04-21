package gql

import (
	"api-gateway/services"
	"context"
)

type Resolver struct{}

func (r *mutationResolver) Signup(ctx context.Context, email string, password string) (*services.User, error) {
	return services.Signup(email, password)
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*services.Token, error) {
	return services.Login(email, password)
}

func (r *mutationResolver) GoogleLogin(ctx context.Context, token string) (*services.Token, error) {
	return services.GoogleLogin(token)
}

func (r *queryResolver) GetUser(ctx context.Context, email string) (*services.User, error) {
	return services.GetUser(email)
}

func (r *queryResolver) GetSessions(ctx context.Context, email string) ([]*services.Session, error) {
	return services.GetSessions(email)
}
