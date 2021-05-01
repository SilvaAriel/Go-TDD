package listing

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestListing_ListAll(t *testing.T) {
	now := time.Now
	reports := []Report{}

	tt := map[string]struct {
		mockrepo *mockRepo
	}{
		"All projects": {&mockRepo{
			ExpectedProject: []Project{{1, "Destroy E-Corp", now(), reports}, {2, "Help Ray", now(), reports}}, ExpectedError: nil}},
		"No projects": {&mockRepo{
			ExpectedProject: []Project{}, ExpectedError: nil}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)

			got := s.ListAll()
			if !reflect.DeepEqual(got, tc.mockrepo.ExpectedProject) {
				t.Errorf("Got %v, but wanted %v", got, tc.mockrepo.ExpectedProject)
			}
		})
	}
}

func TestListing_GetById(t *testing.T) {
	now := time.Now

	reports := []Report{}
	tt := map[string]struct {
		id       int
		mockrepo *mockRepo
	}{
		"Valid ID": {
			id: 1,
			mockrepo: &mockRepo{
				ExpectedProject: []Project{{1, "Destroy E-Corp", now(), reports}},
				ExpectedError:   nil}},
		"Invalid ID": {
			id: 2,
			mockrepo: &mockRepo{
				ExpectedProject: nil,
				ExpectedError:   errors.New("Project not found")},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)

			got, err := s.GetById(tc.id)
			if err != tc.mockrepo.ExpectedError {
				t.Fatalf("Got error %v, but wanted %v", err, tc.mockrepo.ExpectedError)
			}
			if !reflect.DeepEqual(got, tc.mockrepo.ExpectedProject) {
				t.Fatalf("Got %v, but expected %v", got, tc.mockrepo.ExpectedProject)
			}
		})
	}

}

type mockRepo struct {
	ExpectedProject []Project
	ExpectedError   error
}

func (m *mockRepo) ListAll() []Project {
	return m.ExpectedProject
}

func (m *mockRepo) GetById(id int) ([]Project, error) {
	return m.ExpectedProject, m.ExpectedError
}
