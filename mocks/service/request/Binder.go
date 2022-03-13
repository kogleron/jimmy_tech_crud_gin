// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Binder is an autogenerated mock type for the Binder type
type Binder struct {
	mock.Mock
}

// Bind provides a mock function with given fields: ctx, obj
func (_m *Binder) Bind(ctx *gin.Context, obj interface{}) error {
	ret := _m.Called(ctx, obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gin.Context, interface{}) error); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
