package vo_test

import (
	"testing"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewStationNameAndValue(t *testing.T) {
	tests := map[string]struct {
		arg      string
		expected string
		hasErr   error
	}{
		"[正常]:東京": {
			arg:      "東京",
			expected: "東京",
			hasErr:   nil,
		},
		"[エラー]:空文字列": {
			arg:    "",
			hasErr: vo.ErrMinLength,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := vo.NewStationName(tt.arg)
			if tt.hasErr != nil {
				assert.ErrorIs(t, err, tt.hasErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual.Value())
			}
		})
	}
}
