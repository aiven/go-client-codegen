// Code generated by Aiven. DO NOT EDIT.

package kafkamirrormaker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// ServiceKafkaMirrorMakerCreateReplicationFlow create a replication flow
	// POST /v1/project/{project}/service/{service_name}/mirrormaker/replication-flows
	// https://api.aiven.io/doc/#tag/Service:_Kafka_MirrorMaker/operation/ServiceKafkaMirrorMakerCreateReplicationFlow
	ServiceKafkaMirrorMakerCreateReplicationFlow(ctx context.Context, project string, serviceName string, in *ServiceKafkaMirrorMakerCreateReplicationFlowIn) error

	// ServiceKafkaMirrorMakerDeleteReplicationFlow delete a replication flow
	// DELETE /v1/project/{project}/service/{service_name}/mirrormaker/replication-flows/{source_cluster}/{target_cluster}
	// https://api.aiven.io/doc/#tag/Service:_Kafka_MirrorMaker/operation/ServiceKafkaMirrorMakerDeleteReplicationFlow
	ServiceKafkaMirrorMakerDeleteReplicationFlow(ctx context.Context, project string, serviceName string, sourceCluster string, targetCluster string) error

	// ServiceKafkaMirrorMakerGetReplicationFlow get a replication flow
	// GET /v1/project/{project}/service/{service_name}/mirrormaker/replication-flows/{source_cluster}/{target_cluster}
	// https://api.aiven.io/doc/#tag/Service:_Kafka_MirrorMaker/operation/ServiceKafkaMirrorMakerGetReplicationFlow
	ServiceKafkaMirrorMakerGetReplicationFlow(ctx context.Context, project string, serviceName string, sourceCluster string, targetCluster string) (*ServiceKafkaMirrorMakerGetReplicationFlowOut, error)

	// ServiceKafkaMirrorMakerGetReplicationFlows get replication flows
	// GET /v1/project/{project}/service/{service_name}/mirrormaker/replication-flows
	// https://api.aiven.io/doc/#tag/Service:_Kafka_MirrorMaker/operation/ServiceKafkaMirrorMakerGetReplicationFlows
	ServiceKafkaMirrorMakerGetReplicationFlows(ctx context.Context, project string, serviceName string) ([]ReplicationFlowOut, error)

	// ServiceKafkaMirrorMakerPatchReplicationFlow update a replication flow
	// PUT /v1/project/{project}/service/{service_name}/mirrormaker/replication-flows/{source_cluster}/{target_cluster}
	// https://api.aiven.io/doc/#tag/Service:_Kafka_MirrorMaker/operation/ServiceKafkaMirrorMakerPatchReplicationFlow
	ServiceKafkaMirrorMakerPatchReplicationFlow(ctx context.Context, project string, serviceName string, sourceCluster string, targetCluster string, in *ServiceKafkaMirrorMakerPatchReplicationFlowIn) (*ServiceKafkaMirrorMakerPatchReplicationFlowOut, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) KafkaMirrorMakerHandler {
	return KafkaMirrorMakerHandler{doer}
}

type KafkaMirrorMakerHandler struct {
	doer doer
}

func (h *KafkaMirrorMakerHandler) ServiceKafkaMirrorMakerCreateReplicationFlow(ctx context.Context, project string, serviceName string, in *ServiceKafkaMirrorMakerCreateReplicationFlowIn) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/mirrormaker/replication-flows", url.PathEscape(project), url.PathEscape(serviceName))
	_, err := h.doer.Do(ctx, "ServiceKafkaMirrorMakerCreateReplicationFlow", "POST", path, in)
	return err
}
func (h *KafkaMirrorMakerHandler) ServiceKafkaMirrorMakerDeleteReplicationFlow(ctx context.Context, project string, serviceName string, sourceCluster string, targetCluster string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/mirrormaker/replication-flows/%s/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(sourceCluster), url.PathEscape(targetCluster))
	_, err := h.doer.Do(ctx, "ServiceKafkaMirrorMakerDeleteReplicationFlow", "DELETE", path, nil)
	return err
}
func (h *KafkaMirrorMakerHandler) ServiceKafkaMirrorMakerGetReplicationFlow(ctx context.Context, project string, serviceName string, sourceCluster string, targetCluster string) (*ServiceKafkaMirrorMakerGetReplicationFlowOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/mirrormaker/replication-flows/%s/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(sourceCluster), url.PathEscape(targetCluster))
	b, err := h.doer.Do(ctx, "ServiceKafkaMirrorMakerGetReplicationFlow", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaMirrorMakerGetReplicationFlowOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.ReplicationFlow, nil
}
func (h *KafkaMirrorMakerHandler) ServiceKafkaMirrorMakerGetReplicationFlows(ctx context.Context, project string, serviceName string) ([]ReplicationFlowOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/mirrormaker/replication-flows", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceKafkaMirrorMakerGetReplicationFlows", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaMirrorMakerGetReplicationFlowsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.ReplicationFlows, nil
}
func (h *KafkaMirrorMakerHandler) ServiceKafkaMirrorMakerPatchReplicationFlow(ctx context.Context, project string, serviceName string, sourceCluster string, targetCluster string, in *ServiceKafkaMirrorMakerPatchReplicationFlowIn) (*ServiceKafkaMirrorMakerPatchReplicationFlowOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/mirrormaker/replication-flows/%s/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(sourceCluster), url.PathEscape(targetCluster))
	b, err := h.doer.Do(ctx, "ServiceKafkaMirrorMakerPatchReplicationFlow", "PUT", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceKafkaMirrorMakerPatchReplicationFlowOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.ReplicationFlow, nil
}

