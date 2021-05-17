package updating

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestUpdating(t *testing.T) {
	now := time.Time{}
	tt := map[string]struct {
		id       int
		name     string
		mockrepo *mockRepo
	}{
		"With correct ID": {
			id: 1, name: "Work at E-Corp",
			mockrepo: &mockRepo{output: Project{ID: 1, Name: "Work at E-Corp", Status: Status{Success: 10, Failure: 5, Total: 15}, CreatedAt: now}, expectedError: nil}},
		"With incorrect ID": {
			id: 5, name: "Work at E-Corp",
			mockrepo: &mockRepo{output: Project{}, expectedError: errors.New("")}},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)
			got, err := s.Update(tc.id, tc.name)

			if !reflect.DeepEqual(got, tc.mockrepo.output) {
				t.Errorf("Got %v, but expected %v", got, tc.mockrepo.output)
			}
			if err != nil && tc.mockrepo.expectedError == nil {
				t.Errorf("Got error %q, but expected none", err)
			}
		})
	}
}

type mockRepo struct {
	output        Project
	expectedError error
}

func (m *mockRepo) Update(id int, name string) (Project, error) {
	return m.output, m.expectedError
}
