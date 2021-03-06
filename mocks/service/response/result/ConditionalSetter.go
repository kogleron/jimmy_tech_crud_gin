// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// ConditionalSetter is an autogenerated mock type for the ConditionalSetter type
type ConditionalSetter struct {
	mock.Mock
}

// Set provides a mock function with given fields: ctx, status, data
func (_m *ConditionalSetter) Set(ctx *gin.Context, status int, data interface{}) {
	_m.Called(ctx, status, data)
}

// Supports provides a mock function with given fields: ctx
func (_m *ConditionalSetter) Supports(ctx *gin.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gin.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
