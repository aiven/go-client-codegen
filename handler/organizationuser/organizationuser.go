// Code generated by Aiven. DO NOT EDIT.

package organizationuser

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// OrganizationUserAuthenticationMethodsList list authentication methods for a user in the organization
	// GET /v1/organization/{organization_id}/user/{member_user_id}/authentication_methods
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserAuthenticationMethodsList
	OrganizationUserAuthenticationMethodsList(ctx context.Context, organizationId string, memberUserId string) ([]AuthenticationMethodOut, error)

	// OrganizationUserDelete remove a user from the organization
	// DELETE /v1/organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserDelete
	OrganizationUserDelete(ctx context.Context, organizationId string, memberUserId string) error

	// OrganizationUserGet get details on a user of the organization
	// GET /v1/organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserGet
	OrganizationUserGet(ctx context.Context, organizationId string, memberUserId string) (*OrganizationUserGetOut, error)

	// OrganizationUserInvitationAccept accept a user invitation to the organization
	// POST /v1/organization/{organization_id}/invitation/{user_email}
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUserInvitationAccept
	OrganizationUserInvitationAccept(ctx context.Context, organizationId string, userEmail string) error

	// OrganizationUserInvitationDelete remove an invitation to the organization
	// DELETE /v1/organization/{organization_id}/invitation/{user_email}
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUserInvitationDelete
	OrganizationUserInvitationDelete(ctx context.Context, organizationId string, userEmail string) error

	// OrganizationUserInvitationsList list user invitations to the organization
	// GET /v1/organization/{organization_id}/invitation
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUserInvitationsList
	OrganizationUserInvitationsList(ctx context.Context, organizationId string) ([]InvitationOut, error)

	// OrganizationUserInvite invite a user to the organization
	// POST /v1/organization/{organization_id}/invitation
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUserInvite
	OrganizationUserInvite(ctx context.Context, organizationId string, in *OrganizationUserInviteIn) error

	// OrganizationUserList list users of the organization
	// GET /v1/organization/{organization_id}/user
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserList
	OrganizationUserList(ctx context.Context, organizationId string) ([]UserOut, error)

	// OrganizationUserPasswordReset reset the password of a managed user in the organization
	// POST /v1/organization/{organization_id}/user/{member_user_id}/reset_password
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserPasswordReset
	OrganizationUserPasswordReset(ctx context.Context, organizationId string, memberUserId string) error

	// OrganizationUserRevokeToken revoke the token of a managed user in the organization
	// DELETE /v1/organization/{organization_id}/user/{member_user_id}/access-token/{token_prefix}
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserRevokeToken
	OrganizationUserRevokeToken(ctx context.Context, organizationId string, memberUserId string, tokenPrefix string) error

	// OrganizationUserTokensList list tokens from an organization's member
	// GET /v1/organization/{organization_id}/user/{member_user_id}/access-tokens
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserTokensList
	OrganizationUserTokensList(ctx context.Context, organizationId string, memberUserId string) ([]TokenOut, error)

	// OrganizationUserUpdate update details on a user of the organization
	// PATCH /v1/organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserUpdate
	OrganizationUserUpdate(ctx context.Context, organizationId string, memberUserId string, in *OrganizationUserUpdateIn) (*OrganizationUserUpdateOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) OrganizationUserHandler {
	return OrganizationUserHandler{doer}
}

type OrganizationUserHandler struct {
	doer doer
}

