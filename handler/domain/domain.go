// Code generated by Aiven. DO NOT EDIT.

package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Handler interface {
	// OrganizationDomainAdd create a domain
	// POST /organization/{organization_id}/domains
	// https://api.aiven.io/doc/#tag/Domains/operation/OrganizationDomainAdd
	OrganizationDomainAdd(ctx context.Context, organizationId string, in *OrganizationDomainAddIn) (*OrganizationDomainAddOut, error)

	// OrganizationDomainUpdate update a domain
	// PATCH /organization/{organization_id}/domains/{domain_id}
	// https://api.aiven.io/doc/#tag/Domains/operation/OrganizationDomainUpdate
	OrganizationDomainUpdate(ctx context.Context, organizationId string, domainId string, in *OrganizationDomainUpdateIn) (*OrganizationDomainUpdateOut, error)

	// OrganizationDomainVerify verify a domain
	// POST /organization/{organization_id}/domains/{domain_id}/verify
	// https://api.aiven.io/doc/#tag/Domains/operation/OrganizationDomainVerify
	OrganizationDomainVerify(ctx context.Context, organizationId string, domainId string) (*OrganizationDomainVerifyOut, error)

	// OrganizationDomainsList list domains
	// GET /organization/{organization_id}/domains
	// https://api.aiven.io/doc/#tag/Domains/operation/OrganizationDomainsList
	OrganizationDomainsList(ctx context.Context, organizationId string) ([]DomainOut, error)

	// OrganizationDomainsRemove delete a domain
	// DELETE /organization/{organization_id}/domains/{domain_id}
	// https://api.aiven.io/doc/#tag/Domains/operation/OrganizationDomainsRemove
	OrganizationDomainsRemove(ctx context.Context, organizationId string, domainId string) error
}

func NewHandler(doer doer) DomainHandler {
	return DomainHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type DomainHandler struct {
	doer doer
}

func (h *DomainHandler) OrganizationDomainAdd(ctx context.Context, organizationId string, in *OrganizationDomainAddIn) (*OrganizationDomainAddOut, error) {
	path := fmt.Sprintf("/organization/%s/domains", organizationId)
	b, err := h.doer.Do(ctx, "OrganizationDomainAdd", "POST", path, in)
	out := new(OrganizationDomainAddOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *DomainHandler) OrganizationDomainUpdate(ctx context.Context, organizationId string, domainId string, in *OrganizationDomainUpdateIn) (*OrganizationDomainUpdateOut, error) {
	path := fmt.Sprintf("/organization/%s/domains/%s", organizationId, domainId)
	b, err := h.doer.Do(ctx, "OrganizationDomainUpdate", "PATCH", path, in)
	out := new(OrganizationDomainUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *DomainHandler) OrganizationDomainVerify(ctx context.Context, organizationId string, domainId string) (*OrganizationDomainVerifyOut, error) {
	path := fmt.Sprintf("/organization/%s/domains/%s/verify", organizationId, domainId)
	b, err := h.doer.Do(ctx, "OrganizationDomainVerify", "POST", path, nil)
	out := new(OrganizationDomainVerifyOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *DomainHandler) OrganizationDomainsList(ctx context.Context, organizationId string) ([]DomainOut, error) {
	path := fmt.Sprintf("/organization/%s/domains", organizationId)
	b, err := h.doer.Do(ctx, "OrganizationDomainsList", "GET", path, nil)
	out := new(organizationDomainsListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Domains, nil
}
func (h *DomainHandler) OrganizationDomainsRemove(ctx context.Context, organizationId string, domainId string) error {
	path := fmt.Sprintf("/organization/%s/domains/%s", organizationId, domainId)
	_, err := h.doer.Do(ctx, "OrganizationDomainsRemove", "DELETE", path, nil)
	return err
}

type DomainOut struct {
	ChallengeToken   string    `json:"challenge_token"`
	CreateTime       time.Time `json:"create_time"`
	DomainId         string    `json:"domain_id"`
	DomainName       string    `json:"domain_name"`
	OrganizationId   string    `json:"organization_id"`
	State            string    `json:"state"`
	VerificationType string    `json:"verification_type"`
}
type OrganizationDomainAddIn struct {
	DomainName       string           `json:"domain_name"`
	VerificationType VerificationType `json:"verification_type"`
}
type OrganizationDomainAddOut struct {
	ChallengeToken   string    `json:"challenge_token"`
	CreateTime       time.Time `json:"create_time"`
	DomainId         string    `json:"domain_id"`
	DomainName       string    `json:"domain_name"`
	OrganizationId   string    `json:"organization_id"`
	State            string    `json:"state"`
	VerificationType string    `json:"verification_type"`
}
type OrganizationDomainUpdateIn struct {
	VerificationType VerificationType `json:"verification_type,omitempty"`
}
type OrganizationDomainUpdateOut struct {
	ChallengeToken   string    `json:"challenge_token"`
	CreateTime       time.Time `json:"create_time"`
	DomainId         string    `json:"domain_id"`
	DomainName       string    `json:"domain_name"`
	OrganizationId   string    `json:"organization_id"`
	State            string    `json:"state"`
	VerificationType string    `json:"verification_type"`
}
type OrganizationDomainVerifyOut struct {
	ChallengeToken   string    `json:"challenge_token"`
	CreateTime       time.Time `json:"create_time"`
	DomainId         string    `json:"domain_id"`
	DomainName       string    `json:"domain_name"`
	OrganizationId   string    `json:"organization_id"`
	State            string    `json:"state"`
	VerificationType string    `json:"verification_type"`
}
type VerificationType string

const (
	VerificationTypeDns  VerificationType = "dns"
	VerificationTypeHttp VerificationType = "http"
)

func VerificationTypeChoices() []string {
	return []string{"dns", "http"}
}

type organizationDomainsListOut struct {
	Domains []DomainOut `json:"domains"`
}
