package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/thallesp/twitter-golang/entity"
	"github.com/thallesp/twitter-golang/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginResponse struct {
	*entity.User
	Token string `json:"-"`
}

type LoginUseCase struct {
	usersRepository repository.UsersRepository
}

func NewLoginUseCase(usersReposiory repository.UsersRepository) *LoginUseCase {
	return &LoginUseCase{
		usersRepository: usersReposiory,
	}
}

func (l *LoginUseCase) Execute(loginInput LoginInput) (*LoginResponse, error) {
	user, err := l.usersRepository.FindByEmail(loginInput.Email)

	if err != nil {
		return nil, NewException("Email/password incorrect", 400, "emailPasswordIncorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginInput.Password))

	if err != nil {
		return nil, NewException("Email/password incorrect", 400, "emailPasswordIncorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(30 * 24 * time.Hour).Unix(), // 30 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, NewException("Couldn't authenticate you.", 500, "token_generation_error")
	}

	return &LoginResponse{Token: tokenString, User: user}, nil
}
