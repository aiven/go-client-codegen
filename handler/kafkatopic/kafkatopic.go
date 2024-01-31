// Code generated by Aiven. DO NOT EDIT.

package kafkatopic

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// ServiceKafkaTopicCreate create a Kafka topic
	// POST /project/{project}/service/{service_name}/topic
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicCreate
	ServiceKafkaTopicCreate(ctx context.Context, project string, serviceName string, in *ServiceKafkaTopicCreateIn) error

	// ServiceKafkaTopicDelete delete a Kafka topic
	// DELETE /project/{project}/service/{service_name}/topic/{topic_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicDelete
	ServiceKafkaTopicDelete(ctx context.Context, project string, serviceName string, topicName string) error

	// ServiceKafkaTopicGet get Kafka topic info
	// GET /project/{project}/service/{service_name}/topic/{topic_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicGet
	ServiceKafkaTopicGet(ctx context.Context, project string, serviceName string, topicName string) (*ServiceKafkaTopicGetOut, error)

	// ServiceKafkaTopicList get Kafka topic list
	// GET /project/{project}/service/{service_name}/topic
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicList
	ServiceKafkaTopicList(ctx context.Context, project string, serviceName string) ([]TopicOut, error)

	// ServiceKafkaTopicMessageList list kafka topic messages
	// POST /project/{project}/service/{service_name}/kafka/rest/topics/{topic_name}/messages
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicMessageList
	ServiceKafkaTopicMessageList(ctx context.Context, project string, serviceName string, topicName string, in *ServiceKafkaTopicMessageListIn) ([]MessageOut, error)

	// ServiceKafkaTopicMessageProduce produce message into a kafka topic
	// POST /project/{project}/service/{service_name}/kafka/rest/topics/{topic_name}/produce
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicMessageProduce
	ServiceKafkaTopicMessageProduce(ctx context.Context, project string, serviceName string, topicName string, in *ServiceKafkaTopicMessageProduceIn) (*ServiceKafkaTopicMessageProduceOut, error)

	// ServiceKafkaTopicUpdate update a Kafka topic
	// PUT /project/{project}/service/{service_name}/topic/{topic_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceKafkaTopicUpdate
	ServiceKafkaTopicUpdate(ctx context.Context, project string, serviceName string, topicName string, in *ServiceKafkaTopicUpdateIn) error
}

