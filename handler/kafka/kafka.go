// Code generated by Aiven. DO NOT EDIT.

package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// ServiceKafkaAclAdd add Aiven Kafka ACL entry
	// POST /v1/project/{project}/service/{service_name}/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaAclAdd
	ServiceKafkaAclAdd(ctx context.Context, project string, serviceName string, in *ServiceKafkaAclAddIn) ([]AclOut, error)

	// ServiceKafkaAclDelete delete a Kafka ACL entry
	// DELETE /v1/project/{project}/service/{service_name}/acl/{kafka_acl_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaAclDelete
	ServiceKafkaAclDelete(ctx context.Context, project string, serviceName string, kafkaAclId string) ([]AclOut, error)

	// ServiceKafkaAclList list Aiven ACL entries for Kafka service
	// GET /v1/project/{project}/service/{service_name}/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaAclList
	ServiceKafkaAclList(ctx context.Context, project string, serviceName string) ([]AclOut, error)

	// ServiceKafkaNativeAclAdd add a Kafka-native ACL entry
	// POST /v1/project/{project}/service/{service_name}/kafka/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaNativeAclAdd
	ServiceKafkaNativeAclAdd(ctx context.Context, project string, serviceName string, in *ServiceKafkaNativeAclAddIn) (*ServiceKafkaNativeAclAddOut, error)

	// ServiceKafkaNativeAclDelete delete a Kafka-native ACL entry
	// DELETE /v1/project/{project}/service/{service_name}/kafka/acl/{kafka_acl_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaNativeAclDelete
	ServiceKafkaNativeAclDelete(ctx context.Context, project string, serviceName string, kafkaAclId string) error

	// ServiceKafkaNativeAclGet get single Kafka-native ACL entry
	// GET /v1/project/{project}/service/{service_name}/kafka/acl/{kafka_acl_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaNativeAclGet
	ServiceKafkaNativeAclGet(ctx context.Context, project string, serviceName string, kafkaAclId string) (*ServiceKafkaNativeAclGetOut, error)

	// ServiceKafkaNativeAclList list Kafka-native ACL entries
	// GET /v1/project/{project}/service/{service_name}/kafka/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaNativeAclList
	ServiceKafkaNativeAclList(ctx context.Context, project string, serviceName string) (*ServiceKafkaNativeAclListOut, error)

	// ServiceKafkaQuotaCreate create Kafka quota
	// POST /v1/project/{project}/service/{service_name}/quota
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaQuotaCreate
	ServiceKafkaQuotaCreate(ctx context.Context, project string, serviceName string, in *ServiceKafkaQuotaCreateIn) error

	// ServiceKafkaQuotaDelete delete Kafka quota
	// DELETE /v1/project/{project}/service/{service_name}/quota
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaQuotaDelete
	ServiceKafkaQuotaDelete(ctx context.Context, project string, serviceName string, query ...[2]string) error

	// ServiceKafkaQuotaDescribe get service quota configuration
	// GET /v1/project/{project}/service/{service_name}/quota/describe
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaQuotaDescribe
	ServiceKafkaQuotaDescribe(ctx context.Context, project string, serviceName string, query ...[2]string) (*ServiceKafkaQuotaDescribeOut, error)

	// ServiceKafkaQuotaList list Kafka quotas
	// GET /v1/project/{project}/service/{service_name}/quota
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaQuotaList
	ServiceKafkaQuotaList(ctx context.Context, project string, serviceName string) ([]QuotaOut, error)

	// ServiceKafkaTieredStorageStorageUsageByTopic get the Kafka tiered storage object storage usage by topic
	// GET /v1/project/{project}/service/{service_name}/kafka/tiered-storage/storage-usage/by-topic
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTieredStorageStorageUsageByTopic
	ServiceKafkaTieredStorageStorageUsageByTopic(ctx context.Context, project string, serviceName string) (map[string]int, error)

	// ServiceKafkaTieredStorageStorageUsageTotal get the Kafka tiered storage total object storage usage
	// GET /v1/project/{project}/service/{service_name}/kafka/tiered-storage/storage-usage/total
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTieredStorageStorageUsageTotal
	ServiceKafkaTieredStorageStorageUsageTotal(ctx context.Context, project string, serviceName string) (int, error)

	// ServiceKafkaTieredStorageSummary get the Kafka tiered storage summary
	// GET /v1/project/{project}/service/{service_name}/kafka/tiered-storage/summary
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTieredStorageSummary
	ServiceKafkaTieredStorageSummary(ctx context.Context, project string, serviceName string) (*ServiceKafkaTieredStorageSummaryOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) KafkaHandler {
	return KafkaHandler{doer}
}

