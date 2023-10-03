package services

import (
	"context"
	"fmt"

	"github.com/modaniru/html-template-drawer/internal/entity"
	"github.com/modaniru/html-template-drawer/internal/storage"
)

type ArticleService struct {
	articleStorage storage.Articles
}

func NewArticleService(articleStorage storage.Articles) *ArticleService {
	return &ArticleService{articleStorage: articleStorage}
}

// Return course articles
func (a *ArticleService) GetCourseArticles(ctx context.Context, courseId string) ([]entity.Article, error) {
	articles, err := a.articleStorage.GetCourseArticles(ctx, courseId)
	if err != nil {
		return nil, fmt.Errorf("get article articles error: %w", err)
	}
	return articles, nil
}

// Save article
func (a *ArticleService) SaveArticle(ctx context.Context, article entity.ArticleForm) error {
	err := a.articleStorage.SaveArticle(ctx, article)
	if err != nil {
		return fmt.Errorf("save article error: %w", err)
	}
	return nil
}

func (a *ArticleService) DeleteById(ctx context.Context, articleId string) error {
	err := a.articleStorage.DeleteById(ctx, articleId)
	if err != nil {
		return fmt.Errorf("delete article error: %w", err)
	}
	return nil
}
