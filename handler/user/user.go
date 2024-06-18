// Code generated by Aiven. DO NOT EDIT.

package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// AccessTokenCreate create new access token
	// POST /v1/access_token
	// https://api.aiven.io/doc/#tag/Users/operation/AccessTokenCreate
	AccessTokenCreate(ctx context.Context, in *AccessTokenCreateIn) (*AccessTokenCreateOut, error)

	// AccessTokenList list all valid access tokens
	// GET /v1/access_token
	// https://api.aiven.io/doc/#tag/Users/operation/AccessTokenList
	AccessTokenList(ctx context.Context) ([]TokenOut, error)

	// AccessTokenRevoke revoke an access token
	// DELETE /v1/access_token/{token_prefix}
	// https://api.aiven.io/doc/#tag/Users/operation/AccessTokenRevoke
	AccessTokenRevoke(ctx context.Context, tokenPrefix string) error

	// AccessTokenUpdate update an existing access token
	// PUT /v1/access_token/{token_prefix}
	// https://api.aiven.io/doc/#tag/Users/operation/AccessTokenUpdate
	AccessTokenUpdate(ctx context.Context, tokenPrefix string, in *AccessTokenUpdateIn) (*AccessTokenUpdateOut, error)

	// CheckPasswordStrengthExistingUser check password strength for an existing user
	// POST /v1/me/password_strength
	// https://api.aiven.io/doc/#tag/Users/operation/CheckPasswordStrengthExistingUser
	CheckPasswordStrengthExistingUser(ctx context.Context, in *CheckPasswordStrengthExistingUserIn) (*CheckPasswordStrengthExistingUserOut, error)

	// CheckPasswordStrengthNewUser check password strength for a new user
	// POST /v1/user/password_strength
	// https://api.aiven.io/doc/#tag/Users/operation/CheckPasswordStrengthNewUser
	CheckPasswordStrengthNewUser(ctx context.Context, in *CheckPasswordStrengthNewUserIn) (*CheckPasswordStrengthNewUserOut, error)

	// OrganizationMemberGroupsList list user groups of the organization's member
	// GET /v1/organization/{organization_id}/user/{member_user_id}/user-groups
	// https://api.aiven.io/doc/#tag/Users/operation/OrganizationMemberGroupsList
	OrganizationMemberGroupsList(ctx context.Context, organizationId string, memberUserId string) ([]UserGroupOut, error)

	// TwoFactorAuthConfigure configure two-factor authentication
	// PUT /v1/me/2fa
	// https://api.aiven.io/doc/#tag/Users/operation/TwoFactorAuthConfigure
	TwoFactorAuthConfigure(ctx context.Context, in *TwoFactorAuthConfigureIn) (*TwoFactorAuthConfigureOut, error)

	// TwoFactorAuthConfigureOTP complete one-time password configuration
	// PUT /v1/me/2fa/otp
	// https://api.aiven.io/doc/#tag/Users/operation/TwoFactorAuthConfigureOTP
	TwoFactorAuthConfigureOTP(ctx context.Context, in *TwoFactorAuthConfigureOtpIn) (*TwoFactorAuthConfigureOtpOut, error)

	// UserAccountDelete delete user account
	// DELETE /v1/user/{user_id}
	// https://api.aiven.io/doc/#tag/Users/operation/UserAccountDelete
	UserAccountDelete(ctx context.Context, userId string) error

	// UserAccountInvitesAccept accept all invites for a single account
	// POST /v1/me/account/invites/accept
	// https://api.aiven.io/doc/#tag/Users/operation/UserAccountInvitesAccept
	UserAccountInvitesAccept(ctx context.Context, in *UserAccountInvitesAcceptIn) ([]AccountInviteOut, error)

	// UserAccountInvitesList list pending account invites
	// GET /v1/me/account/invites
	// https://api.aiven.io/doc/#tag/Users/operation/UserAccountInvitesList
	UserAccountInvitesList(ctx context.Context) ([]AccountInviteOut, error)

	// UserAuth authenticate user
	// POST /v1/userauth
	// https://api.aiven.io/doc/#tag/Users/operation/UserAuth
	UserAuth(ctx context.Context, in *UserAuthIn) (*UserAuthOut, error)

	// UserAuthLoginOptions get available login options
	// POST /v1/userauth/login_options
	// https://api.aiven.io/doc/#tag/Users/operation/UserAuthLoginOptions
	UserAuthLoginOptions(ctx context.Context, in *UserAuthLoginOptionsIn) (*UserAuthLoginOptionsOut, error)

	// UserAuthenticationMethodDelete delete linked authentication method, and revoke all associated access tokens
	// DELETE /v1/me/authentication_methods/{user_authentication_method_id}
	// https://api.aiven.io/doc/#tag/Users/operation/UserAuthenticationMethodDelete
	UserAuthenticationMethodDelete(ctx context.Context, userAuthenticationMethodId string) error

	// UserAuthenticationMethodsList list linked authentication methods
	// GET /v1/me/authentication_methods
	// https://api.aiven.io/doc/#tag/Users/operation/UserAuthenticationMethodsList
	UserAuthenticationMethodsList(ctx context.Context) ([]AuthenticationMethodOut, error)

	// UserExpireTokens expire all authorization tokens
	// POST /v1/me/expire_tokens
	// https://api.aiven.io/doc/#tag/Users/operation/UserExpireTokens
	UserExpireTokens(ctx context.Context) error

	// UserInfo get information for the current session's user
	// GET /v1/me
	// https://api.aiven.io/doc/#tag/Users/operation/UserInfo
	UserInfo(ctx context.Context) (*UserInfoOut, error)

	// UserLogout logout user, removing current authentication token
	// POST /v1/me/logout
	// https://api.aiven.io/doc/#tag/Users/operation/UserLogout
	UserLogout(ctx context.Context) error

	// UserPasswordChange change user password
	// PUT /v1/me/password
	// https://api.aiven.io/doc/#tag/Users/operation/UserPasswordChange
	UserPasswordChange(ctx context.Context, in *UserPasswordChangeIn) (string, error)

	// UserPasswordReset confirm user password reset
	// POST /v1/user/password_reset/{verification_code}
	// https://api.aiven.io/doc/#tag/Users/operation/UserPasswordReset
	UserPasswordReset(ctx context.Context, verificationCode string, in *UserPasswordResetIn) error

	// UserPasswordResetRequest request user password reset
	// POST /v1/user/password_reset_request
	// https://api.aiven.io/doc/#tag/Users/operation/UserPasswordResetRequest
	UserPasswordResetRequest(ctx context.Context, in *UserPasswordResetRequestIn) error

	// UserUpdate edit profile
	// PATCH /v1/me
	// https://api.aiven.io/doc/#tag/Users/operation/UserUpdate
	UserUpdate(ctx context.Context, in *UserUpdateIn) (*UserUpdateOut, error)

	// UserVerifyEmail confirm user email address
	// POST /v1/user/verify_email/{verification_code}
	// https://api.aiven.io/doc/#tag/Users/operation/UserVerifyEmail
	UserVerifyEmail(ctx context.Context, verificationCode string) (*UserVerifyEmailOut, error)

	// ValidateCreditCode validate campaign credit code
	// GET /v1/user/credit_code/{credit_code}
	// https://api.aiven.io/doc/#tag/Users/operation/ValidateCreditCode
	ValidateCreditCode(ctx context.Context, creditCode string) error

	// ValidateReferralCode validate referral_code code
	// GET /v1/me/referral/validation/{referral_code}
	// https://api.aiven.io/doc/#tag/Users/operation/ValidateReferralCode
	ValidateReferralCode(ctx context.Context, referralCode string) error
}

