// Code generated by Aiven. DO NOT EDIT.

package organization

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// OrganizationAddressCreate [EXPERIMENTAL] Create new address for an organization
	// POST /v1/organizations/{organization_id}/addresses
	// https://api.aiven.io/doc/#tag/Billing/operation/OrganizationAddressCreate
	OrganizationAddressCreate(ctx context.Context, organizationId string, in *OrganizationAddressCreateIn) (*OrganizationAddressCreateOut, error)

	// OrganizationAddressDelete [EXPERIMENTAL] Delete an address of an organization
	// DELETE /v1/organizations/{organization_id}/addresses/{address_id}
	// https://api.aiven.io/doc/#tag/Billing/operation/OrganizationAddressDelete
	OrganizationAddressDelete(ctx context.Context, organizationId string, addressId string) error

	// OrganizationAddressGet [EXPERIMENTAL] Get organization address info
	// GET /v1/organizations/{organization_id}/addresses/{address_id}
	// https://api.aiven.io/doc/#tag/Billing/operation/OrganizationAddressGet
	OrganizationAddressGet(ctx context.Context, organizationId string, addressId string) (*OrganizationAddressGetOut, error)

	// OrganizationAddressList [EXPERIMENTAL] List addresses of an organization
	// GET /v1/organizations/{organization_id}/addresses
	// https://api.aiven.io/doc/#tag/Billing/operation/OrganizationAddressList
	OrganizationAddressList(ctx context.Context, organizationId string) ([]AddresseOut, error)

	// OrganizationAddressUpdate [EXPERIMENTAL] Update an address of an organization
	// PATCH /v1/organizations/{organization_id}/addresses/{address_id}
	// https://api.aiven.io/doc/#tag/Billing/operation/OrganizationAddressUpdate
	OrganizationAddressUpdate(ctx context.Context, organizationId string, addressId string, in *OrganizationAddressUpdateIn) (*OrganizationAddressUpdateOut, error)

	// OrganizationAuthDomainLink link a domain to an organization's identity provider
	// PUT /v1/organization/{organization_id}/authentication-methods/{authentication_method_id}/domains
	// https://api.aiven.io/doc/#tag/Authentication_Methods/operation/OrganizationAuthDomainLink
	OrganizationAuthDomainLink(ctx context.Context, organizationId string, authenticationMethodId string, in *OrganizationAuthDomainLinkIn) error

	// OrganizationAuthDomainList list domains linked to an organization's identity provider
	// GET /v1/organization/{organization_id}/authentication-methods/{authentication_method_id}/domains
	// https://api.aiven.io/doc/#tag/Authentication_Methods/operation/OrganizationAuthDomainList
	OrganizationAuthDomainList(ctx context.Context, organizationId string, authenticationMethodId string) ([]DomainOut, error)

	// OrganizationAuthDomainUnlink unlink domain from authentication method
	// DELETE /v1/organization/{organization_id}/authentication-methods/{authentication_method_id}/domains/{domain_id}
	// https://api.aiven.io/doc/#tag/Authentication_Methods/operation/OrganizationAuthDomainUnlink
	OrganizationAuthDomainUnlink(ctx context.Context, organizationId string, authenticationMethodId string, domainId string) error

	// OrganizationAuthenticationConfigGet retrieve authentication configuration
	// GET /v1/organization/{organization_id}/config/authentication
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationAuthenticationConfigGet
	OrganizationAuthenticationConfigGet(ctx context.Context, organizationId string) (*OrganizationAuthenticationConfigGetOut, error)

	// OrganizationAuthenticationConfigUpdate update authentication configuration
	// PATCH /v1/organization/{organization_id}/config/authentication
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationAuthenticationConfigUpdate
	OrganizationAuthenticationConfigUpdate(ctx context.Context, organizationId string, in *OrganizationAuthenticationConfigUpdateIn) (*OrganizationAuthenticationConfigUpdateOut, error)

	// OrganizationGet get information about an organization
	// GET /v1/organization/{organization_id}
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationGet
	OrganizationGet(ctx context.Context, organizationId string) (*OrganizationGetOut, error)

	// OrganizationUpdate update organization's details
	// PATCH /v1/organization/{organization_id}
	// https://api.aiven.io/doc/#tag/Organizations/operation/OrganizationUpdate
	OrganizationUpdate(ctx context.Context, organizationId string, in *OrganizationUpdateIn) (*OrganizationUpdateOut, error)

	// PermissionsGet list of permissions
	// GET /v1/organization/{organization_id}/permissions/{resource_type}/{resource_id}
	// https://api.aiven.io/doc/#tag/Permissions/operation/PermissionsGet
	PermissionsGet(ctx context.Context, organizationId string, resourceType ResourceType, resourceId string) ([]PermissionOut, error)

	// PermissionsSet set permissions
	// PUT /v1/organization/{organization_id}/permissions/{resource_type}/{resource_id}
	// https://api.aiven.io/doc/#tag/Permissions/operation/PermissionsSet
	PermissionsSet(ctx context.Context, organizationId string, resourceType ResourceType, resourceId string, in *PermissionsSetIn) error

	// PermissionsUpdate update permissions
	// PATCH /v1/organization/{organization_id}/permissions/{resource_type}/{resource_id}
	// https://api.aiven.io/doc/#tag/Permissions/operation/PermissionsUpdate
	PermissionsUpdate(ctx context.Context, organizationId string, resourceType ResourceType, resourceId string, in *PermissionsUpdateIn) error

	// UserOrganizationCreate create an organization
	// POST /v1/organizations
	// https://api.aiven.io/doc/#tag/Organizations/operation/UserOrganizationCreate
	UserOrganizationCreate(ctx context.Context, in *UserOrganizationCreateIn) (*UserOrganizationCreateOut, error)

	// UserOrganizationsList list organizations the user belongs to
	// GET /v1/organizations
	// https://api.aiven.io/doc/#tag/Organizations/operation/UserOrganizationsList
	UserOrganizationsList(ctx context.Context) ([]OrganizationOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) OrganizationHandler {
	return OrganizationHandler{doer}
}

