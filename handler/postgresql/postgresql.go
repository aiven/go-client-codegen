// Code generated by Aiven. DO NOT EDIT.

package postgresql

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// PGServiceAvailableExtensions list PostgreSQL extensions that can be loaded with CREATE EXTENSION in this service
	// GET /v1/project/{project}/service/{service_name}/pg/available-extensions
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/PGServiceAvailableExtensions
	PGServiceAvailableExtensions(ctx context.Context, project string, serviceName string) ([]ExtensionOut, error)

	// PGServiceQueryStatistics fetch PostgreSQL service query statistics
	// POST /v1/project/{project}/service/{service_name}/pg/query/stats
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/PGServiceQueryStatistics
	PGServiceQueryStatistics(ctx context.Context, project string, serviceName string, in *PgserviceQueryStatisticsIn) ([]QueryOut, error)

	// PgAvailableExtensions list PostgreSQL extensions available for this tenant grouped by PG version
	// GET /v1/tenants/{tenant}/pg-available-extensions
	// https://api.aiven.io/doc/#tag/Service/operation/PgAvailableExtensions
	PgAvailableExtensions(ctx context.Context, tenant string) ([]PgOut, error)

	// ServicePGBouncerCreate create a new connection pool for service
	// POST /v1/project/{project}/service/{service_name}/connection_pool
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/ServicePGBouncerCreate
	ServicePGBouncerCreate(ctx context.Context, project string, serviceName string, in *ServicePgbouncerCreateIn) error

	// ServicePGBouncerDelete delete a connection pool
	// DELETE /v1/project/{project}/service/{service_name}/connection_pool/{pool_name}
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/ServicePGBouncerDelete
	ServicePGBouncerDelete(ctx context.Context, project string, serviceName string, poolName string) error

	// ServicePGBouncerUpdate update a connection pool
	// PUT /v1/project/{project}/service/{service_name}/connection_pool/{pool_name}
	// https://api.aiven.io/doc/#tag/Service:_PostgreSQL/operation/ServicePGBouncerUpdate
	ServicePGBouncerUpdate(ctx context.Context, project string, serviceName string, poolName string, in *ServicePgbouncerUpdateIn) error
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) PostgreSQLHandler {
	return PostgreSQLHandler{doer}
}

type PostgreSQLHandler struct {
	doer doer
}

func (h *PostgreSQLHandler) PGServiceAvailableExtensions(ctx context.Context, project string, serviceName string) ([]ExtensionOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/pg/available-extensions", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "PGServiceAvailableExtensions", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(pgserviceAvailableExtensionsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Extensions, nil
}
func (h *PostgreSQLHandler) PGServiceQueryStatistics(ctx context.Context, project string, serviceName string, in *PgserviceQueryStatisticsIn) ([]QueryOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/pg/query/stats", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "PGServiceQueryStatistics", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(pgserviceQueryStatisticsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Queries, nil
}
func (h *PostgreSQLHandler) PgAvailableExtensions(ctx context.Context, tenant string) ([]PgOut, error) {
	path := fmt.Sprintf("/v1/tenants/%s/pg-available-extensions", url.PathEscape(tenant))
	b, err := h.doer.Do(ctx, "PgAvailableExtensions", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(pgAvailableExtensionsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Pg, nil
}
func (h *PostgreSQLHandler) ServicePGBouncerCreate(ctx context.Context, project string, serviceName string, in *ServicePgbouncerCreateIn) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connection_pool", url.PathEscape(project), url.PathEscape(serviceName))
	_, err := h.doer.Do(ctx, "ServicePGBouncerCreate", "POST", path, in)
	return err
}
func (h *PostgreSQLHandler) ServicePGBouncerDelete(ctx context.Context, project string, serviceName string, poolName string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connection_pool/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(poolName))
	_, err := h.doer.Do(ctx, "ServicePGBouncerDelete", "DELETE", path, nil)
	return err
}
func (h *PostgreSQLHandler) ServicePGBouncerUpdate(ctx context.Context, project string, serviceName string, poolName string, in *ServicePgbouncerUpdateIn) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/connection_pool/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(poolName))
	_, err := h.doer.Do(ctx, "ServicePGBouncerUpdate", "PUT", path, in)
	return err
}

type ExtensionOut struct {
	DefaultVersion *string  `json:"default_version,omitempty"` // Default version
	Name           string   `json:"name"`                      // Extension name
	Versions       []string `json:"versions,omitempty"`        // Extension versions available
}
type PgOut struct {
	Extensions []ExtensionOut `json:"extensions"` // Extensions available for loading with CREATE EXTENSION in this service
	Version    string         `json:"version"`    // PostgreSQL version
}

// PgserviceQueryStatisticsIn PGServiceQueryStatisticsRequestBody
type PgserviceQueryStatisticsIn struct {
	Limit   *int    `json:"limit,omitempty"`    // Limit for number of results
	Offset  *int    `json:"offset,omitempty"`   // Offset for retrieved results based on sort order
	OrderBy *string `json:"order_by,omitempty"` // Sort order can be either asc or desc and multiple comma separated columns with their own order can be specified: :asc,:desc. Accepted sort columns are: blk_read_time, blk_write_time, calls, database_name, local_blks_dirtied, local_blks_hit, local_blks_read, local_blks_written, max_plan_time, max_time, mean_plan_time, mean_time, min_plan_time, min_time, query, queryid, rows, shared_blks_dirtied, shared_blks_hit, shared_blks_read, shared_blks_written, stddev_plan_time, stddev_time, temp_blks_read, temp_blks_written, total_plan_time, total_time, user_name, wal_bytes, wal_fpi, wal_records
}
type PoolModeType string

