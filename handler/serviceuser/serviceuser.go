// Code generated by Aiven. DO NOT EDIT.

package serviceuser

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// ServiceUserCreate create a new (sub) user for service
	// POST /project/{project}/service/{service_name}/user
	// https://api.aiven.io/doc/#tag/Service/operation/ServiceUserCreate
	ServiceUserCreate(ctx context.Context, project string, serviceName string, in *ServiceUserCreateIn) (*ServiceUserCreateOut, error)

	// ServiceUserCredentialsModify modify service user credentials
	// PUT /project/{project}/service/{service_name}/user/{service_username}
	// https://api.aiven.io/doc/#tag/Service/operation/ServiceUserCredentialsModify
	ServiceUserCredentialsModify(ctx context.Context, project string, serviceName string, serviceUsername string, in *ServiceUserCredentialsModifyIn) (*ServiceUserCredentialsModifyOut, error)

	// ServiceUserCredentialsReset reset service user credentials
	// PUT /project/{project}/service/{service_name}/user/{service_username}/credentials/reset
	// https://api.aiven.io/doc/#tag/Service/operation/ServiceUserCredentialsReset
	ServiceUserCredentialsReset(ctx context.Context, project string, serviceName string, serviceUsername string) (*ServiceUserCredentialsResetOut, error)

	// ServiceUserDelete delete a service user
	// DELETE /project/{project}/service/{service_name}/user/{service_username}
	// https://api.aiven.io/doc/#tag/Service/operation/ServiceUserDelete
	ServiceUserDelete(ctx context.Context, project string, serviceName string, serviceUsername string) error

	// ServiceUserGet get details for a single user
	// GET /project/{project}/service/{service_name}/user/{service_username}
	// https://api.aiven.io/doc/#tag/Service/operation/ServiceUserGet
	ServiceUserGet(ctx context.Context, project string, serviceName string, serviceUsername string) (*ServiceUserGetOut, error)
}

