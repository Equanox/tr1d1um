// Code generated by mockery v1.0.0. DO NOT EDIT.
package translation

import (
	mock "github.com/stretchr/testify/mock"
	common "github.com/xmidt-org/tr1d1um/common"

	wrp "github.com/xmidt-org/wrp-go/wrp"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// SendWRP provides a mock function with given fields: _a0, _a1
func (_m *MockService) SendWRP(_a0 *wrp.Message, _a1 string) (*common.XmidtResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *common.XmidtResponse
	if rf, ok := ret.Get(0).(func(*wrp.Message, string) *common.XmidtResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.XmidtResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*wrp.Message, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}