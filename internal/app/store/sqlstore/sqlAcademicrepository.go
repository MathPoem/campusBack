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

func (r *AcademicRepository) GetSchoolByUniversity(idUniversity []string) (*[]model.School, error) {
	var sl []model.School
	rows, err := r.store.db.Query(
		"SELECT id, name, university_id, url  FROM school WHERE university_id = $1", &idUniversity[0])
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

func (r *AcademicRepository) GetSchoolList() (*[]model.School, error) {
	var sl []model.School
	rows, err := r.store.db.Query(
		"SELECT id, name, university_id, url  FROM school")
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

func (r *AcademicRepository) GetProgramBySchool(idSchool []string) (*[]model.Program, error) {
	var pl []model.Program
	rows, err := r.store.db.Query(
		"SELECT id, name, school_id, semester, year_start FROM program WHERE school_id = $1", &idSchool[0])
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

func (r *AcademicRepository) GetProgramList() (*[]model.Program, error) {
	var pl []model.Program
	rows, err := r.store.db.Query(
		"SELECT id, name, school_id, semester, year_start FROM program")
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

func (r *AcademicRepository) GetCourseByProgram(idProgram []string) (*[]model.Course, error) {
	var pl []model.Course
	rows, err := r.store.db.Query(
		"SELECT id, program_id, credits, hours_lecture, hours_seminar, estimation_in_diploma, exam, description, url, name FROM course WHERE program_id = $1", &idProgram[0])
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Course
		if err := rows.Scan(&p.ID, &p.ProgramID, &p.Credits, &p.HoursLecture, &p.HoursSeminar, &p.EstimationInDiploma, &p.Exam, &p.Description, &p.URL, &p.Name); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetCourseList() (*[]model.Course, error) {
	var pl []model.Course
	rows, err := r.store.db.Query(
		"SELECT id, program_id, credits, hours_lecture, hours_seminar, estimation_in_diploma, exam, description, url, name FROM course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Course
		if err := rows.Scan(&p.ID, &p.ProgramID, &p.Credits, &p.HoursLecture, &p.HoursSeminar, &p.EstimationInDiploma, &p.Exam, &p.Description, &p.URL, &p.Name); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetDepartmentBySchool(idSchool []string) (*[]model.Department, error) {
	var pl []model.Department
	rows, err := r.store.db.Query(
		"SELECT id, name, school_id, url FROM department WHERE school_id = $1", &idSchool[0])
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Department
		if err := rows.Scan(&p.ID, &p.Name, &p.SchoolID, &p.URL); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetDepartmentList() (*[]model.Department, error) {
	var pl []model.Department
	rows, err := r.store.db.Query(
		"SELECT id, name, school_id, url FROM department")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Department
		if err := rows.Scan(&p.ID, &p.Name, &p.SchoolID, &p.URL); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetPersonByDepartment(idDepartment []string) (*[]model.Person, error) {
	var pl []model.Person
	rows, err := r.store.db.Query(
		"SELECT id, department_id, first_name, middle_name, second_name, age, url, first_degree, second_degree, third_degree FROM person WHERE department_id = $1", &idDepartment[0])
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Person
		if err := rows.Scan(&p.ID, &p.DepartmentID, &p.FirstName, &p.MiddleName, &p.SecondName, &p.Age, &p.URL, &p.FirstDegree, &p.SecondDegree, &p.ThirdDegree); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetPersonList() (*[]model.Person, error) {
	var pl []model.Person
	rows, err := r.store.db.Query(
		"SELECT id, department_id, first_name, middle_name, second_name, age, url, first_degree, second_degree, third_degree FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Person
		if err := rows.Scan(&p.ID, &p.DepartmentID, &p.FirstName, &p.MiddleName, &p.SecondName, &p.Age, &p.URL, &p.FirstDegree, &p.SecondDegree, &p.ThirdDegree); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetLectureByCourse(idCourse []string) (*[]model.Lecture, error) {
	var pl []model.Lecture
	rows, err := r.store.db.Query(
		"SELECT id, person_id, course_id, url, year FROM lecture WHERE course_id = $1", &idCourse[0])
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Lecture
		if err := rows.Scan(&p.ID, &p.PersonID, &p.CourseID, &p.URL, &p.Year); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetLectureList() (*[]model.Lecture, error) {
	var pl []model.Lecture
	rows, err := r.store.db.Query(
		"SELECT id, person_id, course_id, url, year FROM lecture")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Lecture
		if err := rows.Scan(&p.ID, &p.PersonID, &p.CourseID, &p.URL, &p.Year); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetSeminarByCourse(idCourse []string) (*[]model.Seminar, error) {
	var pl []model.Seminar
	rows, err := r.store.db.Query(
		"SELECT id, person_id, course_id, url, year FROM lecture WHERE course_id = $1", &idCourse[0])
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Seminar
		if err := rows.Scan(&p.ID, &p.PersonID, &p.CourseID, &p.URL, &p.Year); err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return &pl, nil
}

func (r *AcademicRepository) GetSeminarList() (*[]model.Seminar, error) {
	var pl []model.Seminar
	rows, err := r.store.db.Query(
		"SELECT id, person_id, course_id, url, year FROM lecture")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Seminar
		if err := rows.Scan(&p.ID, &p.PersonID, &p.CourseID, &p.URL, &p.Year); err != nil {
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
	for _, course := range model.CourseList {
		r.store.db.QueryRow(
			"INSERT INTO course (id, url, name,credits,description,estimation_in_diploma,exam,hours_lecture,hours_seminar,program_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
			&course.ID,
			&course.URL,
			&course.Name,
			&course.Credits,
			&course.Description,
			&course.EstimationInDiploma,
			&course.Exam,
			&course.HoursLecture,
			&course.HoursSeminar,
			&course.ProgramID)
	}
	for _, department := range model.DepartmentList {
		r.store.db.QueryRow(
			"INSERT INTO department (id, name, school_id, url) VALUES ($1, $2, $3, $4)",
			&department.ID,
			&department.Name,
			&department.SchoolID,
			&department.URL)
	}
	for _, person := range model.PersonList {
		r.store.db.QueryRow(
			"INSERT INTO person (id, department_id, first_name, middle_name, second_name, age, url, first_degree, second_degree, third_degree)  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
			&person.ID,
			&person.DepartmentID,
			&person.FirstName,
			&person.MiddleName,
			&person.SecondName,
			&person.Age,
			&person.URL,
			&person.FirstDegree,
			&person.SecondDegree,
			&person.ThirdDegree)
	}
	for _, lecture := range model.LectureList {

		row := r.store.db.QueryRow(
			"INSERT INTO lecture (id, year, person_id, course_id, url) VALUES ($1, $2, $3, $4, $5)",
			&lecture.ID,
			&lecture.Year,
			&lecture.PersonID,
			&lecture.CourseID,
			&lecture.URL)
		if err := row.Err(); err != nil {
			return err
		}
	}
	for _, seminar := range model.SeminarList {
		r.store.db.QueryRow(
			"INSERT INTO seminar (id, year, person_id, course_id, url) VALUES ($1, $2, $3, $4, $5)",
			&seminar.ID,
			&seminar.Year,
			&seminar.PersonID,
			&seminar.CourseID,
			&seminar.URL)
	}
	return nil
}
