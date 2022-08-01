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
)
