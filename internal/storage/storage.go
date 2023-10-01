package storage

import (
	"context"
	"database/sql"
	log "log/slog"
	"os"

	"github.com/modaniru/html-template-drawer/internal/entity"
	"github.com/modaniru/html-template-drawer/internal/storage/repos"
)

type Articles interface {
	GetCourseArticles(ctx context.Context, courseId string) ([]entity.Article, error)
	SaveArticle(ctx context.Context, article entity.ArticleForm) error
}

type Courses interface {
	GetAllCourses(ctx context.Context) ([]entity.Course, error)
	SaveCourse(ctx context.Context, course entity.CourseForm) (string, error)
}

type Storage struct {
	Articles Articles
	Courses  Courses
}

//GetCourseArticles implements Articles.
func (Storage) GetCourseArticles(ctx context.Context, courseId string) ([]entity.Article, error) {
	panic("unimplemented")
}

//SaveArticle implements Articles.
func (Storage) SaveArticle(ctx context.Context, article entity.ArticleForm) error {
	panic("unimplemented")
}

//Create all storages
func NewStorage(db *sql.DB) *Storage {
	_, err := db.Exec(`
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	create table if not exists Courses (
		id uuid DEFAULT uuid_generate_v4() unique primary key,
		image varchar,
		title varchar unique
	);

	create table if not exists Articles (
		id uuid DEFAULT uuid_generate_v4() unique primary key,
		template_name varchar,
		title varchar,
		course_id uuid REFERENCES Courses (id),
		unique(template_name, course_id)
	);`)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	return &Storage{Articles: repos.NewArticleStorage(db), Courses: repos.NewCourseStorage(db)}
}
