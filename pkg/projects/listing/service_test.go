package listing

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestListing_ListAll(t *testing.T) {
	now := time.Now

	t.Run("List all projects", func(t *testing.T) {
		mockrepo := new(mockRepo)
		mockrepo.Projects = []Project{{1, "Destroy E-Corp", now()}, {2, "Help Ray", now()}}

		s := NewService(mockrepo)

		got := s.ListAll()
		want := []Project{
			{1, "Destroy E-Corp", now()},
			{2, "Help Ray", now()},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, but wanted %v", got, want)
		}
	})
}

func TestListing_GetById(t *testing.T) {
	now := time.Now

	tt := []struct {
		id            int
		mockProject   []Project
		want          Project
		expectedError error
	}{
		{
			id:          1,
			mockProject: []Project{{1, "Destroy E-Corp", now()}},
			want:        Project{1, "Destroy E-Corp", now()},
		},
		{
			id:            2,
			mockProject:   []Project{{1, "Destroy E-Corp", now()}},
			want:          Project{},
			expectedError: errors.New("Id not found"),
		},
	}

	for _, tc := range tt {
		t.Run("Get project by id", func(t *testing.T) {
			mockrepo := new(mockRepo)
			project := Project{1, "Destroy E-Corp", now()}
			mockrepo.Projects = tc.mockProject

			s := NewService(mockrepo)

			got, err := s.GetById(project.ID)
			if err != tc.expectedError {
				t.Errorf("Got %v, but wanted %v", got, tc.want)
			}
		})
	}

}

type mockRepo struct {
	ExpectedProject []Project
}

func (m *mockRepo) ListAll() []Project {
	return m.ExpectedProject
}

func (m *mockRepo) GetById(id int) (Project, error) {
	return m.expectedProject, m.expectedError
}