func NewHandler(doer doer) ServiceUserHandler {
	return ServiceUserHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type ServiceUserHandler struct {
	doer doer
}

func (h *ServiceUserHandler) ServiceUserCreate(ctx context.Context, project string, serviceName string, in *ServiceUserCreateIn) (*ServiceUserCreateOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/user", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceUserCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceUserCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.User, nil
}
func (h *ServiceUserHandler) ServiceUserCredentialsModify(ctx context.Context, project string, serviceName string, serviceUsername string, in *ServiceUserCredentialsModifyIn) (*ServiceUserCredentialsModifyOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/user/%s", project, serviceName, serviceUsername)
	b, err := h.doer.Do(ctx, "ServiceUserCredentialsModify", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceUserCredentialsModifyOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Service, nil
}
func (h *ServiceUserHandler) ServiceUserCredentialsReset(ctx context.Context, project string, serviceName string, serviceUsername string) (*ServiceUserCredentialsResetOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/user/%s/credentials/reset", project, serviceName, serviceUsername)
	b, err := h.doer.Do(ctx, "ServiceUserCredentialsReset", "PUT", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceUserCredentialsResetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Service, nil
}
func (h *ServiceUserHandler) ServiceUserDelete(ctx context.Context, project string, serviceName string, serviceUsername string) error {
	path := fmt.Sprintf("/project/%s/service/%s/user/%s", project, serviceName, serviceUsername)
	_, err := h.doer.Do(ctx, "ServiceUserDelete", "DELETE", path, nil)
	return err
}
func (h *ServiceUserHandler) ServiceUserGet(ctx context.Context, project string, serviceName string, serviceUsername string) (*ServiceUserGetOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/user/%s", project, serviceName, serviceUsername)
	b, err := h.doer.Do(ctx, "ServiceUserGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceUserGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.User, nil
}

type AccessControlIn struct {
	M3Group            string    `json:"m3_group,omitempty"`
	PgAllowReplication *bool     `json:"pg_allow_replication,omitempty"`
	RedisAclCategories *[]string `json:"redis_acl_categories,omitempty"`
	RedisAclChannels   *[]string `json:"redis_acl_channels,omitempty"`
	RedisAclCommands   *[]string `json:"redis_acl_commands,omitempty"`
	RedisAclKeys       *[]string `json:"redis_acl_keys,omitempty"`
}
type AccessControlOut struct {
	M3Group            string   `json:"m3_group,omitempty"`
	PgAllowReplication *bool    `json:"pg_allow_replication,omitempty"`
	RedisAclCategories []string `json:"redis_acl_categories,omitempty"`
	RedisAclChannels   []string `json:"redis_acl_channels,omitempty"`
	RedisAclCommands   []string `json:"redis_acl_commands,omitempty"`
	RedisAclKeys       []string `json:"redis_acl_keys,omitempty"`
}
type AclOut struct {
	Id         string         `json:"id,omitempty"`
	Permission PermissionType `json:"permission"`
	Topic      string         `json:"topic"`
	Username   string         `json:"username"`
}
type AdditionalRegionOut struct {
	Cloud       string `json:"cloud"`
	PauseReason string `json:"pause_reason,omitempty"`
	Paused      *bool  `json:"paused,omitempty"`
	Region      string `json:"region,omitempty"`
}
type AuthenticationType string

const (
	AuthenticationTypeNull                AuthenticationType = "null"
	AuthenticationTypeCachingSha2Password AuthenticationType = "caching_sha2_password"
	AuthenticationTypeMysqlNativePassword AuthenticationType = "mysql_native_password"
)

func AuthenticationTypeChoices() []string {
	return []string{"null", "caching_sha2_password", "mysql_native_password"}
}

type BackupOut struct {
	AdditionalRegions []AdditionalRegionOut `json:"additional_regions,omitempty"`
	BackupName        string                `json:"backup_name"`
	BackupTime        time.Time             `json:"backup_time"`
	DataSize          int                   `json:"data_size"`
	StorageLocation   string                `json:"storage_location,omitempty"`
}
type ComponentOut struct {
	Component                 string                        `json:"component"`
	Host                      string                        `json:"host"`
	KafkaAuthenticationMethod KafkaAuthenticationMethodType `json:"kafka_authentication_method,omitempty"`
	Path                      string                        `json:"path,omitempty"`
	Port                      int                           `json:"port"`
	PrivatelinkConnectionId   string                        `json:"privatelink_connection_id,omitempty"`
	Route                     RouteType                     `json:"route"`
	Ssl                       *bool                         `json:"ssl,omitempty"`
	Usage                     UsageType                     `json:"usage"`
}
type ConnectionPoolOut struct {
	ConnectionUri string       `json:"connection_uri"`
	Database      string       `json:"database"`
	PoolMode      PoolModeType `json:"pool_mode"`
	PoolName      string       `json:"pool_name"`
	PoolSize      int          `json:"pool_size"`
	Username      string       `json:"username,omitempty"`
}
type DowType string

const (
	DowTypeMonday    DowType = "monday"
	DowTypeTuesday   DowType = "tuesday"
	DowTypeWednesday DowType = "wednesday"
	DowTypeThursday  DowType = "thursday"
	DowTypeFriday    DowType = "friday"
	DowTypeSaturday  DowType = "saturday"
	DowTypeSunday    DowType = "sunday"
	DowTypeNever     DowType = "never"
)

func DowTypeChoices() []string {
	return []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday", "never"}
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

type KafkaAuthenticationMethodType string

const (
	KafkaAuthenticationMethodTypeCertificate KafkaAuthenticationMethodType = "certificate"
	KafkaAuthenticationMethodTypeSasl        KafkaAuthenticationMethodType = "sasl"
)

func KafkaAuthenticationMethodTypeChoices() []string {
	return []string{"certificate", "sasl"}
}

type LevelType string

const (
	LevelTypeNotice  LevelType = "notice"
	LevelTypeWarning LevelType = "warning"
)

func LevelTypeChoices() []string {
	return []string{"notice", "warning"}
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

type MaintenanceOut struct {
	Dow     DowType     `json:"dow"`
	Time    string      `json:"time"`
	Updates []UpdateOut `json:"updates"`
}
type MetadataOut struct {
	EndOfLifeHelpArticleUrl string     `json:"end_of_life_help_article_url,omitempty"`
	EndOfLifePolicyUrl      string     `json:"end_of_life_policy_url,omitempty"`
	ServiceEndOfLifeTime    *time.Time `json:"service_end_of_life_time,omitempty"`
	UpgradeToServiceType    string     `json:"upgrade_to_service_type,omitempty"`
	UpgradeToVersion        string     `json:"upgrade_to_version,omitempty"`
}
type NodeStateOut struct {
	Name            string              `json:"name"`
	ProgressUpdates []ProgressUpdateOut `json:"progress_updates,omitempty"`
	Role            RoleType            `json:"role,omitempty"`
	Shard           *ShardOut           `json:"shard,omitempty"`
	State           NodeStateType       `json:"state"`
}
type NodeStateType string

const (
	NodeStateTypeLeaving     NodeStateType = "leaving"
	NodeStateTypeRunning     NodeStateType = "running"
	NodeStateTypeSettingUpVm NodeStateType = "setting_up_vm"
	NodeStateTypeSyncingData NodeStateType = "syncing_data"
	NodeStateTypeTimingOut   NodeStateType = "timing_out"
	NodeStateTypeUnknown     NodeStateType = "unknown"
)

func NodeStateTypeChoices() []string {
	return []string{"leaving", "running", "setting_up_vm", "syncing_data", "timing_out", "unknown"}
}

type OperationType string

const (
	OperationTypeAcknowledgeRenewal OperationType = "acknowledge-renewal"
	OperationTypeResetCredentials   OperationType = "reset-credentials"
	OperationTypeSetAccessControl   OperationType = "set-access-control"
)

func OperationTypeChoices() []string {
	return []string{"acknowledge-renewal", "reset-credentials", "set-access-control"}
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

type PermissionTypeAlt string

const (
	PermissionTypeAltSchemaRegistryRead  PermissionTypeAlt = "schema_registry_read"
	PermissionTypeAltSchemaRegistryWrite PermissionTypeAlt = "schema_registry_write"
)

func PermissionTypeAltChoices() []string {
	return []string{"schema_registry_read", "schema_registry_write"}
}

type PhaseType string

const (
	PhaseTypePrepare    PhaseType = "prepare"
	PhaseTypeBasebackup PhaseType = "basebackup"
	PhaseTypeStream     PhaseType = "stream"
	PhaseTypeFinalize   PhaseType = "finalize"
)

func PhaseTypeChoices() []string {
	return []string{"prepare", "basebackup", "stream", "finalize"}
}

type PoolModeType string

const (
	PoolModeTypeSession     PoolModeType = "session"
	PoolModeTypeTransaction PoolModeType = "transaction"
	PoolModeTypeStatement   PoolModeType = "statement"
)

func PoolModeTypeChoices() []string {
	return []string{"session", "transaction", "statement"}
}

type ProgressUpdateOut struct {
	Completed bool      `json:"completed"`
	Current   *int      `json:"current,omitempty"`
	Max       *int      `json:"max,omitempty"`
	Min       *int      `json:"min,omitempty"`
	Phase     PhaseType `json:"phase"`
	Unit      UnitType  `json:"unit,omitempty"`
}
type RoleType string

const (
	RoleTypeMaster      RoleType = "master"
	RoleTypeStandby     RoleType = "standby"
	RoleTypeReadReplica RoleType = "read-replica"
)

func RoleTypeChoices() []string {
	return []string{"master", "standby", "read-replica"}
}

type RouteType string

const (
	RouteTypeDynamic     RouteType = "dynamic"
	RouteTypePublic      RouteType = "public"
	RouteTypePrivate     RouteType = "private"
	RouteTypePrivatelink RouteType = "privatelink"
)

func RouteTypeChoices() []string {
	return []string{"dynamic", "public", "private", "privatelink"}
}

type SchemaRegistryAclOut struct {
	Id         string            `json:"id,omitempty"`
	Permission PermissionTypeAlt `json:"permission"`
	Resource   string            `json:"resource"`
	Username   string            `json:"username"`
}
type ServiceIntegrationOut struct {
	Active               bool                  `json:"active"`
	Description          string                `json:"description"`
	DestEndpoint         string                `json:"dest_endpoint,omitempty"`
	DestEndpointId       string                `json:"dest_endpoint_id,omitempty"`
	DestProject          string                `json:"dest_project"`
	DestService          string                `json:"dest_service,omitempty"`
	DestServiceType      string                `json:"dest_service_type"`
	Enabled              bool                  `json:"enabled"`
	IntegrationStatus    *IntegrationStatusOut `json:"integration_status,omitempty"`
	IntegrationType      string                `json:"integration_type"`
	ServiceIntegrationId string                `json:"service_integration_id"`
	SourceEndpoint       string                `json:"source_endpoint,omitempty"`
	SourceEndpointId     string                `json:"source_endpoint_id,omitempty"`
	SourceProject        string                `json:"source_project"`
	SourceService        string                `json:"source_service"`
	SourceServiceType    string                `json:"source_service_type"`
	UserConfig           map[string]any        `json:"user_config,omitempty"`
}
type ServiceNotificationOut struct {
	Level    LevelType               `json:"level"`
	Message  string                  `json:"message"`
	Metadata MetadataOut             `json:"metadata"`
	Type     ServiceNotificationType `json:"type"`
}
type ServiceNotificationType string

const (
	ServiceNotificationTypeServiceEndOfLife         ServiceNotificationType = "service_end_of_life"
	ServiceNotificationTypeServicePoweredOffRemoval ServiceNotificationType = "service_powered_off_removal"
)

func ServiceNotificationTypeChoices() []string {
	return []string{"service_end_of_life", "service_powered_off_removal"}
}

type ServiceStateType string

const (
	ServiceStateTypePoweroff    ServiceStateType = "POWEROFF"
	ServiceStateTypeRebalancing ServiceStateType = "REBALANCING"
	ServiceStateTypeRebuilding  ServiceStateType = "REBUILDING"
	ServiceStateTypeRunning     ServiceStateType = "RUNNING"
)

func ServiceStateTypeChoices() []string {
	return []string{"POWEROFF", "REBALANCING", "REBUILDING", "RUNNING"}
}

type ServiceUserCreateIn struct {
	AccessControl  *AccessControlIn   `json:"access_control,omitempty"`
	Authentication AuthenticationType `json:"authentication,omitempty"`
	Username       string             `json:"username"`
}
type ServiceUserCreateOut struct {
	AccessCert                    string             `json:"access_cert,omitempty"`
	AccessCertNotValidAfterTime   *time.Time         `json:"access_cert_not_valid_after_time,omitempty"`
	AccessControl                 *AccessControlOut  `json:"access_control,omitempty"`
	AccessKey                     string             `json:"access_key,omitempty"`
	Authentication                AuthenticationType `json:"authentication,omitempty"`
	ExpiringCertNotValidAfterTime *time.Time         `json:"expiring_cert_not_valid_after_time,omitempty"`
	Password                      string             `json:"password"`
	Type                          string             `json:"type"`
	Username                      string             `json:"username"`
}
type ServiceUserCredentialsModifyIn struct {
	AccessControl  *AccessControlIn   `json:"access_control,omitempty"`
	Authentication AuthenticationType `json:"authentication,omitempty"`
	NewPassword    string             `json:"new_password,omitempty"`
	Operation      OperationType      `json:"operation"`
}
type ServiceUserCredentialsModifyOut struct {
	Acl                    []AclOut                 `json:"acl,omitempty"`
	Backups                []BackupOut              `json:"backups,omitempty"`
	CloudDescription       string                   `json:"cloud_description,omitempty"`
	CloudName              string                   `json:"cloud_name"`
	Components             []ComponentOut           `json:"components,omitempty"`
	ConnectionInfo         map[string]any           `json:"connection_info,omitempty"`
	ConnectionPools        []ConnectionPoolOut      `json:"connection_pools,omitempty"`
	CreateTime             time.Time                `json:"create_time"`
	Databases              []string                 `json:"databases,omitempty"`
	DiskSpaceMb            *float64                 `json:"disk_space_mb,omitempty"`
	Features               map[string]any           `json:"features,omitempty"`
	GroupList              []string                 `json:"group_list"`
	Maintenance            *MaintenanceOut          `json:"maintenance,omitempty"`
	Metadata               map[string]any           `json:"metadata,omitempty"`
	NodeCount              *int                     `json:"node_count,omitempty"`
	NodeCpuCount           *int                     `json:"node_cpu_count,omitempty"`
	NodeMemoryMb           *float64                 `json:"node_memory_mb,omitempty"`
	NodeStates             []NodeStateOut           `json:"node_states,omitempty"`
	Plan                   string                   `json:"plan"`
	ProjectVpcId           string                   `json:"project_vpc_id"`
	SchemaRegistryAcl      []SchemaRegistryAclOut   `json:"schema_registry_acl,omitempty"`
	ServiceIntegrations    []ServiceIntegrationOut  `json:"service_integrations"`
	ServiceName            string                   `json:"service_name"`
	ServiceNotifications   []ServiceNotificationOut `json:"service_notifications,omitempty"`
	ServiceType            string                   `json:"service_type"`
	ServiceTypeDescription string                   `json:"service_type_description,omitempty"`
	ServiceUri             string                   `json:"service_uri"`
	ServiceUriParams       map[string]any           `json:"service_uri_params,omitempty"`
	State                  ServiceStateType         `json:"state"`
	Tags                   map[string]string        `json:"tags,omitempty"`
	TechEmails             []TechEmailOut           `json:"tech_emails,omitempty"`
	TerminationProtection  bool                     `json:"termination_protection"`
	Topics                 []TopicOut               `json:"topics,omitempty"`
	UpdateTime             time.Time                `json:"update_time"`
	UserConfig             map[string]any           `json:"user_config"`
	Users                  []UserOut                `json:"users,omitempty"`
}
type ServiceUserCredentialsResetOut struct {
	Acl                    []AclOut                 `json:"acl,omitempty"`
	Backups                []BackupOut              `json:"backups,omitempty"`
	CloudDescription       string                   `json:"cloud_description,omitempty"`
	CloudName              string                   `json:"cloud_name"`
	Components             []ComponentOut           `json:"components,omitempty"`
	ConnectionInfo         map[string]any           `json:"connection_info,omitempty"`
	ConnectionPools        []ConnectionPoolOut      `json:"connection_pools,omitempty"`
	CreateTime             time.Time                `json:"create_time"`
	Databases              []string                 `json:"databases,omitempty"`
	DiskSpaceMb            *float64                 `json:"disk_space_mb,omitempty"`
	Features               map[string]any           `json:"features,omitempty"`
	GroupList              []string                 `json:"group_list"`
	Maintenance            *MaintenanceOut          `json:"maintenance,omitempty"`
	Metadata               map[string]any           `json:"metadata,omitempty"`
	NodeCount              *int                     `json:"node_count,omitempty"`
	NodeCpuCount           *int                     `json:"node_cpu_count,omitempty"`
	NodeMemoryMb           *float64                 `json:"node_memory_mb,omitempty"`
	NodeStates             []NodeStateOut           `json:"node_states,omitempty"`
	Plan                   string                   `json:"plan"`
	ProjectVpcId           string                   `json:"project_vpc_id"`
	SchemaRegistryAcl      []SchemaRegistryAclOut   `json:"schema_registry_acl,omitempty"`
	ServiceIntegrations    []ServiceIntegrationOut  `json:"service_integrations"`
	ServiceName            string                   `json:"service_name"`
	ServiceNotifications   []ServiceNotificationOut `json:"service_notifications,omitempty"`
	ServiceType            string                   `json:"service_type"`
	ServiceTypeDescription string                   `json:"service_type_description,omitempty"`
	ServiceUri             string                   `json:"service_uri"`
	ServiceUriParams       map[string]any           `json:"service_uri_params,omitempty"`
	State                  ServiceStateType         `json:"state"`
	Tags                   map[string]string        `json:"tags,omitempty"`
	TechEmails             []TechEmailOut           `json:"tech_emails,omitempty"`
	TerminationProtection  bool                     `json:"termination_protection"`
	Topics                 []TopicOut               `json:"topics,omitempty"`
	UpdateTime             time.Time                `json:"update_time"`
	UserConfig             map[string]any           `json:"user_config"`
	Users                  []UserOut                `json:"users,omitempty"`
}
type ServiceUserGetOut struct {
	AccessCert                    string             `json:"access_cert,omitempty"`
	AccessCertNotValidAfterTime   *time.Time         `json:"access_cert_not_valid_after_time,omitempty"`
	AccessControl                 *AccessControlOut  `json:"access_control,omitempty"`
	AccessKey                     string             `json:"access_key,omitempty"`
	Authentication                AuthenticationType `json:"authentication,omitempty"`
	ExpiringCertNotValidAfterTime *time.Time         `json:"expiring_cert_not_valid_after_time,omitempty"`
	Password                      string             `json:"password"`
	Type                          string             `json:"type"`
	Username                      string             `json:"username"`
}
type ShardOut struct {
	Name     string `json:"name,omitempty"`
	Position *int   `json:"position,omitempty"`
}
type StateOut struct {
	Errors           []string              `json:"errors"`
	LikelyErrorCause LikelyErrorCauseType  `json:"likely_error_cause,omitempty"`
	Nodes            map[string]any        `json:"nodes"`
	Status           IntegrationStatusType `json:"status"`
}
type TechEmailOut struct {
	Email string `json:"email"`
}
type TopicOut struct {
	CleanupPolicy     string         `json:"cleanup_policy"`
	MinInsyncReplicas int            `json:"min_insync_replicas"`
	Partitions        int            `json:"partitions"`
	Replication       int            `json:"replication"`
	RetentionBytes    int            `json:"retention_bytes"`
	RetentionHours    int            `json:"retention_hours"`
	State             TopicStateType `json:"state,omitempty"`
	TopicName         string         `json:"topic_name"`
}
type TopicStateType string

const (
	TopicStateTypeActive      TopicStateType = "ACTIVE"
	TopicStateTypeConfiguring TopicStateType = "CONFIGURING"
	TopicStateTypeDeleting    TopicStateType = "DELETING"
)

func TopicStateTypeChoices() []string {
	return []string{"ACTIVE", "CONFIGURING", "DELETING"}
}

type UnitType string

const (
	UnitTypeBinlogs           UnitType = "binlogs"
	UnitTypeBytesCompressed   UnitType = "bytes_compressed"
	UnitTypeBytesUncompressed UnitType = "bytes_uncompressed"
	UnitTypeWalLsn            UnitType = "wal_lsn"
)

func UnitTypeChoices() []string {
	return []string{"binlogs", "bytes_compressed", "bytes_uncompressed", "wal_lsn"}
}

type UpdateOut struct {
	Deadline    string     `json:"deadline,omitempty"`
	Description string     `json:"description,omitempty"`
	StartAfter  string     `json:"start_after,omitempty"`
	StartAt     *time.Time `json:"start_at,omitempty"`
}
type UsageType string

const (
	UsageTypePrimary UsageType = "primary"
	UsageTypeReplica UsageType = "replica"
)

func UsageTypeChoices() []string {
	return []string{"primary", "replica"}
}

type UserOut struct {
	AccessCert                    string             `json:"access_cert,omitempty"`
	AccessCertNotValidAfterTime   *time.Time         `json:"access_cert_not_valid_after_time,omitempty"`
	AccessControl                 *AccessControlOut  `json:"access_control,omitempty"`
	AccessKey                     string             `json:"access_key,omitempty"`
	Authentication                AuthenticationType `json:"authentication,omitempty"`
	ExpiringCertNotValidAfterTime *time.Time         `json:"expiring_cert_not_valid_after_time,omitempty"`
	Password                      string             `json:"password"`
	Type                          string             `json:"type"`
	Username                      string             `json:"username"`
}
type serviceUserCreateOut struct {
	User ServiceUserCreateOut `json:"user"`
}
type serviceUserCredentialsModifyOut struct {
	Service ServiceUserCredentialsModifyOut `json:"service"`
}
type serviceUserCredentialsResetOut struct {
	Service ServiceUserCredentialsResetOut `json:"service"`
}
type serviceUserGetOut struct {
	User ServiceUserGetOut `json:"user"`
}
