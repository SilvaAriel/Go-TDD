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
		"Add one Report":   {&mockRepo{ProjectId: 6, Name: "Teste", Success: true, Output: &Report{1, 6, "Teste", now, true}, ExpectedError: nil}},
		"Wrong Project Id": {&mockRepo{ProjectId: 10, Name: "Teste", Success: true, Output: &Report{}, ExpectedError: errors.New("Error")}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)
			report, err := s.Add(tc.mockrepo.ProjectId, tc.mockrepo.Name, tc.mockrepo.Success)

			if tc.mockrepo.ExpectedError == nil {
				if report.ProjectID != tc.mockrepo.ProjectId {
					t.Errorf("Got ProjectId %d, but expected %d", report.ProjectID, tc.mockrepo.ProjectId)
				}

				if report.name != tc.mockrepo.Name {
					t.Errorf("Got name %q, but expected %q", report.name, tc.mockrepo.Name)
				}

				if report.Success != tc.mockrepo.Success {
					t.Errorf("Got Success Status %v, but expected %v", report.Success, tc.mockrepo.Success)
				}
			} else {
				if err != tc.mockrepo.ExpectedError {
					t.Errorf("Got error %q, but expected %q", err, tc.mockrepo.ExpectedError)
				}
			}

			if !reflect.DeepEqual(report, tc.mockrepo.Output) {
				t.Errorf("Got %v, but expected %v", report, tc.mockrepo.Output)
			}

			if err != tc.mockrepo.ExpectedError {
				t.Errorf("Got error %q, but expected %q", err, tc.mockrepo.ExpectedError)
			}
		})
	}
}

type mockRepo struct {
	ProjectId     int
	Name          string
	Success       bool
	Output        *Report
	ExpectedError error
}

func (m *mockRepo) Add(projectId int, name string, success bool) (*Report, error) {
	return m.Output, m.ExpectedError
}
