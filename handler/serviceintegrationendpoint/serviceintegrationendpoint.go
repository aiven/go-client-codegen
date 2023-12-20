// Code generated by Aiven. DO NOT EDIT.

package serviceintegrationendpoint

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// Create create a new service integration endpoint
	// ServiceIntegrationEndpointCreate POST /project/{project}/integration_endpoint
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointCreate
	Create(ctx context.Context, project string, in *CreateIn) (*ServiceIntegrationEndpoint, error)

	// Delete delete a service integration endpoint
	// ServiceIntegrationEndpointDelete DELETE /project/{project}/integration_endpoint/{integration_endpoint_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointDelete
	Delete(ctx context.Context, project string, integrationEndpointId string) error

	// Get get service integration endpoint
	// ServiceIntegrationEndpointGet GET /project/{project}/integration_endpoint/{integration_endpoint_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointGet
	Get(ctx context.Context, project string, integrationEndpointId string) (*ServiceIntegrationEndpoint, error)

	// List list available integration endpoints for project
	// ServiceIntegrationEndpointList GET /project/{project}/integration_endpoint
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointList
	List(ctx context.Context, project string) ([]ServiceIntegrationEndpoint, error)

	// Types list available service integration endpoint types
	// ServiceIntegrationEndpointTypes GET /project/{project}/integration_endpoint_types
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointTypes
	Types(ctx context.Context, project string) ([]EndpointTypeItem, error)

	// Update update service integration endpoint
	// ServiceIntegrationEndpointUpdate PUT /project/{project}/integration_endpoint/{integration_endpoint_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointUpdate
	Update(ctx context.Context, project string, integrationEndpointId string, in *UpdateIn) (*ServiceIntegrationEndpoint, error)
}

func NewHandler(doer doer) Handler {
	return &handler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type handler struct {
	doer doer
}

func (h *handler) Create(ctx context.Context, project string, in *CreateIn) (*ServiceIntegrationEndpoint, error) {
	path := fmt.Sprintf("/project/%s/integration_endpoint", project)
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointCreate", "POST", path, in)
	out := new(createOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ServiceIntegrationEndpoint, nil
}
func (h *handler) Delete(ctx context.Context, project string, integrationEndpointId string) error {
	path := fmt.Sprintf("/project/%s/integration_endpoint/%s", project, integrationEndpointId)
	_, err := h.doer.Do(ctx, "ServiceIntegrationEndpointDelete", "DELETE", path, nil)
	return err
}
func (h *handler) Get(ctx context.Context, project string, integrationEndpointId string) (*ServiceIntegrationEndpoint, error) {
	path := fmt.Sprintf("/project/%s/integration_endpoint/%s", project, integrationEndpointId)
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointGet", "GET", path, nil)
	out := new(getOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ServiceIntegrationEndpoint, nil
}
func (h *handler) List(ctx context.Context, project string) ([]ServiceIntegrationEndpoint, error) {
	path := fmt.Sprintf("/project/%s/integration_endpoint", project)
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointList", "GET", path, nil)
	out := new(listOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ServiceIntegrationEndpoints, nil
}
func (h *handler) Types(ctx context.Context, project string) ([]EndpointTypeItem, error) {
	path := fmt.Sprintf("/project/%s/integration_endpoint_types", project)
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointTypes", "GET", path, nil)
	out := new(typesOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.EndpointTypes, nil
}
func (h *handler) Update(ctx context.Context, project string, integrationEndpointId string, in *UpdateIn) (*ServiceIntegrationEndpoint, error) {
	path := fmt.Sprintf("/project/%s/integration_endpoint/%s", project, integrationEndpointId)
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointUpdate", "PUT", path, in)
	out := new(updateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ServiceIntegrationEndpoint, nil
}

type CreateIn struct {
	EndpointName string         `json:"endpoint_name"`
	EndpointType EndpointType   `json:"endpoint_type"`
	UserConfig   map[string]any `json:"user_config"`
}
type createOut struct {
	ServiceIntegrationEndpoint *ServiceIntegrationEndpoint `json:"service_integration_endpoint"`
}
type EndpointType string

const (
	EndpointTypeAutoscaler                   EndpointType = "autoscaler"
	EndpointTypeDatadog                      EndpointType = "datadog"
	EndpointTypeExternalAwsCloudwatchLogs    EndpointType = "external_aws_cloudwatch_logs"
	EndpointTypeExternalAwsCloudwatchMetrics EndpointType = "external_aws_cloudwatch_metrics"
	EndpointTypeExternalAwsS3                EndpointType = "external_aws_s3"
	EndpointTypeExternalElasticsearchLogs    EndpointType = "external_elasticsearch_logs"
	EndpointTypeExternalGoogleCloudBigquery  EndpointType = "external_google_cloud_bigquery"
	EndpointTypeExternalGoogleCloudLogging   EndpointType = "external_google_cloud_logging"
	EndpointTypeExternalKafka                EndpointType = "external_kafka"
	EndpointTypeExternalMysql                EndpointType = "external_mysql"
	EndpointTypeExternalOpensearchLogs       EndpointType = "external_opensearch_logs"
	EndpointTypeExternalPostgresql           EndpointType = "external_postgresql"
	EndpointTypeExternalRedis                EndpointType = "external_redis"
	EndpointTypeExternalSchemaRegistry       EndpointType = "external_schema_registry"
	EndpointTypeJolokia                      EndpointType = "jolokia"
	EndpointTypePrometheus                   EndpointType = "prometheus"
	EndpointTypeRsyslog                      EndpointType = "rsyslog"
)

type EndpointTypeItem struct {
	EndpointType     string         `json:"endpoint_type"`
	ServiceTypes     []string       `json:"service_types"`
	Title            string         `json:"title"`
	UserConfigSchema map[string]any `json:"user_config_schema"`
}
type getOut struct {
	ServiceIntegrationEndpoint *ServiceIntegrationEndpoint `json:"service_integration_endpoint"`
}
type listOut struct {
	ServiceIntegrationEndpoints []ServiceIntegrationEndpoint `json:"service_integration_endpoints"`
}
type ServiceIntegrationEndpoint struct {
	EndpointConfig map[string]any `json:"endpoint_config"`
	EndpointId     string         `json:"endpoint_id"`
	EndpointName   string         `json:"endpoint_name"`
	EndpointType   EndpointType   `json:"endpoint_type"`
	UserConfig     map[string]any `json:"user_config"`
}
type typesOut struct {
	EndpointTypes []EndpointTypeItem `json:"endpoint_types"`
}
type UpdateIn struct {
	UserConfig map[string]any `json:"user_config"`
}
type updateOut struct {
	ServiceIntegrationEndpoint *ServiceIntegrationEndpoint `json:"service_integration_endpoint"`
}
