// Code generated by Aiven. DO NOT EDIT.

package organization

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// AuthenticationConfigGet retrieve authentication configuration
	// OrganizationAuthenticationConfigGet GET /organization/{organization_id}/config/authentication
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationAuthenticationConfigGet
	AuthenticationConfigGet(ctx context.Context, id string) (*AuthenticationConfigGetOut, error)

	// AuthenticationConfigUpdate update authentication configuration
	// OrganizationAuthenticationConfigUpdate PATCH /organization/{organization_id}/config/authentication
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationAuthenticationConfigUpdate
	AuthenticationConfigUpdate(ctx context.Context, id string, in *AuthenticationConfigUpdateIn) (*AuthenticationConfigUpdateOut, error)

	// Get get information about an organization
	// OrganizationGet GET /organization/{organization_id}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationGet
	Get(ctx context.Context, id string) (*GetOut, error)

	// MemberGroupsList list user groups of the organization's member
	// OrganizationMemberGroupsList GET /organization/{organization_id}/user/{member_user_id}/user-groups
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationMemberGroupsList
	MemberGroupsList(ctx context.Context, id string, memberUserId string) ([]UserGroup, error)

	// ProjectsList list projects under the organization
	// OrganizationProjectsList GET /organization/{organization_id}/projects
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationProjectsList
	ProjectsList(ctx context.Context, id string) (*ProjectsListOut, error)

	// Update update organization's details
	// OrganizationUpdate PATCH /organization/{organization_id}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUpdate
	Update(ctx context.Context, id string, in *UpdateIn) (*UpdateOut, error)

	// UserAuthenticationMethodsList list authentication methods for a user in the organization
	// OrganizationUserAuthenticationMethodsList GET /organization/{organization_id}/user/{member_user_id}/authentication_methods
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserAuthenticationMethodsList
	UserAuthenticationMethodsList(ctx context.Context, id string, memberUserId string) ([]AuthenticationMethod, error)

	// UserDelete remove a user from the organization
	// OrganizationUserDelete DELETE /organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserDelete
	UserDelete(ctx context.Context, id string, memberUserId string) error

	// UserGet get details on a user of the organization
	// OrganizationUserGet GET /organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserGet
	UserGet(ctx context.Context, id string, memberUserId string) (*UserGetOut, error)

	// UserInvitationAccept accept a user invitation to the organization
	// OrganizationUserInvitationAccept POST /organization/{organization_id}/invitation/{user_email}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserInvitationAccept
	UserInvitationAccept(ctx context.Context, id string, userEmail string, in *UserInvitationAcceptIn) error

	// UserInvitationDelete remove an invitation to the organization
	// OrganizationUserInvitationDelete DELETE /organization/{organization_id}/invitation/{user_email}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserInvitationDelete
	UserInvitationDelete(ctx context.Context, id string, userEmail string) error

	// UserInvitationsList list user invitations to the organization
	// OrganizationUserInvitationsList GET /organization/{organization_id}/invitation
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserInvitationsList
	UserInvitationsList(ctx context.Context, id string) ([]Invitation, error)

	// UserInvite invite a user to the organization
	// OrganizationUserInvite POST /organization/{organization_id}/invitation
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserInvite
	UserInvite(ctx context.Context, id string, in *UserInviteIn) error

	// UserList list users of the organization
	// OrganizationUserList GET /organization/{organization_id}/user
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserList
	UserList(ctx context.Context, id string) ([]User, error)

	// UserOrganizationCreate create an organization
	// UserOrganizationCreate POST /organizations
	// https://api.aiven.io/doc/#tag/Organization/operation/UserOrganizationCreate
	UserOrganizationCreate(ctx context.Context, in *UserOrganizationCreateIn) (*UserOrganizationCreateOut, error)

	// UserOrganizationsList list organizations the user belongs to
	// UserOrganizationsList GET /organizations
	// https://api.aiven.io/doc/#tag/Organization/operation/UserOrganizationsList
	UserOrganizationsList(ctx context.Context) ([]Organization, error)

	// UserPasswordReset reset the password of a managed user in the organization
	// OrganizationUserPasswordReset POST /organization/{organization_id}/user/{member_user_id}/reset_password
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserPasswordReset
	UserPasswordReset(ctx context.Context, id string, memberUserId string) error

	// UserRevokeToken revoke the token of a managed user in the organization
	// OrganizationUserRevokeToken DELETE /organization/{organization_id}/user/{member_user_id}/access-token/{token_prefix}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserRevokeToken
	UserRevokeToken(ctx context.Context, id string, memberUserId string, tokenPrefix string) error

	// UserSet add or modify a user of the organization
	// OrganizationUserSet PUT /organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserSet
	UserSet(ctx context.Context, id string, memberUserId string) (*UserSetOut, error)

	// UserTokensList list tokens from an organization's member
	// OrganizationUserTokensList GET /organization/{organization_id}/user/{member_user_id}/access-tokens
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserTokensList
	UserTokensList(ctx context.Context, id string, memberUserId string) ([]Token, error)

	// UserUpdate update details on a user of the organization
	// OrganizationUserUpdate PATCH /organization/{organization_id}/user/{member_user_id}
	// https://api.aiven.io/doc/#tag/Organization/operation/OrganizationUserUpdate
	UserUpdate(ctx context.Context, id string, memberUserId string, in *UserUpdateIn) (*UserUpdateOut, error)
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

