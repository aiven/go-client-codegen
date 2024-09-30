// Code generated by Aiven. DO NOT EDIT.

package billinggroup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// BillingGroupCreate create a billing group
	// POST /v1/billing-group
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupCreate
	BillingGroupCreate(ctx context.Context, in *BillingGroupCreateIn) (*BillingGroupCreateOut, error)

	// BillingGroupCreditsClaim claim a credit code
	// POST /v1/billing-group/{billing_group_id}/credits
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupCreditsClaim
	BillingGroupCreditsClaim(ctx context.Context, billingGroupId string, in *BillingGroupCreditsClaimIn) (*BillingGroupCreditsClaimOut, error)

	// BillingGroupCreditsList list billing group credits
	// GET /v1/billing-group/{billing_group_id}/credits
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupCreditsList
	BillingGroupCreditsList(ctx context.Context, billingGroupId string) ([]CreditOut, error)

	// BillingGroupDelete delete billing group
	// DELETE /v1/billing-group/{billing_group_id}
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupDelete
	BillingGroupDelete(ctx context.Context, billingGroupId string) error

	// BillingGroupEventList list billing group events
	// GET /v1/billing-group/{billing_group_id}/events
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupEventList
	BillingGroupEventList(ctx context.Context, billingGroupId string) ([]EventOut, error)

	// BillingGroupGet get billing group details
	// GET /v1/billing-group/{billing_group_id}
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupGet
	BillingGroupGet(ctx context.Context, billingGroupId string) (*BillingGroupGetOut, error)

	// BillingGroupInvoiceLinesList get invoice lines for a single invoice
	// GET /v1/billing-group/{billing_group_id}/invoice/{invoice_number}/lines
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupInvoiceLinesList
	BillingGroupInvoiceLinesList(ctx context.Context, billingGroupId string, invoiceNumber string) ([]LineOut, error)

	// BillingGroupInvoiceList get invoices generated for billing group
	// GET /v1/billing-group/{billing_group_id}/invoice
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupInvoiceList
	BillingGroupInvoiceList(ctx context.Context, billingGroupId string) ([]InvoiceOut, error)

	// BillingGroupList list billing groups
	// GET /v1/billing-group
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupList
	BillingGroupList(ctx context.Context) ([]BillingGroupOut, error)

	// BillingGroupProjectAssign assign project to billing group
	// POST /v1/billing-group/{billing_group_id}/project-assign/{project}
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupProjectAssign
	BillingGroupProjectAssign(ctx context.Context, billingGroupId string, project string) error

	// BillingGroupProjectList get projects assigned to billing group
	// GET /v1/billing-group/{billing_group_id}/projects
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupProjectList
	BillingGroupProjectList(ctx context.Context, billingGroupId string) ([]ProjectOut, error)

	// BillingGroupProjectsAssign assign projects to billing group
	// POST /v1/billing-group/{billing_group_id}/projects-assign
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupProjectsAssign
	BillingGroupProjectsAssign(ctx context.Context, billingGroupId string, in *BillingGroupProjectsAssignIn) error

	// BillingGroupUpdate update billing group
	// PUT /v1/billing-group/{billing_group_id}
	// https://api.aiven.io/doc/#tag/BillingGroup/operation/BillingGroupUpdate
	BillingGroupUpdate(ctx context.Context, billingGroupId string, in *BillingGroupUpdateIn) (*BillingGroupUpdateOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) BillingGroupHandler {
	return BillingGroupHandler{doer}
}

type BillingGroupHandler struct {
	doer doer
}

