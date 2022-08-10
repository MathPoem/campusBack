package store

import "aws/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	Find(int) (*model.User, error)
}

type AcademicRepository interface {
	GetUniversity() (*[]model.University, error)
	GetSchoolByUniversity(idUniversity []string) (*[]model.School, error)
	GetSchoolList() (*[]model.School, error)
	GetProgramBySchool(idSchool []string) (*[]model.Program, error)
	GetProgramList() (*[]model.Program, error)
	GetCourseList() (*[]model.Course, error)
	GetCourseByProgram(idProgram []string) (*[]model.Course, error)
	GetDepartmentList() (*[]model.Department, error)
	GetDepartmentBySchool(idSchool []string) (*[]model.Department, error)
	GetPersonList() (*[]model.Person, error)
	GetPersonByDepartment(idDepartment []string) (*[]model.Person, error)
	GetLectureList() (*[]model.Lecture, error)
	GetLectureByCourse(idCourse []string) (*[]model.Lecture, error)
	GetSeminarList() (*[]model.Seminar, error)
	GetSeminarByCourse(idCourse []string) (*[]model.Seminar, error)
	InDefaultBase() error
}
