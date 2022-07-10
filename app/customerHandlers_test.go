package app

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/iloncar89/banking-lib/errs"
	"github.com/iloncar89/banking/dto"
	"github.com/iloncar89/banking/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Ivan", "Osijek", "31000", "1989-11-21", "1"},
		{"1002", "Josip", "Amsterdam", "10000", "1992-12-01", "1"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed with testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("database error"))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed with testing the status code")
	}
}
