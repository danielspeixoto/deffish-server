// Code generated by MockGen. DO NOT EDIT.
// Source: Boundary.go

// Package essay is a generated GoMock package.
package essay

import (
	aggregates "deffish-server/src/aggregates"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIRepository is a mock of IRepository interface
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockIRepository) Insert(essay aggregates.Essay) (aggregates.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", essay)
	ret0, _ := ret[0].(aggregates.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockIRepositoryMockRecorder) Insert(essay interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIRepository)(nil).Insert), essay)
}

// Id mocks base method
func (m *MockIRepository) Id(id aggregates.Id) (aggregates.Essay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id", id)
	ret0, _ := ret[0].(aggregates.Essay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Id indicates an expected call of Id
func (mr *MockIRepositoryMockRecorder) Id(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockIRepository)(nil).Id), id)
}

// FilterByTopic mocks base method
func (m *MockIRepository) FilterByTopic(arg0 aggregates.Id) ([]aggregates.Essay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterByTopic", arg0)
	ret0, _ := ret[0].([]aggregates.Essay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterByTopic indicates an expected call of FilterByTopic
func (mr *MockIRepositoryMockRecorder) FilterByTopic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByTopic", reflect.TypeOf((*MockIRepository)(nil).FilterByTopic), arg0)
}

// Comment mocks base method
func (m *MockIRepository) Comment(essayId aggregates.Id, comment aggregates.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Comment", essayId, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Comment indicates an expected call of Comment
func (mr *MockIRepositoryMockRecorder) Comment(essayId, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Comment", reflect.TypeOf((*MockIRepository)(nil).Comment), essayId, comment)
}

// random mocks base method
func (m *MockIRepository) Random(amount int) ([]aggregates.Essay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "random", amount)
	ret0, _ := ret[0].([]aggregates.Essay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// random indicates an expected call of random
func (mr *MockIRepositoryMockRecorder) Random(amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "random", reflect.TypeOf((*MockIRepository)(nil).Random), amount)
}

// MockIRandomUseCase is a mock of IRandomUseCase interface
type MockIRandomUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIRandomUseCaseMockRecorder
}

// MockIRandomUseCaseMockRecorder is the mock recorder for MockIRandomUseCase
type MockIRandomUseCaseMockRecorder struct {
	mock *MockIRandomUseCase
}

// NewMockIRandomUseCase creates a new mock instance
func NewMockIRandomUseCase(ctrl *gomock.Controller) *MockIRandomUseCase {
	mock := &MockIRandomUseCase{ctrl: ctrl}
	mock.recorder = &MockIRandomUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRandomUseCase) EXPECT() *MockIRandomUseCaseMockRecorder {
	return m.recorder
}

// random mocks base method
func (m *MockIRandomUseCase) Random(amount int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "random", amount)
}

// random indicates an expected call of random
func (mr *MockIRandomUseCaseMockRecorder) Random(amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "random", reflect.TypeOf((*MockIRandomUseCase)(nil).Random), amount)
}

// MockIRandomPresenter is a mock of IRandomPresenter interface
type MockIRandomPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockIRandomPresenterMockRecorder
}

// MockIRandomPresenterMockRecorder is the mock recorder for MockIRandomPresenter
type MockIRandomPresenterMockRecorder struct {
	mock *MockIRandomPresenter
}

// NewMockIRandomPresenter creates a new mock instance
func NewMockIRandomPresenter(ctrl *gomock.Controller) *MockIRandomPresenter {
	mock := &MockIRandomPresenter{ctrl: ctrl}
	mock.recorder = &MockIRandomPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRandomPresenter) EXPECT() *MockIRandomPresenterMockRecorder {
	return m.recorder
}

// OnListReceived mocks base method
func (m *MockIRandomPresenter) OnListReceived(arg0 []aggregates.Essay) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnListReceived", arg0)
}

// OnListReceived indicates an expected call of OnListReceived
func (mr *MockIRandomPresenterMockRecorder) OnListReceived(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnListReceived", reflect.TypeOf((*MockIRandomPresenter)(nil).OnListReceived), arg0)
}

// OnError mocks base method
func (m *MockIRandomPresenter) OnError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnError", arg0)
}

// OnError indicates an expected call of OnError
func (mr *MockIRandomPresenterMockRecorder) OnError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockIRandomPresenter)(nil).OnError), arg0)
}

// MockIByIdUseCase is a mock of IByIdUseCase interface
type MockIByIdUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIByIdUseCaseMockRecorder
}

// MockIByIdUseCaseMockRecorder is the mock recorder for MockIByIdUseCase
type MockIByIdUseCaseMockRecorder struct {
	mock *MockIByIdUseCase
}

