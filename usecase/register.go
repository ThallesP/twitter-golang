package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/thallesp/twitter-golang/entity"
	"github.com/thallesp/twitter-golang/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	FullName        string
	Username        string
	Email           string
	Password        string
	ProfileImageURL string
}

type RegisterResponse struct {
	*entity.User
}

type RegisterUseCase struct {
	usersRepository repository.UsersRepository
}

func NewRegisterUseCase(usersRepository repository.UsersRepository) *RegisterUseCase {
	return &RegisterUseCase{
		usersRepository: usersRepository,
	}
}

func (r *RegisterUseCase) Execute(registerInput *RegisterInput) (*RegisterResponse, error) {
	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), 10)

	if err != nil {
		return nil, NewException("Couldn't encrypt your password.", 500, "passwordHashFailed")
	}

	userExists, _ := r.usersRepository.FindByEmail(registerInput.Email)

	if userExists != nil {
		return nil, NewException("E-mail already exists", 400, "emailAlreadyExists")
	}

	user := &entity.User{
		Id:              uuid.NewString(),
		Email:           registerInput.Email,
		FullName:        registerInput.FullName,
		PasswordHash:    string(passwordHashBytes),
		Username:        registerInput.Username,
		ProfileImageURL: registerInput.ProfileImageURL,
		CreatedAt:       time.Now(),
	}

	err = r.usersRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
		User: user,
	}, nil
}
