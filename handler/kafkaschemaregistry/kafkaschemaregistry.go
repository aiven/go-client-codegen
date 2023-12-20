// Code generated by Aiven. DO NOT EDIT.

package kafkaschemaregistry

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	// AclAdd add a Schema Registry ACL entry
	// ServiceSchemaRegistryAclAdd POST /project/{project}/service/{service_name}/kafka/schema-registry/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryAclAdd
	AclAdd(ctx context.Context, project string, serviceName string, in *AclAddIn) ([]Acl, error)

	// AclDelete delete a Schema Registry ACL entry
	// ServiceSchemaRegistryAclDelete DELETE /project/{project}/service/{service_name}/kafka/schema-registry/acl/{schema_registry_acl_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryAclDelete
	AclDelete(ctx context.Context, project string, serviceName string, schemaRegistryAclId string) ([]Acl, error)

	// AclList list Schema Registry ACL entries
	// ServiceSchemaRegistryAclList GET /project/{project}/service/{service_name}/kafka/schema-registry/acl
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryAclList
	AclList(ctx context.Context, project string, serviceName string) ([]Acl, error)

	// Compatibility check compatibility of schema in Schema Registry
	// ServiceSchemaRegistryCompatibility POST /project/{project}/service/{service_name}/kafka/schema/compatibility/subjects/{subject_name}/versions/{version_id:latest|\d+}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryCompatibility
	Compatibility(ctx context.Context, project string, serviceName string, subjectName string, versionId int, in *CompatibilityIn) (bool, error)

	// GlobalConfigGet get global configuration for Schema Registry
	// ServiceSchemaRegistryGlobalConfigGet GET /project/{project}/service/{service_name}/kafka/schema/config
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryGlobalConfigGet
	GlobalConfigGet(ctx context.Context, project string, serviceName string) (CompatibilityLevelType, error)

	// GlobalConfigPut edit global configuration for Schema Registry
	// ServiceSchemaRegistryGlobalConfigPut PUT /project/{project}/service/{service_name}/kafka/schema/config
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistryGlobalConfigPut
	GlobalConfigPut(ctx context.Context, project string, serviceName string, in *GlobalConfigPutIn) (CompatibilityType, error)

	// SchemaGet get schema in Schema Registry
	// ServiceSchemaRegistrySchemaGet GET /project/{project}/service/{service_name}/kafka/schema/schemas/ids/{schema_id}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySchemaGet
	SchemaGet(ctx context.Context, project string, serviceName string, schemaId string) error

	// SubjectConfigGet get configuration for Schema Registry subject
	// ServiceSchemaRegistrySubjectConfigGet GET /project/{project}/service/{service_name}/kafka/schema/config/{subject_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectConfigGet
	SubjectConfigGet(ctx context.Context, project string, serviceName string, subjectName string) (CompatibilityLevelType, error)

	// SubjectConfigPut edit configuration for Schema Registry subject
	// ServiceSchemaRegistrySubjectConfigPut PUT /project/{project}/service/{service_name}/kafka/schema/config/{subject_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectConfigPut
	SubjectConfigPut(ctx context.Context, project string, serviceName string, subjectName string, in *SubjectConfigPutIn) (CompatibilityType, error)

	// SubjectDelete delete Schema Registry subject
	// ServiceSchemaRegistrySubjectDelete DELETE /project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectDelete
	SubjectDelete(ctx context.Context, project string, serviceName string, subjectName string) error

	// SubjectVersionDelete delete Schema Registry subject version
	// ServiceSchemaRegistrySubjectVersionDelete DELETE /project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions/{version_id:latest|\d+}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionDelete
	SubjectVersionDelete(ctx context.Context, project string, serviceName string, subjectName string, versionId int) error

	// SubjectVersionGet get Schema Registry Subject version
	// ServiceSchemaRegistrySubjectVersionGet GET /project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions/{version_id:latest|\d+}
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionGet
	SubjectVersionGet(ctx context.Context, project string, serviceName string, subjectName string, versionId int) error

	// SubjectVersionPost register a new Schema in Schema Registry
	// ServiceSchemaRegistrySubjectVersionPost POST /project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionPost
	SubjectVersionPost(ctx context.Context, project string, serviceName string, subjectName string, in *SubjectVersionPostIn) (int, error)

	// SubjectVersionsGet get Schema Registry subject versions
	// ServiceSchemaRegistrySubjectVersionsGet GET /project/{project}/service/{service_name}/kafka/schema/subjects/{subject_name}/versions
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjectVersionsGet
	SubjectVersionsGet(ctx context.Context, project string, serviceName string, subjectName string) ([]int, error)

	// Subjects lists Schema Registry subjects
	// ServiceSchemaRegistrySubjects GET /project/{project}/service/{service_name}/kafka/schema/subjects
	// https://api.aiven.io/doc/#tag/Service:_Kafka/operation/ServiceSchemaRegistrySubjects
	Subjects(ctx context.Context, project string, serviceName string) ([]string, error)
}

