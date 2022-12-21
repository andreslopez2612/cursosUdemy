// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package prometheusserviceiface provides an interface to enable mocking the Amazon Prometheus Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package prometheusserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/prometheusservice"
)

// PrometheusServiceAPI provides an interface to enable mocking the
// prometheusservice.PrometheusService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Prometheus Service.
//	func myFunc(svc prometheusserviceiface.PrometheusServiceAPI) bool {
//	    // Make svc.CreateAlertManagerDefinition request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := prometheusservice.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockPrometheusServiceClient struct {
//	    prometheusserviceiface.PrometheusServiceAPI
//	}
//	func (m *mockPrometheusServiceClient) CreateAlertManagerDefinition(input *prometheusservice.CreateAlertManagerDefinitionInput) (*prometheusservice.CreateAlertManagerDefinitionOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockPrometheusServiceClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PrometheusServiceAPI interface {
	CreateAlertManagerDefinition(*prometheusservice.CreateAlertManagerDefinitionInput) (*prometheusservice.CreateAlertManagerDefinitionOutput, error)
	CreateAlertManagerDefinitionWithContext(aws.Context, *prometheusservice.CreateAlertManagerDefinitionInput, ...request.Option) (*prometheusservice.CreateAlertManagerDefinitionOutput, error)
	CreateAlertManagerDefinitionRequest(*prometheusservice.CreateAlertManagerDefinitionInput) (*request.Request, *prometheusservice.CreateAlertManagerDefinitionOutput)

	CreateLoggingConfiguration(*prometheusservice.CreateLoggingConfigurationInput) (*prometheusservice.CreateLoggingConfigurationOutput, error)
	CreateLoggingConfigurationWithContext(aws.Context, *prometheusservice.CreateLoggingConfigurationInput, ...request.Option) (*prometheusservice.CreateLoggingConfigurationOutput, error)
	CreateLoggingConfigurationRequest(*prometheusservice.CreateLoggingConfigurationInput) (*request.Request, *prometheusservice.CreateLoggingConfigurationOutput)

	CreateRuleGroupsNamespace(*prometheusservice.CreateRuleGroupsNamespaceInput) (*prometheusservice.CreateRuleGroupsNamespaceOutput, error)
	CreateRuleGroupsNamespaceWithContext(aws.Context, *prometheusservice.CreateRuleGroupsNamespaceInput, ...request.Option) (*prometheusservice.CreateRuleGroupsNamespaceOutput, error)
	CreateRuleGroupsNamespaceRequest(*prometheusservice.CreateRuleGroupsNamespaceInput) (*request.Request, *prometheusservice.CreateRuleGroupsNamespaceOutput)

	CreateWorkspace(*prometheusservice.CreateWorkspaceInput) (*prometheusservice.CreateWorkspaceOutput, error)
	CreateWorkspaceWithContext(aws.Context, *prometheusservice.CreateWorkspaceInput, ...request.Option) (*prometheusservice.CreateWorkspaceOutput, error)
	CreateWorkspaceRequest(*prometheusservice.CreateWorkspaceInput) (*request.Request, *prometheusservice.CreateWorkspaceOutput)

	DeleteAlertManagerDefinition(*prometheusservice.DeleteAlertManagerDefinitionInput) (*prometheusservice.DeleteAlertManagerDefinitionOutput, error)
	DeleteAlertManagerDefinitionWithContext(aws.Context, *prometheusservice.DeleteAlertManagerDefinitionInput, ...request.Option) (*prometheusservice.DeleteAlertManagerDefinitionOutput, error)
	DeleteAlertManagerDefinitionRequest(*prometheusservice.DeleteAlertManagerDefinitionInput) (*request.Request, *prometheusservice.DeleteAlertManagerDefinitionOutput)

	DeleteLoggingConfiguration(*prometheusservice.DeleteLoggingConfigurationInput) (*prometheusservice.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationWithContext(aws.Context, *prometheusservice.DeleteLoggingConfigurationInput, ...request.Option) (*prometheusservice.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationRequest(*prometheusservice.DeleteLoggingConfigurationInput) (*request.Request, *prometheusservice.DeleteLoggingConfigurationOutput)

	DeleteRuleGroupsNamespace(*prometheusservice.DeleteRuleGroupsNamespaceInput) (*prometheusservice.DeleteRuleGroupsNamespaceOutput, error)
	DeleteRuleGroupsNamespaceWithContext(aws.Context, *prometheusservice.DeleteRuleGroupsNamespaceInput, ...request.Option) (*prometheusservice.DeleteRuleGroupsNamespaceOutput, error)
	DeleteRuleGroupsNamespaceRequest(*prometheusservice.DeleteRuleGroupsNamespaceInput) (*request.Request, *prometheusservice.DeleteRuleGroupsNamespaceOutput)

	DeleteWorkspace(*prometheusservice.DeleteWorkspaceInput) (*prometheusservice.DeleteWorkspaceOutput, error)
	DeleteWorkspaceWithContext(aws.Context, *prometheusservice.DeleteWorkspaceInput, ...request.Option) (*prometheusservice.DeleteWorkspaceOutput, error)
	DeleteWorkspaceRequest(*prometheusservice.DeleteWorkspaceInput) (*request.Request, *prometheusservice.DeleteWorkspaceOutput)

	DescribeAlertManagerDefinition(*prometheusservice.DescribeAlertManagerDefinitionInput) (*prometheusservice.DescribeAlertManagerDefinitionOutput, error)
	DescribeAlertManagerDefinitionWithContext(aws.Context, *prometheusservice.DescribeAlertManagerDefinitionInput, ...request.Option) (*prometheusservice.DescribeAlertManagerDefinitionOutput, error)
	DescribeAlertManagerDefinitionRequest(*prometheusservice.DescribeAlertManagerDefinitionInput) (*request.Request, *prometheusservice.DescribeAlertManagerDefinitionOutput)

	DescribeLoggingConfiguration(*prometheusservice.DescribeLoggingConfigurationInput) (*prometheusservice.DescribeLoggingConfigurationOutput, error)
	DescribeLoggingConfigurationWithContext(aws.Context, *prometheusservice.DescribeLoggingConfigurationInput, ...request.Option) (*prometheusservice.DescribeLoggingConfigurationOutput, error)
	DescribeLoggingConfigurationRequest(*prometheusservice.DescribeLoggingConfigurationInput) (*request.Request, *prometheusservice.DescribeLoggingConfigurationOutput)

	DescribeRuleGroupsNamespace(*prometheusservice.DescribeRuleGroupsNamespaceInput) (*prometheusservice.DescribeRuleGroupsNamespaceOutput, error)
	DescribeRuleGroupsNamespaceWithContext(aws.Context, *prometheusservice.DescribeRuleGroupsNamespaceInput, ...request.Option) (*prometheusservice.DescribeRuleGroupsNamespaceOutput, error)
	DescribeRuleGroupsNamespaceRequest(*prometheusservice.DescribeRuleGroupsNamespaceInput) (*request.Request, *prometheusservice.DescribeRuleGroupsNamespaceOutput)

	DescribeWorkspace(*prometheusservice.DescribeWorkspaceInput) (*prometheusservice.DescribeWorkspaceOutput, error)
	DescribeWorkspaceWithContext(aws.Context, *prometheusservice.DescribeWorkspaceInput, ...request.Option) (*prometheusservice.DescribeWorkspaceOutput, error)
	DescribeWorkspaceRequest(*prometheusservice.DescribeWorkspaceInput) (*request.Request, *prometheusservice.DescribeWorkspaceOutput)

	ListRuleGroupsNamespaces(*prometheusservice.ListRuleGroupsNamespacesInput) (*prometheusservice.ListRuleGroupsNamespacesOutput, error)
	ListRuleGroupsNamespacesWithContext(aws.Context, *prometheusservice.ListRuleGroupsNamespacesInput, ...request.Option) (*prometheusservice.ListRuleGroupsNamespacesOutput, error)
	ListRuleGroupsNamespacesRequest(*prometheusservice.ListRuleGroupsNamespacesInput) (*request.Request, *prometheusservice.ListRuleGroupsNamespacesOutput)

	ListRuleGroupsNamespacesPages(*prometheusservice.ListRuleGroupsNamespacesInput, func(*prometheusservice.ListRuleGroupsNamespacesOutput, bool) bool) error
	ListRuleGroupsNamespacesPagesWithContext(aws.Context, *prometheusservice.ListRuleGroupsNamespacesInput, func(*prometheusservice.ListRuleGroupsNamespacesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*prometheusservice.ListTagsForResourceInput) (*prometheusservice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *prometheusservice.ListTagsForResourceInput, ...request.Option) (*prometheusservice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*prometheusservice.ListTagsForResourceInput) (*request.Request, *prometheusservice.ListTagsForResourceOutput)

	ListWorkspaces(*prometheusservice.ListWorkspacesInput) (*prometheusservice.ListWorkspacesOutput, error)
	ListWorkspacesWithContext(aws.Context, *prometheusservice.ListWorkspacesInput, ...request.Option) (*prometheusservice.ListWorkspacesOutput, error)
	ListWorkspacesRequest(*prometheusservice.ListWorkspacesInput) (*request.Request, *prometheusservice.ListWorkspacesOutput)

	ListWorkspacesPages(*prometheusservice.ListWorkspacesInput, func(*prometheusservice.ListWorkspacesOutput, bool) bool) error
	ListWorkspacesPagesWithContext(aws.Context, *prometheusservice.ListWorkspacesInput, func(*prometheusservice.ListWorkspacesOutput, bool) bool, ...request.Option) error

	PutAlertManagerDefinition(*prometheusservice.PutAlertManagerDefinitionInput) (*prometheusservice.PutAlertManagerDefinitionOutput, error)
	PutAlertManagerDefinitionWithContext(aws.Context, *prometheusservice.PutAlertManagerDefinitionInput, ...request.Option) (*prometheusservice.PutAlertManagerDefinitionOutput, error)
	PutAlertManagerDefinitionRequest(*prometheusservice.PutAlertManagerDefinitionInput) (*request.Request, *prometheusservice.PutAlertManagerDefinitionOutput)

	PutRuleGroupsNamespace(*prometheusservice.PutRuleGroupsNamespaceInput) (*prometheusservice.PutRuleGroupsNamespaceOutput, error)
	PutRuleGroupsNamespaceWithContext(aws.Context, *prometheusservice.PutRuleGroupsNamespaceInput, ...request.Option) (*prometheusservice.PutRuleGroupsNamespaceOutput, error)
	PutRuleGroupsNamespaceRequest(*prometheusservice.PutRuleGroupsNamespaceInput) (*request.Request, *prometheusservice.PutRuleGroupsNamespaceOutput)

	TagResource(*prometheusservice.TagResourceInput) (*prometheusservice.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *prometheusservice.TagResourceInput, ...request.Option) (*prometheusservice.TagResourceOutput, error)
	TagResourceRequest(*prometheusservice.TagResourceInput) (*request.Request, *prometheusservice.TagResourceOutput)

	UntagResource(*prometheusservice.UntagResourceInput) (*prometheusservice.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *prometheusservice.UntagResourceInput, ...request.Option) (*prometheusservice.UntagResourceOutput, error)
	UntagResourceRequest(*prometheusservice.UntagResourceInput) (*request.Request, *prometheusservice.UntagResourceOutput)

	UpdateLoggingConfiguration(*prometheusservice.UpdateLoggingConfigurationInput) (*prometheusservice.UpdateLoggingConfigurationOutput, error)
	UpdateLoggingConfigurationWithContext(aws.Context, *prometheusservice.UpdateLoggingConfigurationInput, ...request.Option) (*prometheusservice.UpdateLoggingConfigurationOutput, error)
	UpdateLoggingConfigurationRequest(*prometheusservice.UpdateLoggingConfigurationInput) (*request.Request, *prometheusservice.UpdateLoggingConfigurationOutput)

	UpdateWorkspaceAlias(*prometheusservice.UpdateWorkspaceAliasInput) (*prometheusservice.UpdateWorkspaceAliasOutput, error)
	UpdateWorkspaceAliasWithContext(aws.Context, *prometheusservice.UpdateWorkspaceAliasInput, ...request.Option) (*prometheusservice.UpdateWorkspaceAliasOutput, error)
	UpdateWorkspaceAliasRequest(*prometheusservice.UpdateWorkspaceAliasInput) (*request.Request, *prometheusservice.UpdateWorkspaceAliasOutput)

	WaitUntilWorkspaceActive(*prometheusservice.DescribeWorkspaceInput) error
	WaitUntilWorkspaceActiveWithContext(aws.Context, *prometheusservice.DescribeWorkspaceInput, ...request.WaiterOption) error

	WaitUntilWorkspaceDeleted(*prometheusservice.DescribeWorkspaceInput) error
	WaitUntilWorkspaceDeletedWithContext(aws.Context, *prometheusservice.DescribeWorkspaceInput, ...request.WaiterOption) error
}

var _ PrometheusServiceAPI = (*prometheusservice.PrometheusService)(nil)
