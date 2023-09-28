create table author (
    name varchar,
    image_link varchar
);
insert into author (name, image_link) value ("modaniru", "link");

create table Courses (
    id uuid DEFAULT uuid_generate_v4 (),
    title varchar
);

create table Articles (
    id uuid DEFAULT uuid_generate_v4 (),
    name varchar unique,
    course_id uuid REFERENCES Courses (id)
);