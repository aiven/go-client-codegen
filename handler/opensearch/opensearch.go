// Code generated by Aiven. DO NOT EDIT.

package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// ServiceOpenSearchAclGet show OpenSearch ACL configuration
	// GET /v1/project/{project}/service/{service_name}/opensearch/acl
	// https://api.aiven.io/doc/#tag/Service:_OpenSearch/operation/ServiceOpenSearchAclGet
	ServiceOpenSearchAclGet(ctx context.Context, project string, serviceName string) (*ServiceOpenSearchAclGetOut, error)

	// ServiceOpenSearchAclSet set OpenSearch ACL configuration
	// POST /v1/project/{project}/service/{service_name}/opensearch/acl
	// https://api.aiven.io/doc/#tag/Service:_OpenSearch/operation/ServiceOpenSearchAclSet
	ServiceOpenSearchAclSet(ctx context.Context, project string, serviceName string, in *ServiceOpenSearchAclSetIn) (*ServiceOpenSearchAclSetOut, error)

	// ServiceOpenSearchAclUpdate update OpenSearch ACL configuration
	// PUT /v1/project/{project}/service/{service_name}/opensearch/acl
	// https://api.aiven.io/doc/#tag/Service:_OpenSearch/operation/ServiceOpenSearchAclUpdate
	ServiceOpenSearchAclUpdate(ctx context.Context, project string, serviceName string, in *ServiceOpenSearchAclUpdateIn) (*ServiceOpenSearchAclUpdateOut, error)

	// ServiceOpenSearchIndexDelete delete an OpenSearch index
	// DELETE /v1/project/{project}/service/{service_name}/index/{index_name}
	// https://api.aiven.io/doc/#tag/Service:_OpenSearch/operation/ServiceOpenSearchIndexDelete
	ServiceOpenSearchIndexDelete(ctx context.Context, project string, serviceName string, indexName string) error

	// ServiceOpenSearchIndexList list OpenSearch indexes
	// GET /v1/project/{project}/service/{service_name}/index
	// https://api.aiven.io/doc/#tag/Service:_OpenSearch/operation/ServiceOpenSearchIndexList
	ServiceOpenSearchIndexList(ctx context.Context, project string, serviceName string) ([]IndexeOut, error)

	// ServiceOpenSearchSecurityGet show OpenSearch security configuration status
	// GET /v1/project/{project}/service/{service_name}/opensearch/security
	// https://api.aiven.io/doc/#tag/Service:_OpenSearch/operation/ServiceOpenSearchSecurityGet
	ServiceOpenSearchSecurityGet(ctx context.Context, project string, serviceName string) (*ServiceOpenSearchSecurityGetOut, error)

	// ServiceOpenSearchSecurityReset change Opensearch Security Admin password
	// PUT /v1/project/{project}/service/{service_name}/opensearch/security/admin
	// https://api.aiven.io/doc/#tag/Service:_OpenSearch/operation/ServiceOpenSearchSecurityReset
	ServiceOpenSearchSecurityReset(ctx context.Context, project string, serviceName string, in *ServiceOpenSearchSecurityResetIn) (*ServiceOpenSearchSecurityResetOut, error)

	// ServiceOpenSearchSecuritySet enable Opensearch Security Admin by setting the password
	// POST /v1/project/{project}/service/{service_name}/opensearch/security/admin
	// https://api.aiven.io/doc/#tag/Service:_OpenSearch/operation/ServiceOpenSearchSecuritySet
	ServiceOpenSearchSecuritySet(ctx context.Context, project string, serviceName string, in *ServiceOpenSearchSecuritySetIn) (*ServiceOpenSearchSecuritySetOut, error)
}

