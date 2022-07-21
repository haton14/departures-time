// Code generated by MockGen. DO NOT EDIT.
// Source: repository/nearby_station_detail.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/haton14/departures-time/departures-time-api/domain/model"
)

// MockNearbyStationDetail is a mock of NearbyStationDetail interface.
type MockNearbyStationDetail struct {
	ctrl     *gomock.Controller
	recorder *MockNearbyStationDetailMockRecorder
}

// MockNearbyStationDetailMockRecorder is the mock recorder for MockNearbyStationDetail.
type MockNearbyStationDetailMockRecorder struct {
	mock *MockNearbyStationDetail
}

// NewMockNearbyStationDetail creates a new mock instance.
func NewMockNearbyStationDetail(ctrl *gomock.Controller) *MockNearbyStationDetail {
	mock := &MockNearbyStationDetail{ctrl: ctrl}
	mock.recorder = &MockNearbyStationDetailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNearbyStationDetail) EXPECT() *MockNearbyStationDetailMockRecorder {
	return m.recorder
}

// GetByNearbyStation mocks base method.
func (m *MockNearbyStationDetail) GetByNearbyStation(station model.NearbyStation) (*model.NearbyStation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNearbyStation", station)
	ret0, _ := ret[0].(*model.NearbyStation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNearbyStation indicates an expected call of GetByNearbyStation.
func (mr *MockNearbyStationDetailMockRecorder) GetByNearbyStation(station interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNearbyStation", reflect.TypeOf((*MockNearbyStationDetail)(nil).GetByNearbyStation), station)
}
