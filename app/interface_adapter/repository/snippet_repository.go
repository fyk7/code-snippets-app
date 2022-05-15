package interface_adapter

import (
	"context"
	"strconv"
	"time"

	"github.com/fyk7/code-snippets-app/app/domain/model"
	"github.com/fyk7/code-snippets-app/app/domain/repository"
	"gorm.io/gorm"
)

type snippetRepository struct {
	Conn *gorm.DB
}

func NewSnippetRepository(db *gorm.DB) repository.SnippetRepository {
	return &snippetRepository{
		Conn: db,
	}
}

func (sr *snippetRepository) GetAll(ctx context.Context) ([]model.Snippet, error) {
	snippets := []model.Snippet{}
	query := `SELECT * FROM snippet`
	if err := sr.Conn.Raw(query).Scan(&snippets).Error; err != nil {
		return nil, err
	}

	return snippets, nil
}

func (sr *snippetRepository) GetByID(ctx context.Context, snippetID uint64) (model.Snippet, error) {
	snippet := model.Snippet{}
	query := `
	Select
	  *
	FROM
	  snippet
	WHERE
	  snippet_id = @snippetID
	`
	bindParams := map[string]interface{}{
		"snippetID": snippetID,
	}
	if err := sr.Conn.Raw(query, bindParams).Scan(&snippet).Error; err != nil {
		return model.Snippet{}, err
	}

	return snippet, nil
}

func (sr *snippetRepository) FindByKeyWord(ctx context.Context, keyword string) ([]model.Snippet, error) {
	snippets := []model.Snippet{}
	query := `
	Select
	  *
	FROM
	  snippet
	WHERE
	  title LIKE '%@keyword%'
	OR
	  description LIKE '%@keyword%'
	`
	bindParams := map[string]interface{}{
		"keyword": keyword,
	}
	if err := sr.Conn.Raw(query, bindParams).Scan(&snippets).Error; err != nil {
		return nil, err
	}

	return snippets, nil
}

func (sr *snippetRepository) FindByTag(ctx context.Context, tagID uint64) ([]model.Snippet, error) {
	snippets := []model.Snippet{}
	query := `
	Select
	  *
	FROM
	  snippet s
	INNER JOIN snippet_tag_relation r
	  ON s.snippet_id = r.snippet_id
	WHERE
	  r.tag_id = @tagID;
	`
	bindParams := map[string]interface{}{
		"tagID": tagID,
	}
	if err := sr.Conn.Raw(query, bindParams).Scan(&snippets).Error; err != nil {
		return nil, err
	}

	return snippets, nil
}

func (sr *snippetRepository) AssociateWithTag(ctx context.Context, snippetID, tagID, userID int64) error {
	query := `
	INSERT INTO snippet_tag_relation (
	  snippet_id,
	  tag_id,
	  created_at,
	  created_by,
	  updated_at,
	  updated_by
	) VALUES (
	  @snippetID,
	  @tagID,
	  @now,
	  @UserID,
	  @now,
	  @UserID
	);
	`
	now := time.Now()
	bindParams := map[string]interface{}{
		"snippetID": snippetID,
		"tagID":     tagID,
		"now":       now,
		"UserID":    strconv.Itoa(int(userID)),
	}
	if err := sr.Conn.Exec(query, bindParams).Error; err != nil {
		return err
	}

	return nil
}

func (sr *snippetRepository) Create(ctx context.Context, snippet model.Snippet, UserID uint64) error {
	query := `
	INSERT INTO snippet (
	  title,
	  description,
	  body,
	  programing_language,
	  created_at,
	  created_by,
	  updated_at,
	  updated_by
	) VALUES (
	  @title,
	  @description,
	  @body,
	  @programingLanguage,
	  @now,
	  @UserID,
	  @now,
	  @UserID
	);
	`
	now := time.Now()
	bindParams := map[string]interface{}{
		"title":              snippet.Title,
		"description":        snippet.Description,
		"body":               snippet.Body,
		"programingLanguage": snippet.ProgramingLanguage,
		"now":                now,
		"UserID":             strconv.Itoa(int(UserID)),
	}
	if err := sr.Conn.Exec(query, bindParams).Error; err != nil {
		return err
	}

	return nil
}

func (sr *snippetRepository) Update(ctx context.Context, snippet model.Snippet, UserID uint64) error {
	query := `
	UPDATE
	  snippet
	SET
	  title = @title,
	  description = @description,
	  body = @body,
	  programing_language = @programing_language,
	  updated_at = @now,
	  updated_by = @UserID
	WHERE
	  snippet_id = @id;
	`
	bindParams := map[string]interface{}{
		"id":                 snippet.SnippetID,
		"title":              snippet.Title,
		"description":        snippet.Description,
		"body":               snippet.Body,
		"programingLanguage": snippet.ProgramingLanguage,
		"now":                time.Now(),
		"UserID":             strconv.Itoa(int(UserID)),
	}
	if err := sr.Conn.Exec(query, bindParams).Error; err != nil {
		return err
	}

	return nil
}

func (sr *snippetRepository) Delete(ctx context.Context, id uint64) error {
	query := `
	DELETE FROM snippet WHERE snippet_id = @id
	`
	bindParams := map[string]interface{}{
		"id": id,
	}
	if err := sr.Conn.Exec(query, bindParams).Error; err != nil {
		return err
	}

	return nil
}
