CREATE TABLE "university" (
                              "id" bigserial not null PRIMARY KEY ,
                              "country" varchar not null ,
                              "city" varchar not null ,
                              "name" varchar not null unique ,
                              "url" varchar
);

CREATE TABLE "school" (
                          "id" bigserial not null PRIMARY KEY ,
                          "name" varchar not null ,
                          "university_id" int not null ,
                          "url" varchar
);

CREATE TABLE "department" (
                               "id" bigserial not null PRIMARY KEY,
                               "name" varchar not null ,
                               "school_id" int not null ,
                               "url" varchar
);

CREATE TABLE "person" (
                           "id" bigserial not null PRIMARY KEY ,
                           "department_id" int not null ,
                           "first_name" varchar not null ,
                           "middle_name" varchar not null ,
                           "second_name" varchar not null ,
                           "age" int not null ,
                           "url" varchar ,
                           "first_degree" varchar,
                           "second_degree" varchar,
                           "third_degree" varchar
);

CREATE TABLE "program" (
                           "id" bigserial not null primary key ,
                           "school_id" int not null ,
                           "name" varchar not null ,
                           "year_start" int not null ,
                           "semester" int not null
);

CREATE TABLE "course" (
                          "id" bigserial not null primary key ,
                          "name" varchar not null ,
                          "program_id" int not null ,
                          "credits" int not null ,
                          "hours_lecture" int not null ,
                          "hours_seminar" int not null ,
                          "estimation_in_diploma" boolean not null ,
                          "exam" boolean not null ,
                          "description" text not null ,
                          "url" varchar not null
);

CREATE TABLE "lecture" (
                           "id" bigserial not null primary key ,
                           "year" int not null ,
                           "person_id" int not null ,
                           "course_id" int not null ,
                           "url" varchar
);

CREATE TABLE "seminar" (
                           "id" bigserial not null primary key ,
                           "year" int not null ,
                           "person_id" int not null ,
                           "course_id" int not null ,
                           "url" varchar
);

ALTER TABLE "lecture" ADD FOREIGN KEY (person_id) REFERENCES person ("id");

ALTER TABLE "lecture" ADD FOREIGN KEY ("course_id") REFERENCES "course" ("id");

ALTER TABLE "seminar" ADD FOREIGN KEY (person_id) REFERENCES person ("id");

ALTER TABLE "seminar" ADD FOREIGN KEY ("course_id") REFERENCES "course" ("id");

ALTER TABLE "course" ADD FOREIGN KEY ("program_id") REFERENCES "program" ("id");

ALTER TABLE "program" ADD FOREIGN KEY ("school_id") REFERENCES "school" ("id");

ALTER TABLE person ADD FOREIGN KEY ("department_id") REFERENCES "department" ("id");

ALTER TABLE "school" ADD FOREIGN KEY ("university_id") REFERENCES "university" ("id");

ALTER TABLE "department" ADD FOREIGN KEY ("school_id") REFERENCES "school" ("id");