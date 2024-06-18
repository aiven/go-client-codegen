// Code generated by Aiven. DO NOT EDIT.

package serviceintegration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// ServiceIntegrationCreate create a new service integration
	// POST /v1/project/{project}/integration
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationCreate
	ServiceIntegrationCreate(ctx context.Context, project string, in *ServiceIntegrationCreateIn) (*ServiceIntegrationCreateOut, error)

	// ServiceIntegrationDelete delete a service integration
	// DELETE /v1/project/{project}/integration/{integration_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationDelete
	ServiceIntegrationDelete(ctx context.Context, project string, integrationId string) error

	// ServiceIntegrationEndpointCreate create a new service integration endpoint
	// POST /v1/project/{project}/integration_endpoint
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointCreate
	ServiceIntegrationEndpointCreate(ctx context.Context, project string, in *ServiceIntegrationEndpointCreateIn) (*ServiceIntegrationEndpointCreateOut, error)

	// ServiceIntegrationEndpointDelete delete a service integration endpoint
	// DELETE /v1/project/{project}/integration_endpoint/{integration_endpoint_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointDelete
	ServiceIntegrationEndpointDelete(ctx context.Context, project string, integrationEndpointId string) error

	// ServiceIntegrationEndpointGet get service integration endpoint
	// GET /v1/project/{project}/integration_endpoint/{integration_endpoint_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointGet
	ServiceIntegrationEndpointGet(ctx context.Context, project string, integrationEndpointId string) (*ServiceIntegrationEndpointGetOut, error)

	// ServiceIntegrationEndpointList list available integration endpoints for project
	// GET /v1/project/{project}/integration_endpoint
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointList
	ServiceIntegrationEndpointList(ctx context.Context, project string) ([]ServiceIntegrationEndpointOut, error)

	// ServiceIntegrationEndpointTypes list available service integration endpoint types
	// GET /v1/project/{project}/integration_endpoint_types
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointTypes
	ServiceIntegrationEndpointTypes(ctx context.Context, project string) ([]EndpointTypeOut, error)

	// ServiceIntegrationEndpointUpdate update service integration endpoint
	// PUT /v1/project/{project}/integration_endpoint/{integration_endpoint_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationEndpointUpdate
	ServiceIntegrationEndpointUpdate(ctx context.Context, project string, integrationEndpointId string, in *ServiceIntegrationEndpointUpdateIn) (*ServiceIntegrationEndpointUpdateOut, error)

	// ServiceIntegrationGet get service integration
	// GET /v1/project/{project}/integration/{integration_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationGet
	ServiceIntegrationGet(ctx context.Context, project string, integrationId string) (*ServiceIntegrationGetOut, error)

	// ServiceIntegrationList list available integrations for a service
	// GET /v1/project/{project}/service/{service_name}/integration
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationList
	ServiceIntegrationList(ctx context.Context, project string, serviceName string) ([]ServiceIntegrationOut, error)

	// ServiceIntegrationTypes list available service integration types
	// GET /v1/project/{project}/integration_types
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationTypes
	ServiceIntegrationTypes(ctx context.Context, project string) ([]IntegrationTypeOut, error)

	// ServiceIntegrationUpdate update a service integration
	// PUT /v1/project/{project}/integration/{integration_id}
	// https://api.aiven.io/doc/#tag/Service_Integrations/operation/ServiceIntegrationUpdate
	ServiceIntegrationUpdate(ctx context.Context, project string, integrationId string, in *ServiceIntegrationUpdateIn) (*ServiceIntegrationUpdateOut, error)
}

