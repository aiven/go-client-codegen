// Code generated by Aiven. DO NOT EDIT.

package aiven

import (
	"context"

	account "github.com/aiven/go-client-codegen/handler/account"
	accountauthentication "github.com/aiven/go-client-codegen/handler/accountauthentication"
	accountteam "github.com/aiven/go-client-codegen/handler/accountteam"
	accountteammember "github.com/aiven/go-client-codegen/handler/accountteammember"
	alloydbomni "github.com/aiven/go-client-codegen/handler/alloydbomni"
	applicationuser "github.com/aiven/go-client-codegen/handler/applicationuser"
	billinggroup "github.com/aiven/go-client-codegen/handler/billinggroup"
	clickhouse "github.com/aiven/go-client-codegen/handler/clickhouse"
	cloud "github.com/aiven/go-client-codegen/handler/cloud"
	domain "github.com/aiven/go-client-codegen/handler/domain"
	flink "github.com/aiven/go-client-codegen/handler/flink"
	flinkapplication "github.com/aiven/go-client-codegen/handler/flinkapplication"
	flinkapplicationdeployment "github.com/aiven/go-client-codegen/handler/flinkapplicationdeployment"
	flinkapplicationversion "github.com/aiven/go-client-codegen/handler/flinkapplicationversion"
	flinkjarapplication "github.com/aiven/go-client-codegen/handler/flinkjarapplication"
	flinkjarapplicationdeployment "github.com/aiven/go-client-codegen/handler/flinkjarapplicationdeployment"
	flinkjarapplicationversion "github.com/aiven/go-client-codegen/handler/flinkjarapplicationversion"
	flinkjob "github.com/aiven/go-client-codegen/handler/flinkjob"
	kafka "github.com/aiven/go-client-codegen/handler/kafka"
	kafkaconnect "github.com/aiven/go-client-codegen/handler/kafkaconnect"
	kafkamirrormaker "github.com/aiven/go-client-codegen/handler/kafkamirrormaker"
	kafkaschemaregistry "github.com/aiven/go-client-codegen/handler/kafkaschemaregistry"
	kafkatopic "github.com/aiven/go-client-codegen/handler/kafkatopic"
	mysql "github.com/aiven/go-client-codegen/handler/mysql"
	opensearch "github.com/aiven/go-client-codegen/handler/opensearch"
	organization "github.com/aiven/go-client-codegen/handler/organization"
	organizationuser "github.com/aiven/go-client-codegen/handler/organizationuser"
	postgresql "github.com/aiven/go-client-codegen/handler/postgresql"
	privatelink "github.com/aiven/go-client-codegen/handler/privatelink"
	project "github.com/aiven/go-client-codegen/handler/project"
	projectbilling "github.com/aiven/go-client-codegen/handler/projectbilling"
	service "github.com/aiven/go-client-codegen/handler/service"
	staticip "github.com/aiven/go-client-codegen/handler/staticip"
	thanos "github.com/aiven/go-client-codegen/handler/thanos"
	user "github.com/aiven/go-client-codegen/handler/user"
	usergroup "github.com/aiven/go-client-codegen/handler/usergroup"
	vpc "github.com/aiven/go-client-codegen/handler/vpc"
)

// doer http client
type doer interface {
	Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error)
}

