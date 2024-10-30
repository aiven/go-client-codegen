// Code generated by Aiven. DO NOT EDIT.

package account

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// AccountAttachPaymentMethod attach payment method for account
	// POST /v1/account/{account_id}/payment_methods
	// https://api.aiven.io/doc/#tag/Account/operation/AccountAttachPaymentMethod
	AccountAttachPaymentMethod(ctx context.Context, accountId string, in *AccountAttachPaymentMethodIn) (*AccountAttachPaymentMethodOut, error)

	// AccountBillingGroupList list account billing groups
	// GET /v1/account/{account_id}/billing-group
	// https://api.aiven.io/doc/#tag/Account/operation/AccountBillingGroupList
	AccountBillingGroupList(ctx context.Context, accountId string) ([]AccountBillingGroupOut, error)

	// AccountCreate create a new account
	// POST /v1/account
	// https://api.aiven.io/doc/#tag/Account/operation/AccountCreate
	AccountCreate(ctx context.Context, in *AccountCreateIn) (*AccountCreateOut, error)

	// AccountDelete delete empty account
	// DELETE /v1/account/{account_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountDelete
	AccountDelete(ctx context.Context, accountId string) error

	// AccountEventList list account events
	// GET /v1/account/{account_id}/events
	// https://api.aiven.io/doc/#tag/Account/operation/AccountEventList
	AccountEventList(ctx context.Context, accountId string) ([]EventOut, error)

	// AccountGet get account details
	// GET /v1/account/{account_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountGet
	AccountGet(ctx context.Context, accountId string) (*AccountGetOut, error)

	// AccountList list accounts you have access to
	// GET /v1/account
	// https://api.aiven.io/doc/#tag/Account/operation/AccountList
	AccountList(ctx context.Context) ([]AccountOut, error)

	// AccountMove move an existing organization unitself
	// PUT /v1/account/{account_id}/parent_account
	// https://api.aiven.io/doc/#tag/Account/operation/AccountMove
	AccountMove(ctx context.Context, accountId string, in *AccountMoveIn) (*AccountMoveOut, error)

	// AccountPaymentMethodDelete delete credit card attached to the account as a payment method
	// DELETE /v1/account/{account_id}/payment_method/{card_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountPaymentMethodDelete
	AccountPaymentMethodDelete(ctx context.Context, accountId string, cardId string) error

	// AccountPaymentMethodsList list credit cards attached as a payment method to the account
	// GET /v1/account/{account_id}/payment_methods
	// https://api.aiven.io/doc/#tag/Account/operation/AccountPaymentMethodsList
	AccountPaymentMethodsList(ctx context.Context, accountId string) ([]CardOut, error)

	// AccountProjectsList list projects belonging to account
	// GET /v1/account/{account_id}/projects
	// https://api.aiven.io/doc/#tag/Account/operation/AccountProjectsList
	AccountProjectsList(ctx context.Context, accountId string) (*AccountProjectsListOut, error)

	// Deprecated: AccountProjectsTeamsList list account teams associated to a project
	// GET /v1/account/{account_id}/project/{project_name}/teams
	// https://api.aiven.io/doc/#tag/Account/operation/AccountProjectsTeamsList
	AccountProjectsTeamsList(ctx context.Context, accountId string, projectName string) ([]TeamOut, error)

	// AccountUpdate update existing account
	// PUT /v1/account/{account_id}
	// https://api.aiven.io/doc/#tag/Account/operation/AccountUpdate
	AccountUpdate(ctx context.Context, accountId string, in *AccountUpdateIn) (*AccountUpdateOut, error)

	// AccountUserProjectsList list projects associated with this account that user has access to
	// GET /v1/account/{account_id}/user/{user_id}/projects
	// https://api.aiven.io/doc/#tag/Account/operation/AccountUserProjectsList
	AccountUserProjectsList(ctx context.Context, accountId string, userId string) ([]UserProjectOut, error)

	// Deprecated: AccountUserTeamsList list all teams for user
	// GET /v1/account/{account_id}/user/{user_id}/teams
	// https://api.aiven.io/doc/#tag/Account/operation/AccountUserTeamsList
	AccountUserTeamsList(ctx context.Context, accountId string, userId string) ([]AccountUserTeamsListOut, error)

	// AccountUsersSearch list/search users who are members of any team on this account
	// POST /v1/account/{account_id}/users/search
	// https://api.aiven.io/doc/#tag/Account/operation/AccountUsersSearch
	AccountUsersSearch(ctx context.Context, accountId string, in *AccountUsersSearchIn) ([]UserOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) AccountHandler {
	return AccountHandler{doer}
}