type OffsetSyncsTopicLocationType string

const (
	OffsetSyncsTopicLocationTypeSource OffsetSyncsTopicLocationType = "source"
	OffsetSyncsTopicLocationTypeTarget OffsetSyncsTopicLocationType = "target"
)

func OffsetSyncsTopicLocationTypeChoices() []string {
	return []string{"source", "target"}
}

type ReplicationFlowOut struct {
	ConfigPropertiesExclude         *string                      `json:"config_properties_exclude,omitempty"`           // Topic configuration properties that should not be replicated
	EmitBackwardHeartbeatsEnabled   *bool                        `json:"emit_backward_heartbeats_enabled,omitempty"`    // Emit backward heartbeats enabled
	EmitHeartbeatsEnabled           *bool                        `json:"emit_heartbeats_enabled,omitempty"`             // Emit heartbeats enabled
	Enabled                         bool                         `json:"enabled"`                                       // Is replication flow enabled
	OffsetLagMax                    *int                         `json:"offset_lag_max,omitempty"`                      // How out-of-sync a remote partition can be before it is resynced
	OffsetSyncsTopicLocation        OffsetSyncsTopicLocationType `json:"offset_syncs_topic_location,omitempty"`         // Offset syncs topic location
	ReplicationFactor               *int                         `json:"replication_factor,omitempty"`                  // Replication factor
	ReplicationPolicyClass          ReplicationPolicyClassType   `json:"replication_policy_class,omitempty"`            // Replication policy class
	ReplicationProgress             *float64                     `json:"replication_progress,omitempty"`                // Replication progress
	SourceCluster                   string                       `json:"source_cluster"`                                // Source cluster alias
	SyncGroupOffsetsEnabled         *bool                        `json:"sync_group_offsets_enabled,omitempty"`          // Sync consumer group offsets
	SyncGroupOffsetsIntervalSeconds *int                         `json:"sync_group_offsets_interval_seconds,omitempty"` // Frequency of consumer group offset sync
	TargetCluster                   string                       `json:"target_cluster"`                                // Target cluster alias
	Topics                          []string                     `json:"topics,omitempty"`                              // List of topics and/or regular expressions to replicate. Topic names and regular expressions that match topic names that should be replicated. MirrorMaker will replicate these topics if they are not matched by "topics.exclude". Currently defaults to [".*"].
	TopicsBlacklist                 []string                     `json:"topics.blacklist,omitempty"`                    // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
	TopicsExclude                   []string                     `json:"topics.exclude,omitempty"`                      // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
}
type ReplicationPolicyClassType string

const (
	ReplicationPolicyClassTypeDefault  ReplicationPolicyClassType = "org.apache.kafka.connect.mirror.DefaultReplicationPolicy"
	ReplicationPolicyClassTypeIdentity ReplicationPolicyClassType = "org.apache.kafka.connect.mirror.IdentityReplicationPolicy"
)

func ReplicationPolicyClassTypeChoices() []string {
	return []string{"org.apache.kafka.connect.mirror.DefaultReplicationPolicy", "org.apache.kafka.connect.mirror.IdentityReplicationPolicy"}
}

