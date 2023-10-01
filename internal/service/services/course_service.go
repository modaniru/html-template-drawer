package services

import (
	"context"
	"fmt"

	"github.com/modaniru/html-template-drawer/internal/entity"
	"github.com/modaniru/html-template-drawer/internal/storage"
)

type CourseService struct{
	courseStorage storage.Courses
}

func NewCourseService(courseStorage storage.Courses) *CourseService{
	return &CourseService{courseStorage: courseStorage}
}

//Return all courses
func (c *CourseService) GetAllCourses(ctx context.Context) ([]entity.Course, error){
	courses, err := c.courseStorage.GetAllCourses(ctx)
	if err != nil{
		return nil, fmt.Errorf("get all courses error: %w", err)
	}
	return courses, nil
}

//Save course
func (c *CourseService) SaveCourse(ctx context.Context, course entity.CourseForm) (string, error){
	courseId, err := c.courseStorage.SaveCourse(ctx, course)
	if err != nil{
		return "", fmt.Errorf("save course error: %w", err)
	}
	return courseId, nil
}