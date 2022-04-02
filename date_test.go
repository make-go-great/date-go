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
