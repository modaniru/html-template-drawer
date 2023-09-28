package storage

import (
	"context"
	"database/sql"
	"log"

	"github.com/modaniru/html-template-drawer/internal/entity"
	"github.com/modaniru/html-template-drawer/internal/storage/repos"
)

type Articles interface {
	GetCourseArticles(ctx context.Context, courseId string) ([]entity.Article, error)
}

type Courses interface {
	GetAllCourses(ctx context.Context) ([]entity.Course, error)
}

type Storage struct {
	Articles Articles
	Courses  Courses
}

func NewStorage(db *sql.DB) *Storage {
	_, err := db.Exec(`
	DROP TABLE IF EXISTS Articles;
	DROP TABLE IF EXISTS Courses;
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	create table if not exists Courses (
		id uuid DEFAULT uuid_generate_v4() unique primary key,
		title varchar
	);

	create table if not exists Articles (
		id uuid DEFAULT uuid_generate_v4() unique primary key,
		name varchar unique,
		course_id uuid REFERENCES Courses (id)
	);

	insert into Courses (title) values ('Golang course');
	insert into Courses (title) values ('Python course');
	insert into Courses (title) values ('Java course');
	insert into Courses (title) values ('Kotlin course');


	insert into Articles (name, course_id) values ('go_course_cycle', (select id from Courses where title = 'Golang course'));
	`)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Storage{Articles: repos.NewArticleStorage(db), Courses: repos.NewCourseStorage(db)}
}
