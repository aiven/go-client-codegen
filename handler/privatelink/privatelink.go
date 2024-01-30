// Code generated by Aiven. DO NOT EDIT.

package privatelink

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// PublicPrivatelinkAvailabilityList list privatelink clouds and prices
	// GET /tenants/{tenant}/privatelink-availability
	// https://api.aiven.io/doc/#tag/Cloud_platforms/operation/PublicPrivatelinkAvailabilityList
	PublicPrivatelinkAvailabilityList(ctx context.Context, tenant string) ([]PrivatelinkAvailabilityOut, error)

	// ServicePrivatelinkAWSConnectionList list VPC Endpoint connections for an AWS Privatelink Endpoint Service
	// GET /project/{project}/service/{service_name}/privatelink/aws/connections
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAWSConnectionList
	ServicePrivatelinkAWSConnectionList(ctx context.Context, project string, serviceName string) ([]ConnectionOut, error)

	// ServicePrivatelinkAWSCreate create an AWS Privatelink Endpoint Service
	// POST /project/{project}/service/{service_name}/privatelink/aws
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAWSCreate
	ServicePrivatelinkAWSCreate(ctx context.Context, project string, serviceName string, in *ServicePrivatelinkAwscreateIn) (*ServicePrivatelinkAwscreateOut, error)

	// ServicePrivatelinkAWSDelete delete an AWS Privatelink Endpoint Service
	// DELETE /project/{project}/service/{service_name}/privatelink/aws
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAWSDelete
	ServicePrivatelinkAWSDelete(ctx context.Context, project string, serviceName string) (*ServicePrivatelinkAwsdeleteOut, error)

	// ServicePrivatelinkAWSGet get AWS Privatelink Endpoint Service information
	// GET /project/{project}/service/{service_name}/privatelink/aws
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAWSGet
	ServicePrivatelinkAWSGet(ctx context.Context, project string, serviceName string) (*ServicePrivatelinkAwsgetOut, error)

	// ServicePrivatelinkAWSUpdate update an AWS Privatelink Endpoint Service
	// PUT /project/{project}/service/{service_name}/privatelink/aws
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAWSUpdate
	ServicePrivatelinkAWSUpdate(ctx context.Context, project string, serviceName string, in *ServicePrivatelinkAwsupdateIn) (*ServicePrivatelinkAwsupdateOut, error)

	// ServicePrivatelinkAzureConnectionApproval approve an Azure private endpoint connection pending user approval
	// POST /project/{project}/service/{service_name}/privatelink/azure/connections/{privatelink_connection_id}/approve
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAzureConnectionApproval
	ServicePrivatelinkAzureConnectionApproval(ctx context.Context, project string, serviceName string, privatelinkConnectionId string) (*ServicePrivatelinkAzureConnectionApprovalOut, error)

	// ServicePrivatelinkAzureConnectionList list private endpoint connections for an Azure Privatelink Service
	// GET /project/{project}/service/{service_name}/privatelink/azure/connections
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAzureConnectionList
	ServicePrivatelinkAzureConnectionList(ctx context.Context, project string, serviceName string) ([]ServicePrivatelinkAzureConnectionApprovalOut, error)

	// ServicePrivatelinkAzureConnectionUpdate update a private endpoint connection for an Azure Privatelink Service
	// PUT /project/{project}/service/{service_name}/privatelink/azure/connections/{privatelink_connection_id}
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAzureConnectionUpdate
	ServicePrivatelinkAzureConnectionUpdate(ctx context.Context, project string, serviceName string, privatelinkConnectionId string, in *ServicePrivatelinkAzureConnectionUpdateIn) (*ServicePrivatelinkAzureConnectionUpdateOut, error)

	// ServicePrivatelinkAzureCreate create an Azure Privatelink Service
	// POST /project/{project}/service/{service_name}/privatelink/azure
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAzureCreate
	ServicePrivatelinkAzureCreate(ctx context.Context, project string, serviceName string, in *ServicePrivatelinkAzureCreateIn) (*ServicePrivatelinkAzureCreateOut, error)

	// ServicePrivatelinkAzureDelete delete an Azure Privatelink Service
	// DELETE /project/{project}/service/{service_name}/privatelink/azure
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAzureDelete
	ServicePrivatelinkAzureDelete(ctx context.Context, project string, serviceName string) (*ServicePrivatelinkAzureDeleteOut, error)

	// ServicePrivatelinkAzureGet get Azure Privatelink Service information
	// GET /project/{project}/service/{service_name}/privatelink/azure
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAzureGet
	ServicePrivatelinkAzureGet(ctx context.Context, project string, serviceName string) (*ServicePrivatelinkAzureGetOut, error)

	// ServicePrivatelinkAzureUpdate update an Azure Privatelink Service
	// PUT /project/{project}/service/{service_name}/privatelink/azure
	// https://api.aiven.io/doc/#tag/Service/operation/ServicePrivatelinkAzureUpdate
	ServicePrivatelinkAzureUpdate(ctx context.Context, project string, serviceName string, in *ServicePrivatelinkAzureUpdateIn) (*ServicePrivatelinkAzureUpdateOut, error)
}

