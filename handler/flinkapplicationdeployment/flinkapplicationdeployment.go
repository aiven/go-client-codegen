// Code generated by Aiven. DO NOT EDIT.

package flinkapplicationdeployment

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// ServiceFlinkCancelApplicationDeployment cancel an ApplicationDeployment
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}/cancel
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCancelApplicationDeployment
	ServiceFlinkCancelApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkCancelApplicationDeploymentOut, error)

	// ServiceFlinkCreateApplicationDeployment create an ApplicationDeployment
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/deployment
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCreateApplicationDeployment
	ServiceFlinkCreateApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkCreateApplicationDeploymentIn) (*ServiceFlinkCreateApplicationDeploymentOut, error)

	// ServiceFlinkDeleteApplicationDeployment delete an ApplicationDeployment
	// DELETE /project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkDeleteApplicationDeployment
	ServiceFlinkDeleteApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkDeleteApplicationDeploymentOut, error)

	// ServiceFlinkGetApplicationDeployment get an ApplicationDeployment
	// GET /project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkGetApplicationDeployment
	ServiceFlinkGetApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkGetApplicationDeploymentOut, error)

	// ServiceFlinkListApplicationDeployments get all ApplicationDeployments
	// GET /project/{project}/service/{service_name}/flink/application/{application_id}/deployment
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkListApplicationDeployments
	ServiceFlinkListApplicationDeployments(ctx context.Context, project string, serviceName string, applicationId string) ([]DeploymentOut, error)

	// ServiceFlinkStopApplicationDeployment stop an ApplicationDeployment
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/deployment/{deployment_id}/stop
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkStopApplicationDeployment
	ServiceFlinkStopApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkStopApplicationDeploymentOut, error)
}

func NewHandler(doer doer) FlinkApplicationDeploymentHandler {
	return FlinkApplicationDeploymentHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type FlinkApplicationDeploymentHandler struct {
	doer doer
}

func (h *FlinkApplicationDeploymentHandler) ServiceFlinkCancelApplicationDeployment(ctx context.Context, project string, serviceName string, applicationId string, deploymentId string) (*ServiceFlinkCancelApplicationDeploymentOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment/%s/cancel", project, serviceName, applicationId, deploymentId)
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
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment", project, serviceName, applicationId)
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
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment/%s", project, serviceName, applicationId, deploymentId)
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
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment/%s", project, serviceName, applicationId, deploymentId)
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
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment", project, serviceName, applicationId)
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
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/deployment/%s/stop", project, serviceName, applicationId, deploymentId)
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
	CreatedAt         string `json:"created_at"`
	CreatedBy         string `json:"created_by"`
	ErrorMsg          string `json:"error_msg,omitempty"`
	Id                string `json:"id"`
	JobId             string `json:"job_id,omitempty"`
	LastSavepoint     string `json:"last_savepoint,omitempty"`
	Parallelism       int    `json:"parallelism"`
	RestartEnabled    bool   `json:"restart_enabled"`
	StartingSavepoint string `json:"starting_savepoint,omitempty"`
	Status            string `json:"status"`
	VersionId         string `json:"version_id"`
}
type ServiceFlinkCancelApplicationDeploymentOut struct {
	CreatedAt         string `json:"created_at"`
	CreatedBy         string `json:"created_by"`
	ErrorMsg          string `json:"error_msg,omitempty"`
	Id                string `json:"id"`
	JobId             string `json:"job_id,omitempty"`
	LastSavepoint     string `json:"last_savepoint,omitempty"`
	Parallelism       int    `json:"parallelism"`
	RestartEnabled    bool   `json:"restart_enabled"`
	StartingSavepoint string `json:"starting_savepoint,omitempty"`
	Status            string `json:"status"`
	VersionId         string `json:"version_id"`
}
type ServiceFlinkCreateApplicationDeploymentIn struct {
	Parallelism       *int   `json:"parallelism,omitempty"`
	RestartEnabled    *bool  `json:"restart_enabled,omitempty"`
	StartingSavepoint string `json:"starting_savepoint,omitempty"`
	VersionId         string `json:"version_id"`
}
type ServiceFlinkCreateApplicationDeploymentOut struct {
	CreatedAt         string `json:"created_at"`
	CreatedBy         string `json:"created_by"`
	ErrorMsg          string `json:"error_msg,omitempty"`
	Id                string `json:"id"`
	JobId             string `json:"job_id,omitempty"`
	LastSavepoint     string `json:"last_savepoint,omitempty"`
	Parallelism       int    `json:"parallelism"`
	RestartEnabled    bool   `json:"restart_enabled"`
	StartingSavepoint string `json:"starting_savepoint,omitempty"`
	Status            string `json:"status"`
	VersionId         string `json:"version_id"`
}
type ServiceFlinkDeleteApplicationDeploymentOut struct {
	CreatedAt         string `json:"created_at"`
	CreatedBy         string `json:"created_by"`
	ErrorMsg          string `json:"error_msg,omitempty"`
	Id                string `json:"id"`
	JobId             string `json:"job_id,omitempty"`
	LastSavepoint     string `json:"last_savepoint,omitempty"`
	Parallelism       int    `json:"parallelism"`
	RestartEnabled    bool   `json:"restart_enabled"`
	StartingSavepoint string `json:"starting_savepoint,omitempty"`
	Status            string `json:"status"`
	VersionId         string `json:"version_id"`
}
type ServiceFlinkGetApplicationDeploymentOut struct {
	CreatedAt         string `json:"created_at"`
	CreatedBy         string `json:"created_by"`
	ErrorMsg          string `json:"error_msg,omitempty"`
	Id                string `json:"id"`
	JobId             string `json:"job_id,omitempty"`
	LastSavepoint     string `json:"last_savepoint,omitempty"`
	Parallelism       int    `json:"parallelism"`
	RestartEnabled    bool   `json:"restart_enabled"`
	StartingSavepoint string `json:"starting_savepoint,omitempty"`
	Status            string `json:"status"`
	VersionId         string `json:"version_id"`
}
type ServiceFlinkStopApplicationDeploymentOut struct {
	CreatedAt         string `json:"created_at"`
	CreatedBy         string `json:"created_by"`
	ErrorMsg          string `json:"error_msg,omitempty"`
	Id                string `json:"id"`
	JobId             string `json:"job_id,omitempty"`
	LastSavepoint     string `json:"last_savepoint,omitempty"`
	Parallelism       int    `json:"parallelism"`
	RestartEnabled    bool   `json:"restart_enabled"`
	StartingSavepoint string `json:"starting_savepoint,omitempty"`
	Status            string `json:"status"`
	VersionId         string `json:"version_id"`
}
type serviceFlinkListApplicationDeploymentsOut struct {
	Deployments []DeploymentOut `json:"deployments"`
}