func NewHandler(doer doer) KafkaTopicHandler {
	return KafkaTopicHandler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type KafkaTopicHandler struct {
	doer doer
}

func (h *KafkaTopicHandler) ServiceKafkaTopicCreate(ctx context.Context, project string, serviceName string, in *ServiceKafkaTopicCreateIn) error {
	path := fmt.Sprintf("/project/%s/service/%s/topic", project, serviceName)
	_, err := h.doer.Do(ctx, "ServiceKafkaTopicCreate", "POST", path, in)
	return err
}
func (h *KafkaTopicHandler) ServiceKafkaTopicDelete(ctx context.Context, project string, serviceName string, topicName string) error {
	path := fmt.Sprintf("/project/%s/service/%s/topic/%s", project, serviceName, topicName)
	_, err := h.doer.Do(ctx, "ServiceKafkaTopicDelete", "DELETE", path, nil)
	return err
}
func (h *KafkaTopicHandler) ServiceKafkaTopicGet(ctx context.Context, project string, serviceName string, topicName string) (*ServiceKafkaTopicGetOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/topic/%s", project, serviceName, topicName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTopicGet", "GET", path, nil)
	out := new(serviceKafkaTopicGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Topic, nil
}
func (h *KafkaTopicHandler) ServiceKafkaTopicList(ctx context.Context, project string, serviceName string) ([]TopicOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/topic", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTopicList", "GET", path, nil)
	out := new(serviceKafkaTopicListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Topics, nil
}
func (h *KafkaTopicHandler) ServiceKafkaTopicMessageList(ctx context.Context, project string, serviceName string, topicName string, in *ServiceKafkaTopicMessageListIn) ([]MessageOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/rest/topics/%s/messages", project, serviceName, topicName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTopicMessageList", "POST", path, in)
	out := new(serviceKafkaTopicMessageListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Messages, nil
}
func (h *KafkaTopicHandler) ServiceKafkaTopicMessageProduce(ctx context.Context, project string, serviceName string, topicName string, in *ServiceKafkaTopicMessageProduceIn) (*ServiceKafkaTopicMessageProduceOut, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/rest/topics/%s/produce", project, serviceName, topicName)
	b, err := h.doer.Do(ctx, "ServiceKafkaTopicMessageProduce", "POST", path, in)
	out := new(ServiceKafkaTopicMessageProduceOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *KafkaTopicHandler) ServiceKafkaTopicUpdate(ctx context.Context, project string, serviceName string, topicName string, in *ServiceKafkaTopicUpdateIn) error {
	path := fmt.Sprintf("/project/%s/service/%s/topic/%s", project, serviceName, topicName)
	_, err := h.doer.Do(ctx, "ServiceKafkaTopicUpdate", "PUT", path, in)
	return err
}

type CleanupPolicyOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    string       `json:"value,omitempty"`
}
type CleanupPolicyType string

const (
	CleanupPolicyTypeDelete        CleanupPolicyType = "delete"
	CleanupPolicyTypeCompact       CleanupPolicyType = "compact"
	CleanupPolicyTypeCompactdelete CleanupPolicyType = "compact,delete"
)

func CleanupPolicyTypeChoices() []string {
	return []string{"delete", "compact", "compact,delete"}
}

type CompressionType string

const (
	CompressionTypeSnappy       CompressionType = "snappy"
	CompressionTypeGzip         CompressionType = "gzip"
	CompressionTypeLz4          CompressionType = "lz4"
	CompressionTypeProducer     CompressionType = "producer"
	CompressionTypeUncompressed CompressionType = "uncompressed"
	CompressionTypeZstd         CompressionType = "zstd"
)

func CompressionTypeChoices() []string {
	return []string{"snappy", "gzip", "lz4", "producer", "uncompressed", "zstd"}
}

type CompressionTypeOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    string       `json:"value,omitempty"`
}
type ConfigIn struct {
	CleanupPolicy                   CleanupPolicyType        `json:"cleanup_policy,omitempty"`
	CompressionType                 CompressionType          `json:"compression_type,omitempty"`
	DeleteRetentionMs               *int                     `json:"delete_retention_ms,omitempty"`
	FileDeleteDelayMs               *int                     `json:"file_delete_delay_ms,omitempty"`
	FlushMessages                   *int                     `json:"flush_messages,omitempty"`
	FlushMs                         *int                     `json:"flush_ms,omitempty"`
	IndexIntervalBytes              *int                     `json:"index_interval_bytes,omitempty"`
	LocalRetentionBytes             *int                     `json:"local_retention_bytes,omitempty"`
	LocalRetentionMs                *int                     `json:"local_retention_ms,omitempty"`
	MaxCompactionLagMs              *int                     `json:"max_compaction_lag_ms,omitempty"`
	MaxMessageBytes                 *int                     `json:"max_message_bytes,omitempty"`
	MessageDownconversionEnable     *bool                    `json:"message_downconversion_enable,omitempty"`
	MessageFormatVersion            MessageFormatVersionType `json:"message_format_version,omitempty"`
	MessageTimestampDifferenceMaxMs *int                     `json:"message_timestamp_difference_max_ms,omitempty"`
	MessageTimestampType            MessageTimestampType     `json:"message_timestamp_type,omitempty"`
	MinCleanableDirtyRatio          *float64                 `json:"min_cleanable_dirty_ratio,omitempty"`
	MinCompactionLagMs              *int                     `json:"min_compaction_lag_ms,omitempty"`
	MinInsyncReplicas               *int                     `json:"min_insync_replicas,omitempty"`
	Preallocate                     *bool                    `json:"preallocate,omitempty"`
	RemoteStorageEnable             *bool                    `json:"remote_storage_enable,omitempty"`
	RetentionBytes                  *int                     `json:"retention_bytes,omitempty"`
	RetentionMs                     *int                     `json:"retention_ms,omitempty"`
	SegmentBytes                    *int                     `json:"segment_bytes,omitempty"`
	SegmentIndexBytes               *int                     `json:"segment_index_bytes,omitempty"`
	SegmentJitterMs                 *int                     `json:"segment_jitter_ms,omitempty"`
	SegmentMs                       *int                     `json:"segment_ms,omitempty"`
	UncleanLeaderElectionEnable     *bool                    `json:"unclean_leader_election_enable,omitempty"`
}
type ConfigOut struct {
	CleanupPolicy                   *CleanupPolicyOut                   `json:"cleanup_policy,omitempty"`
	CompressionType                 *CompressionTypeOut                 `json:"compression_type,omitempty"`
	DeleteRetentionMs               *DeleteRetentionMsOut               `json:"delete_retention_ms,omitempty"`
	FileDeleteDelayMs               *FileDeleteDelayMsOut               `json:"file_delete_delay_ms,omitempty"`
	FlushMessages                   *FlushMessagesOut                   `json:"flush_messages,omitempty"`
	FlushMs                         *FlushMsOut                         `json:"flush_ms,omitempty"`
	IndexIntervalBytes              *IndexIntervalBytesOut              `json:"index_interval_bytes,omitempty"`
	LocalRetentionBytes             *LocalRetentionBytesOut             `json:"local_retention_bytes,omitempty"`
	LocalRetentionMs                *LocalRetentionMsOut                `json:"local_retention_ms,omitempty"`
	MaxCompactionLagMs              *MaxCompactionLagMsOut              `json:"max_compaction_lag_ms,omitempty"`
	MaxMessageBytes                 *MaxMessageBytesOut                 `json:"max_message_bytes,omitempty"`
	MessageDownconversionEnable     *MessageDownconversionEnableOut     `json:"message_downconversion_enable,omitempty"`
	MessageFormatVersion            *MessageFormatVersionOut            `json:"message_format_version,omitempty"`
	MessageTimestampDifferenceMaxMs *MessageTimestampDifferenceMaxMsOut `json:"message_timestamp_difference_max_ms,omitempty"`
	MessageTimestampType            *MessageTimestampTypeOut            `json:"message_timestamp_type,omitempty"`
	MinCleanableDirtyRatio          *MinCleanableDirtyRatioOut          `json:"min_cleanable_dirty_ratio,omitempty"`
	MinCompactionLagMs              *MinCompactionLagMsOut              `json:"min_compaction_lag_ms,omitempty"`
	MinInsyncReplicas               *MinInsyncReplicasOut               `json:"min_insync_replicas,omitempty"`
	Preallocate                     *PreallocateOut                     `json:"preallocate,omitempty"`
	RemoteStorageEnable             *RemoteStorageEnableOut             `json:"remote_storage_enable,omitempty"`
	RetentionBytes                  *RetentionBytesOut                  `json:"retention_bytes,omitempty"`
	RetentionMs                     *RetentionMsOut                     `json:"retention_ms,omitempty"`
	SegmentBytes                    *SegmentBytesOut                    `json:"segment_bytes,omitempty"`
	SegmentIndexBytes               *SegmentIndexBytesOut               `json:"segment_index_bytes,omitempty"`
	SegmentJitterMs                 *SegmentJitterMsOut                 `json:"segment_jitter_ms,omitempty"`
	SegmentMs                       *SegmentMsOut                       `json:"segment_ms,omitempty"`
	UncleanLeaderElectionEnable     *UncleanLeaderElectionEnableOut     `json:"unclean_leader_election_enable,omitempty"`
}
type ConsumerGroupOut struct {
	GroupName string `json:"group_name"`
	Offset    int    `json:"offset"`
}
type DeleteRetentionMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type FileDeleteDelayMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type FlushMessagesOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type FlushMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type FormatType string

