// Code generated by Aiven. DO NOT EDIT.

package staticip

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// ProjectStaticIPAssociate associate a static IP address with a service
	// POST /v1/project/{project}/static-ips/{static_ip_address_id}/association
	// https://api.aiven.io/doc/#tag/StaticIP/operation/ProjectStaticIPAssociate
	ProjectStaticIPAssociate(ctx context.Context, project string, staticIpAddressId string, in *ProjectStaticIpassociateIn) (*ProjectStaticIpassociateOut, error)

	// ProjectStaticIPAvailabilityList list static IP address cloud availability and prices for a project
	// GET /v1/project/{project}/static-ip-availability
	// https://api.aiven.io/doc/#tag/StaticIP/operation/ProjectStaticIPAvailabilityList
	ProjectStaticIPAvailabilityList(ctx context.Context, project string) ([]StaticIpAddressAvailabilityOut, error)

	// ProjectStaticIPDissociate dissociate a static IP address from a service
	// DELETE /v1/project/{project}/static-ips/{static_ip_address_id}/association
	// https://api.aiven.io/doc/#tag/StaticIP/operation/ProjectStaticIPDissociate
	ProjectStaticIPDissociate(ctx context.Context, project string, staticIpAddressId string) (*ProjectStaticIpdissociateOut, error)

	// ProjectStaticIPPatch update a static IP address configuration
	// PATCH /v1/project/{project}/static-ips/{static_ip_address_id}
	// https://api.aiven.io/doc/#tag/StaticIP/operation/ProjectStaticIPPatch
	ProjectStaticIPPatch(ctx context.Context, project string, staticIpAddressId string, in *ProjectStaticIppatchIn) (*ProjectStaticIppatchOut, error)

	// PublicStaticIPAvailabilityList list static IP clouds and prices
	// GET /v1/tenants/{tenant}/static-ip-availability
	// https://api.aiven.io/doc/#tag/Cloud_platforms/operation/PublicStaticIPAvailabilityList
	PublicStaticIPAvailabilityList(ctx context.Context, tenant string) ([]StaticIpAddressAvailabilityOut, error)

	// StaticIPCreate create static IP address
	// POST /v1/project/{project}/static-ips
	// https://api.aiven.io/doc/#tag/StaticIP/operation/StaticIPCreate
	StaticIPCreate(ctx context.Context, project string, in *StaticIpcreateIn) (*StaticIpcreateOut, error)

	// StaticIPList list static IP addresses
	// GET /v1/project/{project}/static-ips
	// https://api.aiven.io/doc/#tag/StaticIP/operation/StaticIPList
	StaticIPList(ctx context.Context, project string) ([]StaticIpOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) StaticIPHandler {
	return StaticIPHandler{doer}
}

type StaticIPHandler struct {
	doer doer
}