func NewHandler(doer doer) ServiceIntegrationHandler {
	return ServiceIntegrationHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type ServiceIntegrationHandler struct {
	doer doer
}

func (h *ServiceIntegrationHandler) ServiceIntegrationCreate(ctx context.Context, project string, in *ServiceIntegrationCreateIn) (*ServiceIntegrationCreateOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ServiceIntegrationCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.ServiceIntegration, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationDelete(ctx context.Context, project string, integrationId string) error {
	path := fmt.Sprintf("/v1/project/%s/integration/%s", url.PathEscape(project), url.PathEscape(integrationId))
	_, err := h.doer.Do(ctx, "ServiceIntegrationDelete", "DELETE", path, nil)
	return err
}
func (h *ServiceIntegrationHandler) ServiceIntegrationEndpointCreate(ctx context.Context, project string, in *ServiceIntegrationEndpointCreateIn) (*ServiceIntegrationEndpointCreateOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration_endpoint", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationEndpointCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.ServiceIntegrationEndpoint, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationEndpointDelete(ctx context.Context, project string, integrationEndpointId string) error {
	path := fmt.Sprintf("/v1/project/%s/integration_endpoint/%s", url.PathEscape(project), url.PathEscape(integrationEndpointId))
	_, err := h.doer.Do(ctx, "ServiceIntegrationEndpointDelete", "DELETE", path, nil)
	return err
}
func (h *ServiceIntegrationHandler) ServiceIntegrationEndpointGet(ctx context.Context, project string, integrationEndpointId string) (*ServiceIntegrationEndpointGetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration_endpoint/%s", url.PathEscape(project), url.PathEscape(integrationEndpointId))
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationEndpointGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.ServiceIntegrationEndpoint, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationEndpointList(ctx context.Context, project string) ([]ServiceIntegrationEndpointOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration_endpoint", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationEndpointListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ServiceIntegrationEndpoints, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationEndpointTypes(ctx context.Context, project string) ([]EndpointTypeOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration_endpoint_types", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointTypes", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationEndpointTypesOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.EndpointTypes, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationEndpointUpdate(ctx context.Context, project string, integrationEndpointId string, in *ServiceIntegrationEndpointUpdateIn) (*ServiceIntegrationEndpointUpdateOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration_endpoint/%s", url.PathEscape(project), url.PathEscape(integrationEndpointId))
	b, err := h.doer.Do(ctx, "ServiceIntegrationEndpointUpdate", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationEndpointUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.ServiceIntegrationEndpoint, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationGet(ctx context.Context, project string, integrationId string) (*ServiceIntegrationGetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration/%s", url.PathEscape(project), url.PathEscape(integrationId))
	b, err := h.doer.Do(ctx, "ServiceIntegrationGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.ServiceIntegration, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationList(ctx context.Context, project string, serviceName string) ([]ServiceIntegrationOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/integration", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceIntegrationList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ServiceIntegrations, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationTypes(ctx context.Context, project string) ([]IntegrationTypeOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration_types", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ServiceIntegrationTypes", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationTypesOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.IntegrationTypes, nil
}
func (h *ServiceIntegrationHandler) ServiceIntegrationUpdate(ctx context.Context, project string, integrationId string, in *ServiceIntegrationUpdateIn) (*ServiceIntegrationUpdateOut, error) {
	path := fmt.Sprintf("/v1/project/%s/integration/%s", url.PathEscape(project), url.PathEscape(integrationId))
	b, err := h.doer.Do(ctx, "ServiceIntegrationUpdate", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceIntegrationUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.ServiceIntegration, nil
}

type EndpointType string

const (
	EndpointTypeAutoscaler                   EndpointType = "autoscaler"
	EndpointTypeDatadog                      EndpointType = "datadog"
	EndpointTypeExternalAwsCloudwatchLogs    EndpointType = "external_aws_cloudwatch_logs"
	EndpointTypeExternalAwsCloudwatchMetrics EndpointType = "external_aws_cloudwatch_metrics"
	EndpointTypeExternalAwsS3                EndpointType = "external_aws_s3"
	EndpointTypeExternalClickhouse           EndpointType = "external_clickhouse"
	EndpointTypeExternalElasticsearchLogs    EndpointType = "external_elasticsearch_logs"
	EndpointTypeExternalGoogleCloudBigquery  EndpointType = "external_google_cloud_bigquery"
	EndpointTypeExternalGoogleCloudLogging   EndpointType = "external_google_cloud_logging"
	EndpointTypeExternalKafka                EndpointType = "external_kafka"
	EndpointTypeExternalMysql                EndpointType = "external_mysql"
	EndpointTypeExternalOpensearchLogs       EndpointType = "external_opensearch_logs"
	EndpointTypeExternalPostgresql           EndpointType = "external_postgresql"
	EndpointTypeExternalRedis                EndpointType = "external_redis"
	EndpointTypeExternalSchemaRegistry       EndpointType = "external_schema_registry"
	EndpointTypeExternalSumologicLogs        EndpointType = "external_sumologic_logs"
	EndpointTypeJolokia                      EndpointType = "jolokia"
	EndpointTypePrometheus                   EndpointType = "prometheus"
	EndpointTypeRsyslog                      EndpointType = "rsyslog"
)

func EndpointTypeChoices() []string {
	return []string{"autoscaler", "datadog", "external_aws_cloudwatch_logs", "external_aws_cloudwatch_metrics", "external_aws_s3", "external_clickhouse", "external_elasticsearch_logs", "external_google_cloud_bigquery", "external_google_cloud_logging", "external_kafka", "external_mysql", "external_opensearch_logs", "external_postgresql", "external_redis", "external_schema_registry", "external_sumologic_logs", "jolokia", "prometheus", "rsyslog"}
}

type EndpointTypeOut struct {
	EndpointType     string         `json:"endpoint_type"`
	ServiceTypes     []string       `json:"service_types"`
	Title            string         `json:"title"`
	UserConfigSchema map[string]any `json:"user_config_schema"`
}
type IntegrationStatusOut struct {
	State          StateOut `json:"state"`
	StatusUserDesc string   `json:"status_user_desc"`
}
type IntegrationStatusType string

const (
	IntegrationStatusTypeFailed   IntegrationStatusType = "failed"
	IntegrationStatusTypeInactive IntegrationStatusType = "inactive"
	IntegrationStatusTypeRunning  IntegrationStatusType = "running"
	IntegrationStatusTypeStarting IntegrationStatusType = "starting"
	IntegrationStatusTypeUnknown  IntegrationStatusType = "unknown"
)

func IntegrationStatusTypeChoices() []string {
	return []string{"failed", "inactive", "running", "starting", "unknown"}
}

type IntegrationType string

const (
	IntegrationTypeAlertmanager                      IntegrationType = "alertmanager"
	IntegrationTypeAutoscaler                        IntegrationType = "autoscaler"
	IntegrationTypeCaching                           IntegrationType = "caching"
	IntegrationTypeCassandraCrossServiceCluster      IntegrationType = "cassandra_cross_service_cluster"
	IntegrationTypeClickhouseCredentials             IntegrationType = "clickhouse_credentials"
	IntegrationTypeClickhouseKafka                   IntegrationType = "clickhouse_kafka"
	IntegrationTypeClickhousePostgresql              IntegrationType = "clickhouse_postgresql"
	IntegrationTypeDashboard                         IntegrationType = "dashboard"
	IntegrationTypeDatadog                           IntegrationType = "datadog"
	IntegrationTypeDatasource                        IntegrationType = "datasource"
	IntegrationTypeExternalAwsCloudwatchLogs         IntegrationType = "external_aws_cloudwatch_logs"
	IntegrationTypeExternalAwsCloudwatchMetrics      IntegrationType = "external_aws_cloudwatch_metrics"
	IntegrationTypeExternalElasticsearchLogs         IntegrationType = "external_elasticsearch_logs"
	IntegrationTypeExternalGoogleCloudLogging        IntegrationType = "external_google_cloud_logging"
	IntegrationTypeExternalOpensearchLogs            IntegrationType = "external_opensearch_logs"
	IntegrationTypeFlink                             IntegrationType = "flink"
	IntegrationTypeFlinkExternalBigquery             IntegrationType = "flink_external_bigquery"
	IntegrationTypeFlinkExternalKafka                IntegrationType = "flink_external_kafka"
	IntegrationTypeFlinkExternalPostgresql           IntegrationType = "flink_external_postgresql"
	IntegrationTypeInternalConnectivity              IntegrationType = "internal_connectivity"
	IntegrationTypeJolokia                           IntegrationType = "jolokia"
	IntegrationTypeKafkaConnect                      IntegrationType = "kafka_connect"
	IntegrationTypeKafkaConnectPostgresql            IntegrationType = "kafka_connect_postgresql"
	IntegrationTypeKafkaLogs                         IntegrationType = "kafka_logs"
	IntegrationTypeKafkaMirrormaker                  IntegrationType = "kafka_mirrormaker"
	IntegrationTypeLogs                              IntegrationType = "logs"
	IntegrationTypeM3Aggregator                      IntegrationType = "m3aggregator"
	IntegrationTypeM3Coordinator                     IntegrationType = "m3coordinator"
	IntegrationTypeMetrics                           IntegrationType = "metrics"
	IntegrationTypeOpensearchCrossClusterReplication IntegrationType = "opensearch_cross_cluster_replication"
	IntegrationTypeOpensearchCrossClusterSearch      IntegrationType = "opensearch_cross_cluster_search"
	IntegrationTypePrometheus                        IntegrationType = "prometheus"
	IntegrationTypeReadReplica                       IntegrationType = "read_replica"
	IntegrationTypeRsyslog                           IntegrationType = "rsyslog"
	IntegrationTypeSchemaRegistryProxy               IntegrationType = "schema_registry_proxy"
	IntegrationTypeStresstester                      IntegrationType = "stresstester"
	IntegrationTypeThanosMigrate                     IntegrationType = "thanos_migrate"
	IntegrationTypeThanoscompactor                   IntegrationType = "thanoscompactor"
	IntegrationTypeThanosquery                       IntegrationType = "thanosquery"
	IntegrationTypeThanosstore                       IntegrationType = "thanosstore"
	IntegrationTypeVector                            IntegrationType = "vector"
	IntegrationTypeVmalert                           IntegrationType = "vmalert"
)

func IntegrationTypeChoices() []string {
	return []string{"alertmanager", "autoscaler", "caching", "cassandra_cross_service_cluster", "clickhouse_credentials", "clickhouse_kafka", "clickhouse_postgresql", "dashboard", "datadog", "datasource", "external_aws_cloudwatch_logs", "external_aws_cloudwatch_metrics", "external_elasticsearch_logs", "external_google_cloud_logging", "external_opensearch_logs", "flink", "flink_external_bigquery", "flink_external_kafka", "flink_external_postgresql", "internal_connectivity", "jolokia", "kafka_connect", "kafka_connect_postgresql", "kafka_logs", "kafka_mirrormaker", "logs", "m3aggregator", "m3coordinator", "metrics", "opensearch_cross_cluster_replication", "opensearch_cross_cluster_search", "prometheus", "read_replica", "rsyslog", "schema_registry_proxy", "stresstester", "thanos_migrate", "thanoscompactor", "thanosquery", "thanosstore", "vector", "vmalert"}
}

type IntegrationTypeOut struct {
	DestDescription    string         `json:"dest_description"`
	DestServiceType    string         `json:"dest_service_type"`
	DestServiceTypes   []string       `json:"dest_service_types"`
	IntegrationType    string         `json:"integration_type"`
	SourceDescription  string         `json:"source_description"`
	SourceServiceTypes []string       `json:"source_service_types"`
	UserConfigSchema   map[string]any `json:"user_config_schema"`
}
type LikelyErrorCauseType string

const (
	LikelyErrorCauseTypeNull        LikelyErrorCauseType = "null"
	LikelyErrorCauseTypeDestination LikelyErrorCauseType = "destination"
	LikelyErrorCauseTypeIntegration LikelyErrorCauseType = "integration"
	LikelyErrorCauseTypeSource      LikelyErrorCauseType = "source"
	LikelyErrorCauseTypeUnknown     LikelyErrorCauseType = "unknown"
)

func LikelyErrorCauseTypeChoices() []string {
	return []string{"null", "destination", "integration", "source", "unknown"}
}

type ServiceIntegrationCreateIn struct {
	DestEndpointId   *string         `json:"dest_endpoint_id,omitempty"`
	DestProject      *string         `json:"dest_project,omitempty"`
	DestService      *string         `json:"dest_service,omitempty"`
	IntegrationType  IntegrationType `json:"integration_type"`
	SourceEndpointId *string         `json:"source_endpoint_id,omitempty"`
	SourceProject    *string         `json:"source_project,omitempty"`
	SourceService    *string         `json:"source_service,omitempty"`
	UserConfig       *map[string]any `json:"user_config,omitempty"`
}
type ServiceIntegrationCreateOut struct {
	Active               bool                  `json:"active"`
	Description          string                `json:"description"`
	DestEndpoint         *string               `json:"dest_endpoint,omitempty"`
	DestEndpointId       *string               `json:"dest_endpoint_id,omitempty"`
	DestProject          string                `json:"dest_project"`
	DestService          *string               `json:"dest_service,omitempty"`
	DestServiceType      string                `json:"dest_service_type"`
	Enabled              bool                  `json:"enabled"`
	IntegrationStatus    *IntegrationStatusOut `json:"integration_status,omitempty"`
	IntegrationType      string                `json:"integration_type"`
	ServiceIntegrationId string                `json:"service_integration_id"`
	SourceEndpoint       *string               `json:"source_endpoint,omitempty"`
	SourceEndpointId     *string               `json:"source_endpoint_id,omitempty"`
	SourceProject        string                `json:"source_project"`
	SourceService        string                `json:"source_service"`
	SourceServiceType    string                `json:"source_service_type"`
	UserConfig           map[string]any        `json:"user_config,omitempty"`
}
type ServiceIntegrationEndpointCreateIn struct {
	EndpointName string         `json:"endpoint_name"`
	EndpointType EndpointType   `json:"endpoint_type"`
	UserConfig   map[string]any `json:"user_config"`
}
type ServiceIntegrationEndpointCreateOut struct {
	EndpointConfig map[string]any `json:"endpoint_config"`
	EndpointId     string         `json:"endpoint_id"`
	EndpointName   string         `json:"endpoint_name"`
	EndpointType   EndpointType   `json:"endpoint_type"`
	UserConfig     map[string]any `json:"user_config"`
}
type ServiceIntegrationEndpointGetOut struct {
	EndpointConfig map[string]any `json:"endpoint_config"`
	EndpointId     string         `json:"endpoint_id"`
	EndpointName   string         `json:"endpoint_name"`
	EndpointType   EndpointType   `json:"endpoint_type"`
	UserConfig     map[string]any `json:"user_config"`
}
type ServiceIntegrationEndpointOut struct {
	EndpointConfig map[string]any `json:"endpoint_config"`
	EndpointId     string         `json:"endpoint_id"`
	EndpointName   string         `json:"endpoint_name"`
	EndpointType   EndpointType   `json:"endpoint_type"`
	UserConfig     map[string]any `json:"user_config"`
}
type ServiceIntegrationEndpointUpdateIn struct {
	UserConfig map[string]any `json:"user_config"`
}
type ServiceIntegrationEndpointUpdateOut struct {
	EndpointConfig map[string]any `json:"endpoint_config"`
	EndpointId     string         `json:"endpoint_id"`
	EndpointName   string         `json:"endpoint_name"`
	EndpointType   EndpointType   `json:"endpoint_type"`
	UserConfig     map[string]any `json:"user_config"`
}
type ServiceIntegrationGetOut struct {
	Active               bool                  `json:"active"`
	Description          string                `json:"description"`
	DestEndpoint         *string               `json:"dest_endpoint,omitempty"`
	DestEndpointId       *string               `json:"dest_endpoint_id,omitempty"`
	DestProject          string                `json:"dest_project"`
	DestService          *string               `json:"dest_service,omitempty"`
	DestServiceType      string                `json:"dest_service_type"`
	Enabled              bool                  `json:"enabled"`
	IntegrationStatus    *IntegrationStatusOut `json:"integration_status,omitempty"`
	IntegrationType      string                `json:"integration_type"`
	ServiceIntegrationId string                `json:"service_integration_id"`
	SourceEndpoint       *string               `json:"source_endpoint,omitempty"`
	SourceEndpointId     *string               `json:"source_endpoint_id,omitempty"`
	SourceProject        string                `json:"source_project"`
	SourceService        string                `json:"source_service"`
	SourceServiceType    string                `json:"source_service_type"`
	UserConfig           map[string]any        `json:"user_config,omitempty"`
}
type ServiceIntegrationOut struct {
	Active               bool                  `json:"active"`
	Description          string                `json:"description"`
	DestEndpoint         *string               `json:"dest_endpoint,omitempty"`
	DestEndpointId       *string               `json:"dest_endpoint_id,omitempty"`
	DestProject          string                `json:"dest_project"`
	DestService          *string               `json:"dest_service,omitempty"`
	DestServiceType      string                `json:"dest_service_type"`
	Enabled              bool                  `json:"enabled"`
	IntegrationStatus    *IntegrationStatusOut `json:"integration_status,omitempty"`
	IntegrationType      string                `json:"integration_type"`
	ServiceIntegrationId string                `json:"service_integration_id"`
	SourceEndpoint       *string               `json:"source_endpoint,omitempty"`
	SourceEndpointId     *string               `json:"source_endpoint_id,omitempty"`
	SourceProject        string                `json:"source_project"`
	SourceService        string                `json:"source_service"`
	SourceServiceType    string                `json:"source_service_type"`
	UserConfig           map[string]any        `json:"user_config,omitempty"`
}
type ServiceIntegrationUpdateIn struct {
	UserConfig map[string]any `json:"user_config"`
}
type ServiceIntegrationUpdateOut struct {
	Active               bool                  `json:"active"`
	Description          string                `json:"description"`
	DestEndpoint         *string               `json:"dest_endpoint,omitempty"`
	DestEndpointId       *string               `json:"dest_endpoint_id,omitempty"`
	DestProject          string                `json:"dest_project"`
	DestService          *string               `json:"dest_service,omitempty"`
	DestServiceType      string                `json:"dest_service_type"`
	Enabled              bool                  `json:"enabled"`
	IntegrationStatus    *IntegrationStatusOut `json:"integration_status,omitempty"`
	IntegrationType      string                `json:"integration_type"`
	ServiceIntegrationId string                `json:"service_integration_id"`
	SourceEndpoint       *string               `json:"source_endpoint,omitempty"`
	SourceEndpointId     *string               `json:"source_endpoint_id,omitempty"`
	SourceProject        string                `json:"source_project"`
	SourceService        string                `json:"source_service"`
	SourceServiceType    string                `json:"source_service_type"`
	UserConfig           map[string]any        `json:"user_config,omitempty"`
}
type StateOut struct {
	Errors           []string              `json:"errors"`
	LikelyErrorCause LikelyErrorCauseType  `json:"likely_error_cause,omitempty"`
	Nodes            map[string]any        `json:"nodes"`
	Status           IntegrationStatusType `json:"status"`
}
type serviceIntegrationCreateOut struct {
	ServiceIntegration ServiceIntegrationCreateOut `json:"service_integration"`
}
type serviceIntegrationEndpointCreateOut struct {
	ServiceIntegrationEndpoint ServiceIntegrationEndpointCreateOut `json:"service_integration_endpoint"`
}
type serviceIntegrationEndpointGetOut struct {
	ServiceIntegrationEndpoint ServiceIntegrationEndpointGetOut `json:"service_integration_endpoint"`
}
type serviceIntegrationEndpointListOut struct {
	ServiceIntegrationEndpoints []ServiceIntegrationEndpointOut `json:"service_integration_endpoints"`
}
type serviceIntegrationEndpointTypesOut struct {
	EndpointTypes []EndpointTypeOut `json:"endpoint_types"`
}
type serviceIntegrationEndpointUpdateOut struct {
	ServiceIntegrationEndpoint ServiceIntegrationEndpointUpdateOut `json:"service_integration_endpoint"`
}
type serviceIntegrationGetOut struct {
	ServiceIntegration ServiceIntegrationGetOut `json:"service_integration"`
}
type serviceIntegrationListOut struct {
	ServiceIntegrations []ServiceIntegrationOut `json:"service_integrations"`
}
type serviceIntegrationTypesOut struct {
	IntegrationTypes []IntegrationTypeOut `json:"integration_types"`
}
type serviceIntegrationUpdateOut struct {
	ServiceIntegration ServiceIntegrationUpdateOut `json:"service_integration"`
}