const (
	FormatTypeBinary     FormatType = "binary"
	FormatTypeJson       FormatType = "json"
	FormatTypeAvro       FormatType = "avro"
	FormatTypeProtobuf   FormatType = "protobuf"
	FormatTypeJsonschema FormatType = "jsonschema"
)

func FormatTypeChoices() []string {
	return []string{"binary", "json", "avro", "protobuf", "jsonschema"}
}

type IndexIntervalBytesOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type LocalRetentionBytesOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type LocalRetentionMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type MaxCompactionLagMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type MaxMessageBytesOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type MessageDownconversionEnableOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *bool        `json:"value,omitempty"`
}
type MessageFormatVersionOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    string       `json:"value,omitempty"`
}
type MessageFormatVersionType string

const (
	MessageFormatVersionType080     MessageFormatVersionType = "0.8.0"
	MessageFormatVersionType081     MessageFormatVersionType = "0.8.1"
	MessageFormatVersionType082     MessageFormatVersionType = "0.8.2"
	MessageFormatVersionType090     MessageFormatVersionType = "0.9.0"
	MessageFormatVersionType0100    MessageFormatVersionType = "0.10.0"
	MessageFormatVersionType0100Iv0 MessageFormatVersionType = "0.10.0-IV0"
	MessageFormatVersionType0100Iv1 MessageFormatVersionType = "0.10.0-IV1"
	MessageFormatVersionType0101    MessageFormatVersionType = "0.10.1"
	MessageFormatVersionType0101Iv0 MessageFormatVersionType = "0.10.1-IV0"
	MessageFormatVersionType0101Iv1 MessageFormatVersionType = "0.10.1-IV1"
	MessageFormatVersionType0101Iv2 MessageFormatVersionType = "0.10.1-IV2"
	MessageFormatVersionType0102    MessageFormatVersionType = "0.10.2"
	MessageFormatVersionType0102Iv0 MessageFormatVersionType = "0.10.2-IV0"
	MessageFormatVersionType0110    MessageFormatVersionType = "0.11.0"
	MessageFormatVersionType0110Iv0 MessageFormatVersionType = "0.11.0-IV0"
	MessageFormatVersionType0110Iv1 MessageFormatVersionType = "0.11.0-IV1"
	MessageFormatVersionType0110Iv2 MessageFormatVersionType = "0.11.0-IV2"
	MessageFormatVersionType10      MessageFormatVersionType = "1.0"
	MessageFormatVersionType10Iv0   MessageFormatVersionType = "1.0-IV0"
	MessageFormatVersionType11      MessageFormatVersionType = "1.1"
	MessageFormatVersionType11Iv0   MessageFormatVersionType = "1.1-IV0"
	MessageFormatVersionType20      MessageFormatVersionType = "2.0"
	MessageFormatVersionType20Iv0   MessageFormatVersionType = "2.0-IV0"
	MessageFormatVersionType20Iv1   MessageFormatVersionType = "2.0-IV1"
	MessageFormatVersionType21      MessageFormatVersionType = "2.1"
	MessageFormatVersionType21Iv0   MessageFormatVersionType = "2.1-IV0"
	MessageFormatVersionType21Iv1   MessageFormatVersionType = "2.1-IV1"
	MessageFormatVersionType21Iv2   MessageFormatVersionType = "2.1-IV2"
	MessageFormatVersionType22      MessageFormatVersionType = "2.2"
	MessageFormatVersionType22Iv0   MessageFormatVersionType = "2.2-IV0"
	MessageFormatVersionType22Iv1   MessageFormatVersionType = "2.2-IV1"
	MessageFormatVersionType23      MessageFormatVersionType = "2.3"
	MessageFormatVersionType23Iv0   MessageFormatVersionType = "2.3-IV0"
	MessageFormatVersionType23Iv1   MessageFormatVersionType = "2.3-IV1"
	MessageFormatVersionType24      MessageFormatVersionType = "2.4"
	MessageFormatVersionType24Iv0   MessageFormatVersionType = "2.4-IV0"
	MessageFormatVersionType24Iv1   MessageFormatVersionType = "2.4-IV1"
	MessageFormatVersionType25      MessageFormatVersionType = "2.5"
	MessageFormatVersionType25Iv0   MessageFormatVersionType = "2.5-IV0"
	MessageFormatVersionType26      MessageFormatVersionType = "2.6"
	MessageFormatVersionType26Iv0   MessageFormatVersionType = "2.6-IV0"
	MessageFormatVersionType27      MessageFormatVersionType = "2.7"
	MessageFormatVersionType27Iv0   MessageFormatVersionType = "2.7-IV0"
	MessageFormatVersionType27Iv1   MessageFormatVersionType = "2.7-IV1"
	MessageFormatVersionType27Iv2   MessageFormatVersionType = "2.7-IV2"
	MessageFormatVersionType28      MessageFormatVersionType = "2.8"
	MessageFormatVersionType28Iv0   MessageFormatVersionType = "2.8-IV0"
	MessageFormatVersionType28Iv1   MessageFormatVersionType = "2.8-IV1"
	MessageFormatVersionType30      MessageFormatVersionType = "3.0"
	MessageFormatVersionType30Iv0   MessageFormatVersionType = "3.0-IV0"
	MessageFormatVersionType30Iv1   MessageFormatVersionType = "3.0-IV1"
	MessageFormatVersionType31      MessageFormatVersionType = "3.1"
	MessageFormatVersionType31Iv0   MessageFormatVersionType = "3.1-IV0"
	MessageFormatVersionType32      MessageFormatVersionType = "3.2"
	MessageFormatVersionType32Iv0   MessageFormatVersionType = "3.2-IV0"
	MessageFormatVersionType33      MessageFormatVersionType = "3.3"
	MessageFormatVersionType33Iv0   MessageFormatVersionType = "3.3-IV0"
	MessageFormatVersionType33Iv1   MessageFormatVersionType = "3.3-IV1"
	MessageFormatVersionType33Iv2   MessageFormatVersionType = "3.3-IV2"
	MessageFormatVersionType33Iv3   MessageFormatVersionType = "3.3-IV3"
	MessageFormatVersionType34      MessageFormatVersionType = "3.4"
	MessageFormatVersionType34Iv0   MessageFormatVersionType = "3.4-IV0"
	MessageFormatVersionType35      MessageFormatVersionType = "3.5"
	MessageFormatVersionType35Iv0   MessageFormatVersionType = "3.5-IV0"
	MessageFormatVersionType35Iv1   MessageFormatVersionType = "3.5-IV1"
	MessageFormatVersionType35Iv2   MessageFormatVersionType = "3.5-IV2"
	MessageFormatVersionType36      MessageFormatVersionType = "3.6"
	MessageFormatVersionType36Iv0   MessageFormatVersionType = "3.6-IV0"
	MessageFormatVersionType36Iv1   MessageFormatVersionType = "3.6-IV1"
	MessageFormatVersionType36Iv2   MessageFormatVersionType = "3.6-IV2"
)

