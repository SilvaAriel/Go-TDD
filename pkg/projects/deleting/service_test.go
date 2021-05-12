package deleting

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestDeleting(t *testing.T) {
	now := time.Time{}
	tt := map[string]struct {
		mockrepo *mockRepo
	}{
		"Valid ID":   {&mockRepo{Input: 1, Output: &Project{ID: 1, Name: "Deleted", CreatedAt: now, Reports: []Report{}}, ExpectedError: nil}},
		"Invalid ID": {&mockRepo{Input: 5, Output: &Project{}, ExpectedError: errors.New("Error")}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)
			p, err := s.Delete(tc.mockrepo.Input)

			if !reflect.DeepEqual(p, tc.mockrepo.Output) {
				t.Errorf("Got %v, but expected %v", err, tc.mockrepo.Output)
			}
			if err != nil && tc.mockrepo.ExpectedError == nil {
				t.Errorf("Got error %q, but expected none", err)
			}
		})
	}
}

type mockRepo struct {
	Input         int
	Output        *Project
	ExpectedError error
}

func (m *mockRepo) Delete(id int) (*Project, error) {
	return m.Output, m.ExpectedError
}
