// Code generated by Aiven. DO NOT EDIT.

package clickhouse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// ServiceClickHouseCurrentQueries list active queries
	// GET /v1/project/{project}/service/{service_name}/clickhouse/query
	// https://api.aiven.io/doc/#tag/Service:_ClickHouse/operation/ServiceClickHouseCurrentQueries
	ServiceClickHouseCurrentQueries(ctx context.Context, project string, serviceName string) ([]QueryOut, error)

	// ServiceClickHouseDatabaseCreate create a database
	// POST /v1/project/{project}/service/{service_name}/clickhouse/db
	// https://api.aiven.io/doc/#tag/Service:_ClickHouse/operation/ServiceClickHouseDatabaseCreate
	ServiceClickHouseDatabaseCreate(ctx context.Context, project string, serviceName string, in *ServiceClickHouseDatabaseCreateIn) error

	// ServiceClickHouseDatabaseDelete delete a database
	// DELETE /v1/project/{project}/service/{service_name}/clickhouse/db/{database}
	// https://api.aiven.io/doc/#tag/Service:_ClickHouse/operation/ServiceClickHouseDatabaseDelete
	ServiceClickHouseDatabaseDelete(ctx context.Context, project string, serviceName string, database string) error

	// ServiceClickHouseDatabaseList list all databases
	// GET /v1/project/{project}/service/{service_name}/clickhouse/db
	// https://api.aiven.io/doc/#tag/Service:_ClickHouse/operation/ServiceClickHouseDatabaseList
	ServiceClickHouseDatabaseList(ctx context.Context, project string, serviceName string) ([]DatabaseOut, error)

	// ServiceClickHouseQuery execute an SQL query
	// POST /v1/project/{project}/service/{service_name}/clickhouse/query
	// https://api.aiven.io/doc/#tag/Service:_ClickHouse/operation/ServiceClickHouseQuery
	ServiceClickHouseQuery(ctx context.Context, project string, serviceName string, in *ServiceClickHouseQueryIn) (*ServiceClickHouseQueryOut, error)

	// ServiceClickHouseQueryStats return statistics on recent queries
	// GET /v1/project/{project}/service/{service_name}/clickhouse/query/stats
	// https://api.aiven.io/doc/#tag/Service:_ClickHouse/operation/ServiceClickHouseQueryStats
	ServiceClickHouseQueryStats(ctx context.Context, project string, serviceName string, query ...serviceClickHouseQueryStatsQuery) ([]QueryOutAlt, error)

	// ServiceClickHouseTieredStorageSummary get the ClickHouse tiered storage summary
	// GET /v1/project/{project}/service/{service_name}/clickhouse/tiered-storage/summary
	// https://api.aiven.io/doc/#tag/Service:_ClickHouse/operation/ServiceClickHouseTieredStorageSummary
	ServiceClickHouseTieredStorageSummary(ctx context.Context, project string, serviceName string) (*ServiceClickHouseTieredStorageSummaryOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) ClickHouseHandler {
	return ClickHouseHandler{doer}
}

type ClickHouseHandler struct {
	doer doer
}

