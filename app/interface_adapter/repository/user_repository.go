package interface_adapter

import (
	"context"
	"time"

	"github.com/fyk7/code-snippets-app/app/domain/model"
	"github.com/fyk7/code-snippets-app/app/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		Conn: db,
	}
}

func (sr *userRepository) GetAll(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	query := `SELECT * FROM user`
	if err := sr.Conn.Raw(query).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (sr *userRepository) GetByID(ctx context.Context, userID uint64) (model.User, error) {
	snippet := model.User{}
	query := `
	Select
	  *
	FROM
	  user
	WHERE
	  user_id = @userID
	`
	bindParams := map[string]interface{}{
		"userID": userID,
	}
	if err := sr.Conn.Raw(query, bindParams).Scan(&snippet).Error; err != nil {
		return model.User{}, err
	}

	return snippet, nil
}

func (sr *userRepository) FindByName(ctx context.Context, userName string) ([]model.User, error) {
	users := []model.User{}
	query := `
	Select
	  *
	FROM
	  user
	WHERE
	  user_name LIKE '%@userName%'
	`
	bindParams := map[string]interface{}{
		"userName": userName,
	}
	if err := sr.Conn.Raw(query, bindParams).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (sr *userRepository) Create(ctx context.Context, user model.User) error {
	query := `
	INSERT INTO snippet (
	  user_name,
	  is_superuser,
	  email,
	  created_at,
	  updated_at,
	) VALUES (
	  @userName,
	  @isSuperUser,
	  @email,
	  @now,
	  @now,
	);
	`
	bindParams := map[string]interface{}{
		"userName":    user.UserName,
		"isSuperUser": user.IsSuperUser,
		"email":       user.Email,
		"now":         time.Now(),
	}
	if err := sr.Conn.Exec(query, bindParams).Error; err != nil {
		return err
	}

	return nil
}

func (sr *userRepository) Update(ctx context.Context, user model.User) error {
	query := `
	UPDATE
	  user
	SET
	  user_name = @userName,
	  is_superuser = @IsSuperUser,
	  email = @email,
	  updated_at = @updatedAt,
	  updated_by = @userID
	WHERE
	  user_id = @userID;
	`
	bindParams := map[string]interface{}{
		"userID":      user.UserID,
		"userName":    user.UserName,
		"IsSuperUser": user.IsSuperUser,
		"email":       user.Email,
		"updatedAt":   time.Now(),
	}
	if err := sr.Conn.Exec(query, bindParams).Error; err != nil {
		return err
	}

	return nil
}
