// Code generated by Aiven. DO NOT EDIT.

package mysql

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// MySQLServiceQueryStatistics fetch MySQL service query statistics
	// POST /project/{project}/service/{service_name}/mysql/query/stats
	// https://api.aiven.io/doc/#tag/Service:_MySQL/operation/MySQLServiceQueryStatistics
	MySQLServiceQueryStatistics(ctx context.Context, project string, serviceName string, in *MySqlserviceQueryStatisticsIn) ([]QueryOut, error)
}

func NewHandler(doer doer) MySQLHandler {
	return MySQLHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type MySQLHandler struct {
	doer doer
}

func (h *MySQLHandler) MySQLServiceQueryStatistics(ctx context.Context, project string, serviceName string, in *MySqlserviceQueryStatisticsIn) ([]QueryOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/mysql/query/stats", project, serviceName)
	b, err := h.doer.Do(ctx, "MySQLServiceQueryStatistics", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(mySqlserviceQueryStatisticsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Queries, nil
}

type MySqlserviceQueryStatisticsIn struct {
	Limit   *int   `json:"limit,omitempty"`
	Offset  *int   `json:"offset,omitempty"`
	OrderBy string `json:"order_by,omitempty"`
}
type QueryOut struct {
	AvgTimerWait            *float64 `json:"avg_timer_wait,omitempty"`
	CountStar               *float64 `json:"count_star,omitempty"`
	Digest                  string   `json:"digest,omitempty"`
	DigestText              string   `json:"digest_text,omitempty"`
	FirstSeen               string   `json:"first_seen,omitempty"`
	LastSeen                string   `json:"last_seen,omitempty"`
	MaxTimerWait            *float64 `json:"max_timer_wait,omitempty"`
	MinTimerWait            *float64 `json:"min_timer_wait,omitempty"`
	Quantile95              *float64 `json:"quantile_95,omitempty"`
	Quantile99              *float64 `json:"quantile_99,omitempty"`
	Quantile999             *float64 `json:"quantile_999,omitempty"`
	QuerySampleSeen         string   `json:"query_sample_seen,omitempty"`
	QuerySampleText         string   `json:"query_sample_text,omitempty"`
	QuerySampleTimerWait    string   `json:"query_sample_timer_wait,omitempty"`
	SchemaName              string   `json:"schema_name,omitempty"`
	SumCreatedTmpDiskTables *float64 `json:"sum_created_tmp_disk_tables,omitempty"`
	SumCreatedTmpTables     *float64 `json:"sum_created_tmp_tables,omitempty"`
	SumErrors               *float64 `json:"sum_errors,omitempty"`
	SumLockTime             *float64 `json:"sum_lock_time,omitempty"`
	SumNoGoodIndexUsed      *float64 `json:"sum_no_good_index_used,omitempty"`
	SumNoIndexUsed          *float64 `json:"sum_no_index_used,omitempty"`
	SumRowsAffected         *float64 `json:"sum_rows_affected,omitempty"`
	SumRowsExamined         *float64 `json:"sum_rows_examined,omitempty"`
	SumRowsSent             *float64 `json:"sum_rows_sent,omitempty"`
	SumSelectFullJoin       *float64 `json:"sum_select_full_join,omitempty"`
	SumSelectFullRangeJoin  *float64 `json:"sum_select_full_range_join,omitempty"`
	SumSelectRange          *float64 `json:"sum_select_range,omitempty"`
	SumSelectRangeCheck     *float64 `json:"sum_select_range_check,omitempty"`
	SumSelectScan           *float64 `json:"sum_select_scan,omitempty"`
	SumSortMergePasses      *float64 `json:"sum_sort_merge_passes,omitempty"`
	SumSortRange            *float64 `json:"sum_sort_range,omitempty"`
	SumSortRows             *float64 `json:"sum_sort_rows,omitempty"`
	SumSortScan             *float64 `json:"sum_sort_scan,omitempty"`
	SumTimerWait            *float64 `json:"sum_timer_wait,omitempty"`
	SumWarnings             *float64 `json:"sum_warnings,omitempty"`
}
type mySqlserviceQueryStatisticsOut struct {
	Queries []QueryOut `json:"queries"`
}
