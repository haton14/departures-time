package request_test

import (
	"net/http"
	"testing"

	"github.com/haton14/departures-time/departures-time-api/handler/request"
	"github.com/stretchr/testify/assert"
)

func TestNewDestinationGet(t *testing.T) {

	t.Run("[正常]:名前が大森", func(t *testing.T) {
		// setup
		c := testHelper.createTestContext(
			http.MethodGet,
			"/v1/destinations",
			map[string]string{
				"name": "大森",
			},
		)

		expected := &request.DestinationGet{
			Name: "大森",
		}
		// check
		actual, err := request.NewDestinationGet(c)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	testsNG := map[string]struct {
		query map[string]string
	}{
		"[エラー]:名前が空文字列": {
			query: map[string]string{
				"name": "",
			},
		},
		"[エラー]:名前指定なし": {
			query: map[string]string{},
		},
	}

	for name, tt := range testsNG {
		t.Run(name, func(t *testing.T) {
			// setup
			c := testHelper.createTestContext(
				http.MethodGet,
				"/v1/destinations",
				tt.query,
			)

			// check
			actual, err := request.NewDestinationGet(c)
			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	}

	t.Run("【エラー】bindエラー", func(t *testing.T) {
		// setup
		c := testHelper.createTestContextBindError(
			http.MethodPatch,
			"/v1/destinations",
			nil,
		)
		// check
		actual, err := request.NewDestinationGet(c)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
