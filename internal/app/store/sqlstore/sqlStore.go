package sqlstore

import (
	"aws/internal/app/store"
	"database/sql"
)

type SqlStore struct {
	db                 *sql.DB
	userRepository     *UserRepository
	academicRepository *AcademicRepository
}

func New(db *sql.DB) *SqlStore {
	return &SqlStore{
		db: db,
	}
}

func (s *SqlStore) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{store: s}
	return s.userRepository
}

func (s *SqlStore) Academic() store.AcademicRepository {
	if s.academicRepository != nil {
		return s.academicRepository
	}
	s.academicRepository = &AcademicRepository{store: s}
	return s.academicRepository
}
