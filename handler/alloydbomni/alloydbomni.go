// Code generated by Aiven. DO NOT EDIT.

package alloydbomni

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// AlloyDbOmniGoogleCloudPrivateKeyIdentify get Google service account key
	// GET /v1/project/{project}/service/{service_name}/alloydbomni/google_cloud_private_key
	// https://api.aiven.io/doc/#tag/Service:_AlloyDB_Omni/operation/AlloyDbOmniGoogleCloudPrivateKeyIdentify
	AlloyDbOmniGoogleCloudPrivateKeyIdentify(ctx context.Context, project string, serviceName string) (*AlloyDbOmniGoogleCloudPrivateKeyIdentifyOut, error)

	// AlloyDbOmniGoogleCloudPrivateKeyRemove delete Google service account key
	// DELETE /v1/project/{project}/service/{service_name}/alloydbomni/google_cloud_private_key
	// https://api.aiven.io/doc/#tag/Service:_AlloyDB_Omni/operation/AlloyDbOmniGoogleCloudPrivateKeyRemove
	AlloyDbOmniGoogleCloudPrivateKeyRemove(ctx context.Context, project string, serviceName string) (*AlloyDbOmniGoogleCloudPrivateKeyRemoveOut, error)

	// AlloyDbOmniGoogleCloudPrivateKeySet add Google service account key
	// POST /v1/project/{project}/service/{service_name}/alloydbomni/google_cloud_private_key
	// https://api.aiven.io/doc/#tag/Service:_AlloyDB_Omni/operation/AlloyDbOmniGoogleCloudPrivateKeySet
	AlloyDbOmniGoogleCloudPrivateKeySet(ctx context.Context, project string, serviceName string, in *AlloyDbOmniGoogleCloudPrivateKeySetIn) (*AlloyDbOmniGoogleCloudPrivateKeySetOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) AlloyDBOmniHandler {
	return AlloyDBOmniHandler{doer}
}

type AlloyDBOmniHandler struct {
	doer doer
}

func (h *AlloyDBOmniHandler) AlloyDbOmniGoogleCloudPrivateKeyIdentify(ctx context.Context, project string, serviceName string) (*AlloyDbOmniGoogleCloudPrivateKeyIdentifyOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/alloydbomni/google_cloud_private_key", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "AlloyDbOmniGoogleCloudPrivateKeyIdentify", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(AlloyDbOmniGoogleCloudPrivateKeyIdentifyOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *AlloyDBOmniHandler) AlloyDbOmniGoogleCloudPrivateKeyRemove(ctx context.Context, project string, serviceName string) (*AlloyDbOmniGoogleCloudPrivateKeyRemoveOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/alloydbomni/google_cloud_private_key", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "AlloyDbOmniGoogleCloudPrivateKeyRemove", "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(AlloyDbOmniGoogleCloudPrivateKeyRemoveOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *AlloyDBOmniHandler) AlloyDbOmniGoogleCloudPrivateKeySet(ctx context.Context, project string, serviceName string, in *AlloyDbOmniGoogleCloudPrivateKeySetIn) (*AlloyDbOmniGoogleCloudPrivateKeySetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/alloydbomni/google_cloud_private_key", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "AlloyDbOmniGoogleCloudPrivateKeySet", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(AlloyDbOmniGoogleCloudPrivateKeySetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlloyDbOmniGoogleCloudPrivateKeyIdentifyOut AlloyDbOmniGoogleCloudPrivateKeyIdentifyResponse
type AlloyDbOmniGoogleCloudPrivateKeyIdentifyOut struct {
	ClientEmail  string `json:"client_email"`   // Email address of Google service account key
	PrivateKeyId string `json:"private_key_id"` // Google service account key ID
}

// AlloyDbOmniGoogleCloudPrivateKeyRemoveOut AlloyDbOmniGoogleCloudPrivateKeyRemoveResponse
type AlloyDbOmniGoogleCloudPrivateKeyRemoveOut struct {
	ClientEmail  string `json:"client_email"`   // Email address of Google service account key
	PrivateKeyId string `json:"private_key_id"` // Google service account key ID
}

// AlloyDbOmniGoogleCloudPrivateKeySetIn AlloyDbOmniGoogleCloudPrivateKeySetRequestBody
type AlloyDbOmniGoogleCloudPrivateKeySetIn struct {
	PrivateKey string `json:"private_key"` // Google Service Account Credentials
}

// AlloyDbOmniGoogleCloudPrivateKeySetOut AlloyDbOmniGoogleCloudPrivateKeySetResponse
type AlloyDbOmniGoogleCloudPrivateKeySetOut struct {
	ClientEmail  string `json:"client_email"`   // Email address of Google service account key
	PrivateKeyId string `json:"private_key_id"` // Google service account key ID
}