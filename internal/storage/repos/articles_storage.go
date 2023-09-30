package repos

import (
	"context"
	"database/sql"

	"github.com/modaniru/html-template-drawer/internal/entity"
)

type articleStorage struct {
	db *sql.DB
}

func NewArticleStorage(db *sql.DB) *articleStorage {
	return &articleStorage{db: db}
}

func (a *articleStorage) GetCourseArticles(ctx context.Context, courseId string) ([]entity.Article, error) {
	rows, err := a.db.QueryContext(ctx, "select name from Articles where course_id::uuid = $1::uuid;", courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []entity.Article{}
	for rows.Next() {
		article := entity.Article{}
		err := rows.Scan(&article.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, article)
	}
	return res, nil
}
