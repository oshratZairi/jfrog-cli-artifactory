// Code generated by MockGen. DO NOT EDIT.
// Source: command.go

// Package mock_cli is a generated GoMock package.
package mock_cli

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/jfrog/jfrog-cli-core/v2/utils/config"
)

// MockEvidenceCommands is a mock of EvidenceCommands interface.
type MockEvidenceCommands struct {
	ctrl     *gomock.Controller
	recorder *MockEvidenceCommandsMockRecorder
}

// MockEvidenceCommandsMockRecorder is the mock recorder for MockEvidenceCommands.
type MockEvidenceCommandsMockRecorder struct {
	mock *MockEvidenceCommands
}

// NewMockEvidenceCommands creates a new mock instance.
func NewMockEvidenceCommands(ctrl *gomock.Controller) *MockEvidenceCommands {
	mock := &MockEvidenceCommands{ctrl: ctrl}
	mock.recorder = &MockEvidenceCommandsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvidenceCommands) EXPECT() *MockEvidenceCommandsMockRecorder {
	return m.recorder
}

// CreateEvidence mocks base method.
func (m *MockEvidenceCommands) CreateEvidence(arg0 *config.ServerDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvidence", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEvidence indicates an expected call of CreateEvidence.
func (mr *MockEvidenceCommandsMockRecorder) CreateEvidence(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvidence", reflect.TypeOf((*MockEvidenceCommands)(nil).CreateEvidence), arg0)
}