func (h *OrganizationUserHandler) OrganizationUserAuthenticationMethodsList(ctx context.Context, organizationId string, memberUserId string) ([]AuthenticationMethodOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user/%s/authentication_methods", url.PathEscape(organizationId), url.PathEscape(memberUserId))
	b, err := h.doer.Do(ctx, "OrganizationUserAuthenticationMethodsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(organizationUserAuthenticationMethodsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AuthenticationMethods, nil
}
func (h *OrganizationUserHandler) OrganizationUserDelete(ctx context.Context, organizationId string, memberUserId string) error {
	path := fmt.Sprintf("/v1/organization/%s/user/%s", url.PathEscape(organizationId), url.PathEscape(memberUserId))
	_, err := h.doer.Do(ctx, "OrganizationUserDelete", "DELETE", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserGet(ctx context.Context, organizationId string, memberUserId string) (*OrganizationUserGetOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user/%s", url.PathEscape(organizationId), url.PathEscape(memberUserId))
	b, err := h.doer.Do(ctx, "OrganizationUserGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationUserGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationUserHandler) OrganizationUserInvitationAccept(ctx context.Context, organizationId string, userEmail string) error {
	path := fmt.Sprintf("/v1/organization/%s/invitation/%s", url.PathEscape(organizationId), url.PathEscape(userEmail))
	_, err := h.doer.Do(ctx, "OrganizationUserInvitationAccept", "POST", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserInvitationDelete(ctx context.Context, organizationId string, userEmail string) error {
	path := fmt.Sprintf("/v1/organization/%s/invitation/%s", url.PathEscape(organizationId), url.PathEscape(userEmail))
	_, err := h.doer.Do(ctx, "OrganizationUserInvitationDelete", "DELETE", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserInvitationsList(ctx context.Context, organizationId string) ([]InvitationOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/invitation", url.PathEscape(organizationId))
	b, err := h.doer.Do(ctx, "OrganizationUserInvitationsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(organizationUserInvitationsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Invitations, nil
}
func (h *OrganizationUserHandler) OrganizationUserInvite(ctx context.Context, organizationId string, in *OrganizationUserInviteIn) error {
	path := fmt.Sprintf("/v1/organization/%s/invitation", url.PathEscape(organizationId))
	_, err := h.doer.Do(ctx, "OrganizationUserInvite", "POST", path, in)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserList(ctx context.Context, organizationId string) ([]UserOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user", url.PathEscape(organizationId))
	b, err := h.doer.Do(ctx, "OrganizationUserList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(organizationUserListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Users, nil
}
func (h *OrganizationUserHandler) OrganizationUserPasswordReset(ctx context.Context, organizationId string, memberUserId string) error {
	path := fmt.Sprintf("/v1/organization/%s/user/%s/reset_password", url.PathEscape(organizationId), url.PathEscape(memberUserId))
	_, err := h.doer.Do(ctx, "OrganizationUserPasswordReset", "POST", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserRevokeToken(ctx context.Context, organizationId string, memberUserId string, tokenPrefix string) error {
	path := fmt.Sprintf("/v1/organization/%s/user/%s/access-token/%s", url.PathEscape(organizationId), url.PathEscape(memberUserId), url.PathEscape(tokenPrefix))
	_, err := h.doer.Do(ctx, "OrganizationUserRevokeToken", "DELETE", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserTokensList(ctx context.Context, organizationId string, memberUserId string) ([]TokenOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user/%s/access-tokens", url.PathEscape(organizationId), url.PathEscape(memberUserId))
	b, err := h.doer.Do(ctx, "OrganizationUserTokensList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(organizationUserTokensListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Tokens, nil
}
func (h *OrganizationUserHandler) OrganizationUserUpdate(ctx context.Context, organizationId string, memberUserId string, in *OrganizationUserUpdateIn) (*OrganizationUserUpdateOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user/%s", url.PathEscape(organizationId), url.PathEscape(memberUserId))
	b, err := h.doer.Do(ctx, "OrganizationUserUpdate", "PATCH", path, in)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationUserUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type AuthenticationMethodOut struct {
	IsEnabled2Fa     *bool      `json:"is_enabled_2fa,omitempty"`    // Verifies if 2FA is enabled for the user
	LastUsedTime     *time.Time `json:"last_used_time,omitempty"`    // Last activity time with the authentication method
	LocalProviderId  *string    `json:"local_provider_id,omitempty"` // Local authentication method provider resource ID
	MethodId         *string    `json:"method_id,omitempty"`         // User authentication method ID
	Name             *string    `json:"name,omitempty"`              // Name of the organization authentication method
	OrganizationId   *string    `json:"organization_id,omitempty"`   // Organization ID
	RemoteProviderId string     `json:"remote_provider_id"`          // Remote authentication method provider ID
	Type             *string    `json:"type,omitempty"`              // Type of the organization authentication method
	UserEmail        *string    `json:"user_email,omitempty"`        // User's email address for the authentication method
	UserId           *string    `json:"user_id,omitempty"`           // User ID
}
type InvitationOut struct {
	CreateTime time.Time `json:"create_time"` // Time of creating the invitation
	ExpiryTime time.Time `json:"expiry_time"` // By when the invitation is valid
	InvitedBy  string    `json:"invited_by"`  // Name of the invitation creator
	UserEmail  string    `json:"user_email"`  // User Email
}

// OrganizationUserGetOut OrganizationUserGetResponse
type OrganizationUserGetOut struct {
	IsSuperAdmin     bool        `json:"is_super_admin"`     // Super admin state of the organization user
	JoinTime         time.Time   `json:"join_time"`          // Join time
	LastActivityTime time.Time   `json:"last_activity_time"` // Last activity time
	UserId           string      `json:"user_id"`            // User ID
	UserInfo         UserInfoOut `json:"user_info"`          // OrganizationUserInfo
}

// OrganizationUserInviteIn OrganizationUserInviteRequestBody
type OrganizationUserInviteIn struct {
	UserEmail string `json:"user_email"` // User Email
}
type OrganizationUserStateType string

const (
	OrganizationUserStateTypeActive      OrganizationUserStateType = "active"
	OrganizationUserStateTypeDeactivated OrganizationUserStateType = "deactivated"
	OrganizationUserStateTypeDeleted     OrganizationUserStateType = "deleted"
)

func OrganizationUserStateTypeChoices() []string {
	return []string{"active", "deactivated", "deleted"}
}

// OrganizationUserUpdateIn OrganizationUserUpdateRequestBody
type OrganizationUserUpdateIn struct {
	City         *string                   `json:"city,omitempty"`
	Country      *string                   `json:"country,omitempty"`
	Department   *string                   `json:"department,omitempty"`
	IsSuperAdmin *bool                     `json:"is_super_admin,omitempty"` // Alters super admin state of the organization user
	JobTitle     *string                   `json:"job_title,omitempty"`      // Job Title
	RealName     *string                   `json:"real_name,omitempty"`      // Real Name
	State        OrganizationUserStateType `json:"state,omitempty"`          // State of the user in the organization
}

// OrganizationUserUpdateOut OrganizationUserUpdateResponse
type OrganizationUserUpdateOut struct {
	IsSuperAdmin     bool        `json:"is_super_admin"`     // Super admin state of the organization user
	JoinTime         time.Time   `json:"join_time"`          // Join time
	LastActivityTime time.Time   `json:"last_activity_time"` // Last activity time
	UserId           string      `json:"user_id"`            // User ID
	UserInfo         UserInfoOut `json:"user_info"`          // OrganizationUserInfo
}
type TokenOut struct {
	Description   string    `json:"description"`
	LastIp        string    `json:"last_ip"`         // Last-used IP
	LastUsedTime  time.Time `json:"last_used_time"`  // Last-used time
	LastUserAgent string    `json:"last_user_agent"` // Last-used user agent
	TokenPrefix   string    `json:"token_prefix"`    // Token prefix
}

// UserInfoOut OrganizationUserInfo
type UserInfoOut struct {
	City                   *string   `json:"city,omitempty"`
	Country                *string   `json:"country,omitempty"`
	CreateTime             time.Time `json:"create_time"` // Creation time
	Department             *string   `json:"department,omitempty"`
	IsApplicationUser      bool      `json:"is_application_user"`                // Is Application User
	JobTitle               *string   `json:"job_title,omitempty"`                // Job Title
	ManagedByScim          bool      `json:"managed_by_scim"`                    // Managed By Scim
	ManagingOrganizationId *string   `json:"managing_organization_id,omitempty"` // Managing Organization ID
	RealName               string    `json:"real_name"`                          // Real Name
	State                  string    `json:"state"`
	UserEmail              string    `json:"user_email"` // User Email
}
type UserOut struct {
	IsSuperAdmin     bool        `json:"is_super_admin"`     // Super admin state of the organization user
	JoinTime         time.Time   `json:"join_time"`          // Join time
	LastActivityTime time.Time   `json:"last_activity_time"` // Last activity time
	UserId           string      `json:"user_id"`            // User ID
	UserInfo         UserInfoOut `json:"user_info"`          // OrganizationUserInfo
}

// organizationUserAuthenticationMethodsListOut OrganizationUserAuthenticationMethodsListResponse
type organizationUserAuthenticationMethodsListOut struct {
	AuthenticationMethods []AuthenticationMethodOut `json:"authentication_methods"` // List of authentication methods for the organization user
}

// organizationUserInvitationsListOut OrganizationUserInvitationsListResponse
type organizationUserInvitationsListOut struct {
	Invitations []InvitationOut `json:"invitations"` // List of user invitations for the organization
}

// organizationUserListOut OrganizationUserListResponse
type organizationUserListOut struct {
	Users []UserOut `json:"users"` // List of users of the organization
}

// organizationUserTokensListOut OrganizationUserTokensListResponse
type organizationUserTokensListOut struct {
	Tokens []TokenOut `json:"tokens"` // List of user tokens accessible to the organization
}
