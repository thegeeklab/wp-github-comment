// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	github "github.com/google/go-github/v71/github"
	mock "github.com/stretchr/testify/mock"
)

// MockIssueService is an autogenerated mock type for the IssueService type
type MockIssueService struct {
	mock.Mock
}

type MockIssueService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIssueService) EXPECT() *MockIssueService_Expecter {
	return &MockIssueService_Expecter{mock: &_m.Mock}
}

// CreateComment provides a mock function with given fields: ctx, owner, repo, number, comment
func (_m *MockIssueService) CreateComment(ctx context.Context, owner string, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, number, comment)

	if len(ret) == 0 {
		panic("no return value specified for CreateComment")
	}

	var r0 *github.IssueComment
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, *github.IssueComment) (*github.IssueComment, *github.Response, error)); ok {
		return rf(ctx, owner, repo, number, comment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, *github.IssueComment) *github.IssueComment); ok {
		r0 = rf(ctx, owner, repo, number, comment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.IssueComment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, *github.IssueComment) *github.Response); ok {
		r1 = rf(ctx, owner, repo, number, comment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, int, *github.IssueComment) error); ok {
		r2 = rf(ctx, owner, repo, number, comment)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockIssueService_CreateComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateComment'
type MockIssueService_CreateComment_Call struct {
	*mock.Call
}

// CreateComment is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - number int
//   - comment *github.IssueComment
func (_e *MockIssueService_Expecter) CreateComment(ctx interface{}, owner interface{}, repo interface{}, number interface{}, comment interface{}) *MockIssueService_CreateComment_Call {
	return &MockIssueService_CreateComment_Call{Call: _e.mock.On("CreateComment", ctx, owner, repo, number, comment)}
}

func (_c *MockIssueService_CreateComment_Call) Run(run func(ctx context.Context, owner string, repo string, number int, comment *github.IssueComment)) *MockIssueService_CreateComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(int), args[4].(*github.IssueComment))
	})
	return _c
}

func (_c *MockIssueService_CreateComment_Call) Return(_a0 *github.IssueComment, _a1 *github.Response, _a2 error) *MockIssueService_CreateComment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockIssueService_CreateComment_Call) RunAndReturn(run func(context.Context, string, string, int, *github.IssueComment) (*github.IssueComment, *github.Response, error)) *MockIssueService_CreateComment_Call {
	_c.Call.Return(run)
	return _c
}

// EditComment provides a mock function with given fields: ctx, owner, repo, commentID, comment
func (_m *MockIssueService) EditComment(ctx context.Context, owner string, repo string, commentID int64, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, commentID, comment)

	if len(ret) == 0 {
		panic("no return value specified for EditComment")
	}

	var r0 *github.IssueComment
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, *github.IssueComment) (*github.IssueComment, *github.Response, error)); ok {
		return rf(ctx, owner, repo, commentID, comment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, *github.IssueComment) *github.IssueComment); ok {
		r0 = rf(ctx, owner, repo, commentID, comment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.IssueComment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64, *github.IssueComment) *github.Response); ok {
		r1 = rf(ctx, owner, repo, commentID, comment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, int64, *github.IssueComment) error); ok {
		r2 = rf(ctx, owner, repo, commentID, comment)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockIssueService_EditComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditComment'
type MockIssueService_EditComment_Call struct {
	*mock.Call
}

// EditComment is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - commentID int64
//   - comment *github.IssueComment
func (_e *MockIssueService_Expecter) EditComment(ctx interface{}, owner interface{}, repo interface{}, commentID interface{}, comment interface{}) *MockIssueService_EditComment_Call {
	return &MockIssueService_EditComment_Call{Call: _e.mock.On("EditComment", ctx, owner, repo, commentID, comment)}
}

func (_c *MockIssueService_EditComment_Call) Run(run func(ctx context.Context, owner string, repo string, commentID int64, comment *github.IssueComment)) *MockIssueService_EditComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(int64), args[4].(*github.IssueComment))
	})
	return _c
}

func (_c *MockIssueService_EditComment_Call) Return(_a0 *github.IssueComment, _a1 *github.Response, _a2 error) *MockIssueService_EditComment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockIssueService_EditComment_Call) RunAndReturn(run func(context.Context, string, string, int64, *github.IssueComment) (*github.IssueComment, *github.Response, error)) *MockIssueService_EditComment_Call {
	_c.Call.Return(run)
	return _c
}

// ListComments provides a mock function with given fields: ctx, owner, repo, number, opts
func (_m *MockIssueService) ListComments(ctx context.Context, owner string, repo string, number int, opts *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, number, opts)

	if len(ret) == 0 {
		panic("no return value specified for ListComments")
	}

	var r0 []*github.IssueComment
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error)); ok {
		return rf(ctx, owner, repo, number, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, *github.IssueListCommentsOptions) []*github.IssueComment); ok {
		r0 = rf(ctx, owner, repo, number, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.IssueComment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, *github.IssueListCommentsOptions) *github.Response); ok {
		r1 = rf(ctx, owner, repo, number, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, int, *github.IssueListCommentsOptions) error); ok {
		r2 = rf(ctx, owner, repo, number, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockIssueService_ListComments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListComments'
type MockIssueService_ListComments_Call struct {
	*mock.Call
}

// ListComments is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - number int
//   - opts *github.IssueListCommentsOptions
func (_e *MockIssueService_Expecter) ListComments(ctx interface{}, owner interface{}, repo interface{}, number interface{}, opts interface{}) *MockIssueService_ListComments_Call {
	return &MockIssueService_ListComments_Call{Call: _e.mock.On("ListComments", ctx, owner, repo, number, opts)}
}

func (_c *MockIssueService_ListComments_Call) Run(run func(ctx context.Context, owner string, repo string, number int, opts *github.IssueListCommentsOptions)) *MockIssueService_ListComments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(int), args[4].(*github.IssueListCommentsOptions))
	})
	return _c
}

func (_c *MockIssueService_ListComments_Call) Return(_a0 []*github.IssueComment, _a1 *github.Response, _a2 error) *MockIssueService_ListComments_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockIssueService_ListComments_Call) RunAndReturn(run func(context.Context, string, string, int, *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error)) *MockIssueService_ListComments_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIssueService creates a new instance of MockIssueService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIssueService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIssueService {
	mock := &MockIssueService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
