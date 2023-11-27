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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
	"net/http"
	"net/url"
	"regexp"
)

// BuildpackBindingServer is a fake server for instances of the armappplatform.BuildpackBindingClient type.
type BuildpackBindingServer struct {
	// BeginCreateOrUpdate is the fake for method BuildpackBindingClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, buildpackBinding armappplatform.BuildpackBindingResource, options *armappplatform.BuildpackBindingClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappplatform.BuildpackBindingClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BuildpackBindingClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, options *armappplatform.BuildpackBindingClientBeginDeleteOptions) (resp azfake.PollerResponder[armappplatform.BuildpackBindingClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BuildpackBindingClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, options *armappplatform.BuildpackBindingClientGetOptions) (resp azfake.Responder[armappplatform.BuildpackBindingClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BuildpackBindingClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *armappplatform.BuildpackBindingClientListOptions) (resp azfake.PagerResponder[armappplatform.BuildpackBindingClientListResponse])
}

// NewBuildpackBindingServerTransport creates a new instance of BuildpackBindingServerTransport with the provided implementation.
// The returned BuildpackBindingServerTransport instance is connected to an instance of armappplatform.BuildpackBindingClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBuildpackBindingServerTransport(srv *BuildpackBindingServer) *BuildpackBindingServerTransport {
	return &BuildpackBindingServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armappplatform.BuildpackBindingClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armappplatform.BuildpackBindingClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armappplatform.BuildpackBindingClientListResponse]](),
	}
}

// BuildpackBindingServerTransport connects instances of armappplatform.BuildpackBindingClient to instances of BuildpackBindingServer.
// Don't use this type directly, use NewBuildpackBindingServerTransport instead.
type BuildpackBindingServerTransport struct {
	srv                 *BuildpackBindingServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armappplatform.BuildpackBindingClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armappplatform.BuildpackBindingClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armappplatform.BuildpackBindingClientListResponse]]
}

// Do implements the policy.Transporter interface for BuildpackBindingServerTransport.
func (b *BuildpackBindingServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BuildpackBindingClient.BeginCreateOrUpdate":
		resp, err = b.dispatchBeginCreateOrUpdate(req)
	case "BuildpackBindingClient.BeginDelete":
		resp, err = b.dispatchBeginDelete(req)
	case "BuildpackBindingClient.Get":
		resp, err = b.dispatchGet(req)
	case "BuildpackBindingClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BuildpackBindingServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := b.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildpackBindings/(?P<buildpackBindingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.BuildpackBindingResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		buildpackBindingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildpackBindingName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, buildServiceNameParam, builderNameParam, buildpackBindingNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		b.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		b.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		b.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (b *BuildpackBindingServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildpackBindings/(?P<buildpackBindingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		buildpackBindingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildpackBindingName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), resourceGroupNameParam, serviceNameParam, buildServiceNameParam, builderNameParam, buildpackBindingNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *BuildpackBindingServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildpackBindings/(?P<buildpackBindingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
	if err != nil {
		return nil, err
	}
	builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
	if err != nil {
		return nil, err
	}
	buildpackBindingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildpackBindingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, buildServiceNameParam, builderNameParam, buildpackBindingNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BuildpackBindingResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BuildpackBindingServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildpackBindings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListPager(resourceGroupNameParam, serviceNameParam, buildServiceNameParam, builderNameParam, nil)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappplatform.BuildpackBindingClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}
