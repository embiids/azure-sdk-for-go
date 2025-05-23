// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationassessment

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

// WebAppAssessmentV2SummaryOperationsClient contains the methods for the WebAppAssessmentV2SummaryOperations group.
// Don't use this type directly, use NewWebAppAssessmentV2SummaryOperationsClient() instead.
type WebAppAssessmentV2SummaryOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebAppAssessmentV2SummaryOperationsClient creates a new instance of WebAppAssessmentV2SummaryOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebAppAssessmentV2SummaryOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebAppAssessmentV2SummaryOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebAppAssessmentV2SummaryOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a WebAppAssessmentV2Summary
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - Web app Assessment arm name.
//   - summaryName - Gets the Name of the Web app Summary.
//   - options - WebAppAssessmentV2SummaryOperationsClientGetOptions contains the optional parameters for the WebAppAssessmentV2SummaryOperationsClient.Get
//     method.
func (client *WebAppAssessmentV2SummaryOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, summaryName string, options *WebAppAssessmentV2SummaryOperationsClientGetOptions) (WebAppAssessmentV2SummaryOperationsClientGetResponse, error) {
	var err error
	const operationName = "WebAppAssessmentV2SummaryOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, summaryName, options)
	if err != nil {
		return WebAppAssessmentV2SummaryOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppAssessmentV2SummaryOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebAppAssessmentV2SummaryOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WebAppAssessmentV2SummaryOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, summaryName string, _ *WebAppAssessmentV2SummaryOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/webAppAssessments/{assessmentName}/summaries/{summaryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if summaryName == "" {
		return nil, errors.New("parameter summaryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{summaryName}", url.PathEscape(summaryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebAppAssessmentV2SummaryOperationsClient) getHandleResponse(resp *http.Response) (WebAppAssessmentV2SummaryOperationsClientGetResponse, error) {
	result := WebAppAssessmentV2SummaryOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppAssessmentV2Summary); err != nil {
		return WebAppAssessmentV2SummaryOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWebAppAssessmentV2Pager - List WebAppAssessmentV2Summary resources by WebAppAssessmentV2
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - Web app Assessment arm name.
//   - options - WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Options contains the optional parameters for
//     the WebAppAssessmentV2SummaryOperationsClient.NewListByWebAppAssessmentV2Pager method.
func (client *WebAppAssessmentV2SummaryOperationsClient) NewListByWebAppAssessmentV2Pager(resourceGroupName string, projectName string, groupName string, assessmentName string, options *WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Options) *runtime.Pager[WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response] {
	return runtime.NewPager(runtime.PagingHandler[WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response]{
		More: func(page WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response) (WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebAppAssessmentV2SummaryOperationsClient.NewListByWebAppAssessmentV2Pager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWebAppAssessmentV2CreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, options)
			}, nil)
			if err != nil {
				return WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response{}, err
			}
			return client.listByWebAppAssessmentV2HandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWebAppAssessmentV2CreateRequest creates the ListByWebAppAssessmentV2 request.
func (client *WebAppAssessmentV2SummaryOperationsClient) listByWebAppAssessmentV2CreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, _ *WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Options) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/webAppAssessments/{assessmentName}/summaries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWebAppAssessmentV2HandleResponse handles the ListByWebAppAssessmentV2 response.
func (client *WebAppAssessmentV2SummaryOperationsClient) listByWebAppAssessmentV2HandleResponse(resp *http.Response) (WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response, error) {
	result := WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppAssessmentV2SummaryListResult); err != nil {
		return WebAppAssessmentV2SummaryOperationsClientListByWebAppAssessmentV2Response{}, err
	}
	return result, nil
}