func newClient(doer doer) Client {
	return &client{
		AccountAuthenticationHandler:         accountauthentication.NewHandler(doer),
		AccountHandler:                       account.NewHandler(doer),
		AccountTeamHandler:                   accountteam.NewHandler(doer),
		AccountTeamMemberHandler:             accountteammember.NewHandler(doer),
		AlloyDBOmniHandler:                   alloydbomni.NewHandler(doer),
		ApplicationUserHandler:               applicationuser.NewHandler(doer),
		BillingGroupHandler:                  billinggroup.NewHandler(doer),
		ClickHouseHandler:                    clickhouse.NewHandler(doer),
		CloudHandler:                         cloud.NewHandler(doer),
		DomainHandler:                        domain.NewHandler(doer),
		FlinkApplicationDeploymentHandler:    flinkapplicationdeployment.NewHandler(doer),
		FlinkApplicationHandler:              flinkapplication.NewHandler(doer),
		FlinkApplicationVersionHandler:       flinkapplicationversion.NewHandler(doer),
		FlinkHandler:                         flink.NewHandler(doer),
		FlinkJarApplicationDeploymentHandler: flinkjarapplicationdeployment.NewHandler(doer),
		FlinkJarApplicationHandler:           flinkjarapplication.NewHandler(doer),
		FlinkJarApplicationVersionHandler:    flinkjarapplicationversion.NewHandler(doer),
		FlinkJobHandler:                      flinkjob.NewHandler(doer),
		KafkaConnectHandler:                  kafkaconnect.NewHandler(doer),
		KafkaHandler:                         kafka.NewHandler(doer),
		KafkaMirrorMakerHandler:              kafkamirrormaker.NewHandler(doer),
		KafkaSchemaRegistryHandler:           kafkaschemaregistry.NewHandler(doer),
		KafkaTopicHandler:                    kafkatopic.NewHandler(doer),
		MySQLHandler:                         mysql.NewHandler(doer),
		OpenSearchHandler:                    opensearch.NewHandler(doer),
		OrganizationHandler:                  organization.NewHandler(doer),
		OrganizationUserHandler:              organizationuser.NewHandler(doer),
		PostgreSQLHandler:                    postgresql.NewHandler(doer),
		PrivatelinkHandler:                   privatelink.NewHandler(doer),
		ProjectBillingHandler:                projectbilling.NewHandler(doer),
		ProjectHandler:                       project.NewHandler(doer),
		ServiceHandler:                       service.NewHandler(doer),
		StaticIPHandler:                      staticip.NewHandler(doer),
		ThanosHandler:                        thanos.NewHandler(doer),
		UserGroupHandler:                     usergroup.NewHandler(doer),
		UserHandler:                          user.NewHandler(doer),
		VpcHandler:                           vpc.NewHandler(doer),
	}
}

type client struct {
	account.AccountHandler
	accountauthentication.AccountAuthenticationHandler
	accountteam.AccountTeamHandler
	accountteammember.AccountTeamMemberHandler
	alloydbomni.AlloyDBOmniHandler
	applicationuser.ApplicationUserHandler
	billinggroup.BillingGroupHandler
	clickhouse.ClickHouseHandler
	cloud.CloudHandler
	domain.DomainHandler
	flink.FlinkHandler
	flinkapplication.FlinkApplicationHandler
	flinkapplicationdeployment.FlinkApplicationDeploymentHandler
	flinkapplicationversion.FlinkApplicationVersionHandler
	flinkjarapplication.FlinkJarApplicationHandler
	flinkjarapplicationdeployment.FlinkJarApplicationDeploymentHandler
	flinkjarapplicationversion.FlinkJarApplicationVersionHandler
	flinkjob.FlinkJobHandler
	kafka.KafkaHandler
	kafkaconnect.KafkaConnectHandler
	kafkamirrormaker.KafkaMirrorMakerHandler
	kafkaschemaregistry.KafkaSchemaRegistryHandler
	kafkatopic.KafkaTopicHandler
	mysql.MySQLHandler
	opensearch.OpenSearchHandler
	organization.OrganizationHandler
	organizationuser.OrganizationUserHandler
	postgresql.PostgreSQLHandler
	privatelink.PrivatelinkHandler
	project.ProjectHandler
	projectbilling.ProjectBillingHandler
	service.ServiceHandler
	staticip.StaticIPHandler
	thanos.ThanosHandler
	user.UserHandler
	usergroup.UserGroupHandler
	vpc.VpcHandler
}
type Client interface {
	account.Handler
	accountauthentication.Handler
	accountteam.Handler
	accountteammember.Handler
	alloydbomni.Handler
	applicationuser.Handler
	billinggroup.Handler
	clickhouse.Handler
	cloud.Handler
	domain.Handler
	flink.Handler
	flinkapplication.Handler
	flinkapplicationdeployment.Handler
	flinkapplicationversion.Handler
	flinkjarapplication.Handler
	flinkjarapplicationdeployment.Handler
	flinkjarapplicationversion.Handler
	flinkjob.Handler
	kafka.Handler
	kafkaconnect.Handler
	kafkamirrormaker.Handler
	kafkaschemaregistry.Handler
	kafkatopic.Handler
	mysql.Handler
	opensearch.Handler
	organization.Handler
	organizationuser.Handler
	postgresql.Handler
	privatelink.Handler
	project.Handler
	projectbilling.Handler
	service.Handler
	staticip.Handler
	thanos.Handler
	user.Handler
	usergroup.Handler
	vpc.Handler
}
