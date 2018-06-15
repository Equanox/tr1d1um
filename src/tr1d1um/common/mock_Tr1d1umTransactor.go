// Code generated by mockery v1.0.0. DO NOT EDIT.
package common

import http "net/http"
import mock "github.com/stretchr/testify/mock"

// MockTr1d1umTransactor is an autogenerated mock type for the Tr1d1umTransactor type
type MockTr1d1umTransactor struct {
	mock.Mock
}

// Transact provides a mock function with given fields: _a0
func (_m *MockTr1d1umTransactor) Transact(_a0 *http.Request) (*XmidtResponse, error) {
	ret := _m.Called(_a0)

	var r0 *XmidtResponse
	if rf, ok := ret.Get(0).(func(*http.Request) *XmidtResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*XmidtResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
