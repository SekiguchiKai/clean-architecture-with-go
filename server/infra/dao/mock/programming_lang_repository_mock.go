// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/programming_lang_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProgrammingLangRepository is a mock of ProgrammingLangRepository interface
type MockProgrammingLangRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProgrammingLangRepositoryMockRecorder
}

// MockProgrammingLangRepositoryMockRecorder is the mock recorder for MockProgrammingLangRepository
type MockProgrammingLangRepositoryMockRecorder struct {
	mock *MockProgrammingLangRepository
}

// NewMockProgrammingLangRepository creates a new mock instance
func NewMockProgrammingLangRepository(ctrl *gomock.Controller) *MockProgrammingLangRepository {
	mock := &MockProgrammingLangRepository{ctrl: ctrl}
	mock.recorder = &MockProgrammingLangRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProgrammingLangRepository) EXPECT() *MockProgrammingLangRepositoryMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockProgrammingLangRepository) List(ctx context.Context, limit int) ([]*model.ProgrammingLang, error) {
	ret := m.ctrl.Call(m, "List", ctx, limit)
	ret0, _ := ret[0].([]*model.ProgrammingLang)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockProgrammingLangRepositoryMockRecorder) List(ctx, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProgrammingLangRepository)(nil).List), ctx, limit)
}

// Create mocks base method
func (m *MockProgrammingLangRepository) Create(ctx context.Context, lang *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	ret := m.ctrl.Call(m, "Create", ctx, lang)
	ret0, _ := ret[0].(*model.ProgrammingLang)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockProgrammingLangRepositoryMockRecorder) Create(ctx, lang interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProgrammingLangRepository)(nil).Create), ctx, lang)
}

// Read mocks base method
func (m *MockProgrammingLangRepository) Read(ctx context.Context, id int) (*model.ProgrammingLang, error) {
	ret := m.ctrl.Call(m, "Read", ctx, id)
	ret0, _ := ret[0].(*model.ProgrammingLang)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockProgrammingLangRepositoryMockRecorder) Read(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockProgrammingLangRepository)(nil).Read), ctx, id)
}

// ReadByName mocks base method
func (m *MockProgrammingLangRepository) ReadByName(ctx context.Context, name string) (*model.ProgrammingLang, error) {
	ret := m.ctrl.Call(m, "ReadByName", ctx, name)
	ret0, _ := ret[0].(*model.ProgrammingLang)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByName indicates an expected call of ReadByName
func (mr *MockProgrammingLangRepositoryMockRecorder) ReadByName(ctx, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByName", reflect.TypeOf((*MockProgrammingLangRepository)(nil).ReadByName), ctx, name)
}

// Update mocks base method
func (m *MockProgrammingLangRepository) Update(ctx context.Context, lang *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	ret := m.ctrl.Call(m, "Update", ctx, lang)
	ret0, _ := ret[0].(*model.ProgrammingLang)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockProgrammingLangRepositoryMockRecorder) Update(ctx, lang interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProgrammingLangRepository)(nil).Update), ctx, lang)
}

// Delete mocks base method
func (m *MockProgrammingLangRepository) Delete(ctx context.Context, id int) error {
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockProgrammingLangRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProgrammingLangRepository)(nil).Delete), ctx, id)
}
