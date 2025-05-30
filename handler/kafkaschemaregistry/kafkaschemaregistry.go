// Code generated by Aiven. DO NOT EDIT.

package kafkaschemaregistry

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Handler interface {
	// ServiceSchemaRegistryAclAdd add a Schema Registry ACL entry
	// POST /v1/project/{project}/service/{service_name}/kafka/schema-registry/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryAclAdd
	ServiceSchemaRegistryAclAdd(ctx context.Context, project string, serviceName string, in *ServiceSchemaRegistryAclAddIn) ([]AclOut, error)

	// ServiceSchemaRegistryAclDelete delete a Schema Registry ACL entry
	// DELETE /v1/project/{project}/service/{service_name}/kafka/schema-registry/acl/{schema_registry_acl_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryAclDelete
	ServiceSchemaRegistryAclDelete(ctx context.Context, project string, serviceName string, schemaRegistryAclId string) ([]AclOut, error)

	// ServiceSchemaRegistryAclList list Schema Registry ACL entries
	// GET /v1/project/{project}/service/{service_name}/kafka/schema-registry/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryAclList
	ServiceSchemaRegistryAclList(ctx context.Context, project string, serviceName string) ([]AclOut, error)

	// ServiceSchemaRegistryCompatibility check compatibility of schema in Schema Registry
	// POST /v1/project/{project}/service/{service_name}/kafka/schema/compatibility/subjects/{subject_name}/versions/{version_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryCompatibility
	ServiceSchemaRegistryCompatibility(ctx context.Context, project string, serviceName string, subjectName string, versionId int, in *ServiceSchemaRegistryCompatibilityIn) (*ServiceSchemaRegistryCompatibilityOut, error)

	// ServiceSchemaRegistryGlobalConfigGet get global configuration for Schema Registry
	// GET /v1/project/{project}/service/{service_name}/kafka/schema/config
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryGlobalConfigGet
	ServiceSchemaRegistryGlobalConfigGet(ctx context.Context, project string, serviceName string) (CompatibilityType, error)

	// ServiceSchemaRegistryGlobalConfigPut edit global configuration for Schema Registry
	// PUT /v1/project/{project}/service/{service_name}/kafka/schema/config
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryGlobalConfigPut
	ServiceSchemaRegistryGlobalConfigPut(ctx context.Context, project string, serviceName string, in *ServiceSchemaRegistryGlobalConfigPutIn) (CompatibilityType, error)

	// ServiceSchemaRegistrySchemaGet get schema in Schema Registry
	// GET /v1/project/{project}/service/{service_name}/kafka/schema/schemas/ids/{schema_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySchemaGet
	ServiceSchemaRegistrySchemaGet(ctx context.Context, project string, serviceName string, schemaId string) error

	// ServiceSchemaRegistrySubjectConfigGet get configuration for Schema Registry subject
	// GET /v1/project/{project}/service/{service_name}/kafka/schema/config/{subject_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectConfigGet
	ServiceSchemaRegistrySubjectConfigGet(ctx context.Context, project string, serviceName string, subjectName string) (CompatibilityType, error)

	// ServiceSchemaRegistrySubjectConfigPut edit configuration for Schema Registry subject
	// PUT /v1/project/{project}/service/{service_name}/kafka/schema/config/{subject_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectConfigPut
	ServiceSchemaRegistrySubjectConfigPut(ctx context.Context, project string, serviceName string, subjectName string, in *ServiceSchemaRegistrySubjectConfigPutIn) (CompatibilityType, error)

	// ServiceSchemaRegistrySubjectDelete delete Schema Registry subject
	// DELETE /v1/project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectDelete
	ServiceSchemaRegistrySubjectDelete(ctx context.Context, project string, serviceName string, subjectName string) error

	// ServiceSchemaRegistrySubjectVersionDelete delete Schema Registry subject version
	// DELETE /v1/project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions/{version_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionDelete
	ServiceSchemaRegistrySubjectVersionDelete(ctx context.Context, project string, serviceName string, subjectName string, versionId int) error

	// ServiceSchemaRegistrySubjectVersionGet get Schema Registry Subject version
	// GET /v1/project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions/{version_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionGet
	ServiceSchemaRegistrySubjectVersionGet(ctx context.Context, project string, serviceName string, subjectName string, versionId int) (*ServiceSchemaRegistrySubjectVersionGetOut, error)

	// ServiceSchemaRegistrySubjectVersionPost register a new Schema in Schema Registry
	// POST /v1/project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionPost
	ServiceSchemaRegistrySubjectVersionPost(ctx context.Context, project string, serviceName string, subjectName string, in *ServiceSchemaRegistrySubjectVersionPostIn) (int, error)

	// Deprecated: ServiceSchemaRegistrySubjectVersionSchemaGet dEPRECATED: Get raw schema of a specific version in Schema Registry
	// GET /v1/project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions/{version_id}/schema
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionSchemaGet
	ServiceSchemaRegistrySubjectVersionSchemaGet(ctx context.Context, project string, serviceName string, subjectName string, versionId int) error

	// ServiceSchemaRegistrySubjectVersionsGet get Schema Registry subject versions
	// GET /v1/project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionsGet
	ServiceSchemaRegistrySubjectVersionsGet(ctx context.Context, project string, serviceName string, subjectName string) ([]int, error)

	// ServiceSchemaRegistrySubjects lists Schema Registry subjects
	// GET /v1/project/{project}/service/{service_name}/kafka/schema/subjects
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjects
	ServiceSchemaRegistrySubjects(ctx context.Context, project string, serviceName string) ([]string, error)
}

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func NewHandler(doer doer) KafkaSchemaRegistryHandler {
	return KafkaSchemaRegistryHandler{doer}
}

