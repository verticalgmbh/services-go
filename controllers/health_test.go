package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func goodServiceCheck() *ServiceResult {
	return &ServiceResult{
		Service: "TestService",
		Message: "Service is running fine",
		Status:  HealthOK,
	}
}

func problematicServiceCheck() *ServiceResult {
	return &ServiceResult{
		Service: "TestService",
		Message: "Service has some slight issues",
		Status:  HealthWarning,
	}
}

func badServiceCheck() *ServiceResult {
	return &ServiceResult{
		Service: "TestService",
		Message: "Service has crashed",
		Status:  HealthError,
	}
}

func TestHealthCheckOK(t *testing.T) {
	controller := NewHealthController(goodServiceCheck)

	router := mux.NewRouter()
	router.HandleFunc("/api/health", controller.Check)

	req, err := http.NewRequest("GET", "/api/health", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, `{"services":[{"service":"TestService","message":"Service is running fine","status":0}]}`+"\n", rr.Body.String())
}

func TestHealthCheckProblematic(t *testing.T) {
	controller := NewHealthController(problematicServiceCheck)

	router := mux.NewRouter()
	router.HandleFunc("/api/health", controller.Check)

	req, err := http.NewRequest("GET", "/api/health", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, `{"services":[{"service":"TestService","message":"Service has some slight issues","status":1}]}`+"\n", rr.Body.String())
}

func TestHealthCheckBad(t *testing.T) {
	controller := NewHealthController(badServiceCheck)

	router := mux.NewRouter()
	router.HandleFunc("/api/health", controller.Check)

	req, err := http.NewRequest("GET", "/api/health", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusServiceUnavailable, rr.Code)
	require.Equal(t, `{"services":[{"service":"TestService","message":"Service has crashed","status":2}]}`+"\n", rr.Body.String())
}
