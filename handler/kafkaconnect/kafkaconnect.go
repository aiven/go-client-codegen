// Code generated by Aiven. DO NOT EDIT.

package kafkaconnect

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// ServiceKafkaConnectCreateConnector create a Kafka Connect connector
	// POST /v1/project/{project}/service/{service_name}/connectors
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectCreateConnector
	ServiceKafkaConnectCreateConnector(ctx context.Context, project string, serviceName string, in *ServiceKafkaConnectCreateConnectorIn) (*ServiceKafkaConnectCreateConnectorOut, error)

	// ServiceKafkaConnectDeleteConnector delete Kafka Connect connector
	// DELETE /v1/project/{project}/service/{service_name}/connectors/{connector_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectDeleteConnector
	ServiceKafkaConnectDeleteConnector(ctx context.Context, project string, serviceName string, connectorName string) error

	// ServiceKafkaConnectEditConnector edit Kafka Connect connector
	// PUT /v1/project/{project}/service/{service_name}/connectors/{connector_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectEditConnector
	ServiceKafkaConnectEditConnector(ctx context.Context, project string, serviceName string, connectorName string, in *ServiceKafkaConnectEditConnectorIn) (*ServiceKafkaConnectEditConnectorOut, error)

	// ServiceKafkaConnectGetAvailableConnectors get available Kafka Connect connectors
	// GET /v1/project/{project}/service/{service_name}/available-connectors
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectGetAvailableConnectors
	ServiceKafkaConnectGetAvailableConnectors(ctx context.Context, project string, serviceName string) ([]ServiceKafkaConnectGetAvailableConnectorsOut, error)

	// ServiceKafkaConnectGetConnectorConfiguration get Kafka Connect connector configuration schema
	// GET /v1/project/{project}/service/{service_name}/connector-plugins/{connector_name}/configuration
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectGetConnectorConfiguration
	ServiceKafkaConnectGetConnectorConfiguration(ctx context.Context, project string, serviceName string, connectorName string) ([]ConfigurationSchemaOut, error)

	// ServiceKafkaConnectGetConnectorStatus get a Kafka Connect Connector status
	// GET /v1/project/{project}/service/{service_name}/connectors/{connector_name}/status
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectGetConnectorStatus
	ServiceKafkaConnectGetConnectorStatus(ctx context.Context, project string, serviceName string, connectorName string) (*ServiceKafkaConnectGetConnectorStatusOut, error)

	// ServiceKafkaConnectList lists Kafka connectors
	// GET /v1/project/{project}/service/{service_name}/connectors
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectList
	ServiceKafkaConnectList(ctx context.Context, project string, serviceName string) ([]ConnectorOut, error)

	// ServiceKafkaConnectPauseConnector pause a Kafka Connect Connector
	// POST /v1/project/{project}/service/{service_name}/connectors/{connector_name}/pause
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectPauseConnector
	ServiceKafkaConnectPauseConnector(ctx context.Context, project string, serviceName string, connectorName string) error

	// ServiceKafkaConnectRestartConnector restart a Kafka Connect Connector
	// POST /v1/project/{project}/service/{service_name}/connectors/{connector_name}/restart
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectRestartConnector
	ServiceKafkaConnectRestartConnector(ctx context.Context, project string, serviceName string, connectorName string) error

	// ServiceKafkaConnectRestartConnectorTask restart a Kafka Connect Connector task
	// POST /v1/project/{project}/service/{service_name}/connectors/{connector_name}/tasks/{task_id}/restart
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectRestartConnectorTask
	ServiceKafkaConnectRestartConnectorTask(ctx context.Context, project string, serviceName string, connectorName string, taskId string) error

	// ServiceKafkaConnectResumeConnector resume a Kafka Connect Connector
	// POST /v1/project/{project}/service/{service_name}/connectors/{connector_name}/resume
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaConnectResumeConnector
	ServiceKafkaConnectResumeConnector(ctx context.Context, project string, serviceName string, connectorName string) error
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) KafkaConnectHandler {
	return KafkaConnectHandler{doer}
}

type KafkaConnectHandler struct {
	doer doer
}

