// Code generated by mockery v2.1.0. DO NOT EDIT.

package dao

import (
	context "context"

	dao "github.com/goharbor/harbor/src/pkg/accessory/dao"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// DAO is an autogenerated mock type for the DAO type
type DAO struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *DAO) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, accessory
func (_m *DAO) Create(ctx context.Context, accessory *dao.Accessory) (int64, error) {
	ret := _m.Called(ctx, accessory)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *dao.Accessory) int64); ok {
		r0 = rf(ctx, accessory)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dao.Accessory) error); ok {
		r1 = rf(ctx, accessory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *DAO) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOfArtifact provides a mock function with given fields: ctx, subArtifactID
func (_m *DAO) DeleteOfArtifact(ctx context.Context, subArtifactID int64) error {
	ret := _m.Called(ctx, subArtifactID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, subArtifactID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *DAO) Get(ctx context.Context, id int64) (*dao.Accessory, error) {
	ret := _m.Called(ctx, id)

	var r0 *dao.Accessory
	if rf, ok := ret.Get(0).(func(context.Context, int64) *dao.Accessory); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.Accessory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *DAO) List(ctx context.Context, query *q.Query) ([]*dao.Accessory, error) {
	ret := _m.Called(ctx, query)

	var r0 []*dao.Accessory
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*dao.Accessory); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dao.Accessory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
