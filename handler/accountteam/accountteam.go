// Code generated by Aiven. DO NOT EDIT.

package accountteam

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// AccountTeamDelete delete a team
	// DELETE /account/{account_id}/team/{team_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamDelete
	AccountTeamDelete(ctx context.Context, accountId string, teamId string) error

	// AccountTeamGet get details for a single team
	// GET /account/{account_id}/team/{team_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamGet
	AccountTeamGet(ctx context.Context, accountId string, teamId string) (*AccountTeamGetOut, error)

	// AccountTeamInvitesList list pending invites
	// GET /account/{account_id}/team/{team_id}/invites
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamInvitesList
	AccountTeamInvitesList(ctx context.Context, accountId string, teamId string) ([]AccountInviteOut, error)

	// AccountTeamList list teams belonging to an account
	// GET /account/{account_id}/teams
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamList
	AccountTeamList(ctx context.Context, accountId string) ([]TeamOut, error)

	// AccountTeamProjectAssociate associate team to a project
	// POST /account/{account_id}/team/{team_id}/project/{project}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamProjectAssociate
	AccountTeamProjectAssociate(ctx context.Context, accountId string, teamId string, project string, in *AccountTeamProjectAssociateIn) error

	// AccountTeamProjectDisassociate disassociate team from a project
	// DELETE /account/{account_id}/team/{team_id}/project/{project}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamProjectDisassociate
	AccountTeamProjectDisassociate(ctx context.Context, accountId string, teamId string, project string) error

	// AccountTeamUpdate update team details
	// PUT /account/{account_id}/team/{team_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountTeamUpdate
	AccountTeamUpdate(ctx context.Context, accountId string, teamId string, in *AccountTeamUpdateIn) (*AccountTeamUpdateOut, error)
}

func NewHandler(doer doer) AccountTeamHandler {
	return AccountTeamHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type AccountTeamHandler struct {
	doer doer
}

func (h *AccountTeamHandler) AccountTeamDelete(ctx context.Context, accountId string, teamId string) error {
	path := fmt.Sprintf("/account/%s/team/%s", accountId, teamId)
	_, err := h.doer.Do(ctx, "AccountTeamDelete", "DELETE", path, nil)
	return err
}
func (h *AccountTeamHandler) AccountTeamGet(ctx context.Context, accountId string, teamId string) (*AccountTeamGetOut, error) {
	path := fmt.Sprintf("/account/%s/team/%s", accountId, teamId)
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
	path := fmt.Sprintf("/account/%s/team/%s/invites", accountId, teamId)
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
	path := fmt.Sprintf("/account/%s/teams", accountId)
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
	path := fmt.Sprintf("/account/%s/team/%s/project/%s", accountId, teamId, project)
	_, err := h.doer.Do(ctx, "AccountTeamProjectAssociate", "POST", path, in)
	return err
}
func (h *AccountTeamHandler) AccountTeamProjectDisassociate(ctx context.Context, accountId string, teamId string, project string) error {
	path := fmt.Sprintf("/account/%s/team/%s/project/%s", accountId, teamId, project)
	_, err := h.doer.Do(ctx, "AccountTeamProjectDisassociate", "DELETE", path, nil)
	return err
}
func (h *AccountTeamHandler) AccountTeamUpdate(ctx context.Context, accountId string, teamId string, in *AccountTeamUpdateIn) (*AccountTeamUpdateOut, error) {
	path := fmt.Sprintf("/account/%s/team/%s", accountId, teamId)
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
	AccountId          string `json:"account_id"`
	AccountName        string `json:"account_name"`
	CreateTime         string `json:"create_time"`
	InvitedByUserEmail string `json:"invited_by_user_email"`
	TeamId             string `json:"team_id"`
	TeamName           string `json:"team_name"`
	UserEmail          string `json:"user_email"`
}
type AccountTeamGetOut struct {
	AccountId  string `json:"account_id,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	TeamId     string `json:"team_id"`
	TeamName   string `json:"team_name"`
	UpdateTime string `json:"update_time,omitempty"`
}
type AccountTeamProjectAssociateIn struct {
	TeamType TeamType `json:"team_type"`
}
type AccountTeamUpdateIn struct {
	TeamName string `json:"team_name"`
}
type AccountTeamUpdateOut struct {
	AccountId  string `json:"account_id,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	TeamId     string `json:"team_id"`
	TeamName   string `json:"team_name"`
	UpdateTime string `json:"update_time,omitempty"`
}
type TeamOut struct {
	AccountId  string `json:"account_id,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	TeamId     string `json:"team_id"`
	TeamName   string `json:"team_name"`
	UpdateTime string `json:"update_time,omitempty"`
}
type TeamType string

const (
	TeamTypeAdmin     TeamType = "admin"
	TeamTypeOperator  TeamType = "operator"
	TeamTypeDeveloper TeamType = "developer"
	TeamTypeReadOnly  TeamType = "read_only"
)

func TeamTypeChoices() []string {
	return []string{"admin", "operator", "developer", "read_only"}
}

type accountTeamGetOut struct {
	Team AccountTeamGetOut `json:"team"`
}
type accountTeamInvitesListOut struct {
	AccountInvites []AccountInviteOut `json:"account_invites"`
}
type accountTeamListOut struct {
	Teams []TeamOut `json:"teams"`
}
type accountTeamUpdateOut struct {
	Team AccountTeamUpdateOut `json:"team"`
}
