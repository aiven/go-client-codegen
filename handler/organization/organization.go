// Code generated by Aiven. DO NOT EDIT.

package organization

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// OrganizationAuthenticationConfigGet retrieve authentication configuration
	// GET /organization/{organization_id}/config/authentication
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationAuthenticationConfigGet
	OrganizationAuthenticationConfigGet(ctx context.Context, organizationId string) (*OrganizationAuthenticationConfigGetOut, error)

	// OrganizationAuthenticationConfigUpdate update authentication configuration
	// PATCH /organization/{organization_id}/config/authentication
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationAuthenticationConfigUpdate
	OrganizationAuthenticationConfigUpdate(ctx context.Context, organizationId string, in *OrganizationAuthenticationConfigUpdateIn) (*OrganizationAuthenticationConfigUpdateOut, error)

	// OrganizationGet get information about an organization
	// GET /organization/{organization_id}
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationGet
	OrganizationGet(ctx context.Context, organizationId string) (*OrganizationGetOut, error)

	// OrganizationProjectsList list projects under the organization
	// GET /organization/{organization_id}/projects
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationProjectsList
	OrganizationProjectsList(ctx context.Context, organizationId string) (*OrganizationProjectsListOut, error)

	// OrganizationUpdate update organization's details
	// PATCH /organization/{organization_id}
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUpdate
	OrganizationUpdate(ctx context.Context, organizationId string, in *OrganizationUpdateIn) (*OrganizationUpdateOut, error)

	// UserOrganizationCreate create an organization
	// POST /organizations
	// https://api.aiven.io/doc/#tag/Organizations/operation/UserOrganizationCreate
	UserOrganizationCreate(ctx context.Context, in *UserOrganizationCreateIn) (*UserOrganizationCreateOut, error)

	// UserOrganizationsList list organizations the user belongs to
	// GET /organizations
	// https://api.aiven.io/doc/#tag/Organizations/operation/UserOrganizationsList
	UserOrganizationsList(ctx context.Context) ([]OrganizationOut, error)
}