type KafkaSchemaRegistryHandler struct {
	doer doer
}

func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistryAclAdd(ctx context.Context, project string, serviceName string, in *ServiceSchemaRegistryAclAddIn) ([]AclOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema-registry/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryAclAdd", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(serviceSchemaRegistryAclAddOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistryAclDelete(ctx context.Context, project string, serviceName string, schemaRegistryAclId string) ([]AclOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema-registry/acl/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(schemaRegistryAclId))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryAclDelete", "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceSchemaRegistryAclDeleteOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistryAclList(ctx context.Context, project string, serviceName string) ([]AclOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema-registry/acl", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryAclList", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceSchemaRegistryAclListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistryCompatibility(ctx context.Context, project string, serviceName string, subjectName string, versionId int, in *ServiceSchemaRegistryCompatibilityIn) (*ServiceSchemaRegistryCompatibilityOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/compatibility/subjects/%s/versions/%d", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName), versionId)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryCompatibility", "POST", path, in)
	if err != nil {
		return nil, err
	}
	out := new(ServiceSchemaRegistryCompatibilityOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistryGlobalConfigGet(ctx context.Context, project string, serviceName string) (CompatibilityType, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/config", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryGlobalConfigGet", "GET", path, nil)
	if err != nil {
		return "", err
	}
	out := new(serviceSchemaRegistryGlobalConfigGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.CompatibilityLevel, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistryGlobalConfigPut(ctx context.Context, project string, serviceName string, in *ServiceSchemaRegistryGlobalConfigPutIn) (CompatibilityType, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/config", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryGlobalConfigPut", "PUT", path, in)
	if err != nil {
		return "", err
	}
	out := new(serviceSchemaRegistryGlobalConfigPutOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.Compatibility, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySchemaGet(ctx context.Context, project string, serviceName string, schemaId string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/schemas/ids/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(schemaId))
	_, err := h.doer.Do(ctx, "ServiceSchemaRegistrySchemaGet", "GET", path, nil)
	return err
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjectConfigGet(ctx context.Context, project string, serviceName string, subjectName string) (CompatibilityType, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/config/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectConfigGet", "GET", path, nil)
	if err != nil {
		return "", err
	}
	out := new(serviceSchemaRegistrySubjectConfigGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.CompatibilityLevel, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjectConfigPut(ctx context.Context, project string, serviceName string, subjectName string, in *ServiceSchemaRegistrySubjectConfigPutIn) (CompatibilityType, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/config/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectConfigPut", "PUT", path, in)
	if err != nil {
		return "", err
	}
	out := new(serviceSchemaRegistrySubjectConfigPutOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.Compatibility, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjectDelete(ctx context.Context, project string, serviceName string, subjectName string) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/subjects/%s", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName))
	_, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectDelete", "DELETE", path, nil)
	return err
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjectVersionDelete(ctx context.Context, project string, serviceName string, subjectName string, versionId int) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/subjects/%s/versions/%d", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName), versionId)
	_, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionDelete", "DELETE", path, nil)
	return err
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjectVersionGet(ctx context.Context, project string, serviceName string, subjectName string, versionId int) (*ServiceSchemaRegistrySubjectVersionGetOut, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/subjects/%s/versions/%d", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName), versionId)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceSchemaRegistrySubjectVersionGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return &out.Version, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjectVersionPost(ctx context.Context, project string, serviceName string, subjectName string, in *ServiceSchemaRegistrySubjectVersionPostIn) (int, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/subjects/%s/versions", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionPost", "POST", path, in)
	if err != nil {
		return 0, err
	}
	out := new(serviceSchemaRegistrySubjectVersionPostOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return 0, err
	}
	return out.Id, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjectVersionSchemaGet(ctx context.Context, project string, serviceName string, subjectName string, versionId int) error {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/subjects/%s/versions/%d/schema", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName), versionId)
	_, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionSchemaGet", "GET", path, nil)
	return err
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjectVersionsGet(ctx context.Context, project string, serviceName string, subjectName string) ([]int, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/subjects/%s/versions", url.PathEscape(project), url.PathEscape(serviceName), url.PathEscape(subjectName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionsGet", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceSchemaRegistrySubjectVersionsGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Versions, nil
}
func (h *KafkaSchemaRegistryHandler) ServiceSchemaRegistrySubjects(ctx context.Context, project string, serviceName string) ([]string, error) {
	path := fmt.Sprintf("/v1/project/%s/service/%s/kafka/schema/subjects", url.PathEscape(project), url.PathEscape(serviceName))
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjects", "GET", path, nil)
	if err != nil {
		return nil, err
	}
	out := new(serviceSchemaRegistrySubjectsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Subjects, nil
}

type AclOut struct {
	Id         *string        `json:"id,omitempty"` // ID
	Permission PermissionType `json:"permission"`   // ACL entry for Schema Registry
	Resource   string         `json:"resource"`     // Schema Registry ACL entry resource name pattern
	Username   string         `json:"username"`
}
type CompatibilityType string

const (
	CompatibilityTypeBackward           CompatibilityType = "BACKWARD"
	CompatibilityTypeBackwardTransitive CompatibilityType = "BACKWARD_TRANSITIVE"
	CompatibilityTypeForward            CompatibilityType = "FORWARD"
	CompatibilityTypeForwardTransitive  CompatibilityType = "FORWARD_TRANSITIVE"
	CompatibilityTypeFull               CompatibilityType = "FULL"
	CompatibilityTypeFullTransitive     CompatibilityType = "FULL_TRANSITIVE"
	CompatibilityTypeNone               CompatibilityType = "NONE"
)

func CompatibilityTypeChoices() []string {
	return []string{"BACKWARD", "BACKWARD_TRANSITIVE", "FORWARD", "FORWARD_TRANSITIVE", "FULL", "FULL_TRANSITIVE", "NONE"}
}

type PermissionType string

const (
	PermissionTypeSchemaRegistryRead  PermissionType = "schema_registry_read"
	PermissionTypeSchemaRegistryWrite PermissionType = "schema_registry_write"
)

func PermissionTypeChoices() []string {
	return []string{"schema_registry_read", "schema_registry_write"}
}

type ReferenceIn struct {
	Name    string `json:"name"` // The name used to reference the provided subject and version
	Subject string `json:"subject"`
	Version int    `json:"version"`
}
type ReferenceOut struct {
	Name    string `json:"name"` // The name used to reference the provided subject and version
	Subject string `json:"subject"`
	Version int    `json:"version"`
}
type SchemaType string

const (
	SchemaTypeAvro     SchemaType = "AVRO"
	SchemaTypeJson     SchemaType = "JSON"
	SchemaTypeProtobuf SchemaType = "PROTOBUF"
)

func SchemaTypeChoices() []string {
	return []string{"AVRO", "JSON", "PROTOBUF"}
}

// ServiceSchemaRegistryAclAddIn ServiceSchemaRegistryAclAddRequestBody
type ServiceSchemaRegistryAclAddIn struct {
	Permission PermissionType `json:"permission"` // ACL entry for Schema Registry
	Resource   string         `json:"resource"`   // Schema Registry ACL entry resource name pattern
	Username   string         `json:"username"`
}

// ServiceSchemaRegistryCompatibilityIn ServiceSchemaRegistryCompatibilityRequestBody
type ServiceSchemaRegistryCompatibilityIn struct {
	Schema     string     `json:"schema"`
	SchemaType SchemaType `json:"schemaType,omitempty"` // Schema type
}

// ServiceSchemaRegistryCompatibilityOut ServiceSchemaRegistryCompatibilityResponse
type ServiceSchemaRegistryCompatibilityOut struct {
	IsCompatible bool     `json:"is_compatible"`      // Compatibility
	Messages     []string `json:"messages,omitempty"` // Compatibility check messages
}

// ServiceSchemaRegistryGlobalConfigPutIn ServiceSchemaRegistryGlobalConfigPutRequestBody
type ServiceSchemaRegistryGlobalConfigPutIn struct {
	Compatibility CompatibilityType `json:"compatibility"` // Compatibility level
}

// ServiceSchemaRegistrySubjectConfigPutIn ServiceSchemaRegistrySubjectConfigPutRequestBody
type ServiceSchemaRegistrySubjectConfigPutIn struct {
	Compatibility CompatibilityType `json:"compatibility"` // Compatibility level
}

// ServiceSchemaRegistrySubjectVersionGetOut Version
type ServiceSchemaRegistrySubjectVersionGetOut struct {
	Id         int            `json:"id"`                   // Schema Id
	References []ReferenceOut `json:"references,omitempty"` // Schema references
	Schema     string         `json:"schema"`
	SchemaType SchemaType     `json:"schemaType,omitempty"` // Schema type
	Subject    string         `json:"subject"`
	Version    int            `json:"version"`
}

// ServiceSchemaRegistrySubjectVersionPostIn ServiceSchemaRegistrySubjectVersionPostRequestBody
type ServiceSchemaRegistrySubjectVersionPostIn struct {
	References *[]ReferenceIn `json:"references,omitempty"` // Schema references
	Schema     string         `json:"schema"`
	SchemaType SchemaType     `json:"schemaType,omitempty"` // Schema type
}

// serviceSchemaRegistryAclAddOut ServiceSchemaRegistryAclAddResponse
type serviceSchemaRegistryAclAddOut struct {
	Acl []AclOut `json:"acl"` // List of Schema Registry ACL entries
}

// serviceSchemaRegistryAclDeleteOut ServiceSchemaRegistryAclDeleteResponse
type serviceSchemaRegistryAclDeleteOut struct {
	Acl []AclOut `json:"acl"` // List of Schema Registry ACL entries
}

// serviceSchemaRegistryAclListOut ServiceSchemaRegistryAclListResponse
type serviceSchemaRegistryAclListOut struct {
	Acl []AclOut `json:"acl"` // List of Schema Registry ACL entries
}

// serviceSchemaRegistryGlobalConfigGetOut ServiceSchemaRegistryGlobalConfigGetResponse
type serviceSchemaRegistryGlobalConfigGetOut struct {
	CompatibilityLevel CompatibilityType `json:"compatibilityLevel"` // Compatibility level
}

// serviceSchemaRegistryGlobalConfigPutOut ServiceSchemaRegistryGlobalConfigPutResponse
type serviceSchemaRegistryGlobalConfigPutOut struct {
	Compatibility CompatibilityType `json:"compatibility"` // Compatibility level
}

// serviceSchemaRegistrySubjectConfigGetOut ServiceSchemaRegistrySubjectConfigGetResponse
type serviceSchemaRegistrySubjectConfigGetOut struct {
	CompatibilityLevel CompatibilityType `json:"compatibilityLevel"` // Compatibility level
}

// serviceSchemaRegistrySubjectConfigPutOut ServiceSchemaRegistrySubjectConfigPutResponse
type serviceSchemaRegistrySubjectConfigPutOut struct {
	Compatibility CompatibilityType `json:"compatibility"` // Compatibility level
}

// serviceSchemaRegistrySubjectVersionGetOut ServiceSchemaRegistrySubjectVersionGetResponse
type serviceSchemaRegistrySubjectVersionGetOut struct {
	Version ServiceSchemaRegistrySubjectVersionGetOut `json:"version"` // Version
}

// serviceSchemaRegistrySubjectVersionPostOut ServiceSchemaRegistrySubjectVersionPostResponse
type serviceSchemaRegistrySubjectVersionPostOut struct {
	Id int `json:"id"` // Schema Id
}

// serviceSchemaRegistrySubjectVersionsGetOut ServiceSchemaRegistrySubjectVersionsGetResponse
type serviceSchemaRegistrySubjectVersionsGetOut struct {
	Versions []int `json:"versions"` // List of available versions for a Schema Registry subject
}

// serviceSchemaRegistrySubjectsOut ServiceSchemaRegistrySubjectsResponse
type serviceSchemaRegistrySubjectsOut struct {
	Subjects []string `json:"subjects"` // List of available Schema Registry subjects
}
