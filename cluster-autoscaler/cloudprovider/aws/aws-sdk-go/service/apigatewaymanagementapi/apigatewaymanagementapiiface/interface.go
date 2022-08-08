// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package apigatewaymanagementapiiface provides an interface to enable mocking the AmazonApiGatewayManagementApi service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package apigatewaymanagementapiiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

// ApiGatewayManagementApiAPI provides an interface to enable mocking the
// apigatewaymanagementapi.ApiGatewayManagementApi service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AmazonApiGatewayManagementApi.
//    func myFunc(svc apigatewaymanagementapiiface.ApiGatewayManagementApiAPI) bool {
//        // Make svc.DeleteConnection request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := apigatewaymanagementapi.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockApiGatewayManagementApiClient struct {
//        apigatewaymanagementapiiface.ApiGatewayManagementApiAPI
//    }
//    func (m *mockApiGatewayManagementApiClient) DeleteConnection(input *apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockApiGatewayManagementApiClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ApiGatewayManagementApiAPI interface {
	DeleteConnection(*apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error)
	DeleteConnectionWithContext(aws.Context, *apigatewaymanagementapi.DeleteConnectionInput, ...request.Option) (*apigatewaymanagementapi.DeleteConnectionOutput, error)
	DeleteConnectionRequest(*apigatewaymanagementapi.DeleteConnectionInput) (*request.Request, *apigatewaymanagementapi.DeleteConnectionOutput)

	GetConnection(*apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error)
	GetConnectionWithContext(aws.Context, *apigatewaymanagementapi.GetConnectionInput, ...request.Option) (*apigatewaymanagementapi.GetConnectionOutput, error)
	GetConnectionRequest(*apigatewaymanagementapi.GetConnectionInput) (*request.Request, *apigatewaymanagementapi.GetConnectionOutput)

	PostToConnection(*apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error)
	PostToConnectionWithContext(aws.Context, *apigatewaymanagementapi.PostToConnectionInput, ...request.Option) (*apigatewaymanagementapi.PostToConnectionOutput, error)
	PostToConnectionRequest(*apigatewaymanagementapi.PostToConnectionInput) (*request.Request, *apigatewaymanagementapi.PostToConnectionOutput)
}

var _ ApiGatewayManagementApiAPI = (*apigatewaymanagementapi.ApiGatewayManagementApi)(nil)