package sqlstore

import (
	"aws/internal/app/model"
)

type AcademicRepository struct {
	store *SqlStore
}

func (r *AcademicRepository) GetUniversity() (*[]model.University, error) {
	var ul []model.University
	rows, err := r.store.db.Query(
		"SELECT id, country, city, name, url FROM university")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u model.University
		if err := rows.Scan(&u.ID, &u.Country, &u.City, &u.Name, &u.URL); err != nil {
			return nil, err
		}
		ul = append(ul, u)
	}
	return &ul, nil
}

func (r *AcademicRepository) GetSchool(idUniversity string) (*[]model.School, error) {
	var sl []model.School
	rows, err := r.store.db.Query(
		"SELECT id, name, university_id, url  FROM school WHERE university_id = $1", &idUniversity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s model.School
		if err := rows.Scan(&s.ID, &s.Name, &s.UniversityID, &s.URL); err != nil {
			return nil, err
		}
		sl = append(sl, s)
	}
	return &sl, nil
}

func (r *AcademicRepository) GetProgram(idSchool string) (*[]model.Program, error) {
	var pl []model.Program
	rows, err := r.store.db.Query(
		"SELECT id, name, school_id, semester, year_start FROM program WHERE school_id = $1", &idSchool)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Program
		if err := rows.Scan(&p.ID, &p.Name, &p.SchoolID, &p.Semester, &p.YearStart); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) InDefaultBase() error {
	for _, univ := range model.UnivList {
		r.store.db.QueryRow(
			"INSERT INTO university (id, country, city, name, url) VALUES ($1, $2, $3, $4, $5)",
			&univ.ID,
			&univ.Country,
			&univ.City,
			&univ.Name,
			&univ.URL)
	}
	for _, school := range model.SchoolList {
		r.store.db.QueryRow(
			"INSERT INTO school (id, name, university_id, url) VALUES ($1, $2, $3, $4)",
			&school.ID,
			&school.Name,
			&school.UniversityID,
			&school.URL)
	}
	for _, program := range model.ProgramList {
		r.store.db.QueryRow(
			"INSERT INTO program (id, school_id, name, year_start, semester) VALUES ($1, $2, $3, $4, $5)",
			&program.ID,
			&program.SchoolID,
			&program.Name,
			&program.YearStart,
			&program.Semester)
	}
	return nil
}
