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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
	"net/http"
)

// RuntimeVersionsServer is a fake server for instances of the armappplatform.RuntimeVersionsClient type.
type RuntimeVersionsServer struct {
	// ListRuntimeVersions is the fake for method RuntimeVersionsClient.ListRuntimeVersions
	// HTTP status codes to indicate success: http.StatusOK
	ListRuntimeVersions func(ctx context.Context, options *armappplatform.RuntimeVersionsClientListRuntimeVersionsOptions) (resp azfake.Responder[armappplatform.RuntimeVersionsClientListRuntimeVersionsResponse], errResp azfake.ErrorResponder)
}

// NewRuntimeVersionsServerTransport creates a new instance of RuntimeVersionsServerTransport with the provided implementation.
// The returned RuntimeVersionsServerTransport instance is connected to an instance of armappplatform.RuntimeVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRuntimeVersionsServerTransport(srv *RuntimeVersionsServer) *RuntimeVersionsServerTransport {
	return &RuntimeVersionsServerTransport{srv: srv}
}

// RuntimeVersionsServerTransport connects instances of armappplatform.RuntimeVersionsClient to instances of RuntimeVersionsServer.
// Don't use this type directly, use NewRuntimeVersionsServerTransport instead.
type RuntimeVersionsServerTransport struct {
	srv *RuntimeVersionsServer
}

// Do implements the policy.Transporter interface for RuntimeVersionsServerTransport.
func (r *RuntimeVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RuntimeVersionsClient.ListRuntimeVersions":
		resp, err = r.dispatchListRuntimeVersions(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RuntimeVersionsServerTransport) dispatchListRuntimeVersions(req *http.Request) (*http.Response, error) {
	if r.srv.ListRuntimeVersions == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListRuntimeVersions not implemented")}
	}
	respr, errRespr := r.srv.ListRuntimeVersions(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AvailableRuntimeVersions, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
