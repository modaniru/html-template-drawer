package repos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/modaniru/html-template-drawer/internal/entity"
)

type articleStorage struct {
	db *sql.DB
}

func NewArticleStorage(db *sql.DB) *articleStorage {
	return &articleStorage{db: db}
}

// Return course articles
func (a *articleStorage) GetCourseArticles(ctx context.Context, courseId string) ([]entity.Article, error) {
	rows, err := a.db.QueryContext(ctx, "select template_name, title  from Articles where course_id::uuid = $1::uuid;", courseId)
	if err != nil {
		return nil, fmt.Errorf("execute query error: %w", err)
	}
	defer rows.Close()

	articles := []entity.Article{}
	for rows.Next() {
		article := entity.Article{}
		err := rows.Scan(&article.TemplateName, &article.Title)
		if err != nil {
			return nil, fmt.Errorf("scan query error: %w", err)
		}
		articles = append(articles, article)
	}
	return articles, nil
}

// Save article
// TODO create custom error when article in course already exists
func (a *articleStorage) SaveArticle(ctx context.Context, article entity.ArticleForm) error {
	_, err := a.db.Exec("insert into Articles (template_name, title, course_id) values ($1, $2, $3::uuid);", article.TemplateName, article.Title, article.CourseId)
	if err != nil {
		return fmt.Errorf("execute query error: %w", err)
	}
	return err
}
