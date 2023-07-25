//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccessReviewHistoryDefinitionInstancesClient contains the methods for the AccessReviewHistoryDefinitionInstances group.
// Don't use this type directly, use NewAccessReviewHistoryDefinitionInstancesClient() instead.
type AccessReviewHistoryDefinitionInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessReviewHistoryDefinitionInstancesClient creates a new instance of AccessReviewHistoryDefinitionInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessReviewHistoryDefinitionInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewHistoryDefinitionInstancesClient, error) {
	cl, err := arm.NewClient(moduleName+".AccessReviewHistoryDefinitionInstancesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewHistoryDefinitionInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get access review history definition instances by definition Id
//
// Generated from API version 2021-12-01-preview
//   - historyDefinitionID - The id of the access review history definition.
//   - options - AccessReviewHistoryDefinitionInstancesClientListOptions contains the optional parameters for the AccessReviewHistoryDefinitionInstancesClient.NewListPager
//     method.
func (client *AccessReviewHistoryDefinitionInstancesClient) NewListPager(historyDefinitionID string, options *AccessReviewHistoryDefinitionInstancesClientListOptions) *runtime.Pager[AccessReviewHistoryDefinitionInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessReviewHistoryDefinitionInstancesClientListResponse]{
		More: func(page AccessReviewHistoryDefinitionInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewHistoryDefinitionInstancesClientListResponse) (AccessReviewHistoryDefinitionInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, historyDefinitionID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessReviewHistoryDefinitionInstancesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AccessReviewHistoryDefinitionInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessReviewHistoryDefinitionInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewHistoryDefinitionInstancesClient) listCreateRequest(ctx context.Context, historyDefinitionID string, options *AccessReviewHistoryDefinitionInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/{historyDefinitionId}/instances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if historyDefinitionID == "" {
		return nil, errors.New("parameter historyDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyDefinitionId}", url.PathEscape(historyDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewHistoryDefinitionInstancesClient) listHandleResponse(resp *http.Response) (AccessReviewHistoryDefinitionInstancesClientListResponse, error) {
	result := AccessReviewHistoryDefinitionInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewHistoryDefinitionInstanceListResult); err != nil {
		return AccessReviewHistoryDefinitionInstancesClientListResponse{}, err
	}
	return result, nil
}