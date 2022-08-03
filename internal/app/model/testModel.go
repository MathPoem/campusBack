package model

var (
	UnivList = []University{
		{
			Name:    "МГУ",
			ID:      1,
			City:    "Москва",
			Country: "Россия",
			URL:     "msu.ru",
		},
		{
			Name:    "ВШЭ",
			ID:      2,
			City:    "Москва",
			Country: "Россия",
			URL:     "msu.ru",
		},
		{
			Name:    "МФТИ",
			ID:      3,
			City:    "Москва",
			Country: "Россия",
			URL:     "msu.ru",
		},
		{
			Name:    "МГТУ",
			ID:      4,
			City:    "Москва",
			Country: "Россия",
			URL:     "msu.ru",
		},
		{
			Name:    "МИСИС",
			ID:      5,
			City:    "Москва",
			Country: "Россия",
			URL:     "msu.ru",
		},
		{
			Name:    "МИФИ",
			ID:      6,
			City:    "Москва",
			Country: "Россия",
			URL:     "msu.ru",
		},
		{
			Name:    "МГИМО",
			ID:      7,
			City:    "Москва",
			Country: "Россия",
			URL:     "msu.ru",
		},
	}

	SchoolList = []School{
		{
			ID:           1,
			Name:         "МехМат",
			UniversityID: 1,
			URL:          "math.msu.ru",
		},
		{
			ID:           2,
			Name:         "ВМК",
			UniversityID: 1,
			URL:          "cs.msy.ru",
		},
		{
			ID:           3,
			Name:         "Физфак",
			UniversityID: 1,
			URL:          "phys.msu.ru",
		},
		{
			ID:           4,
			Name:         "ФКН",
			UniversityID: 2,
			URL:          "cs.hse",
		},
		{
			ID:           5,
			Name:         "Матфак",
			UniversityID: 2,
			URL:          "math.hse",
		},
		{
			ID:           6,
			Name:         "Физфак",
			UniversityID: 2,
			URL:          "phys.hse",
		},
		{
			ID:           7,
			Name:         "ФПМИ",
			UniversityID: 3,
			URL:          "fpmi.mipt",
		},
		{
			ID:           8,
			Name:         "ЛФИ",
			UniversityID: 3,
			URL:          "lfi.mipt",
		},
		{
			ID:           9,
			Name:         "ФРКТ",
			UniversityID: 3,
			URL:          "frkt.mipt",
		},
	}
	ProgramList = []Program{
		{
			ID:        1,
			SchoolID:  1,
			Name:      "Математика",
			YearStart: 0,
			Semester:  0,
		},
		{
			ID:        2,
			SchoolID:  1,
			Name:      "Механика",
			YearStart: 0,
			Semester:  0,
		},
	}
	PersonList = []Pearson{
		{
			ID:           1,
			DepartmentID: 1,
			FirstName:    "Александр",
			MiddleName:   "Алексеевич",
			SecondName:   "Флеров",
			Age:          33,
			URL:          "https://istina.msu.ru/profile/AAFlerov/",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		},
		{
			ID:           2,
			DepartmentID: 2,
			FirstName:    "Наталия",
			MiddleName:   "Евгеньевна",
			SecondName:   "Шавгулидзе",
			Age:          40,
			URL:          "https://istina.msu.ru/profile/Natalia_Shavgulidze/",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		},
		{
			ID:           3,
			DepartmentID: 2,
			FirstName:    "Юрий",
			MiddleName:   "Викторович",
			SecondName:   "Садовничий",
			Age:          56,
			URL:          "https://istina.msu.ru/profile/uvs/",
			FirstDegree:  "",
			SecondDegree: "",
			ThirdDegree:  "",
		},
	}
	DepartmentList = []Department{
		{
			ID:       1,
			Name:     "Кафедра МатематическогоАнализа",
			SchoolID: 1,
			URL:      "",
		},
		{
			ID:       2,
			Name:     "Кафедра Общей Геометрии и Топологии",
			SchoolID: 1,
			URL:      "",
		},
	}
	CourseList = []Course{
		{
			ID:                  1,
			Name:                "Маттематический Анализ 1",
			ProgramID:           1,
			Credits:             8,
			HoursLecture:        4,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                true,
			Description:         "первая часть общего курса математического анализа",
			URL:                 "",
		},
		{
			ID:                  2,
			Name:                "Математичексий Анализ 2",
			ProgramID:           1,
			Credits:             8,
			HoursLecture:        4,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                true,
			Description:         "вторая часть общего курса математического анализа",
			URL:                 "",
		},
		{
			ID:                  3,
			Name:                "Аналитическа Геометрия",
			ProgramID:           1,
			Credits:             8,
			HoursLecture:        4,
			HoursSeminar:        4,
			EstimationInDiploma: false,
			Exam:                true,
			Description:         "Аналичическая геометрия и введение в линейную алгебру",
			URL:                 "",
		},
	}
	LectureList = []Lecture{
		{
			ID:       1,
			Year:     2021,
			PersonID: 3,
			CourseID: 3,
			URL:      "",
		},
	}
	SeminarList = []Seminar{
		{
			ID:       1,
			Year:     2021,
			PersonID: 2,
			CourseID: 3,
			URL:      "",
		},
	}
)
