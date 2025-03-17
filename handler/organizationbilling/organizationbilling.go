// Code generated by Aiven. DO NOT EDIT.

package organizationbilling

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// OrganizationBillingGroupCreate [EXPERIMENTAL] Create an organization billing group
	// POST /v1/organization/{organization_id}/billing-groups
	// https://api.aiven.io/doc/#tag/OrganizationBillingGroup/operation/OrganizationBillingGroupCreate
	OrganizationBillingGroupCreate(ctx context.Context, organizationId string, in *OrganizationBillingGroupCreateIn) (*OrganizationBillingGroupCreateOut, error)

	// OrganizationBillingGroupDelete [EXPERIMENTAL] Delete an organization billing group
	// DELETE /v1/organization/{organization_id}/billing-group/{billing_group_id}
	// https://api.aiven.io/doc/#tag/OrganizationBillingGroup/operation/OrganizationBillingGroupDelete
	OrganizationBillingGroupDelete(ctx context.Context, organizationId string, billingGroupId string) error

	// OrganizationBillingGroupGet [EXPERIMENTAL] Get organization billing group details
	// GET /v1/organization/{organization_id}/billing-group/{billing_group_id}
	// https://api.aiven.io/doc/#tag/OrganizationBillingGroup/operation/OrganizationBillingGroupGet
	OrganizationBillingGroupGet(ctx context.Context, organizationId string, billingGroupId string) (*OrganizationBillingGroupGetOut, error)

	// OrganizationBillingGroupList [EXPERIMENTAL] List billing groups in an organization
	// GET /v1/organization/{organization_id}/billing-group
	// https://api.aiven.io/doc/#tag/OrganizationBillingGroup/operation/OrganizationBillingGroupList
	OrganizationBillingGroupList(ctx context.Context, organizationId string) ([]BillingGroupOut, error)

	// OrganizationBillingGroupUpdate [EXPERIMENTAL] Update organization billing group details
	// PUT /v1/organization/{organization_id}/billing-group/{billing_group_id}
	// https://api.aiven.io/doc/#tag/OrganizationBillingGroup/operation/OrganizationBillingGroupUpdate
	OrganizationBillingGroupUpdate(ctx context.Context, organizationId string, billingGroupId string, in *OrganizationBillingGroupUpdateIn) (*OrganizationBillingGroupUpdateOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) OrganizationBillingHandler {
	return OrganizationBillingHandler{doer}
}

type OrganizationBillingHandler struct {
	doer doer
}

