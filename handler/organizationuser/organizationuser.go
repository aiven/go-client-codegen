// Code generated by Aiven. DO NOT EDIT.

package organizationuser

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// OrganizationUserAuthenticationMethodsList list authentication methods for a user in the organization
	// GET /organization/{organization_id}/user/{member_user_id}/authentication_methods
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserAuthenticationMethodsList
	OrganizationUserAuthenticationMethodsList(ctx context.Context, organizationId string, memberUserId string) ([]AuthenticationMethodOut, error)

	// OrganizationUserDelete remove a user from the organization
	// DELETE /organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserDelete
	OrganizationUserDelete(ctx context.Context, organizationId string, memberUserId string) error

	// OrganizationUserGet get details on a user of the organization
	// GET /organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserGet
	OrganizationUserGet(ctx context.Context, organizationId string, memberUserId string) (*OrganizationUserGetOut, error)

	// OrganizationUserInvitationAccept accept a user invitation to the organization
	// POST /organization/{organization_id}/invitation/{user_email}
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUserInvitationAccept
	OrganizationUserInvitationAccept(ctx context.Context, organizationId string, userEmail string, in *OrganizationUserInvitationAcceptIn) error

	// OrganizationUserInvitationDelete remove an invitation to the organization
	// DELETE /organization/{organization_id}/invitation/{user_email}
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUserInvitationDelete
	OrganizationUserInvitationDelete(ctx context.Context, organizationId string, userEmail string) error

	// OrganizationUserInvitationsList list user invitations to the organization
	// GET /organization/{organization_id}/invitation
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUserInvitationsList
	OrganizationUserInvitationsList(ctx context.Context, organizationId string) ([]InvitationOut, error)

	// OrganizationUserInvite invite a user to the organization
	// POST /organization/{organization_id}/invitation
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUserInvite
	OrganizationUserInvite(ctx context.Context, organizationId string, in *OrganizationUserInviteIn) error

	// OrganizationUserList list users of the organization
	// GET /organization/{organization_id}/user
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserList
	OrganizationUserList(ctx context.Context, organizationId string) ([]UserOut, error)

	// OrganizationUserPasswordReset reset the password of a managed user in the organization
	// POST /organization/{organization_id}/user/{member_user_id}/reset_password
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserPasswordReset
	OrganizationUserPasswordReset(ctx context.Context, organizationId string, memberUserId string) error

	// OrganizationUserRevokeToken revoke the token of a managed user in the organization
	// DELETE /organization/{organization_id}/user/{member_user_id}/access-token/{token_prefix}
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserRevokeToken
	OrganizationUserRevokeToken(ctx context.Context, organizationId string, memberUserId string, tokenPrefix string) error

	// OrganizationUserTokensList list tokens from an organization's member
	// GET /organization/{organization_id}/user/{member_user_id}/access-tokens
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserTokensList
	OrganizationUserTokensList(ctx context.Context, organizationId string, memberUserId string) ([]TokenOut, error)

	// OrganizationUserUpdate update details on a user of the organization
	// PATCH /organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationUserUpdate
	OrganizationUserUpdate(ctx context.Context, organizationId string, memberUserId string, in *OrganizationUserUpdateIn) (*OrganizationUserUpdateOut, error)
}

