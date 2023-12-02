// Code generated by MockGen. DO NOT EDIT.
// Source: adapter/mongoadapter.go

// Package mock is a generated GoMock package.
package mock

import (
	adapter "common-utils/adapter"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockMongoAdapter is a mock of MongoAdapter interface.
type MockMongoAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockMongoAdapterMockRecorder
}

// MockMongoAdapterMockRecorder is the mock recorder for MockMongoAdapter.
type MockMongoAdapterMockRecorder struct {
	mock *MockMongoAdapter
}

// NewMockMongoAdapter creates a new mock instance.
func NewMockMongoAdapter(ctrl *gomock.Controller) *MockMongoAdapter {
	mock := &MockMongoAdapter{ctrl: ctrl}
	mock.recorder = &MockMongoAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMongoAdapter) EXPECT() *MockMongoAdapterMockRecorder {
	return m.recorder
}

// DeleteOne mocks base method.
func (m *MockMongoAdapter) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOne", varargs...)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockMongoAdapterMockRecorder) DeleteOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockMongoAdapter)(nil).DeleteOne), varargs...)
}

// Find mocks base method.
func (m *MockMongoAdapter) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (adapter.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(adapter.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMongoAdapterMockRecorder) Find(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMongoAdapter)(nil).Find), varargs...)
}

// FindOne mocks base method.
func (m *MockMongoAdapter) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*mongo.SingleResult)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockMongoAdapterMockRecorder) FindOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockMongoAdapter)(nil).FindOne), varargs...)
}

// InsertOne mocks base method.
func (m *MockMongoAdapter) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, document}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockMongoAdapterMockRecorder) InsertOne(ctx, document interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, document}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockMongoAdapter)(nil).InsertOne), varargs...)
}

// UpdateOne mocks base method.
func (m *MockMongoAdapter) UpdateOne(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOne", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockMongoAdapterMockRecorder) UpdateOne(ctx, filter, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockMongoAdapter)(nil).UpdateOne), varargs...)
}