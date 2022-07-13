package vo_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewLatitudeAndValue(t *testing.T) {
	tests := map[string]struct {
		arg      float64
		expected float64
		hasErr   error
	}{
		"[正常]:-180.0": {
			arg:      -180.0,
			expected: -180.0,
			hasErr:   nil,
		},
		"[正常]:180.0": {
			arg:      180.0,
			expected: 180.0,
			hasErr:   nil,
		},
		"[エラー]:-180.1": {
			arg:    -180.1,
			hasErr: vo.ErrMinRange,
		},
		"[エラー]:180.1": {
			arg:    180.1,
			hasErr: vo.ErrMaxRange,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := vo.NewLatitude(tt.arg)
			if tt.hasErr != nil {
				assert.ErrorIs(t, err, tt.hasErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual.Value())
			}
		})
	}
}
