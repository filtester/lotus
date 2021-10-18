// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	os "os"
	reflect "reflect"

	abi "github.com/filecoin-project/go-state-types/abi"
	fsutil "github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	partialfile "github.com/filecoin-project/lotus/extern/sector-storage/partialfile"
	stores "github.com/filecoin-project/lotus/extern/sector-storage/stores"
	storiface "github.com/filecoin-project/lotus/extern/sector-storage/storiface"
	storage "github.com/filecoin-project/specs-storage/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockpartialFileHandler is a mock of partialFileHandler interface.
type MockpartialFileHandler struct {
	ctrl     *gomock.Controller
	recorder *MockpartialFileHandlerMockRecorder
}

// MockpartialFileHandlerMockRecorder is the mock recorder for MockpartialFileHandler.
type MockpartialFileHandlerMockRecorder struct {
	mock *MockpartialFileHandler
}

// NewMockpartialFileHandler creates a new mock instance.
func NewMockpartialFileHandler(ctrl *gomock.Controller) *MockpartialFileHandler {
	mock := &MockpartialFileHandler{ctrl: ctrl}
	mock.recorder = &MockpartialFileHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpartialFileHandler) EXPECT() *MockpartialFileHandlerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockpartialFileHandler) Close(pf *partialfile.PartialFile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", pf)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockpartialFileHandlerMockRecorder) Close(pf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockpartialFileHandler)(nil).Close), pf)
}

// HasAllocated mocks base method.
func (m *MockpartialFileHandler) HasAllocated(pf *partialfile.PartialFile, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAllocated", pf, offset, size)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAllocated indicates an expected call of HasAllocated.
func (mr *MockpartialFileHandlerMockRecorder) HasAllocated(pf, offset, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAllocated", reflect.TypeOf((*MockpartialFileHandler)(nil).HasAllocated), pf, offset, size)
}

// OpenPartialFile mocks base method.
func (m *MockpartialFileHandler) OpenPartialFile(maxPieceSize abi.PaddedPieceSize, path string) (*partialfile.PartialFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenPartialFile", maxPieceSize, path)
	ret0, _ := ret[0].(*partialfile.PartialFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenPartialFile indicates an expected call of OpenPartialFile.
func (mr *MockpartialFileHandlerMockRecorder) OpenPartialFile(maxPieceSize, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPartialFile", reflect.TypeOf((*MockpartialFileHandler)(nil).OpenPartialFile), maxPieceSize, path)
}

// Reader mocks base method.
func (m *MockpartialFileHandler) Reader(pf *partialfile.PartialFile, offset storiface.PaddedByteIndex, size abi.PaddedPieceSize) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader", pf, offset, size)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader.
func (mr *MockpartialFileHandlerMockRecorder) Reader(pf, offset, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockpartialFileHandler)(nil).Reader), pf, offset, size)
}

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

// AcquireSector mocks base method.
func (m *MockStore) AcquireSector(ctx context.Context, s storage.SectorRef, existing, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (storiface.SectorPaths, storiface.SectorPaths, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireSector", ctx, s, existing, allocate, sealing, op)
	ret0, _ := ret[0].(storiface.SectorPaths)
	ret1, _ := ret[1].(storiface.SectorPaths)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AcquireSector indicates an expected call of AcquireSector.
func (mr *MockStoreMockRecorder) AcquireSector(ctx, s, existing, allocate, sealing, op interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireSector", reflect.TypeOf((*MockStore)(nil).AcquireSector), ctx, s, existing, allocate, sealing, op)
}

// FsStat mocks base method.
func (m *MockStore) FsStat(ctx context.Context, id stores.ID) (fsutil.FsStat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FsStat", ctx, id)
	ret0, _ := ret[0].(fsutil.FsStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FsStat indicates an expected call of FsStat.
func (mr *MockStoreMockRecorder) FsStat(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FsStat", reflect.TypeOf((*MockStore)(nil).FsStat), ctx, id)
}

// MoveStorage mocks base method.
func (m *MockStore) MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveStorage", ctx, s, types)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveStorage indicates an expected call of MoveStorage.
func (mr *MockStoreMockRecorder) MoveStorage(ctx, s, types interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveStorage", reflect.TypeOf((*MockStore)(nil).MoveStorage), ctx, s, types)
}

// Remove mocks base method.
func (m *MockStore) Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, s, types, force)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockStoreMockRecorder) Remove(ctx, s, types, force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStore)(nil).Remove), ctx, s, types, force)
}

// RemoveCopies mocks base method.
func (m *MockStore) RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCopies", ctx, s, types)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCopies indicates an expected call of RemoveCopies.
func (mr *MockStoreMockRecorder) RemoveCopies(ctx, s, types interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCopies", reflect.TypeOf((*MockStore)(nil).RemoveCopies), ctx, s, types)
}

// Reserve mocks base method.
func (m *MockStore) Reserve(ctx context.Context, sid storage.SectorRef, ft storiface.SectorFileType, storageIDs storiface.SectorPaths, overheadTab map[storiface.SectorFileType]int) (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reserve", ctx, sid, ft, storageIDs, overheadTab)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reserve indicates an expected call of Reserve.
func (mr *MockStoreMockRecorder) Reserve(ctx, sid, ft, storageIDs, overheadTab interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reserve", reflect.TypeOf((*MockStore)(nil).Reserve), ctx, sid, ft, storageIDs, overheadTab)
}
