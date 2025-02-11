// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockreceipt

import (
	http "net/http"

	httprouter "github.com/julienschmidt/httprouter"
	mock "github.com/stretchr/testify/mock"
)

// ReceiptStore is an autogenerated mock type for the ReceiptStore type
type ReceiptStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ReceiptStore) Close() {
	_m.Called()
}

// GetReceipt provides a mock function with given fields: res, req, params
func (_m *ReceiptStore) GetReceipt(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	_m.Called(res, req, params)
}

// GetReceipts provides a mock function with given fields: res, req, params
func (_m *ReceiptStore) GetReceipts(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	_m.Called(res, req, params)
}

// Init provides a mock function with given fields:
func (_m *ReceiptStore) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessReceipt provides a mock function with given fields: msgBytes
func (_m *ReceiptStore) ProcessReceipt(msgBytes []byte) {
	_m.Called(msgBytes)
}

// ValidateConf provides a mock function with given fields:
func (_m *ReceiptStore) ValidateConf() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
