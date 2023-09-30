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
	rows, err := c.db.QueryContext(ctx, "select c.id, c.title, c.image, count(c.title) from Courses as c inner join Articles as a on c.id::uuid = a.course_id::uuid group by c.id;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []entity.Course{}
	for rows.Next() {
		course := entity.Course{}
		err := rows.Scan(&course.Id, &course.Title, &course.Image, &course.ArticlesCount)
		if err != nil {
			return nil, err
		}
		res = append(res, course)
	}
	return res, nil
}