func NewHandler(doer doer) UserHandler {
	return UserHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type UserHandler struct {
	doer doer
}

func (h *UserHandler) AccessTokenCreate(ctx context.Context, in *AccessTokenCreateIn) (*AccessTokenCreateOut, error) {
	path := fmt.Sprintf("/v1/access_token")
	b, err := h.doer.Do(ctx, "AccessTokenCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(AccessTokenCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserHandler) AccessTokenList(ctx context.Context) ([]TokenOut, error) {
	path := fmt.Sprintf("/v1/access_token")
	b, err := h.doer.Do(ctx, "AccessTokenList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accessTokenListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Tokens, nil
}
func (h *UserHandler) AccessTokenRevoke(ctx context.Context, tokenPrefix string) error {
	path := fmt.Sprintf("/v1/access_token/%s", url.PathEscape(tokenPrefix))
	_, err := h.doer.Do(ctx, "AccessTokenRevoke", "DELETE", path, nil)
	return err
}
func (h *UserHandler) AccessTokenUpdate(ctx context.Context, tokenPrefix string, in *AccessTokenUpdateIn) (*AccessTokenUpdateOut, error) {
	path := fmt.Sprintf("/v1/access_token/%s", url.PathEscape(tokenPrefix))
	b, err := h.doer.Do(ctx, "AccessTokenUpdate", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(AccessTokenUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserHandler) CheckPasswordStrengthExistingUser(ctx context.Context, in *CheckPasswordStrengthExistingUserIn) (*CheckPasswordStrengthExistingUserOut, error) {
	path := fmt.Sprintf("/v1/me/password_strength")
	b, err := h.doer.Do(ctx, "CheckPasswordStrengthExistingUser", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(checkPasswordStrengthExistingUserOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.PasswordStrength, nil
}
func (h *UserHandler) CheckPasswordStrengthNewUser(ctx context.Context, in *CheckPasswordStrengthNewUserIn) (*CheckPasswordStrengthNewUserOut, error) {
	path := fmt.Sprintf("/v1/user/password_strength")
	b, err := h.doer.Do(ctx, "CheckPasswordStrengthNewUser", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(checkPasswordStrengthNewUserOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.PasswordStrength, nil
}
func (h *UserHandler) OrganizationMemberGroupsList(ctx context.Context, organizationId string, memberUserId string) ([]UserGroupOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user/%s/user-groups", url.PathEscape(organizationId), url.PathEscape(memberUserId))
	b, err := h.doer.Do(ctx, "OrganizationMemberGroupsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(organizationMemberGroupsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.UserGroups, nil
}
func (h *UserHandler) TwoFactorAuthConfigure(ctx context.Context, in *TwoFactorAuthConfigureIn) (*TwoFactorAuthConfigureOut, error) {
	path := fmt.Sprintf("/v1/me/2fa")
	b, err := h.doer.Do(ctx, "TwoFactorAuthConfigure", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(TwoFactorAuthConfigureOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserHandler) TwoFactorAuthConfigureOTP(ctx context.Context, in *TwoFactorAuthConfigureOtpIn) (*TwoFactorAuthConfigureOtpOut, error) {
	path := fmt.Sprintf("/v1/me/2fa/otp")
	b, err := h.doer.Do(ctx, "TwoFactorAuthConfigureOTP", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(TwoFactorAuthConfigureOtpOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserHandler) UserAccountDelete(ctx context.Context, userId string) error {
	path := fmt.Sprintf("/v1/user/%s", url.PathEscape(userId))
	_, err := h.doer.Do(ctx, "UserAccountDelete", "DELETE", path, nil)
	return err
}
func (h *UserHandler) UserAccountInvitesAccept(ctx context.Context, in *UserAccountInvitesAcceptIn) ([]AccountInviteOut, error) {
	path := fmt.Sprintf("/v1/me/account/invites/accept")
	b, err := h.doer.Do(ctx, "UserAccountInvitesAccept", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(userAccountInvitesAcceptOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AccountInvites, nil
}
func (h *UserHandler) UserAccountInvitesList(ctx context.Context) ([]AccountInviteOut, error) {
	path := fmt.Sprintf("/v1/me/account/invites")
	b, err := h.doer.Do(ctx, "UserAccountInvitesList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(userAccountInvitesListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AccountInvites, nil
}
func (h *UserHandler) UserAuth(ctx context.Context, in *UserAuthIn) (*UserAuthOut, error) {
	path := fmt.Sprintf("/v1/userauth")
	b, err := h.doer.Do(ctx, "UserAuth", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(UserAuthOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserHandler) UserAuthLoginOptions(ctx context.Context, in *UserAuthLoginOptionsIn) (*UserAuthLoginOptionsOut, error) {
	path := fmt.Sprintf("/v1/userauth/login_options")
	b, err := h.doer.Do(ctx, "UserAuthLoginOptions", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(UserAuthLoginOptionsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserHandler) UserAuthenticationMethodDelete(ctx context.Context, userAuthenticationMethodId string) error {
	path := fmt.Sprintf("/v1/me/authentication_methods/%s", url.PathEscape(userAuthenticationMethodId))
	_, err := h.doer.Do(ctx, "UserAuthenticationMethodDelete", "DELETE", path, nil)
	return err
}
func (h *UserHandler) UserAuthenticationMethodsList(ctx context.Context) ([]AuthenticationMethodOut, error) {
	path := fmt.Sprintf("/v1/me/authentication_methods")
	b, err := h.doer.Do(ctx, "UserAuthenticationMethodsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(userAuthenticationMethodsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AuthenticationMethods, nil
}
func (h *UserHandler) UserExpireTokens(ctx context.Context) error {
	path := fmt.Sprintf("/v1/me/expire_tokens")
	_, err := h.doer.Do(ctx, "UserExpireTokens", "POST", path, nil)
	return err
}
func (h *UserHandler) UserInfo(ctx context.Context) (*UserInfoOut, error) {
	path := fmt.Sprintf("/v1/me")
	b, err := h.doer.Do(ctx, "UserInfo", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(userInfoOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.User, nil
}
func (h *UserHandler) UserLogout(ctx context.Context) error {
	path := fmt.Sprintf("/v1/me/logout")
	_, err := h.doer.Do(ctx, "UserLogout", "POST", path, nil)
	return err
}
func (h *UserHandler) UserPasswordChange(ctx context.Context, in *UserPasswordChangeIn) (string, error) {
	path := fmt.Sprintf("/v1/me/password")
	b, err := h.doer.Do(ctx, "UserPasswordChange", "PUT", path, in)
	if err != nil {
		return "", err
	}
	out := new(userPasswordChangeOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.Token, nil
}
func (h *UserHandler) UserPasswordReset(ctx context.Context, verificationCode string, in *UserPasswordResetIn) error {
	path := fmt.Sprintf("/v1/user/password_reset/%s", url.PathEscape(verificationCode))
	_, err := h.doer.Do(ctx, "UserPasswordReset", "POST", path, in)
	return err
}
func (h *UserHandler) UserPasswordResetRequest(ctx context.Context, in *UserPasswordResetRequestIn) error {
	path := fmt.Sprintf("/v1/user/password_reset_request")
	_, err := h.doer.Do(ctx, "UserPasswordResetRequest", "POST", path, in)
	return err
}
func (h *UserHandler) UserUpdate(ctx context.Context, in *UserUpdateIn) (*UserUpdateOut, error) {
	path := fmt.Sprintf("/v1/me")
	b, err := h.doer.Do(ctx, "UserUpdate", "PATCH", path, in)
	if err != nil {
		return nil, err
	}
	out := new(userUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.User, nil
}
func (h *UserHandler) UserVerifyEmail(ctx context.Context, verificationCode string) (*UserVerifyEmailOut, error) {
	path := fmt.Sprintf("/v1/user/verify_email/%s", url.PathEscape(verificationCode))
	b, err := h.doer.Do(ctx, "UserVerifyEmail", "POST", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(userVerifyEmailOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.InviteDetails, nil
}
func (h *UserHandler) ValidateCreditCode(ctx context.Context, creditCode string) error {
	path := fmt.Sprintf("/v1/user/credit_code/%s", url.PathEscape(creditCode))
	_, err := h.doer.Do(ctx, "ValidateCreditCode", "GET", path, nil)
	return err
}
func (h *UserHandler) ValidateReferralCode(ctx context.Context, referralCode string) error {
	path := fmt.Sprintf("/v1/me/referral/validation/%s", url.PathEscape(referralCode))
	_, err := h.doer.Do(ctx, "ValidateReferralCode", "GET", path, nil)
	return err
}

type AccessTokenCreateIn struct {
	Description    string    `json:"description"`
	ExtendWhenUsed *bool     `json:"extend_when_used,omitempty"`
	MaxAgeSeconds  *float64  `json:"max_age_seconds,omitempty"`
	Scopes         *[]string `json:"scopes,omitempty"`
}
type AccessTokenCreateOut struct {
	CreateTime                 time.Time  `json:"create_time"`
	CreatedManually            bool       `json:"created_manually"`
	CurrentlyActive            *bool      `json:"currently_active,omitempty"`
	Description                *string    `json:"description,omitempty"`
	ExpiryTime                 *time.Time `json:"expiry_time,omitempty"`
	ExtendWhenUsed             bool       `json:"extend_when_used"`
	FullToken                  string     `json:"full_token"`
	LastIp                     *string    `json:"last_ip,omitempty"`
	LastUsedTime               *time.Time `json:"last_used_time,omitempty"`
	LastUserAgent              *string    `json:"last_user_agent,omitempty"`
	LastUserAgentHumanReadable *string    `json:"last_user_agent_human_readable,omitempty"`
	MaxAgeSeconds              float64    `json:"max_age_seconds"`
	Scopes                     []string   `json:"scopes,omitempty"`
	TokenPrefix                string     `json:"token_prefix"`
}
type AccessTokenUpdateIn struct {
	Description string `json:"description"`
}
type AccessTokenUpdateOut struct {
	CreateTime                 time.Time  `json:"create_time"`
	CreatedManually            *bool      `json:"created_manually,omitempty"`
	CurrentlyActive            *bool      `json:"currently_active,omitempty"`
	Description                *string    `json:"description,omitempty"`
	ExpiryTime                 *time.Time `json:"expiry_time,omitempty"`
	ExtendWhenUsed             bool       `json:"extend_when_used"`
	LastIp                     *string    `json:"last_ip,omitempty"`
	LastUsedTime               *time.Time `json:"last_used_time,omitempty"`
	LastUserAgent              *string    `json:"last_user_agent,omitempty"`
	LastUserAgentHumanReadable *string    `json:"last_user_agent_human_readable,omitempty"`
	MaxAgeSeconds              float64    `json:"max_age_seconds"`
	Scopes                     []string   `json:"scopes,omitempty"`
	TokenPrefix                string     `json:"token_prefix"`
}
type AccountInviteOut struct {
	AccountId          string    `json:"account_id"`
	AccountName        string    `json:"account_name"`
	CreateTime         time.Time `json:"create_time"`
	InvitedByUserEmail string    `json:"invited_by_user_email"`
	TeamId             string    `json:"team_id"`
	TeamName           string    `json:"team_name"`
	UserEmail          string    `json:"user_email"`
}
type ActionType string

const (
	ActionTypeAzureOauth  ActionType = "azure_oauth"
	ActionTypeGithubOauth ActionType = "github_oauth"
	ActionTypeGoogleOauth ActionType = "google_oauth"
	ActionTypeHasuraOauth ActionType = "hasura_oauth"
	ActionTypePassword    ActionType = "password"
	ActionTypeSaml        ActionType = "saml"
	ActionTypeSignup      ActionType = "signup"
)

func ActionTypeChoices() []string {
	return []string{"azure_oauth", "github_oauth", "google_oauth", "hasura_oauth", "password", "saml", "signup"}
}

type AnyType string

const (
	AnyTypeAdmin     AnyType = "admin"
	AnyTypeDeveloper AnyType = "developer"
	AnyTypeOperator  AnyType = "operator"
	AnyTypeReadOnly  AnyType = "read_only"
)

func AnyTypeChoices() []string {
	return []string{"admin", "developer", "operator", "read_only"}
}

type AuthenticationMethodOut struct {
	AuthenticationMethodAccountId string                        `json:"authentication_method_account_id"`
	CreateTime                    time.Time                     `json:"create_time"`
	CurrentlyActive               bool                          `json:"currently_active"`
	DeleteTime                    time.Time                     `json:"delete_time"`
	LastUsedTime                  time.Time                     `json:"last_used_time"`
	MethodId                      string                        `json:"method_id"`
	Name                          *string                       `json:"name,omitempty"`
	PublicRemoteIdentity          string                        `json:"public_remote_identity"`
	RemoteProviderId              string                        `json:"remote_provider_id"`
	State                         AuthenticationMethodStateType `json:"state"`
	UpdateTime                    time.Time                     `json:"update_time"`
	UserEmail                     string                        `json:"user_email"`
}
type AuthenticationMethodStateType string

const (
	AuthenticationMethodStateTypeActive  AuthenticationMethodStateType = "active"
	AuthenticationMethodStateTypeDeleted AuthenticationMethodStateType = "deleted"
)

func AuthenticationMethodStateTypeChoices() []string {
	return []string{"active", "deleted"}
}

type CheckPasswordStrengthExistingUserIn struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
type CheckPasswordStrengthExistingUserOut struct {
	IsAcceptable *bool  `json:"is_acceptable,omitempty"`
	Message      string `json:"message"`
	Score        int    `json:"score"`
}
type CheckPasswordStrengthNewUserIn struct {
	Email    *string `json:"email,omitempty"`
	Password string  `json:"password"`
	RealName *string `json:"real_name,omitempty"`
}
type CheckPasswordStrengthNewUserOut struct {
	IsAcceptable *bool  `json:"is_acceptable,omitempty"`
	Message      string `json:"message"`
	Score        int    `json:"score"`
}
type IntercomOut struct {
	AppId string `json:"app_id"`
	Hmac  string `json:"hmac"`
}
type InvitationOut struct {
	InviteCode        string    `json:"invite_code"`
	InviteTime        time.Time `json:"invite_time"`
	InvitingUserEmail string    `json:"inviting_user_email"`
	ProjectName       string    `json:"project_name"`
}
type MethodType string

const (
	MethodTypePost MethodType = "POST"
	MethodTypeGet  MethodType = "GET"
)

func MethodTypeChoices() []string {
	return []string{"POST", "GET"}
}

type ProjectMembershipOut struct {
	Any AnyType `json:"ANY,omitempty"`
}
type ProjectMembershipsOut struct {
	Any []string `json:"ANY,omitempty"`
}
type TokenOut struct {
	CreateTime                 time.Time  `json:"create_time"`
	CreatedManually            bool       `json:"created_manually"`
	CurrentlyActive            *bool      `json:"currently_active,omitempty"`
	Description                *string    `json:"description,omitempty"`
	ExpiryTime                 *time.Time `json:"expiry_time,omitempty"`
	ExtendWhenUsed             bool       `json:"extend_when_used"`
	LastIp                     *string    `json:"last_ip,omitempty"`
	LastUsedTime               *time.Time `json:"last_used_time,omitempty"`
	LastUserAgent              *string    `json:"last_user_agent,omitempty"`
	LastUserAgentHumanReadable *string    `json:"last_user_agent_human_readable,omitempty"`
	MaxAgeSeconds              float64    `json:"max_age_seconds"`
	Scopes                     []string   `json:"scopes,omitempty"`
	TokenPrefix                string     `json:"token_prefix"`
}
type TwoFactorAuthConfigureIn struct {
	Method   string `json:"method"`
	Password string `json:"password"`
}
type TwoFactorAuthConfigureOtpIn struct {
	Otp      string `json:"otp"`
	Password string `json:"password"`
	Uri      string `json:"uri"`
}
type TwoFactorAuthConfigureOtpOut struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}
type TwoFactorAuthConfigureOut struct {
	Method string  `json:"method"`
	Qrcode *string `json:"qrcode,omitempty"`
	Uri    *string `json:"uri,omitempty"`
}
type UserAccountInvitesAcceptIn struct {
	AccountId string  `json:"account_id"`
	TeamId    *string `json:"team_id,omitempty"`
}
type UserAuthIn struct {
	Email    string  `json:"email"`
	Otp      *string `json:"otp,omitempty"`
	Password string  `json:"password"`
}
type UserAuthLoginOptionsIn struct {
	Email *string `json:"email,omitempty"`
}
type UserAuthLoginOptionsOut struct {
	None        []map[string]any `json:"None,omitempty"`
	Action      ActionType       `json:"action"`
	Method      MethodType       `json:"method,omitempty"`
	Name        *string          `json:"name,omitempty"`
	RedirectUrl *string          `json:"redirect_url,omitempty"`
}
type UserAuthOut struct {
	ReturnUrl *string `json:"return_url,omitempty"`
	State     string  `json:"state"`
	Token     string  `json:"token"`
	UserEmail string  `json:"user_email"`
}
type UserGroupOut struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	ManagedByScim bool      `json:"managed_by_scim"`
	UpdateTime    time.Time `json:"update_time"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
}
type UserInfoOut struct {
	Auth                   []string               `json:"auth"`
	City                   *string                `json:"city,omitempty"`
	Country                *string                `json:"country,omitempty"`
	CreateTime             *time.Time             `json:"create_time,omitempty"`
	Department             *string                `json:"department,omitempty"`
	Features               map[string]any         `json:"features,omitempty"`
	Intercom               IntercomOut            `json:"intercom"`
	Invitations            []InvitationOut        `json:"invitations"`
	JobTitle               *string                `json:"job_title,omitempty"`
	ManagedByScim          *bool                  `json:"managed_by_scim,omitempty"`
	ManagingOrganizationId *string                `json:"managing_organization_id,omitempty"`
	ProjectMembership      ProjectMembershipOut   `json:"project_membership"`
	ProjectMemberships     *ProjectMembershipsOut `json:"project_memberships,omitempty"`
	Projects               []string               `json:"projects"`
	RealName               string                 `json:"real_name"`
	State                  string                 `json:"state"`
	TokenValidityBegin     *string                `json:"token_validity_begin,omitempty"`
	User                   string                 `json:"user"`
	UserId                 string                 `json:"user_id"`
}
type UserPasswordChangeIn struct {
	NewPassword string `json:"new_password"`
	Password    string `json:"password"`
}
type UserPasswordResetIn struct {
	NewPassword string `json:"new_password"`
}
type UserPasswordResetRequestIn struct {
	Email string `json:"email"`
}
type UserUpdateIn struct {
	City       *string `json:"city,omitempty"`
	Country    *string `json:"country,omitempty"`
	Department *string `json:"department,omitempty"`
	JobTitle   *string `json:"job_title,omitempty"`
	RealName   string  `json:"real_name"`
}
type UserUpdateOut struct {
	Auth                   []string               `json:"auth"`
	City                   *string                `json:"city,omitempty"`
	Country                *string                `json:"country,omitempty"`
	CreateTime             *time.Time             `json:"create_time,omitempty"`
	Department             *string                `json:"department,omitempty"`
	Features               map[string]any         `json:"features,omitempty"`
	Intercom               IntercomOut            `json:"intercom"`
	Invitations            []InvitationOut        `json:"invitations"`
	JobTitle               *string                `json:"job_title,omitempty"`
	ManagedByScim          *bool                  `json:"managed_by_scim,omitempty"`
	ManagingOrganizationId *string                `json:"managing_organization_id,omitempty"`
	ProjectMembership      ProjectMembershipOut   `json:"project_membership"`
	ProjectMemberships     *ProjectMembershipsOut `json:"project_memberships,omitempty"`
	Projects               []string               `json:"projects"`
	RealName               string                 `json:"real_name"`
	State                  string                 `json:"state"`
	TokenValidityBegin     *string                `json:"token_validity_begin,omitempty"`
	User                   string                 `json:"user"`
	UserId                 string                 `json:"user_id"`
}
type UserVerifyEmailOut struct {
	UserEmail string `json:"user_email"`
}
type accessTokenListOut struct {
	Tokens []TokenOut `json:"tokens"`
}
type checkPasswordStrengthExistingUserOut struct {
	PasswordStrength CheckPasswordStrengthExistingUserOut `json:"password_strength"`
}
type checkPasswordStrengthNewUserOut struct {
	PasswordStrength CheckPasswordStrengthNewUserOut `json:"password_strength"`
}
type organizationMemberGroupsListOut struct {
	UserGroups []UserGroupOut `json:"user_groups"`
}
type userAccountInvitesAcceptOut struct {
	AccountInvites []AccountInviteOut `json:"account_invites"`
}
type userAccountInvitesListOut struct {
	AccountInvites []AccountInviteOut `json:"account_invites"`
}
type userAuthenticationMethodsListOut struct {
	AuthenticationMethods []AuthenticationMethodOut `json:"authentication_methods"`
}
type userInfoOut struct {
	User UserInfoOut `json:"user"`
}
type userPasswordChangeOut struct {
	Token string `json:"token"`
}
type userUpdateOut struct {
	User UserUpdateOut `json:"user"`
}
type userVerifyEmailOut struct {
	InviteDetails UserVerifyEmailOut `json:"invite_details"`
}
