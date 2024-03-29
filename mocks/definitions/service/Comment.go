// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	request "go-api/definitions/request"

	mock "github.com/stretchr/testify/mock"

	response "go-api/definitions/response"
)

// Comment is an autogenerated mock type for the Comment type
type Comment struct {
	mock.Mock
}

// List provides a mock function with given fields: req
func (_m *Comment) List(req request.CommentList) (response.CommentList, error) {
	ret := _m.Called(req)

	var r0 response.CommentList
	if rf, ok := ret.Get(0).(func(request.CommentList) response.CommentList); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.CommentList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(request.CommentList) error); ok {
		r1 = rf(req)
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
