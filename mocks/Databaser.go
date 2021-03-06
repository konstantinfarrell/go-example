// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	orm "github.com/go-pg/pg/v9/orm"
	mock "github.com/stretchr/testify/mock"
)

// Databaser is an autogenerated mock type for the Databaser type
type Databaser struct {
	mock.Mock
}

// Query provides a mock function with given fields: _a0, _a1, _a2
func (_m *Databaser) Query(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
