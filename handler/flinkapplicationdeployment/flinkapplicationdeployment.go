// Code generated by Aiven. DO NOT EDIT.

package flinkapplicationdeployment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// ServiceFlinkCancelApplicationDeployment cancel an ApplicationDeployment
	// POST /v1/project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}/cancel
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCancelApplicationDeployment
	ServiceFlinkCancelApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkCancelApplicationDeploymentOut, error)

	// ServiceFlinkCreateApplicationDeployment create an ApplicationDeployment
	// POST /v1/project/{project}/service/{service_name}/flink/application/{application_id}/deployment
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCreateApplicationDeployment
	ServiceFlinkCreateApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkCreateApplicationDeploymentIn) (*ServiceFlinkCreateApplicationDeploymentOut, error)

	// ServiceFlinkDeleteApplicationDeployment delete an ApplicationDeployment
	// DELETE /v1/project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkDeleteApplicationDeployment
	ServiceFlinkDeleteApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkDeleteApplicationDeploymentOut, error)

	// ServiceFlinkGetApplicationDeployment get an ApplicationDeployment
	// GET /v1/project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkGetApplicationDeployment
	ServiceFlinkGetApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkGetApplicationDeploymentOut, error)

	// ServiceFlinkListApplicationDeployments get all ApplicationDeployments
	// GET /v1/project/{project}/service/{service_name}/flink/application/{application_id}/deployment
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkListApplicationDeployments
	ServiceFlinkListApplicationDeployments(ctx context.Context, project string, serviceName string, applicationId string) ([]DeploymentOut, error)

	// ServiceFlinkStopApplicationDeployment stop an ApplicationDeployment
	// POST /v1/project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}/stop
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkStopApplicationDeployment
	ServiceFlinkStopApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkStopApplicationDeploymentOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) FlinkApplicationDeploymentHandler {
	return FlinkApplicationDeploymentHandler{doer}
}

type FlinkApplicationDeploymentHandler struct {
	doer doer
}

