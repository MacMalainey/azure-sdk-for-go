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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"net/http"
	"net/url"
	"regexp"
)

// RoleEligibilityScheduleRequestsServer is a fake server for instances of the armauthorization.RoleEligibilityScheduleRequestsClient type.
type RoleEligibilityScheduleRequestsServer struct {
	// Cancel is the fake for method RoleEligibilityScheduleRequestsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *armauthorization.RoleEligibilityScheduleRequestsClientCancelOptions) (resp azfake.Responder[armauthorization.RoleEligibilityScheduleRequestsClientCancelResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method RoleEligibilityScheduleRequestsClient.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, parameters armauthorization.RoleEligibilityScheduleRequest, options *armauthorization.RoleEligibilityScheduleRequestsClientCreateOptions) (resp azfake.Responder[armauthorization.RoleEligibilityScheduleRequestsClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RoleEligibilityScheduleRequestsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *armauthorization.RoleEligibilityScheduleRequestsClientGetOptions) (resp azfake.Responder[armauthorization.RoleEligibilityScheduleRequestsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListForScopePager is the fake for method RoleEligibilityScheduleRequestsClient.NewListForScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForScopePager func(scope string, options *armauthorization.RoleEligibilityScheduleRequestsClientListForScopeOptions) (resp azfake.PagerResponder[armauthorization.RoleEligibilityScheduleRequestsClientListForScopeResponse])

	// Validate is the fake for method RoleEligibilityScheduleRequestsClient.Validate
	// HTTP status codes to indicate success: http.StatusOK
	Validate func(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, parameters armauthorization.RoleEligibilityScheduleRequest, options *armauthorization.RoleEligibilityScheduleRequestsClientValidateOptions) (resp azfake.Responder[armauthorization.RoleEligibilityScheduleRequestsClientValidateResponse], errResp azfake.ErrorResponder)
}

// NewRoleEligibilityScheduleRequestsServerTransport creates a new instance of RoleEligibilityScheduleRequestsServerTransport with the provided implementation.
// The returned RoleEligibilityScheduleRequestsServerTransport instance is connected to an instance of armauthorization.RoleEligibilityScheduleRequestsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRoleEligibilityScheduleRequestsServerTransport(srv *RoleEligibilityScheduleRequestsServer) *RoleEligibilityScheduleRequestsServerTransport {
	return &RoleEligibilityScheduleRequestsServerTransport{
		srv:                  srv,
		newListForScopePager: newTracker[azfake.PagerResponder[armauthorization.RoleEligibilityScheduleRequestsClientListForScopeResponse]](),
	}
}

// RoleEligibilityScheduleRequestsServerTransport connects instances of armauthorization.RoleEligibilityScheduleRequestsClient to instances of RoleEligibilityScheduleRequestsServer.
// Don't use this type directly, use NewRoleEligibilityScheduleRequestsServerTransport instead.
type RoleEligibilityScheduleRequestsServerTransport struct {
	srv                  *RoleEligibilityScheduleRequestsServer
	newListForScopePager *tracker[azfake.PagerResponder[armauthorization.RoleEligibilityScheduleRequestsClientListForScopeResponse]]
}

// Do implements the policy.Transporter interface for RoleEligibilityScheduleRequestsServerTransport.
func (r *RoleEligibilityScheduleRequestsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RoleEligibilityScheduleRequestsClient.Cancel":
		resp, err = r.dispatchCancel(req)
	case "RoleEligibilityScheduleRequestsClient.Create":
		resp, err = r.dispatchCreate(req)
	case "RoleEligibilityScheduleRequestsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RoleEligibilityScheduleRequestsClient.NewListForScopePager":
		resp, err = r.dispatchNewListForScopePager(req)
	case "RoleEligibilityScheduleRequestsClient.Validate":
		resp, err = r.dispatchValidate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RoleEligibilityScheduleRequestsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if r.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleEligibilityScheduleRequests/(?P<roleEligibilityScheduleRequestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	roleEligibilityScheduleRequestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleEligibilityScheduleRequestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Cancel(req.Context(), scopeParam, roleEligibilityScheduleRequestNameParam, nil)
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

func (r *RoleEligibilityScheduleRequestsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if r.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleEligibilityScheduleRequests/(?P<roleEligibilityScheduleRequestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.RoleEligibilityScheduleRequest](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	roleEligibilityScheduleRequestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleEligibilityScheduleRequestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Create(req.Context(), scopeParam, roleEligibilityScheduleRequestNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleEligibilityScheduleRequest, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleEligibilityScheduleRequestsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleEligibilityScheduleRequests/(?P<roleEligibilityScheduleRequestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	roleEligibilityScheduleRequestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleEligibilityScheduleRequestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), scopeParam, roleEligibilityScheduleRequestNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleEligibilityScheduleRequest, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleEligibilityScheduleRequestsServerTransport) dispatchNewListForScopePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListForScopePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForScopePager not implemented")}
	}
	newListForScopePager := r.newListForScopePager.get(req)
	if newListForScopePager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleEligibilityScheduleRequests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.RoleEligibilityScheduleRequestsClientListForScopeOptions
		if filterParam != nil {
			options = &armauthorization.RoleEligibilityScheduleRequestsClientListForScopeOptions{
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListForScopePager(scopeParam, options)
		newListForScopePager = &resp
		r.newListForScopePager.add(req, newListForScopePager)
		server.PagerResponderInjectNextLinks(newListForScopePager, req, func(page *armauthorization.RoleEligibilityScheduleRequestsClientListForScopeResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForScopePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListForScopePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForScopePager) {
		r.newListForScopePager.remove(req)
	}
	return resp, nil
}

func (r *RoleEligibilityScheduleRequestsServerTransport) dispatchValidate(req *http.Request) (*http.Response, error) {
	if r.srv.Validate == nil {
		return nil, &nonRetriableError{errors.New("fake for method Validate not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleEligibilityScheduleRequests/(?P<roleEligibilityScheduleRequestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.RoleEligibilityScheduleRequest](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	roleEligibilityScheduleRequestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleEligibilityScheduleRequestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Validate(req.Context(), scopeParam, roleEligibilityScheduleRequestNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleEligibilityScheduleRequest, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
