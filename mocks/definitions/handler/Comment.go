// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	request "go-api/definitions/request"
)

// Comment is an autogenerated mock type for the Comment type
type Comment struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx
func (_m *Comment) List(ctx *gin.Context) (request.CommentList, error) {
	ret := _m.Called(ctx)

	var r0 request.CommentList
	if rf, ok := ret.Get(0).(func(*gin.Context) request.CommentList); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(request.CommentList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewComment interface {
	mock.TestingT
	Cleanup(func())
}

// NewComment creates a new instance of Comment. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewComment(t mockConstructorTestingTNewComment) *Comment {
	mock := &Comment{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}