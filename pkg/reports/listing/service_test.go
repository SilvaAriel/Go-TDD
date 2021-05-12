package listing

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestListing(t *testing.T) {
	now := time.Time{}
	tt := map[string]struct {
		mockrepo *mockRepo
	}{
		"List by ID":         {&mockRepo{ProjectId: 1, ReportId: 1, Output: &Report{ID: 1, name: "Teste", CreatedAt: now, Success: true}, ExpectedError: nil}},
		"Invalid Project ID": {&mockRepo{ProjectId: 5, ReportId: 1, Output: &Report{}, ExpectedError: errors.New("")}},
		"Invalid Report ID":  {&mockRepo{ProjectId: 1, ReportId: 5, Output: &Report{}, ExpectedError: errors.New("")}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)
			report, err := s.GetById(tc.mockrepo.ProjectId, tc.mockrepo.ReportId)

			if !reflect.DeepEqual(report, tc.mockrepo.Output) {
				t.Errorf("Got %v, but expected %v", err, tc.mockrepo.Output)
			}
			if err != nil && tc.mockrepo.ExpectedError == nil {
				t.Errorf("Got error %q, but expected none", err)
			}
		})
	}
}

type mockRepo struct {
	ProjectId     int
	ReportId      int
	Output        *Report
	ExpectedError error
}

func (m *mockRepo) GetById(projectId int, reportId int) (*Report, error) {
	return m.Output, m.ExpectedError
}
