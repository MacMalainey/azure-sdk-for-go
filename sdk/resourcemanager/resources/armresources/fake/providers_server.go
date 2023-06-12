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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// ProvidersServer is a fake server for instances of the armresources.ProvidersClient type.
type ProvidersServer struct {
	// Get is the fake for method ProvidersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceProviderNamespace string, options *armresources.ProvidersClientGetOptions) (resp azfake.Responder[armresources.ProvidersClientGetResponse], errResp azfake.ErrorResponder)

	// GetAtTenantScope is the fake for method ProvidersClient.GetAtTenantScope
	// HTTP status codes to indicate success: http.StatusOK
	GetAtTenantScope func(ctx context.Context, resourceProviderNamespace string, options *armresources.ProvidersClientGetAtTenantScopeOptions) (resp azfake.Responder[armresources.ProvidersClientGetAtTenantScopeResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ProvidersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armresources.ProvidersClientListOptions) (resp azfake.PagerResponder[armresources.ProvidersClientListResponse])

	// NewListAtTenantScopePager is the fake for method ProvidersClient.NewListAtTenantScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAtTenantScopePager func(options *armresources.ProvidersClientListAtTenantScopeOptions) (resp azfake.PagerResponder[armresources.ProvidersClientListAtTenantScopeResponse])

	// ProviderPermissions is the fake for method ProvidersClient.ProviderPermissions
	// HTTP status codes to indicate success: http.StatusOK
	ProviderPermissions func(ctx context.Context, resourceProviderNamespace string, options *armresources.ProvidersClientProviderPermissionsOptions) (resp azfake.Responder[armresources.ProvidersClientProviderPermissionsResponse], errResp azfake.ErrorResponder)

	// Register is the fake for method ProvidersClient.Register
	// HTTP status codes to indicate success: http.StatusOK
	Register func(ctx context.Context, resourceProviderNamespace string, options *armresources.ProvidersClientRegisterOptions) (resp azfake.Responder[armresources.ProvidersClientRegisterResponse], errResp azfake.ErrorResponder)

	// RegisterAtManagementGroupScope is the fake for method ProvidersClient.RegisterAtManagementGroupScope
	// HTTP status codes to indicate success: http.StatusOK
	RegisterAtManagementGroupScope func(ctx context.Context, resourceProviderNamespace string, groupID string, options *armresources.ProvidersClientRegisterAtManagementGroupScopeOptions) (resp azfake.Responder[armresources.ProvidersClientRegisterAtManagementGroupScopeResponse], errResp azfake.ErrorResponder)

	// Unregister is the fake for method ProvidersClient.Unregister
	// HTTP status codes to indicate success: http.StatusOK
	Unregister func(ctx context.Context, resourceProviderNamespace string, options *armresources.ProvidersClientUnregisterOptions) (resp azfake.Responder[armresources.ProvidersClientUnregisterResponse], errResp azfake.ErrorResponder)
}

// NewProvidersServerTransport creates a new instance of ProvidersServerTransport with the provided implementation.
// The returned ProvidersServerTransport instance is connected to an instance of armresources.ProvidersClient by way of the
// undefined.Transporter field.
func NewProvidersServerTransport(srv *ProvidersServer) *ProvidersServerTransport {
	return &ProvidersServerTransport{srv: srv}
}

// ProvidersServerTransport connects instances of armresources.ProvidersClient to instances of ProvidersServer.
// Don't use this type directly, use NewProvidersServerTransport instead.
type ProvidersServerTransport struct {
	srv                       *ProvidersServer
	newListPager              *azfake.PagerResponder[armresources.ProvidersClientListResponse]
	newListAtTenantScopePager *azfake.PagerResponder[armresources.ProvidersClientListAtTenantScopeResponse]
}

// Do implements the policy.Transporter interface for ProvidersServerTransport.
func (p *ProvidersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProvidersClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProvidersClient.GetAtTenantScope":
		resp, err = p.dispatchGetAtTenantScope(req)
	case "ProvidersClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "ProvidersClient.NewListAtTenantScopePager":
		resp, err = p.dispatchNewListAtTenantScopePager(req)
	case "ProvidersClient.ProviderPermissions":
		resp, err = p.dispatchProviderPermissions(req)
	case "ProvidersClient.Register":
		resp, err = p.dispatchRegister(req)
	case "ProvidersClient.RegisterAtManagementGroupScope":
		resp, err = p.dispatchRegisterAtManagementGroupScope(req)
	case "ProvidersClient.Unregister":
		resp, err = p.dispatchUnregister(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProvidersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	resourceProviderNamespaceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	var options *armresources.ProvidersClientGetOptions
	if expandParam != nil {
		options = &armresources.ProvidersClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceProviderNamespaceUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Provider, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProvidersServerTransport) dispatchGetAtTenantScope(req *http.Request) (*http.Response, error) {
	if p.srv.GetAtTenantScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAtTenantScope not implemented")}
	}
	const regexStr = `/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	resourceProviderNamespaceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	var options *armresources.ProvidersClientGetAtTenantScopeOptions
	if expandParam != nil {
		options = &armresources.ProvidersClientGetAtTenantScopeOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.GetAtTenantScope(req.Context(), resourceProviderNamespaceUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Provider, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProvidersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if p.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armresources.ProvidersClientListOptions
		if expandParam != nil {
			options = &armresources.ProvidersClientListOptions{
				Expand: expandParam,
			}
		}
		resp := p.srv.NewListPager(options)
		p.newListPager = &resp
		server.PagerResponderInjectNextLinks(p.newListPager, req, func(page *armresources.ProvidersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListPager) {
		p.newListPager = nil
	}
	return resp, nil
}

func (p *ProvidersServerTransport) dispatchNewListAtTenantScopePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListAtTenantScopePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAtTenantScopePager not implemented")}
	}
	if p.newListAtTenantScopePager == nil {
		qp := req.URL.Query()
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armresources.ProvidersClientListAtTenantScopeOptions
		if expandParam != nil {
			options = &armresources.ProvidersClientListAtTenantScopeOptions{
				Expand: expandParam,
			}
		}
		resp := p.srv.NewListAtTenantScopePager(options)
		p.newListAtTenantScopePager = &resp
		server.PagerResponderInjectNextLinks(p.newListAtTenantScopePager, req, func(page *armresources.ProvidersClientListAtTenantScopeResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListAtTenantScopePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListAtTenantScopePager) {
		p.newListAtTenantScopePager = nil
	}
	return resp, nil
}

func (p *ProvidersServerTransport) dispatchProviderPermissions(req *http.Request) (*http.Response, error) {
	if p.srv.ProviderPermissions == nil {
		return nil, &nonRetriableError{errors.New("fake for method ProviderPermissions not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providerPermissions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceProviderNamespaceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ProviderPermissions(req.Context(), resourceProviderNamespaceUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProviderPermissionListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProvidersServerTransport) dispatchRegister(req *http.Request) (*http.Response, error) {
	if p.srv.Register == nil {
		return nil, &nonRetriableError{errors.New("fake for method Register not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/register`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armresources.ProviderRegistrationRequest](req)
	if err != nil {
		return nil, err
	}
	resourceProviderNamespaceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	var options *armresources.ProvidersClientRegisterOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armresources.ProvidersClientRegisterOptions{
			Properties: &body,
		}
	}
	respr, errRespr := p.srv.Register(req.Context(), resourceProviderNamespaceUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Provider, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProvidersServerTransport) dispatchRegisterAtManagementGroupScope(req *http.Request) (*http.Response, error) {
	if p.srv.RegisterAtManagementGroupScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegisterAtManagementGroupScope not implemented")}
	}
	const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/register`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceProviderNamespaceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	groupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.RegisterAtManagementGroupScope(req.Context(), resourceProviderNamespaceUnescaped, groupIDUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProvidersServerTransport) dispatchUnregister(req *http.Request) (*http.Response, error) {
	if p.srv.Unregister == nil {
		return nil, &nonRetriableError{errors.New("fake for method Unregister not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unregister`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceProviderNamespaceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Unregister(req.Context(), resourceProviderNamespaceUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Provider, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