// ServiceKafkaMirrorMakerCreateReplicationFlowIn ServiceKafkaMirrorMakerCreateReplicationFlowRequestBody
type ServiceKafkaMirrorMakerCreateReplicationFlowIn struct {
	ConfigPropertiesExclude         *string                      `json:"config_properties_exclude,omitempty"`           // Topic configuration properties that should not be replicated
	EmitBackwardHeartbeatsEnabled   *bool                        `json:"emit_backward_heartbeats_enabled,omitempty"`    // Emit backward heartbeats enabled
	EmitHeartbeatsEnabled           *bool                        `json:"emit_heartbeats_enabled,omitempty"`             // Emit heartbeats enabled
	Enabled                         bool                         `json:"enabled"`                                       // Is replication flow enabled
	OffsetLagMax                    *int                         `json:"offset_lag_max,omitempty"`                      // How out-of-sync a remote partition can be before it is resynced
	OffsetSyncsTopicLocation        OffsetSyncsTopicLocationType `json:"offset_syncs_topic_location,omitempty"`         // Offset syncs topic location
	ReplicationFactor               *int                         `json:"replication_factor,omitempty"`                  // Replication factor
	ReplicationPolicyClass          ReplicationPolicyClassType   `json:"replication_policy_class,omitempty"`            // Replication policy class
	SourceCluster                   string                       `json:"source_cluster"`                                // Source cluster alias
	SyncGroupOffsetsEnabled         *bool                        `json:"sync_group_offsets_enabled,omitempty"`          // Sync consumer group offsets
	SyncGroupOffsetsIntervalSeconds *int                         `json:"sync_group_offsets_interval_seconds,omitempty"` // Frequency of consumer group offset sync
	TargetCluster                   string                       `json:"target_cluster"`                                // Target cluster alias
	Topics                          *[]string                    `json:"topics,omitempty"`                              // List of topics and/or regular expressions to replicate. Topic names and regular expressions that match topic names that should be replicated. MirrorMaker will replicate these topics if they are not matched by "topics.exclude". Currently defaults to [".*"].
	TopicsBlacklist                 *[]string                    `json:"topics.blacklist,omitempty"`                    // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
	TopicsExclude                   *[]string                    `json:"topics.exclude,omitempty"`                      // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
}

// ServiceKafkaMirrorMakerGetReplicationFlowOut Replication flow
type ServiceKafkaMirrorMakerGetReplicationFlowOut struct {
	ConfigPropertiesExclude         *string                      `json:"config_properties_exclude,omitempty"`           // Topic configuration properties that should not be replicated
	EmitBackwardHeartbeatsEnabled   *bool                        `json:"emit_backward_heartbeats_enabled,omitempty"`    // Emit backward heartbeats enabled
	EmitHeartbeatsEnabled           *bool                        `json:"emit_heartbeats_enabled,omitempty"`             // Emit heartbeats enabled
	Enabled                         bool                         `json:"enabled"`                                       // Is replication flow enabled
	OffsetLagMax                    *int                         `json:"offset_lag_max,omitempty"`                      // How out-of-sync a remote partition can be before it is resynced
	OffsetSyncsTopicLocation        OffsetSyncsTopicLocationType `json:"offset_syncs_topic_location,omitempty"`         // Offset syncs topic location
	ReplicationFactor               *int                         `json:"replication_factor,omitempty"`                  // Replication factor
	ReplicationPolicyClass          ReplicationPolicyClassType   `json:"replication_policy_class,omitempty"`            // Replication policy class
	ReplicationProgress             *float64                     `json:"replication_progress,omitempty"`                // Replication progress
	SourceCluster                   string                       `json:"source_cluster"`                                // Source cluster alias
	SyncGroupOffsetsEnabled         *bool                        `json:"sync_group_offsets_enabled,omitempty"`          // Sync consumer group offsets
	SyncGroupOffsetsIntervalSeconds *int                         `json:"sync_group_offsets_interval_seconds,omitempty"` // Frequency of consumer group offset sync
	TargetCluster                   string                       `json:"target_cluster"`                                // Target cluster alias
	Topics                          []string                     `json:"topics,omitempty"`                              // List of topics and/or regular expressions to replicate. Topic names and regular expressions that match topic names that should be replicated. MirrorMaker will replicate these topics if they are not matched by "topics.exclude". Currently defaults to [".*"].
	TopicsBlacklist                 []string                     `json:"topics.blacklist,omitempty"`                    // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
	TopicsExclude                   []string                     `json:"topics.exclude,omitempty"`                      // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
}

