//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// OutboundFirewallRulesServer is a fake server for instances of the armsql.OutboundFirewallRulesClient type.
type OutboundFirewallRulesServer struct {
	// BeginCreateOrUpdate is the fake for method OutboundFirewallRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, parameters armsql.OutboundFirewallRule, options *armsql.OutboundFirewallRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.OutboundFirewallRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method OutboundFirewallRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *armsql.OutboundFirewallRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armsql.OutboundFirewallRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OutboundFirewallRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *armsql.OutboundFirewallRulesClientGetOptions) (resp azfake.Responder[armsql.OutboundFirewallRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method OutboundFirewallRulesClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armsql.OutboundFirewallRulesClientListByServerOptions) (resp azfake.PagerResponder[armsql.OutboundFirewallRulesClientListByServerResponse])
}

// NewOutboundFirewallRulesServerTransport creates a new instance of OutboundFirewallRulesServerTransport with the provided implementation.
// The returned OutboundFirewallRulesServerTransport instance is connected to an instance of armsql.OutboundFirewallRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOutboundFirewallRulesServerTransport(srv *OutboundFirewallRulesServer) *OutboundFirewallRulesServerTransport {
	return &OutboundFirewallRulesServerTransport{
		srv:                  srv,
		beginCreateOrUpdate:  newTracker[azfake.PollerResponder[armsql.OutboundFirewallRulesClientCreateOrUpdateResponse]](),
		beginDelete:          newTracker[azfake.PollerResponder[armsql.OutboundFirewallRulesClientDeleteResponse]](),
		newListByServerPager: newTracker[azfake.PagerResponder[armsql.OutboundFirewallRulesClientListByServerResponse]](),
	}
}

// OutboundFirewallRulesServerTransport connects instances of armsql.OutboundFirewallRulesClient to instances of OutboundFirewallRulesServer.
// Don't use this type directly, use NewOutboundFirewallRulesServerTransport instead.
type OutboundFirewallRulesServerTransport struct {
	srv                  *OutboundFirewallRulesServer
	beginCreateOrUpdate  *tracker[azfake.PollerResponder[armsql.OutboundFirewallRulesClientCreateOrUpdateResponse]]
	beginDelete          *tracker[azfake.PollerResponder[armsql.OutboundFirewallRulesClientDeleteResponse]]
	newListByServerPager *tracker[azfake.PagerResponder[armsql.OutboundFirewallRulesClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for OutboundFirewallRulesServerTransport.
func (o *OutboundFirewallRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OutboundFirewallRulesClient.BeginCreateOrUpdate":
		resp, err = o.dispatchBeginCreateOrUpdate(req)
	case "OutboundFirewallRulesClient.BeginDelete":
		resp, err = o.dispatchBeginDelete(req)
	case "OutboundFirewallRulesClient.Get":
		resp, err = o.dispatchGet(req)
	case "OutboundFirewallRulesClient.NewListByServerPager":
		resp, err = o.dispatchNewListByServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OutboundFirewallRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := o.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundFirewallRules/(?P<outboundRuleFqdn>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.OutboundFirewallRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		outboundRuleFqdnParam, err := url.PathUnescape(matches[regex.SubexpIndex("outboundRuleFqdn")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, outboundRuleFqdnParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		o.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		o.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		o.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (o *OutboundFirewallRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if o.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := o.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundFirewallRules/(?P<outboundRuleFqdn>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		outboundRuleFqdnParam, err := url.PathUnescape(matches[regex.SubexpIndex("outboundRuleFqdn")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginDelete(req.Context(), resourceGroupNameParam, serverNameParam, outboundRuleFqdnParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		o.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		o.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		o.beginDelete.remove(req)
	}

	return resp, nil
}

func (o *OutboundFirewallRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundFirewallRules/(?P<outboundRuleFqdn>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	outboundRuleFqdnParam, err := url.PathUnescape(matches[regex.SubexpIndex("outboundRuleFqdn")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, outboundRuleFqdnParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OutboundFirewallRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OutboundFirewallRulesServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := o.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundFirewallRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := o.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		o.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armsql.OutboundFirewallRulesClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		o.newListByServerPager.remove(req)
	}
	return resp, nil
}