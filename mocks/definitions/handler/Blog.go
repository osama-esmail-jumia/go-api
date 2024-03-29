// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	request "go-api/definitions/request"
)

// Blog is an autogenerated mock type for the Blog type
type Blog struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx
func (_m *Blog) Create(ctx *gin.Context) (request.BlogCreate, error) {
	ret := _m.Called(ctx)

	var r0 request.BlogCreate
	if rf, ok := ret.Get(0).(func(*gin.Context) request.BlogCreate); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(request.BlogCreate)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx
func (_m *Blog) Delete(ctx *gin.Context) (request.BlogDelete, error) {
	ret := _m.Called(ctx)

	var r0 request.BlogDelete
	if rf, ok := ret.Get(0).(func(*gin.Context) request.BlogDelete); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(request.BlogDelete)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *Blog) List(ctx *gin.Context) (request.BlogList, error) {
	ret := _m.Called(ctx)

	var r0 request.BlogList
	if rf, ok := ret.Get(0).(func(*gin.Context) request.BlogList); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(request.BlogList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx
func (_m *Blog) Update(ctx *gin.Context) (request.BlogUpdate, error) {
	ret := _m.Called(ctx)

	var r0 request.BlogUpdate
	if rf, ok := ret.Get(0).(func(*gin.Context) request.BlogUpdate); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(request.BlogUpdate)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBlog interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlog creates a new instance of Blog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlog(t mockConstructorTestingTNewBlog) *Blog {
	mock := &Blog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