func (h *StaticIPHandler) ProjectStaticIPAssociate(ctx context.Context, project string, staticIpAddressId string, in *ProjectStaticIpassociateIn) (*ProjectStaticIpassociateOut, error) {
	path := fmt.Sprintf("/v1/project/%s/static-ips/%s/association", url.PathEscape(project), url.PathEscape(staticIpAddressId))
	b, err := h.doer.Do(ctx, "ProjectStaticIPAssociate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ProjectStaticIpassociateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *StaticIPHandler) ProjectStaticIPAvailabilityList(ctx context.Context, project string) ([]StaticIpAddressAvailabilityOut, error) {
	path := fmt.Sprintf("/v1/project/%s/static-ip-availability", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ProjectStaticIPAvailabilityList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(projectStaticIpavailabilityListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.StaticIpAddressAvailability, nil
}
func (h *StaticIPHandler) ProjectStaticIPDissociate(ctx context.Context, project string, staticIpAddressId string) (*ProjectStaticIpdissociateOut, error) {
	path := fmt.Sprintf("/v1/project/%s/static-ips/%s/association", url.PathEscape(project), url.PathEscape(staticIpAddressId))
	b, err := h.doer.Do(ctx, "ProjectStaticIPDissociate", "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ProjectStaticIpdissociateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *StaticIPHandler) ProjectStaticIPPatch(ctx context.Context, project string, staticIpAddressId string, in *ProjectStaticIppatchIn) (*ProjectStaticIppatchOut, error) {
	path := fmt.Sprintf("/v1/project/%s/static-ips/%s", url.PathEscape(project), url.PathEscape(staticIpAddressId))
	b, err := h.doer.Do(ctx, "ProjectStaticIPPatch", "PATCH", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ProjectStaticIppatchOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *StaticIPHandler) PublicStaticIPAvailabilityList(ctx context.Context, tenant string) ([]StaticIpAddressAvailabilityOut, error) {
	path := fmt.Sprintf("/v1/tenants/%s/static-ip-availability", url.PathEscape(tenant))
	b, err := h.doer.Do(ctx, "PublicStaticIPAvailabilityList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(publicStaticIpavailabilityListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.StaticIpAddressAvailability, nil
}
func (h *StaticIPHandler) StaticIPCreate(ctx context.Context, project string, in *StaticIpcreateIn) (*StaticIpcreateOut, error) {
	path := fmt.Sprintf("/v1/project/%s/static-ips", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "StaticIPCreate", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(StaticIpcreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *StaticIPHandler) StaticIPList(ctx context.Context, project string) ([]StaticIpOut, error) {
	path := fmt.Sprintf("/v1/project/%s/static-ips", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "StaticIPList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(staticIplistOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.StaticIps, nil
}

// ProjectStaticIpassociateIn ProjectStaticIPAssociateRequestBody
type ProjectStaticIpassociateIn struct {
	ServiceName string `json:"service_name"` // Service name
}

// ProjectStaticIpassociateOut ProjectStaticIPAssociateResponse
type ProjectStaticIpassociateOut struct {
	CloudName             string                            `json:"cloud_name"`             // Target cloud
	IpAddress             string                            `json:"ip_address"`             // IPv4 address
	ServiceName           string                            `json:"service_name"`           // Service name
	State                 ProjectStaticIpassociateStateType `json:"state"`                  // Static IP address state
	StaticIpAddressId     string                            `json:"static_ip_address_id"`   // Static IP address identifier
	TerminationProtection bool                              `json:"termination_protection"` // Static IP address is protected against deletion
}
type ProjectStaticIpassociateStateType string

const (
	ProjectStaticIpassociateStateTypeAssigned  ProjectStaticIpassociateStateType = "assigned"
	ProjectStaticIpassociateStateTypeAvailable ProjectStaticIpassociateStateType = "available"
	ProjectStaticIpassociateStateTypeCreated   ProjectStaticIpassociateStateType = "created"
	ProjectStaticIpassociateStateTypeCreating  ProjectStaticIpassociateStateType = "creating"
	ProjectStaticIpassociateStateTypeDeleted   ProjectStaticIpassociateStateType = "deleted"
	ProjectStaticIpassociateStateTypeDeleting  ProjectStaticIpassociateStateType = "deleting"
)

func ProjectStaticIpassociateStateTypeChoices() []string {
	return []string{"assigned", "available", "created", "creating", "deleted", "deleting"}
}

// ProjectStaticIpdissociateOut ProjectStaticIPDissociateResponse
type ProjectStaticIpdissociateOut struct {
	CloudName             string                             `json:"cloud_name"`             // Target cloud
	IpAddress             string                             `json:"ip_address"`             // IPv4 address
	ServiceName           string                             `json:"service_name"`           // Service name
	State                 ProjectStaticIpdissociateStateType `json:"state"`                  // Static IP address state
	StaticIpAddressId     string                             `json:"static_ip_address_id"`   // Static IP address identifier
	TerminationProtection bool                               `json:"termination_protection"` // Static IP address is protected against deletion
}
type ProjectStaticIpdissociateStateType string

const (
	ProjectStaticIpdissociateStateTypeAssigned  ProjectStaticIpdissociateStateType = "assigned"
	ProjectStaticIpdissociateStateTypeAvailable ProjectStaticIpdissociateStateType = "available"
	ProjectStaticIpdissociateStateTypeCreated   ProjectStaticIpdissociateStateType = "created"
	ProjectStaticIpdissociateStateTypeCreating  ProjectStaticIpdissociateStateType = "creating"
	ProjectStaticIpdissociateStateTypeDeleted   ProjectStaticIpdissociateStateType = "deleted"
	ProjectStaticIpdissociateStateTypeDeleting  ProjectStaticIpdissociateStateType = "deleting"
)

func ProjectStaticIpdissociateStateTypeChoices() []string {
	return []string{"assigned", "available", "created", "creating", "deleted", "deleting"}
}

// ProjectStaticIppatchIn ProjectStaticIPPatchRequestBody
type ProjectStaticIppatchIn struct {
	TerminationProtection *bool `json:"termination_protection,omitempty"` // Static IP address is protected against deletion
}

// ProjectStaticIppatchOut ProjectStaticIPPatchResponse
type ProjectStaticIppatchOut struct {
	CloudName             string                        `json:"cloud_name"`             // Target cloud
	IpAddress             string                        `json:"ip_address"`             // IPv4 address
	ServiceName           string                        `json:"service_name"`           // Service name
	State                 ProjectStaticIppatchStateType `json:"state"`                  // Static IP address state
	StaticIpAddressId     string                        `json:"static_ip_address_id"`   // Static IP address identifier
	TerminationProtection bool                          `json:"termination_protection"` // Static IP address is protected against deletion
}
type ProjectStaticIppatchStateType string

const (
	ProjectStaticIppatchStateTypeAssigned  ProjectStaticIppatchStateType = "assigned"
	ProjectStaticIppatchStateTypeAvailable ProjectStaticIppatchStateType = "available"
	ProjectStaticIppatchStateTypeCreated   ProjectStaticIppatchStateType = "created"
	ProjectStaticIppatchStateTypeCreating  ProjectStaticIppatchStateType = "creating"
	ProjectStaticIppatchStateTypeDeleted   ProjectStaticIppatchStateType = "deleted"
	ProjectStaticIppatchStateTypeDeleting  ProjectStaticIppatchStateType = "deleting"
)

func ProjectStaticIppatchStateTypeChoices() []string {
	return []string{"assigned", "available", "created", "creating", "deleted", "deleting"}
}

type StaticIpAddressAvailabilityOut struct {
	CloudName string `json:"cloud_name"` // Target cloud
	PriceUsd  string `json:"price_usd"`  // Hourly static IP address price in this cloud region
}
type StaticIpOut struct {
	CloudName             string            `json:"cloud_name"`             // Target cloud
	IpAddress             string            `json:"ip_address"`             // IPv4 address
	ServiceName           string            `json:"service_name"`           // Service name
	State                 StaticIpStateType `json:"state"`                  // Static IP address state
	StaticIpAddressId     string            `json:"static_ip_address_id"`   // Static IP address identifier
	TerminationProtection bool              `json:"termination_protection"` // Static IP address is protected against deletion
}
type StaticIpStateType string

const (
	StaticIpStateTypeAssigned  StaticIpStateType = "assigned"
	StaticIpStateTypeAvailable StaticIpStateType = "available"
	StaticIpStateTypeCreated   StaticIpStateType = "created"
	StaticIpStateTypeCreating  StaticIpStateType = "creating"
	StaticIpStateTypeDeleted   StaticIpStateType = "deleted"
	StaticIpStateTypeDeleting  StaticIpStateType = "deleting"
)

func StaticIpStateTypeChoices() []string {
	return []string{"assigned", "available", "created", "creating", "deleted", "deleting"}
}

// StaticIpcreateIn StaticIPCreateRequestBody
type StaticIpcreateIn struct {
	CloudName             string `json:"cloud_name"`                       // Target cloud
	TerminationProtection *bool  `json:"termination_protection,omitempty"` // Static IP address is protected against deletion
}

// StaticIpcreateOut StaticIPCreateResponse
type StaticIpcreateOut struct {
	CloudName             string                  `json:"cloud_name"`             // Target cloud
	IpAddress             string                  `json:"ip_address"`             // IPv4 address
	ServiceName           string                  `json:"service_name"`           // Service name
	State                 StaticIpcreateStateType `json:"state"`                  // Static IP address state
	StaticIpAddressId     string                  `json:"static_ip_address_id"`   // Static IP address identifier
	TerminationProtection bool                    `json:"termination_protection"` // Static IP address is protected against deletion
}
type StaticIpcreateStateType string

const (
	StaticIpcreateStateTypeAssigned  StaticIpcreateStateType = "assigned"
	StaticIpcreateStateTypeAvailable StaticIpcreateStateType = "available"
	StaticIpcreateStateTypeCreated   StaticIpcreateStateType = "created"
	StaticIpcreateStateTypeCreating  StaticIpcreateStateType = "creating"
	StaticIpcreateStateTypeDeleted   StaticIpcreateStateType = "deleted"
	StaticIpcreateStateTypeDeleting  StaticIpcreateStateType = "deleting"
)

func StaticIpcreateStateTypeChoices() []string {
	return []string{"assigned", "available", "created", "creating", "deleted", "deleting"}
}

// projectStaticIpavailabilityListOut ProjectStaticIPAvailabilityListResponse
type projectStaticIpavailabilityListOut struct {
	StaticIpAddressAvailability []StaticIpAddressAvailabilityOut `json:"static_ip_address_availability"` // Paginated array
}

// publicStaticIpavailabilityListOut PublicStaticIPAvailabilityListResponse
type publicStaticIpavailabilityListOut struct {
	StaticIpAddressAvailability []StaticIpAddressAvailabilityOut `json:"static_ip_address_availability"` // Paginated array
}

// staticIplistOut StaticIPListResponse
type staticIplistOut struct {
	StaticIps []StaticIpOut `json:"static_ips"` // Paginated array
}
