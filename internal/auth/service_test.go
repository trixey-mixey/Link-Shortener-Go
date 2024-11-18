package auth_test

import (
	"go/projcet-Adv/internal/auth"
	"go/projcet-Adv/internal/users"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *users.User) (*users.User, error) {
	return &users.User{
		Email: "trixey@mail.com",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*users.User, error) {
	return nil, nil
}
func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "trixey@mail.com"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "Vasya")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatal("Email dont match")
	}
}
