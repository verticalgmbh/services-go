package controllers

import (
	"encoding/json"
	"net/http"
)

// HealthStatus - health status of service
type HealthStatus int8

const (
	// HealthOK - service is working like expected
	HealthOK HealthStatus = iota
	// HealthWarning - service might have some issues but is responding
	HealthWarning
	// HealthError - service does not respond
	HealthError
)

// ServiceResult - result of health check
type ServiceResult struct {
	// Service - name of service for which health check is executed
	Service string `json:"service"`
	// Message - message describing health status
	Message string `json:"message"`
	// Status - status value which represents result of health check
	Status HealthStatus `json:"status"`
}

// HealthResponse - response of a health check
type HealthResponse struct {
	// Services - detailed service checks
	Services []*ServiceResult `json:"services"`
}

// HealthController - controller executing health check for a rest service
type HealthController struct {
	checks []func() *ServiceResult
}

// NewHealthController - creates a new health controller
func NewHealthController(checks ...func() *ServiceResult) *HealthController {
	return &HealthController{checks: checks}
}

// Check checks health of system and returns status
//
// **Parameters**
//   - w: writer used to write response
//   - r: http request
func (controller *HealthController) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var checks []*ServiceResult
	haserror := false

	for _, check := range controller.checks {
		result := check()
		if result.Status == HealthError {
			haserror = true
		}
		checks = append(checks, result)
	}

	if haserror {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	enc := json.NewEncoder(w)
	enc.Encode(&HealthResponse{
		Services: checks,
	})
}
