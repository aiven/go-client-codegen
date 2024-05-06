// Code generated by Aiven. DO NOT EDIT.

package usergroup

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// UserGroupCreate create a group
	// POST /v1/organization/{organization_id}/user-groups
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupCreate
	UserGroupCreate(ctx context.Context, organizationId string, in *UserGroupCreateIn) (*UserGroupCreateOut, error)

	// UserGroupDelete delete a group
	// DELETE /v1/organization/{organization_id}/user-groups/{user_group_id}
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupDelete
	UserGroupDelete(ctx context.Context, organizationId string, userGroupId string) error

	// UserGroupGet retrieve a group
	// GET /v1/organization/{organization_id}/user-groups/{user_group_id}
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupGet
	UserGroupGet(ctx context.Context, organizationId string, userGroupId string) (*UserGroupGetOut, error)

	// UserGroupMemberList list group members
	// GET /v1/organization/{organization_id}/user-groups/{user_group_id}/members
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupMemberList
	UserGroupMemberList(ctx context.Context, organizationId string, userGroupId string) ([]MemberOut, error)

	// UserGroupMembersUpdate add or remove group members
	// PATCH /v1/organization/{organization_id}/user-groups/{user_group_id}/members
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupMembersUpdate
	UserGroupMembersUpdate(ctx context.Context, organizationId string, userGroupId string, in *UserGroupMembersUpdateIn) error

	// UserGroupUpdate update a group
	// PATCH /v1/organization/{organization_id}/user-groups/{user_group_id}
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupUpdate
	UserGroupUpdate(ctx context.Context, organizationId string, userGroupId string, in *UserGroupUpdateIn) (*UserGroupUpdateOut, error)

	// UserGroupsList list groups
	// GET /v1/organization/{organization_id}/user-groups
	// https://api.aiven.io/doc/#tag/Groups/operation/UserGroupsList
	UserGroupsList(ctx context.Context, organizationId string) ([]UserGroupOut, error)
}

func NewHandler(doer doer) UserGroupHandler {
	return UserGroupHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type UserGroupHandler struct {
	doer doer
}

func (h *UserGroupHandler) UserGroupCreate(ctx context.Context, organizationId string, in *UserGroupCreateIn) (*UserGroupCreateOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user-groups", organizationId)
	b, err := h.doer.Do(ctx, "UserGroupCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(UserGroupCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserGroupHandler) UserGroupDelete(ctx context.Context, organizationId string, userGroupId string) error {
	path := fmt.Sprintf("/v1/organization/%s/user-groups/%s", organizationId, userGroupId)
	_, err := h.doer.Do(ctx, "UserGroupDelete", "DELETE", path, nil)
	return err
}
func (h *UserGroupHandler) UserGroupGet(ctx context.Context, organizationId string, userGroupId string) (*UserGroupGetOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user-groups/%s", organizationId, userGroupId)
	b, err := h.doer.Do(ctx, "UserGroupGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(UserGroupGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserGroupHandler) UserGroupMemberList(ctx context.Context, organizationId string, userGroupId string) ([]MemberOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user-groups/%s/members", organizationId, userGroupId)
	b, err := h.doer.Do(ctx, "UserGroupMemberList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(userGroupMemberListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Members, nil
}
func (h *UserGroupHandler) UserGroupMembersUpdate(ctx context.Context, organizationId string, userGroupId string, in *UserGroupMembersUpdateIn) error {
	path := fmt.Sprintf("/v1/organization/%s/user-groups/%s/members", organizationId, userGroupId)
	_, err := h.doer.Do(ctx, "UserGroupMembersUpdate", "PATCH", path, in)
	return err
}
func (h *UserGroupHandler) UserGroupUpdate(ctx context.Context, organizationId string, userGroupId string, in *UserGroupUpdateIn) (*UserGroupUpdateOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user-groups/%s", organizationId, userGroupId)
	b, err := h.doer.Do(ctx, "UserGroupUpdate", "PATCH", path, in)
	if err != nil {
		return nil, err
	}
	out := new(UserGroupUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *UserGroupHandler) UserGroupsList(ctx context.Context, organizationId string) ([]UserGroupOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/user-groups", organizationId)
	b, err := h.doer.Do(ctx, "UserGroupsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(userGroupsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.UserGroups, nil
}

type MemberOut struct {
	LastActivityTime *time.Time  `json:"last_activity_time,omitempty"`
	UserId           string      `json:"user_id"`
	UserInfo         UserInfoOut `json:"user_info"`
}
type OperationType string

const (
	OperationTypeAddMembers    OperationType = "add_members"
	OperationTypeRemoveMembers OperationType = "remove_members"
)

func OperationTypeChoices() []string {
	return []string{"add_members", "remove_members"}
}

type UserGroupCreateIn struct {
	Description   string `json:"description"`
	UserGroupName string `json:"user_group_name"`
}
type UserGroupCreateOut struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	ManagedByScim bool      `json:"managed_by_scim"`
	UpdateTime    time.Time `json:"update_time"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
}
type UserGroupGetOut struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	ManagedByScim bool      `json:"managed_by_scim"`
	UpdateTime    time.Time `json:"update_time"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
}
type UserGroupMembersUpdateIn struct {
	MemberIds []string      `json:"member_ids"`
	Operation OperationType `json:"operation"`
}
type UserGroupOut struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	ManagedByScim bool      `json:"managed_by_scim"`
	MemberCount   int       `json:"member_count"`
	UpdateTime    time.Time `json:"update_time"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
}
type UserGroupUpdateIn struct {
	Description   string `json:"description,omitempty"`
	UserGroupName string `json:"user_group_name,omitempty"`
}
type UserGroupUpdateOut struct {
	CreateTime    time.Time `json:"create_time"`
	Description   string    `json:"description"`
	ManagedByScim bool      `json:"managed_by_scim"`
	UpdateTime    time.Time `json:"update_time"`
	UserGroupId   string    `json:"user_group_id"`
	UserGroupName string    `json:"user_group_name"`
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
type userGroupMemberListOut struct {
	Members []MemberOut `json:"members"`
}
type userGroupsListOut struct {
	UserGroups []UserGroupOut `json:"user_groups"`
}
