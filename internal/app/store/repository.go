package store

import "aws/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	Find(int) (*model.User, error)
}

type AcademicRepository interface {
	GetUniversity() (*[]model.University, error)
	GetSchool(idUniversity string) (*[]model.School, error)
	GetProgram(idSchool string) (*[]model.Program, error)
	InDefaultBase() error
}
