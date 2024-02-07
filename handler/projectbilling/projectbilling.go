// Code generated by Aiven. DO NOT EDIT.

package projectbilling

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// ProjectCreditsClaim claim a credit code
	// POST /project/{project}/credits
	// https://api.aiven.io/doc/#tag/Project_Billing/operation/ProjectCreditsClaim
	ProjectCreditsClaim(ctx context.Context, project string, in *ProjectCreditsClaimIn) (*ProjectCreditsClaimOut, error)

	// ProjectCreditsList list project credits
	// GET /project/{project}/credits
	// https://api.aiven.io/doc/#tag/Project_Billing/operation/ProjectCreditsList
	ProjectCreditsList(ctx context.Context, project string) ([]CreditOut, error)

	// ProjectInvoiceList list project invoices
	// GET /project/{project}/invoice
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

func (h *ProjectBillingHandler) ProjectCreditsClaim(ctx context.Context, project string, in *ProjectCreditsClaimIn) (*ProjectCreditsClaimOut, error) {
	path := fmt.Sprintf("/project/%s/credits", project)
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
	path := fmt.Sprintf("/project/%s/credits", project)
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
	path := fmt.Sprintf("/project/%s/invoice", project)
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

type CreditOut struct {
	Code           string     `json:"code,omitempty"`
	ExpireTime     *time.Time `json:"expire_time,omitempty"`
	RemainingValue string     `json:"remaining_value,omitempty"`
	StartTime      *time.Time `json:"start_time,omitempty"`
	Type           string     `json:"type,omitempty"`
	Value          string     `json:"value,omitempty"`
}
type InvoiceOut struct {
	BillingGroupId    string     `json:"billing_group_id"`
	BillingGroupName  string     `json:"billing_group_name"`
	BillingGroupState string     `json:"billing_group_state"`
	Currency          string     `json:"currency"`
	DownloadCookie    string     `json:"download_cookie"`
	GeneratedAt       *time.Time `json:"generated_at,omitempty"`
	InvoiceNumber     string     `json:"invoice_number"`
	PeriodBegin       string     `json:"period_begin"`
	PeriodEnd         string     `json:"period_end"`
	State             string     `json:"state"`
	TotalIncVat       string     `json:"total_inc_vat"`
	TotalVatZero      string     `json:"total_vat_zero"`
}
type ProjectCreditsClaimIn struct {
	Code string `json:"code"`
}
type ProjectCreditsClaimOut struct {
	Code           string     `json:"code,omitempty"`
	ExpireTime     *time.Time `json:"expire_time,omitempty"`
	RemainingValue string     `json:"remaining_value,omitempty"`
	StartTime      *time.Time `json:"start_time,omitempty"`
	Type           string     `json:"type,omitempty"`
	Value          string     `json:"value,omitempty"`
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