const (
	PoolModeTypeSession     PoolModeType = "session"
	PoolModeTypeStatement   PoolModeType = "statement"
	PoolModeTypeTransaction PoolModeType = "transaction"
)

func PoolModeTypeChoices() []string {
	return []string{"session", "statement", "transaction"}
}

type QueryOut struct {
	BlkReadTime       *float64 `json:"blk_read_time,omitempty"`       // Query statistic
	BlkWriteTime      *float64 `json:"blk_write_time,omitempty"`      // Query statistic
	Calls             *float64 `json:"calls,omitempty"`               // Query statistic
	DatabaseName      *string  `json:"database_name,omitempty"`       // Query statistic
	LocalBlksDirtied  *float64 `json:"local_blks_dirtied,omitempty"`  // Query statistic
	LocalBlksHit      *float64 `json:"local_blks_hit,omitempty"`      // Query statistic
	LocalBlksRead     *float64 `json:"local_blks_read,omitempty"`     // Query statistic
	LocalBlksWritten  *float64 `json:"local_blks_written,omitempty"`  // Query statistic
	MaxPlanTime       *float64 `json:"max_plan_time,omitempty"`       // Query statistic
	MaxTime           *float64 `json:"max_time,omitempty"`            // Query statistic
	MeanPlanTime      *float64 `json:"mean_plan_time,omitempty"`      // Query statistic
	MeanTime          *float64 `json:"mean_time,omitempty"`           // Query statistic
	MinPlanTime       *float64 `json:"min_plan_time,omitempty"`       // Query statistic
	MinTime           *float64 `json:"min_time,omitempty"`            // Query statistic
	Query             *string  `json:"query,omitempty"`               // Query statistic
	Queryid           *float64 `json:"queryid,omitempty"`             // Query statistic
	Rows              *float64 `json:"rows,omitempty"`                // Query statistic
	SharedBlksDirtied *float64 `json:"shared_blks_dirtied,omitempty"` // Query statistic
	SharedBlksHit     *float64 `json:"shared_blks_hit,omitempty"`     // Query statistic
	SharedBlksRead    *float64 `json:"shared_blks_read,omitempty"`    // Query statistic
	SharedBlksWritten *float64 `json:"shared_blks_written,omitempty"` // Query statistic
	StddevPlanTime    *float64 `json:"stddev_plan_time,omitempty"`    // Query statistic
	StddevTime        *float64 `json:"stddev_time,omitempty"`         // Query statistic
	TempBlksRead      *float64 `json:"temp_blks_read,omitempty"`      // Query statistic
	TempBlksWritten   *float64 `json:"temp_blks_written,omitempty"`   // Query statistic
	TotalPlanTime     *float64 `json:"total_plan_time,omitempty"`     // Query statistic
	TotalTime         *float64 `json:"total_time,omitempty"`          // Query statistic
	UserName          *string  `json:"user_name,omitempty"`           // Query statistic
	WalBytes          *string  `json:"wal_bytes,omitempty"`           // Query statistic
	WalFpi            *float64 `json:"wal_fpi,omitempty"`             // Query statistic
	WalRecords        *float64 `json:"wal_records,omitempty"`         // Query statistic
}

// ServicePgbouncerCreateIn ServicePGBouncerCreateRequestBody
type ServicePgbouncerCreateIn struct {
	Database string       `json:"database"`            // Service database name
	PoolMode PoolModeType `json:"pool_mode,omitempty"` // PGBouncer pool mode
	PoolName string       `json:"pool_name"`           // Connection pool name
	PoolSize *int         `json:"pool_size,omitempty"` // Size of PGBouncer's PostgreSQL side connection pool
	Username *string      `json:"username,omitempty"`  // Service username
}

// ServicePgbouncerUpdateIn ServicePGBouncerUpdateRequestBody
type ServicePgbouncerUpdateIn struct {
	Database *string      `json:"database,omitempty"`  // Service database name
	PoolMode PoolModeType `json:"pool_mode,omitempty"` // PGBouncer pool mode
	PoolSize *int         `json:"pool_size,omitempty"` // Size of PGBouncer's PostgreSQL side connection pool
	Username *string      `json:"username,omitempty"`  // Service username
}

// pgAvailableExtensionsOut PgAvailableExtensionsResponse
type pgAvailableExtensionsOut struct {
	Pg []PgOut `json:"pg,omitempty"` // Supported PostgreSQL versions
}

// pgserviceAvailableExtensionsOut PGServiceAvailableExtensionsResponse
type pgserviceAvailableExtensionsOut struct {
	Extensions []ExtensionOut `json:"extensions"` // Extensions available for loading with CREATE EXTENSION in this service
}

// pgserviceQueryStatisticsOut PGServiceQueryStatisticsResponse
type pgserviceQueryStatisticsOut struct {
	Queries []QueryOut `json:"queries"` // List of query statistics
}