// ServiceKafkaMirrorMakerPatchReplicationFlowIn ServiceKafkaMirrorMakerPatchReplicationFlowRequestBody
type ServiceKafkaMirrorMakerPatchReplicationFlowIn struct {
	ConfigPropertiesExclude         *string                      `json:"config_properties_exclude,omitempty"`           // Topic configuration properties that should not be replicated
	EmitBackwardHeartbeatsEnabled   *bool                        `json:"emit_backward_heartbeats_enabled,omitempty"`    // Emit backward heartbeats enabled
	EmitHeartbeatsEnabled           *bool                        `json:"emit_heartbeats_enabled,omitempty"`             // Emit heartbeats enabled
	Enabled                         *bool                        `json:"enabled,omitempty"`                             // Is replication flow enabled
	OffsetLagMax                    *int                         `json:"offset_lag_max,omitempty"`                      // How out-of-sync a remote partition can be before it is resynced
	OffsetSyncsTopicLocation        OffsetSyncsTopicLocationType `json:"offset_syncs_topic_location,omitempty"`         // Offset syncs topic location
	ReplicationFactor               *int                         `json:"replication_factor,omitempty"`                  // Replication factor
	ReplicationPolicyClass          ReplicationPolicyClassType   `json:"replication_policy_class,omitempty"`            // Replication policy class
	SyncGroupOffsetsEnabled         *bool                        `json:"sync_group_offsets_enabled,omitempty"`          // Sync consumer group offsets
	SyncGroupOffsetsIntervalSeconds *int                         `json:"sync_group_offsets_interval_seconds,omitempty"` // Frequency of consumer group offset sync
	Topics                          *[]string                    `json:"topics,omitempty"`                              // List of topics and/or regular expressions to replicate. Topic names and regular expressions that match topic names that should be replicated. MirrorMaker will replicate these topics if they are not matched by "topics.exclude". Currently defaults to [".*"].
	TopicsBlacklist                 *[]string                    `json:"topics.blacklist,omitempty"`                    // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
	TopicsExclude                   *[]string                    `json:"topics.exclude,omitempty"`                      // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
}

// ServiceKafkaMirrorMakerPatchReplicationFlowOut Replication flow
type ServiceKafkaMirrorMakerPatchReplicationFlowOut struct {
	ConfigPropertiesExclude         *string                      `json:"config_properties_exclude,omitempty"`           // Topic configuration properties that should not be replicated
	EmitBackwardHeartbeatsEnabled   *bool                        `json:"emit_backward_heartbeats_enabled,omitempty"`    // Emit backward heartbeats enabled
	EmitHeartbeatsEnabled           *bool                        `json:"emit_heartbeats_enabled,omitempty"`             // Emit heartbeats enabled
	Enabled                         bool                         `json:"enabled"`                                       // Is replication flow enabled
	OffsetLagMax                    *int                         `json:"offset_lag_max,omitempty"`                      // How out-of-sync a remote partition can be before it is resynced
	OffsetSyncsTopicLocation        OffsetSyncsTopicLocationType `json:"offset_syncs_topic_location,omitempty"`         // Offset syncs topic location
	ReplicationFactor               *int                         `json:"replication_factor,omitempty"`                  // Replication factor
	ReplicationPolicyClass          ReplicationPolicyClassType   `json:"replication_policy_class,omitempty"`            // Replication policy class
	ReplicationProgress             *float64                     `json:"replication_progress,omitempty"`                // Replication progress
	SourceCluster                   string                       `json:"source_cluster"`                                // Source cluster alias
	SyncGroupOffsetsEnabled         *bool                        `json:"sync_group_offsets_enabled,omitempty"`          // Sync consumer group offsets
	SyncGroupOffsetsIntervalSeconds *int                         `json:"sync_group_offsets_interval_seconds,omitempty"` // Frequency of consumer group offset sync
	TargetCluster                   string                       `json:"target_cluster"`                                // Target cluster alias
	Topics                          []string                     `json:"topics,omitempty"`                              // List of topics and/or regular expressions to replicate. Topic names and regular expressions that match topic names that should be replicated. MirrorMaker will replicate these topics if they are not matched by "topics.exclude". Currently defaults to [".*"].
	TopicsBlacklist                 []string                     `json:"topics.blacklist,omitempty"`                    // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
	TopicsExclude                   []string                     `json:"topics.exclude,omitempty"`                      // List of topics and/or regular expressions to not replicate. Topic names and regular expressions that match topic names that should not be replicated. MirrorMaker will not replicate these topics even if they are matched by "topics". If not set, MM2 uses the default exclusion.
}

// serviceKafkaMirrorMakerGetReplicationFlowOut ServiceKafkaMirrorMakerGetReplicationFlowResponse
type serviceKafkaMirrorMakerGetReplicationFlowOut struct {
	ReplicationFlow ServiceKafkaMirrorMakerGetReplicationFlowOut `json:"replication_flow"` // Replication flow
}

// serviceKafkaMirrorMakerGetReplicationFlowsOut ServiceKafkaMirrorMakerGetReplicationFlowsResponse
type serviceKafkaMirrorMakerGetReplicationFlowsOut struct {
	ReplicationFlows []ReplicationFlowOut `json:"replication_flows"` // Replication flows. Describes data replication flows between Kafka clusters
}

// serviceKafkaMirrorMakerPatchReplicationFlowOut ServiceKafkaMirrorMakerPatchReplicationFlowResponse
type serviceKafkaMirrorMakerPatchReplicationFlowOut struct {
	ReplicationFlow ServiceKafkaMirrorMakerPatchReplicationFlowOut `json:"replication_flow"` // Replication flow
}
