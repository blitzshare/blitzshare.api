// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Rnd is an autogenerated mock type for the Rnd type
type Rnd struct {
	mock.Mock
}

// GenerateRandomWordSequence provides a mock function with given fields:
func (_m *Rnd) GenerateRandomWordSequence() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}
