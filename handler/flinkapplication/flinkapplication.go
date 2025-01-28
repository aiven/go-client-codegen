// Code generated by Aiven. DO NOT EDIT.

package flinkapplication

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// ServiceFlinkCreateApplication create a Flink Application
	// POST /v1/project/{project}/service/{service_name}/flink/application
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCreateApplication
	ServiceFlinkCreateApplication(ctx context.Context, project string, serviceName string, in *ServiceFlinkCreateApplicationIn) (*ServiceFlinkCreateApplicationOut, error)

	// ServiceFlinkDeleteApplication delete a Flink Application
	// DELETE /v1/project/{project}/service/{service_name}/flink/application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkDeleteApplication
	ServiceFlinkDeleteApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkDeleteApplicationOut, error)

	// ServiceFlinkGetApplication get a Flink Application
	// GET /v1/project/{project}/service/{service_name}/flink/application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkGetApplication
	ServiceFlinkGetApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkGetApplicationOut, error)

	// ServiceFlinkListApplications get all Flink Applications
	// GET /v1/project/{project}/service/{service_name}/flink/application
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkListApplications
	ServiceFlinkListApplications(ctx context.Context, project string, serviceName string) ([]ApplicationOut, error)

	// ServiceFlinkUpdateApplication update a Flink Application
	// PUT /v1/project/{project}/service/{service_name}/flink/application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkUpdateApplication
	ServiceFlinkUpdateApplication(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkUpdateApplicationIn) (*ServiceFlinkUpdateApplicationOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) FlinkApplicationHandler {
	return FlinkApplicationHandler{doer}
}

type FlinkApplicationHandler struct {
	doer doer
}

