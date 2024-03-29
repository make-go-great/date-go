package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToRFC3339(t *testing.T) {
	tests := []struct {
		name        string
		date        string
		wantResult  string
		wantIsError bool
	}{
		{
			name:       "YYYY-MM-DD",
			date:       "1998-01-18",
			wantResult: "1998-01-18T00:00:00Z",
		},
		{
			name:       "YYYY/MM/DD",
			date:       "1998/04/01",
			wantResult: "1998-04-01T00:00:00Z",
		},
		{
			name:       "YYYY-MM-DD",
			date:       "1998.01.18",
			wantResult: "1998-01-18T00:00:00Z",
		},
		{
			name:       "YYYY MM DD",
			date:       "1998 04 01",
			wantResult: "1998-04-01T00:00:00Z",
		},
		{
			name:       "YYYYMMDD",
			date:       "19980401",
			wantResult: "1998-04-01T00:00:00Z",
		},
		{
			name:       "DD-MM-YYYY",
			date:       "18-01-1998",
			wantResult: "1998-01-18T00:00:00Z",
		},
		{
			name:       "DD/MM/YYYY",
			date:       "01/04/1998",
			wantResult: "1998-04-01T00:00:00Z",
		},
		{
			name:       "DD.MM.YYYY",
			date:       "18.01.1998",
			wantResult: "1998-01-18T00:00:00Z",
		},
		{
			name:       "DD MM YYYY",
			date:       "01 04 1998",
			wantResult: "1998-04-01T00:00:00Z",
		},
		{
			name:       "DDMMYYYY",
			date:       "01041998",
			wantResult: "1998-04-01T00:00:00Z",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotErr := ToRFC3339(tc.date, time.UTC)
			if gotErr != nil {
				assert.Equal(t, tc.wantIsError, true)
				return
			}
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}

func TestFromRFC3339(t *testing.T) {
	tests := []struct {
		name        string
		rfc3339     string
		wantResult  string
		wantIsError bool
	}{
		{
			name:       "1998-04-01T00:00:00Z",
			rfc3339:    "1998-04-01T00:00:00Z",
			wantResult: "1998-04-01",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotErr := FromRFC3339(tc.rfc3339, time.UTC)
			if gotErr != nil {
				assert.Equal(t, tc.wantIsError, true)
				return
			}
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}

func TestToDefaultDate(t *testing.T) {
	tests := []struct {
		name        string
		date        string
		wantResult  string
		wantIsError bool
	}{
		{
			name:       "YYYY-MM-DD",
			date:       "1998-01-18",
			wantResult: "1998-01-18",
		},
		{
			name:       "YYYY/MM/DD",
			date:       "1998/04/01",
			wantResult: "1998-04-01",
		},
		{
			name:       "YYYY-MM-DD",
			date:       "1998.01.18",
			wantResult: "1998-01-18",
		},
		{
			name:       "YYYY MM DD",
			date:       "1998 04 01",
			wantResult: "1998-04-01",
		},
		{
			name:       "YYYYMMDD",
			date:       "19980401",
			wantResult: "1998-04-01",
		},
		{
			name:       "DD-MM-YYYY",
			date:       "18-01-1998",
			wantResult: "1998-01-18",
		},
		{
			name:       "DD/MM/YYYY",
			date:       "01/04/1998",
			wantResult: "1998-04-01",
		},
		{
			name:       "DD.MM.YYYY",
			date:       "18.01.1998",
			wantResult: "1998-01-18",
		},
		{
			name:       "DD MM YYYY",
			date:       "01 04 1998",
			wantResult: "1998-04-01",
		},
		{
			name:       "DDMMYYYY",
			date:       "01041998",
			wantResult: "1998-04-01",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotErr := ToDefaultDate(tc.date, time.UTC)
			if gotErr != nil {
				assert.Equal(t, tc.wantIsError, true)
				return
			}
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}

func TestFormatDateByDefault(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{
			name: "1997-04-01",
			t:    time.Date(1997, 4, 1, 0, 0, 0, 0, time.UTC),
			want: "1997-04-01",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatDateByDefault(tc.t, time.UTC)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFormatDateTimeByDefault(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{
			name: "1997-04-01 01:02:03",
			t:    time.Date(1997, 4, 1, 1, 2, 3, 0, time.UTC),
			want: "1997-04-01 01:02:03",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatDateTimeByDefault(tc.t, time.UTC)
			assert.Equal(t, tc.want, got)
		})
	}
}
