// Code generated by Aiven. DO NOT EDIT.

package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// ListClouds list cloud platforms
	// GET /v1/clouds
	// https://api.aiven.io/doc/#tag/Cloud_platforms/operation/ListClouds
	ListClouds(ctx context.Context) ([]CloudOut, error)

	// ListProjectClouds list cloud platforms for a project
	// GET /v1/project/{project}/clouds
	// https://api.aiven.io/doc/#tag/Cloud_platforms/operation/ListProjectClouds
	ListProjectClouds(ctx context.Context, project string) ([]CloudOut, error)
}

func NewHandler(doer doer) CloudHandler {
	return CloudHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type CloudHandler struct {
	doer doer
}

func (h *CloudHandler) ListClouds(ctx context.Context) ([]CloudOut, error) {
	path := fmt.Sprintf("/v1/clouds")
	b, err := h.doer.Do(ctx, "ListClouds", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(listCloudsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Clouds, nil
}
func (h *CloudHandler) ListProjectClouds(ctx context.Context, project string) ([]CloudOut, error) {
	path := fmt.Sprintf("/v1/project/%s/clouds", url.PathEscape(project))
	b, err := h.doer.Do(ctx, "ListProjectClouds", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(listProjectCloudsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Clouds, nil
}

type CloudOut struct {
	CloudDescription    string   `json:"cloud_description,omitempty"`
	CloudName           string   `json:"cloud_name"`
	GeoLatitude         *float64 `json:"geo_latitude,omitempty"`
	GeoLongitude        *float64 `json:"geo_longitude,omitempty"`
	GeoRegion           string   `json:"geo_region"`
	Provider            string   `json:"provider,omitempty"`
	ProviderDescription string   `json:"provider_description,omitempty"`
}
type listCloudsOut struct {
	Clouds []CloudOut `json:"clouds"`
}
type listProjectCloudsOut struct {
	Clouds []CloudOut `json:"clouds"`
}
