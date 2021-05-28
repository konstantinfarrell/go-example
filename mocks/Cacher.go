// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	redis "github.com/go-redis/redis/v8"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Cacher is an autogenerated mock type for the Cacher type
type Cacher struct {
	mock.Mock
}

// Del provides a mock function with given fields: _a0, _a1
func (_m *Cacher) Del(_a0 context.Context, _a1 ...string) *redis.IntCmd {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *redis.IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *redis.IntCmd); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.IntCmd)
		}
	}

	return r0
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *Cacher) Get(_a0 context.Context, _a1 string) *redis.StringCmd {
	ret := _m.Called(_a0, _a1)

	var r0 *redis.StringCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *redis.StringCmd); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StringCmd)
		}
	}

	return r0
}

// Set provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Cacher) Set(_a0 context.Context, _a1 string, _a2 interface{}, _a3 time.Duration) *redis.StatusCmd {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *redis.StatusCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StatusCmd)
		}
	}

	return r0
}
