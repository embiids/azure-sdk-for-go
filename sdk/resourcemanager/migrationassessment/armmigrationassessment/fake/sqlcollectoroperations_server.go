// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
	"net/http"
	"net/url"
	"regexp"
)

// SQLCollectorOperationsServer is a fake server for instances of the armmigrationassessment.SQLCollectorOperationsClient type.
type SQLCollectorOperationsServer struct {
	// BeginCreate is the fake for method SQLCollectorOperationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, projectName string, collectorName string, resource armmigrationassessment.SQLCollector, options *armmigrationassessment.SQLCollectorOperationsClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrationassessment.SQLCollectorOperationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SQLCollectorOperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *armmigrationassessment.SQLCollectorOperationsClientDeleteOptions) (resp azfake.Responder[armmigrationassessment.SQLCollectorOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SQLCollectorOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *armmigrationassessment.SQLCollectorOperationsClientGetOptions) (resp azfake.Responder[armmigrationassessment.SQLCollectorOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAssessmentProjectPager is the fake for method SQLCollectorOperationsClient.NewListByAssessmentProjectPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAssessmentProjectPager func(resourceGroupName string, projectName string, options *armmigrationassessment.SQLCollectorOperationsClientListByAssessmentProjectOptions) (resp azfake.PagerResponder[armmigrationassessment.SQLCollectorOperationsClientListByAssessmentProjectResponse])
}

// NewSQLCollectorOperationsServerTransport creates a new instance of SQLCollectorOperationsServerTransport with the provided implementation.
// The returned SQLCollectorOperationsServerTransport instance is connected to an instance of armmigrationassessment.SQLCollectorOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLCollectorOperationsServerTransport(srv *SQLCollectorOperationsServer) *SQLCollectorOperationsServerTransport {
	return &SQLCollectorOperationsServerTransport{
		srv:                             srv,
		beginCreate:                     newTracker[azfake.PollerResponder[armmigrationassessment.SQLCollectorOperationsClientCreateResponse]](),
		newListByAssessmentProjectPager: newTracker[azfake.PagerResponder[armmigrationassessment.SQLCollectorOperationsClientListByAssessmentProjectResponse]](),
	}
}

// SQLCollectorOperationsServerTransport connects instances of armmigrationassessment.SQLCollectorOperationsClient to instances of SQLCollectorOperationsServer.
// Don't use this type directly, use NewSQLCollectorOperationsServerTransport instead.
type SQLCollectorOperationsServerTransport struct {
	srv                             *SQLCollectorOperationsServer
	beginCreate                     *tracker[azfake.PollerResponder[armmigrationassessment.SQLCollectorOperationsClientCreateResponse]]
	newListByAssessmentProjectPager *tracker[azfake.PagerResponder[armmigrationassessment.SQLCollectorOperationsClientListByAssessmentProjectResponse]]
}

// Do implements the policy.Transporter interface for SQLCollectorOperationsServerTransport.
func (s *SQLCollectorOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SQLCollectorOperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sqlCollectorOperationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sqlCollectorOperationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SQLCollectorOperationsClient.BeginCreate":
				res.resp, res.err = s.dispatchBeginCreate(req)
			case "SQLCollectorOperationsClient.Delete":
				res.resp, res.err = s.dispatchDelete(req)
			case "SQLCollectorOperationsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SQLCollectorOperationsClient.NewListByAssessmentProjectPager":
				res.resp, res.err = s.dispatchNewListByAssessmentProjectPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (s *SQLCollectorOperationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := s.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlcollectors/(?P<collectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationassessment.SQLCollector](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		collectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreate(req.Context(), resourceGroupNameParam, projectNameParam, collectorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		s.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		s.beginCreate.remove(req)
	}

	return resp, nil
}

func (s *SQLCollectorOperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlcollectors/(?P<collectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	collectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, projectNameParam, collectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLCollectorOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlcollectors/(?P<collectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	collectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, collectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLCollector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLCollectorOperationsServerTransport) dispatchNewListByAssessmentProjectPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByAssessmentProjectPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAssessmentProjectPager not implemented")}
	}
	newListByAssessmentProjectPager := s.newListByAssessmentProjectPager.get(req)
	if newListByAssessmentProjectPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlcollectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByAssessmentProjectPager(resourceGroupNameParam, projectNameParam, nil)
		newListByAssessmentProjectPager = &resp
		s.newListByAssessmentProjectPager.add(req, newListByAssessmentProjectPager)
		server.PagerResponderInjectNextLinks(newListByAssessmentProjectPager, req, func(page *armmigrationassessment.SQLCollectorOperationsClientListByAssessmentProjectResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAssessmentProjectPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByAssessmentProjectPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAssessmentProjectPager) {
		s.newListByAssessmentProjectPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SQLCollectorOperationsServerTransport
var sqlCollectorOperationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
