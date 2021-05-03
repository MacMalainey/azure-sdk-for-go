// Package documentdb implements the Azure ARM Documentdb service API version 2021-04-01-preview.
//
//
package documentdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Documentdb
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Documentdb.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// LocationGet get the properties of an existing Cosmos DB location
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
func (client BaseClient) LocationGet(ctx context.Context, location string) (result LocationGetResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.LocationGet")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.BaseClient", "LocationGet", err.Error())
	}

	req, err := client.LocationGetPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.BaseClient", "LocationGet", nil, "Failure preparing request")
		return
	}

	resp, err := client.LocationGetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.BaseClient", "LocationGet", resp, "Failure sending request")
		return
	}

	result, err = client.LocationGetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.BaseClient", "LocationGet", resp, "Failure responding to request")
		return
	}

	return
}

// LocationGetPreparer prepares the LocationGet request.
func (client BaseClient) LocationGetPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// LocationGetSender sends the LocationGet request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) LocationGetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// LocationGetResponder handles the response to the LocationGet request. The method always
// closes the http.Response Body.
func (client BaseClient) LocationGetResponder(resp *http.Response) (result LocationGetResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// LocationList list Cosmos DB locations and their properties
func (client BaseClient) LocationList(ctx context.Context) (result LocationListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.LocationList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.BaseClient", "LocationList", err.Error())
	}

	req, err := client.LocationListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.BaseClient", "LocationList", nil, "Failure preparing request")
		return
	}

	resp, err := client.LocationListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.BaseClient", "LocationList", resp, "Failure sending request")
		return
	}

	result, err = client.LocationListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.BaseClient", "LocationList", resp, "Failure responding to request")
		return
	}

	return
}

// LocationListPreparer prepares the LocationList request.
func (client BaseClient) LocationListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// LocationListSender sends the LocationList request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) LocationListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// LocationListResponder handles the response to the LocationList request. The method always
// closes the http.Response Body.
func (client BaseClient) LocationListResponder(resp *http.Response) (result LocationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
