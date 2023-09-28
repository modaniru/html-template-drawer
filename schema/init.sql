create table Courses (
    id uuid DEFAULT uuid_generate_v4 (),
    title varchar
);

create table Articles (
    id uuid DEFAULT uuid_generate_v4 (),
    name varchar unique,
    course_id uuid REFERENCES Courses (id)
);

insert into Courses (title) values ("Golang course");
insert into Articles (name, course_id) values ("go_course_cycle", (select id from Courses where title = "go_course_cycle";));