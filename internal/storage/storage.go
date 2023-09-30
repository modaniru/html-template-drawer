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
		image varchar,
		title varchar
	);

	create table if not exists Articles (
		id uuid DEFAULT uuid_generate_v4() unique primary key,
		template_name varchar,
		title varchar,
		course_id uuid REFERENCES Courses (id)
	);

	insert into Courses (title, image) values ('Golang course', 'https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png');
	insert into Courses (title, image) values ('Python course', 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1869px-Python-logo-notext.svg.png');
	insert into Courses (title, image) values ('Java course', 'https://static.vecteezy.com/system/resources/previews/022/101/050/original/java-logo-transparent-free-png.png');
	insert into Courses (title, image) values ('Kotlin course', 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/06/Kotlin_Icon.svg/2048px-Kotlin_Icon.svg.png');


	insert into Articles (template_name, course_id, title) values ('go_course_cycle', (select id from Courses where title = 'Golang course'), 'gos cycles');
	insert into Articles (template_name, course_id, title) values ('go_course_cycle', (select id from Courses where title = 'Python course'), 'gos cycles');
	insert into Articles (template_name, course_id, title) values ('go_course_cycle', (select id from Courses where title = 'Java course'), 'gos cycles');
	insert into Articles (template_name, course_id, title) values ('go_course_cycle', (select id from Courses where title = 'Kotlin course'), 'gos cycles');
	`)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Storage{Articles: repos.NewArticleStorage(db), Courses: repos.NewCourseStorage(db)}
}
