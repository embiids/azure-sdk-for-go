//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

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

// VirtualNetworkLinksClient contains the methods for the VirtualNetworkLinks group.
// Don't use this type directly, use NewVirtualNetworkLinksClient() instead.
type VirtualNetworkLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworkLinksClient creates a new instance of VirtualNetworkLinksClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualNetworkLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkLinksClient, error) {
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
	client := &VirtualNetworkLinksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a virtual network link to a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// virtualNetworkLinkName - The name of the virtual network link.
// parameters - Parameters supplied to the CreateOrUpdate operation.
// options - VirtualNetworkLinksClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginCreateOrUpdate
// method.
func (client *VirtualNetworkLinksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginCreateOrUpdateOptions) (*armruntime.Poller[VirtualNetworkLinksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, dnsForwardingRulesetName, virtualNetworkLinkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualNetworkLinksClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualNetworkLinksClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a virtual network link to a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkLinksClient) createOrUpdate(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, virtualNetworkLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkLinksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a virtual network link to a DNS forwarding ruleset. WARNING: This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// virtualNetworkLinkName - The name of the virtual network link.
// options - VirtualNetworkLinksClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginDelete
// method.
func (client *VirtualNetworkLinksClient) BeginDelete(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientBeginDeleteOptions) (*armruntime.Poller[VirtualNetworkLinksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, dnsForwardingRulesetName, virtualNetworkLinkName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualNetworkLinksClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualNetworkLinksClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a virtual network link to a DNS forwarding ruleset. WARNING: This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkLinksClient) deleteOperation(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, virtualNetworkLinkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets properties of a virtual network link to a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// virtualNetworkLinkName - The name of the virtual network link.
// options - VirtualNetworkLinksClientGetOptions contains the optional parameters for the VirtualNetworkLinksClient.Get method.
func (client *VirtualNetworkLinksClient) Get(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientGetOptions) (VirtualNetworkLinksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, virtualNetworkLinkName, options)
	if err != nil {
		return VirtualNetworkLinksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkLinksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkLinksClient) getHandleResponse(resp *http.Response) (VirtualNetworkLinksClientGetResponse, error) {
	result := VirtualNetworkLinksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkLink); err != nil {
		return VirtualNetworkLinksClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists virtual network links to a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// options - VirtualNetworkLinksClientListOptions contains the optional parameters for the VirtualNetworkLinksClient.List
// method.
func (client *VirtualNetworkLinksClient) List(resourceGroupName string, dnsForwardingRulesetName string, options *VirtualNetworkLinksClientListOptions) *runtime.Pager[VirtualNetworkLinksClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[VirtualNetworkLinksClientListResponse]{
		More: func(page VirtualNetworkLinksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkLinksClientListResponse) (VirtualNetworkLinksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworkLinksClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualNetworkLinksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworkLinksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualNetworkLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, options *VirtualNetworkLinksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/virtualNetworkLinks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualNetworkLinksClient) listHandleResponse(resp *http.Response) (VirtualNetworkLinksClientListResponse, error) {
	result := VirtualNetworkLinksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkLinkListResult); err != nil {
		return VirtualNetworkLinksClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a virtual network link to a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// virtualNetworkLinkName - The name of the virtual network link.
// parameters - Parameters supplied to the Update operation.
// options - VirtualNetworkLinksClientBeginUpdateOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginUpdate
// method.
func (client *VirtualNetworkLinksClient) BeginUpdate(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, parameters VirtualNetworkLinkPatch, options *VirtualNetworkLinksClientBeginUpdateOptions) (*armruntime.Poller[VirtualNetworkLinksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, dnsForwardingRulesetName, virtualNetworkLinkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[VirtualNetworkLinksClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualNetworkLinksClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a virtual network link to a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkLinksClient) update(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, parameters VirtualNetworkLinkPatch, options *VirtualNetworkLinksClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, virtualNetworkLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualNetworkLinksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, virtualNetworkLinkName string, parameters VirtualNetworkLinkPatch, options *VirtualNetworkLinksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