func MessageFormatVersionTypeChoices() []string {
	return []string{"0.8.0", "0.8.1", "0.8.2", "0.9.0", "0.10.0", "0.10.0-IV0", "0.10.0-IV1", "0.10.1", "0.10.1-IV0", "0.10.1-IV1", "0.10.1-IV2", "0.10.2", "0.10.2-IV0", "0.11.0", "0.11.0-IV0", "0.11.0-IV1", "0.11.0-IV2", "1.0", "1.0-IV0", "1.1", "1.1-IV0", "2.0", "2.0-IV0", "2.0-IV1", "2.1", "2.1-IV0", "2.1-IV1", "2.1-IV2", "2.2", "2.2-IV0", "2.2-IV1", "2.3", "2.3-IV0", "2.3-IV1", "2.4", "2.4-IV0", "2.4-IV1", "2.5", "2.5-IV0", "2.6", "2.6-IV0", "2.7", "2.7-IV0", "2.7-IV1", "2.7-IV2", "2.8", "2.8-IV0", "2.8-IV1", "3.0", "3.0-IV0", "3.0-IV1", "3.1", "3.1-IV0", "3.2", "3.2-IV0", "3.3", "3.3-IV0", "3.3-IV1", "3.3-IV2", "3.3-IV3", "3.4", "3.4-IV0", "3.5", "3.5-IV0", "3.5-IV1", "3.5-IV2", "3.6", "3.6-IV0", "3.6-IV1", "3.6-IV2"}
}

