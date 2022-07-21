// Code generated by MockGen. DO NOT EDIT.
// Source: ./external/nea_rest_api.go

// Package mock_external is a generated GoMock package.
package mock_external

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	vo "github.com/haton14/departures-time/departures-time-api/domain/vo"
	external "github.com/haton14/departures-time/departures-time-api/external"
)

// MockNeaRestApi is a mock of NeaRestApi interface.
type MockNeaRestApi struct {
	ctrl     *gomock.Controller
	recorder *MockNeaRestApiMockRecorder
}

// MockNeaRestApiMockRecorder is the mock recorder for MockNeaRestApi.
type MockNeaRestApiMockRecorder struct {
	mock *MockNeaRestApi
}

// NewMockNeaRestApi creates a new mock instance.
func NewMockNeaRestApi(ctrl *gomock.Controller) *MockNeaRestApi {
	mock := &MockNeaRestApi{ctrl: ctrl}
	mock.recorder = &MockNeaRestApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNeaRestApi) EXPECT() *MockNeaRestApiMockRecorder {
	return m.recorder
}

// GetNearbyStations mocks base method.
func (m *MockNeaRestApi) GetNearbyStations(lo vo.Longitude, la vo.Latitude) ([]external.NeaRestApiDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearbyStations", lo, la)
	ret0, _ := ret[0].([]external.NeaRestApiDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNearbyStations indicates an expected call of GetNearbyStations.
func (mr *MockNeaRestApiMockRecorder) GetNearbyStations(lo, la interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearbyStations", reflect.TypeOf((*MockNeaRestApi)(nil).GetNearbyStations), lo, la)
}