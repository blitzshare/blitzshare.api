// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "blitzshare.api/app/model"
	mock "github.com/stretchr/testify/mock"
)

// EventEmit is an autogenerated mock type for the EventEmit type
type EventEmit struct {
	mock.Mock
}

// EmitP2pPeerRegistryCmd provides a mock function with given fields: queueUrl, clientId, _a2
func (_m *EventEmit) EmitP2pPeerRegistryCmd(queueUrl string, clientId string, _a2 *model.P2pPeerRegistryCmd) (string, error) {
	ret := _m.Called(queueUrl, clientId, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, *model.P2pPeerRegistryCmd) string); ok {
		r0 = rf(queueUrl, clientId, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.P2pPeerRegistryCmd) error); ok {
		r1 = rf(queueUrl, clientId, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}