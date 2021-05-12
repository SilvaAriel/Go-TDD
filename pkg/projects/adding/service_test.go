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
		"Valid adding":     {&mockRepo{Input: &Project{Name: "Learn Regex"}, Output: &Project{ID: 1, Name: "Learn Regex", CreatedAt: now, Reports: []Report{}}, ExpectedError: nil}},
		"Empty name field": {&mockRepo{Input: &Project{}, ExpectedError: errors.New("Shit happens")}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)
			got, err := s.Add(tc.mockrepo.Input.Name)
			if !reflect.DeepEqual(got, tc.mockrepo.Output) {
				t.Fatalf("Got %v, but wanted %v", got, tc.mockrepo.Output)
			}
			if err != nil && tc.mockrepo.ExpectedError == nil {
				t.Fatalf("Got error %q, but expected none", err)
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
