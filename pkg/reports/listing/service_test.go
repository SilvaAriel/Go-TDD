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
		"List by ID":              {&mockRepo{1, 6, &Report{1, 6, "Teste", now, true}, nil}},
		"Return wrong Project ID": {&mockRepo{1, 6, &Report{1, 6, "Teste", now, true}, nil}},
		"Return wrong Report ID":  {&mockRepo{2, 7, &Report{2, 7, "Teste", now, true}, nil}},
		"Invalid Project ID":      {&mockRepo{5, 1, nil, errors.New("")}},
		"Invalid Report ID":       {&mockRepo{1, 5, nil, errors.New("")}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := NewService(tc.mockrepo)
			report, err := s.GetById(tc.mockrepo.ProjectId, tc.mockrepo.ReportId)

			if tc.mockrepo.Output != nil {
				if tc.mockrepo.ProjectId != tc.mockrepo.Output.ProjectID {
					t.Errorf("Got Project ID %d, but expected %d", tc.mockrepo.Output.ProjectID, tc.mockrepo.ProjectId)
				}

				if tc.mockrepo.ReportId != tc.mockrepo.Output.ID {
					t.Errorf("Got Report ID %d, but expected %d", tc.mockrepo.Output.ProjectID, tc.mockrepo.ProjectId)
				}
			}

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
	ReportId      int
	ProjectId     int
	Output        *Report
	ExpectedError error
}

func (m *mockRepo) GetById(projectId int, reportId int) (*Report, error) {
	return m.Output, m.ExpectedError
}
