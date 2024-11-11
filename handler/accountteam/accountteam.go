// Code generated by Aiven. DO NOT EDIT.

package accountteam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// Deprecated: AccountTeamCreate create a new team
	// POST /v1/account/{account_id}/teams
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamCreate
	AccountTeamCreate(ctx context.Context, accountId string, in *AccountTeamCreateIn) (*AccountTeamCreateOut, error)

	// Deprecated: AccountTeamDelete delete a team
	// DELETE /v1/account/{account_id}/team/{team_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamDelete
	AccountTeamDelete(ctx context.Context, accountId string, teamId string) error

	// Deprecated: AccountTeamGet get details for a single team
	// GET /v1/account/{account_id}/team/{team_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamGet
	AccountTeamGet(ctx context.Context, accountId string, teamId string) (*AccountTeamGetOut, error)

	// Deprecated: AccountTeamInvitesList list pending invites
	// GET /v1/account/{account_id}/team/{team_id}/invites
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamInvitesList
	AccountTeamInvitesList(ctx context.Context, accountId string, teamId string) ([]AccountInviteOut, error)

	// Deprecated: AccountTeamList list teams belonging to an account
	// GET /v1/account/{account_id}/teams
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamList
	AccountTeamList(ctx context.Context, accountId string) ([]TeamOut, error)

	// Deprecated: AccountTeamProjectAssociate associate team to a project
	// POST /v1/account/{account_id}/team/{team_id}/project/{project}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamProjectAssociate
	AccountTeamProjectAssociate(ctx context.Context, accountId string, teamId string, project string, in *AccountTeamProjectAssociateIn) error

	// Deprecated: AccountTeamProjectAssociationUpdate update team-project association
	// PUT /v1/account/{account_id}/team/{team_id}/project/{project}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamProjectAssociationUpdate
	AccountTeamProjectAssociationUpdate(ctx context.Context, accountId string, teamId string, project string, in *AccountTeamProjectAssociationUpdateIn) error

	// Deprecated: AccountTeamProjectDisassociate disassociate team from a project
	// DELETE /v1/account/{account_id}/team/{team_id}/project/{project}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamProjectDisassociate
	AccountTeamProjectDisassociate(ctx context.Context, accountId string, teamId string, project string) error

	// Deprecated: AccountTeamProjectList list projects associated to a team
	// GET /v1/account/{account_id}/team/{team_id}/projects
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamProjectList
	AccountTeamProjectList(ctx context.Context, accountId string, teamId string) ([]ProjectOut, error)

	// Deprecated: AccountTeamUpdate update team details
	// PUT /v1/account/{account_id}/team/{team_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamUpdate
	AccountTeamUpdate(ctx context.Context, accountId string, teamId string, in *AccountTeamUpdateIn) (*AccountTeamUpdateOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) AccountTeamHandler {
	return AccountTeamHandler{doer}
}

type AccountTeamHandler struct {
	doer doer
}

