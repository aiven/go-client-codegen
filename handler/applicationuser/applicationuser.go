// Code generated by Aiven. DO NOT EDIT.

package applicationuser

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// ApplicationUserAccessTokenCreate create an application token
	// POST /v1/organization/{organization_id}/application-users/{user_id}/access-tokens
	// https://api.aiven.io/doc/#tag/Application_Users/operation/ApplicationUserAccessTokenCreate
	ApplicationUserAccessTokenCreate(ctx context.Context, organizationId string, userId string, in *ApplicationUserAccessTokenCreateIn) (*ApplicationUserAccessTokenCreateOut, error)

	// ApplicationUserAccessTokenDelete delete an application token
	// DELETE /v1/organization/{organization_id}/application-users/{user_id}/access-tokens/{token_prefix}
	// https://api.aiven.io/doc/#tag/Application_Users/operation/ApplicationUserAccessTokenDelete
	ApplicationUserAccessTokenDelete(ctx context.Context, organizationId string, userId string, tokenPrefix string) error

	// ApplicationUserAccessTokensList list application tokens
	// GET /v1/organization/{organization_id}/application-users/{user_id}/access-tokens
	// https://api.aiven.io/doc/#tag/Application_Users/operation/ApplicationUserAccessTokensList
	ApplicationUserAccessTokensList(ctx context.Context, organizationId string, userId string) ([]TokenOut, error)

	// ApplicationUserCreate create an application user
	// POST /v1/organization/{organization_id}/application-users
	// https://api.aiven.io/doc/#tag/Application_Users/operation/ApplicationUserCreate
	ApplicationUserCreate(ctx context.Context, organizationId string, in *ApplicationUserCreateIn) (*ApplicationUserCreateOut, error)

	// ApplicationUserDelete delete an application user
	// DELETE /v1/organization/{organization_id}/application-users/{user_id}
	// https://api.aiven.io/doc/#tag/Application_Users/operation/ApplicationUserDelete
	ApplicationUserDelete(ctx context.Context, organizationId string, userId string) error

	// ApplicationUserGet get an application user
	// GET /v1/organization/{organization_id}/application-users/{user_id}
	// https://api.aiven.io/doc/#tag/Application_Users/operation/ApplicationUserGet
	ApplicationUserGet(ctx context.Context, organizationId string, userId string) (*ApplicationUserGetOut, error)

	// ApplicationUsersList list application users
	// GET /v1/organization/{organization_id}/application-users
	// https://api.aiven.io/doc/#tag/Application_Users/operation/ApplicationUsersList
	ApplicationUsersList(ctx context.Context, organizationId string) ([]ApplicationUserOut, error)
}

func NewHandler(doer doer) ApplicationUserHandler {
	return ApplicationUserHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type ApplicationUserHandler struct {
	doer doer
}

func (h *ApplicationUserHandler) ApplicationUserAccessTokenCreate(ctx context.Context, organizationId string, userId string, in *ApplicationUserAccessTokenCreateIn) (*ApplicationUserAccessTokenCreateOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/application-users/%s/access-tokens", organizationId, userId)
	b, err := h.doer.Do(ctx, "ApplicationUserAccessTokenCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ApplicationUserAccessTokenCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *ApplicationUserHandler) ApplicationUserAccessTokenDelete(ctx context.Context, organizationId string, userId string, tokenPrefix string) error {
	path := fmt.Sprintf("/v1/organization/%s/application-users/%s/access-tokens/%s", organizationId, userId, tokenPrefix)
	_, err := h.doer.Do(ctx, "ApplicationUserAccessTokenDelete", "DELETE", path, nil)
	return err
}
func (h *ApplicationUserHandler) ApplicationUserAccessTokensList(ctx context.Context, organizationId string, userId string) ([]TokenOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/application-users/%s/access-tokens", organizationId, userId)
	b, err := h.doer.Do(ctx, "ApplicationUserAccessTokensList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(applicationUserAccessTokensListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Tokens, nil
}
func (h *ApplicationUserHandler) ApplicationUserCreate(ctx context.Context, organizationId string, in *ApplicationUserCreateIn) (*ApplicationUserCreateOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/application-users", organizationId)
	b, err := h.doer.Do(ctx, "ApplicationUserCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ApplicationUserCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *ApplicationUserHandler) ApplicationUserDelete(ctx context.Context, organizationId string, userId string) error {
	path := fmt.Sprintf("/v1/organization/%s/application-users/%s", organizationId, userId)
	_, err := h.doer.Do(ctx, "ApplicationUserDelete", "DELETE", path, nil)
	return err
}
func (h *ApplicationUserHandler) ApplicationUserGet(ctx context.Context, organizationId string, userId string) (*ApplicationUserGetOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/application-users/%s", organizationId, userId)
	b, err := h.doer.Do(ctx, "ApplicationUserGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ApplicationUserGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *ApplicationUserHandler) ApplicationUsersList(ctx context.Context, organizationId string) ([]ApplicationUserOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/application-users", organizationId)
	b, err := h.doer.Do(ctx, "ApplicationUsersList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(applicationUsersListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ApplicationUsers, nil
}

type ApplicationUserAccessTokenCreateIn struct {
	Description    string    `json:"description"`
	ExtendWhenUsed *bool     `json:"extend_when_used,omitempty"`
	MaxAgeSeconds  *int      `json:"max_age_seconds,omitempty"`
	Scopes         *[]string `json:"scopes,omitempty"`
}
type ApplicationUserAccessTokenCreateOut struct {
	FullToken   string `json:"full_token"`
	TokenPrefix string `json:"token_prefix"`
}
type ApplicationUserCreateIn struct {
	Name string `json:"name"`
}
type ApplicationUserCreateOut struct {
	Name      string `json:"name"`
	UserEmail string `json:"user_email"`
	UserId    string `json:"user_id"`
}
type ApplicationUserGetOut struct {
	Name      string `json:"name"`
	UserEmail string `json:"user_email"`
	UserId    string `json:"user_id"`
}
type ApplicationUserOut struct {
	Name      string `json:"name"`
	UserEmail string `json:"user_email"`
	UserId    string `json:"user_id"`
}
type TokenOut struct {
	CreateTime                 time.Time  `json:"create_time"`
	CreatedManually            bool       `json:"created_manually"`
	CurrentlyActive            bool       `json:"currently_active"`
	Description                string     `json:"description,omitempty"`
	ExpiryTime                 *time.Time `json:"expiry_time,omitempty"`
	ExtendWhenUsed             *bool      `json:"extend_when_used,omitempty"`
	LastIp                     string     `json:"last_ip,omitempty"`
	LastUsedTime               *time.Time `json:"last_used_time,omitempty"`
	LastUserAgent              string     `json:"last_user_agent,omitempty"`
	LastUserAgentHumanReadable string     `json:"last_user_agent_human_readable,omitempty"`
	MaxAgeSeconds              *int       `json:"max_age_seconds,omitempty"`
	Scopes                     []string   `json:"scopes,omitempty"`
	TokenPrefix                string     `json:"token_prefix"`
}
type applicationUserAccessTokensListOut struct {
	Tokens []TokenOut `json:"tokens"`
}
type applicationUsersListOut struct {
	ApplicationUsers []ApplicationUserOut `json:"application_users"`
}