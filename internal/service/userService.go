package service

import (
	"context"

	"github.com/Angstreminus/cinema/internal/apperrors"
	"github.com/Angstreminus/cinema/internal/dto"
	"github.com/Angstreminus/cinema/internal/entity"
	"github.com/Angstreminus/cinema/internal/repository"
	"github.com/Angstreminus/cinema/logger"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository *repository.UserRepository
	Logger         *logger.Logger
}

func NewUserService(logger *logger.Logger, repo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
		Logger:         logger,
	}
}

func (us *UserService) RegiterUser(toRegistrate *dto.RegisterRequest) (*entity.User, apperrors.AppError) {
	var signature dto.UserSignature
	signature.Login = toRegistrate.Login
	hasedPassword, err := HashPassword(toRegistrate.Password)
	if err != nil {
		us.Logger.ZapLogger.Error("Error to hash password")
		return nil, &apperrors.HashError{
			Message: err.Error(),
		}
	}
	signature.Password = hasedPassword
	ctx := context.Background()
	exists, err := us.UserRepository.IsUserExists(&ctx, &signature)
	if !exists {
		var user entity.User
		user.Id = uuid.New()
		user.Login = toRegistrate.Login
		user.Password = hasedPassword
		// admin role created manually via sql query
		user.Role = "user"
		user.IsDeleted = false
		return us.UserRepository.RegistrateUser(&ctx, &user)
	} else {
		return nil, &apperrors.AlreadyExistsError{
			Message: "User already exists",
		}
	}
}

func (us *UserService) IsUserExists(ctx *context.Context, signature *dto.UserSignature) (bool, apperrors.AppError) {
	return us.UserRepository.IsUserExists(ctx, signature)

}
