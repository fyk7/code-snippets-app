package interface_adapter

import (
	"context"
	"strconv"
	"time"

	"github.com/fyk7/code-snippets-app/app/domain/model"
	"github.com/fyk7/code-snippets-app/app/domain/repository"
	"gorm.io/gorm"
)

type tagRepository struct {
	Conn *gorm.DB
}

func NewTagRepository(db *gorm.DB) repository.TagRepository {
	return &tagRepository{
		Conn: db,
	}
}

func (tr *tagRepository) GetAll(ctx context.Context) ([]model.Tag, error) {
	tags := []model.Tag{}
	query := `SELECT * FROM tag`
	if err := tr.Conn.Raw(query).Scan(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (tr *tagRepository) GetByID(ctx context.Context, tagID uint64) (model.Tag, error) {
	tag := model.Tag{}
	query := `
	Select
	  *
	FROM
	  snippet
	WHERE
	  tag_id = @tagID
	`
	bindParams := map[string]interface{}{
		"tagID": tagID,
	}
	if err := tr.Conn.Raw(query, bindParams).Scan(&tag).Error; err != nil {
		return model.Tag{}, err
	}

	return tag, nil
}

func (tr *tagRepository) FindByKeyWord(ctx context.Context, keyword string) ([]model.Tag, error) {
	tags := []model.Tag{}
	query := `
	Select
	  *
	FROM
	  snippet
	WHERE
	  tag_name LIKE '%@keyword%'
	`
	bindParams := map[string]interface{}{
		"keyword": keyword,
	}
	if err := tr.Conn.Raw(query, bindParams).Scan(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (tr *tagRepository) Create(ctx context.Context, tag model.Tag, UserID uint64) error {
	query := `
	INSERT INTO tag (
	  tag_name,
	  created_at,
	  created_by,
	  updated_at,
	  updated_by
	) VALUES (
	  @tagName,
	  @now,
	  @UserID,
	  @now,
	  @UserID
	);
	`
	now := time.Now()
	bindParams := map[string]interface{}{
		"tagName": tag.TagName,
		"now":     now,
		"UserID":  strconv.Itoa(int(UserID)),
	}
	if err := tr.Conn.Exec(query, bindParams).Error; err != nil {
		return err
	}

	return nil
}
