package sqlstore_test

import (
	"aws/internal/app/model"
	"aws/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcademicRepository_GetUniversity(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("university")

	s := sqlstore.New(db)

	err := s.Academic().InDefaultBase()
	if err != nil {
		t.Fatal("error create test model")
	}
	d, err := s.Academic().GetUniversity()
	assert.NoError(t, err)
	assert.Equal(t, &model.UnivList, d)
}

func TestAcademicRepository_GetSchool(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("university")

	s := sqlstore.New(db)

	err := s.Academic().InDefaultBase()
	if err != nil {
		t.Fatal("error create test model")
	}
	t.Run("fullList", func(t *testing.T) {
		d, err := s.Academic().GetSchoolList()
		assert.NoError(t, err)
		assert.Equal(t, &model.SchoolList, d)
	})

	testCases := []struct {
		name          string
		payload       string
		expectedModel []model.School
		err           error
	}{
		{
			name:          "exist",
			payload:       "1",
			expectedModel: model.Filter(model.SchoolList, model.FilterSchool, 1),
			err:           nil,
		},
		{
			name:          "exist",
			payload:       "2",
			expectedModel: model.Filter(model.SchoolList, model.FilterSchool, 2),
			err:           nil,
		},
		{
			name:          "notexist",
			payload:       "100",
			expectedModel: nil,
			err:           nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := s.Academic().GetSchoolByUniversity([]string{tc.payload})
			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, &tc.expectedModel, u)
		})
	}
}
