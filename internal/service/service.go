package service

import (
	"context"

	"github.com/modaniru/html-template-drawer/internal/entity"
	"github.com/modaniru/html-template-drawer/internal/service/services"
	"github.com/modaniru/html-template-drawer/internal/storage"
)

type CourseService interface {
	DeleteCourse(ctx context.Context, courseId string) error
	GetAllCourses(ctx context.Context) ([]entity.Course, error)
	SaveCourse(ctx context.Context, course entity.CourseForm) (string, error)
}

type ArticleService interface {
	GetCourseArticles(ctx context.Context, courseId string) ([]entity.Article, error)
	SaveArticle(ctx context.Context, article entity.ArticleForm) error
	DeleteById(ctx context.Context, articleId string) error
}

type Service struct {
	CourseService
	ArticleService
}

// Create all services
func CreateService(storage *storage.Storage) *Service {
	return &Service{
		ArticleService: services.NewArticleService(storage),
		CourseService:  services.NewCourseService(storage),
	}
}