func (h *FlinkApplicationHandler) ServiceFlinkCreateApplication(ctx context.Context, project string, serviceName string, in *ServiceFlinkCreateApplicationIn) (*ServiceFlinkCreateApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceFlinkCreateApplication", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkCreateApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkDeleteApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkDeleteApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId))
	b, err := h.doer.Do(ctx, "ServiceFlinkDeleteApplication", "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkDeleteApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkGetApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkGetApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId))
	b, err := h.doer.Do(ctx, "ServiceFlinkGetApplication", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkGetApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkListApplications(ctx context.Context, project string, serviceName string) ([]ApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceFlinkListApplications", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceFlinkListApplicationsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Applications, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkUpdateApplication(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkUpdateApplicationIn) (*ServiceFlinkUpdateApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/application/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId))
	b, err := h.doer.Do(ctx, "ServiceFlinkUpdateApplication", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkUpdateApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ApplicationOut struct {
	CreatedAt *time.Time `json:"created_at,omitempty"` // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy *string    `json:"created_by,omitempty"` // The creator of this entity
	Id        string     `json:"id"`                   // Application ID
	Name      string     `json:"name"`                 // Application name
	UpdatedAt *time.Time `json:"updated_at,omitempty"` // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy *string    `json:"updated_by,omitempty"` // The latest updater of this entity
}

// ApplicationVersionIn Flink ApplicationVersion
type ApplicationVersionIn struct {
	Sinks     []SinkIn   `json:"sinks"`
	Sources   []SourceIn `json:"sources"`
	Statement string     `json:"statement"` // The INSERT INTO SQL statement that will be performed as a job
}
type ApplicationVersionOut struct {
	CreatedAt time.Time   `json:"created_at"` // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy string      `json:"created_by"` // The creator of this entity
	Id        string      `json:"id"`         // ApplicationVersion ID
	Sinks     []SinkOut   `json:"sinks"`
	Sources   []SourceOut `json:"sources"`
	Statement string      `json:"statement"` // The INSERT INTO SQL statement that will be performed as a job
	Version   int         `json:"version"`   // Version number
}
type ColumnOut struct {
	DataType  string  `json:"data_type"`           // The data type of the column
	Extras    *string `json:"extras,omitempty"`    // Column extra information
	Key       *string `json:"key,omitempty"`       // The key info of the column
	Name      string  `json:"name"`                // The name of the column
	Nullable  bool    `json:"nullable"`            // Whether the column is nullable, i.e. if true, the column is NOT NULL
	Watermark *string `json:"watermark,omitempty"` // Information of the watermark if the column is used for watermark.
}

// CurrentDeploymentOut Flink ApplicationDeployment
type CurrentDeploymentOut struct {
	CreatedAt         time.Time                   `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy         string                      `json:"created_by"`                   // The creator of this entity
	ErrorMsg          *string                     `json:"error_msg,omitempty"`          // Error message describing what caused deployment to fail
	Id                string                      `json:"id"`                           // Deployment ID
	JobId             *string                     `json:"job_id,omitempty"`             // Job ID
	LastSavepoint     *string                     `json:"last_savepoint,omitempty"`     // Job savepoint
	Parallelism       int                         `json:"parallelism"`                  // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	RestartEnabled    bool                        `json:"restart_enabled"`              // Specifies whether a Flink Job is restarted in case it fails
	StartingSavepoint *string                     `json:"starting_savepoint,omitempty"` // Job savepoint
	Status            CurrentDeploymentStatusType `json:"status"`                       // Deployment status
	VersionId         string                      `json:"version_id"`                   // ApplicationVersion ID
}
type CurrentDeploymentStatusType string

const (
	CurrentDeploymentStatusTypeCanceled               CurrentDeploymentStatusType = "CANCELED"
	CurrentDeploymentStatusTypeCancelling             CurrentDeploymentStatusType = "CANCELLING"
	CurrentDeploymentStatusTypeCancellingRequested    CurrentDeploymentStatusType = "CANCELLING_REQUESTED"
	CurrentDeploymentStatusTypeCreated                CurrentDeploymentStatusType = "CREATED"
	CurrentDeploymentStatusTypeDeleteRequested        CurrentDeploymentStatusType = "DELETE_REQUESTED"
	CurrentDeploymentStatusTypeDeleting               CurrentDeploymentStatusType = "DELETING"
	CurrentDeploymentStatusTypeFailed                 CurrentDeploymentStatusType = "FAILED"
	CurrentDeploymentStatusTypeFailing                CurrentDeploymentStatusType = "FAILING"
	CurrentDeploymentStatusTypeFinished               CurrentDeploymentStatusType = "FINISHED"
	CurrentDeploymentStatusTypeInitializing           CurrentDeploymentStatusType = "INITIALIZING"
	CurrentDeploymentStatusTypeReconciling            CurrentDeploymentStatusType = "RECONCILING"
	CurrentDeploymentStatusTypeRestarting             CurrentDeploymentStatusType = "RESTARTING"
	CurrentDeploymentStatusTypeRunning                CurrentDeploymentStatusType = "RUNNING"
	CurrentDeploymentStatusTypeSaving                 CurrentDeploymentStatusType = "SAVING"
	CurrentDeploymentStatusTypeSavingAndStop          CurrentDeploymentStatusType = "SAVING_AND_STOP"
	CurrentDeploymentStatusTypeSavingAndStopRequested CurrentDeploymentStatusType = "SAVING_AND_STOP_REQUESTED"
	CurrentDeploymentStatusTypeSuspended              CurrentDeploymentStatusType = "SUSPENDED"
)

func CurrentDeploymentStatusTypeChoices() []string {
	return []string{"CANCELED", "CANCELLING", "CANCELLING_REQUESTED", "CREATED", "DELETE_REQUESTED", "DELETING", "FAILED", "FAILING", "FINISHED", "INITIALIZING", "RECONCILING", "RESTARTING", "RUNNING", "SAVING", "SAVING_AND_STOP", "SAVING_AND_STOP_REQUESTED", "SUSPENDED"}
}

// ServiceFlinkCreateApplicationIn ServiceFlinkCreateApplicationRequestBody
type ServiceFlinkCreateApplicationIn struct {
	ApplicationVersion *ApplicationVersionIn `json:"application_version,omitempty"` // Flink ApplicationVersion
	Name               string                `json:"name"`                          // Application name
}

// ServiceFlinkCreateApplicationOut ServiceFlinkCreateApplicationResponse
type ServiceFlinkCreateApplicationOut struct {
	ApplicationVersions []ApplicationVersionOut `json:"application_versions"`
	CreatedAt           time.Time               `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy           string                  `json:"created_by"`                   // The creator of this entity
	CurrentDeployment   *CurrentDeploymentOut   `json:"current_deployment,omitempty"` // Flink ApplicationDeployment
	Id                  string                  `json:"id"`                           // Application ID
	Name                string                  `json:"name"`                         // Application name
	UpdatedAt           time.Time               `json:"updated_at"`                   // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy           string                  `json:"updated_by"`                   // The latest updater of this entity
}

// ServiceFlinkDeleteApplicationOut ServiceFlinkDeleteApplicationResponse
type ServiceFlinkDeleteApplicationOut struct {
	ApplicationVersions []ApplicationVersionOut `json:"application_versions"`
	CreatedAt           time.Time               `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy           string                  `json:"created_by"`                   // The creator of this entity
	CurrentDeployment   *CurrentDeploymentOut   `json:"current_deployment,omitempty"` // Flink ApplicationDeployment
	Id                  string                  `json:"id"`                           // Application ID
	Name                string                  `json:"name"`                         // Application name
	UpdatedAt           time.Time               `json:"updated_at"`                   // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy           string                  `json:"updated_by"`                   // The latest updater of this entity
}

// ServiceFlinkGetApplicationOut ServiceFlinkGetApplicationResponse
type ServiceFlinkGetApplicationOut struct {
	ApplicationVersions []ApplicationVersionOut `json:"application_versions"`
	CreatedAt           time.Time               `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy           string                  `json:"created_by"`                   // The creator of this entity
	CurrentDeployment   *CurrentDeploymentOut   `json:"current_deployment,omitempty"` // Flink ApplicationDeployment
	Id                  string                  `json:"id"`                           // Application ID
	Name                string                  `json:"name"`                         // Application name
	UpdatedAt           time.Time               `json:"updated_at"`                   // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy           string                  `json:"updated_by"`                   // The latest updater of this entity
}

// ServiceFlinkUpdateApplicationIn ServiceFlinkUpdateApplicationRequestBody
type ServiceFlinkUpdateApplicationIn struct {
	Name string `json:"name"` // Application name
}

// ServiceFlinkUpdateApplicationOut ServiceFlinkUpdateApplicationResponse
type ServiceFlinkUpdateApplicationOut struct {
	ApplicationVersions []ApplicationVersionOut `json:"application_versions"`
	CreatedAt           time.Time               `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy           string                  `json:"created_by"`                   // The creator of this entity
	CurrentDeployment   *CurrentDeploymentOut   `json:"current_deployment,omitempty"` // Flink ApplicationDeployment
	Id                  string                  `json:"id"`                           // Application ID
	Name                string                  `json:"name"`                         // Application name
	UpdatedAt           time.Time               `json:"updated_at"`                   // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy           string                  `json:"updated_by"`                   // The latest updater of this entity
}
type SinkIn struct {
	CreateTable   string  `json:"create_table"`             // The CREATE TABLE statement
	IntegrationId *string `json:"integration_id,omitempty"` // Integration ID
}
type SinkOut struct {
	Columns       []ColumnOut    `json:"columns"`
	CreateTable   string         `json:"create_table"`             // The CREATE TABLE statement
	IntegrationId *string        `json:"integration_id,omitempty"` // Integration ID
	Options       map[string]any `json:"options"`                  // Option
	TableId       string         `json:"table_id"`                 // Sink ID
	TableName     string         `json:"table_name"`               // Table name
}
type SourceIn struct {
	CreateTable   string  `json:"create_table"`             // The CREATE TABLE statement
	IntegrationId *string `json:"integration_id,omitempty"` // Integration ID
}
type SourceOut struct {
	Columns       []ColumnOut    `json:"columns"`
	CreateTable   string         `json:"create_table"`             // The CREATE TABLE statement
	IntegrationId *string        `json:"integration_id,omitempty"` // Integration ID
	Options       map[string]any `json:"options"`                  // Option
	TableId       string         `json:"table_id"`                 // Source ID
	TableName     string         `json:"table_name"`               // Table name
}

// serviceFlinkListApplicationsOut ServiceFlinkListApplicationsResponse
type serviceFlinkListApplicationsOut struct {
	Applications []ApplicationOut `json:"applications"` // Flink Applications
}
