package adding

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestAdding(t *testing.T) {
	now := time.Time{}
	tt := map[string]struct {
		mockrepo *mockRepo
	}{
		"Valid post": {&mockRepo{Input: &Project{Name: "Learn Regex"}, Output: &Project{ID: 1, Name: "Learn Regex", CreatedAt: now, Reports: []Report{}}, ExpectedError: nil}},
		"Empty post": {&mockRepo{Input: &Project{}, ExpectedError: &FieldError{Field: "name"}}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)
			got, err := s.Add(tc.mockrepo.Input.Name)
			if !reflect.DeepEqual(got, tc.mockrepo.Output) {
				t.Fatalf("Got %v, but wanted %v", got, tc.mockrepo.Output)
			}
			if err != nil {
				if !errors.Is(err, tc.mockrepo.ExpectedError) {
					t.Fatalf("Got error %q, but wanted error %q", err, tc.mockrepo.ExpectedError)
				}
			}

		})
	}
}

type mockRepo struct {
	Input         *Project
	Output        *Project
	ExpectedError error
}

func (m *mockRepo) Add(name string) (*Project, error) {
	return m.Output, m.ExpectedError
}
