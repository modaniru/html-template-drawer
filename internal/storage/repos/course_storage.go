package repos

import (
	"context"
	"database/sql"

	"github.com/modaniru/html-template-drawer/internal/entity"
)

type courseStorage struct {
	db *sql.DB
}

func NewCourseStorage(db *sql.DB) *courseStorage {
	return &courseStorage{db: db}
}

// TODO pagination??
func (c *courseStorage) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	rows, err := c.db.QueryContext(ctx, "select id, title from Courses;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []entity.Course{}
	for rows.Next() {
		course := entity.Course{}
		err := rows.Scan(&course.Id, &course.Title)
		if err != nil {
			return nil, err
		}
		res = append(res, course)
	}
	return res, nil
}
