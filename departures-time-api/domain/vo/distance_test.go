package vo_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewDistanceAndValue(t *testing.T) {
	tests := map[string]struct {
		arg      int
		expected int
		hasErr   error
	}{
		"[正常]:0": {
			arg:      0,
			expected: 0,
			hasErr:   nil,
		},
		"[正常]:100000": {
			arg:      100000,
			expected: 100000,
			hasErr:   nil,
		},
		"[エラー]:-1": {
			arg:    -1,
			hasErr: vo.ErrMinRange,
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