func (h *handler) AuthenticationConfigGet(ctx context.Context, id string) (*AuthenticationConfigGetOut, error) {
	path := fmt.Sprintf("/organization/%s/config/authentication", id)
	b, err := h.doer.Do(ctx, "OrganizationAuthenticationConfigGet", "GET", path, nil)
	out := new(AuthenticationConfigGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) AuthenticationConfigUpdate(ctx context.Context, id string, in *AuthenticationConfigUpdateIn) (*AuthenticationConfigUpdateOut, error) {
	path := fmt.Sprintf("/organization/%s/config/authentication", id)
	b, err := h.doer.Do(ctx, "OrganizationAuthenticationConfigUpdate", "PATCH", path, in)
	out := new(AuthenticationConfigUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) Get(ctx context.Context, id string) (*GetOut, error) {
	path := fmt.Sprintf("/organization/%s", id)
	b, err := h.doer.Do(ctx, "OrganizationGet", "GET", path, nil)
	out := new(GetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) MemberGroupsList(ctx context.Context, id string, memberUserId string) ([]UserGroup, error) {
	path := fmt.Sprintf("/organization/%s/user/%s/user-groups", id, memberUserId)
	b, err := h.doer.Do(ctx, "OrganizationMemberGroupsList", "GET", path, nil)
	out := new(MemberGroupsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.UserGroups, nil
}
func (h *handler) ProjectsList(ctx context.Context, id string) (*ProjectsListOut, error) {
	path := fmt.Sprintf("/organization/%s/projects", id)
	b, err := h.doer.Do(ctx, "OrganizationProjectsList", "GET", path, nil)
	out := new(ProjectsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) Update(ctx context.Context, id string, in *UpdateIn) (*UpdateOut, error) {
	path := fmt.Sprintf("/organization/%s", id)
	b, err := h.doer.Do(ctx, "OrganizationUpdate", "PATCH", path, in)
	out := new(UpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) UserAuthenticationMethodsList(ctx context.Context, id string, memberUserId string) ([]AuthenticationMethod, error) {
	path := fmt.Sprintf("/organization/%s/user/%s/authentication_methods", id, memberUserId)
	b, err := h.doer.Do(ctx, "OrganizationUserAuthenticationMethodsList", "GET", path, nil)
	out := new(UserAuthenticationMethodsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AuthenticationMethods, nil
}
func (h *handler) UserDelete(ctx context.Context, id string, memberUserId string) error {
	path := fmt.Sprintf("/organization/%s/user/%s", id, memberUserId)
	_, err := h.doer.Do(ctx, "OrganizationUserDelete", "DELETE", path, nil)
	return err
}
func (h *handler) UserGet(ctx context.Context, id string, memberUserId string) (*UserGetOut, error) {
	path := fmt.Sprintf("/organization/%s/user/%s", id, memberUserId)
	b, err := h.doer.Do(ctx, "OrganizationUserGet", "GET", path, nil)
	out := new(UserGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) UserInvitationAccept(ctx context.Context, id string, userEmail string, in *UserInvitationAcceptIn) error {
	path := fmt.Sprintf("/organization/%s/invitation/%s", id, userEmail)
	_, err := h.doer.Do(ctx, "OrganizationUserInvitationAccept", "POST", path, in)
	return err
}
func (h *handler) UserInvitationDelete(ctx context.Context, id string, userEmail string) error {
	path := fmt.Sprintf("/organization/%s/invitation/%s", id, userEmail)
	_, err := h.doer.Do(ctx, "OrganizationUserInvitationDelete", "DELETE", path, nil)
	return err
}
func (h *handler) UserInvitationsList(ctx context.Context, id string) ([]Invitation, error) {
	path := fmt.Sprintf("/organization/%s/invitation", id)
	b, err := h.doer.Do(ctx, "OrganizationUserInvitationsList", "GET", path, nil)
	out := new(UserInvitationsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Invitations, nil
}
func (h *handler) UserInvite(ctx context.Context, id string, in *UserInviteIn) error {
	path := fmt.Sprintf("/organization/%s/invitation", id)
	_, err := h.doer.Do(ctx, "OrganizationUserInvite", "POST", path, in)
	return err
}
func (h *handler) UserList(ctx context.Context, id string) ([]User, error) {
	path := fmt.Sprintf("/organization/%s/user", id)
	b, err := h.doer.Do(ctx, "OrganizationUserList", "GET", path, nil)
	out := new(UserListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Users, nil
}
func (h *handler) UserOrganizationCreate(ctx context.Context, in *UserOrganizationCreateIn) (*UserOrganizationCreateOut, error) {
	path := fmt.Sprintf("/organizations")
	b, err := h.doer.Do(ctx, "UserOrganizationCreate", "POST", path, in)
	out := new(UserOrganizationCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) UserOrganizationsList(ctx context.Context) ([]Organization, error) {
	path := fmt.Sprintf("/organizations")
	b, err := h.doer.Do(ctx, "UserOrganizationsList", "GET", path, nil)
	out := new(UserOrganizationsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Organizations, nil
}
func (h *handler) UserPasswordReset(ctx context.Context, id string, memberUserId string) error {
	path := fmt.Sprintf("/organization/%s/user/%s/reset_password", id, memberUserId)
	_, err := h.doer.Do(ctx, "OrganizationUserPasswordReset", "POST", path, nil)
	return err
}
func (h *handler) UserRevokeToken(ctx context.Context, id string, memberUserId string, tokenPrefix string) error {
	path := fmt.Sprintf("/organization/%s/user/%s/access-token/%s", id, memberUserId, tokenPrefix)
	_, err := h.doer.Do(ctx, "OrganizationUserRevokeToken", "DELETE", path, nil)
	return err
}
func (h *handler) UserSet(ctx context.Context, id string, memberUserId string) (*UserSetOut, error) {
	path := fmt.Sprintf("/organization/%s/user/%s", id, memberUserId)
	b, err := h.doer.Do(ctx, "OrganizationUserSet", "PUT", path, nil)
	out := new(UserSetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *handler) UserTokensList(ctx context.Context, id string, memberUserId string) ([]Token, error) {
	path := fmt.Sprintf("/organization/%s/user/%s/access-tokens", id, memberUserId)
	b, err := h.doer.Do(ctx, "OrganizationUserTokensList", "GET", path, nil)
	out := new(UserTokensListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Tokens, nil
}
func (h *handler) UserUpdate(ctx context.Context, id string, memberUserId string, in *UserUpdateIn) (*UserUpdateOut, error) {
	path := fmt.Sprintf("/organization/%s/user/%s", id, memberUserId)
	b, err := h.doer.Do(ctx, "OrganizationUserUpdate", "PATCH", path, in)
	out := new(UserUpdateOut)
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

type AuthenticationConfigGetOut struct {
	OauthEnabled        *bool `json:"oauth_enabled,omitempty"`
	PasswordAuthEnabled *bool `json:"password_auth_enabled,omitempty"`
	SamlEnabled         *bool `json:"saml_enabled,omitempty"`
	TwoFactorRequired   *bool `json:"two_factor_required,omitempty"`
}
type AuthenticationConfigUpdateIn struct {
	OauthEnabled        *bool `json:"oauth_enabled,omitempty"`
	PasswordAuthEnabled *bool `json:"password_auth_enabled,omitempty"`
	SamlEnabled         *bool `json:"saml_enabled,omitempty"`
	TwoFactorRequired   *bool `json:"two_factor_required,omitempty"`
}
type AuthenticationConfigUpdateOut struct {
	OauthEnabled        *bool `json:"oauth_enabled,omitempty"`
	PasswordAuthEnabled *bool `json:"password_auth_enabled,omitempty"`
	SamlEnabled         *bool `json:"saml_enabled,omitempty"`
	TwoFactorRequired   *bool `json:"two_factor_required,omitempty"`
}
type AuthenticationMethod struct {
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
type BillingCurrencyType string

const (
	BillingCurrencyTypeAud BillingCurrencyType = "AUD"
	BillingCurrencyTypeCad BillingCurrencyType = "CAD"
	BillingCurrencyTypeChf BillingCurrencyType = "CHF"
	BillingCurrencyTypeDkk BillingCurrencyType = "DKK"
	BillingCurrencyTypeEur BillingCurrencyType = "EUR"
	BillingCurrencyTypeGbp BillingCurrencyType = "GBP"
	BillingCurrencyTypeJpy BillingCurrencyType = "JPY"
	BillingCurrencyTypeNok BillingCurrencyType = "NOK"
	BillingCurrencyTypeNzd BillingCurrencyType = "NZD"
	BillingCurrencyTypeSek BillingCurrencyType = "SEK"
	BillingCurrencyTypeSgd BillingCurrencyType = "SGD"
	BillingCurrencyTypeUsd BillingCurrencyType = "USD"
)

type BillingEmail struct {
	Email string `json:"email"`
}
type CardInfo struct {
	Brand       string `json:"brand"`
	CardId      string `json:"card_id"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	ExpMonth    int    `json:"exp_month"`
	ExpYear     int    `json:"exp_year"`
	Last4       string `json:"last4"`
	Name        string `json:"name"`
	UserEmail   string `json:"user_email"`
}
type Elasticsearch struct {
	EolDate string `json:"eol_date"`
	Version string `json:"version"`
}
type EndOfLifeExtension struct {
	Elasticsearch *Elasticsearch `json:"elasticsearch,omitempty"`
}
type GetOut struct {
	AccountId        string    `json:"account_id"`
	CreateTime       time.Time `json:"create_time"`
	OrganizationId   string    `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	Tier             TierType  `json:"tier"`
	UpdateTime       time.Time `json:"update_time"`
}
type Info struct {
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
type Invitation struct {
	CreateTime time.Time `json:"create_time"`
	ExpiryTime time.Time `json:"expiry_time"`
	InvitedBy  string    `json:"invited_by"`
	UserEmail  string    `json:"user_email"`
}
type MemberGroupsListOut struct {
	UserGroups []UserGroup `json:"user_groups"`
}
type Organization struct {
	AccountId        string    `json:"account_id"`
	CreateTime       time.Time `json:"create_time"`
	OrganizationId   string    `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	Tier             TierType  `json:"tier"`
	UpdateTime       time.Time `json:"update_time"`
}
type Project struct {
	AccountId             string              `json:"account_id"`
	AccountName           string              `json:"account_name,omitempty"`
	AddressLines          []string            `json:"address_lines,omitempty"`
	AvailableCredits      string              `json:"available_credits,omitempty"`
	BillingAddress        string              `json:"billing_address"`
	BillingCurrency       BillingCurrencyType `json:"billing_currency,omitempty"`
	BillingEmails         []BillingEmail      `json:"billing_emails"`
	BillingExtraText      string              `json:"billing_extra_text,omitempty"`
	BillingGroupId        string              `json:"billing_group_id"`
	BillingGroupName      string              `json:"billing_group_name"`
	CardInfo              *CardInfo           `json:"card_info,omitempty"`
	City                  string              `json:"city,omitempty"`
	Company               string              `json:"company,omitempty"`
	Country               string              `json:"country"`
	CountryCode           string              `json:"country_code"`
	DefaultCloud          string              `json:"default_cloud"`
	EndOfLifeExtension    *EndOfLifeExtension `json:"end_of_life_extension,omitempty"`
	EstimatedBalance      string              `json:"estimated_balance"`
	EstimatedBalanceLocal string              `json:"estimated_balance_local,omitempty"`
	Features              map[string]any      `json:"features,omitempty"`
	ProjectName           string              `json:"project_name"`
	OrganizationId        string              `json:"organization_id"`
	PaymentMethod         string              `json:"payment_method"`
	State                 string              `json:"state,omitempty"`
	Tags                  map[string]string   `json:"tags,omitempty"`
	TechEmails            []TechEmail         `json:"tech_emails,omitempty"`
	TenantId              string              `json:"tenant_id,omitempty"`
	TrialExpirationTime   *time.Time          `json:"trial_expiration_time,omitempty"`
	VatId                 string              `json:"vat_id"`
	ZipCode               string              `json:"zip_code,omitempty"`
}
type ProjectsListOut struct {
	Projects          []Project `json:"projects"`
	TotalProjectCount *int      `json:"total_project_count,omitempty"`
}
type TechEmail struct {
	Email string `json:"email"`
}
type TierType string

const (
	TierTypeBusiness TierType = "business"
	TierTypePersonal TierType = "personal"
)

type Token struct {
	Description   string    `json:"description"`
	LastIp        string    `json:"last_ip"`
	LastUsedTime  time.Time `json:"last_used_time"`
	LastUserAgent string    `json:"last_user_agent"`
	TokenPrefix   string    `json:"token_prefix"`
}
type UpdateIn struct {
	Name string   `json:"name,omitempty"`
	Tier TierType `json:"tier,omitempty"`
}
type UpdateOut struct {
	AccountId        string    `json:"account_id"`
	CreateTime       time.Time `json:"create_time"`
	OrganizationId   string    `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	Tier             TierType  `json:"tier"`
	UpdateTime       time.Time `json:"update_time"`
}
type User struct {
	UserId           string    `json:"user_id"`
	UserInfo         *Info     `json:"user_info,omitempty"`
	IsSuperAdmin     bool      `json:"is_super_admin"`
	JoinTime         time.Time `json:"join_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
}
type UserAuthenticationMethodsListOut struct {
	AuthenticationMethods []AuthenticationMethod `json:"authentication_methods"`
}
type UserGetOut struct {
	IsSuperAdmin     bool      `json:"is_super_admin"`
	JoinTime         time.Time `json:"join_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
	UserId           string    `json:"user_id"`
	UserInfo         *UserInfo `json:"user_info,omitempty"`
}
type UserGroup struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
	UpdateTime    time.Time `json:"update_time"`
}
type UserInfo struct {
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
type UserInvitationAcceptIn struct {
	Action ActionType `json:"action,omitempty"`
}
type UserInvitationsListOut struct {
	Invitations []Invitation `json:"invitations"`
}
type UserInviteIn struct {
	UserEmail string `json:"user_email"`
}
type UserListOut struct {
	Users []User `json:"users"`
}
type UserOrganizationCreateIn struct {
	OrganizationName      string   `json:"organization_name"`
	PrimaryBillingGroupId string   `json:"primary_billing_group_id,omitempty"`
	Tier                  TierType `json:"tier"`
}
type UserOrganizationCreateOut struct {
	AccountId        string    `json:"account_id"`
	CreateTime       time.Time `json:"create_time"`
	OrganizationId   string    `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	Tier             TierType  `json:"tier"`
	UpdateTime       time.Time `json:"update_time"`
}
type UserOrganizationsListOut struct {
	Organizations []Organization `json:"organizations"`
}
type UserSetOut struct {
	IsSuperAdmin     bool      `json:"is_super_admin"`
	JoinTime         time.Time `json:"join_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
	UserId           string    `json:"user_id"`
	UserInfo         *UserInfo `json:"user_info,omitempty"`
}
type UserTokensListOut struct {
	Tokens []Token `json:"tokens"`
}
type UserUpdateIn struct {
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	Department   string `json:"department,omitempty"`
	IsSuperAdmin *bool  `json:"is_super_admin,omitempty"`
	JobTitle     string `json:"job_title,omitempty"`
	RealName     string `json:"real_name,omitempty"`
}
type UserUpdateOut struct {
	IsSuperAdmin     bool      `json:"is_super_admin"`
	JoinTime         time.Time `json:"join_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
	UserId           string    `json:"user_id"`
	UserInfo         *UserInfo `json:"user_info,omitempty"`
}