func (h *ClickHouseHandler) ServiceClickHouseCurrentQueries(ctx context.Context, project string, serviceName string) ([]QueryOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/clickhouse/query", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceClickHouseCurrentQueries", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceClickHouseCurrentQueriesOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Queries, nil
}
func (h *ClickHouseHandler) ServiceClickHouseDatabaseCreate(ctx context.Context, project string, serviceName string, in *ServiceClickHouseDatabaseCreateIn) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/clickhouse/db", url.PathEscape(project), url.PathEscape(serviceName))
	_, err := h.doer.Do(ctx, "ServiceClickHouseDatabaseCreate", "POST", path, in)
	return err
}
func (h *ClickHouseHandler) ServiceClickHouseDatabaseDelete(ctx context.Context, project string, serviceName string, database string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/clickhouse/db/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(database))
	_, err := h.doer.Do(ctx, "ServiceClickHouseDatabaseDelete", "DELETE", path, nil)
	return err
}
func (h *ClickHouseHandler) ServiceClickHouseDatabaseList(ctx context.Context, project string, serviceName string) ([]DatabaseOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/clickhouse/db", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceClickHouseDatabaseList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceClickHouseDatabaseListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Databases, nil
}
func (h *ClickHouseHandler) ServiceClickHouseQuery(ctx context.Context, project string, serviceName string, in *ServiceClickHouseQueryIn) (*ServiceClickHouseQueryOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/clickhouse/query", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceClickHouseQuery", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceClickHouseQueryOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// serviceClickHouseQueryStatsQuery http query params private type

type serviceClickHouseQueryStatsQuery [2]string

// ServiceClickHouseQueryStatsLimit Limit for number of results
func ServiceClickHouseQueryStatsLimit(limit int) serviceClickHouseQueryStatsQuery {
	return serviceClickHouseQueryStatsQuery{"limit", fmt.Sprintf("%d", limit)}
}

// ServiceClickHouseQueryStatsOffset Offset for retrieved results based on sort order
func ServiceClickHouseQueryStatsOffset(offset int) serviceClickHouseQueryStatsQuery {
	return serviceClickHouseQueryStatsQuery{"offset", fmt.Sprintf("%d", offset)}
}

// ServiceClickHouseQueryStatsOrderByType Order in which to sort retrieved results
func ServiceClickHouseQueryStatsOrderByType(orderByType OrderByType) serviceClickHouseQueryStatsQuery {
	return serviceClickHouseQueryStatsQuery{"order_by", fmt.Sprintf("%s", orderByType)}
}
func (h *ClickHouseHandler) ServiceClickHouseQueryStats(ctx context.Context, project string, serviceName string, query ...serviceClickHouseQueryStatsQuery) ([]QueryOutAlt, error) {
	p := make([][2]string, 0, len(query))
	for _, v := range query {
		p = append(p, v)
	}
	path := fmt.Sprintf("/v1/project/%s/service/%s/clickhouse/query/stats", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceClickHouseQueryStats", "GET", path, nil, p...)
	if err != nil {
		return nil, err
	}
	out := new(serviceClickHouseQueryStatsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Queries, nil
}
func (h *ClickHouseHandler) ServiceClickHouseTieredStorageSummary(ctx context.Context, project string, serviceName string) (*ServiceClickHouseTieredStorageSummaryOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/clickhouse/tiered-storage/summary", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceClickHouseTieredStorageSummary", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(ServiceClickHouseTieredStorageSummaryOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DatabaseOut struct {
	Engine   string            `json:"engine"`          // Database engine
	Name     string            `json:"name"`            // Database name
	Required bool              `json:"required"`        // Required database
	State    DatabaseStateType `json:"state,omitempty"` // Database state
}
type DatabaseStateType string

const (
	DatabaseStateTypeOk              DatabaseStateType = "ok"
	DatabaseStateTypePendingCreation DatabaseStateType = "pending_creation"
	DatabaseStateTypePendingRemoval  DatabaseStateType = "pending_removal"
)

func DatabaseStateTypeChoices() []string {
	return []string{"ok", "pending_creation", "pending_removal"}
}

type HourlyOut struct {
	EstimatedCost   *string `json:"estimated_cost,omitempty"` // The estimated cost in USD of tiered storage for this hour
	HourStart       string  `json:"hour_start"`               // Timestamp in ISO 8601 format, always in UTC
	PeakStoredBytes int     `json:"peak_stored_bytes"`        // Peak bytes stored on object storage at this hour
}
type MetaOut struct {
	Name string `json:"name"` // Column name
	Type string `json:"type"` // Column type
}
type OrderByType string

const (
	OrderByTypeCallsasc       OrderByType = "calls:asc"
	OrderByTypeCallsdesc      OrderByType = "calls:desc"
	OrderByTypeMinTimeasc     OrderByType = "min_time:asc"
	OrderByTypeMinTimedesc    OrderByType = "min_time:desc"
	OrderByTypeMaxTimeasc     OrderByType = "max_time:asc"
	OrderByTypeMaxTimedesc    OrderByType = "max_time:desc"
	OrderByTypeMeanTimeasc    OrderByType = "mean_time:asc"
	OrderByTypeMeanTimedesc   OrderByType = "mean_time:desc"
	OrderByTypeP95Timeasc     OrderByType = "p95_time:asc"
	OrderByTypeP95Timedesc    OrderByType = "p95_time:desc"
	OrderByTypeStddevTimeasc  OrderByType = "stddev_time:asc"
	OrderByTypeStddevTimedesc OrderByType = "stddev_time:desc"
	OrderByTypeTotalTimeasc   OrderByType = "total_time:asc"
	OrderByTypeTotalTimedesc  OrderByType = "total_time:desc"
)

func OrderByTypeChoices() []string {
	return []string{"calls:asc", "calls:desc", "min_time:asc", "min_time:desc", "max_time:asc", "max_time:desc", "mean_time:asc", "mean_time:desc", "p95_time:asc", "p95_time:desc", "stddev_time:asc", "stddev_time:desc", "total_time:asc", "total_time:desc"}
}

type QueryOut struct {
	ClientName *string  `json:"client_name,omitempty"` // Client name, if set
	Database   *string  `json:"database,omitempty"`
	Elapsed    *float64 `json:"elapsed,omitempty"` // The time in seconds since request execution started
	Query      *string  `json:"query,omitempty"`   // The query text
	User       *string  `json:"user,omitempty"`    // The user who made the query
}
type QueryOutAlt struct {
	Calls      *int     `json:"calls,omitempty"` // Number of calls
	Database   *string  `json:"database,omitempty"`
	MaxTime    *int     `json:"max_time,omitempty"`    // Maximum query duration in milliseconds
	MeanTime   *int     `json:"mean_time,omitempty"`   // Average query duration in milliseconds
	MinTime    *int     `json:"min_time,omitempty"`    // Minimum query duration in milliseconds
	P95Time    *int     `json:"p95_time,omitempty"`    // Query duration 95th percentile in milliseconds
	Query      *string  `json:"query,omitempty"`       // Normalized query
	Rows       *float64 `json:"rows,omitempty"`        // Average number of rows per call
	StddevTime *int     `json:"stddev_time,omitempty"` // Query duration standard deviation in milliseconds
	TotalTime  *int     `json:"total_time,omitempty"`  // Total duration of all calls in milliseconds
}

// ServiceClickHouseDatabaseCreateIn ServiceClickHouseDatabaseCreateRequestBody
type ServiceClickHouseDatabaseCreateIn struct {
	Database string `json:"database"` // Service database name
}

// ServiceClickHouseQueryIn ServiceClickHouseQueryRequestBody
type ServiceClickHouseQueryIn struct {
	Database string `json:"database"` // Service database name
	Query    string `json:"query"`
}

// ServiceClickHouseQueryOut ServiceClickHouseQueryResponse
type ServiceClickHouseQueryOut struct {
	Data    [][]any    `json:"data"`
	Meta    []MetaOut  `json:"meta"`
	Summary SummaryOut `json:"summary"` // Summary
}

// ServiceClickHouseTieredStorageSummaryOut ServiceClickHouseTieredStorageSummaryResponse
type ServiceClickHouseTieredStorageSummaryOut struct {
	CurrentCost         string                 `json:"current_cost"`              // The current cost in USD of tiered storage since the beginning of the billing period
	ForecastedCost      string                 `json:"forecasted_cost"`           // The forecasted cost in USD of tiered storage in the billing period
	ForecastedRate      *string                `json:"forecasted_rate,omitempty"` // The rate on GBs/hour used to calculate the forecasted cost
	StorageUsageHistory StorageUsageHistoryOut `json:"storage_usage_history"`     // History of usage and cumulative costs in the billing period
	TotalStorageUsage   int                    `json:"total_storage_usage"`       // Total storage usage by tiered storage, in bytes
}

// StorageUsageHistoryOut History of usage and cumulative costs in the billing period
type StorageUsageHistoryOut struct {
	Hourly []HourlyOut `json:"hourly"` // History by hour
}

// SummaryOut Summary
type SummaryOut struct {
	ElapsedNs    *int `json:"elapsed_ns,omitempty"`    // Elapsed time in nanoseconds
	ReadBytes    *int `json:"read_bytes,omitempty"`    // Number of bytes read
	ReadRows     *int `json:"read_rows,omitempty"`     // Number of rows read
	ResultBytes  *int `json:"result_bytes,omitempty"`  // Number of bytes in the result
	ResultRows   *int `json:"result_rows,omitempty"`   // Number of rows in the result
	WrittenBytes *int `json:"written_bytes,omitempty"` // Number of bytes written
	WrittenRows  *int `json:"written_rows,omitempty"`  // Number of rows written
}

// serviceClickHouseCurrentQueriesOut ServiceClickHouseCurrentQueriesResponse
type serviceClickHouseCurrentQueriesOut struct {
	Queries []QueryOut `json:"queries"` // List of currently running queries
}

// serviceClickHouseDatabaseListOut ServiceClickHouseDatabaseListResponse
type serviceClickHouseDatabaseListOut struct {
	Databases []DatabaseOut `json:"databases"` // List of databases
}

// serviceClickHouseQueryStatsOut ServiceClickHouseQueryStatsResponse
type serviceClickHouseQueryStatsOut struct {
	Queries []QueryOutAlt `json:"queries"` // List of query statistics
}
