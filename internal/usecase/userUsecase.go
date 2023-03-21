package usecase

import (
	"github.com/timofeyka25/test-jungle/internal/repository"
	"github.com/timofeyka25/test-jungle/pkg/hash"
	"github.com/timofeyka25/test-jungle/pkg/jwt"
	"time"
)

type UserUseCaseInterface interface {
	Login(params LoginParams) (string, error)
}

type UserUseCase struct {
	repository     repository.UserRepositoryInterface
	tokenGenerator *jwt.TokenGenerator
}

func NewUserUseCase(repository repository.UserRepositoryInterface,
	tokenGenerator *jwt.TokenGenerator) UserUseCase {
	return UserUseCase{repository: repository, tokenGenerator: tokenGenerator}
}

func (uc UserUseCase) Login(params LoginParams) (string, error) {
	user, err := uc.repository.GetByUsername(params.Username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", err
	}
	if !hash.IsEqualWithHash(params.Password, user.PasswordHash) {
		return "", err
	}
	return uc.tokenGenerator.GenerateNewAccessToken(jwt.Params{
		Id:  user.Id,
		Ttl: 12 * time.Hour,
	})
}

type LoginParams struct {
	Username string
	Password string
}
