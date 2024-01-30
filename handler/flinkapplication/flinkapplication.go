// Code generated by Aiven. DO NOT EDIT.

package flinkapplication

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// ServiceFlinkCreateApplicationVersion create a Flink ApplicationVersion
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/version
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkCreateApplicationVersion
	ServiceFlinkCreateApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkCreateApplicationVersionIn) (*ServiceFlinkCreateApplicationVersionOut, error)

	// ServiceFlinkDeleteApplicationVersion delete a Flink ApplicationVersion
	// DELETE /project/{project}/service/{service_name}/flink/application/{application_id}/version/{application_version_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkDeleteApplicationVersion
	ServiceFlinkDeleteApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, applicationVersionId string) (*ServiceFlinkDeleteApplicationVersionOut, error)

	// ServiceFlinkGetApplicationVersion get a Flink ApplicationVersion
	// GET /project/{project}/service/{service_name}/flink/application/{application_id}/version/{application_version_id}
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkGetApplicationVersion
	ServiceFlinkGetApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, applicationVersionId string) (*ServiceFlinkGetApplicationVersionOut, error)

	// ServiceFlinkValidateApplicationVersion validate a Flink ApplicationVersion
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/version/validate
	// https://api.aiven.io/doc/#tag/Service:_Flink/operation/ServiceFlinkValidateApplicationVersion
	ServiceFlinkValidateApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkValidateApplicationVersionIn) (*ServiceFlinkValidateApplicationVersionOut, error)
}

func NewHandler(doer doer) FlinkApplicationHandler {
	return FlinkApplicationHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type FlinkApplicationHandler struct {
	doer doer
}

func (h *FlinkApplicationHandler) ServiceFlinkCreateApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkCreateApplicationVersionIn) (*ServiceFlinkCreateApplicationVersionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/version", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkCreateApplicationVersion", "POST", path, in)
	out := new(ServiceFlinkCreateApplicationVersionOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkDeleteApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, applicationVersionId string) (*ServiceFlinkDeleteApplicationVersionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/version/%s", project, serviceName, applicationId, applicationVersionId)
	b, err := h.doer.Do(ctx, "ServiceFlinkDeleteApplicationVersion", "DELETE", path, nil)
	out := new(ServiceFlinkDeleteApplicationVersionOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkGetApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, applicationVersionId string) (*ServiceFlinkGetApplicationVersionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/version/%s", project, serviceName, applicationId, applicationVersionId)
	b, err := h.doer.Do(ctx, "ServiceFlinkGetApplicationVersion", "GET", path, nil)
	out := new(ServiceFlinkGetApplicationVersionOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *FlinkApplicationHandler) ServiceFlinkValidateApplicationVersion(ctx context.Context, project string, serviceName string, applicationId string, in *ServiceFlinkValidateApplicationVersionIn) (*ServiceFlinkValidateApplicationVersionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/flink/application/%s/version/validate", project, serviceName, applicationId)
	b, err := h.doer.Do(ctx, "ServiceFlinkValidateApplicationVersion", "POST", path, in)
	out := new(ServiceFlinkValidateApplicationVersionOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ColumnOut struct {
	DataType  string `json:"data_type"`
	Extras    string `json:"extras,omitempty"`
	Key       string `json:"key,omitempty"`
	Name      string `json:"name"`
	Nullable  bool   `json:"nullable"`
	Watermark string `json:"watermark,omitempty"`
}
type PositionOut struct {
	CharacterNumber    int `json:"character_number"`
	EndCharacterNumber int `json:"end_character_number"`
	EndLineNumber      int `json:"end_line_number"`
	LineNumber         int `json:"line_number"`
}
type ServiceFlinkCreateApplicationVersionIn struct {
	Sinks     []Sink `json:"sinks"`
	Sources   []Sink `json:"sources"`
	Statement string `json:"statement"`
}
type ServiceFlinkCreateApplicationVersionOut struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Id        string    `json:"id"`
	Sinks     []SinkOut `json:"sinks"`
	Sources   []SinkOut `json:"sources"`
	Statement string    `json:"statement"`
	Version   int       `json:"version"`
}
type ServiceFlinkDeleteApplicationVersionOut struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Id        string    `json:"id"`
	Sinks     []SinkOut `json:"sinks"`
	Sources   []SinkOut `json:"sources"`
	Statement string    `json:"statement"`
	Version   int       `json:"version"`
}
type ServiceFlinkGetApplicationVersionOut struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Id        string    `json:"id"`
	Sinks     []SinkOut `json:"sinks"`
	Sources   []SinkOut `json:"sources"`
	Statement string    `json:"statement"`
	Version   int       `json:"version"`
}
type ServiceFlinkValidateApplicationVersionIn struct {
	Sinks     []Sink `json:"sinks"`
	Sources   []Sink `json:"sources"`
	Statement string `json:"statement,omitempty"`
}
type ServiceFlinkValidateApplicationVersionOut struct {
	Sinks          []SinkOutItem      `json:"sinks"`
	Sources        []SinkOutItem      `json:"sources"`
	Statement      string             `json:"statement,omitempty"`
	StatementError *StatementErrorOut `json:"statement_error,omitempty"`
}
type Sink struct {
	CreateTable   string `json:"create_table"`
	IntegrationId string `json:"integration_id,omitempty"`
}
type SinkOut struct {
	Columns       []ColumnOut    `json:"columns"`
	CreateTable   string         `json:"create_table"`
	IntegrationId string         `json:"integration_id,omitempty"`
	Options       map[string]any `json:"options"`
	TableId       string         `json:"table_id"`
	TableName     string         `json:"table_name"`
}
type SinkOutItem struct {
	Columns       []ColumnOut    `json:"columns,omitempty"`
	CreateTable   string         `json:"create_table"`
	IntegrationId string         `json:"integration_id,omitempty"`
	Message       string         `json:"message,omitempty"`
	Options       map[string]any `json:"options,omitempty"`
	Position      *PositionOut   `json:"position,omitempty"`
	TableName     string         `json:"table_name,omitempty"`
}
type StatementErrorOut struct {
	Message  string       `json:"message"`
	Position *PositionOut `json:"position,omitempty"`
}