func NewHandler(doer doer) PrivatelinkHandler {
	return PrivatelinkHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type PrivatelinkHandler struct {
	doer doer
}

func (h *PrivatelinkHandler) PublicPrivatelinkAvailabilityList(ctx context.Context, tenant string) ([]PrivatelinkAvailabilityOut, error) {
	path := fmt.Sprintf("/tenants/%s/privatelink-availability", tenant)
	b, err := h.doer.Do(ctx, "PublicPrivatelinkAvailabilityList", "GET", path, nil)
	out := new(PublicPrivatelinkAvailabilityListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.PrivatelinkAvailability, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAWSConnectionList(ctx context.Context, project string, serviceName string) ([]ConnectionOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/aws/connections", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAWSConnectionList", "GET", path, nil)
	out := new(ServicePrivatelinkAwsconnectionListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Connections, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAWSCreate(ctx context.Context, project string, serviceName string, in *ServicePrivatelinkAwscreateIn) (*ServicePrivatelinkAwscreateOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/aws", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAWSCreate", "POST", path, in)
	out := new(ServicePrivatelinkAwscreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAWSDelete(ctx context.Context, project string, serviceName string) (*ServicePrivatelinkAwsdeleteOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/aws", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAWSDelete", "DELETE", path, nil)
	out := new(ServicePrivatelinkAwsdeleteOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAWSGet(ctx context.Context, project string, serviceName string) (*ServicePrivatelinkAwsgetOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/aws", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAWSGet", "GET", path, nil)
	out := new(ServicePrivatelinkAwsgetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAWSUpdate(ctx context.Context, project string, serviceName string, in *ServicePrivatelinkAwsupdateIn) (*ServicePrivatelinkAwsupdateOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/aws", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAWSUpdate", "PUT", path, in)
	out := new(ServicePrivatelinkAwsupdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAzureConnectionApproval(ctx context.Context, project string, serviceName string, privatelinkConnectionId string) (*ServicePrivatelinkAzureConnectionApprovalOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/azure/connections/%s/approve", project, serviceName, privatelinkConnectionId)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAzureConnectionApproval", "POST", path, nil)
	out := new(ServicePrivatelinkAzureConnectionApprovalOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAzureConnectionList(ctx context.Context, project string, serviceName string) ([]ServicePrivatelinkAzureConnectionApprovalOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/azure/connections", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAzureConnectionList", "GET", path, nil)
	out := new(ServicePrivatelinkAzureConnectionListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Connections, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAzureConnectionUpdate(ctx context.Context, project string, serviceName string, privatelinkConnectionId string, in *ServicePrivatelinkAzureConnectionUpdateIn) (*ServicePrivatelinkAzureConnectionUpdateOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/azure/connections/%s", project, serviceName, privatelinkConnectionId)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAzureConnectionUpdate", "PUT", path, in)
	out := new(ServicePrivatelinkAzureConnectionUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAzureCreate(ctx context.Context, project string, serviceName string, in *ServicePrivatelinkAzureCreateIn) (*ServicePrivatelinkAzureCreateOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/azure", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAzureCreate", "POST", path, in)
	out := new(ServicePrivatelinkAzureCreateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAzureDelete(ctx context.Context, project string, serviceName string) (*ServicePrivatelinkAzureDeleteOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/azure", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAzureDelete", "DELETE", path, nil)
	out := new(ServicePrivatelinkAzureDeleteOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAzureGet(ctx context.Context, project string, serviceName string) (*ServicePrivatelinkAzureGetOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/azure", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAzureGet", "GET", path, nil)
	out := new(ServicePrivatelinkAzureGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *PrivatelinkHandler) ServicePrivatelinkAzureUpdate(ctx context.Context, project string, serviceName string, in *ServicePrivatelinkAzureUpdateIn) (*ServicePrivatelinkAzureUpdateOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/privatelink/azure", project, serviceName)
	b, err := h.doer.Do(ctx, "ServicePrivatelinkAzureUpdate", "PUT", path, in)
	out := new(ServicePrivatelinkAzureUpdateOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ConnectionOut struct {
	DnsName                 string    `json:"dns_name"`
	PrivatelinkConnectionId string    `json:"privatelink_connection_id,omitempty"`
	State                   StateType `json:"state"`
	VpcEndpointId           string    `json:"vpc_endpoint_id"`
}
type PrivatelinkAvailabilityOut struct {
	CloudName string `json:"cloud_name"`
	PriceUsd  string `json:"price_usd"`
}
type PublicPrivatelinkAvailabilityListOut struct {
	PrivatelinkAvailability []PrivatelinkAvailabilityOut `json:"privatelink_availability"`
}
type ServicePrivatelinkAwsconnectionListOut struct {
	Connections []ConnectionOut `json:"connections"`
}
type ServicePrivatelinkAwscreateIn struct {
	Principals []string `json:"principals"`
}
type ServicePrivatelinkAwscreateOut struct {
	AwsServiceId   string                               `json:"aws_service_id,omitempty"`
	AwsServiceName string                               `json:"aws_service_name,omitempty"`
	Principals     []string                             `json:"principals"`
	State          ServicePrivatelinkAwscreateStateType `json:"state"`
}
type ServicePrivatelinkAwscreateStateType string

const (
	ServicePrivatelinkAwscreateStateTypeCreating ServicePrivatelinkAwscreateStateType = "creating"
	ServicePrivatelinkAwscreateStateTypeActive   ServicePrivatelinkAwscreateStateType = "active"
	ServicePrivatelinkAwscreateStateTypeDeleting ServicePrivatelinkAwscreateStateType = "deleting"
)

type ServicePrivatelinkAwsdeleteOut struct {
	AwsServiceId   string                               `json:"aws_service_id,omitempty"`
	AwsServiceName string                               `json:"aws_service_name,omitempty"`
	Principals     []string                             `json:"principals"`
	State          ServicePrivatelinkAwscreateStateType `json:"state"`
}
type ServicePrivatelinkAwsgetOut struct {
	AwsServiceId   string                               `json:"aws_service_id,omitempty"`
	AwsServiceName string                               `json:"aws_service_name,omitempty"`
	Principals     []string                             `json:"principals"`
	State          ServicePrivatelinkAwscreateStateType `json:"state"`
}
type ServicePrivatelinkAwsupdateIn struct {
	Principals []string `json:"principals"`
}
type ServicePrivatelinkAwsupdateOut struct {
	AwsServiceId   string                               `json:"aws_service_id,omitempty"`
	AwsServiceName string                               `json:"aws_service_name,omitempty"`
	Principals     []string                             `json:"principals"`
	State          ServicePrivatelinkAwscreateStateType `json:"state"`
}
type ServicePrivatelinkAzureConnectionApprovalOut struct {
	PrivateEndpointId       string    `json:"private_endpoint_id"`
	PrivatelinkConnectionId string    `json:"privatelink_connection_id,omitempty"`
	State                   StateType `json:"state"`
	UserIpAddress           string    `json:"user_ip_address"`
}
type ServicePrivatelinkAzureConnectionListOut struct {
	Connections []ServicePrivatelinkAzureConnectionApprovalOut `json:"connections"`
}
type ServicePrivatelinkAzureConnectionUpdateIn struct {
	UserIpAddress string `json:"user_ip_address"`
}
type ServicePrivatelinkAzureConnectionUpdateOut struct {
	PrivateEndpointId       string    `json:"private_endpoint_id"`
	PrivatelinkConnectionId string    `json:"privatelink_connection_id,omitempty"`
	State                   StateType `json:"state"`
	UserIpAddress           string    `json:"user_ip_address"`
}
type ServicePrivatelinkAzureCreateIn struct {
	UserSubscriptionIds []string `json:"user_subscription_ids"`
}
type ServicePrivatelinkAzureCreateOut struct {
	AzureServiceAlias   string                               `json:"azure_service_alias,omitempty"`
	AzureServiceId      string                               `json:"azure_service_id,omitempty"`
	State               ServicePrivatelinkAwscreateStateType `json:"state"`
	UserSubscriptionIds []string                             `json:"user_subscription_ids"`
}
type ServicePrivatelinkAzureDeleteOut struct {
	AzureServiceAlias   string                               `json:"azure_service_alias,omitempty"`
	AzureServiceId      string                               `json:"azure_service_id,omitempty"`
	State               ServicePrivatelinkAwscreateStateType `json:"state"`
	UserSubscriptionIds []string                             `json:"user_subscription_ids"`
}
type ServicePrivatelinkAzureGetOut struct {
	AzureServiceAlias   string                               `json:"azure_service_alias,omitempty"`
	AzureServiceId      string                               `json:"azure_service_id,omitempty"`
	State               ServicePrivatelinkAwscreateStateType `json:"state"`
	UserSubscriptionIds []string                             `json:"user_subscription_ids"`
}
type ServicePrivatelinkAzureUpdateIn struct {
	UserSubscriptionIds []string `json:"user_subscription_ids"`
}
type ServicePrivatelinkAzureUpdateOut struct {
	AzureServiceAlias   string                               `json:"azure_service_alias,omitempty"`
	AzureServiceId      string                               `json:"azure_service_id,omitempty"`
	State               ServicePrivatelinkAwscreateStateType `json:"state"`
	UserSubscriptionIds []string                             `json:"user_subscription_ids"`
}
type StateType string

const (
	StateTypePendingUserApproval StateType = "pending-user-approval"
	StateTypeUserApproved        StateType = "user-approved"
	StateTypeConnected           StateType = "connected"
	StateTypeActive              StateType = "active"
)