// NewMockIByIdUseCase creates a new mock instance
func NewMockIByIdUseCase(ctrl *gomock.Controller) *MockIByIdUseCase {
	mock := &MockIByIdUseCase{ctrl: ctrl}
	mock.recorder = &MockIByIdUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIByIdUseCase) EXPECT() *MockIByIdUseCaseMockRecorder {
	return m.recorder
}

// Id mocks base method
func (m *MockIByIdUseCase) Id(id aggregates.Id) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Id", id)
}

// Id indicates an expected call of Id
func (mr *MockIByIdUseCaseMockRecorder) Id(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockIByIdUseCase)(nil).Id), id)
}

// MockIByIdPresenter is a mock of IByIdPresenter interface
type MockIByIdPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockIByIdPresenterMockRecorder
}

// MockIByIdPresenterMockRecorder is the mock recorder for MockIByIdPresenter
type MockIByIdPresenterMockRecorder struct {
	mock *MockIByIdPresenter
}

// NewMockIByIdPresenter creates a new mock instance
func NewMockIByIdPresenter(ctrl *gomock.Controller) *MockIByIdPresenter {
	mock := &MockIByIdPresenter{ctrl: ctrl}
	mock.recorder = &MockIByIdPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIByIdPresenter) EXPECT() *MockIByIdPresenterMockRecorder {
	return m.recorder
}

// OnReceived mocks base method
func (m *MockIByIdPresenter) OnReceived(arg0 aggregates.Essay) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnReceived", arg0)
}

// OnReceived indicates an expected call of OnReceived
func (mr *MockIByIdPresenterMockRecorder) OnReceived(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnReceived", reflect.TypeOf((*MockIByIdPresenter)(nil).OnReceived), arg0)
}

// OnError mocks base method
func (m *MockIByIdPresenter) OnError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnError", arg0)
}

// OnError indicates an expected call of OnError
func (mr *MockIByIdPresenterMockRecorder) OnError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockIByIdPresenter)(nil).OnError), arg0)
}

// MockIUploadUseCase is a mock of IUploadUseCase interface
type MockIUploadUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIUploadUseCaseMockRecorder
}

// MockIUploadUseCaseMockRecorder is the mock recorder for MockIUploadUseCase
type MockIUploadUseCaseMockRecorder struct {
	mock *MockIUploadUseCase
}

// NewMockIUploadUseCase creates a new mock instance
func NewMockIUploadUseCase(ctrl *gomock.Controller) *MockIUploadUseCase {
	mock := &MockIUploadUseCase{ctrl: ctrl}
	mock.recorder = &MockIUploadUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUploadUseCase) EXPECT() *MockIUploadUseCaseMockRecorder {
	return m.recorder
}

// Upload mocks base method
func (m *MockIUploadUseCase) Upload(arg0 aggregates.Essay) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Upload", arg0)
}

// Upload indicates an expected call of Upload
func (mr *MockIUploadUseCaseMockRecorder) Upload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockIUploadUseCase)(nil).Upload), arg0)
}

// MockIUploadPresenter is a mock of IUploadPresenter interface
type MockIUploadPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockIUploadPresenterMockRecorder
}

// MockIUploadPresenterMockRecorder is the mock recorder for MockIUploadPresenter
type MockIUploadPresenterMockRecorder struct {
	mock *MockIUploadPresenter
}

// NewMockIUploadPresenter creates a new mock instance
func NewMockIUploadPresenter(ctrl *gomock.Controller) *MockIUploadPresenter {
	mock := &MockIUploadPresenter{ctrl: ctrl}
	mock.recorder = &MockIUploadPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUploadPresenter) EXPECT() *MockIUploadPresenterMockRecorder {
	return m.recorder
}

// OnUploaded mocks base method
func (m *MockIUploadPresenter) OnUploaded() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnUploaded")
}

// OnUploaded indicates an expected call of OnUploaded
func (mr *MockIUploadPresenterMockRecorder) OnUploaded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUploaded", reflect.TypeOf((*MockIUploadPresenter)(nil).OnUploaded))
}

// OnError mocks base method
func (m *MockIUploadPresenter) OnError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnError", arg0)
}

// OnError indicates an expected call of OnError
func (mr *MockIUploadPresenterMockRecorder) OnError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockIUploadPresenter)(nil).OnError), arg0)
}

// MockIFilterByTopicPresenter is a mock of IFilterByTopicPresenter interface
type MockIFilterByTopicPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockIFilterByTopicPresenterMockRecorder
}

// MockIFilterByTopicPresenterMockRecorder is the mock recorder for MockIFilterByTopicPresenter
type MockIFilterByTopicPresenterMockRecorder struct {
	mock *MockIFilterByTopicPresenter
}