func (h *FlinkApplicationDeploymentHandler) ServiceFlinkCancelApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkCancelApplicationDeploymentOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s/deployment/%s/cancel", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId), url.PathEscape(deploymentId))
	b, err := h.doer.Do(ctx, "ServiceFlinkCancelApplicationDeployment", "POST", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkCancelApplicationDeploymentOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationDeploymentHandler) ServiceFlinkCreateApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkCreateApplicationDeploymentIn) (*ServiceFlinkCreateApplicationDeploymentOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s/deployment", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId))
	b, err := h.doer.Do(ctx, "ServiceFlinkCreateApplicationDeployment", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkCreateApplicationDeploymentOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationDeploymentHandler) ServiceFlinkDeleteApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkDeleteApplicationDeploymentOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s/deployment/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId), url.PathEscape(deploymentId))
	b, err := h.doer.Do(ctx, "ServiceFlinkDeleteApplicationDeployment", "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkDeleteApplicationDeploymentOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationDeploymentHandler) ServiceFlinkGetApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkGetApplicationDeploymentOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s/deployment/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId), url.PathEscape(deploymentId))
	b, err := h.doer.Do(ctx, "ServiceFlinkGetApplicationDeployment", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkGetApplicationDeploymentOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationDeploymentHandler) ServiceFlinkListApplicationDeployments(ctx context.Context, project string, serviceName string, applicationId string) ([]DeploymentOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s/deployment", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId))
	b, err := h.doer.Do(ctx, "ServiceFlinkListApplicationDeployments", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceFlinkListApplicationDeploymentsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Deployments, nil
}
func (h *FlinkApplicationDeploymentHandler) ServiceFlinkStopApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkStopApplicationDeploymentOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s/deployment/%s/stop", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId), url.PathEscape(deploymentId))
	b, err := h.doer.Do(ctx, "ServiceFlinkStopApplicationDeployment", "POST", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkStopApplicationDeploymentOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DeploymentOut struct {
	CreatedAt         time.Time            `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy         string               `json:"created_by"`                   // The creator of this entity
	ErrorMsg          *string              `json:"error_msg,omitempty"`          // Error message describing what caused deployment to fail
	Id                string               `json:"id"`                           // Deployment ID
	JobId             *string              `json:"job_id,omitempty"`             // Job ID
	LastSavepoint     *string              `json:"last_savepoint,omitempty"`     // Job savepoint
	Parallelism       int                  `json:"parallelism"`                  // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	RestartEnabled    bool                 `json:"restart_enabled"`              // Specifies whether a Flink Job is restarted in case it fails
	StartingSavepoint *string              `json:"starting_savepoint,omitempty"` // Job savepoint
	Status            DeploymentStatusType `json:"status"`                       // Deployment status
	VersionId         string               `json:"version_id"`                   // ApplicationVersion ID
}
type DeploymentStatusType string

const (
	DeploymentStatusTypeCanceled               DeploymentStatusType = "CANCELED"
	DeploymentStatusTypeCancelling             DeploymentStatusType = "CANCELLING"
	DeploymentStatusTypeCancellingRequested    DeploymentStatusType = "CANCELLING_REQUESTED"
	DeploymentStatusTypeCreated                DeploymentStatusType = "CREATED"
	DeploymentStatusTypeDeleteRequested        DeploymentStatusType = "DELETE_REQUESTED"
	DeploymentStatusTypeDeleting               DeploymentStatusType = "DELETING"
	DeploymentStatusTypeFailed                 DeploymentStatusType = "FAILED"
	DeploymentStatusTypeFailing                DeploymentStatusType = "FAILING"
	DeploymentStatusTypeFinished               DeploymentStatusType = "FINISHED"
	DeploymentStatusTypeInitializing           DeploymentStatusType = "INITIALIZING"
	DeploymentStatusTypeReconciling            DeploymentStatusType = "RECONCILING"
	DeploymentStatusTypeRestarting             DeploymentStatusType = "RESTARTING"
	DeploymentStatusTypeRunning                DeploymentStatusType = "RUNNING"
	DeploymentStatusTypeSaving                 DeploymentStatusType = "SAVING"
	DeploymentStatusTypeSavingAndStop          DeploymentStatusType = "SAVING_AND_STOP"
	DeploymentStatusTypeSavingAndStopRequested DeploymentStatusType = "SAVING_AND_STOP_REQUESTED"
	DeploymentStatusTypeSuspended              DeploymentStatusType = "SUSPENDED"
)

func DeploymentStatusTypeChoices() []string {
	return []string{"CANCELED", "CANCELLING", "CANCELLING_REQUESTED", "CREATED", "DELETE_REQUESTED", "DELETING", "FAILED", "FAILING", "FINISHED", "INITIALIZING", "RECONCILING", "RESTARTING", "RUNNING", "SAVING", "SAVING_AND_STOP", "SAVING_AND_STOP_REQUESTED", "SUSPENDED"}
}

type ServiceFlinkApplicationDeploymentStatusType string

const (
	ServiceFlinkApplicationDeploymentStatusTypeCanceled               ServiceFlinkApplicationDeploymentStatusType = "CANCELED"
	ServiceFlinkApplicationDeploymentStatusTypeCancelling             ServiceFlinkApplicationDeploymentStatusType = "CANCELLING"
	ServiceFlinkApplicationDeploymentStatusTypeCancellingRequested    ServiceFlinkApplicationDeploymentStatusType = "CANCELLING_REQUESTED"
	ServiceFlinkApplicationDeploymentStatusTypeCreated                ServiceFlinkApplicationDeploymentStatusType = "CREATED"
	ServiceFlinkApplicationDeploymentStatusTypeDeleteRequested        ServiceFlinkApplicationDeploymentStatusType = "DELETE_REQUESTED"
	ServiceFlinkApplicationDeploymentStatusTypeDeleting               ServiceFlinkApplicationDeploymentStatusType = "DELETING"
	ServiceFlinkApplicationDeploymentStatusTypeFailed                 ServiceFlinkApplicationDeploymentStatusType = "FAILED"
	ServiceFlinkApplicationDeploymentStatusTypeFailing                ServiceFlinkApplicationDeploymentStatusType = "FAILING"
	ServiceFlinkApplicationDeploymentStatusTypeFinished               ServiceFlinkApplicationDeploymentStatusType = "FINISHED"
	ServiceFlinkApplicationDeploymentStatusTypeInitializing           ServiceFlinkApplicationDeploymentStatusType = "INITIALIZING"
	ServiceFlinkApplicationDeploymentStatusTypeReconciling            ServiceFlinkApplicationDeploymentStatusType = "RECONCILING"
	ServiceFlinkApplicationDeploymentStatusTypeRestarting             ServiceFlinkApplicationDeploymentStatusType = "RESTARTING"
	ServiceFlinkApplicationDeploymentStatusTypeRunning                ServiceFlinkApplicationDeploymentStatusType = "RUNNING"
	ServiceFlinkApplicationDeploymentStatusTypeSaving                 ServiceFlinkApplicationDeploymentStatusType = "SAVING"
	ServiceFlinkApplicationDeploymentStatusTypeSavingAndStop          ServiceFlinkApplicationDeploymentStatusType = "SAVING_AND_STOP"
	ServiceFlinkApplicationDeploymentStatusTypeSavingAndStopRequested ServiceFlinkApplicationDeploymentStatusType = "SAVING_AND_STOP_REQUESTED"
	ServiceFlinkApplicationDeploymentStatusTypeSuspended              ServiceFlinkApplicationDeploymentStatusType = "SUSPENDED"
)

func ServiceFlinkApplicationDeploymentStatusTypeChoices() []string {
	return []string{"CANCELED", "CANCELLING", "CANCELLING_REQUESTED", "CREATED", "DELETE_REQUESTED", "DELETING", "FAILED", "FAILING", "FINISHED", "INITIALIZING", "RECONCILING", "RESTARTING", "RUNNING", "SAVING", "SAVING_AND_STOP", "SAVING_AND_STOP_REQUESTED", "SUSPENDED"}
}

// ServiceFlinkCancelApplicationDeploymentOut ServiceFlinkCancelApplicationDeploymentResponse
type ServiceFlinkCancelApplicationDeploymentOut struct {
	CreatedAt         time.Time                                   `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy         string                                      `json:"created_by"`                   // The creator of this entity
	ErrorMsg          *string                                     `json:"error_msg,omitempty"`          // Error message describing what caused deployment to fail
	Id                string                                      `json:"id"`                           // Deployment ID
	JobId             *string                                     `json:"job_id,omitempty"`             // Job ID
	LastSavepoint     *string                                     `json:"last_savepoint,omitempty"`     // Job savepoint
	Parallelism       int                                         `json:"parallelism"`                  // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	RestartEnabled    bool                                        `json:"restart_enabled"`              // Specifies whether a Flink Job is restarted in case it fails
	StartingSavepoint *string                                     `json:"starting_savepoint,omitempty"` // Job savepoint
	Status            ServiceFlinkApplicationDeploymentStatusType `json:"status"`                       // Deployment status
	VersionId         string                                      `json:"version_id"`                   // ApplicationVersion ID
}

// ServiceFlinkCreateApplicationDeploymentIn ServiceFlinkCreateApplicationDeploymentRequestBody
type ServiceFlinkCreateApplicationDeploymentIn struct {
	Parallelism       *int    `json:"parallelism,omitempty"`        // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	RestartEnabled    *bool   `json:"restart_enabled,omitempty"`    // Specifies whether a Flink Job is restarted in case it fails
	StartingSavepoint *string `json:"starting_savepoint,omitempty"` // Job savepoint
	VersionId         string  `json:"version_id"`                   // ApplicationVersion ID
}

// ServiceFlinkCreateApplicationDeploymentOut ServiceFlinkCreateApplicationDeploymentResponse
type ServiceFlinkCreateApplicationDeploymentOut struct {
	CreatedAt         time.Time                                   `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy         string                                      `json:"created_by"`                   // The creator of this entity
	ErrorMsg          *string                                     `json:"error_msg,omitempty"`          // Error message describing what caused deployment to fail
	Id                string                                      `json:"id"`                           // Deployment ID
	JobId             *string                                     `json:"job_id,omitempty"`             // Job ID
	LastSavepoint     *string                                     `json:"last_savepoint,omitempty"`     // Job savepoint
	Parallelism       int                                         `json:"parallelism"`                  // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	RestartEnabled    bool                                        `json:"restart_enabled"`              // Specifies whether a Flink Job is restarted in case it fails
	StartingSavepoint *string                                     `json:"starting_savepoint,omitempty"` // Job savepoint
	Status            ServiceFlinkApplicationDeploymentStatusType `json:"status"`                       // Deployment status
	VersionId         string                                      `json:"version_id"`                   // ApplicationVersion ID
}

// ServiceFlinkDeleteApplicationDeploymentOut ServiceFlinkDeleteApplicationDeploymentResponse
type ServiceFlinkDeleteApplicationDeploymentOut struct {
	CreatedAt         time.Time                                   `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy         string                                      `json:"created_by"`                   // The creator of this entity
	ErrorMsg          *string                                     `json:"error_msg,omitempty"`          // Error message describing what caused deployment to fail
	Id                string                                      `json:"id"`                           // Deployment ID
	JobId             *string                                     `json:"job_id,omitempty"`             // Job ID
	LastSavepoint     *string                                     `json:"last_savepoint,omitempty"`     // Job savepoint
	Parallelism       int                                         `json:"parallelism"`                  // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	RestartEnabled    bool                                        `json:"restart_enabled"`              // Specifies whether a Flink Job is restarted in case it fails
	StartingSavepoint *string                                     `json:"starting_savepoint,omitempty"` // Job savepoint
	Status            ServiceFlinkApplicationDeploymentStatusType `json:"status"`                       // Deployment status
	VersionId         string                                      `json:"version_id"`                   // ApplicationVersion ID
}

// ServiceFlinkGetApplicationDeploymentOut ServiceFlinkGetApplicationDeploymentResponse
type ServiceFlinkGetApplicationDeploymentOut struct {
	CreatedAt         time.Time                                   `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy         string                                      `json:"created_by"`                   // The creator of this entity
	ErrorMsg          *string                                     `json:"error_msg,omitempty"`          // Error message describing what caused deployment to fail
	Id                string                                      `json:"id"`                           // Deployment ID
	JobId             *string                                     `json:"job_id,omitempty"`             // Job ID
	LastSavepoint     *string                                     `json:"last_savepoint,omitempty"`     // Job savepoint
	Parallelism       int                                         `json:"parallelism"`                  // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	RestartEnabled    bool                                        `json:"restart_enabled"`              // Specifies whether a Flink Job is restarted in case it fails
	StartingSavepoint *string                                     `json:"starting_savepoint,omitempty"` // Job savepoint
	Status            ServiceFlinkApplicationDeploymentStatusType `json:"status"`                       // Deployment status
	VersionId         string                                      `json:"version_id"`                   // ApplicationVersion ID
}

// ServiceFlinkStopApplicationDeploymentOut ServiceFlinkStopApplicationDeploymentResponse
type ServiceFlinkStopApplicationDeploymentOut struct {
	CreatedAt         time.Time                                   `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy         string                                      `json:"created_by"`                   // The creator of this entity
	ErrorMsg          *string                                     `json:"error_msg,omitempty"`          // Error message describing what caused deployment to fail
	Id                string                                      `json:"id"`                           // Deployment ID
	JobId             *string                                     `json:"job_id,omitempty"`             // Job ID
	LastSavepoint     *string                                     `json:"last_savepoint,omitempty"`     // Job savepoint
	Parallelism       int                                         `json:"parallelism"`                  // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	RestartEnabled    bool                                        `json:"restart_enabled"`              // Specifies whether a Flink Job is restarted in case it fails
	StartingSavepoint *string                                     `json:"starting_savepoint,omitempty"` // Job savepoint
	Status            ServiceFlinkApplicationDeploymentStatusType `json:"status"`                       // Deployment status
	VersionId         string                                      `json:"version_id"`                   // ApplicationVersion ID
}

// serviceFlinkListApplicationDeploymentsOut ServiceFlinkListApplicationDeploymentsResponse
type serviceFlinkListApplicationDeploymentsOut struct {
	Deployments []DeploymentOut `json:"deployments"` // Flink ApplicationDeployments
}
