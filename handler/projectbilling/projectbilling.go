// Code generated by Aiven. DO NOT EDIT.

package projectbilling

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Handler interface {
	// InvoiceGet get a single invoice
	// GET /v1/invoices/{invoice_number}
	// https://api.aiven.io/doc/#tag/Billing/operation/InvoiceGet
	InvoiceGet(ctx context.Context, invoiceNumber string) (*InvoiceGetOut, error)

	// ProjectCreditsClaim claim a credit code
	// POST /v1/project/{project}/credits
	// https://api.aiven.io/doc/#tag/Project_Billing/operation/ProjectCreditsClaim
	ProjectCreditsClaim(ctx context.Context, project string, in *ProjectCreditsClaimIn) (*ProjectCreditsClaimOut, error)

	// ProjectCreditsList list project credits
	// GET /v1/project/{project}/credits
	// https://api.aiven.io/doc/#tag/Project_Billing/operation/ProjectCreditsList
	ProjectCreditsList(ctx context.Context, project string) ([]CreditOut, error)

	// ProjectInvoiceList list project invoices
	// GET /v1/project/{project}/invoice
	// https://api.aiven.io/doc/#tag/Project_Billing/operation/ProjectInvoiceList
	ProjectInvoiceList(ctx context.Context, project string) ([]InvoiceOut, error)
}

func NewHandler(doer doer) ProjectBillingHandler {
	return ProjectBillingHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type ProjectBillingHandler struct {
	doer doer
}

func (h *ProjectBillingHandler) InvoiceGet(ctx context.Context, invoiceNumber string) (*InvoiceGetOut, error) {
	path := fmt.Sprintf("/v1/invoices/%s", url.PathEscape(invoiceNumber))
	b, err := h.doer.Do(ctx, "InvoiceGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(invoiceGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Invoice, nil
}
func (h *ProjectBillingHandler) ProjectCreditsClaim(ctx context.Context, project string, in *ProjectCreditsClaimIn) (*ProjectCreditsClaimOut, error) {
	path := fmt.Sprintf("/v1/project/%s/credits", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ProjectCreditsClaim", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(projectCreditsClaimOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Credit, nil
}
func (h *ProjectBillingHandler) ProjectCreditsList(ctx context.Context, project string) ([]CreditOut, error) {
	path := fmt.Sprintf("/v1/project/%s/credits", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ProjectCreditsList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(projectCreditsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Credits, nil
}
func (h *ProjectBillingHandler) ProjectInvoiceList(ctx context.Context, project string) ([]InvoiceOut, error) {
	path := fmt.Sprintf("/v1/project/%s/invoice", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ProjectInvoiceList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(projectInvoiceListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Invoices, nil
}

type BillingGroupStateType string

const (
	BillingGroupStateTypeActive  BillingGroupStateType = "active"
	BillingGroupStateTypeDeleted BillingGroupStateType = "deleted"
)

func BillingGroupStateTypeChoices() []string {
	return []string{"active", "deleted"}
}

type CreditOut struct {
	Code           *string    `json:"code,omitempty"`
	ExpireTime     *time.Time `json:"expire_time,omitempty"`
	RemainingValue *string    `json:"remaining_value,omitempty"`
	StartTime      *time.Time `json:"start_time,omitempty"`
	Type           CreditType `json:"type,omitempty"`
	Value          *string    `json:"value,omitempty"`
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

type InvoiceGetOut struct {
	BillingGroupId    string    `json:"billing_group_id"`
	BillingGroupName  string    `json:"billing_group_name"`
	BillingGroupState string    `json:"billing_group_state"`
	Currency          string    `json:"currency"`
	DownloadCookie    *string   `json:"download_cookie,omitempty"`
	GeneratedAt       time.Time `json:"generated_at"`
	InvoiceNumber     string    `json:"invoice_number"`
	PeriodBegin       string    `json:"period_begin"`
	PeriodEnd         string    `json:"period_end"`
	State             string    `json:"state"`
	TotalIncVat       string    `json:"total_inc_vat"`
	TotalVatZero      string    `json:"total_vat_zero"`
}
type InvoiceOut struct {
	BillingGroupId    string                `json:"billing_group_id"`
	BillingGroupName  string                `json:"billing_group_name"`
	BillingGroupState BillingGroupStateType `json:"billing_group_state"`
	Currency          CurrencyType          `json:"currency"`
	DownloadCookie    string                `json:"download_cookie"`
	GeneratedAt       *time.Time            `json:"generated_at,omitempty"`
	InvoiceNumber     string                `json:"invoice_number"`
	PeriodBegin       string                `json:"period_begin"`
	PeriodEnd         string                `json:"period_end"`
	State             InvoiceStateType      `json:"state"`
	TotalIncVat       string                `json:"total_inc_vat"`
	TotalVatZero      string                `json:"total_vat_zero"`
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

type ProjectCreditsClaimIn struct {
	Code string `json:"code"`
}
type ProjectCreditsClaimOut struct {
	Code           *string    `json:"code,omitempty"`
	ExpireTime     *time.Time `json:"expire_time,omitempty"`
	RemainingValue *string    `json:"remaining_value,omitempty"`
	StartTime      *time.Time `json:"start_time,omitempty"`
	Type           CreditType `json:"type,omitempty"`
	Value          *string    `json:"value,omitempty"`
}
type invoiceGetOut struct {
	Invoice InvoiceGetOut `json:"invoice"`
}
type projectCreditsClaimOut struct {
	Credit ProjectCreditsClaimOut `json:"credit"`
}
type projectCreditsListOut struct {
	Credits []CreditOut `json:"credits"`
}
type projectInvoiceListOut struct {
	Invoices []InvoiceOut `json:"invoices"`
}