type KafkaHandler struct {
	doer doer
}

func (h *KafkaHandler) ServiceKafkaAclAdd(ctx context.Context, project string, serviceName string, in *ServiceKafkaAclAddIn) ([]AclOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaAclAdd", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaAclAddOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaHandler) ServiceKafkaAclDelete(ctx context.Context, project string, serviceName string, kafkaAclId string) ([]AclOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/acl/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(kafkaAclId))
	b, err := h.doer.Do(ctx, "ServiceKafkaAclDelete", "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaAclDeleteOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaHandler) ServiceKafkaAclList(ctx context.Context, project string, serviceName string) ([]AclOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaAclList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaAclListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaHandler) ServiceKafkaNativeAclAdd(ctx context.Context, project string, serviceName string, in *ServiceKafkaNativeAclAddIn) (*ServiceKafkaNativeAclAddOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaNativeAclAdd", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaNativeAclAddOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Acl, nil
}
func (h *KafkaHandler) ServiceKafkaNativeAclDelete(ctx context.Context, project string, serviceName string, kafkaAclId string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/acl/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(kafkaAclId))
	_, err := h.doer.Do(ctx, "ServiceKafkaNativeAclDelete", "DELETE", path, nil)
	return err
}
func (h *KafkaHandler) ServiceKafkaNativeAclGet(ctx context.Context, project string, serviceName string, kafkaAclId string) (*ServiceKafkaNativeAclGetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/acl/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(kafkaAclId))
	b, err := h.doer.Do(ctx, "ServiceKafkaNativeAclGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaNativeAclGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Acl, nil
}
func (h *KafkaHandler) ServiceKafkaNativeAclList(ctx context.Context, project string, serviceName string) (*ServiceKafkaNativeAclListOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaNativeAclList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceKafkaNativeAclListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *KafkaHandler) ServiceKafkaQuotaCreate(ctx context.Context, project string, serviceName string, in *ServiceKafkaQuotaCreateIn) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/quota", url.PathEscape(project), url.PathEscape(serviceName))
	_, err := h.doer.Do(ctx, "ServiceKafkaQuotaCreate", "POST", path, in)
	return err
}

// ServiceKafkaQuotaDeleteClientId Client ID.
func ServiceKafkaQuotaDeleteClientId(clientId string) [2]string {
	return [2]string{"client-id", clientId}
}

// ServiceKafkaQuotaDeleteUser Username.
func ServiceKafkaQuotaDeleteUser(user string) [2]string {
	return [2]string{"user", user}
}
func (h *KafkaHandler) ServiceKafkaQuotaDelete(ctx context.Context, project string, serviceName string, query ...[2]string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/quota", url.PathEscape(project), url.PathEscape(serviceName))
	_, err := h.doer.Do(ctx, "ServiceKafkaQuotaDelete", "DELETE", path, nil, query...)
	return err
}

// ServiceKafkaQuotaDescribeUser
func ServiceKafkaQuotaDescribeUser(user string) [2]string {
	return [2]string{"user", user}
}