type OrganizationHandler struct {
	doer doer
}

func (h *OrganizationHandler) OrganizationAddressCreate(ctx context.Context, organizationId string, in *OrganizationAddressCreateIn) (*OrganizationAddressCreateOut, error) {
	path := fmt.Sprintf("/v1/organizations/%s/addresses", url.PathEscape(organizationId))
	b, err := h.doer.Do(ctx, "OrganizationAddressCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationAddressCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) OrganizationAddressDelete(ctx context.Context, organizationId string, addressId string) error {
	path := fmt.Sprintf("/v1/organizations/%s/addresses/%s", url.PathEscape(organizationId), url.PathEscape(addressId))
	_, err := h.doer.Do(ctx, "OrganizationAddressDelete", "DELETE", path, nil)
	return err
}
func (h *OrganizationHandler) OrganizationAddressGet(ctx context.Context, organizationId string, addressId string) (*OrganizationAddressGetOut, error) {
	path := fmt.Sprintf("/v1/organizations/%s/addresses/%s", url.PathEscape(organizationId), url.PathEscape(addressId))
	b, err := h.doer.Do(ctx, "OrganizationAddressGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationAddressGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) OrganizationAddressList(ctx context.Context, organizationId string) ([]AddresseOut, error) {
	path := fmt.Sprintf("/v1/organizations/%s/addresses", url.PathEscape(organizationId))
	b, err := h.doer.Do(ctx, "OrganizationAddressList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(organizationAddressListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Addresses, nil
}
func (h *OrganizationHandler) OrganizationAddressUpdate(ctx context.Context, organizationId string, addressId string, in *OrganizationAddressUpdateIn) (*OrganizationAddressUpdateOut, error) {
	path := fmt.Sprintf("/v1/organizations/%s/addresses/%s", url.PathEscape(organizationId), url.PathEscape(addressId))
	b, err := h.doer.Do(ctx, "OrganizationAddressUpdate", "PATCH", path, in)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationAddressUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationHandler) OrganizationAuthDomainLink(ctx context.Context, organizationId string, authenticationMethodId string, in *OrganizationAuthDomainLinkIn) error {
	path := fmt.Sprintf("/v1/organization/%s/authentication-methods/%s/domains", url.PathEscape(organizationId), url.PathEscape(authenticationMethodId))
	_, err := h.doer.Do(ctx, "OrganizationAuthDomainLink", "PUT", path, in)
	return err
}
func (h *OrganizationHandler) OrganizationAuthDomainList(ctx context.Context, organizationId string, authenticationMethodId string) ([]DomainOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/authentication-methods/%s/domains", url.PathEscape(organizationId), url.PathEscape(authenticationMethodId))
	b, err := h.doer.Do(ctx, "OrganizationAuthDomainList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(organizationAuthDomainListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Domains, nil
}
func (h *OrganizationHandler) OrganizationAuthDomainUnlink(ctx context.Context, organizationId string, authenticationMethodId string, domainId string) error {
	path := fmt.Sprintf("/v1/organization/%s/authentication-methods/%s/domains/%s", url.PathEscape(organizationId), url.PathEscape(authenticationMethodId), url.PathEscape(domainId))
	_, err := h.doer.Do(ctx, "OrganizationAuthDomainUnlink", "DELETE", path, nil)
	return err
}
func (h *OrganizationHandler) OrganizationAuthenticationConfigGet(ctx context.Context, organizationId string) (*OrganizationAuthenticationConfigGetOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/config/authentication", url.PathEscape(organizationId))
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
	path := fmt.Sprintf("/v1/organization/%s/config/authentication", url.PathEscape(organizationId))
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
	path := fmt.Sprintf("/v1/organization/%s", url.PathEscape(organizationId))
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
func (h *OrganizationHandler) OrganizationUpdate(ctx context.Context, organizationId string, in *OrganizationUpdateIn) (*OrganizationUpdateOut, error) {
	path := fmt.Sprintf("/v1/organization/%s", url.PathEscape(organizationId))
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
func (h *OrganizationHandler) PermissionsGet(ctx context.Context, organizationId string, resourceType ResourceType, resourceId string) ([]PermissionOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/permissions/%s/%s", url.PathEscape(organizationId), url.PathEscape(string(resourceType)), url.PathEscape(resourceId))
	b, err := h.doer.Do(ctx, "PermissionsGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(permissionsGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Permissions, nil
}
func (h *OrganizationHandler) PermissionsSet(ctx context.Context, organizationId string, resourceType ResourceType, resourceId string, in *PermissionsSetIn) error {
	path := fmt.Sprintf("/v1/organization/%s/permissions/%s/%s", url.PathEscape(organizationId), url.PathEscape(string(resourceType)), url.PathEscape(resourceId))
	_, err := h.doer.Do(ctx, "PermissionsSet", "PUT", path, in)
	return err
}
func (h *OrganizationHandler) PermissionsUpdate(ctx context.Context, organizationId string, resourceType ResourceType, resourceId string, in *PermissionsUpdateIn) error {
	path := fmt.Sprintf("/v1/organization/%s/permissions/%s/%s", url.PathEscape(organizationId), url.PathEscape(string(resourceType)), url.PathEscape(resourceId))
	_, err := h.doer.Do(ctx, "PermissionsUpdate", "PATCH", path, in)
	return err
}
func (h *OrganizationHandler) UserOrganizationCreate(ctx context.Context, in *UserOrganizationCreateIn) (*UserOrganizationCreateOut, error) {
	path := fmt.Sprintf("/v1/organizations")
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
	path := fmt.Sprintf("/v1/organizations")
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

type AddresseOut struct {
	AddressId      string    `json:"address_id"`              // Address ID
	AddressLines   []string  `json:"address_lines,omitempty"` // Address Lines
	City           *string   `json:"city,omitempty"`
	CompanyName    string    `json:"company_name"`              // Name of a company
	CountryCode    string    `json:"country_code"`              // Country Code
	CreateTime     time.Time `json:"create_time"`               // Create Time
	OrganizationId *string   `json:"organization_id,omitempty"` // Organization ID
	State          *string   `json:"state,omitempty"`
	UpdateTime     time.Time `json:"update_time"`        // Update Time
	ZipCode        *string   `json:"zip_code,omitempty"` // Zip Code
}
type DomainOut struct {
	ChallengeToken                string           `json:"challenge_token"`                  // Random string to be used for validation
	CreateTime                    time.Time        `json:"create_time"`                      // Time of creating the domain
	DomainId                      string           `json:"domain_id"`                        // ID of the domain
	DomainName                    string           `json:"domain_name"`                      // Name of the domain
	LinkedAuthenticationMethodIds []string         `json:"linked_authentication_method_ids"` // Linked Authentication Method Ids
	OrganizationId                string           `json:"organization_id"`                  // ID of the organization owning this domain
	State                         DomainStateType  `json:"state"`                            // An enumeration.
	VerificationType              VerificationType `json:"verification_type"`                // An enumeration.
}
type DomainStateType string

const (
	DomainStateTypeDeleted    DomainStateType = "deleted"
	DomainStateTypeUnverified DomainStateType = "unverified"
	DomainStateTypeVerified   DomainStateType = "verified"
)

func DomainStateTypeChoices() []string {
	return []string{"deleted", "unverified", "verified"}
}

type OperationType string

const (
	OperationTypeGrant  OperationType = "grant"
	OperationTypeRevoke OperationType = "revoke"
	OperationTypeSet    OperationType = "set"
)

func OperationTypeChoices() []string {
	return []string{"grant", "revoke", "set"}
}

// OrganizationAddressCreateIn OrganizationAddressCreateRequestBody
type OrganizationAddressCreateIn struct {
	AddressLines []string `json:"address_lines"` // Address lines
	City         string   `json:"city"`
	CompanyName  *string  `json:"company_name,omitempty"` // Name of a company
	CountryCode  string   `json:"country_code"`           // Country Code
	State        *string  `json:"state,omitempty"`
	ZipCode      *string  `json:"zip_code,omitempty"` // Zip Code
}

// OrganizationAddressCreateOut OrganizationAddressCreateResponse
type OrganizationAddressCreateOut struct {
	AddressId      string    `json:"address_id"`              // Address ID
	AddressLines   []string  `json:"address_lines,omitempty"` // Address Lines
	City           *string   `json:"city,omitempty"`
	CompanyName    string    `json:"company_name"`              // Name of a company
	CountryCode    string    `json:"country_code"`              // Country Code
	CreateTime     time.Time `json:"create_time"`               // Create Time
	OrganizationId *string   `json:"organization_id,omitempty"` // Organization ID
	State          *string   `json:"state,omitempty"`
	UpdateTime     time.Time `json:"update_time"`        // Update Time
	ZipCode        *string   `json:"zip_code,omitempty"` // Zip Code
}

// OrganizationAddressGetOut OrganizationAddressGetResponse
type OrganizationAddressGetOut struct {
	AddressId      string    `json:"address_id"`              // Address ID
	AddressLines   []string  `json:"address_lines,omitempty"` // Address Lines
	City           *string   `json:"city,omitempty"`
	CompanyName    string    `json:"company_name"`              // Name of a company
	CountryCode    string    `json:"country_code"`              // Country Code
	CreateTime     time.Time `json:"create_time"`               // Create Time
	OrganizationId *string   `json:"organization_id,omitempty"` // Organization ID
	State          *string   `json:"state,omitempty"`
	UpdateTime     time.Time `json:"update_time"`        // Update Time
	ZipCode        *string   `json:"zip_code,omitempty"` // Zip Code
}

// OrganizationAddressUpdateIn OrganizationAddressUpdateRequestBody
type OrganizationAddressUpdateIn struct {
	AddressLines *[]string `json:"address_lines,omitempty"` // Address Lines
	City         *string   `json:"city,omitempty"`
	CompanyName  *string   `json:"company_name,omitempty"` // Name of a company
	CountryCode  *string   `json:"country_code,omitempty"` // Country Code
	State        *string   `json:"state,omitempty"`
	ZipCode      *string   `json:"zip_code,omitempty"` // Zip Code
}

// OrganizationAddressUpdateOut OrganizationAddressUpdateResponse
type OrganizationAddressUpdateOut struct {
	AddressId      string    `json:"address_id"`              // Address ID
	AddressLines   []string  `json:"address_lines,omitempty"` // Address Lines
	City           *string   `json:"city,omitempty"`
	CompanyName    string    `json:"company_name"`              // Name of a company
	CountryCode    string    `json:"country_code"`              // Country Code
	CreateTime     time.Time `json:"create_time"`               // Create Time
	OrganizationId *string   `json:"organization_id,omitempty"` // Organization ID
	State          *string   `json:"state,omitempty"`
	UpdateTime     time.Time `json:"update_time"`        // Update Time
	ZipCode        *string   `json:"zip_code,omitempty"` // Zip Code
}

// OrganizationAuthDomainLinkIn OrganizationAuthDomainLinkRequestBody
type OrganizationAuthDomainLinkIn struct {
	DomainId string `json:"domain_id"` // ID of the domain
}

// OrganizationAuthenticationConfigGetOut OrganizationAuthenticationConfigGetResponse
type OrganizationAuthenticationConfigGetOut struct {
	OauthEnabled                           *bool `json:"oauth_enabled,omitempty"`                               // Organization users are able to use OAuth authentication.
	PasswordAuthEnabled                    *bool `json:"password_auth_enabled,omitempty"`                       // Organization users are able to use password authentication.
	PersonalTokensEnabled                  *bool `json:"personal_tokens_enabled,omitempty"`                     // Organization users can use their personal tokens to access the organization through the Aiven API or other applications.
	PersonalTokensRequireAllowedAuthMethod *bool `json:"personal_tokens_require_allowed_auth_method,omitempty"` // Organization users are able to use personal tokens that were generated from one of the allowed authentication methods.
	SamlAllowExternal                      *bool `json:"saml_allow_external,omitempty"`                         // Organization users are able to use SAML authentication of other organizations.
	SamlEnabled                            *bool `json:"saml_enabled,omitempty"`                                // Organization users are able to use SAML authentication.
	TwoFactorRequired                      *bool `json:"two_factor_required,omitempty"`                         // 2FA is required to access resources in this organization.
}

// OrganizationAuthenticationConfigUpdateIn OrganizationAuthenticationConfigUpdateRequestBody
type OrganizationAuthenticationConfigUpdateIn struct {
	OauthEnabled                           *bool `json:"oauth_enabled,omitempty"`                               // Organization users are able to use OAuth authentication.
	PasswordAuthEnabled                    *bool `json:"password_auth_enabled,omitempty"`                       // Organization users are able to use password authentication.
	PersonalTokensEnabled                  *bool `json:"personal_tokens_enabled,omitempty"`                     // Organization users can use their personal tokens to access the organization through the Aiven API or other applications.
	PersonalTokensRequireAllowedAuthMethod *bool `json:"personal_tokens_require_allowed_auth_method,omitempty"` // Organization users are able to use personal tokens that were generated from one of the allowed authentication methods.
	SamlAllowExternal                      *bool `json:"saml_allow_external,omitempty"`                         // Organization users are able to use SAML authentication of other organizations.
	SamlEnabled                            *bool `json:"saml_enabled,omitempty"`                                // Organization users are able to use SAML authentication.
	TwoFactorRequired                      *bool `json:"two_factor_required,omitempty"`                         // 2FA is required to access resources in this organization.
}

// OrganizationAuthenticationConfigUpdateOut OrganizationAuthenticationConfigUpdateResponse
type OrganizationAuthenticationConfigUpdateOut struct {
	OauthEnabled                           *bool `json:"oauth_enabled,omitempty"`                               // Organization users are able to use OAuth authentication.
	PasswordAuthEnabled                    *bool `json:"password_auth_enabled,omitempty"`                       // Organization users are able to use password authentication.
	PersonalTokensEnabled                  *bool `json:"personal_tokens_enabled,omitempty"`                     // Organization users can use their personal tokens to access the organization through the Aiven API or other applications.
	PersonalTokensRequireAllowedAuthMethod *bool `json:"personal_tokens_require_allowed_auth_method,omitempty"` // Organization users are able to use personal tokens that were generated from one of the allowed authentication methods.
	SamlAllowExternal                      *bool `json:"saml_allow_external,omitempty"`                         // Organization users are able to use SAML authentication of other organizations.
	SamlEnabled                            *bool `json:"saml_enabled,omitempty"`                                // Organization users are able to use SAML authentication.
	TwoFactorRequired                      *bool `json:"two_factor_required,omitempty"`                         // 2FA is required to access resources in this organization.
}

// OrganizationGetOut OrganizationGetResponse
type OrganizationGetOut struct {
	AccountId                    string    `json:"account_id"`                                 // Account ID of the organization's root unit
	CreateTime                   time.Time `json:"create_time"`                                // Time of creating the organization
	DefaultGovernanceUserGroupId *string   `json:"default_governance_user_group_id,omitempty"` // Default governance user group ID
	OrganizationId               string    `json:"organization_id"`                            // Organization's ID
	OrganizationName             string    `json:"organization_name"`                          // Organization's name
	Tier                         TierType  `json:"tier"`                                       // An enumeration.
	UpdateTime                   time.Time `json:"update_time"`                                // Time of the organization's latest update
}
type OrganizationOut struct {
	AccountId                    string    `json:"account_id"`                                 // Account ID of the organization's root unit
	CreateTime                   time.Time `json:"create_time"`                                // Time of creating the organization
	DefaultGovernanceUserGroupId *string   `json:"default_governance_user_group_id,omitempty"` // Default governance user group ID
	OrganizationId               string    `json:"organization_id"`                            // Organization's ID
	OrganizationName             string    `json:"organization_name"`                          // Organization's name
	Tier                         TierType  `json:"tier"`                                       // An enumeration.
	UpdateTime                   time.Time `json:"update_time"`                                // Time of the organization's latest update
}

// OrganizationUpdateIn OrganizationUpdateRequestBody
type OrganizationUpdateIn struct {
	Name *string  `json:"name,omitempty"` // New name of the organization
	Tier TierType `json:"tier,omitempty"` // An enumeration.
}

// OrganizationUpdateOut OrganizationUpdateResponse
type OrganizationUpdateOut struct {
	AccountId                    string    `json:"account_id"`                                 // Account ID of the organization's root unit
	CreateTime                   time.Time `json:"create_time"`                                // Time of creating the organization
	DefaultGovernanceUserGroupId *string   `json:"default_governance_user_group_id,omitempty"` // Default governance user group ID
	OrganizationId               string    `json:"organization_id"`                            // Organization's ID
	OrganizationName             string    `json:"organization_name"`                          // Organization's name
	Tier                         TierType  `json:"tier"`                                       // An enumeration.
	UpdateTime                   time.Time `json:"update_time"`                                // Time of the organization's latest update
}
type PermissionIn struct {
	Permissions   []string      `json:"permissions"`    // List of roles
	PrincipalId   string        `json:"principal_id"`   // ID of the principal
	PrincipalType PrincipalType `json:"principal_type"` // An enumeration.
}
type PermissionOut struct {
	CreateTime    time.Time     `json:"create_time"`    // Create Time
	Permissions   []string      `json:"permissions"`    // List of roles
	PrincipalId   string        `json:"principal_id"`   // ID of the principal
	PrincipalType PrincipalType `json:"principal_type"` // An enumeration.
	UpdateTime    time.Time     `json:"update_time"`    // Update Time
}

// PermissionsSetIn PermissionsSetRequestBody
type PermissionsSetIn struct {
	Permissions []PermissionIn `json:"permissions"` // List of roles to set
}

// PermissionsUpdateIn PermissionsUpdateRequestBody
type PermissionsUpdateIn struct {
	Operation   OperationType  `json:"operation,omitempty"` // An enumeration.
	Permissions []PermissionIn `json:"permissions"`         // List of roles to adjust
}
type PrincipalType string

const (
	PrincipalTypeUser      PrincipalType = "user"
	PrincipalTypeUserGroup PrincipalType = "user_group"
)

func PrincipalTypeChoices() []string {
	return []string{"user", "user_group"}
}

type ResourceType string

const (
	ResourceTypeOrganization     ResourceType = "organization"
	ResourceTypeOrganizationUnit ResourceType = "organization_unit"
	ResourceTypeProject          ResourceType = "project"
)

func ResourceTypeChoices() []string {
	return []string{"organization", "organization_unit", "project"}
}

type TierType string

const (
	TierTypeBusiness TierType = "business"
	TierTypePersonal TierType = "personal"
)

func TierTypeChoices() []string {
	return []string{"business", "personal"}
}

// UserOrganizationCreateIn UserOrganizationCreateRequestBody
type UserOrganizationCreateIn struct {
	OrganizationName      string   `json:"organization_name"`                  // Organization's name
	PrimaryBillingGroupId *string  `json:"primary_billing_group_id,omitempty"` // Billing group ID
	Tier                  TierType `json:"tier"`                               // An enumeration.
}

// UserOrganizationCreateOut UserOrganizationCreateResponse
type UserOrganizationCreateOut struct {
	AccountId                    string    `json:"account_id"`                                 // Account ID of the organization's root unit
	CreateTime                   time.Time `json:"create_time"`                                // Time of creating the organization
	DefaultGovernanceUserGroupId *string   `json:"default_governance_user_group_id,omitempty"` // Default governance user group ID
	OrganizationId               string    `json:"organization_id"`                            // Organization's ID
	OrganizationName             string    `json:"organization_name"`                          // Organization's name
	Tier                         TierType  `json:"tier"`                                       // An enumeration.
	UpdateTime                   time.Time `json:"update_time"`                                // Time of the organization's latest update
}
type VerificationType string

const (
	VerificationTypeDns  VerificationType = "dns"
	VerificationTypeHttp VerificationType = "http"
)

func VerificationTypeChoices() []string {
	return []string{"dns", "http"}
}

// organizationAddressListOut OrganizationAddressListResponse
type organizationAddressListOut struct {
	Addresses []AddresseOut `json:"addresses"` // Addresses
}

// organizationAuthDomainListOut OrganizationAuthDomainListResponse
type organizationAuthDomainListOut struct {
	Domains []DomainOut `json:"domains"` // List of domains for the organization
}

// permissionsGetOut PermissionsGetResponse
type permissionsGetOut struct {
	Permissions []PermissionOut `json:"permissions"` // List of roles
}

// userOrganizationsListOut UserOrganizationsListResponse
type userOrganizationsListOut struct {
	Organizations []OrganizationOut `json:"organizations"` // Organizations
}