func NewHandler(doer doer) OrganizationHandler {
	return OrganizationHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type OrganizationHandler struct {
	doer doer
}

func (h *OrganizationHandler) OrganizationAuthenticationConfigGet(ctx context.Context, organizationId string) (*OrganizationAuthenticationConfigGetOut, error) {
	path := fmt.Sprintf("/organization/%s/config/authentication", organizationId)
	b, err := h.doer.Do(ctx, "OrganizationAuthenticationConfigGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationAuthenticationConfigGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) OrganizationAuthenticationConfigUpdate(ctx context.Context, organizationId string, in *OrganizationAuthenticationConfigUpdateIn) (*OrganizationAuthenticationConfigUpdateOut, error) {
	path := fmt.Sprintf("/organization/%s/config/authentication", organizationId)
	b, err := h.doer.Do(ctx, "OrganizationAuthenticationConfigUpdate", "PATCH", path, in)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationAuthenticationConfigUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) OrganizationGet(ctx context.Context, organizationId string) (*OrganizationGetOut, error) {
	path := fmt.Sprintf("/organization/%s", organizationId)
	b, err := h.doer.Do(ctx, "OrganizationGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) OrganizationProjectsList(ctx context.Context, organizationId string) (*OrganizationProjectsListOut, error) {
	path := fmt.Sprintf("/organization/%s/projects", organizationId)
	b, err := h.doer.Do(ctx, "OrganizationProjectsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationProjectsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) OrganizationUpdate(ctx context.Context, organizationId string, in *OrganizationUpdateIn) (*OrganizationUpdateOut, error) {
	path := fmt.Sprintf("/organization/%s", organizationId)
	b, err := h.doer.Do(ctx, "OrganizationUpdate", "PATCH", path, in)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) UserOrganizationCreate(ctx context.Context, in *UserOrganizationCreateIn) (*UserOrganizationCreateOut, error) {
	path := fmt.Sprintf("/organizations")
	b, err := h.doer.Do(ctx, "UserOrganizationCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(UserOrganizationCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) UserOrganizationsList(ctx context.Context) ([]OrganizationOut, error) {
	path := fmt.Sprintf("/organizations")
	b, err := h.doer.Do(ctx, "UserOrganizationsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(userOrganizationsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Organizations, nil
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
	Email string `json:"email"`
}
type CardInfoOut struct {
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
type ElasticsearchOut struct {
	EolDate string `json:"eol_date"`
	Version string `json:"version"`
}
type EndOfLifeExtensionOut struct {
	Elasticsearch *ElasticsearchOut `json:"elasticsearch,omitempty"`
}
type OrganizationAuthenticationConfigGetOut struct {
	OauthEnabled        *bool `json:"oauth_enabled,omitempty"`
	PasswordAuthEnabled *bool `json:"password_auth_enabled,omitempty"`
	SamlEnabled         *bool `json:"saml_enabled,omitempty"`
	TwoFactorRequired   *bool `json:"two_factor_required,omitempty"`
}
type OrganizationAuthenticationConfigUpdateIn struct {
	OauthEnabled        *bool `json:"oauth_enabled,omitempty"`
	PasswordAuthEnabled *bool `json:"password_auth_enabled,omitempty"`
	SamlEnabled         *bool `json:"saml_enabled,omitempty"`
	TwoFactorRequired   *bool `json:"two_factor_required,omitempty"`
}
type OrganizationAuthenticationConfigUpdateOut struct {
	OauthEnabled        *bool `json:"oauth_enabled,omitempty"`
	PasswordAuthEnabled *bool `json:"password_auth_enabled,omitempty"`
	SamlEnabled         *bool `json:"saml_enabled,omitempty"`
	TwoFactorRequired   *bool `json:"two_factor_required,omitempty"`
}
type OrganizationGetOut struct {
	AccountId                    string    `json:"account_id"`
	CreateTime                   time.Time `json:"create_time"`
	DefaultGovernanceUserGroupId string    `json:"default_governance_user_group_id,omitempty"`
	OrganizationId               string    `json:"organization_id"`
	OrganizationName             string    `json:"organization_name"`
	Tier                         TierType  `json:"tier"`
	UpdateTime                   time.Time `json:"update_time"`
}
type OrganizationOut struct {
	AccountId                    string    `json:"account_id"`
	CreateTime                   time.Time `json:"create_time"`
	DefaultGovernanceUserGroupId string    `json:"default_governance_user_group_id,omitempty"`
	OrganizationId               string    `json:"organization_id"`
	OrganizationName             string    `json:"organization_name"`
	Tier                         TierType  `json:"tier"`
	UpdateTime                   time.Time `json:"update_time"`
}
type OrganizationProjectsListOut struct {
	Projects          []ProjectOut `json:"projects"`
	TotalProjectCount *int         `json:"total_project_count,omitempty"`
}
type OrganizationUpdateIn struct {
	DefaultGovernanceUserGroupId string   `json:"default_governance_user_group_id,omitempty"`
	KafkaGovernanceEnabled       *bool    `json:"kafka_governance_enabled,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Tier                         TierType `json:"tier,omitempty"`
}
type OrganizationUpdateOut struct {
	AccountId                    string    `json:"account_id"`
	CreateTime                   time.Time `json:"create_time"`
	DefaultGovernanceUserGroupId string    `json:"default_governance_user_group_id,omitempty"`
	OrganizationId               string    `json:"organization_id"`
	OrganizationName             string    `json:"organization_name"`
	Tier                         TierType  `json:"tier"`
	UpdateTime                   time.Time `json:"update_time"`
}
type ProjectOut struct {
	AccountId             string                 `json:"account_id"`
	AccountName           string                 `json:"account_name,omitempty"`
	AddressLines          []string               `json:"address_lines,omitempty"`
	AvailableCredits      string                 `json:"available_credits,omitempty"`
	BillingAddress        string                 `json:"billing_address"`
	BillingCurrency       BillingCurrencyType    `json:"billing_currency,omitempty"`
	BillingEmails         []BillingEmailOut      `json:"billing_emails"`
	BillingExtraText      string                 `json:"billing_extra_text,omitempty"`
	BillingGroupId        string                 `json:"billing_group_id"`
	BillingGroupName      string                 `json:"billing_group_name"`
	CardInfo              *CardInfoOut           `json:"card_info,omitempty"`
	City                  string                 `json:"city,omitempty"`
	Company               string                 `json:"company,omitempty"`
	Country               string                 `json:"country"`
	CountryCode           string                 `json:"country_code"`
	DefaultCloud          string                 `json:"default_cloud"`
	EndOfLifeExtension    *EndOfLifeExtensionOut `json:"end_of_life_extension,omitempty"`
	EstimatedBalance      string                 `json:"estimated_balance"`
	EstimatedBalanceLocal string                 `json:"estimated_balance_local,omitempty"`
	Features              map[string]any         `json:"features,omitempty"`
	OrganizationId        string                 `json:"organization_id"`
	PaymentMethod         string                 `json:"payment_method"`
	ProjectName           string                 `json:"project_name"`
	State                 string                 `json:"state,omitempty"`
	Tags                  map[string]string      `json:"tags,omitempty"`
	TechEmails            []TechEmailOut         `json:"tech_emails,omitempty"`
	TenantId              string                 `json:"tenant_id,omitempty"`
	TrialExpirationTime   *time.Time             `json:"trial_expiration_time,omitempty"`
	VatId                 string                 `json:"vat_id"`
	ZipCode               string                 `json:"zip_code,omitempty"`
}
type TechEmailOut struct {
	Email string `json:"email"`
}
type TierType string

const (
	TierTypeBusiness TierType = "business"
	TierTypePersonal TierType = "personal"
)

func TierTypeChoices() []string {
	return []string{"business", "personal"}
}

type UserOrganizationCreateIn struct {
	OrganizationName      string   `json:"organization_name"`
	PrimaryBillingGroupId string   `json:"primary_billing_group_id,omitempty"`
	Tier                  TierType `json:"tier"`
}
type UserOrganizationCreateOut struct {
	AccountId                    string    `json:"account_id"`
	CreateTime                   time.Time `json:"create_time"`
	DefaultGovernanceUserGroupId string    `json:"default_governance_user_group_id,omitempty"`
	OrganizationId               string    `json:"organization_id"`
	OrganizationName             string    `json:"organization_name"`
	Tier                         TierType  `json:"tier"`
	UpdateTime                   time.Time `json:"update_time"`
}
type userOrganizationsListOut struct {
	Organizations []OrganizationOut `json:"organizations"`
}
