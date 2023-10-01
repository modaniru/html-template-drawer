package repos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/modaniru/html-template-drawer/internal/entity"
)

type courseStorage struct {
	db *sql.DB
}

func NewCourseStorage(db *sql.DB) *courseStorage {
	return &courseStorage{db: db}
}

// TODO pagination??
// Return all courses
func (c *courseStorage) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	rows, err := c.db.QueryContext(ctx, "select c.id, c.title, c.image, count(a.id) from Courses as c left join Articles as a on c.id::uuid = a.course_id::uuid group by c.id order by count(c.title) desc;")
	if err != nil {
		return nil, fmt.Errorf("execute query error: %w", err)
	}
	defer rows.Close()

	courses := []entity.Course{}
	for rows.Next() {
		course := entity.Course{}
		err := rows.Scan(&course.Id, &course.Title, &course.Image, &course.ArticlesCount)
		if err != nil {
			return nil, fmt.Errorf("scan query error: %w", err)
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// Save Course from entity.CourseFrom
// TODO create custom err when course already exists
func (c *courseStorage) SaveCourse(ctx context.Context, course entity.CourseForm) (string, error) {
	query := "insert into Courses (title, image) values ($1, $2) returning id;"
	stmt, err := c.db.Prepare(query)
	if err != nil {
		return "", fmt.Errorf("prepare query error: %w", err)
	}
	row := stmt.QueryRowContext(ctx, course.Name, course.Image)
	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", fmt.Errorf("scan query error: %w", err)
	}
	return id, nil
}
