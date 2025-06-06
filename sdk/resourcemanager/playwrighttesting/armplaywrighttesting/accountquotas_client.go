//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armplaywrighttesting

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

// AccountQuotasClient contains the methods for the AccountQuotas group.
// Don't use this type directly, use NewAccountQuotasClient() instead.
type AccountQuotasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccountQuotasClient creates a new instance of AccountQuotasClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccountQuotasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccountQuotasClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccountQuotasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get quota by name for an account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Name of account.
//   - quotaName - The Playwright service account quota name.
//   - options - AccountQuotasClientGetOptions contains the optional parameters for the AccountQuotasClient.Get method.
func (client *AccountQuotasClient) Get(ctx context.Context, resourceGroupName string, accountName string, quotaName QuotaNames, options *AccountQuotasClientGetOptions) (AccountQuotasClientGetResponse, error) {
	var err error
	const operationName = "AccountQuotasClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, quotaName, options)
	if err != nil {
		return AccountQuotasClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountQuotasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccountQuotasClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AccountQuotasClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, quotaName QuotaNames, options *AccountQuotasClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzurePlaywrightService/accounts/{accountName}/quotas/{quotaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if quotaName == "" {
		return nil, errors.New("parameter quotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaName}", url.PathEscape(string(quotaName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountQuotasClient) getHandleResponse(resp *http.Response) (AccountQuotasClientGetResponse, error) {
	result := AccountQuotasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountQuota); err != nil {
		return AccountQuotasClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - List quotas for a given account.
//
// Generated from API version 2024-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Name of account.
//   - options - AccountQuotasClientListByAccountOptions contains the optional parameters for the AccountQuotasClient.NewListByAccountPager
//     method.
func (client *AccountQuotasClient) NewListByAccountPager(resourceGroupName string, accountName string, options *AccountQuotasClientListByAccountOptions) *runtime.Pager[AccountQuotasClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccountQuotasClientListByAccountResponse]{
		More: func(page AccountQuotasClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountQuotasClientListByAccountResponse) (AccountQuotasClientListByAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccountQuotasClient.NewListByAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return AccountQuotasClientListByAccountResponse{}, err
			}
			return client.listByAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *AccountQuotasClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountQuotasClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzurePlaywrightService/accounts/{accountName}/quotas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *AccountQuotasClient) listByAccountHandleResponse(resp *http.Response) (AccountQuotasClientListByAccountResponse, error) {
	result := AccountQuotasClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountQuotaListResult); err != nil {
		return AccountQuotasClientListByAccountResponse{}, err
	}
	return result, nil
}