func (h *AccountTeamHandler) AccountTeamCreate(ctx context.Context, accountId string, in *AccountTeamCreateIn) (*AccountTeamCreateOut, error) {
	path := fmt.Sprintf("/v1/account/%s/teams", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountTeamCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(accountTeamCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Team, nil
}
func (h *AccountTeamHandler) AccountTeamDelete(ctx context.Context, accountId string, teamId string) error {
	path := fmt.Sprintf("/v1/account/%s/team/%s", url.PathEscape(accountId), url.PathEscape(teamId))
	_, err := h.doer.Do(ctx, "AccountTeamDelete", "DELETE", path, nil)
	return err
}
func (h *AccountTeamHandler) AccountTeamGet(ctx context.Context, accountId string, teamId string) (*AccountTeamGetOut, error) {
	path := fmt.Sprintf("/v1/account/%s/team/%s", url.PathEscape(accountId), url.PathEscape(teamId))
	b, err := h.doer.Do(ctx, "AccountTeamGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountTeamGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Team, nil
}
func (h *AccountTeamHandler) AccountTeamInvitesList(ctx context.Context, accountId string, teamId string) ([]AccountInviteOut, error) {
	path := fmt.Sprintf("/v1/account/%s/team/%s/invites", url.PathEscape(accountId), url.PathEscape(teamId))
	b, err := h.doer.Do(ctx, "AccountTeamInvitesList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountTeamInvitesListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AccountInvites, nil
}
func (h *AccountTeamHandler) AccountTeamList(ctx context.Context, accountId string) ([]TeamOut, error) {
	path := fmt.Sprintf("/v1/account/%s/teams", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountTeamList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountTeamListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Teams, nil
}
func (h *AccountTeamHandler) AccountTeamProjectAssociate(ctx context.Context, accountId string, teamId string, project string, in *AccountTeamProjectAssociateIn) error {
	path := fmt.Sprintf("/v1/account/%s/team/%s/project/%s", url.PathEscape(accountId), url.PathEscape(teamId), url.PathEscape(project))
	_, err := h.doer.Do(ctx, "AccountTeamProjectAssociate", "POST", path, in)
	return err
}
func (h *AccountTeamHandler) AccountTeamProjectAssociationUpdate(ctx context.Context, accountId string, teamId string, project string, in *AccountTeamProjectAssociationUpdateIn) error {
	path := fmt.Sprintf("/v1/account/%s/team/%s/project/%s", url.PathEscape(accountId), url.PathEscape(teamId), url.PathEscape(project))
	_, err := h.doer.Do(ctx, "AccountTeamProjectAssociationUpdate", "PUT", path, in)
	return err
}
func (h *AccountTeamHandler) AccountTeamProjectDisassociate(ctx context.Context, accountId string, teamId string, project string) error {
	path := fmt.Sprintf("/v1/account/%s/team/%s/project/%s", url.PathEscape(accountId), url.PathEscape(teamId), url.PathEscape(project))
	_, err := h.doer.Do(ctx, "AccountTeamProjectDisassociate", "DELETE", path, nil)
	return err
}
func (h *AccountTeamHandler) AccountTeamProjectList(ctx context.Context, accountId string, teamId string) ([]ProjectOut, error) {
	path := fmt.Sprintf("/v1/account/%s/team/%s/projects", url.PathEscape(accountId), url.PathEscape(teamId))
	b, err := h.doer.Do(ctx, "AccountTeamProjectList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountTeamProjectListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Projects, nil
}
func (h *AccountTeamHandler) AccountTeamUpdate(ctx context.Context, accountId string, teamId string, in *AccountTeamUpdateIn) (*AccountTeamUpdateOut, error) {
	path := fmt.Sprintf("/v1/account/%s/team/%s", url.PathEscape(accountId), url.PathEscape(teamId))
	b, err := h.doer.Do(ctx, "AccountTeamUpdate", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(accountTeamUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Team, nil
}

type AccountInviteOut struct {
	AccountId          string    `json:"account_id"`            // Account ID
	AccountName        string    `json:"account_name"`          // Account name
	CreateTime         time.Time `json:"create_time"`           // Timestamp in ISO 8601 format, always in UTC
	InvitedByUserEmail string    `json:"invited_by_user_email"` // User email address
	TeamId             string    `json:"team_id"`               // Team ID
	TeamName           string    `json:"team_name"`             // Team name
	UserEmail          string    `json:"user_email"`            // User email address
}

// AccountTeamCreateIn AccountTeamCreateRequestBody
type AccountTeamCreateIn struct {
	TeamName string `json:"team_name"` // Team name
}

// AccountTeamCreateOut Account Team details
type AccountTeamCreateOut struct {
	AccountId  *string    `json:"account_id,omitempty"`  // Account ID
	CreateTime *time.Time `json:"create_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
	TeamId     string     `json:"team_id"`               // Team ID
	TeamName   string     `json:"team_name"`             // Team name
	UpdateTime *time.Time `json:"update_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
}

// AccountTeamGetOut Account Team details
type AccountTeamGetOut struct {
	AccountId  *string    `json:"account_id,omitempty"`  // Account ID
	CreateTime *time.Time `json:"create_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
	TeamId     string     `json:"team_id"`               // Team ID
	TeamName   string     `json:"team_name"`             // Team name
	UpdateTime *time.Time `json:"update_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
}

// AccountTeamProjectAssociateIn AccountTeamProjectAssociateRequestBody
type AccountTeamProjectAssociateIn struct {
	TeamType TeamType `json:"team_type"` // Team type (permission level)
}

// AccountTeamProjectAssociationUpdateIn AccountTeamProjectAssociationUpdateRequestBody
type AccountTeamProjectAssociationUpdateIn struct {
	TeamType TeamType `json:"team_type,omitempty"` // Team type (permission level)
}

// AccountTeamUpdateIn AccountTeamUpdateRequestBody
type AccountTeamUpdateIn struct {
	TeamName string `json:"team_name"` // Team name
}

// AccountTeamUpdateOut Account Team details
type AccountTeamUpdateOut struct {
	AccountId  *string    `json:"account_id,omitempty"`  // Account ID
	CreateTime *time.Time `json:"create_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
	TeamId     string     `json:"team_id"`               // Team ID
	TeamName   string     `json:"team_name"`             // Team name
	UpdateTime *time.Time `json:"update_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
}
type ProjectOut struct {
	ProjectName string   `json:"project_name"` // Project name
	TeamType    TeamType `json:"team_type"`    // Team type (permission level)
}
type TeamOut struct {
	AccountId  *string    `json:"account_id,omitempty"`  // Account ID
	CreateTime *time.Time `json:"create_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
	TeamId     string     `json:"team_id"`               // Team ID
	TeamName   string     `json:"team_name"`             // Team name
	UpdateTime *time.Time `json:"update_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
}
type TeamType string

const (
	TeamTypeAdmin                        TeamType = "admin"
	TeamTypeOperator                     TeamType = "operator"
	TeamTypeDeveloper                    TeamType = "developer"
	TeamTypeReadOnly                     TeamType = "read_only"
	TeamTypeProjectIntegrationsRead      TeamType = "project:integrations:read"
	TeamTypeProjectIntegrationsWrite     TeamType = "project:integrations:write"
	TeamTypeProjectNetworkingRead        TeamType = "project:networking:read"
	TeamTypeProjectNetworkingWrite       TeamType = "project:networking:write"
	TeamTypeProjectPermissionsRead       TeamType = "project:permissions:read"
	TeamTypeServiceConfigurationWrite    TeamType = "service:configuration:write"
	TeamTypeServiceLogsRead              TeamType = "service:logs:read"
	TeamTypeProjectServicesRead          TeamType = "project:services:read"
	TeamTypeProjectServicesWrite         TeamType = "project:services:write"
	TeamTypeProjectAuditLogsRead         TeamType = "project:audit_logs:read"
	TeamTypeServiceDataWrite             TeamType = "service:data:write"
	TeamTypeServiceSecretsRead           TeamType = "service:secrets:read"
	TeamTypeRoleServicesMaintenance      TeamType = "role:services:maintenance"
	TeamTypeRoleServicesRecover          TeamType = "role:services:recover"
	TeamTypeOrganizationBillingRead      TeamType = "organization:billing:read"
	TeamTypeOrganizationBillingWrite     TeamType = "organization:billing:write"
	TeamTypeOrganizationAuditLogsRead    TeamType = "organization:audit_logs:read"
	TeamTypeOrganizationProjectsRead     TeamType = "organization:projects:read"
	TeamTypeOrganizationProjectsWrite    TeamType = "organization:projects:write"
	TeamTypeOrganizationUsersWrite       TeamType = "organization:users:write"
	TeamTypeOrganizationPermissionsRead  TeamType = "organization:permissions:read"
	TeamTypeOrganizationPermissionsWrite TeamType = "organization:permissions:write"
	TeamTypeOrganizationAppUsersWrite    TeamType = "organization:app_users:write"
	TeamTypeOrganizationGroupsWrite      TeamType = "organization:groups:write"
	TeamTypeOrganizationIdpsWrite        TeamType = "organization:idps:write"
	TeamTypeOrganizationDomainsWrite     TeamType = "organization:domains:write"
	TeamTypeOrganizationNetworkRead      TeamType = "organization:network:read"
	TeamTypeOrganizationNetworkWrite     TeamType = "organization:network:write"
	TeamTypeRoleOrganizationAdmin        TeamType = "role:organization:admin"
	TeamTypeServiceUsersWrite            TeamType = "service:users:write"
)

func TeamTypeChoices() []string {
	return []string{"admin", "operator", "developer", "read_only", "project:integrations:read", "project:integrations:write", "project:networking:read", "project:networking:write", "project:permissions:read", "service:configuration:write", "service:logs:read", "project:services:read", "project:services:write", "project:audit_logs:read", "service:data:write", "service:secrets:read", "role:services:maintenance", "role:services:recover", "organization:billing:read", "organization:billing:write", "organization:audit_logs:read", "organization:projects:read", "organization:projects:write", "organization:users:write", "organization:permissions:read", "organization:permissions:write", "organization:app_users:write", "organization:groups:write", "organization:idps:write", "organization:domains:write", "organization:network:read", "organization:network:write", "role:organization:admin", "service:users:write"}
}

// accountTeamCreateOut AccountTeamCreateResponse
type accountTeamCreateOut struct {
	Team AccountTeamCreateOut `json:"team"` // Account Team details
}

// accountTeamGetOut AccountTeamGetResponse
type accountTeamGetOut struct {
	Team AccountTeamGetOut `json:"team"` // Account Team details
}

// accountTeamInvitesListOut AccountTeamInvitesListResponse
type accountTeamInvitesListOut struct {
	AccountInvites []AccountInviteOut `json:"account_invites"` // List of invites
}

// accountTeamListOut AccountTeamListResponse
type accountTeamListOut struct {
	Teams []TeamOut `json:"teams"` // List of teams
}

// accountTeamProjectListOut AccountTeamProjectListResponse
type accountTeamProjectListOut struct {
	Projects []ProjectOut `json:"projects"` // List of projects associated to a team
}

// accountTeamUpdateOut AccountTeamUpdateResponse
type accountTeamUpdateOut struct {
	Team AccountTeamUpdateOut `json:"team"` // Account Team details
}