func NewHandler(doer doer) Handler {
	return &handler{doer}
}

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

type handler struct {
	doer doer
}

func (h *handler) AclAdd(ctx context.Context, project string, serviceName string, in *AclAddIn) ([]Acl, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema-registry/acl", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryAclAdd", "POST", path, in)
	out := new(aclAddOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *handler) AclDelete(ctx context.Context, project string, serviceName string, schemaRegistryAclId string) ([]Acl, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema-registry/acl/%s", project, serviceName, schemaRegistryAclId)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryAclDelete", "DELETE", path, nil)
	out := new(aclDeleteOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *handler) AclList(ctx context.Context, project string, serviceName string) ([]Acl, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema-registry/acl", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryAclList", "GET", path, nil)
	out := new(aclListOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Acl, nil
}
func (h *handler) Compatibility(ctx context.Context, project string, serviceName string, subjectName string, versionId int, in *CompatibilityIn) (bool, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/compatibility/subjects/%s/versions/%d", project, serviceName, subjectName, versionId)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryCompatibility", "POST", path, in)
	out := new(compatibilityOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return false, err
	}
	return out.IsCompatible, nil
}
func (h *handler) GlobalConfigGet(ctx context.Context, project string, serviceName string) (CompatibilityLevelType, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/config", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryGlobalConfigGet", "GET", path, nil)
	out := new(globalConfigGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.CompatibilityLevel, nil
}
func (h *handler) GlobalConfigPut(ctx context.Context, project string, serviceName string, in *GlobalConfigPutIn) (CompatibilityType, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/config", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistryGlobalConfigPut", "PUT", path, in)
	out := new(globalConfigPutOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.Compatibility, nil
}
func (h *handler) SchemaGet(ctx context.Context, project string, serviceName string, schemaId string) error {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/schemas/ids/%s", project, serviceName, schemaId)
	_, err := h.doer.Do(ctx, "ServiceSchemaRegistrySchemaGet", "GET", path, nil)
	return err
}
func (h *handler) SubjectConfigGet(ctx context.Context, project string, serviceName string, subjectName string) (CompatibilityLevelType, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/config/%s", project, serviceName, subjectName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectConfigGet", "GET", path, nil)
	out := new(subjectConfigGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.CompatibilityLevel, nil
}
func (h *handler) SubjectConfigPut(ctx context.Context, project string, serviceName string, subjectName string, in *SubjectConfigPutIn) (CompatibilityType, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/config/%s", project, serviceName, subjectName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectConfigPut", "PUT", path, in)
	out := new(subjectConfigPutOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return "", err
	}
	return out.Compatibility, nil
}
func (h *handler) SubjectDelete(ctx context.Context, project string, serviceName string, subjectName string) error {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/subjects/%s", project, serviceName, subjectName)
	_, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectDelete", "DELETE", path, nil)
	return err
}
func (h *handler) SubjectVersionDelete(ctx context.Context, project string, serviceName string, subjectName string, versionId int) error {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/subjects/%s/versions/%d", project, serviceName, subjectName, versionId)
	_, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionDelete", "DELETE", path, nil)
	return err
}
func (h *handler) SubjectVersionGet(ctx context.Context, project string, serviceName string, subjectName string, versionId int) error {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/subjects/%s/versions/%d", project, serviceName, subjectName, versionId)
	_, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionGet", "GET", path, nil)
	return err
}
func (h *handler) SubjectVersionPost(ctx context.Context, project string, serviceName string, subjectName string, in *SubjectVersionPostIn) (int, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/subjects/%s/versions", project, serviceName, subjectName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionPost", "POST", path, in)
	out := new(subjectVersionPostOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return 0, err
	}
	return out.Id, nil
}
func (h *handler) SubjectVersionsGet(ctx context.Context, project string, serviceName string, subjectName string) ([]int, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/subjects/%s/versions", project, serviceName, subjectName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjectVersionsGet", "GET", path, nil)
	out := new(subjectVersionsGetOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Versions, nil
}
func (h *handler) Subjects(ctx context.Context, project string, serviceName string) ([]string, error) {
	path := fmt.Sprintf("/project/%s/service/%s/kafka/schema/subjects", project, serviceName)
	b, err := h.doer.Do(ctx, "ServiceSchemaRegistrySubjects", "GET", path, nil)
	out := new(subjectsOut)
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out.Subjects, nil
}