type AccountHandler struct {
	doer doer
}

func (h *AccountHandler) AccountAttachPaymentMethod(ctx context.Context, accountId string, in *AccountAttachPaymentMethodIn) (*AccountAttachPaymentMethodOut, error) {
	path := fmt.Sprintf("/v1/account/%s/payment_methods", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountAttachPaymentMethod", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(accountAttachPaymentMethodOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Card, nil
}
func (h *AccountHandler) AccountBillingGroupList(ctx context.Context, accountId string) ([]AccountBillingGroupOut, error) {
	path := fmt.Sprintf("/v1/account/%s/billing-group", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountBillingGroupList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountBillingGroupListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.AccountBillingGroups, nil
}
func (h *AccountHandler) AccountCreate(ctx context.Context, in *AccountCreateIn) (*AccountCreateOut, error) {
	path := fmt.Sprintf("/v1/account")
	b, err := h.doer.Do(ctx, "AccountCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(accountCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Account, nil
}
func (h *AccountHandler) AccountDelete(ctx context.Context, accountId string) error {
	path := fmt.Sprintf("/v1/account/%s", url.PathEscape(accountId))
	_, err := h.doer.Do(ctx, "AccountDelete", "DELETE", path, nil)
	return err
}
func (h *AccountHandler) AccountEventList(ctx context.Context, accountId string) ([]EventOut, error) {
	path := fmt.Sprintf("/v1/account/%s/events", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountEventList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountEventListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Events, nil
}
func (h *AccountHandler) AccountGet(ctx context.Context, accountId string) (*AccountGetOut, error) {
	path := fmt.Sprintf("/v1/account/%s", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Account, nil
}
func (h *AccountHandler) AccountList(ctx context.Context) ([]AccountOut, error) {
	path := fmt.Sprintf("/v1/account")
	b, err := h.doer.Do(ctx, "AccountList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Accounts, nil
}
func (h *AccountHandler) AccountMove(ctx context.Context, accountId string, in *AccountMoveIn) (*AccountMoveOut, error) {
	path := fmt.Sprintf("/v1/account/%s/parent_account", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountMove", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(accountMoveOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Account, nil
}
func (h *AccountHandler) AccountPaymentMethodDelete(ctx context.Context, accountId string, cardId string) error {
	path := fmt.Sprintf("/v1/account/%s/payment_method/%s", url.PathEscape(accountId), url.PathEscape(cardId))
	_, err := h.doer.Do(ctx, "AccountPaymentMethodDelete", "DELETE", path, nil)
	return err
}
func (h *AccountHandler) AccountPaymentMethodsList(ctx context.Context, accountId string) ([]CardOut, error) {
	path := fmt.Sprintf("/v1/account/%s/payment_methods", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountPaymentMethodsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountPaymentMethodsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Cards, nil
}
func (h *AccountHandler) AccountProjectsList(ctx context.Context, accountId string) (*AccountProjectsListOut, error) {
	path := fmt.Sprintf("/v1/account/%s/projects", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountProjectsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(AccountProjectsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *AccountHandler) AccountProjectsTeamsList(ctx context.Context, accountId string, projectName string) ([]TeamOut, error) {
	path := fmt.Sprintf("/v1/account/%s/project/%s/teams", url.PathEscape(accountId), url.PathEscape(projectName))
	b, err := h.doer.Do(ctx, "AccountProjectsTeamsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountProjectsTeamsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Teams, nil
}
func (h *AccountHandler) AccountUpdate(ctx context.Context, accountId string, in *AccountUpdateIn) (*AccountUpdateOut, error) {
	path := fmt.Sprintf("/v1/account/%s", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountUpdate", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(accountUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Account, nil
}
func (h *AccountHandler) AccountUserProjectsList(ctx context.Context, accountId string, userId string) ([]UserProjectOut, error) {
	path := fmt.Sprintf("/v1/account/%s/user/%s/projects", url.PathEscape(accountId), url.PathEscape(userId))
	b, err := h.doer.Do(ctx, "AccountUserProjectsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountUserProjectsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.UserProjects, nil
}
func (h *AccountHandler) AccountUserTeamsList(ctx context.Context, accountId string, userId string) ([]AccountUserTeamsListOut, error) {
	path := fmt.Sprintf("/v1/account/%s/user/%s/teams", url.PathEscape(accountId), url.PathEscape(userId))
	b, err := h.doer.Do(ctx, "AccountUserTeamsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(accountUserTeamsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Teams, nil
}
func (h *AccountHandler) AccountUsersSearch(ctx context.Context, accountId string, in *AccountUsersSearchIn) ([]UserOut, error) {
	path := fmt.Sprintf("/v1/account/%s/users/search", url.PathEscape(accountId))
	b, err := h.doer.Do(ctx, "AccountUsersSearch", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(accountUsersSearchOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Users, nil
}

type AccessSourceType string

const (
	AccessSourceTypeDescendantMembership   AccessSourceType = "descendant_membership"
	AccessSourceTypeOrganizationMembership AccessSourceType = "organization_membership"
	AccessSourceTypeProjectMembership      AccessSourceType = "project_membership"
	AccessSourceTypeTeamMembership         AccessSourceType = "team_membership"
)

func AccessSourceTypeChoices() []string {
	return []string{"descendant_membership", "organization_membership", "project_membership", "team_membership"}
}

// AccountAttachPaymentMethodIn AccountAttachPaymentMethodRequestBody
type AccountAttachPaymentMethodIn struct {
	PaymentMethodId string `json:"payment_method_id"` // Unique identifier for a Stripe payment method
}

// AccountAttachPaymentMethodOut User credit card information
type AccountAttachPaymentMethodOut struct {
	Brand          string   `json:"brand"`
	CardId         string   `json:"card_id"` // Credit card ID
	Country        string   `json:"country"`
	CountryCode    string   `json:"country_code"`              // Two letter ISO country code
	ExpMonth       int      `json:"exp_month"`                 // Expiration month
	ExpYear        int      `json:"exp_year"`                  // Expiration year
	Last4          string   `json:"last4"`                     // Credit card last four digits
	Name           string   `json:"name"`                      // Name on the credit card
	OrganizationId *string  `json:"organization_id,omitempty"` // Organization ID
	Projects       []string `json:"projects"`                  // List of projects the card is assigned to
}
type AccountBillingGroupOut struct {
	AccountId             string              `json:"account_id"`                // Account ID
	AccountName           string              `json:"account_name"`              // Account name
	AddressLines          []string            `json:"address_lines"`             // Address lines
	BillingAddress        *string             `json:"billing_address,omitempty"` // DEPRECATED: use split address fields like company, address_lines, zip_code, city and state instead
	BillingCurrency       BillingCurrencyType `json:"billing_currency"`          // Billing currency
	BillingEmails         []BillingEmailOut   `json:"billing_emails"`            // List of project billing email addresses
	BillingExtraText      string              `json:"billing_extra_text"`        // Extra text to be included in all project invoices, e.g. purchase order or cost center number
	BillingGroupId        string              `json:"billing_group_id"`          // Billing group ID
	BillingGroupName      string              `json:"billing_group_name"`        // Billing group name
	BillingType           string              `json:"billing_type"`              // Method of charging/invoicing this project
	CardInfo              CardInfoOut         `json:"card_info"`                 // Credit card assigned to the project
	City                  string              `json:"city"`                      // Address city
	Company               string              `json:"company"`                   // Name of a company
	Country               string              `json:"country"`                   // Billing country
	CountryCode           string              `json:"country_code"`              // Two letter ISO country code
	CreateTime            time.Time           `json:"create_time"`               // Timestamp in ISO 8601 format, always in UTC
	EstimatedBalanceLocal string              `json:"estimated_balance_local"`   // Estimated balance in billing currency, before tax
	EstimatedBalanceUsd   string              `json:"estimated_balance_usd"`     // Estimated balance in USD, before tax
	PaymentMethod         PaymentMethodType   `json:"payment_method"`            // Payment method
	State                 string              `json:"state"`                     // Address state
	VatId                 string              `json:"vat_id"`                    // EU VAT Identification Number
	ZipCode               string              `json:"zip_code"`                  // Address zip code
}

// AccountCreateIn AccountCreateRequestBody
type AccountCreateIn struct {
	AccountName           string  `json:"account_name"`                       // Account name
	ParentAccountId       *string `json:"parent_account_id,omitempty"`        // Account ID
	PrimaryBillingGroupId *string `json:"primary_billing_group_id,omitempty"` // Billing group ID
}

// AccountCreateOut Account details
type AccountCreateOut struct {
	AccessSource          AccessSourceType `json:"access_source,omitempty"`     // Describe the source of the account
	AccountId             string           `json:"account_id"`                  // Account ID
	AccountName           string           `json:"account_name"`                // Account name
	AccountOwnerTeamId    string           `json:"account_owner_team_id"`       // Team ID
	CreateTime            time.Time        `json:"create_time"`                 // Timestamp in ISO 8601 format, always in UTC
	Features              map[string]any   `json:"features,omitempty"`          // Feature flags
	IsAccountMember       *bool            `json:"is_account_member,omitempty"` // If true, user is part of a team of this or a parent account
	IsAccountOwner        bool             `json:"is_account_owner"`            // If true, user is part of the owners team for this account
	OrganizationId        string           `json:"organization_id"`             // Organization ID
	ParentAccountId       *string          `json:"parent_account_id,omitempty"` // Account ID
	PrimaryBillingGroupId string           `json:"primary_billing_group_id"`    // Billing group ID
	RootAccountId         string           `json:"root_account_id"`             // Account ID
	TenantId              *string          `json:"tenant_id,omitempty"`         // Tenant identifier
	UpdateTime            time.Time        `json:"update_time"`                 // Timestamp in ISO 8601 format, always in UTC
}

// AccountGetOut Account details
type AccountGetOut struct {
	AccessSource          AccessSourceType `json:"access_source,omitempty"`     // Describe the source of the account
	AccountId             string           `json:"account_id"`                  // Account ID
	AccountName           string           `json:"account_name"`                // Account name
	AccountOwnerTeamId    string           `json:"account_owner_team_id"`       // Team ID
	CreateTime            time.Time        `json:"create_time"`                 // Timestamp in ISO 8601 format, always in UTC
	Features              map[string]any   `json:"features,omitempty"`          // Feature flags
	IsAccountMember       *bool            `json:"is_account_member,omitempty"` // If true, user is part of a team of this or a parent account
	IsAccountOwner        bool             `json:"is_account_owner"`            // If true, user is part of the owners team for this account
	OrganizationId        string           `json:"organization_id"`             // Organization ID
	ParentAccountId       *string          `json:"parent_account_id,omitempty"` // Account ID
	PrimaryBillingGroupId string           `json:"primary_billing_group_id"`    // Billing group ID
	RootAccountId         string           `json:"root_account_id"`             // Account ID
	TenantId              *string          `json:"tenant_id,omitempty"`         // Tenant identifier
	UpdateTime            time.Time        `json:"update_time"`                 // Timestamp in ISO 8601 format, always in UTC
}

// AccountMoveIn AccountMoveRequestBody
type AccountMoveIn struct {
	ParentAccountId string `json:"parent_account_id"` // Account ID
}

// AccountMoveOut Account details
type AccountMoveOut struct {
	AccessSource          AccessSourceType `json:"access_source,omitempty"`     // Describe the source of the account
	AccountId             string           `json:"account_id"`                  // Account ID
	AccountName           string           `json:"account_name"`                // Account name
	AccountOwnerTeamId    string           `json:"account_owner_team_id"`       // Team ID
	CreateTime            time.Time        `json:"create_time"`                 // Timestamp in ISO 8601 format, always in UTC
	Features              map[string]any   `json:"features,omitempty"`          // Feature flags
	IsAccountMember       *bool            `json:"is_account_member,omitempty"` // If true, user is part of a team of this or a parent account
	IsAccountOwner        bool             `json:"is_account_owner"`            // If true, user is part of the owners team for this account
	OrganizationId        string           `json:"organization_id"`             // Organization ID
	ParentAccountId       *string          `json:"parent_account_id,omitempty"` // Account ID
	PrimaryBillingGroupId string           `json:"primary_billing_group_id"`    // Billing group ID
	RootAccountId         string           `json:"root_account_id"`             // Account ID
	TenantId              *string          `json:"tenant_id,omitempty"`         // Tenant identifier
	UpdateTime            time.Time        `json:"update_time"`                 // Timestamp in ISO 8601 format, always in UTC
}
type AccountOut struct {
	AccessSource          AccessSourceType `json:"access_source,omitempty"`     // Describe the source of the account
	AccountId             string           `json:"account_id"`                  // Account ID
	AccountName           string           `json:"account_name"`                // Account name
	AccountOwnerTeamId    string           `json:"account_owner_team_id"`       // Team ID
	CreateTime            time.Time        `json:"create_time"`                 // Timestamp in ISO 8601 format, always in UTC
	Features              map[string]any   `json:"features,omitempty"`          // Feature flags
	IsAccountMember       *bool            `json:"is_account_member,omitempty"` // If true, user is part of a team of this or a parent account
	IsAccountOwner        bool             `json:"is_account_owner"`            // If true, user is part of the owners team for this account
	OrganizationId        string           `json:"organization_id"`             // Organization ID
	ParentAccountId       *string          `json:"parent_account_id,omitempty"` // Account ID
	PrimaryBillingGroupId string           `json:"primary_billing_group_id"`    // Billing group ID
	RootAccountId         string           `json:"root_account_id"`             // Account ID
	TenantId              *string          `json:"tenant_id,omitempty"`         // Tenant identifier
	UpdateTime            time.Time        `json:"update_time"`                 // Timestamp in ISO 8601 format, always in UTC
}

// AccountProjectsListOut AccountProjectsListResponse
type AccountProjectsListOut struct {
	Projects          []ProjectOut `json:"projects"`                      // List of projects
	TotalProjectCount *int         `json:"total_project_count,omitempty"` // Total count of projects associated to account.
}

// AccountUpdateIn AccountUpdateRequestBody
type AccountUpdateIn struct {
	AccountName           *string `json:"account_name,omitempty"`             // Account name
	PrimaryBillingGroupId *string `json:"primary_billing_group_id,omitempty"` // Billing group ID
}

// AccountUpdateOut Account details
type AccountUpdateOut struct {
	AccessSource          AccessSourceType `json:"access_source,omitempty"`     // Describe the source of the account
	AccountId             string           `json:"account_id"`                  // Account ID
	AccountName           string           `json:"account_name"`                // Account name
	AccountOwnerTeamId    string           `json:"account_owner_team_id"`       // Team ID
	CreateTime            time.Time        `json:"create_time"`                 // Timestamp in ISO 8601 format, always in UTC
	Features              map[string]any   `json:"features,omitempty"`          // Feature flags
	IsAccountMember       *bool            `json:"is_account_member,omitempty"` // If true, user is part of a team of this or a parent account
	IsAccountOwner        bool             `json:"is_account_owner"`            // If true, user is part of the owners team for this account
	OrganizationId        string           `json:"organization_id"`             // Organization ID
	ParentAccountId       *string          `json:"parent_account_id,omitempty"` // Account ID
	PrimaryBillingGroupId string           `json:"primary_billing_group_id"`    // Billing group ID
	RootAccountId         string           `json:"root_account_id"`             // Account ID
	TenantId              *string          `json:"tenant_id,omitempty"`         // Tenant identifier
	UpdateTime            time.Time        `json:"update_time"`                 // Timestamp in ISO 8601 format, always in UTC
}
type AccountUserTeamsListOut struct {
	AccountId   string `json:"account_id"`   // Account ID
	AccountName string `json:"account_name"` // Account name
	TeamId      string `json:"team_id"`      // Team ID
	TeamName    string `json:"team_name"`    // Team name
}

// AccountUsersSearchIn AccountUsersSearchRequestBody
type AccountUsersSearchIn struct {
	Limit   *int        `json:"limit,omitempty"`    // Maximum number of results to return
	OrderBy OrderByType `json:"order_by,omitempty"` // Sorting criteria; desc is descending order and asc ascending
	Query   *string     `json:"query,omitempty"`    // Filter keyword
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

func BillingCurrencyTypeChoices() []string {
	return []string{"AUD", "CAD", "CHF", "DKK", "EUR", "GBP", "JPY", "NOK", "NZD", "SEK", "SGD", "USD"}
}

type BillingEmailOut struct {
	Email string `json:"email"` // User email address
}

// CardInfoOut Credit card assigned to the project
type CardInfoOut struct {
	Brand       string `json:"brand"`
	CardId      string `json:"card_id"` // Credit card ID
	Country     string `json:"country"`
	CountryCode string `json:"country_code"` // Two letter ISO country code
	ExpMonth    int    `json:"exp_month"`    // Expiration month
	ExpYear     int    `json:"exp_year"`     // Expiration year
	Last4       string `json:"last4"`        // Credit card last four digits
	Name        string `json:"name"`         // Name on the credit card
	UserEmail   string `json:"user_email"`   // User email address
}
type CardOut struct {
	Brand       string `json:"brand"`
	CardId      string `json:"card_id"` // Credit card ID
	Country     string `json:"country"`
	CountryCode string `json:"country_code"` // Two letter ISO country code
	ExpMonth    int    `json:"exp_month"`    // Expiration month
	ExpYear     int    `json:"exp_year"`     // Expiration year
	Last4       string `json:"last4"`        // Credit card last four digits
	Name        string `json:"name"`         // Name on the credit card
}

// ElasticsearchOut Service EOL extension
type ElasticsearchOut struct {
	EolDate string `json:"eol_date"` // Extended EOL date
	Version string `json:"version"`  // Service version
}

// EndOfLifeExtensionOut End of life extension information
type EndOfLifeExtensionOut struct {
	Elasticsearch *ElasticsearchOut `json:"elasticsearch,omitempty"` // Service EOL extension
}
type EventOut struct {
	AccountId         string    `json:"account_id"`         // Account ID
	ActionDescription string    `json:"action_description"` // Event description
	ActionType        string    `json:"action_type"`        // Event type
	Actor             string    `json:"actor"`              // Actor details
	ActorUserId       string    `json:"actor_user_id"`      // User ID
	CreateTime        time.Time `json:"create_time"`        // Timestamp in ISO 8601 format, always in UTC
	LogEntryId        int       `json:"log_entry_id"`       // Entry ID
	TeamId            string    `json:"team_id"`            // Team ID
}
type MemberType string

const (
	MemberTypeAdmin                     MemberType = "admin"
	MemberTypeDeveloper                 MemberType = "developer"
	MemberTypeOperator                  MemberType = "operator"
	MemberTypeProjectAuditLogsRead      MemberType = "project:audit_logs:read"
	MemberTypeProjectIntegrationsRead   MemberType = "project:integrations:read"
	MemberTypeProjectIntegrationsWrite  MemberType = "project:integrations:write"
	MemberTypeProjectNetworkingRead     MemberType = "project:networking:read"
	MemberTypeProjectNetworkingWrite    MemberType = "project:networking:write"
	MemberTypeProjectPermissionsRead    MemberType = "project:permissions:read"
	MemberTypeProjectServicesRead       MemberType = "project:services:read"
	MemberTypeReadOnly                  MemberType = "read_only"
	MemberTypeServiceConfigurationWrite MemberType = "service:configuration:write"
	MemberTypeServiceLogsRead           MemberType = "service:logs:read"
	MemberTypeServicesMaintenance       MemberType = "services:maintenance"
)

func MemberTypeChoices() []string {
	return []string{"admin", "developer", "operator", "project:audit_logs:read", "project:integrations:read", "project:integrations:write", "project:networking:read", "project:networking:write", "project:permissions:read", "project:services:read", "read_only", "service:configuration:write", "service:logs:read", "services:maintenance"}
}

type OrderByType string

const (
	OrderByTypeUserEmailAsc  OrderByType = "user_email:asc"
	OrderByTypeUserEmailDesc OrderByType = "user_email:desc"
	OrderByTypeUserIdAsc     OrderByType = "user_id:asc"
	OrderByTypeUserIdDesc    OrderByType = "user_id:desc"
	OrderByTypeRealNameAsc   OrderByType = "real_name:asc"
	OrderByTypeRealNameDesc  OrderByType = "real_name:desc"
)

func OrderByTypeChoices() []string {
	return []string{"user_email:asc", "user_email:desc", "user_id:asc", "user_id:desc", "real_name:asc", "real_name:desc"}
}

type PaymentMethodType string

const (
	PaymentMethodTypeAccrual           PaymentMethodType = "accrual"
	PaymentMethodTypeCard              PaymentMethodType = "card"
	PaymentMethodTypeDisabled          PaymentMethodType = "disabled"
	PaymentMethodTypeEmail             PaymentMethodType = "email"
	PaymentMethodTypeNoPaymentExpected PaymentMethodType = "no_payment_expected"
	PaymentMethodTypePartner           PaymentMethodType = "partner"
)

func PaymentMethodTypeChoices() []string {
	return []string{"accrual", "card", "disabled", "email", "no_payment_expected", "partner"}
}

type ProjectOut struct {
	AccountId             string                 `json:"account_id"`                        // Account ID
	AccountName           *string                `json:"account_name,omitempty"`            // Account name
	AddressLines          []string               `json:"address_lines,omitempty"`           // Address lines
	AvailableCredits      *string                `json:"available_credits,omitempty"`       // Available credits, in USD
	BillingAddress        string                 `json:"billing_address"`                   // DEPRECATED: use split address fields like company, address_lines, zip_code, city and state instead
	BillingCurrency       BillingCurrencyType    `json:"billing_currency,omitempty"`        // Billing currency
	BillingEmails         []BillingEmailOut      `json:"billing_emails"`                    // List of project billing email addresses
	BillingExtraText      *string                `json:"billing_extra_text,omitempty"`      // Extra text to be included in all project invoices, e.g. purchase order or cost center number
	BillingGroupId        string                 `json:"billing_group_id"`                  // Billing group ID
	BillingGroupName      string                 `json:"billing_group_name"`                // Billing group name
	CardInfo              *CardInfoOut           `json:"card_info,omitempty"`               // Credit card assigned to the project
	City                  *string                `json:"city,omitempty"`                    // Address city
	Company               *string                `json:"company,omitempty"`                 // Name of a company
	Country               string                 `json:"country"`                           // Billing country
	CountryCode           string                 `json:"country_code"`                      // Two letter ISO country code
	DefaultCloud          string                 `json:"default_cloud"`                     // Default cloud to use when launching services
	EndOfLifeExtension    *EndOfLifeExtensionOut `json:"end_of_life_extension,omitempty"`   // End of life extension information
	EstimatedBalance      string                 `json:"estimated_balance"`                 // Estimated balance, in USD
	EstimatedBalanceLocal *string                `json:"estimated_balance_local,omitempty"` // Estimated balance, in billing currency
	Features              map[string]any         `json:"features,omitempty"`                // Feature flags
	OrganizationId        string                 `json:"organization_id"`                   // Organization ID
	PaymentMethod         string                 `json:"payment_method"`                    // Payment method
	ProjectName           string                 `json:"project_name"`                      // Project name
	State                 *string                `json:"state,omitempty"`                   // Address state
	Tags                  map[string]string      `json:"tags,omitempty"`                    // Set of resource tags
	TechEmails            []TechEmailOut         `json:"tech_emails,omitempty"`             // List of project tech email addresses
	TenantId              *string                `json:"tenant_id,omitempty"`               // Tenant ID
	TrialExpirationTime   *time.Time             `json:"trial_expiration_time,omitempty"`   // Trial expiration time (ISO 8601)
	VatId                 string                 `json:"vat_id"`                            // EU VAT Identification Number
	ZipCode               *string                `json:"zip_code,omitempty"`                // Address zip code
}
type TeamOut struct {
	AccountId  *string    `json:"account_id,omitempty"`  // Account ID
	CreateTime *time.Time `json:"create_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
	TeamId     string     `json:"team_id"`               // Team ID
	TeamName   string     `json:"team_name"`             // Team name
	TeamType   TeamType   `json:"team_type,omitempty"`   // Team type (permission level)
	UpdateTime *time.Time `json:"update_time,omitempty"` // Timestamp in ISO 8601 format, always in UTC
}
type TeamType string

const (
	TeamTypeAdmin                     TeamType = "admin"
	TeamTypeOperator                  TeamType = "operator"
	TeamTypeDeveloper                 TeamType = "developer"
	TeamTypeReadOnly                  TeamType = "read_only"
	TeamTypeProjectIntegrationsRead   TeamType = "project:integrations:read"
	TeamTypeProjectIntegrationsWrite  TeamType = "project:integrations:write"
	TeamTypeProjectNetworkingRead     TeamType = "project:networking:read"
	TeamTypeProjectNetworkingWrite    TeamType = "project:networking:write"
	TeamTypeProjectPermissionsRead    TeamType = "project:permissions:read"
	TeamTypeServiceConfigurationWrite TeamType = "service:configuration:write"
	TeamTypeServicesMaintenance       TeamType = "services:maintenance"
	TeamTypeServiceLogsRead           TeamType = "service:logs:read"
	TeamTypeProjectServicesRead       TeamType = "project:services:read"
	TeamTypeProjectAuditLogsRead      TeamType = "project:audit_logs:read"
)

func TeamTypeChoices() []string {
	return []string{"admin", "operator", "developer", "read_only", "project:integrations:read", "project:integrations:write", "project:networking:read", "project:networking:write", "project:permissions:read", "service:configuration:write", "services:maintenance", "service:logs:read", "project:services:read", "project:audit_logs:read"}
}

type TechEmailOut struct {
	Email string `json:"email"` // User email address
}
type UserOut struct {
	RealName  string `json:"real_name"`  // User real name
	UserEmail string `json:"user_email"` // User email address
	UserId    string `json:"user_id"`    // User ID
}
type UserProjectOut struct {
	AccessType  *string    `json:"access_type,omitempty"` // Access type
	AccountId   string     `json:"account_id"`            // Account ID
	CreateTime  time.Time  `json:"create_time"`           // Timestamp in ISO 8601 format, always in UTC
	MemberType  MemberType `json:"member_type"`           // Project member type
	ProjectName string     `json:"project_name"`          // Project name
	RealName    string     `json:"real_name"`             // User real name
	TeamId      string     `json:"team_id"`               // Team ID
	TeamName    string     `json:"team_name"`             // Team name
	UserEmail   string     `json:"user_email"`            // User email address
}

// accountAttachPaymentMethodOut AccountAttachPaymentMethodResponse
type accountAttachPaymentMethodOut struct {
	Card AccountAttachPaymentMethodOut `json:"card"` // User credit card information
}

// accountBillingGroupListOut AccountBillingGroupListResponse
type accountBillingGroupListOut struct {
	AccountBillingGroups []AccountBillingGroupOut `json:"account_billing_groups"` // List of billing groups
}

// accountCreateOut AccountCreateResponse
type accountCreateOut struct {
	Account AccountCreateOut `json:"account"` // Account details
}

// accountEventListOut AccountEventListResponse
type accountEventListOut struct {
	Events []EventOut `json:"events"` // List of events
}

// accountGetOut AccountGetResponse
type accountGetOut struct {
	Account AccountGetOut `json:"account"` // Account details
}

// accountListOut AccountListResponse
type accountListOut struct {
	Accounts []AccountOut `json:"accounts"` // List of accounts
}

// accountMoveOut AccountMoveResponse
type accountMoveOut struct {
	Account AccountMoveOut `json:"account"` // Account details
}

// accountPaymentMethodsListOut AccountPaymentMethodsListResponse
type accountPaymentMethodsListOut struct {
	Cards []CardOut `json:"cards"` // List of account's credit cards
}

// accountProjectsTeamsListOut AccountProjectsTeamsListResponse
type accountProjectsTeamsListOut struct {
	Teams []TeamOut `json:"teams"` // List of teams
}

// accountUpdateOut AccountUpdateResponse
type accountUpdateOut struct {
	Account AccountUpdateOut `json:"account"` // Account details
}

// accountUserProjectsListOut AccountUserProjectsListResponse
type accountUserProjectsListOut struct {
	UserProjects []UserProjectOut `json:"user_projects"` // List of user's projects
}

// accountUserTeamsListOut AccountUserTeamsListResponse
type accountUserTeamsListOut struct {
	Teams []AccountUserTeamsListOut `json:"teams"` // List of teams
}

// accountUsersSearchOut AccountUsersSearchResponse
type accountUsersSearchOut struct {
	Users []UserOut `json:"users"` // List of users
}