type MessageOut struct {
	Key       map[string]any `json:"key,omitempty"`
	Offset    *int           `json:"offset,omitempty"`
	Partition *int           `json:"partition,omitempty"`
	Topic     string         `json:"topic,omitempty"`
	Value     map[string]any `json:"value,omitempty"`
}
type MessageTimestampDifferenceMaxMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type MessageTimestampType string

const (
	MessageTimestampTypeCreateTime    MessageTimestampType = "CreateTime"
	MessageTimestampTypeLogAppendTime MessageTimestampType = "LogAppendTime"
)

func MessageTimestampTypeChoices() []string {
	return []string{"CreateTime", "LogAppendTime"}
}

type MessageTimestampTypeOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    string       `json:"value,omitempty"`
}
type MinCleanableDirtyRatioOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *float64     `json:"value,omitempty"`
}
type MinCompactionLagMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type MinInsyncReplicasOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type OffsetOut struct {
	Error     string `json:"error,omitempty"`
	ErrorCode *int   `json:"error_code,omitempty"`
	Offset    *int   `json:"offset,omitempty"`
	Partition *int   `json:"partition,omitempty"`
}
type PartitionOut struct {
	ConsumerGroups []ConsumerGroupOut `json:"consumer_groups"`
	EarliestOffset int                `json:"earliest_offset"`
	Isr            int                `json:"isr"`
	LatestOffset   int                `json:"latest_offset"`
	Partition      int                `json:"partition"`
	Size           int                `json:"size"`
}
type PreallocateOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *bool        `json:"value,omitempty"`
}
type RecordIn struct {
	Key       *map[string]any `json:"key,omitempty"`
	Partition *int            `json:"partition,omitempty"`
	Value     *map[string]any `json:"value,omitempty"`
}
type RemoteStorageEnableOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *bool        `json:"value,omitempty"`
}
type RetentionBytesOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type RetentionMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type SegmentBytesOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type SegmentIndexBytesOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type SegmentJitterMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type SegmentMsOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *int         `json:"value,omitempty"`
}
type ServiceKafkaTopicCreateIn struct {
	CleanupPolicy     CleanupPolicyType `json:"cleanup_policy,omitempty"`
	Config            *ConfigIn         `json:"config,omitempty"`
	MinInsyncReplicas *int              `json:"min_insync_replicas,omitempty"`
	Partitions        *int              `json:"partitions,omitempty"`
	Replication       *int              `json:"replication,omitempty"`
	RetentionBytes    *int              `json:"retention_bytes,omitempty"`
	RetentionHours    *int              `json:"retention_hours,omitempty"`
	Tags              *[]TagIn          `json:"tags,omitempty"`
	TopicName         string            `json:"topic_name"`
}
type ServiceKafkaTopicGetOut struct {
	CleanupPolicy     string         `json:"cleanup_policy"`
	Config            ConfigOut      `json:"config"`
	MinInsyncReplicas int            `json:"min_insync_replicas"`
	Partitions        []PartitionOut `json:"partitions"`
	Replication       int            `json:"replication"`
	RetentionBytes    int            `json:"retention_bytes"`
	RetentionHours    int            `json:"retention_hours"`
	State             string         `json:"state"`
	Tags              []TagOut       `json:"tags"`
	TopicName         string         `json:"topic_name"`
}
type ServiceKafkaTopicMessageListIn struct {
	Format     FormatType     `json:"format,omitempty"`
	MaxBytes   *int           `json:"max_bytes,omitempty"`
	Partitions map[string]any `json:"partitions"`
	Timeout    *int           `json:"timeout,omitempty"`
}
type ServiceKafkaTopicMessageProduceIn struct {
	Format        FormatType `json:"format"`
	KeySchema     string     `json:"key_schema,omitempty"`
	KeySchemaId   *int       `json:"key_schema_id,omitempty"`
	Records       []RecordIn `json:"records"`
	ValueSchema   string     `json:"value_schema,omitempty"`
	ValueSchemaId *int       `json:"value_schema_id,omitempty"`
}
type ServiceKafkaTopicMessageProduceOut struct {
	KeySchemaId   *int        `json:"key_schema_id,omitempty"`
	Offsets       []OffsetOut `json:"offsets,omitempty"`
	ValueSchemaId *int        `json:"value_schema_id,omitempty"`
}
type ServiceKafkaTopicUpdateIn struct {
	Config            *ConfigIn `json:"config,omitempty"`
	MinInsyncReplicas *int      `json:"min_insync_replicas,omitempty"`
	Partitions        *int      `json:"partitions,omitempty"`
	Replication       *int      `json:"replication,omitempty"`
	RetentionBytes    *int      `json:"retention_bytes,omitempty"`
	RetentionHours    *int      `json:"retention_hours,omitempty"`
	Tags              *[]TagIn  `json:"tags,omitempty"`
}
type SynonymOut struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
	Value  *bool  `json:"value,omitempty"`
}
type TagIn struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type TagOut struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type TopicOut struct {
	CleanupPolicy       string   `json:"cleanup_policy"`
	MinInsyncReplicas   int      `json:"min_insync_replicas"`
	Partitions          int      `json:"partitions"`
	RemoteStorageEnable *bool    `json:"remote_storage_enable,omitempty"`
	Replication         int      `json:"replication"`
	RetentionBytes      int      `json:"retention_bytes"`
	RetentionHours      int      `json:"retention_hours"`
	State               string   `json:"state"`
	Tags                []TagOut `json:"tags"`
	TopicName           string   `json:"topic_name"`
}
type UncleanLeaderElectionEnableOut struct {
	Source   string       `json:"source,omitempty"`
	Synonyms []SynonymOut `json:"synonyms,omitempty"`
	Value    *bool        `json:"value,omitempty"`
}
type serviceKafkaTopicGetOut struct {
	Topic ServiceKafkaTopicGetOut `json:"topic"`
}
type serviceKafkaTopicListOut struct {
	Topics []TopicOut `json:"topics"`
}
type serviceKafkaTopicMessageListOut struct {
	Messages []MessageOut `json:"messages,omitempty"`
}
