package request_test

import (
	"fmt"
	"net/http/httptest"
	"net/url"

	custom_middleware "github.com/haton14/departures-time/departures-time-api/middleware"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

var testHelper = helper{}

type helper struct {
}

func (h helper) createTestContext(method, path string, queryParams map[string]string) echo.Context {
	e := echo.New()
	e.Validator = custom_middleware.NewValidator()
	e.Pre(middleware.RemoveTrailingSlash())

	query := make(url.Values)
	for k, v := range queryParams {
		query.Add(k, v)
	}
	if len(query.Encode()) > 0 {
		path += "?" + query.Encode()
	}

	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()

	return e.NewContext(req, rec)
}

type customBinderForErrorTest struct {
	result error
}

func (cb customBinderForErrorTest) Bind(c echo.Context, i interface{}) error {
	return cb.result
}

func (h *helper) createTestContextBindError(method, path string, queryParams map[string]string) echo.Context {
	e := echo.New()
	e.Validator = custom_middleware.NewValidator()
	e.Binder = &customBinderForErrorTest{
		result: fmt.Errorf("bind error"),
	}
	e.Pre(middleware.RemoveTrailingSlash())

	query := make(url.Values)
	for k, v := range queryParams {
		query.Add(k, v)
	}
	if len(query.Encode()) > 0 {
		path += "?" + query.Encode()
	}

	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec)
}
