package pgxpoolmock

import (
	"context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Ensure that the pgx.Tx interface is implemented by MockPgxTx
var _ pgx.Tx = (*MockPgxTx)(nil)

// MockPgxTx is a mock of pgx.Tx interface.
type MockPgxTx struct {
	ctrl     *gomock.Controller
	recorder *MockPgxTxMockRecorder
}

// MockPgxTxMockRecorder is the mock recorder for MockPgxTx.
type MockPgxTxMockRecorder struct {
	mock *MockPgxTx
}

// NewMockPgxTx creates a new mock instance.
func NewMockPgxTx(ctrl *gomock.Controller) *MockPgxTx {
	mock := &MockPgxTx{ctrl: ctrl}
	mock.recorder = &MockPgxTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPgxTx) EXPECT() *MockPgxTxMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockPgxTx) Begin(arg0 context.Context) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockPgxTxMockRecorder) Begin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockPgxTx)(nil).Begin), ctx)
}

// Commit mocks base method.
func (m *MockPgxTx) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockPgxTxMockRecorder) Commit(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockPgxTx)(nil).Commit), ctx)
}

// Conn mocks base method.
func (m *MockPgxTx) Conn() *pgx.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn")
	ret0, _ := ret[0].(*pgx.Conn)
	return ret0
}

// Conn indicates an expected call of Conn.
func (mr *MockPgxTxMockRecorder) Conn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockPgxTx)(nil).Conn))
}

// CopyFrom mocks base method.
func (m *MockPgxTx) CopyFrom(arg0 context.Context, arg1 pgx.Identifier, arg2 []string, arg3 pgx.CopyFromSource) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFrom", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopyFrom indicates an expected call of CopyFrom.
func (mr *MockPgxTxMockRecorder) CopyFrom(ctx, tableName, columnNames, rowSrc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFrom", reflect.TypeOf((*MockPgxTx)(nil).CopyFrom), ctx, tableName, columnNames, rowSrc)
}

// Exec mocks base method.
func (m *MockPgxTx) Exec(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1, arg2)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockPgxTxMockRecorder) Exec(ctx, sql, arguments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockPgxTx)(nil).Exec), ctx, sql, arguments)
}

// LargeObjects mocks base method.
func (m *MockPgxTx) LargeObjects() pgx.LargeObjects {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LargeObjects")
	ret0, _ := ret[0].(pgx.LargeObjects)
	return ret0
}

// LargeObjects indicates an expected call of LargeObjects.
func (mr *MockPgxTxMockRecorder) LargeObjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LargeObjects", reflect.TypeOf((*MockPgxTx)(nil).LargeObjects))
}

// Prepare mocks base method.
func (m *MockPgxTx) Prepare(arg0 context.Context, arg1 string, arg2 string) (*pgconn.StatementDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", arg0, arg1, arg2)
	ret0, _ := ret[0].(*pgconn.StatementDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare.
func (mr *MockPgxTxMockRecorder) Prepare(ctx, name, sql interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockPgxTx)(nil).Prepare), ctx, name, sql)
}

// Query mocks base method.
func (m *MockPgxTx) Query(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1, arg2)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockPgxTxMockRecorder) Query(ctx, sql, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockPgxTx)(nil).Query), ctx, sql, args)
}

// QueryRow mocks base method.
func (m *MockPgxTx) QueryRow(arg0 context.Context, arg1 string, arg2 ...interface{}) pgx.Row {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRow", arg0, arg1, arg2)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockPgxTxMockRecorder) QueryRow(ctx, sql, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockPgxTx)(nil).QueryRow), ctx, sql, args)
}

// Rollback mocks base method.
func (m *MockPgxTx) Rollback(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockPgxTxMockRecorder) Rollback(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockPgxTx)(nil).Rollback), ctx)
}

// SendBatch mocks base method.
func (m *MockPgxTx) SendBatch(arg0 context.Context, arg1 *pgx.Batch) pgx.BatchResults {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendBatch", arg0, arg1)
	ret0, _ := ret[0].(pgx.BatchResults)
	return ret0
}

// SendBatch indicates an expected call of SendBatch.
func (mr *MockPgxTxMockRecorder) SendBatch(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatch", reflect.TypeOf((*MockPgxTx)(nil).SendBatch), ctx, b)
}
