// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	context "context"

	git "github.com/argoproj-labs/argocd-autopilot/pkg/git"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Persist provides a mock function with given fields: ctx, opts
func (_m *Repository) Persist(ctx context.Context, opts *git.PushOptions) (string, error) {
	ret := _m.Called(ctx, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *git.PushOptions) string); ok {
		r0 = rf(ctx, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *git.PushOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
