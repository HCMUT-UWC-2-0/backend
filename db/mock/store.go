// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/HCMUT-UWC-2-0/backend/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateBackOfficer mocks base method.
func (m *MockStore) CreateBackOfficer(arg0 context.Context, arg1 db.CreateBackOfficerParams) (db.BackOfficer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBackOfficer", arg0, arg1)
	ret0, _ := ret[0].(db.BackOfficer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBackOfficer indicates an expected call of CreateBackOfficer.
func (mr *MockStoreMockRecorder) CreateBackOfficer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackOfficer", reflect.TypeOf((*MockStore)(nil).CreateBackOfficer), arg0, arg1)
}

// CreateMCP mocks base method.
func (m *MockStore) CreateMCP(arg0 context.Context, arg1 db.CreateMCPParams) (db.MCP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMCP", arg0, arg1)
	ret0, _ := ret[0].(db.MCP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMCP indicates an expected call of CreateMCP.
func (mr *MockStoreMockRecorder) CreateMCP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMCP", reflect.TypeOf((*MockStore)(nil).CreateMCP), arg0, arg1)
}

// CreateRoute mocks base method.
func (m *MockStore) CreateRoute(arg0 context.Context, arg1 db.CreateRouteParams) (db.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoute", arg0, arg1)
	ret0, _ := ret[0].(db.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoute indicates an expected call of CreateRoute.
func (mr *MockStoreMockRecorder) CreateRoute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoute", reflect.TypeOf((*MockStore)(nil).CreateRoute), arg0, arg1)
}

// CreateTask mocks base method.
func (m *MockStore) CreateTask(arg0 context.Context, arg1 db.CreateTaskParams) (db.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0, arg1)
	ret0, _ := ret[0].(db.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockStoreMockRecorder) CreateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockStore)(nil).CreateTask), arg0, arg1)
}

// CreateVehicle mocks base method.
func (m *MockStore) CreateVehicle(arg0 context.Context, arg1 db.CreateVehicleParams) (db.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVehicle", arg0, arg1)
	ret0, _ := ret[0].(db.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVehicle indicates an expected call of CreateVehicle.
func (mr *MockStoreMockRecorder) CreateVehicle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVehicle", reflect.TypeOf((*MockStore)(nil).CreateVehicle), arg0, arg1)
}

// CreateWorker mocks base method.
func (m *MockStore) CreateWorker(arg0 context.Context, arg1 db.CreateWorkerParams) (db.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorker", arg0, arg1)
	ret0, _ := ret[0].(db.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorker indicates an expected call of CreateWorker.
func (mr *MockStoreMockRecorder) CreateWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorker", reflect.TypeOf((*MockStore)(nil).CreateWorker), arg0, arg1)
}

// CreateWorkerStatus mocks base method.
func (m *MockStore) CreateWorkerStatus(arg0 context.Context, arg1 int32) (db.WorkerStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkerStatus", arg0, arg1)
	ret0, _ := ret[0].(db.WorkerStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkerStatus indicates an expected call of CreateWorkerStatus.
func (mr *MockStoreMockRecorder) CreateWorkerStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkerStatus", reflect.TypeOf((*MockStore)(nil).CreateWorkerStatus), arg0, arg1)
}

// GetBackOfficer mocks base method.
func (m *MockStore) GetBackOfficer(arg0 context.Context, arg1 string) (db.BackOfficer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackOfficer", arg0, arg1)
	ret0, _ := ret[0].(db.BackOfficer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackOfficer indicates an expected call of GetBackOfficer.
func (mr *MockStoreMockRecorder) GetBackOfficer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackOfficer", reflect.TypeOf((*MockStore)(nil).GetBackOfficer), arg0, arg1)
}

// InsertTaskTx mocks base method.
func (m *MockStore) InsertTaskTx(arg0 context.Context, arg1 db.CreateTaskParams) (db.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTaskTx", arg0, arg1)
	ret0, _ := ret[0].(db.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTaskTx indicates an expected call of InsertTaskTx.
func (mr *MockStoreMockRecorder) InsertTaskTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTaskTx", reflect.TypeOf((*MockStore)(nil).InsertTaskTx), arg0, arg1)
}

// ListAllMCPs mocks base method.
func (m *MockStore) ListAllMCPs(arg0 context.Context) ([]db.MCP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllMCPs", arg0)
	ret0, _ := ret[0].([]db.MCP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllMCPs indicates an expected call of ListAllMCPs.
func (mr *MockStoreMockRecorder) ListAllMCPs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllMCPs", reflect.TypeOf((*MockStore)(nil).ListAllMCPs), arg0)
}

// ListAllVehicles mocks base method.
func (m *MockStore) ListAllVehicles(arg0 context.Context) ([]db.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllVehicles", arg0)
	ret0, _ := ret[0].([]db.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllVehicles indicates an expected call of ListAllVehicles.
func (mr *MockStoreMockRecorder) ListAllVehicles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllVehicles", reflect.TypeOf((*MockStore)(nil).ListAllVehicles), arg0)
}

// ListAllWorkers mocks base method.
func (m *MockStore) ListAllWorkers(arg0 context.Context, arg1 db.WorkerType) ([]db.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllWorkers", arg0, arg1)
	ret0, _ := ret[0].([]db.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllWorkers indicates an expected call of ListAllWorkers.
func (mr *MockStoreMockRecorder) ListAllWorkers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllWorkers", reflect.TypeOf((*MockStore)(nil).ListAllWorkers), arg0, arg1)
}

// UpdateWorkerStatus mocks base method.
func (m *MockStore) UpdateWorkerStatus(arg0 context.Context, arg1 db.UpdateWorkerStatusParams) (db.WorkerStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkerStatus", arg0, arg1)
	ret0, _ := ret[0].(db.WorkerStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkerStatus indicates an expected call of UpdateWorkerStatus.
func (mr *MockStoreMockRecorder) UpdateWorkerStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkerStatus", reflect.TypeOf((*MockStore)(nil).UpdateWorkerStatus), arg0, arg1)
}
