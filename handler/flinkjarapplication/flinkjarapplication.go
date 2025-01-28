// Code generated by Aiven. DO NOT EDIT.

package flinkjarapplication

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// ServiceFlinkCreateJarApplication [EXPERIMENTAL] Create a Flink JarApplication
	// POST /v1/project/{project}/service/{service_name}/flink/jar_application
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCreateJarApplication
	ServiceFlinkCreateJarApplication(ctx context.Context, project string, serviceName string, in *ServiceFlinkCreateJarApplicationIn) (*ServiceFlinkCreateJarApplicationOut, error)

	// ServiceFlinkDeleteJarApplication [EXPERIMENTAL] Delete a Flink JarApplication
	// DELETE /v1/project/{project}/service/{service_name}/flink/jar_application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkDeleteJarApplication
	ServiceFlinkDeleteJarApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkDeleteJarApplicationOut, error)

	// ServiceFlinkGetJarApplication [EXPERIMENTAL] Get a Flink JarApplication
	// GET /v1/project/{project}/service/{service_name}/flink/jar_application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkGetJarApplication
	ServiceFlinkGetJarApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkGetJarApplicationOut, error)

	// ServiceFlinkListJarApplications [EXPERIMENTAL] Get all Flink JarApplications
	// GET /v1/project/{project}/service/{service_name}/flink/jar_application
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkListJarApplications
	ServiceFlinkListJarApplications(ctx context.Context, project string, serviceName string) ([]ApplicationOut, error)

	// ServiceFlinkUpdateJarApplication [EXPERIMENTAL] Update a Flink JarApplication
	// PUT /v1/project/{project}/service/{service_name}/flink/jar_application/{application_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkUpdateJarApplication
	ServiceFlinkUpdateJarApplication(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkUpdateJarApplicationIn) (*ServiceFlinkUpdateJarApplicationOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) FlinkJarApplicationHandler {
	return FlinkJarApplicationHandler{doer}
}

type FlinkJarApplicationHandler struct {
	doer doer
}