func (h *OrganizationBillingHandler) OrganizationBillingGroupCreate(ctx context.Context, organizationId string, in *OrganizationBillingGroupCreateIn) (*OrganizationBillingGroupCreateOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/billing-groups", url.PathEscape(organizationId))
	b, err := h.doer.Do(ctx, "OrganizationBillingGroupCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationBillingGroupCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationBillingHandler) OrganizationBillingGroupDelete(ctx context.Context, organizationId string, billingGroupId string) error {
	path := fmt.Sprintf("/v1/organization/%s/billing-group/%s", url.PathEscape(organizationId), url.PathEscape(billingGroupId))
	_, err := h.doer.Do(ctx, "OrganizationBillingGroupDelete", "DELETE", path, nil)
	return err
}
func (h *OrganizationBillingHandler) OrganizationBillingGroupGet(ctx context.Context, organizationId string, billingGroupId string) (*OrganizationBillingGroupGetOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/billing-group/%s", url.PathEscape(organizationId), url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "OrganizationBillingGroupGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationBillingGroupGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *OrganizationBillingHandler) OrganizationBillingGroupList(ctx context.Context, organizationId string) ([]BillingGroupOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/billing-group", url.PathEscape(organizationId))
	b, err := h.doer.Do(ctx, "OrganizationBillingGroupList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(organizationBillingGroupListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.BillingGroups, nil
}
func (h *OrganizationBillingHandler) OrganizationBillingGroupUpdate(ctx context.Context, organizationId string, billingGroupId string, in *OrganizationBillingGroupUpdateIn) (*OrganizationBillingGroupUpdateOut, error) {
	path := fmt.Sprintf("/v1/organization/%s/billing-group/%s", url.PathEscape(organizationId), url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "OrganizationBillingGroupUpdate", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(OrganizationBillingGroupUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type BillingContactEmailIn struct {
	Email string `json:"email"`
}
type BillingContactEmailOut struct {
	Email string `json:"email"`
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

type BillingEmailIn struct {
	Email string `json:"email"`
}
type BillingEmailOut struct {
	Email string `json:"email"`
}
type BillingGroupOut struct {
	BillingAddressId     string                   `json:"billing_address_id"`            // Billing address ID
	BillingContactEmails []BillingContactEmailOut `json:"billing_contact_emails"`        // List of billing contact emails
	BillingCurrency      BillingCurrencyType      `json:"billing_currency,omitempty"`    // Acceptable currencies for a billing group.
	BillingEmails        []BillingEmailOut        `json:"billing_emails"`                // List of billing contact emails
	BillingGroupId       string                   `json:"billing_group_id"`              // Billing group ID
	BillingGroupName     string                   `json:"billing_group_name"`            // Billing Group Name
	CustomInvoiceText    *string                  `json:"custom_invoice_text,omitempty"` // Extra billing text
	OrganizationId       string                   `json:"organization_id"`               // Organization ID
	PaymentMethodId      *string                  `json:"payment_method_id,omitempty"`   // Payment method ID
	ShippingAddressId    string                   `json:"shipping_address_id"`           // Shipping address ID
	VatId                *string                  `json:"vat_id,omitempty"`              // VAT ID
}

// OrganizationBillingGroupCreateIn OrganizationBillingGroupCreateRequestBody
type OrganizationBillingGroupCreateIn struct {
	BillingAddressId     string                  `json:"billing_address_id"`            // Billing address ID
	BillingContactEmails []BillingContactEmailIn `json:"billing_contact_emails"`        // List of billing contact emails
	BillingCurrency      BillingCurrencyType     `json:"billing_currency,omitempty"`    // Acceptable currencies for a billing group.
	BillingEmails        []BillingEmailIn        `json:"billing_emails"`                // List of billing contact emails
	BillingGroupName     string                  `json:"billing_group_name"`            // Billing Group Name
	CustomInvoiceText    *string                 `json:"custom_invoice_text,omitempty"` // Extra billing text
	PaymentMethodId      *string                 `json:"payment_method_id,omitempty"`   // Payment method ID
	ShippingAddressId    string                  `json:"shipping_address_id"`           // Shipping address ID
	VatId                *string                 `json:"vat_id,omitempty"`              // VAT ID
}

// OrganizationBillingGroupCreateOut OrganizationBillingGroupCreateResponse
type OrganizationBillingGroupCreateOut struct {
	BillingAddressId     string                   `json:"billing_address_id"`            // Billing address ID
	BillingContactEmails []BillingContactEmailOut `json:"billing_contact_emails"`        // List of billing contact emails
	BillingCurrency      BillingCurrencyType      `json:"billing_currency,omitempty"`    // Acceptable currencies for a billing group.
	BillingEmails        []BillingEmailOut        `json:"billing_emails"`                // List of billing contact emails
	BillingGroupId       string                   `json:"billing_group_id"`              // Billing group ID
	BillingGroupName     string                   `json:"billing_group_name"`            // Billing Group Name
	CustomInvoiceText    *string                  `json:"custom_invoice_text,omitempty"` // Extra billing text
	OrganizationId       string                   `json:"organization_id"`               // Organization ID
	PaymentMethodId      *string                  `json:"payment_method_id,omitempty"`   // Payment method ID
	ShippingAddressId    string                   `json:"shipping_address_id"`           // Shipping address ID
	VatId                *string                  `json:"vat_id,omitempty"`              // VAT ID
}

// OrganizationBillingGroupGetOut OrganizationBillingGroupGetResponse
type OrganizationBillingGroupGetOut struct {
	BillingAddressId     string                   `json:"billing_address_id"`            // Billing address ID
	BillingContactEmails []BillingContactEmailOut `json:"billing_contact_emails"`        // List of billing contact emails
	BillingCurrency      BillingCurrencyType      `json:"billing_currency,omitempty"`    // Acceptable currencies for a billing group.
	BillingEmails        []BillingEmailOut        `json:"billing_emails"`                // List of billing contact emails
	BillingGroupId       string                   `json:"billing_group_id"`              // Billing group ID
	BillingGroupName     string                   `json:"billing_group_name"`            // Billing Group Name
	CustomInvoiceText    *string                  `json:"custom_invoice_text,omitempty"` // Extra billing text
	OrganizationId       string                   `json:"organization_id"`               // Organization ID
	PaymentMethodId      *string                  `json:"payment_method_id,omitempty"`   // Payment method ID
	ShippingAddressId    string                   `json:"shipping_address_id"`           // Shipping address ID
	VatId                *string                  `json:"vat_id,omitempty"`              // VAT ID
}

// OrganizationBillingGroupUpdateIn OrganizationBillingGroupUpdateRequestBody
type OrganizationBillingGroupUpdateIn struct {
	BillingAddressId     *string                 `json:"billing_address_id,omitempty"`  // Billing address ID
	BillingContactEmails []BillingContactEmailIn `json:"billing_contact_emails"`        // List of billing contact emails
	BillingCurrency      BillingCurrencyType     `json:"billing_currency"`              // Acceptable currencies for a billing group.
	BillingEmails        []BillingEmailIn        `json:"billing_emails"`                // List of billing contact emails
	BillingGroupName     string                  `json:"billing_group_name"`            // Billing Group Name
	CustomInvoiceText    *string                 `json:"custom_invoice_text,omitempty"` // Extra billing text
	PaymentMethodId      *string                 `json:"payment_method_id,omitempty"`   // Payment method ID
	ShippingAddressId    *string                 `json:"shipping_address_id,omitempty"` // Shipping address ID
	VatId                *string                 `json:"vat_id,omitempty"`              // VAT ID
}

// OrganizationBillingGroupUpdateOut OrganizationBillingGroupUpdateResponse
type OrganizationBillingGroupUpdateOut struct {
	BillingAddressId     string                   `json:"billing_address_id"`            // Billing address ID
	BillingContactEmails []BillingContactEmailOut `json:"billing_contact_emails"`        // List of billing contact emails
	BillingCurrency      BillingCurrencyType      `json:"billing_currency,omitempty"`    // Acceptable currencies for a billing group.
	BillingEmails        []BillingEmailOut        `json:"billing_emails"`                // List of billing contact emails
	BillingGroupId       string                   `json:"billing_group_id"`              // Billing group ID
	BillingGroupName     string                   `json:"billing_group_name"`            // Billing Group Name
	CustomInvoiceText    *string                  `json:"custom_invoice_text,omitempty"` // Extra billing text
	OrganizationId       string                   `json:"organization_id"`               // Organization ID
	PaymentMethodId      *string                  `json:"payment_method_id,omitempty"`   // Payment method ID
	ShippingAddressId    string                   `json:"shipping_address_id"`           // Shipping address ID
	VatId                *string                  `json:"vat_id,omitempty"`              // VAT ID
}

// organizationBillingGroupListOut OrganizationBillingGroupListResponse
type organizationBillingGroupListOut struct {
	BillingGroups []BillingGroupOut `json:"billing_groups"` // A list of all billing groups belonging to the organization
}
