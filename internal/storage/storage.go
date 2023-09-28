package storage

import (
	"context"
	"database/sql"

	"github.com/modaniru/html-template-drawer/internal/entity"
	"github.com/modaniru/html-template-drawer/internal/storage/repos"
)

type Articles interface {
	GetCourseArticles(ctx context.Context, courseId string) ([]entity.Article, error)
}

type Courses interface {
	GetAllCourses(ctx context.Context) ([]entity.Course, error)
}

type Storage struct{
	Articles Articles
	Cources Courses
}

func NewStorage(db *sql.DB) *Storage{
	return &Storage{Articles: repos.NewArticleStorage(db), Cources: repos.NewCourseStorage(db)}
}