func (h *BillingGroupHandler) BillingGroupCreate(ctx context.Context, in *BillingGroupCreateIn) (*BillingGroupCreateOut, error) {
	path := fmt.Sprintf("/v1/billing-group")
	b, err := h.doer.Do(ctx, "BillingGroupCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.BillingGroup, nil
}
func (h *BillingGroupHandler) BillingGroupCreditsClaim(ctx context.Context, billingGroupId string, in *BillingGroupCreditsClaimIn) (*BillingGroupCreditsClaimOut, error) {
	path := fmt.Sprintf("/v1/billing-group/%s/credits", url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "BillingGroupCreditsClaim", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupCreditsClaimOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Credit, nil
}
func (h *BillingGroupHandler) BillingGroupCreditsList(ctx context.Context, billingGroupId string) ([]CreditOut, error) {
	path := fmt.Sprintf("/v1/billing-group/%s/credits", url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "BillingGroupCreditsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupCreditsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Credits, nil
}
func (h *BillingGroupHandler) BillingGroupDelete(ctx context.Context, billingGroupId string) error {
	path := fmt.Sprintf("/v1/billing-group/%s", url.PathEscape(billingGroupId))
	_, err := h.doer.Do(ctx, "BillingGroupDelete", "DELETE", path, nil)
	return err
}
func (h *BillingGroupHandler) BillingGroupEventList(ctx context.Context, billingGroupId string) ([]EventOut, error) {
	path := fmt.Sprintf("/v1/billing-group/%s/events", url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "BillingGroupEventList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupEventListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Events, nil
}
func (h *BillingGroupHandler) BillingGroupGet(ctx context.Context, billingGroupId string) (*BillingGroupGetOut, error) {
	path := fmt.Sprintf("/v1/billing-group/%s", url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "BillingGroupGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.BillingGroup, nil
}
func (h *BillingGroupHandler) BillingGroupInvoiceLinesList(ctx context.Context, billingGroupId string, invoiceNumber string) ([]LineOut, error) {
	path := fmt.Sprintf("/v1/billing-group/%s/invoice/%s/lines", url.PathEscape(billingGroupId), url.PathEscape(invoiceNumber))
	b, err := h.doer.Do(ctx, "BillingGroupInvoiceLinesList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupInvoiceLinesListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Lines, nil
}
func (h *BillingGroupHandler) BillingGroupInvoiceList(ctx context.Context, billingGroupId string) ([]InvoiceOut, error) {
	path := fmt.Sprintf("/v1/billing-group/%s/invoice", url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "BillingGroupInvoiceList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupInvoiceListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Invoices, nil
}
func (h *BillingGroupHandler) BillingGroupList(ctx context.Context) ([]BillingGroupOut, error) {
	path := fmt.Sprintf("/v1/billing-group")
	b, err := h.doer.Do(ctx, "BillingGroupList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.BillingGroups, nil
}
func (h *BillingGroupHandler) BillingGroupProjectAssign(ctx context.Context, billingGroupId string, project string) error {
	path := fmt.Sprintf("/v1/billing-group/%s/project-assign/%s", url.PathEscape(billingGroupId), url.PathEscape(project))
	_, err := h.doer.Do(ctx, "BillingGroupProjectAssign", "POST", path, nil)
	return err
}
func (h *BillingGroupHandler) BillingGroupProjectList(ctx context.Context, billingGroupId string) ([]ProjectOut, error) {
	path := fmt.Sprintf("/v1/billing-group/%s/projects", url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "BillingGroupProjectList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupProjectListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Projects, nil
}
func (h *BillingGroupHandler) BillingGroupProjectsAssign(ctx context.Context, billingGroupId string, in *BillingGroupProjectsAssignIn) error {
	path := fmt.Sprintf("/v1/billing-group/%s/projects-assign", url.PathEscape(billingGroupId))
	_, err := h.doer.Do(ctx, "BillingGroupProjectsAssign", "POST", path, in)
	return err
}
func (h *BillingGroupHandler) BillingGroupUpdate(ctx context.Context, billingGroupId string, in *BillingGroupUpdateIn) (*BillingGroupUpdateOut, error) {
	path := fmt.Sprintf("/v1/billing-group/%s", url.PathEscape(billingGroupId))
	b, err := h.doer.Do(ctx, "BillingGroupUpdate", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(billingGroupUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.BillingGroup, nil
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
	Email string `json:"email"` // User email address
}
type BillingEmailOut struct {
	Email string `json:"email"` // User email address
}

// BillingGroupCreateIn BillingGroupCreateRequestBody
type BillingGroupCreateIn struct {
	AccountId            *string             `json:"account_id,omitempty"`              // Account ID
	AddressLines         *[]string           `json:"address_lines,omitempty"`           // Address lines
	BillingCurrency      BillingCurrencyType `json:"billing_currency,omitempty"`        // Billing currency
	BillingEmails        *[]BillingEmailIn   `json:"billing_emails,omitempty"`          // List of project billing email addresses
	BillingExtraText     *string             `json:"billing_extra_text,omitempty"`      // Extra text to be included in all project invoices, e.g. purchase order or cost center number
	BillingGroupName     string              `json:"billing_group_name"`                // Billing group name
	CardId               *string             `json:"card_id,omitempty"`                 // Credit card ID
	City                 *string             `json:"city,omitempty"`                    // Address city
	Company              *string             `json:"company,omitempty"`                 // Name of a company
	CopyFromBillingGroup *string             `json:"copy_from_billing_group,omitempty"` // Billing group ID
	CountryCode          *string             `json:"country_code,omitempty"`            // Two letter country code for billing country
	State                *string             `json:"state,omitempty"`                   // Address state
	VatId                *string             `json:"vat_id,omitempty"`                  // EU VAT Identification Number
	ZipCode              *string             `json:"zip_code,omitempty"`                // Address zip code
}

// BillingGroupCreateOut Billing group information
type BillingGroupCreateOut struct {
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

// BillingGroupCreditsClaimIn BillingGroupCreditsClaimRequestBody
type BillingGroupCreditsClaimIn struct {
	Code string `json:"code"` // Credit code
}

// BillingGroupCreditsClaimOut Assigned credit
type BillingGroupCreditsClaimOut struct {
	Code           *string    `json:"code,omitempty"`            // Credit code
	ExpireTime     *time.Time `json:"expire_time,omitempty"`     // Timestamp in ISO 8601 format, always in UTC
	RemainingValue *string    `json:"remaining_value,omitempty"` // Remaining credit value
	StartTime      *time.Time `json:"start_time,omitempty"`      // Timestamp in ISO 8601 format, always in UTC
	Type           CreditType `json:"type,omitempty"`            // Credit type
	Value          *string    `json:"value,omitempty"`           // Original credit value, or for expired credits, the consumed credit value
}

// BillingGroupGetOut Billing group information
type BillingGroupGetOut struct {
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
type BillingGroupOut struct {
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

// BillingGroupProjectsAssignIn BillingGroupProjectsAssignRequestBody
type BillingGroupProjectsAssignIn struct {
	ProjectsNames []string `json:"projects_names"` // Projects names
}
type BillingGroupStateType string

const (
	BillingGroupStateTypeActive  BillingGroupStateType = "active"
	BillingGroupStateTypeDeleted BillingGroupStateType = "deleted"
)

func BillingGroupStateTypeChoices() []string {
	return []string{"active", "deleted"}
}

// BillingGroupUpdateIn BillingGroupUpdateRequestBody
type BillingGroupUpdateIn struct {
	AccountId        *string             `json:"account_id,omitempty"`         // Account ID
	AddressLines     *[]string           `json:"address_lines,omitempty"`      // Address lines
	BillingCurrency  BillingCurrencyType `json:"billing_currency,omitempty"`   // Billing currency
	BillingEmails    *[]BillingEmailIn   `json:"billing_emails,omitempty"`     // List of project billing email addresses
	BillingExtraText *string             `json:"billing_extra_text,omitempty"` // Extra text to be included in all project invoices, e.g. purchase order or cost center number
	BillingGroupName *string             `json:"billing_group_name,omitempty"` // Billing group name
	CardId           *string             `json:"card_id,omitempty"`            // Credit card ID
	City             *string             `json:"city,omitempty"`               // Address city
	Company          *string             `json:"company,omitempty"`            // Name of a company
	CountryCode      *string             `json:"country_code,omitempty"`       // Two letter country code for billing country
	State            *string             `json:"state,omitempty"`              // Address state
	VatId            *string             `json:"vat_id,omitempty"`             // EU VAT Identification Number
	ZipCode          *string             `json:"zip_code,omitempty"`           // Address zip code
}

// BillingGroupUpdateOut Billing group information
type BillingGroupUpdateOut struct {
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
type CreditOut struct {
	Code           *string    `json:"code,omitempty"`            // Credit code
	ExpireTime     *time.Time `json:"expire_time,omitempty"`     // Timestamp in ISO 8601 format, always in UTC
	RemainingValue *string    `json:"remaining_value,omitempty"` // Remaining credit value
	StartTime      *time.Time `json:"start_time,omitempty"`      // Timestamp in ISO 8601 format, always in UTC
	Type           CreditType `json:"type,omitempty"`            // Credit type
	Value          *string    `json:"value,omitempty"`           // Original credit value, or for expired credits, the consumed credit value
}
type CreditType string

const (
	CreditTypeDiscount    CreditType = "discount"
	CreditTypeEmployee    CreditType = "employee"
	CreditTypeEvaluation  CreditType = "evaluation"
	CreditTypeInternal    CreditType = "internal"
	CreditTypeOther       CreditType = "other"
	CreditTypeOutage      CreditType = "outage"
	CreditTypePartner     CreditType = "partner"
	CreditTypePromotion   CreditType = "promotion"
	CreditTypePurchase    CreditType = "purchase"
	CreditTypeReferral    CreditType = "referral"
	CreditTypeSponsorship CreditType = "sponsorship"
	CreditTypeTrial       CreditType = "trial"
	CreditTypeTrialOver   CreditType = "trial_over"
)

func CreditTypeChoices() []string {
	return []string{"discount", "employee", "evaluation", "internal", "other", "outage", "partner", "promotion", "purchase", "referral", "sponsorship", "trial", "trial_over"}
}

type CurrencyType string

const (
	CurrencyTypeAud CurrencyType = "AUD"
	CurrencyTypeCad CurrencyType = "CAD"
	CurrencyTypeChf CurrencyType = "CHF"
	CurrencyTypeDkk CurrencyType = "DKK"
	CurrencyTypeEur CurrencyType = "EUR"
	CurrencyTypeGbp CurrencyType = "GBP"
	CurrencyTypeJpy CurrencyType = "JPY"
	CurrencyTypeNok CurrencyType = "NOK"
	CurrencyTypeNzd CurrencyType = "NZD"
	CurrencyTypeSek CurrencyType = "SEK"
	CurrencyTypeSgd CurrencyType = "SGD"
	CurrencyTypeUsd CurrencyType = "USD"
)

func CurrencyTypeChoices() []string {
	return []string{"AUD", "CAD", "CHF", "DKK", "EUR", "GBP", "JPY", "NOK", "NZD", "SEK", "SGD", "USD"}
}

type EventOut struct {
	Actor          *string    `json:"actor,omitempty"`            // Initiator of the event
	BillingGroupId *string    `json:"billing_group_id,omitempty"` // Billing group ID
	CreateTime     *time.Time `json:"create_time,omitempty"`      // Timestamp in ISO 8601 format, always in UTC
	EventDesc      *string    `json:"event_desc,omitempty"`       // Event description
	EventType      *string    `json:"event_type,omitempty"`       // Event type identifier
	LogEntryId     *int       `json:"log_entry_id,omitempty"`     // Entry ID
	ProjectId      *string    `json:"project_id,omitempty"`       // Identifier of a project
	ProjectName    *string    `json:"project_name,omitempty"`     // Project name
}
type InvoiceOut struct {
	BillingGroupId    string                `json:"billing_group_id"`       // Billing group ID
	BillingGroupName  string                `json:"billing_group_name"`     // Billing group name
	BillingGroupState BillingGroupStateType `json:"billing_group_state"`    // Billing group state
	Currency          CurrencyType          `json:"currency"`               // Billing currency
	DownloadCookie    string                `json:"download_cookie"`        // Authentication cookie for downloads
	GeneratedAt       *time.Time            `json:"generated_at,omitempty"` // The time when the invoice was generated
	InvoiceNumber     string                `json:"invoice_number"`         // Unique invoice reference code
	PeriodBegin       string                `json:"period_begin"`           // Period begin
	PeriodEnd         string                `json:"period_end"`             // Period end
	State             InvoiceStateType      `json:"state"`                  // State of this invoice
	TotalIncVat       string                `json:"total_inc_vat"`          // Total including taxes
	TotalVatZero      string                `json:"total_vat_zero"`         // Total excluding taxes
}
type InvoiceStateType string

const (
	InvoiceStateTypeAccrual                              InvoiceStateType = "accrual"
	InvoiceStateTypeConsolidated                         InvoiceStateType = "consolidated"
	InvoiceStateTypeDue                                  InvoiceStateType = "due"
	InvoiceStateTypeEstimate                             InvoiceStateType = "estimate"
	InvoiceStateTypeFailedCreditCardCharge               InvoiceStateType = "failed_credit_card_charge"
	InvoiceStateTypeFailedNoCreditCard                   InvoiceStateType = "failed_no_credit_card"
	InvoiceStateTypeMailed                               InvoiceStateType = "mailed"
	InvoiceStateTypeNoPaymentExpected                    InvoiceStateType = "no_payment_expected"
	InvoiceStateTypePaid                                 InvoiceStateType = "paid"
	InvoiceStateTypePartnerMetering                      InvoiceStateType = "partner_metering"
	InvoiceStateTypeUncollectible                        InvoiceStateType = "uncollectible"
	InvoiceStateTypeWaived                               InvoiceStateType = "waived"
	InvoiceStateTypeDueOnlyProjectChargesCalculated      InvoiceStateType = "due_only_project_charges_calculated"
	InvoiceStateTypeEstimateOnlyProjectChargesCalculated InvoiceStateType = "estimate_only_project_charges_calculated"
)

func InvoiceStateTypeChoices() []string {
	return []string{"accrual", "consolidated", "due", "estimate", "failed_credit_card_charge", "failed_no_credit_card", "mailed", "no_payment_expected", "paid", "partner_metering", "uncollectible", "waived", "due_only_project_charges_calculated", "estimate_only_project_charges_calculated"}
}

type LineOut struct {
	CloudName            *string           `json:"cloud_name,omitempty"`              // Name of the cloud, if billed resource is associated with a cloud resource
	CommitmentName       *string           `json:"commitment_name,omitempty"`         // Name of the commitment which is referred to this invoice line
	Description          string            `json:"description"`                       // Human-readable short description of the invoice line
	LinePreDiscountLocal *string           `json:"line_pre_discount_local,omitempty"` // Pre-tax sum of invoice line, in local currency, before any discounts
	LineTotalLocal       *string           `json:"line_total_local,omitempty"`        // Pre-tax sum of invoice line, in the local currency configured for the invoice
	LineTotalUsd         string            `json:"line_total_usd"`                    // Pre-tax sum of invoice line, in USD
	LineType             LineType          `json:"line_type"`                         // Type of the invoice line
	LocalCurrency        *string           `json:"local_currency,omitempty"`          // Currency used for line_local_total
	ProjectName          *string           `json:"project_name,omitempty"`            // Name of the project this line is associated with, if any
	ServiceName          *string           `json:"service_name,omitempty"`            // Name of the service, if invoice line is for service use
	ServicePlan          *string           `json:"service_plan,omitempty"`            // Service plan name, if invoice line is for service use
	ServiceType          ServiceType       `json:"service_type,omitempty"`            // Service type, if invoice line is for service use
	Tags                 map[string]string `json:"tags,omitempty"`                    // Billing tags
	TimestampBegin       *string           `json:"timestamp_begin,omitempty"`         // Begin timestamp of the billed time period, for resources billed by time
	TimestampEnd         *string           `json:"timestamp_end,omitempty"`           // End timestamp of the billed time period, for resources billed by time
}
type LineType string

const (
	LineTypeCommitmentFee     LineType = "commitment_fee"
	LineTypeCreditConsumption LineType = "credit_consumption"
	LineTypeExtraCharge       LineType = "extra_charge"
	LineTypeMultiplier        LineType = "multiplier"
	LineTypeOtherEvent        LineType = "other_event"
	LineTypeProPlatformCharge LineType = "pro_platform_charge"
	LineTypeRounding          LineType = "rounding"
	LineTypeServiceCharge     LineType = "service_charge"
	LineTypeSupportCharge     LineType = "support_charge"
)

func LineTypeChoices() []string {
	return []string{"commitment_fee", "credit_consumption", "extra_charge", "multiplier", "other_event", "pro_platform_charge", "rounding", "service_charge", "support_charge"}
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
	AvailableCredits string `json:"available_credits"` // Available credits
	EstimatedBalance string `json:"estimated_balance"` // Estimated balance
	ProjectName      string `json:"project_name"`      // Project name
}
type ServiceType string

const (
	ServiceTypeAlertmanager     ServiceType = "alertmanager"
	ServiceTypeCassandra        ServiceType = "cassandra"
	ServiceTypeClickhouse       ServiceType = "clickhouse"
	ServiceTypeDragonfly        ServiceType = "dragonfly"
	ServiceTypeElasticsearch    ServiceType = "elasticsearch"
	ServiceTypeFlink            ServiceType = "flink"
	ServiceTypeGrafana          ServiceType = "grafana"
	ServiceTypeInfluxdb         ServiceType = "influxdb"
	ServiceTypeKafka            ServiceType = "kafka"
	ServiceTypeKafkaConnect     ServiceType = "kafka_connect"
	ServiceTypeKafkaMirrormaker ServiceType = "kafka_mirrormaker"
	ServiceTypeM3Aggregator     ServiceType = "m3aggregator"
	ServiceTypeM3Db             ServiceType = "m3db"
	ServiceTypeMysql            ServiceType = "mysql"
	ServiceTypeOpensearch       ServiceType = "opensearch"
	ServiceTypeParca            ServiceType = "parca"
	ServiceTypePg               ServiceType = "pg"
	ServiceTypeRedis            ServiceType = "redis"
	ServiceTypeStresstester     ServiceType = "stresstester"
	ServiceTypeSw               ServiceType = "sw"
	ServiceTypeThanos           ServiceType = "thanos"
	ServiceTypeThanoscompactor  ServiceType = "thanoscompactor"
	ServiceTypeThanosquery      ServiceType = "thanosquery"
	ServiceTypeThanosreceiver   ServiceType = "thanosreceiver"
	ServiceTypeThanosruler      ServiceType = "thanosruler"
	ServiceTypeThanosstore      ServiceType = "thanosstore"
	ServiceTypeValkey           ServiceType = "valkey"
	ServiceTypeVector           ServiceType = "vector"
	ServiceTypeVmalert          ServiceType = "vmalert"
	ServiceTypeWarpstream       ServiceType = "warpstream"
)

func ServiceTypeChoices() []string {
	return []string{"alertmanager", "cassandra", "clickhouse", "dragonfly", "elasticsearch", "flink", "grafana", "influxdb", "kafka", "kafka_connect", "kafka_mirrormaker", "m3aggregator", "m3db", "mysql", "opensearch", "parca", "pg", "redis", "stresstester", "sw", "thanos", "thanoscompactor", "thanosquery", "thanosreceiver", "thanosruler", "thanosstore", "valkey", "vector", "vmalert", "warpstream"}
}

// billingGroupCreateOut BillingGroupCreateResponse
type billingGroupCreateOut struct {
	BillingGroup BillingGroupCreateOut `json:"billing_group"` // Billing group information
}

// billingGroupCreditsClaimOut BillingGroupCreditsClaimResponse
type billingGroupCreditsClaimOut struct {
	Credit BillingGroupCreditsClaimOut `json:"credit"` // Assigned credit
}

// billingGroupCreditsListOut BillingGroupCreditsListResponse
type billingGroupCreditsListOut struct {
	Credits []CreditOut `json:"credits"` // List of credits assigned to a billing group
}

// billingGroupEventListOut BillingGroupEventListResponse
type billingGroupEventListOut struct {
	Events []EventOut `json:"events"` // List of events related to a billing group
}

// billingGroupGetOut BillingGroupGetResponse
type billingGroupGetOut struct {
	BillingGroup BillingGroupGetOut `json:"billing_group"` // Billing group information
}

// billingGroupInvoiceLinesListOut BillingGroupInvoiceLinesListResponse
type billingGroupInvoiceLinesListOut struct {
	Lines []LineOut `json:"lines,omitempty"` // List of invoice lines
}

// billingGroupInvoiceListOut BillingGroupInvoiceListResponse
type billingGroupInvoiceListOut struct {
	Invoices []InvoiceOut `json:"invoices"` // List of billing group invoices
}

// billingGroupListOut BillingGroupListResponse
type billingGroupListOut struct {
	BillingGroups []BillingGroupOut `json:"billing_groups"` // List of billing groups
}

// billingGroupProjectListOut BillingGroupProjectListResponse
type billingGroupProjectListOut struct {
	Projects []ProjectOut `json:"projects"` // List of projects assigned to billing group
}

// billingGroupUpdateOut BillingGroupUpdateResponse
type billingGroupUpdateOut struct {
	BillingGroup BillingGroupUpdateOut `json:"billing_group"` // Billing group information
}
