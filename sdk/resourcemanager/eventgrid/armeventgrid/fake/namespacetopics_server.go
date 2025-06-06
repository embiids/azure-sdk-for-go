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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// NamespaceTopicsServer is a fake server for instances of the armeventgrid.NamespaceTopicsClient type.
type NamespaceTopicsServer struct {
	// BeginCreateOrUpdate is the fake for method NamespaceTopicsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, namespaceTopicInfo armeventgrid.NamespaceTopic, options *armeventgrid.NamespaceTopicsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armeventgrid.NamespaceTopicsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NamespaceTopicsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *armeventgrid.NamespaceTopicsClientBeginDeleteOptions) (resp azfake.PollerResponder[armeventgrid.NamespaceTopicsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NamespaceTopicsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *armeventgrid.NamespaceTopicsClientGetOptions) (resp azfake.Responder[armeventgrid.NamespaceTopicsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByNamespacePager is the fake for method NamespaceTopicsClient.NewListByNamespacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNamespacePager func(resourceGroupName string, namespaceName string, options *armeventgrid.NamespaceTopicsClientListByNamespaceOptions) (resp azfake.PagerResponder[armeventgrid.NamespaceTopicsClientListByNamespaceResponse])

	// ListSharedAccessKeys is the fake for method NamespaceTopicsClient.ListSharedAccessKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListSharedAccessKeys func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *armeventgrid.NamespaceTopicsClientListSharedAccessKeysOptions) (resp azfake.Responder[armeventgrid.NamespaceTopicsClientListSharedAccessKeysResponse], errResp azfake.ErrorResponder)

	// BeginRegenerateKey is the fake for method NamespaceTopicsClient.BeginRegenerateKey
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRegenerateKey func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, regenerateKeyRequest armeventgrid.TopicRegenerateKeyRequest, options *armeventgrid.NamespaceTopicsClientBeginRegenerateKeyOptions) (resp azfake.PollerResponder[armeventgrid.NamespaceTopicsClientRegenerateKeyResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method NamespaceTopicsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, namespaceTopicUpdateParameters armeventgrid.NamespaceTopicUpdateParameters, options *armeventgrid.NamespaceTopicsClientBeginUpdateOptions) (resp azfake.PollerResponder[armeventgrid.NamespaceTopicsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewNamespaceTopicsServerTransport creates a new instance of NamespaceTopicsServerTransport with the provided implementation.
// The returned NamespaceTopicsServerTransport instance is connected to an instance of armeventgrid.NamespaceTopicsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNamespaceTopicsServerTransport(srv *NamespaceTopicsServer) *NamespaceTopicsServerTransport {
	return &NamespaceTopicsServerTransport{
		srv:                     srv,
		beginCreateOrUpdate:     newTracker[azfake.PollerResponder[armeventgrid.NamespaceTopicsClientCreateOrUpdateResponse]](),
		beginDelete:             newTracker[azfake.PollerResponder[armeventgrid.NamespaceTopicsClientDeleteResponse]](),
		newListByNamespacePager: newTracker[azfake.PagerResponder[armeventgrid.NamespaceTopicsClientListByNamespaceResponse]](),
		beginRegenerateKey:      newTracker[azfake.PollerResponder[armeventgrid.NamespaceTopicsClientRegenerateKeyResponse]](),
		beginUpdate:             newTracker[azfake.PollerResponder[armeventgrid.NamespaceTopicsClientUpdateResponse]](),
	}
}

// NamespaceTopicsServerTransport connects instances of armeventgrid.NamespaceTopicsClient to instances of NamespaceTopicsServer.
// Don't use this type directly, use NewNamespaceTopicsServerTransport instead.
type NamespaceTopicsServerTransport struct {
	srv                     *NamespaceTopicsServer
	beginCreateOrUpdate     *tracker[azfake.PollerResponder[armeventgrid.NamespaceTopicsClientCreateOrUpdateResponse]]
	beginDelete             *tracker[azfake.PollerResponder[armeventgrid.NamespaceTopicsClientDeleteResponse]]
	newListByNamespacePager *tracker[azfake.PagerResponder[armeventgrid.NamespaceTopicsClientListByNamespaceResponse]]
	beginRegenerateKey      *tracker[azfake.PollerResponder[armeventgrid.NamespaceTopicsClientRegenerateKeyResponse]]
	beginUpdate             *tracker[azfake.PollerResponder[armeventgrid.NamespaceTopicsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for NamespaceTopicsServerTransport.
func (n *NamespaceTopicsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NamespaceTopicsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if namespaceTopicsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = namespaceTopicsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "NamespaceTopicsClient.BeginCreateOrUpdate":
				res.resp, res.err = n.dispatchBeginCreateOrUpdate(req)
			case "NamespaceTopicsClient.BeginDelete":
				res.resp, res.err = n.dispatchBeginDelete(req)
			case "NamespaceTopicsClient.Get":
				res.resp, res.err = n.dispatchGet(req)
			case "NamespaceTopicsClient.NewListByNamespacePager":
				res.resp, res.err = n.dispatchNewListByNamespacePager(req)
			case "NamespaceTopicsClient.ListSharedAccessKeys":
				res.resp, res.err = n.dispatchListSharedAccessKeys(req)
			case "NamespaceTopicsClient.BeginRegenerateKey":
				res.resp, res.err = n.dispatchBeginRegenerateKey(req)
			case "NamespaceTopicsClient.BeginUpdate":
				res.resp, res.err = n.dispatchBeginUpdate(req)
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

func (n *NamespaceTopicsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.NamespaceTopic](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NamespaceTopicsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NamespaceTopicsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NamespaceTopic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespaceTopicsServerTransport) dispatchNewListByNamespacePager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByNamespacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNamespacePager not implemented")}
	}
	newListByNamespacePager := n.newListByNamespacePager.get(req)
	if newListByNamespacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armeventgrid.NamespaceTopicsClientListByNamespaceOptions
		if filterParam != nil || topParam != nil {
			options = &armeventgrid.NamespaceTopicsClientListByNamespaceOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := n.srv.NewListByNamespacePager(resourceGroupNameParam, namespaceNameParam, options)
		newListByNamespacePager = &resp
		n.newListByNamespacePager.add(req, newListByNamespacePager)
		server.PagerResponderInjectNextLinks(newListByNamespacePager, req, func(page *armeventgrid.NamespaceTopicsClientListByNamespaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNamespacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByNamespacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNamespacePager) {
		n.newListByNamespacePager.remove(req)
	}
	return resp, nil
}

func (n *NamespaceTopicsServerTransport) dispatchListSharedAccessKeys(req *http.Request) (*http.Response, error) {
	if n.srv.ListSharedAccessKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSharedAccessKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.ListSharedAccessKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TopicSharedAccessKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespaceTopicsServerTransport) dispatchBeginRegenerateKey(req *http.Request) (*http.Response, error) {
	if n.srv.BeginRegenerateKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRegenerateKey not implemented")}
	}
	beginRegenerateKey := n.beginRegenerateKey.get(req)
	if beginRegenerateKey == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKey`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.TopicRegenerateKeyRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginRegenerateKey(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRegenerateKey = &respr
		n.beginRegenerateKey.add(req, beginRegenerateKey)
	}

	resp, err := server.PollerResponderNext(beginRegenerateKey, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginRegenerateKey.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRegenerateKey) {
		n.beginRegenerateKey.remove(req)
	}

	return resp, nil
}

func (n *NamespaceTopicsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := n.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.NamespaceTopicUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpdate(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		n.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		n.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to NamespaceTopicsServerTransport
var namespaceTopicsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
