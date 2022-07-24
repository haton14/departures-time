package request_test

import (
	"net/http"
	"testing"

	"github.com/haton14/departures-time/departures-time-api/handler/request"
	"github.com/stretchr/testify/assert"
)

func TestNewRouteGet(t *testing.T) {
	t.Run("[正常]:正しくコードを指定", func(t *testing.T) {
		// setup
		c := testHelper.createTestContext(
			http.MethodGet,
			"/v1/routes",
			map[string]string{
				"from": "22566",
				"to":   "22567",
			},
		)

		expected := &request.RouteGet{
			From: "22566",
			To:   "22567",
		}
		// check
		actual, err := request.NewRouteGet(c)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	testsNG := map[string]struct {
		query map[string]string
	}{
		"[エラー]:From指定なし": {
			query: map[string]string{
				"from": "22566",
			},
		},
		"[エラー]:To指定なし": {
			query: map[string]string{
				"to": "22567"},
		},
	}

	for name, tt := range testsNG {
		t.Run(name, func(t *testing.T) {
			// setup
			c := testHelper.createTestContext(
				http.MethodGet,
				"/v1/routes",
				tt.query,
			)

			// check
			actual, err := request.NewRouteGet(c)
			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	}

	t.Run("【エラー】bindエラー", func(t *testing.T) {
		// setup
		c := testHelper.createTestContextBindError(
			http.MethodPatch,
			"/v1/routes",
			nil,
		)
		// check
		actual, err := request.NewRouteGet(c)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
