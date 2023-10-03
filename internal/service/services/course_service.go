package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/modaniru/html-template-drawer/internal/entity"
	"github.com/modaniru/html-template-drawer/internal/storage"
)

type CourseService struct {
	courseStorage storage.Courses
}

func NewCourseService(courseStorage storage.Courses) *CourseService {
	return &CourseService{courseStorage: courseStorage}
}

// Return all courses
func (c *CourseService) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	courses, err := c.courseStorage.GetAllCourses(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all courses error: %w", err)
	}
	return courses, nil
}

// Save course
func (c *CourseService) SaveCourse(ctx context.Context, course entity.CourseForm) (string, error) {
	courseId, err := c.courseStorage.SaveCourse(ctx, &entity.SaveCourse{Title: course.Name, TitleId: toTitleId(course.Name), Image: course.Image})
	if err != nil {
		return "", fmt.Errorf("save course error: %w", err)
	}
	return courseId, nil
}

func (c *CourseService) DeleteCourse(ctx context.Context, courseId string) error {
	err := c.courseStorage.DeleteCourse(ctx, courseId)
	if err != nil {
		return fmt.Errorf("delete course error: %w", err)
	}
	return nil
}

func toTitleId(title string) string {
	result := strings.ToLower(title)
	result = strings.ReplaceAll(result, " ", "_")
	return result
}