func NewHandler(doer doer) OrganizationUserHandler {
	return OrganizationUserHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type OrganizationUserHandler struct {
	doer doer
}

func (h *OrganizationUserHandler) OrganizationUserAuthenticationMethodsList(ctx context.Context, organizationId string, memberUserId string) ([]AuthenticationMethodOut, error) {
	path := fmt.Sprintf("/organization/%s/user/%s/authentication_methods", organizationId, memberUserId)
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
	path := fmt.Sprintf("/organization/%s/user/%s", organizationId, memberUserId)
	_, err := h.doer.Do(ctx, "OrganizationUserDelete", "DELETE", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserGet(ctx context.Context, organizationId string, memberUserId string) (*OrganizationUserGetOut, error) {
	path := fmt.Sprintf("/organization/%s/user/%s", organizationId, memberUserId)
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
func (h *OrganizationUserHandler) OrganizationUserInvitationAccept(ctx context.Context, organizationId string, userEmail string, in *OrganizationUserInvitationAcceptIn) error {
	path := fmt.Sprintf("/organization/%s/invitation/%s", organizationId, userEmail)
	_, err := h.doer.Do(ctx, "OrganizationUserInvitationAccept", "POST", path, in)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserInvitationDelete(ctx context.Context, organizationId string, userEmail string) error {
	path := fmt.Sprintf("/organization/%s/invitation/%s", organizationId, userEmail)
	_, err := h.doer.Do(ctx, "OrganizationUserInvitationDelete", "DELETE", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserInvitationsList(ctx context.Context, organizationId string) ([]InvitationOut, error) {
	path := fmt.Sprintf("/organization/%s/invitation", organizationId)
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
	path := fmt.Sprintf("/organization/%s/invitation", organizationId)
	_, err := h.doer.Do(ctx, "OrganizationUserInvite", "POST", path, in)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserList(ctx context.Context, organizationId string) ([]UserOut, error) {
	path := fmt.Sprintf("/organization/%s/user", organizationId)
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
	path := fmt.Sprintf("/organization/%s/user/%s/reset_password", organizationId, memberUserId)
	_, err := h.doer.Do(ctx, "OrganizationUserPasswordReset", "POST", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserRevokeToken(ctx context.Context, organizationId string, memberUserId string, tokenPrefix string) error {
	path := fmt.Sprintf("/organization/%s/user/%s/access-token/%s", organizationId, memberUserId, tokenPrefix)
	_, err := h.doer.Do(ctx, "OrganizationUserRevokeToken", "DELETE", path, nil)
	return err
}
func (h *OrganizationUserHandler) OrganizationUserTokensList(ctx context.Context, organizationId string, memberUserId string) ([]TokenOut, error) {
	path := fmt.Sprintf("/organization/%s/user/%s/access-tokens", organizationId, memberUserId)
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
	path := fmt.Sprintf("/organization/%s/user/%s", organizationId, memberUserId)
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

type ActionType string

const (
	ActionTypeAccept ActionType = "accept"
)

func ActionTypeChoices() []string {
	return []string{"accept"}
}

type AuthenticationMethodOut struct {
	IsEnabled2Fa     *bool      `json:"is_enabled_2fa,omitempty"`
	LastUsedTime     *time.Time `json:"last_used_time,omitempty"`
	LocalProviderId  string     `json:"local_provider_id,omitempty"`
	MethodId         string     `json:"method_id,omitempty"`
	Name             string     `json:"name,omitempty"`
	OrganizationId   string     `json:"organization_id,omitempty"`
	RemoteProviderId string     `json:"remote_provider_id"`
	Type             string     `json:"type,omitempty"`
	UserEmail        string     `json:"user_email,omitempty"`
	UserId           string     `json:"user_id,omitempty"`
}
type InvitationOut struct {
	CreateTime time.Time `json:"create_time"`
	ExpiryTime time.Time `json:"expiry_time"`
	InvitedBy  string    `json:"invited_by"`
	UserEmail  string    `json:"user_email"`
}
type OrganizationUserGetOut struct {
	IsSuperAdmin     bool        `json:"is_super_admin"`
	JoinTime         time.Time   `json:"join_time"`
	LastActivityTime time.Time   `json:"last_activity_time"`
	UserId           string      `json:"user_id"`
	UserInfo         UserInfoOut `json:"user_info"`
}
type OrganizationUserInvitationAcceptIn struct {
	Action ActionType `json:"action,omitempty"`
}
type OrganizationUserInviteIn struct {
	UserEmail string `json:"user_email"`
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

type OrganizationUserUpdateIn struct {
	City         string                    `json:"city,omitempty"`
	Country      string                    `json:"country,omitempty"`
	Department   string                    `json:"department,omitempty"`
	IsSuperAdmin *bool                     `json:"is_super_admin,omitempty"`
	JobTitle     string                    `json:"job_title,omitempty"`
	RealName     string                    `json:"real_name,omitempty"`
	State        OrganizationUserStateType `json:"state,omitempty"`
}
type OrganizationUserUpdateOut struct {
	IsSuperAdmin     bool        `json:"is_super_admin"`
	JoinTime         time.Time   `json:"join_time"`
	LastActivityTime time.Time   `json:"last_activity_time"`
	UserId           string      `json:"user_id"`
	UserInfo         UserInfoOut `json:"user_info"`
}
type TokenOut struct {
	Description   string    `json:"description"`
	LastIp        string    `json:"last_ip"`
	LastUsedTime  time.Time `json:"last_used_time"`
	LastUserAgent string    `json:"last_user_agent"`
	TokenPrefix   string    `json:"token_prefix"`
}
type UserInfoOut struct {
	City                   string    `json:"city,omitempty"`
	Country                string    `json:"country,omitempty"`
	CreateTime             time.Time `json:"create_time"`
	Department             string    `json:"department,omitempty"`
	IsApplicationUser      bool      `json:"is_application_user"`
	JobTitle               string    `json:"job_title,omitempty"`
	ManagedByScim          bool      `json:"managed_by_scim"`
	ManagingOrganizationId string    `json:"managing_organization_id,omitempty"`
	RealName               string    `json:"real_name"`
	State                  string    `json:"state"`
	UserEmail              string    `json:"user_email"`
}
type UserOut struct {
	IsSuperAdmin     bool        `json:"is_super_admin"`
	JoinTime         time.Time   `json:"join_time"`
	LastActivityTime time.Time   `json:"last_activity_time"`
	UserId           string      `json:"user_id"`
	UserInfo         UserInfoOut `json:"user_info"`
}
type organizationUserAuthenticationMethodsListOut struct {
	AuthenticationMethods []AuthenticationMethodOut `json:"authentication_methods"`
}
type organizationUserInvitationsListOut struct {
	Invitations []InvitationOut `json:"invitations"`
}
type organizationUserListOut struct {
	Users []UserOut `json:"users"`
}
type organizationUserTokensListOut struct {
	Tokens []TokenOut `json:"tokens"`
}