func (h *FlinkJarApplicationHandler) ServiceFlinkCreateJarApplication(ctx context.Context, project string, serviceName string, in *ServiceFlinkCreateJarApplicationIn) (*ServiceFlinkCreateJarApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/jar_application", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceFlinkCreateJarApplication", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkCreateJarApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkJarApplicationHandler) ServiceFlinkDeleteJarApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkDeleteJarApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/jar_application/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId))
	b, err := h.doer.Do(ctx, "ServiceFlinkDeleteJarApplication", "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkDeleteJarApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkJarApplicationHandler) ServiceFlinkGetJarApplication(ctx context.Context, project string, serviceName string, applicationId string) (*ServiceFlinkGetJarApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/jar_application/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId))
	b, err := h.doer.Do(ctx, "ServiceFlinkGetJarApplication", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkGetJarApplicationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkJarApplicationHandler) ServiceFlinkListJarApplications(ctx context.Context, project string, serviceName string) ([]ApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/jar_application", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceFlinkListJarApplications", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceFlinkListJarApplicationsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Applications, nil
}
func (h *FlinkJarApplicationHandler) ServiceFlinkUpdateJarApplication(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkUpdateJarApplicationIn) (*ServiceFlinkUpdateJarApplicationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/flink/jar_application/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(applicationId))
	b, err := h.doer.Do(ctx, "ServiceFlinkUpdateJarApplication", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceFlinkUpdateJarApplicationOut)
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
type ApplicationVersionOut struct {
	CreatedAt time.Time    `json:"created_at"`          // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy string       `json:"created_by"`          // The creator of this entity
	FileInfo  *FileInfoOut `json:"file_info,omitempty"` // Flink JarApplicationVersion FileInfo
	Id        string       `json:"id"`                  // ApplicationVersion ID
	Version   int          `json:"version"`             // Version number
}

// CurrentDeploymentOut Flink JarApplicationDeployment
type CurrentDeploymentOut struct {
	CreatedAt         time.Time                   `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy         string                      `json:"created_by"`                   // The creator of this entity
	EntryClass        *string                     `json:"entry_class,omitempty"`        // The fully qualified name of the entry class to pass during Flink job submission through the entryClass parameter
	ErrorMsg          *string                     `json:"error_msg,omitempty"`          // Error message describing what caused deployment to fail
	Id                string                      `json:"id"`                           // Deployment ID
	JobId             *string                     `json:"job_id,omitempty"`             // Job ID
	LastSavepoint     *string                     `json:"last_savepoint,omitempty"`     // Job savepoint
	Parallelism       int                         `json:"parallelism"`                  // Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.
	ProgramArgs       []string                    `json:"program_args,omitempty"`       // Arguments to pass during Flink job submission through the programArgsList parameter
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

// FileInfoOut Flink JarApplicationVersion FileInfo
type FileInfoOut struct {
	FileSha256         *string             `json:"file_sha256,omitempty"`          // sha256 of the file if known
	FileSize           *int                `json:"file_size,omitempty"`            // The size of the file in bytes
	FileStatus         FileStatusType      `json:"file_status,omitempty"`          // Indicates whether the uploaded .jar file has been verified by the system and deployment ready
	Url                *string             `json:"url,omitempty"`                  // The pre-signed url of the bucket where the .jar file is uploaded. Becomes null when the JarApplicationVersion is ready or failed.
	VerifyErrorCode    VerifyErrorCodeType `json:"verify_error_code,omitempty"`    // In the case file_status is FAILED, the error code of the failure.
	VerifyErrorMessage *string             `json:"verify_error_message,omitempty"` // In the case file_status is FAILED, may contain details about the failure.
}
type FileStatusType string

const (
	FileStatusTypeFailed  FileStatusType = "FAILED"
	FileStatusTypeInitial FileStatusType = "INITIAL"
	FileStatusTypeReady   FileStatusType = "READY"
)

func FileStatusTypeChoices() []string {
	return []string{"FAILED", "INITIAL", "READY"}
}

// ServiceFlinkCreateJarApplicationIn ServiceFlinkCreateJarApplicationRequestBody
type ServiceFlinkCreateJarApplicationIn struct {
	Name string `json:"name"` // Application name
}

// ServiceFlinkCreateJarApplicationOut ServiceFlinkCreateJarApplicationResponse
type ServiceFlinkCreateJarApplicationOut struct {
	ApplicationVersions []ApplicationVersionOut `json:"application_versions"`         // JarApplicationVersions
	CreatedAt           time.Time               `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy           string                  `json:"created_by"`                   // The creator of this entity
	CurrentDeployment   *CurrentDeploymentOut   `json:"current_deployment,omitempty"` // Flink JarApplicationDeployment
	Id                  string                  `json:"id"`                           // Application ID
	Name                string                  `json:"name"`                         // Application name
	UpdatedAt           time.Time               `json:"updated_at"`                   // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy           string                  `json:"updated_by"`                   // The latest updater of this entity
}

// ServiceFlinkDeleteJarApplicationOut ServiceFlinkDeleteJarApplicationResponse
type ServiceFlinkDeleteJarApplicationOut struct {
	ApplicationVersions []ApplicationVersionOut `json:"application_versions"`         // JarApplicationVersions
	CreatedAt           time.Time               `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy           string                  `json:"created_by"`                   // The creator of this entity
	CurrentDeployment   *CurrentDeploymentOut   `json:"current_deployment,omitempty"` // Flink JarApplicationDeployment
	Id                  string                  `json:"id"`                           // Application ID
	Name                string                  `json:"name"`                         // Application name
	UpdatedAt           time.Time               `json:"updated_at"`                   // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy           string                  `json:"updated_by"`                   // The latest updater of this entity
}

// ServiceFlinkGetJarApplicationOut ServiceFlinkGetJarApplicationResponse
type ServiceFlinkGetJarApplicationOut struct {
	ApplicationVersions []ApplicationVersionOut `json:"application_versions"`         // JarApplicationVersions
	CreatedAt           time.Time               `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy           string                  `json:"created_by"`                   // The creator of this entity
	CurrentDeployment   *CurrentDeploymentOut   `json:"current_deployment,omitempty"` // Flink JarApplicationDeployment
	Id                  string                  `json:"id"`                           // Application ID
	Name                string                  `json:"name"`                         // Application name
	UpdatedAt           time.Time               `json:"updated_at"`                   // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy           string                  `json:"updated_by"`                   // The latest updater of this entity
}

// ServiceFlinkUpdateJarApplicationIn ServiceFlinkUpdateJarApplicationRequestBody
type ServiceFlinkUpdateJarApplicationIn struct {
	Name string `json:"name"` // Application name
}

// ServiceFlinkUpdateJarApplicationOut ServiceFlinkUpdateJarApplicationResponse
type ServiceFlinkUpdateJarApplicationOut struct {
	ApplicationVersions []ApplicationVersionOut `json:"application_versions"`         // JarApplicationVersions
	CreatedAt           time.Time               `json:"created_at"`                   // The creation timestamp of this entity in ISO 8601 format, always in UTC
	CreatedBy           string                  `json:"created_by"`                   // The creator of this entity
	CurrentDeployment   *CurrentDeploymentOut   `json:"current_deployment,omitempty"` // Flink JarApplicationDeployment
	Id                  string                  `json:"id"`                           // Application ID
	Name                string                  `json:"name"`                         // Application name
	UpdatedAt           time.Time               `json:"updated_at"`                   // The update timestamp of this entity in ISO 8601 format, always in UTC
	UpdatedBy           string                  `json:"updated_by"`                   // The latest updater of this entity
}
type VerifyErrorCodeType int

const (
	VerifyErrorCodeType1 VerifyErrorCodeType = 1
	VerifyErrorCodeType2 VerifyErrorCodeType = 2
	VerifyErrorCodeType3 VerifyErrorCodeType = 3
)

func VerifyErrorCodeTypeChoices() []int {
	return []int{1, 2, 3}
}

// serviceFlinkListJarApplicationsOut ServiceFlinkListJarApplicationsResponse
type serviceFlinkListJarApplicationsOut struct {
	Applications []ApplicationOut `json:"applications"` // Flink JarApplications
}