func NewHandler(doer doer) OpenSearchHandler {
	return OpenSearchHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type OpenSearchHandler struct {
	doer doer
}

func (h *OpenSearchHandler) ServiceOpenSearchAclGet(ctx context.Context, project string, serviceName string) (*ServiceOpenSearchAclGetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/opensearch/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceOpenSearchAclGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceOpenSearchAclGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.OpensearchAclConfig, nil
}
func (h *OpenSearchHandler) ServiceOpenSearchAclSet(ctx context.Context, project string, serviceName string, in *ServiceOpenSearchAclSetIn) (*ServiceOpenSearchAclSetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/opensearch/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceOpenSearchAclSet", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceOpenSearchAclSetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.OpensearchAclConfig, nil
}
func (h *OpenSearchHandler) ServiceOpenSearchAclUpdate(ctx context.Context, project string, serviceName string, in *ServiceOpenSearchAclUpdateIn) (*ServiceOpenSearchAclUpdateOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/opensearch/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceOpenSearchAclUpdate", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceOpenSearchAclUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.OpensearchAclConfig, nil
}
func (h *OpenSearchHandler) ServiceOpenSearchIndexDelete(ctx context.Context, project string, serviceName string, indexName string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/index/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(indexName))
	_, err := h.doer.Do(ctx, "ServiceOpenSearchIndexDelete", "DELETE", path, nil)
	return err
}
func (h *OpenSearchHandler) ServiceOpenSearchIndexList(ctx context.Context, project string, serviceName string) ([]IndexeOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/index", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceOpenSearchIndexList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceOpenSearchIndexListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Indexes, nil
}
func (h *OpenSearchHandler) ServiceOpenSearchSecurityGet(ctx context.Context, project string, serviceName string) (*ServiceOpenSearchSecurityGetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/opensearch/security", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceOpenSearchSecurityGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceOpenSearchSecurityGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OpenSearchHandler) ServiceOpenSearchSecurityReset(ctx context.Context, project string, serviceName string, in *ServiceOpenSearchSecurityResetIn) (*ServiceOpenSearchSecurityResetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/opensearch/security/admin", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceOpenSearchSecurityReset", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceOpenSearchSecurityResetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OpenSearchHandler) ServiceOpenSearchSecuritySet(ctx context.Context, project string, serviceName string, in *ServiceOpenSearchSecuritySetIn) (*ServiceOpenSearchSecuritySetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/opensearch/security/admin", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceOpenSearchSecuritySet", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceOpenSearchSecuritySetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type AclIn struct {
	Rules    []RuleIn `json:"rules"`
	Username string   `json:"username"`
}
type AclOut struct {
	Rules    []RuleOut `json:"rules"`
	Username string    `json:"username"`
}
type HealthType string

const (
	HealthTypeGreen       HealthType = "green"
	HealthTypeYellow      HealthType = "yellow"
	HealthTypeRed         HealthType = "red"
	HealthTypeRedAsterisk HealthType = "red*"
	HealthTypeUnknown     HealthType = "unknown"
)

func HealthTypeChoices() []string {
	return []string{"green", "yellow", "red", "red*", "unknown"}
}

type IndexeOut struct {
	CreateTime          time.Time        `json:"create_time"`
	Docs                *int             `json:"docs,omitempty"`
	Health              HealthType       `json:"health,omitempty"`
	IndexName           string           `json:"index_name"`
	NumberOfReplicas    int              `json:"number_of_replicas"`
	NumberOfShards      int              `json:"number_of_shards"`
	ReadOnlyAllowDelete *bool            `json:"read_only_allow_delete,omitempty"`
	Replication         *ReplicationOut  `json:"replication,omitempty"`
	Size                *int             `json:"size,omitempty"`
	Status              IndexeStatusType `json:"status,omitempty"`
}
type IndexeStatusType string

const (
	IndexeStatusTypeUnknown IndexeStatusType = "unknown"
	IndexeStatusTypeOpen    IndexeStatusType = "open"
	IndexeStatusTypeClose   IndexeStatusType = "close"
	IndexeStatusTypeNone    IndexeStatusType = "none"
)

