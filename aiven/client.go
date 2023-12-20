// Code generated by Aiven. DO NOT EDIT.

package aiven

import (
	"context"

	account "github.com/aiven/aiven-go-client-v2/handler/account"
	billing "github.com/aiven/aiven-go-client-v2/handler/billing"
	billinggroup "github.com/aiven/aiven-go-client-v2/handler/billinggroup"
	clickhouse "github.com/aiven/aiven-go-client-v2/handler/clickhouse"
	cloudplatform "github.com/aiven/aiven-go-client-v2/handler/cloudplatform"
	domain "github.com/aiven/aiven-go-client-v2/handler/domain"
	flink "github.com/aiven/aiven-go-client-v2/handler/flink"
	flinkapplication "github.com/aiven/aiven-go-client-v2/handler/flinkapplication"
	flinkapplicationdeployment "github.com/aiven/aiven-go-client-v2/handler/flinkapplicationdeployment"
	flinkapplicationversion "github.com/aiven/aiven-go-client-v2/handler/flinkapplicationversion"
	flinkjob "github.com/aiven/aiven-go-client-v2/handler/flinkjob"
	kafka "github.com/aiven/aiven-go-client-v2/handler/kafka"
	kafkaconnect "github.com/aiven/aiven-go-client-v2/handler/kafkaconnect"
	kafkamirrormaker "github.com/aiven/aiven-go-client-v2/handler/kafkamirrormaker"
	kafkaschemaregistry "github.com/aiven/aiven-go-client-v2/handler/kafkaschemaregistry"
	kafkatopic "github.com/aiven/aiven-go-client-v2/handler/kafkatopic"
	mysql "github.com/aiven/aiven-go-client-v2/handler/mysql"
	opensearch "github.com/aiven/aiven-go-client-v2/handler/opensearch"
	organization "github.com/aiven/aiven-go-client-v2/handler/organization"
	organizationuser "github.com/aiven/aiven-go-client-v2/handler/organizationuser"
	postgresql "github.com/aiven/aiven-go-client-v2/handler/postgresql"
	privatelink "github.com/aiven/aiven-go-client-v2/handler/privatelink"
	project "github.com/aiven/aiven-go-client-v2/handler/project"
	projectbilling "github.com/aiven/aiven-go-client-v2/handler/projectbilling"
	projectkeymanagement "github.com/aiven/aiven-go-client-v2/handler/projectkeymanagement"
	service "github.com/aiven/aiven-go-client-v2/handler/service"
	serviceintegration "github.com/aiven/aiven-go-client-v2/handler/serviceintegration"
	serviceintegrationendpoint "github.com/aiven/aiven-go-client-v2/handler/serviceintegrationendpoint"
	staticip "github.com/aiven/aiven-go-client-v2/handler/staticip"
	user "github.com/aiven/aiven-go-client-v2/handler/user"
	usergroup "github.com/aiven/aiven-go-client-v2/handler/usergroup"
)

type doer interface {
	Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error)
}

func newClient(doer doer) *Client {
	return &Client{
		Account:                    account.NewHandler(doer),
		Billing:                    billing.NewHandler(doer),
		BillingGroup:               billinggroup.NewHandler(doer),
		ClickHouse:                 clickhouse.NewHandler(doer),
		CloudPlatform:              cloudplatform.NewHandler(doer),
		Domain:                     domain.NewHandler(doer),
		Flink:                      flink.NewHandler(doer),
		FlinkApplication:           flinkapplication.NewHandler(doer),
		FlinkApplicationDeployment: flinkapplicationdeployment.NewHandler(doer),
		FlinkApplicationVersion:    flinkapplicationversion.NewHandler(doer),
		FlinkJob:                   flinkjob.NewHandler(doer),
		Kafka:                      kafka.NewHandler(doer),
		KafkaConnect:               kafkaconnect.NewHandler(doer),
		KafkaMirrorMaker:           kafkamirrormaker.NewHandler(doer),
		KafkaSchemaRegistry:        kafkaschemaregistry.NewHandler(doer),
		KafkaTopic:                 kafkatopic.NewHandler(doer),
		MySql:                      mysql.NewHandler(doer),
		OpenSearch:                 opensearch.NewHandler(doer),
		Organization:               organization.NewHandler(doer),
		OrganizationUser:           organizationuser.NewHandler(doer),
		PostgreSql:                 postgresql.NewHandler(doer),
		Privatelink:                privatelink.NewHandler(doer),
		Project:                    project.NewHandler(doer),
		ProjectBilling:             projectbilling.NewHandler(doer),
		ProjectKeyManagement:       projectkeymanagement.NewHandler(doer),
		Service:                    service.NewHandler(doer),
		ServiceIntegration:         serviceintegration.NewHandler(doer),
		ServiceIntegrationEndpoint: serviceintegrationendpoint.NewHandler(doer),
		StaticIp:                   staticip.NewHandler(doer),
		User:                       user.NewHandler(doer),
		UserGroup:                  usergroup.NewHandler(doer),
	}
}

type Client struct {
	Account                    account.Handler
	Billing                    billing.Handler
	BillingGroup               billinggroup.Handler
	ClickHouse                 clickhouse.Handler
	CloudPlatform              cloudplatform.Handler
	Domain                     domain.Handler
	Flink                      flink.Handler
	FlinkApplication           flinkapplication.Handler
	FlinkApplicationDeployment flinkapplicationdeployment.Handler
	FlinkApplicationVersion    flinkapplicationversion.Handler
	FlinkJob                   flinkjob.Handler
	Kafka                      kafka.Handler
	KafkaConnect               kafkaconnect.Handler
	KafkaMirrorMaker           kafkamirrormaker.Handler
	KafkaSchemaRegistry        kafkaschemaregistry.Handler
	KafkaTopic                 kafkatopic.Handler
	MySql                      mysql.Handler
	OpenSearch                 opensearch.Handler
	Organization               organization.Handler
	OrganizationUser           organizationuser.Handler
	PostgreSql                 postgresql.Handler
	Privatelink                privatelink.Handler
	Project                    project.Handler
	ProjectBilling             projectbilling.Handler
	ProjectKeyManagement       projectkeymanagement.Handler
	Service                    service.Handler
	ServiceIntegration         serviceintegration.Handler
	ServiceIntegrationEndpoint serviceintegrationendpoint.Handler
	StaticIp                   staticip.Handler
	User                       user.Handler
	UserGroup                  usergroup.Handler
}
