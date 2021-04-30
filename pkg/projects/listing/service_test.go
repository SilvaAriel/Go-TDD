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
		mockrepo.Projects = []Project{{"Destroy E-Corp", now()}}

		s := NewService(mockrepo)

		got := s.ListAllProjects()
		want := []Project{
			{"Destroy E-Corp", now()},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %s, but wanted %s", got, want)
		}
	})
}

type mockRepo struct {
	Projects []Project
}

func (m *mockRepo) ListAllProjects() []Project {
	return m.Projects
}