func IndexeStatusTypeChoices() []string {
	return []string{"unknown", "open", "close", "none"}
}

type OpensearchAclConfigIn struct {
	Acls    []AclIn `json:"acls"`
	Enabled bool    `json:"enabled"`
}
type OpensearchAclConfigInAlt struct {
	Acls    *[]AclIn `json:"acls,omitempty"`
	Enabled *bool    `json:"enabled,omitempty"`
}
type PermissionType string

const (
	PermissionTypeDeny      PermissionType = "deny"
	PermissionTypeAdmin     PermissionType = "admin"
	PermissionTypeRead      PermissionType = "read"
	PermissionTypeReadwrite PermissionType = "readwrite"
	PermissionTypeWrite     PermissionType = "write"
)

func PermissionTypeChoices() []string {
	return []string{"deny", "admin", "read", "readwrite", "write"}
}

type ReplicationOut struct {
	LeaderIndex   string `json:"leader_index,omitempty"`
	LeaderProject string `json:"leader_project,omitempty"`
	LeaderService string `json:"leader_service,omitempty"`
}
type RuleIn struct {
	Index      string         `json:"index"`
	Permission PermissionType `json:"permission"`
}
type RuleOut struct {
	Index      string         `json:"index"`
	Permission PermissionType `json:"permission"`
}
type ServiceOpenSearchAclGetOut struct {
	Acls    []AclOut `json:"acls"`
	Enabled bool     `json:"enabled"`
}
type ServiceOpenSearchAclSetIn struct {
	OpensearchAclConfig OpensearchAclConfigIn `json:"opensearch_acl_config"`
}
type ServiceOpenSearchAclSetOut struct {
	Acls    []AclOut `json:"acls"`
	Enabled bool     `json:"enabled"`
}
type ServiceOpenSearchAclUpdateIn struct {
	OpensearchAclConfig OpensearchAclConfigInAlt `json:"opensearch_acl_config"`
}
type ServiceOpenSearchAclUpdateOut struct {
	Acls    []AclOut `json:"acls"`
	Enabled bool     `json:"enabled"`
}
type ServiceOpenSearchSecurityGetOut struct {
	SecurityPluginAdminEnabled bool  `json:"security_plugin_admin_enabled"`
	SecurityPluginAvailable    bool  `json:"security_plugin_available"`
	SecurityPluginEnabled      *bool `json:"security_plugin_enabled,omitempty"`
}
type ServiceOpenSearchSecurityResetIn struct {
	AdminPassword string `json:"admin_password"`
	NewPassword   string `json:"new_password"`
}
type ServiceOpenSearchSecurityResetOut struct {
	SecurityPluginAdminEnabled bool  `json:"security_plugin_admin_enabled"`
	SecurityPluginAvailable    bool  `json:"security_plugin_available"`
	SecurityPluginEnabled      *bool `json:"security_plugin_enabled,omitempty"`
}
type ServiceOpenSearchSecuritySetIn struct {
	AdminPassword string `json:"admin_password"`
}
type ServiceOpenSearchSecuritySetOut struct {
	SecurityPluginAdminEnabled bool  `json:"security_plugin_admin_enabled"`
	SecurityPluginAvailable    bool  `json:"security_plugin_available"`
	SecurityPluginEnabled      *bool `json:"security_plugin_enabled,omitempty"`
}
type serviceOpenSearchAclGetOut struct {
	OpensearchAclConfig ServiceOpenSearchAclGetOut `json:"opensearch_acl_config"`
}
type serviceOpenSearchAclSetOut struct {
	OpensearchAclConfig ServiceOpenSearchAclSetOut `json:"opensearch_acl_config"`
}
type serviceOpenSearchAclUpdateOut struct {
	OpensearchAclConfig ServiceOpenSearchAclUpdateOut `json:"opensearch_acl_config"`
}
type serviceOpenSearchIndexListOut struct {
	Indexes []IndexeOut `json:"indexes"`
}
