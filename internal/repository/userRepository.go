package repository

import (
	"context"
	"time"

	"github.com/Angstreminus/cinema/internal/apperrors"
	"github.com/Angstreminus/cinema/internal/dto"
	"github.com/Angstreminus/cinema/internal/entity"
	"github.com/Angstreminus/cinema/logger"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	Db     *sqlx.DB
	Logger *logger.Logger
}

func NewUserRepository(db *sqlx.DB, log *logger.Logger) *UserRepository {
	return &UserRepository{
		Db:     db,
		Logger: log,
	}
}

func (ur *UserRepository) RegistrateUser(ctx *context.Context, user *entity.User) (*entity.User, apperrors.AppError) {
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05.999")
	query := `INSERT INTO users VALUES(id, login, role, hashed_password, is_deleted, created_at) $1, $2, $3, $4, $5, $6 RETURNING *;`
	stmt, err := ur.Db.PrepareContext(*ctx, query)
	if err != nil {
		ur.Logger.ZapLogger.Error("Preparing statement error")
		return nil, &apperrors.DBoperationErr{
			Message: err.Error(),
		}
	}
	var usr entity.User
	err = stmt.QueryRowContext(*ctx, &user).Scan(&usr)
	if err != nil {
		ur.Logger.ZapLogger.Error("Create query error")
		return nil, &apperrors.DBoperationErr{
			Message: err.Error(),
		}
	}
	ur.Logger.ZapLogger.Info("User created")
	return &usr, nil
}

func (ur *UserRepository) IsUserExists(ctx *context.Context, signature *dto.UserSignature) (bool, apperrors.AppError) {
	query := `SELECT EXIST(SELECT 1 FROM users WHERE login = $1 AND hashed_password = $2);`
	stmt, err := ur.Db.PrepareContext(*ctx, query)
	if err != nil {
		ur.Logger.ZapLogger.Error("Preparing statement error")
		return true, &apperrors.DBoperationErr{
			Message: err.Error(),
		}
	}
	var exist bool
	err = stmt.QueryRowContext(*ctx, &signature.Login, &signature.Password).Scan(&exist)
	if err != nil {
		ur.Logger.ZapLogger.Error("Query error")
		return true, &apperrors.DBoperationErr{
			Message: err.Error(),
		}
	}
	ur.Logger.ZapLogger.Info("User does not exist")
	return exist, nil
}
