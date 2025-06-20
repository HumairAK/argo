// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	labels "k8s.io/apimachinery/pkg/labels"

	time "time"

	utils "github.com/argoproj/argo-workflows/v3/server/utils"

	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// WorkflowArchive is an autogenerated mock type for the WorkflowArchive type
type WorkflowArchive struct {
	mock.Mock
}

// ArchiveWorkflow provides a mock function with given fields: wf
func (_m *WorkflowArchive) ArchiveWorkflow(wf *v1alpha1.Workflow) error {
	ret := _m.Called(wf)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Workflow) error); ok {
		r0 = rf(wf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountWorkflows provides a mock function with given fields: options
func (_m *WorkflowArchive) CountWorkflows(options utils.ListOptions) (int64, error) {
	ret := _m.Called(options)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(utils.ListOptions) (int64, error)); ok {
		return rf(options)
	}
	if rf, ok := ret.Get(0).(func(utils.ListOptions) int64); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(utils.ListOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteExpiredWorkflows provides a mock function with given fields: ttl
func (_m *WorkflowArchive) DeleteExpiredWorkflows(ttl time.Duration) error {
	ret := _m.Called(ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWorkflow provides a mock function with given fields: uid
func (_m *WorkflowArchive) DeleteWorkflow(uid string) error {
	ret := _m.Called(uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetWorkflow provides a mock function with given fields: uid, namespace, name
func (_m *WorkflowArchive) GetWorkflow(uid string, namespace string, name string) (*v1alpha1.Workflow, error) {
	ret := _m.Called(uid, namespace, name)

	var r0 *v1alpha1.Workflow
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*v1alpha1.Workflow, error)); ok {
		return rf(uid, namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *v1alpha1.Workflow); ok {
		r0 = rf(uid, namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(uid, namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowForEstimator provides a mock function with given fields: namespace, requirements
func (_m *WorkflowArchive) GetWorkflowForEstimator(namespace string, requirements []labels.Requirement) (*v1alpha1.Workflow, error) {
	ret := _m.Called(namespace, requirements)

	var r0 *v1alpha1.Workflow
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []labels.Requirement) (*v1alpha1.Workflow, error)); ok {
		return rf(namespace, requirements)
	}
	if rf, ok := ret.Get(0).(func(string, []labels.Requirement) *v1alpha1.Workflow); ok {
		r0 = rf(namespace, requirements)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []labels.Requirement) error); ok {
		r1 = rf(namespace, requirements)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEnabled provides a mock function with given fields:
func (_m *WorkflowArchive) IsEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListWorkflows provides a mock function with given fields: options
func (_m *WorkflowArchive) ListWorkflows(options utils.ListOptions) (v1alpha1.Workflows, error) {
	ret := _m.Called(options)

	var r0 v1alpha1.Workflows
	var r1 error
	if rf, ok := ret.Get(0).(func(utils.ListOptions) (v1alpha1.Workflows, error)); ok {
		return rf(options)
	}
	if rf, ok := ret.Get(0).(func(utils.ListOptions) v1alpha1.Workflows); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.Workflows)
		}
	}

	if rf, ok := ret.Get(1).(func(utils.ListOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkflowsLabelKeys provides a mock function with given fields:
func (_m *WorkflowArchive) ListWorkflowsLabelKeys() (*v1alpha1.LabelKeys, error) {
	ret := _m.Called()

	var r0 *v1alpha1.LabelKeys
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1alpha1.LabelKeys, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1alpha1.LabelKeys); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.LabelKeys)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkflowsLabelValues provides a mock function with given fields: key
func (_m *WorkflowArchive) ListWorkflowsLabelValues(key string) (*v1alpha1.LabelValues, error) {
	ret := _m.Called(key)

	var r0 *v1alpha1.LabelValues
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v1alpha1.LabelValues, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.LabelValues); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.LabelValues)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWorkflowArchive interface {
	mock.TestingT
	Cleanup(func())
}

// NewWorkflowArchive creates a new instance of WorkflowArchive. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWorkflowArchive(t mockConstructorTestingTNewWorkflowArchive) *WorkflowArchive {
	mock := &WorkflowArchive{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
