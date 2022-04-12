//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// RolloutsClient contains the methods for the Rollouts group.
// Don't use this type directly, use NewRolloutsClient() instead.
type RolloutsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRolloutsClient creates a new instance of RolloutsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRolloutsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RolloutsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RolloutsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Cancel - Only running rollouts can be canceled.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// rolloutName - The rollout name.
// options - RolloutsClientCancelOptions contains the optional parameters for the RolloutsClient.Cancel method.
func (client *RolloutsClient) Cancel(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientCancelOptions) (RolloutsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, rolloutName, options)
	if err != nil {
		return RolloutsClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RolloutsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RolloutsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *RolloutsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelHandleResponse handles the Cancel response.
func (client *RolloutsClient) cancelHandleResponse(resp *http.Response) (RolloutsClientCancelResponse, error) {
	result := RolloutsClientCancelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Rollout); err != nil {
		return RolloutsClientCancelResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - This is an asynchronous operation and can be polled to completion using the location header returned
// by this operation.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// rolloutName - The rollout name.
// options - RolloutsClientBeginCreateOrUpdateOptions contains the optional parameters for the RolloutsClient.BeginCreateOrUpdate
// method.
func (client *RolloutsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[RolloutsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, rolloutName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[RolloutsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[RolloutsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - This is an asynchronous operation and can be polled to completion using the location header returned by
// this operation.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RolloutsClient) createOrUpdate(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, rolloutName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RolloutsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.RolloutRequest != nil {
		return req, runtime.MarshalAsJSON(req, *options.RolloutRequest)
	}
	return req, nil
}

// Delete - Only rollouts in terminal state can be deleted.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// rolloutName - The rollout name.
// options - RolloutsClientDeleteOptions contains the optional parameters for the RolloutsClient.Delete method.
func (client *RolloutsClient) Delete(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientDeleteOptions) (RolloutsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, rolloutName, options)
	if err != nil {
		return RolloutsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RolloutsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RolloutsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RolloutsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RolloutsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets detailed information of a rollout.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// rolloutName - The rollout name.
// options - RolloutsClientGetOptions contains the optional parameters for the RolloutsClient.Get method.
func (client *RolloutsClient) Get(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientGetOptions) (RolloutsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, rolloutName, options)
	if err != nil {
		return RolloutsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RolloutsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RolloutsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RolloutsClient) getCreateRequest(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	if options != nil && options.RetryAttempt != nil {
		reqQP.Set("retryAttempt", strconv.FormatInt(int64(*options.RetryAttempt), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RolloutsClient) getHandleResponse(resp *http.Response) (RolloutsClientGetResponse, error) {
	result := RolloutsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Rollout); err != nil {
		return RolloutsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the rollouts in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - RolloutsClientListOptions contains the optional parameters for the RolloutsClient.List method.
func (client *RolloutsClient) List(ctx context.Context, resourceGroupName string, options *RolloutsClientListOptions) (RolloutsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return RolloutsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RolloutsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RolloutsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RolloutsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *RolloutsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RolloutsClient) listHandleResponse(resp *http.Response) (RolloutsClientListResponse, error) {
	result := RolloutsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RolloutArray); err != nil {
		return RolloutsClientListResponse{}, err
	}
	return result, nil
}

// Restart - Only failed rollouts can be restarted.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// rolloutName - The rollout name.
// options - RolloutsClientRestartOptions contains the optional parameters for the RolloutsClient.Restart method.
func (client *RolloutsClient) Restart(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientRestartOptions) (RolloutsClientRestartResponse, error) {
	req, err := client.restartCreateRequest(ctx, resourceGroupName, rolloutName, options)
	if err != nil {
		return RolloutsClientRestartResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RolloutsClientRestartResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RolloutsClientRestartResponse{}, runtime.NewResponseError(resp)
	}
	return client.restartHandleResponse(resp)
}

// restartCreateRequest creates the Restart request.
func (client *RolloutsClient) restartCreateRequest(ctx context.Context, resourceGroupName string, rolloutName string, options *RolloutsClientRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipSucceeded != nil {
		reqQP.Set("skipSucceeded", strconv.FormatBool(*options.SkipSucceeded))
	}
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// restartHandleResponse handles the Restart response.
func (client *RolloutsClient) restartHandleResponse(resp *http.Response) (RolloutsClientRestartResponse, error) {
	result := RolloutsClientRestartResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Rollout); err != nil {
		return RolloutsClientRestartResponse{}, err
	}
	return result, nil
}