// NewMockIFilterByTopicPresenter creates a new mock instance
func NewMockIFilterByTopicPresenter(ctrl *gomock.Controller) *MockIFilterByTopicPresenter {
	mock := &MockIFilterByTopicPresenter{ctrl: ctrl}
	mock.recorder = &MockIFilterByTopicPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIFilterByTopicPresenter) EXPECT() *MockIFilterByTopicPresenterMockRecorder {
	return m.recorder
}

// OnListReceived mocks base method
func (m *MockIFilterByTopicPresenter) OnListReceived(arg0 []aggregates.Essay) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnListReceived", arg0)
}

// OnListReceived indicates an expected call of OnListReceived
func (mr *MockIFilterByTopicPresenterMockRecorder) OnListReceived(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnListReceived", reflect.TypeOf((*MockIFilterByTopicPresenter)(nil).OnListReceived), arg0)
}

// OnError mocks base method
func (m *MockIFilterByTopicPresenter) OnError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnError", arg0)
}

// OnError indicates an expected call of OnError
func (mr *MockIFilterByTopicPresenterMockRecorder) OnError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockIFilterByTopicPresenter)(nil).OnError), arg0)
}

// MockIFilterByTopicUseCase is a mock of IFilterByTopicUseCase interface
type MockIFilterByTopicUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIFilterByTopicUseCaseMockRecorder
}

// MockIFilterByTopicUseCaseMockRecorder is the mock recorder for MockIFilterByTopicUseCase
type MockIFilterByTopicUseCaseMockRecorder struct {
	mock *MockIFilterByTopicUseCase
}

// NewMockIFilterByTopicUseCase creates a new mock instance
func NewMockIFilterByTopicUseCase(ctrl *gomock.Controller) *MockIFilterByTopicUseCase {
	mock := &MockIFilterByTopicUseCase{ctrl: ctrl}
	mock.recorder = &MockIFilterByTopicUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIFilterByTopicUseCase) EXPECT() *MockIFilterByTopicUseCaseMockRecorder {
	return m.recorder
}

// FilterByTopic mocks base method
func (m *MockIFilterByTopicUseCase) FilterByTopic(arg0 aggregates.Id) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FilterByTopic", arg0)
}

// FilterByTopic indicates an expected call of FilterByTopic
func (mr *MockIFilterByTopicUseCaseMockRecorder) FilterByTopic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByTopic", reflect.TypeOf((*MockIFilterByTopicUseCase)(nil).FilterByTopic), arg0)
}

// MockICommentPresenter is a mock of ICommentPresenter interface
type MockICommentPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockICommentPresenterMockRecorder
}

// MockICommentPresenterMockRecorder is the mock recorder for MockICommentPresenter
type MockICommentPresenterMockRecorder struct {
	mock *MockICommentPresenter
}

// NewMockICommentPresenter creates a new mock instance
func NewMockICommentPresenter(ctrl *gomock.Controller) *MockICommentPresenter {
	mock := &MockICommentPresenter{ctrl: ctrl}
	mock.recorder = &MockICommentPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommentPresenter) EXPECT() *MockICommentPresenterMockRecorder {
	return m.recorder
}

// OnSuccess mocks base method
func (m *MockICommentPresenter) OnSuccess() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnSuccess")
}

// OnSuccess indicates an expected call of OnSuccess
func (mr *MockICommentPresenterMockRecorder) OnSuccess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSuccess", reflect.TypeOf((*MockICommentPresenter)(nil).OnSuccess))
}

// OnError mocks base method
func (m *MockICommentPresenter) OnError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnError", arg0)
}

// OnError indicates an expected call of OnError
func (mr *MockICommentPresenterMockRecorder) OnError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockICommentPresenter)(nil).OnError), arg0)
}

// MockICommentUseCase is a mock of ICommentUseCase interface
type MockICommentUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockICommentUseCaseMockRecorder
}

// MockICommentUseCaseMockRecorder is the mock recorder for MockICommentUseCase
type MockICommentUseCaseMockRecorder struct {
	mock *MockICommentUseCase
}

// NewMockICommentUseCase creates a new mock instance
func NewMockICommentUseCase(ctrl *gomock.Controller) *MockICommentUseCase {
	mock := &MockICommentUseCase{ctrl: ctrl}
	mock.recorder = &MockICommentUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommentUseCase) EXPECT() *MockICommentUseCaseMockRecorder {
	return m.recorder
}

// Comment mocks base method
func (m *MockICommentUseCase) Comment(essayId aggregates.Id, comment aggregates.Comment) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Comment", essayId, comment)
}

// Comment indicates an expected call of Comment
func (mr *MockICommentUseCaseMockRecorder) Comment(essayId, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Comment", reflect.TypeOf((*MockICommentUseCase)(nil).Comment), essayId, comment)
}