// ServiceKafkaQuotaDescribeClientId
func ServiceKafkaQuotaDescribeClientId(clientId string) [2]string {
	return [2]string{"client-id", clientId}
}
func (h *KafkaHandler) ServiceKafkaQuotaDescribe(ctx context.Context, project string, serviceName string, query ...[2]string) (*ServiceKafkaQuotaDescribeOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/quota/describe", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaQuotaDescribe", "GET", path, nil, query...)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaQuotaDescribeOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Quota, nil
}
func (h *KafkaHandler) ServiceKafkaQuotaList(ctx context.Context, project string, serviceName string) ([]QuotaOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/quota", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaQuotaList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaQuotaListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Quotas, nil
}
func (h *KafkaHandler) ServiceKafkaTieredStorageStorageUsageByTopic(ctx context.Context, project string, serviceName string) (map[string]int, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/tiered-storage/storage-usage/by-topic", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaTieredStorageStorageUsageByTopic", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaTieredStorageUsageByTopicOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.StorageUsage, nil
}
func (h *KafkaHandler) ServiceKafkaTieredStorageStorageUsageTotal(ctx context.Context, project string, serviceName string) (int, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/tiered-storage/storage-usage/total", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaTieredStorageStorageUsageTotal", "GET", path, nil)
	if err != nil {
		return 0, err
	}
	out := new(serviceKafkaTieredStorageUsageTotalOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return 0, err
	}
	return out.TotalStorageUsage, nil
}
func (h *KafkaHandler) ServiceKafkaTieredStorageSummary(ctx context.Context, project string, serviceName string) (*ServiceKafkaTieredStorageSummaryOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/tiered-storage/summary", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaTieredStorageSummary", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceKafkaTieredStorageSummaryOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type AclOut struct {
	Id         *string        `json:"id,omitempty"` // ID
	Permission PermissionType `json:"permission"`   // Kafka permission
	Topic      string         `json:"topic"`        // Topic name pattern
	Username   string         `json:"username"`
}
type HourlyOut struct {
	EstimatedCost   *string `json:"estimated_cost,omitempty"` // The estimated cost in USD of tiered storage for this hour
	HourStart       string  `json:"hour_start"`               // Timestamp in ISO 8601 format, always in UTC
	PeakStoredBytes int     `json:"peak_stored_bytes"`        // Peak bytes stored on object storage at this hour
}
type KafkaAclOut struct {
	Host           string                 `json:"host"`            // the host or * for all hosts
	Id             string                 `json:"id"`              // ID
	Operation      OperationType          `json:"operation"`       // Kafka ACL operation represents an operation which an ACL grants or denies permission to perform
	PatternType    PatternType            `json:"pattern_type"`    // Kafka ACL pattern type of resource name
	PermissionType KafkaAclPermissionType `json:"permission_type"` // Kafka ACL permission type
	Principal      string                 `json:"principal"`       // principal is in 'principalType:name' format
	ResourceName   string                 `json:"resource_name"`   // Resource pattern used to match specified resources
	ResourceType   ResourceType           `json:"resource_type"`   // Kafka ACL resource type represents a type of resource which an ACL can be applied to
}
type KafkaAclPermissionType string

const (
	KafkaAclPermissionTypeAllow KafkaAclPermissionType = "ALLOW"
	KafkaAclPermissionTypeDeny  KafkaAclPermissionType = "DENY"
)

func KafkaAclPermissionTypeChoices() []string {
	return []string{"ALLOW", "DENY"}
}

type OperationType string

const (
	OperationTypeAll             OperationType = "All"
	OperationTypeAlter           OperationType = "Alter"
	OperationTypeAlterConfigs    OperationType = "AlterConfigs"
	OperationTypeClusterAction   OperationType = "ClusterAction"
	OperationTypeCreate          OperationType = "Create"
	OperationTypeCreateTokens    OperationType = "CreateTokens"
	OperationTypeDelete          OperationType = "Delete"
	OperationTypeDescribe        OperationType = "Describe"
	OperationTypeDescribeConfigs OperationType = "DescribeConfigs"
	OperationTypeDescribeTokens  OperationType = "DescribeTokens"
	OperationTypeIdempotentWrite OperationType = "IdempotentWrite"
	OperationTypeRead            OperationType = "Read"
	OperationTypeWrite           OperationType = "Write"
)

func OperationTypeChoices() []string {
	return []string{"All", "Alter", "AlterConfigs", "ClusterAction", "Create", "CreateTokens", "Delete", "Describe", "DescribeConfigs", "DescribeTokens", "IdempotentWrite", "Read", "Write"}
}

type PatternType string

const (
	PatternTypeLiteral  PatternType = "LITERAL"
	PatternTypePrefixed PatternType = "PREFIXED"
)

func PatternTypeChoices() []string {
	return []string{"LITERAL", "PREFIXED"}
}

type PermissionType string

const (
	PermissionTypeAdmin     PermissionType = "admin"
	PermissionTypeRead      PermissionType = "read"
	PermissionTypeReadwrite PermissionType = "readwrite"
	PermissionTypeWrite     PermissionType = "write"
)

func PermissionTypeChoices() []string {
	return []string{"admin", "read", "readwrite", "write"}
}

type QuotaOut struct {
	ClientId          string  `json:"client-id"`          // Represents a logical group of clients, assigned a unique name by the client application. Quotas can be applied based on user, client-id, or both. The most relevant quota is chosen for each connection. All connections within a quota group share the same quota.
	ConsumerByteRate  float64 `json:"consumer_byte_rate"` // Defines the bandwidth limit in bytes/sec for each group of clients sharing a quota. Every distinct client group is allocated a specific quota, as defined by the cluster, on a per-broker basis. Exceeding this limit results in client throttling.
	ProducerByteRate  float64 `json:"producer_byte_rate"` // Defines the bandwidth limit in bytes/sec for each group of clients sharing a quota. Every distinct client group is allocated a specific quota, as defined by the cluster, on a per-broker basis. Exceeding this limit results in client throttling.
	RequestPercentage float64 `json:"request_percentage"` // Sets the maximum percentage of CPU time that a client group can use on request handler I/O and network threads per broker within a quota window. Exceeding this limit triggers throttling. The quota, expressed as a percentage, also indicates the total allowable CPU usage for the client groups sharing the quota.
	User              string  `json:"user"`               // Represents a logical group of clients, assigned a unique name by the client application. Quotas can be applied based on user, client-id, or both. The most relevant quota is chosen for each connection. All connections within a quota group share the same quota.
}
type ResourceType string

const (
	ResourceTypeCluster         ResourceType = "Cluster"
	ResourceTypeDelegationToken ResourceType = "DelegationToken"
	ResourceTypeGroup           ResourceType = "Group"
	ResourceTypeTopic           ResourceType = "Topic"
	ResourceTypeTransactionalId ResourceType = "TransactionalId"
	ResourceTypeUser            ResourceType = "User"
)

func ResourceTypeChoices() []string {
	return []string{"Cluster", "DelegationToken", "Group", "Topic", "TransactionalId", "User"}
}

// ServiceKafkaAclAddIn ServiceKafkaAclAddRequestBody
type ServiceKafkaAclAddIn struct {
	Permission PermissionType `json:"permission"` // Kafka permission
	Topic      string         `json:"topic"`      // Topic name pattern
	Username   string         `json:"username"`
}

// ServiceKafkaNativeAclAddIn ServiceKafkaNativeAclAddRequestBody
type ServiceKafkaNativeAclAddIn struct {
	Host           *string                                `json:"host,omitempty"`  // the host or * for all hosts
	Operation      OperationType                          `json:"operation"`       // Kafka ACL operation represents an operation which an ACL grants or denies permission to perform
	PatternType    PatternType                            `json:"pattern_type"`    // Kafka ACL pattern type of resource name
	PermissionType ServiceKafkaNativeAclAddPermissionType `json:"permission_type"` // Kafka ACL permission type
	Principal      string                                 `json:"principal"`       // principal is in 'PrincipalType:name' format
	ResourceName   string                                 `json:"resource_name"`   // Resource pattern used to match specified resources
	ResourceType   ResourceType                           `json:"resource_type"`   // Kafka ACL resource type represents a type of resource which an ACL can be applied to
}

// ServiceKafkaNativeAclAddOut Kafka-native ACL entry for Kafka service
type ServiceKafkaNativeAclAddOut struct {
	Host           string                                 `json:"host"`            // the host or * for all hosts
	Id             string                                 `json:"id"`              // ID
	Operation      OperationType                          `json:"operation"`       // Kafka ACL operation represents an operation which an ACL grants or denies permission to perform
	PatternType    PatternType                            `json:"pattern_type"`    // Kafka ACL pattern type of resource name
	PermissionType ServiceKafkaNativeAclAddPermissionType `json:"permission_type"` // Kafka ACL permission type
	Principal      string                                 `json:"principal"`       // principal is in 'principalType:name' format
	ResourceName   string                                 `json:"resource_name"`   // Resource pattern used to match specified resources
	ResourceType   ResourceType                           `json:"resource_type"`   // Kafka ACL resource type represents a type of resource which an ACL can be applied to
}
type ServiceKafkaNativeAclAddPermissionType string

const (
	ServiceKafkaNativeAclAddPermissionTypeAllow ServiceKafkaNativeAclAddPermissionType = "ALLOW"
	ServiceKafkaNativeAclAddPermissionTypeDeny  ServiceKafkaNativeAclAddPermissionType = "DENY"
)

func ServiceKafkaNativeAclAddPermissionTypeChoices() []string {
	return []string{"ALLOW", "DENY"}
}

// ServiceKafkaNativeAclGetOut Kafka-native ACL entry for Kafka service
type ServiceKafkaNativeAclGetOut struct {
	Host           string                                 `json:"host"`            // the host or * for all hosts
	Id             string                                 `json:"id"`              // ID
	Operation      OperationType                          `json:"operation"`       // Kafka ACL operation represents an operation which an ACL grants or denies permission to perform
	PatternType    PatternType                            `json:"pattern_type"`    // Kafka ACL pattern type of resource name
	PermissionType ServiceKafkaNativeAclGetPermissionType `json:"permission_type"` // Kafka ACL permission type
	Principal      string                                 `json:"principal"`       // principal is in 'principalType:name' format
	ResourceName   string                                 `json:"resource_name"`   // Resource pattern used to match specified resources
	ResourceType   ResourceType                           `json:"resource_type"`   // Kafka ACL resource type represents a type of resource which an ACL can be applied to
}
type ServiceKafkaNativeAclGetPermissionType string

const (
	ServiceKafkaNativeAclGetPermissionTypeAllow ServiceKafkaNativeAclGetPermissionType = "ALLOW"
	ServiceKafkaNativeAclGetPermissionTypeDeny  ServiceKafkaNativeAclGetPermissionType = "DENY"
)

func ServiceKafkaNativeAclGetPermissionTypeChoices() []string {
	return []string{"ALLOW", "DENY"}
}

// ServiceKafkaNativeAclListOut ServiceKafkaNativeAclListResponse
type ServiceKafkaNativeAclListOut struct {
	Acl      []AclOut      `json:"acl"`       // List of Aiven ACL entries for Kafka service
	KafkaAcl []KafkaAclOut `json:"kafka_acl"` // List of Kafka-native ACL entries
}

// ServiceKafkaQuotaCreateIn ServiceKafkaQuotaCreateRequestBody
type ServiceKafkaQuotaCreateIn struct {
	ClientId          *string  `json:"client-id,omitempty"`          // Represents a logical group of clients, assigned a unique name by the client application. Quotas can be applied based on user, client-id, or both. The most relevant quota is chosen for each connection. All connections within a quota group share the same quota.
	ConsumerByteRate  *float64 `json:"consumer_byte_rate,omitempty"` // Defines the bandwidth limit in bytes/sec for each group of clients sharing a quota. Every distinct client group is allocated a specific quota, as defined by the cluster, on a per-broker basis. Exceeding this limit results in client throttling.
	ProducerByteRate  *float64 `json:"producer_byte_rate,omitempty"` // Defines the bandwidth limit in bytes/sec for each group of clients sharing a quota. Every distinct client group is allocated a specific quota, as defined by the cluster, on a per-broker basis. Exceeding this limit results in client throttling.
	RequestPercentage *float64 `json:"request_percentage,omitempty"` // Sets the maximum percentage of CPU time that a client group can use on request handler I/O and network threads per broker within a quota window. Exceeding this limit triggers throttling. The quota, expressed as a percentage, also indicates the total allowable CPU usage for the client groups sharing the quota.
	User              *string  `json:"user,omitempty"`               // Represents a logical group of clients, assigned a unique name by the client application. Quotas can be applied based on user, client-id, or both. The most relevant quota is chosen for each connection. All connections within a quota group share the same quota.
}

// ServiceKafkaQuotaDescribeOut kafka quota
type ServiceKafkaQuotaDescribeOut struct {
	ClientId          *string  `json:"client-id,omitempty"`          // Represents a logical group of clients, assigned a unique name by the client application. Quotas can be applied based on user, client-id, or both. The most relevant quota is chosen for each connection. All connections within a quota group share the same quota.
	ConsumerByteRate  *float64 `json:"consumer_byte_rate,omitempty"` // Defines the bandwidth limit in bytes/sec for each group of clients sharing a quota. Every distinct client group is allocated a specific quota, as defined by the cluster, on a per-broker basis. Exceeding this limit results in client throttling.
	ProducerByteRate  *float64 `json:"producer_byte_rate,omitempty"` // Defines the bandwidth limit in bytes/sec for each group of clients sharing a quota. Every distinct client group is allocated a specific quota, as defined by the cluster, on a per-broker basis. Exceeding this limit results in client throttling.
	RequestPercentage *float64 `json:"request_percentage,omitempty"` // Sets the maximum percentage of CPU time that a client group can use on request handler I/O and network threads per broker within a quota window. Exceeding this limit triggers throttling. The quota, expressed as a percentage, also indicates the total allowable CPU usage for the client groups sharing the quota.
	User              *string  `json:"user,omitempty"`               // Represents a logical group of clients, assigned a unique name by the client application. Quotas can be applied based on user, client-id, or both. The most relevant quota is chosen for each connection. All connections within a quota group share the same quota.
}

// ServiceKafkaTieredStorageSummaryOut ServiceKafkaTieredStorageSummaryResponse
type ServiceKafkaTieredStorageSummaryOut struct {
	CurrentCost         string                 `json:"current_cost"`              // The current cost in USD of tiered storage since the beginning of the billing period
	ForecastedCost      string                 `json:"forecasted_cost"`           // The forecasted cost in USD of tiered storage in the billing period
	ForecastedRate      *string                `json:"forecasted_rate,omitempty"` // The rate on GBs/hour used to calculate the forecasted cost
	StorageUsageHistory StorageUsageHistoryOut `json:"storage_usage_history"`     // History of usage and cumulative costs in the billing period
	TotalStorageUsage   int                    `json:"total_storage_usage"`       // Total storage usage by tiered storage, in bytes
}

// StorageUsageHistoryOut History of usage and cumulative costs in the billing period
type StorageUsageHistoryOut struct {
	Hourly []HourlyOut `json:"hourly"` // History by hour
}

// serviceKafkaAclAddOut ServiceKafkaAclAddResponse
type serviceKafkaAclAddOut struct {
	Acl []AclOut `json:"acl"` // List of Aiven ACL entries for Kafka service
}

// serviceKafkaAclDeleteOut ServiceKafkaAclDeleteResponse
type serviceKafkaAclDeleteOut struct {
	Acl []AclOut `json:"acl"` // List of Aiven ACL entries for Kafka service
}

// serviceKafkaAclListOut ServiceKafkaAclListResponse
type serviceKafkaAclListOut struct {
	Acl []AclOut `json:"acl"` // List of Aiven ACL entries for Kafka service
}

// serviceKafkaNativeAclAddOut ServiceKafkaNativeAclAddResponse
type serviceKafkaNativeAclAddOut struct {
	Acl ServiceKafkaNativeAclAddOut `json:"acl"` // Kafka-native ACL entry for Kafka service
}

// serviceKafkaNativeAclGetOut ServiceKafkaNativeAclGetResponse
type serviceKafkaNativeAclGetOut struct {
	Acl ServiceKafkaNativeAclGetOut `json:"acl"` // Kafka-native ACL entry for Kafka service
}

// serviceKafkaQuotaDescribeOut ServiceKafkaQuotaDescribeResponse
type serviceKafkaQuotaDescribeOut struct {
	Quota ServiceKafkaQuotaDescribeOut `json:"quota"` // kafka quota
}

// serviceKafkaQuotaListOut ServiceKafkaQuotaListResponse
type serviceKafkaQuotaListOut struct {
	Quotas []QuotaOut `json:"quotas"` // List of kafka quotas
}

// serviceKafkaTieredStorageUsageByTopicOut ServiceKafkaTieredStorageStorageUsageByTopicResponse
type serviceKafkaTieredStorageUsageByTopicOut struct {
	StorageUsage map[string]int `json:"storage_usage"` // Storage usage by tiered storage by topics
}

// serviceKafkaTieredStorageUsageTotalOut ServiceKafkaTieredStorageStorageUsageTotalResponse
type serviceKafkaTieredStorageUsageTotalOut struct {
	TotalStorageUsage int `json:"total_storage_usage"` // Total storage usage by tiered storage, in bytes
}
