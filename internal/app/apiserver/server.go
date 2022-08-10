package apiserver

import (
	_ "aws/docs"
	"aws/internal/app/model"
	"aws/internal/app/store"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(handlers.CORS(
		handlers.AllowedMethods([]string{http.MethodGet}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
	))

	s.router.HandleFunc("/users", s.handleAcademic()).Methods("GET")
	s.router.HandleFunc("/university", s.handleUniversity()).Methods("GET")
	s.router.HandleFunc("/school", s.handleSchool()).Methods("GET")
	s.router.HandleFunc("/program", s.handleProgram()).Methods("GET")
	s.router.HandleFunc("/course", s.handleCourse()).Methods("GET")
	s.router.HandleFunc("/department", s.handleDepartment()).Methods("GET")
	s.router.HandleFunc("/person", s.handlePerson()).Methods("GET")
	s.router.HandleFunc("/lecture", s.handleLecture()).Methods("GET")
	s.router.HandleFunc("/seminar", s.handleSeminar()).Methods("GET")

	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

}

func handlerAcademicTemplate[M model.AcademicModelTypeOrdered](
	academicModel []M,
	errorHandler func(w http.ResponseWriter, r *http.Request, code int, err error),
	respondHandler func(w http.ResponseWriter, r *http.Request, code int, data interface{}),
	filterNameField string,
	GetWithFilter func([]string) (*[]M, error),
	Get func() (*[]M, error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		if err := validateQuery(vars, filterNameField); err != nil {
			errorHandler(w, r, http.StatusBadRequest, err)
		}
		school, err := &academicModel, errors.New("")
		if len(vars) == 1 {
			school, err = GetWithFilter(vars[filterNameField])
		} else {
			school, err = Get()
		}
		if err != nil {
			errorHandler(w, r, http.StatusBadRequest, err)
			return
		}
		respondHandler(w, r, http.StatusOK, school)
	}
}

func (s *server) handleAcademic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		fmt.Println(v)
	}
}

// @Summary      Get university list
// @Tags		 Public
// @Description  get university list
// @ID           university-list
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.University
// @Failure      400  {object}  string
// @Router       /university [get]
func (s *server) handleUniversity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.store.Academic().GetUniversity()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		s.respond(w, r, http.StatusOK, u)
	}
}

// @Summary      Get school list
// @Param        universityID   query      int  false  "University ID"
// @Tags		 Public
// @Description  get school list
// @ID           school-list
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.School
// @Failure      400  {object}  string
// @Router       /school [get]
func (s *server) handleSchool() http.HandlerFunc {
	return handlerAcademicTemplate(
		[]model.School{},
		s.error,
		s.respond,
		"universityID",
		s.store.Academic().GetSchoolByUniversity,
		s.store.Academic().GetSchoolList,
	)
}

// @Summary      Get program list
// @Param        schoolID   query      int  false  "School ID"
// @Tags		 Public
// @Description  get program list
// @ID           program-list
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Program
// @Failure      400  {object}  string
// @Router       /program [get]
func (s *server) handleProgram() http.HandlerFunc {
	return handlerAcademicTemplate(
		[]model.Program{},
		s.error,
		s.respond,
		"schoolID",
		s.store.Academic().GetProgramBySchool,
		s.store.Academic().GetProgramList,
	)
}

// @Summary      Get course list
// @Param        courseID   query      int  false  "Course ID"
// @Tags		 Public
// @Description  get course list
// @ID           course-list
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Course
// @Failure      400  {object}  string
// @Router       /course [get]
func (s *server) handleCourse() http.HandlerFunc {
	return handlerAcademicTemplate(
		[]model.Course{},
		s.error,
		s.respond,
		"programID",
		s.store.Academic().GetCourseByProgram,
		s.store.Academic().GetCourseList,
	)
}

// @Summary      Get department list
// @Param        schoolID   query      int  false  "school ID"
// @Tags		 Public
// @Description  get department list
// @ID           department-list
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Department
// @Failure      400  {object}  string
// @Router       /department [get]
func (s *server) handleDepartment() http.HandlerFunc {
	return handlerAcademicTemplate(
		[]model.Department{},
		s.error,
		s.respond,
		"schoolID",
		s.store.Academic().GetDepartmentBySchool,
		s.store.Academic().GetDepartmentList,
	)
}

// @Summary      Get pearson list
// @Param        departmentID   query      int  false  "department ID"
// @Tags		 Public
// @Description  get person list
// @ID           person-list
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Person
// @Failure      400  {object}  string
// @Router       /person [get]
func (s *server) handlePerson() http.HandlerFunc {
	return handlerAcademicTemplate(
		[]model.Person{},
		s.error,
		s.respond,
		"departmentID",
		s.store.Academic().GetPersonByDepartment,
		s.store.Academic().GetPersonList,
	)
}

// @Summary      Get lecture list
// @Param        courseID   query      int  false  "course ID"
// @Tags		 Public
// @Description  get lecture list
// @ID           lecture-list
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Lecture
// @Failure      400  {object}  string
// @Router       /lecture [get]
func (s *server) handleLecture() http.HandlerFunc {
	return handlerAcademicTemplate(
		[]model.Lecture{},
		s.error,
		s.respond,
		"courseID",
		s.store.Academic().GetLectureByCourse,
		s.store.Academic().GetLectureList,
	)
}

// @Summary      Get seminar list
// @Param        courseID   query      int  false  "course ID"
// @Tags		 Public
// @Description  get seminar list
// @ID           seminar-list
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Seminar
// @Failure      400  {object}  string
// @Router       /seminar [get]
func (s *server) handleSeminar() http.HandlerFunc {
	return handlerAcademicTemplate(
		[]model.Seminar{},
		s.error,
		s.respond,
		"courseID",
		s.store.Academic().GetSeminarByCourse,
		s.store.Academic().GetSeminarList,
	)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
