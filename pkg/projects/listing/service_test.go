package listing

import (
	"reflect"
	"testing"
	"time"
)

func TestListing(t *testing.T) {
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
	t.Run("Get project by id", func(t *testing.T) {
		mockrepo := new(mockRepo)
		project := Project{1, "Destroy E-Corp", now()}
		mockrepo.Projects = []Project{project}

		s := NewService(mockrepo)

		got := s.GetById(project.ID)
		want := Project{1, "Destroy E-Corp", now()}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, but wanted %v", got, want)
		}
	})
}

type mockRepo struct {
	Projects []Project
}

func (m *mockRepo) ListAll() []Project {
	return m.Projects
}

func (m *mockRepo) GetById(id int) Project {
	return m.Projects[0]
}
