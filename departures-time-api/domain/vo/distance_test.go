package vo_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewDistanceAndValue(t *testing.T) {
	tests := map[string]struct {
		arg      float64
		expected float64
		hasErr   error
	}{
		"[正常]:0": {
			arg:      0.0,
			expected: 0.0,
			hasErr:   nil,
		},
		"[正常]:20": {
			arg:      20.0,
			expected: 20.0,
			hasErr:   nil,
		},
		"[エラー]:-1": {
			arg:    -1.0,
			hasErr: vo.ErrMinRange,
		},
		"[エラー]:21": {
			arg:    21.0,
			hasErr: vo.ErrMaxRange,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := vo.NewDistance(tt.arg)
			if tt.hasErr != nil {
				assert.ErrorIs(t, err, tt.hasErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual.Value())
			}
		})
	}
}

func TestNewDistanceForMeterAndValue(t *testing.T) {
	tests := map[string]struct {
		arg      int
		expected float64
		hasErr   error
	}{
		"[正常]:0": {
			arg:      0,
			expected: 0.0,
			hasErr:   nil,
		},
		"[正常]:20": {
			arg:      20000,
			expected: 20.0,
			hasErr:   nil,
		},
		"[エラー]:-1": {
			arg:    -1,
			hasErr: vo.ErrMinRange,
		},
		"[エラー]:21": {
			arg:    20001,
			hasErr: vo.ErrMaxRange,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := vo.NewDistanceForMeter(tt.arg)
			if tt.hasErr != nil {
				assert.ErrorIs(t, err, tt.hasErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual.Value())
			}
		})
	}
}
