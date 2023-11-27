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

// RoleEligibilitySchedulesServer is a fake server for instances of the armauthorization.RoleEligibilitySchedulesClient type.
type RoleEligibilitySchedulesServer struct {
	// Get is the fake for method RoleEligibilitySchedulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, roleEligibilityScheduleName string, options *armauthorization.RoleEligibilitySchedulesClientGetOptions) (resp azfake.Responder[armauthorization.RoleEligibilitySchedulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListForScopePager is the fake for method RoleEligibilitySchedulesClient.NewListForScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForScopePager func(scope string, options *armauthorization.RoleEligibilitySchedulesClientListForScopeOptions) (resp azfake.PagerResponder[armauthorization.RoleEligibilitySchedulesClientListForScopeResponse])
}

// NewRoleEligibilitySchedulesServerTransport creates a new instance of RoleEligibilitySchedulesServerTransport with the provided implementation.
// The returned RoleEligibilitySchedulesServerTransport instance is connected to an instance of armauthorization.RoleEligibilitySchedulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRoleEligibilitySchedulesServerTransport(srv *RoleEligibilitySchedulesServer) *RoleEligibilitySchedulesServerTransport {
	return &RoleEligibilitySchedulesServerTransport{
		srv:                  srv,
		newListForScopePager: newTracker[azfake.PagerResponder[armauthorization.RoleEligibilitySchedulesClientListForScopeResponse]](),
	}
}

// RoleEligibilitySchedulesServerTransport connects instances of armauthorization.RoleEligibilitySchedulesClient to instances of RoleEligibilitySchedulesServer.
// Don't use this type directly, use NewRoleEligibilitySchedulesServerTransport instead.
type RoleEligibilitySchedulesServerTransport struct {
	srv                  *RoleEligibilitySchedulesServer
	newListForScopePager *tracker[azfake.PagerResponder[armauthorization.RoleEligibilitySchedulesClientListForScopeResponse]]
}

// Do implements the policy.Transporter interface for RoleEligibilitySchedulesServerTransport.
func (r *RoleEligibilitySchedulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RoleEligibilitySchedulesClient.Get":
		resp, err = r.dispatchGet(req)
	case "RoleEligibilitySchedulesClient.NewListForScopePager":
		resp, err = r.dispatchNewListForScopePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RoleEligibilitySchedulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleEligibilitySchedules/(?P<roleEligibilityScheduleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	roleEligibilityScheduleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleEligibilityScheduleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), scopeParam, roleEligibilityScheduleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleEligibilitySchedule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleEligibilitySchedulesServerTransport) dispatchNewListForScopePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListForScopePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForScopePager not implemented")}
	}
	newListForScopePager := r.newListForScopePager.get(req)
	if newListForScopePager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/roleEligibilitySchedules`
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
		var options *armauthorization.RoleEligibilitySchedulesClientListForScopeOptions
		if filterParam != nil {
			options = &armauthorization.RoleEligibilitySchedulesClientListForScopeOptions{
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListForScopePager(scopeParam, options)
		newListForScopePager = &resp
		r.newListForScopePager.add(req, newListForScopePager)
		server.PagerResponderInjectNextLinks(newListForScopePager, req, func(page *armauthorization.RoleEligibilitySchedulesClientListForScopeResponse, createLink func() string) {
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