type Acl struct {
	Id         string         `json:"id,omitempty"`
	Permission PermissionType `json:"permission"`
	Resource   string         `json:"resource"`
	Username   string         `json:"username"`
}
type AclAddIn struct {
	Permission PermissionType `json:"permission"`
	Resource   string         `json:"resource"`
	Username   string         `json:"username"`
}
type aclAddOut struct {
	Acl []Acl `json:"acl"`
}
type aclDeleteOut struct {
	Acl []Acl `json:"acl"`
}
type aclListOut struct {
	Acl []Acl `json:"acl"`
}
type CompatibilityIn struct {
	Schema     string     `json:"schema"`
	SchemaType SchemaType `json:"schemaType,omitempty"`
}
type CompatibilityLevelType string

const (
	CompatibilityLevelTypeBackward           CompatibilityLevelType = "BACKWARD"
	CompatibilityLevelTypeBackwardTransitive CompatibilityLevelType = "BACKWARD_TRANSITIVE"
	CompatibilityLevelTypeForward            CompatibilityLevelType = "FORWARD"
	CompatibilityLevelTypeForwardTransitive  CompatibilityLevelType = "FORWARD_TRANSITIVE"
	CompatibilityLevelTypeFull               CompatibilityLevelType = "FULL"
	CompatibilityLevelTypeFullTransitive     CompatibilityLevelType = "FULL_TRANSITIVE"
	CompatibilityLevelTypeNone               CompatibilityLevelType = "NONE"
)

type compatibilityOut struct {
	IsCompatible bool `json:"is_compatible"`
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

type globalConfigGetOut struct {
	CompatibilityLevel CompatibilityLevelType `json:"compatibilityLevel"`
}
type GlobalConfigPutIn struct {
	Compatibility CompatibilityType `json:"compatibility"`
}
type globalConfigPutOut struct {
	Compatibility CompatibilityType `json:"compatibility"`
}
type PermissionType string

const (
	PermissionTypeSchemaRegistryRead  PermissionType = "schema_registry_read"
	PermissionTypeSchemaRegistryWrite PermissionType = "schema_registry_write"
)

type Reference struct {
	Name    string `json:"name"`
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

type subjectConfigGetOut struct {
	CompatibilityLevel CompatibilityLevelType `json:"compatibilityLevel"`
}
type SubjectConfigPutIn struct {
	Compatibility CompatibilityType `json:"compatibility"`
}
type subjectConfigPutOut struct {
	Compatibility CompatibilityType `json:"compatibility"`
}
type SubjectVersionPostIn struct {
	References []Reference `json:"references"`
	Schema     string      `json:"schema"`
	SchemaType SchemaType  `json:"schemaType,omitempty"`
}
type subjectVersionPostOut struct {
	Id int `json:"id"`
}
type subjectVersionsGetOut struct {
	Versions []int `json:"versions"`
}
type subjectsOut struct {
	Subjects []string `json:"subjects"`
}
