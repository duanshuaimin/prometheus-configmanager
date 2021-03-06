/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	alert "prometheus-configmanager/prometheus/alert"

	mock "github.com/stretchr/testify/mock"

	rulefmt "github.com/prometheus/prometheus/pkg/rulefmt"
)

// PrometheusAlertClient is an autogenerated mock type for the PrometheusAlertClient type
type PrometheusAlertClient struct {
	mock.Mock
}

// BulkUpdateRules provides a mock function with given fields: filePrefix, rules
func (_m *PrometheusAlertClient) BulkUpdateRules(filePrefix string, rules []rulefmt.Rule) (alert.BulkUpdateResults, error) {
	ret := _m.Called(filePrefix, rules)

	var r0 alert.BulkUpdateResults
	if rf, ok := ret.Get(0).(func(string, []rulefmt.Rule) alert.BulkUpdateResults); ok {
		r0 = rf(filePrefix, rules)
	} else {
		r0 = ret.Get(0).(alert.BulkUpdateResults)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []rulefmt.Rule) error); ok {
		r1 = rf(filePrefix, rules)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRule provides a mock function with given fields: filePrefix, ruleName
func (_m *PrometheusAlertClient) DeleteRule(filePrefix string, ruleName string) error {
	ret := _m.Called(filePrefix, ruleName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(filePrefix, ruleName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadRules provides a mock function with given fields: filePrefix, ruleName
func (_m *PrometheusAlertClient) ReadRules(filePrefix string, ruleName string) ([]rulefmt.Rule, error) {
	ret := _m.Called(filePrefix, ruleName)

	var r0 []rulefmt.Rule
	if rf, ok := ret.Get(0).(func(string, string) []rulefmt.Rule); ok {
		r0 = rf(filePrefix, ruleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rulefmt.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(filePrefix, ruleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReloadPrometheus provides a mock function with given fields:
func (_m *PrometheusAlertClient) ReloadPrometheus() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RuleExists provides a mock function with given fields: filePrefix, rulename
func (_m *PrometheusAlertClient) RuleExists(filePrefix string, rulename string) bool {
	ret := _m.Called(filePrefix, rulename)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(filePrefix, rulename)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Tenancy provides a mock function with given fields:
func (_m *PrometheusAlertClient) Tenancy() alert.TenancyConfig {
	ret := _m.Called()

	var r0 alert.TenancyConfig
	if rf, ok := ret.Get(0).(func() alert.TenancyConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(alert.TenancyConfig)
	}

	return r0
}

// UpdateRule provides a mock function with given fields: filePrefix, rule
func (_m *PrometheusAlertClient) UpdateRule(filePrefix string, rule rulefmt.Rule) error {
	ret := _m.Called(filePrefix, rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, rulefmt.Rule) error); ok {
		r0 = rf(filePrefix, rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateRule provides a mock function with given fields: rule
func (_m *PrometheusAlertClient) ValidateRule(rule rulefmt.Rule) error {
	ret := _m.Called(rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(rulefmt.Rule) error); ok {
		r0 = rf(rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteRule provides a mock function with given fields: filePrefix, rule
func (_m *PrometheusAlertClient) WriteRule(filePrefix string, rule rulefmt.Rule) error {
	ret := _m.Called(filePrefix, rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, rulefmt.Rule) error); ok {
		r0 = rf(filePrefix, rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
