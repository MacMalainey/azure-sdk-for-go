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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PriceSheetServer is a fake server for instances of the armconsumption.PriceSheetClient type.
type PriceSheetServer struct {
	// Get is the fake for method PriceSheetClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *armconsumption.PriceSheetClientGetOptions) (resp azfake.Responder[armconsumption.PriceSheetClientGetResponse], errResp azfake.ErrorResponder)

	// GetByBillingPeriod is the fake for method PriceSheetClient.GetByBillingPeriod
	// HTTP status codes to indicate success: http.StatusOK
	GetByBillingPeriod func(ctx context.Context, billingPeriodName string, options *armconsumption.PriceSheetClientGetByBillingPeriodOptions) (resp azfake.Responder[armconsumption.PriceSheetClientGetByBillingPeriodResponse], errResp azfake.ErrorResponder)
}

// NewPriceSheetServerTransport creates a new instance of PriceSheetServerTransport with the provided implementation.
// The returned PriceSheetServerTransport instance is connected to an instance of armconsumption.PriceSheetClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPriceSheetServerTransport(srv *PriceSheetServer) *PriceSheetServerTransport {
	return &PriceSheetServerTransport{srv: srv}
}

// PriceSheetServerTransport connects instances of armconsumption.PriceSheetClient to instances of PriceSheetServer.
// Don't use this type directly, use NewPriceSheetServerTransport instead.
type PriceSheetServerTransport struct {
	srv *PriceSheetServer
}

// Do implements the policy.Transporter interface for PriceSheetServerTransport.
func (p *PriceSheetServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PriceSheetClient.Get":
		resp, err = p.dispatchGet(req)
	case "PriceSheetClient.GetByBillingPeriod":
		resp, err = p.dispatchGetByBillingPeriod(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PriceSheetServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/pricesheets/default`
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
	skiptokenUnescaped, err := url.QueryUnescape(qp.Get("$skiptoken"))
	if err != nil {
		return nil, err
	}
	skiptokenParam := getOptional(skiptokenUnescaped)
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
	var options *armconsumption.PriceSheetClientGetOptions
	if expandParam != nil || skiptokenParam != nil || topParam != nil {
		options = &armconsumption.PriceSheetClientGetOptions{
			Expand:    expandParam,
			Skiptoken: skiptokenParam,
			Top:       topParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PriceSheetResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PriceSheetServerTransport) dispatchGetByBillingPeriod(req *http.Request) (*http.Response, error) {
	if p.srv.GetByBillingPeriod == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByBillingPeriod not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Billing/billingPeriods/(?P<billingPeriodName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/pricesheets/default`
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
	skiptokenUnescaped, err := url.QueryUnescape(qp.Get("$skiptoken"))
	if err != nil {
		return nil, err
	}
	skiptokenParam := getOptional(skiptokenUnescaped)
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
	billingPeriodNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingPeriodName")])
	if err != nil {
		return nil, err
	}
	var options *armconsumption.PriceSheetClientGetByBillingPeriodOptions
	if expandParam != nil || skiptokenParam != nil || topParam != nil {
		options = &armconsumption.PriceSheetClientGetByBillingPeriodOptions{
			Expand:    expandParam,
			Skiptoken: skiptokenParam,
			Top:       topParam,
		}
	}
	respr, errRespr := p.srv.GetByBillingPeriod(req.Context(), billingPeriodNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PriceSheetResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}