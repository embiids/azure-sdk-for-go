// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armservicenetworking

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AssociationsInterfaceClient contains the methods for the AssociationsInterface group.
// Don't use this type directly, use NewAssociationsInterfaceClient() instead.
type AssociationsInterfaceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssociationsInterfaceClient creates a new instance of AssociationsInterfaceClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssociationsInterfaceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssociationsInterfaceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssociationsInterfaceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - associationName - Name of Association
//   - resource - Resource create parameters.
//   - options - AssociationsInterfaceClientBeginCreateOrUpdateOptions contains the optional parameters for the AssociationsInterfaceClient.BeginCreateOrUpdate
//     method.
func (client *AssociationsInterfaceClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, resource Association, options *AssociationsInterfaceClientBeginCreateOrUpdateOptions) (*runtime.Poller[AssociationsInterfaceClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, trafficControllerName, associationName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AssociationsInterfaceClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AssociationsInterfaceClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *AssociationsInterfaceClient) createOrUpdate(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, resource Association, options *AssociationsInterfaceClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AssociationsInterfaceClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, trafficControllerName, associationName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssociationsInterfaceClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, resource Association, _ *AssociationsInterfaceClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - associationName - Name of Association
//   - options - AssociationsInterfaceClientBeginDeleteOptions contains the optional parameters for the AssociationsInterfaceClient.BeginDelete
//     method.
func (client *AssociationsInterfaceClient) BeginDelete(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *AssociationsInterfaceClientBeginDeleteOptions) (*runtime.Poller[AssociationsInterfaceClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, trafficControllerName, associationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AssociationsInterfaceClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AssociationsInterfaceClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *AssociationsInterfaceClient) deleteOperation(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *AssociationsInterfaceClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AssociationsInterfaceClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, trafficControllerName, associationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AssociationsInterfaceClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, _ *AssociationsInterfaceClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - associationName - Name of Association
//   - options - AssociationsInterfaceClientGetOptions contains the optional parameters for the AssociationsInterfaceClient.Get
//     method.
func (client *AssociationsInterfaceClient) Get(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, options *AssociationsInterfaceClientGetOptions) (AssociationsInterfaceClientGetResponse, error) {
	var err error
	const operationName = "AssociationsInterfaceClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, trafficControllerName, associationName, options)
	if err != nil {
		return AssociationsInterfaceClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssociationsInterfaceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssociationsInterfaceClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssociationsInterfaceClient) getCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, _ *AssociationsInterfaceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssociationsInterfaceClient) getHandleResponse(resp *http.Response) (AssociationsInterfaceClientGetResponse, error) {
	result := AssociationsInterfaceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Association); err != nil {
		return AssociationsInterfaceClientGetResponse{}, err
	}
	return result, nil
}

// NewListByTrafficControllerPager - List Association resources by TrafficController
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - options - AssociationsInterfaceClientListByTrafficControllerOptions contains the optional parameters for the AssociationsInterfaceClient.NewListByTrafficControllerPager
//     method.
func (client *AssociationsInterfaceClient) NewListByTrafficControllerPager(resourceGroupName string, trafficControllerName string, options *AssociationsInterfaceClientListByTrafficControllerOptions) *runtime.Pager[AssociationsInterfaceClientListByTrafficControllerResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssociationsInterfaceClientListByTrafficControllerResponse]{
		More: func(page AssociationsInterfaceClientListByTrafficControllerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssociationsInterfaceClientListByTrafficControllerResponse) (AssociationsInterfaceClientListByTrafficControllerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssociationsInterfaceClient.NewListByTrafficControllerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByTrafficControllerCreateRequest(ctx, resourceGroupName, trafficControllerName, options)
			}, nil)
			if err != nil {
				return AssociationsInterfaceClientListByTrafficControllerResponse{}, err
			}
			return client.listByTrafficControllerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTrafficControllerCreateRequest creates the ListByTrafficController request.
func (client *AssociationsInterfaceClient) listByTrafficControllerCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, _ *AssociationsInterfaceClientListByTrafficControllerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTrafficControllerHandleResponse handles the ListByTrafficController response.
func (client *AssociationsInterfaceClient) listByTrafficControllerHandleResponse(resp *http.Response) (AssociationsInterfaceClientListByTrafficControllerResponse, error) {
	result := AssociationsInterfaceClientListByTrafficControllerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssociationListResult); err != nil {
		return AssociationsInterfaceClientListByTrafficControllerResponse{}, err
	}
	return result, nil
}

// Update - Update a Association
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - associationName - Name of Association
//   - properties - The resource properties to be updated.
//   - options - AssociationsInterfaceClientUpdateOptions contains the optional parameters for the AssociationsInterfaceClient.Update
//     method.
func (client *AssociationsInterfaceClient) Update(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, properties AssociationUpdate, options *AssociationsInterfaceClientUpdateOptions) (AssociationsInterfaceClientUpdateResponse, error) {
	var err error
	const operationName = "AssociationsInterfaceClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, trafficControllerName, associationName, properties, options)
	if err != nil {
		return AssociationsInterfaceClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssociationsInterfaceClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssociationsInterfaceClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AssociationsInterfaceClient) updateCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, associationName string, properties AssociationUpdate, _ *AssociationsInterfaceClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/associations/{associationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AssociationsInterfaceClient) updateHandleResponse(resp *http.Response) (AssociationsInterfaceClientUpdateResponse, error) {
	result := AssociationsInterfaceClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Association); err != nil {
		return AssociationsInterfaceClientUpdateResponse{}, err
	}
	return result, nil
}