func (h *KafkaConnectHandler) ServiceKafkaConnectCreateConnector(ctx context.Context, project string, serviceName string, in *ServiceKafkaConnectCreateConnectorIn) (*ServiceKafkaConnectCreateConnectorOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaConnectCreateConnector", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaConnectCreateConnectorOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Connector, nil
}
func (h *KafkaConnectHandler) ServiceKafkaConnectDeleteConnector(ctx context.Context, project string, serviceName string, connectorName string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(connectorName))
	_, err := h.doer.Do(ctx, "ServiceKafkaConnectDeleteConnector", "DELETE", path, nil)
	return err
}
func (h *KafkaConnectHandler) ServiceKafkaConnectEditConnector(ctx context.Context, project string, serviceName string, connectorName string, in *ServiceKafkaConnectEditConnectorIn) (*ServiceKafkaConnectEditConnectorOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(connectorName))
	b, err := h.doer.Do(ctx, "ServiceKafkaConnectEditConnector", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaConnectEditConnectorOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Connector, nil
}
func (h *KafkaConnectHandler) ServiceKafkaConnectGetAvailableConnectors(ctx context.Context, project string, serviceName string) ([]ServiceKafkaConnectGetAvailableConnectorsOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/available-connectors", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaConnectGetAvailableConnectors", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaConnectGetAvailableConnectorsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Plugins, nil
}
func (h *KafkaConnectHandler) ServiceKafkaConnectGetConnectorConfiguration(ctx context.Context, project string, serviceName string, connectorName string) ([]ConfigurationSchemaOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connector-plugins/%s/configuration", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(connectorName))
	b, err := h.doer.Do(ctx, "ServiceKafkaConnectGetConnectorConfiguration", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaConnectGetConnectorConfigurationOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ConfigurationSchema, nil
}
func (h *KafkaConnectHandler) ServiceKafkaConnectGetConnectorStatus(ctx context.Context, project string, serviceName string, connectorName string) (*ServiceKafkaConnectGetConnectorStatusOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors/%s/status", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(connectorName))
	b, err := h.doer.Do(ctx, "ServiceKafkaConnectGetConnectorStatus", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaConnectGetConnectorStatusOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Status, nil
}
func (h *KafkaConnectHandler) ServiceKafkaConnectList(ctx context.Context, project string, serviceName string) ([]ConnectorOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaConnectList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaConnectListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Connectors, nil
}
func (h *KafkaConnectHandler) ServiceKafkaConnectPauseConnector(ctx context.Context, project string, serviceName string, connectorName string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors/%s/pause", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(connectorName))
	_, err := h.doer.Do(ctx, "ServiceKafkaConnectPauseConnector", "POST", path, nil)
	return err
}
func (h *KafkaConnectHandler) ServiceKafkaConnectRestartConnector(ctx context.Context, project string, serviceName string, connectorName string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors/%s/restart", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(connectorName))
	_, err := h.doer.Do(ctx, "ServiceKafkaConnectRestartConnector", "POST", path, nil)
	return err
}
func (h *KafkaConnectHandler) ServiceKafkaConnectRestartConnectorTask(ctx context.Context, project string, serviceName string, connectorName string, taskId string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors/%s/tasks/%s/restart", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(connectorName), url.PathEscape(taskId))
	_, err := h.doer.Do(ctx, "ServiceKafkaConnectRestartConnectorTask", "POST", path, nil)
	return err
}
func (h *KafkaConnectHandler) ServiceKafkaConnectResumeConnector(ctx context.Context, project string, serviceName string, connectorName string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connectors/%s/resume", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(connectorName))
	_, err := h.doer.Do(ctx, "ServiceKafkaConnectResumeConnector", "POST", path, nil)
	return err
}

type AvailableVersionOut struct {
	Deprecated       *bool   `json:"deprecated,omitempty"`        // If the version is deprecated.
	NeedsMaintenance *bool   `json:"needs_maintenance,omitempty"` // Using this version requires a maintenance update.
	Version          *string `json:"version,omitempty"`           // Connector version number
}

// ConfigOut Connector configuration parameters
type ConfigOut struct {
	ConnectorClass string `json:"connector.class"` // The Java class for the connector
	Name           string `json:"name"`            // Unique name for the connector
}
type ConfigurationSchemaOut struct {
	DefaultValue  string                  `json:"default_value"` // Default value to be set if field omitted in configuration
	DisplayName   string                  `json:"display_name"`  // Human friendly name of the field
	Documentation string                  `json:"documentation"` // Assisting help text
	Group         string                  `json:"group"`         // Name of the field group to which the field belongs to
	Importance    ImportanceType          `json:"importance"`    // How important is the field
	Name          string                  `json:"name"`          // Machine friendly name of the field
	Order         int                     `json:"order"`         // Position of the field in the configuration form
	Required      bool                    `json:"required"`      // Defines if the field value is mandatory or not
	Type          ConfigurationSchemaType `json:"type"`          // Configuration value type
	Width         WidthType               `json:"width"`         // Expected length of the input value
}
type ConfigurationSchemaType string

const (
	ConfigurationSchemaTypeString   ConfigurationSchemaType = "STRING"
	ConfigurationSchemaTypeInt      ConfigurationSchemaType = "INT"
	ConfigurationSchemaTypeShort    ConfigurationSchemaType = "SHORT"
	ConfigurationSchemaTypeLong     ConfigurationSchemaType = "LONG"
	ConfigurationSchemaTypeDouble   ConfigurationSchemaType = "DOUBLE"
	ConfigurationSchemaTypeBoolean  ConfigurationSchemaType = "BOOLEAN"
	ConfigurationSchemaTypeList     ConfigurationSchemaType = "LIST"
	ConfigurationSchemaTypeClass    ConfigurationSchemaType = "CLASS"
	ConfigurationSchemaTypePassword ConfigurationSchemaType = "PASSWORD"
)

func ConfigurationSchemaTypeChoices() []string {
	return []string{"STRING", "INT", "SHORT", "LONG", "DOUBLE", "BOOLEAN", "LIST", "CLASS", "PASSWORD"}
}

type ConnectorOut struct {
	Config ConfigOut `json:"config"` // Connector configuration parameters
	Name   string    `json:"name"`   // Connector name
	Plugin PluginOut `json:"plugin"` // Kafka Connector plugin information
	Tasks  []TaskOut `json:"tasks"`  // List of tasks of a connector
}
type ImportanceType string

const (
	ImportanceTypeLow    ImportanceType = "LOW"
	ImportanceTypeMedium ImportanceType = "MEDIUM"
	ImportanceTypeHigh   ImportanceType = "HIGH"
)

func ImportanceTypeChoices() []string {
	return []string{"LOW", "MEDIUM", "HIGH"}
}

// PluginOut Kafka Connector plugin information
type PluginOut struct {
	Author            string                `json:"author"`                       // Connector author name
	AvailableVersions []AvailableVersionOut `json:"available_versions,omitempty"` // Versions available on the service
	Class             string                `json:"class"`                        // Connector class name
	DocUrl            string                `json:"docURL"`                       // Connector documentation URL
	PluginName        *string               `json:"plugin_name,omitempty"`        // Connector plugin name
	Preview           *bool                 `json:"preview,omitempty"`            // Describes if connector is in beta
	PreviewInfo       *string               `json:"preview_info,omitempty"`       // Information about beta stage of connector
	Title             string                `json:"title"`                        // Descriptive human readable name defined by Aiven
	Type              PluginType            `json:"type"`                         // Describes whether data flows from or to Kafka
	Version           string                `json:"version"`                      // Connector version number
}
type PluginType string

const (
	PluginTypeSink    PluginType = "sink"
	PluginTypeSource  PluginType = "source"
	PluginTypeUnknown PluginType = "unknown"
)

func PluginTypeChoices() []string {
	return []string{"sink", "source", "unknown"}
}

type ServiceKafkaConnectConnectorStateType string

const (
	ServiceKafkaConnectConnectorStateTypeFailed     ServiceKafkaConnectConnectorStateType = "FAILED"
	ServiceKafkaConnectConnectorStateTypePaused     ServiceKafkaConnectConnectorStateType = "PAUSED"
	ServiceKafkaConnectConnectorStateTypeRunning    ServiceKafkaConnectConnectorStateType = "RUNNING"
	ServiceKafkaConnectConnectorStateTypeUnassigned ServiceKafkaConnectConnectorStateType = "UNASSIGNED"
)

func ServiceKafkaConnectConnectorStateTypeChoices() []string {
	return []string{"FAILED", "PAUSED", "RUNNING", "UNASSIGNED"}
}

// ServiceKafkaConnectCreateConnectorIn ServiceKafkaConnectCreateConnectorRequestBody
type ServiceKafkaConnectCreateConnectorIn struct {
	ConnectorClass string `json:"connector.class"` // The Java class for the connector
	Name           string `json:"name"`            // Unique name for the connector
}

// ServiceKafkaConnectCreateConnectorOut Kafka connector information
type ServiceKafkaConnectCreateConnectorOut struct {
	Config ConfigOut `json:"config"` // Connector configuration parameters
	Name   string    `json:"name"`   // Connector name
	Plugin PluginOut `json:"plugin"` // Kafka Connector plugin information
	Tasks  []TaskOut `json:"tasks"`  // List of tasks of a connector
}

// ServiceKafkaConnectEditConnectorIn ServiceKafkaConnectEditConnectorRequestBody
type ServiceKafkaConnectEditConnectorIn struct {
	ConnectorClass string `json:"connector.class"` // The Java class for the connector
	Name           string `json:"name"`            // Unique name for the connector
}

// ServiceKafkaConnectEditConnectorOut Kafka connector information
type ServiceKafkaConnectEditConnectorOut struct {
	Config ConfigOut `json:"config"` // Connector configuration parameters
	Name   string    `json:"name"`   // Connector name
	Plugin PluginOut `json:"plugin"` // Kafka Connector plugin information
	Tasks  []TaskOut `json:"tasks"`  // List of tasks of a connector
}
type ServiceKafkaConnectGetAvailableConnectorsOut struct {
	Author            string                `json:"author"`                       // Connector author name
	AvailableVersions []AvailableVersionOut `json:"available_versions,omitempty"` // Versions available on the service
	Class             string                `json:"class"`                        // Connector class name
	DocUrl            string                `json:"docURL"`                       // Connector documentation URL
	PluginName        *string               `json:"plugin_name,omitempty"`        // Connector plugin name
	Preview           *bool                 `json:"preview,omitempty"`            // Describes if connector is in beta
	PreviewInfo       *string               `json:"preview_info,omitempty"`       // Information about beta stage of connector
	Title             string                `json:"title"`                        // Descriptive human readable name defined by Aiven
	Type              PluginType            `json:"type"`                         // Describes whether data flows from or to Kafka
	Version           string                `json:"version"`                      // Connector version number
}

// ServiceKafkaConnectGetConnectorStatusOut Connector status information
type ServiceKafkaConnectGetConnectorStatusOut struct {
	State ServiceKafkaConnectConnectorStateType          `json:"state"` // Current status of the connector
	Tasks []ServiceKafkaConnectGetConnectorStatusTaskOut `json:"tasks"` // List of tasks currently running for the connector
}
type ServiceKafkaConnectGetConnectorStatusTaskOut struct {
	Id    int           `json:"id"`    // Task identifier
	State TaskStateType `json:"state"` // Current status of the task
	Trace string        `json:"trace"` // Task error information
}
type TaskOut struct {
	Connector string `json:"connector"` // Related connector name
	Task      int    `json:"task"`      // Task id / number
}
type TaskStateType string

const (
	TaskStateTypeFailed     TaskStateType = "FAILED"
	TaskStateTypePaused     TaskStateType = "PAUSED"
	TaskStateTypeRunning    TaskStateType = "RUNNING"
	TaskStateTypeUnassigned TaskStateType = "UNASSIGNED"
)

func TaskStateTypeChoices() []string {
	return []string{"FAILED", "PAUSED", "RUNNING", "UNASSIGNED"}
}

type WidthType string

const (
	WidthTypeNone   WidthType = "NONE"
	WidthTypeShort  WidthType = "SHORT"
	WidthTypeMedium WidthType = "MEDIUM"
	WidthTypeLong   WidthType = "LONG"
)

func WidthTypeChoices() []string {
	return []string{"NONE", "SHORT", "MEDIUM", "LONG"}
}

// serviceKafkaConnectCreateConnectorOut ServiceKafkaConnectCreateConnectorResponse
type serviceKafkaConnectCreateConnectorOut struct {
	Connector ServiceKafkaConnectCreateConnectorOut `json:"connector"` // Kafka connector information
}

// serviceKafkaConnectEditConnectorOut ServiceKafkaConnectEditConnectorResponse
type serviceKafkaConnectEditConnectorOut struct {
	Connector ServiceKafkaConnectEditConnectorOut `json:"connector"` // Kafka connector information
}

// serviceKafkaConnectGetAvailableConnectorsOut ServiceKafkaConnectGetAvailableConnectorsResponse
type serviceKafkaConnectGetAvailableConnectorsOut struct {
	Plugins []ServiceKafkaConnectGetAvailableConnectorsOut `json:"plugins"` // List of available Kafka Connect connector plugins
}

// serviceKafkaConnectGetConnectorConfigurationOut ServiceKafkaConnectGetConnectorConfigurationResponse
type serviceKafkaConnectGetConnectorConfigurationOut struct {
	ConfigurationSchema []ConfigurationSchemaOut `json:"configuration_schema"` // List of connector configuration field definitions
}

// serviceKafkaConnectGetConnectorStatusOut ServiceKafkaConnectGetConnectorStatusResponse
type serviceKafkaConnectGetConnectorStatusOut struct {
	Status ServiceKafkaConnectGetConnectorStatusOut `json:"status"` // Connector status information
}

// serviceKafkaConnectListOut ServiceKafkaConnectListResponse
type serviceKafkaConnectListOut struct {
	Connectors []ConnectorOut `json:"connectors"` // List of active Kafka Connect connectors
}
