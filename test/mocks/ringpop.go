package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/temporalio/ringpop-go/events"
	"github.com/temporalio/ringpop-go/forward"
	"github.com/temporalio/ringpop-go/swim"
	"github.com/temporalio/tchannel-go"
)

type Ringpop struct {
	mock.Mock
}

// Destroy provides a mock function with given fields:
func (_m *Ringpop) Destroy() {
	_m.Called()
}

// App provides a mock function with given fields:
func (_m *Ringpop) App() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WhoAmI provides a mock function with given fields:
func (_m *Ringpop) WhoAmI() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Uptime provides a mock function with given fields:
func (_m *Ringpop) Uptime() (time.Duration, error) {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Bootstrap provides a mock function with given fields: opts
func (_m *Ringpop) Bootstrap(opts *swim.BootstrapOptions) ([]string, error) {
	ret := _m.Called(opts)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*swim.BootstrapOptions) []string); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swim.BootstrapOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Checksum provides a mock function with given fields:
func (_m *Ringpop) Checksum() (uint32, error) {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lookup provides a mock function with given fields: key
func (_m *Ringpop) Lookup(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupN provides a mock function with given fields: key, n
func (_m *Ringpop) LookupN(key string, n int) ([]string, error) {
	ret := _m.Called(key, n)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, int) []string); ok {
		r0 = rf(key, n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(key, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReachableMembers provides a mock function with given fields: predicates
func (_m *Ringpop) GetReachableMembers(predicates ...swim.MemberPredicate) ([]string, error) {
	ret := _m.Called(predicates)

	var r0 []string
	if rf, ok := ret.Get(0).(func(...swim.MemberPredicate) []string); ok {
		r0 = rf(predicates...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...swim.MemberPredicate) error); ok {
		r1 = rf(predicates...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReachableMemberObjects provides a mock function with given fields: predicates
func (_m *Ringpop) GetReachableMemberObjects(predicates ...swim.MemberPredicate) ([]swim.Member, error) {
	ret := _m.Called(predicates)

	var r0 []swim.Member
	if rf, ok := ret.Get(0).(func(...swim.MemberPredicate) []swim.Member); ok {
		r0 = rf(predicates...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]swim.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...swim.MemberPredicate) error); ok {
		r1 = rf(predicates...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountReachableMembers provides a mock function with given fields: predicates
func (_m *Ringpop) CountReachableMembers(predicates ...swim.MemberPredicate) (int, error) {
	ret := _m.Called(predicates)

	var r0 int
	if rf, ok := ret.Get(0).(func(...swim.MemberPredicate) int); ok {
		r0 = rf(predicates...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...swim.MemberPredicate) error); ok {
		r1 = rf(predicates...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleOrForward provides a mock function with given fields: key, request, response, service, endpoint, format, opts
func (_m *Ringpop) HandleOrForward(key string, request []byte, response *[]byte, service string, endpoint string, format tchannel.Format, opts *forward.Options) (bool, error) {
	ret := _m.Called(key, request, response, service, endpoint, format, opts)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, []byte, *[]byte, string, string, tchannel.Format, *forward.Options) bool); ok {
		r0 = rf(key, request, response, service, endpoint, format, opts)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte, *[]byte, string, string, tchannel.Format, *forward.Options) error); ok {
		r1 = rf(key, request, response, service, endpoint, format, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Forward provides a mock function with given fields: dest, keys, request, service, endpoint, format, opts
func (_m *Ringpop) Forward(dest string, keys []string, request []byte, service string, endpoint string, format tchannel.Format, opts *forward.Options) ([]byte, error) {
	ret := _m.Called(dest, keys, request, service, endpoint, format, opts)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, []string, []byte, string, string, tchannel.Format, *forward.Options) []byte); ok {
		r0 = rf(dest, keys, request, service, endpoint, format, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, []byte, string, string, tchannel.Format, *forward.Options) error); ok {
		r1 = rf(dest, keys, request, service, endpoint, format, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Labels provides a mock function with given fields:
func (_m *Ringpop) Labels() (*swim.NodeLabels, error) {
	ret := _m.Called()

	var r0 *swim.NodeLabels
	if rf, ok := ret.Get(0).(func() *swim.NodeLabels); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swim.NodeLabels)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddListener provides a mock function with given fields: _a0
func (_m *Ringpop) AddListener(_a0 events.EventListener) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(events.EventListener) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveListener provides a mock function with given fields: _a0
func (_m *Ringpop) RemoveListener(_a0 events.EventListener) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(events.EventListener) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RegisterListener provides a mock function with given fields: _a0
func (_m *Ringpop) RegisterListener(_a0 events.EventListener) {
	_m.Called(_a0)
}

// RegisterSelfEvictHook provides a mock function with given fields: hooks
func (_m *Ringpop) RegisterSelfEvictHook(hooks swim.SelfEvictHook) error {
	ret := _m.Called(hooks)

	var r0 error
	if rf, ok := ret.Get(0).(func(swim.SelfEvictHook) error); ok {
		r0 = rf(hooks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelfEvict provides a mock function with given fields:
func (_m *Ringpop) SelfEvict() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
