// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package booking_mock is a generated GoMock package.
package booking_mock

import (
	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// AddBooking mocks base method
func (m *MockUsecase) AddBooking(booking bookingModel.BookingAdd) (bookingModel.BookingID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBooking", booking)
	ret0, _ := ret[0].(bookingModel.BookingID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBooking indicates an expected call of AddBooking
func (mr *MockUsecaseMockRecorder) AddBooking(booking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBooking", reflect.TypeOf((*MockUsecase)(nil).AddBooking), booking)
}

// DeleteBooking mocks base method
func (m *MockUsecase) DeleteBooking(bookingID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBooking", bookingID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBooking indicates an expected call of DeleteBooking
func (mr *MockUsecaseMockRecorder) DeleteBooking(bookingID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBooking", reflect.TypeOf((*MockUsecase)(nil).DeleteBooking), bookingID)
}

// GetBookings mocks base method
func (m *MockUsecase) GetBookings(roomID int) ([]bookingModel.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookings", roomID)
	ret0, _ := ret[0].([]bookingModel.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookings indicates an expected call of GetBookings
func (mr *MockUsecaseMockRecorder) GetBookings(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookings", reflect.TypeOf((*MockUsecase)(nil).GetBookings), roomID)
}

// CheckBookingExist mocks base method
func (m *MockUsecase) CheckBookingExist(bookingID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBookingExist", bookingID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBookingExist indicates an expected call of CheckBookingExist
func (mr *MockUsecaseMockRecorder) CheckBookingExist(bookingID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBookingExist", reflect.TypeOf((*MockUsecase)(nil).CheckBookingExist), bookingID)
}
