//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GovernanceAssignmentsClient contains the methods for the GovernanceAssignments group.
// Don't use this type directly, use NewGovernanceAssignmentsClient() instead.
type GovernanceAssignmentsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGovernanceAssignmentsClient creates a new instance of GovernanceAssignmentsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGovernanceAssignmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GovernanceAssignmentsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &GovernanceAssignmentsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or update a security GovernanceAssignment on the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// assessmentName - The Assessment Key - Unique key for the assessment type
// assignmentKey - The security governance assignment key - the assessment key of the required governance assignment
// governanceAssignment - GovernanceAssignment over a subscription scope
// options - GovernanceAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the GovernanceAssignmentsClient.CreateOrUpdate
// method.
func (client *GovernanceAssignmentsClient) CreateOrUpdate(ctx context.Context, scope string, assessmentName string, assignmentKey string, governanceAssignment GovernanceAssignment, options *GovernanceAssignmentsClientCreateOrUpdateOptions) (GovernanceAssignmentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, assessmentName, assignmentKey, governanceAssignment, options)
	if err != nil {
		return GovernanceAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GovernanceAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return GovernanceAssignmentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GovernanceAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, assessmentName string, assignmentKey string, governanceAssignment GovernanceAssignment, options *GovernanceAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments/{assignmentKey}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assignmentKey == "" {
		return nil, errors.New("parameter assignmentKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentKey}", url.PathEscape(assignmentKey))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, governanceAssignment)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GovernanceAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (GovernanceAssignmentsClientCreateOrUpdateResponse, error) {
	result := GovernanceAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceAssignment); err != nil {
		return GovernanceAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a GovernanceAssignment over a given scope
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// assessmentName - The Assessment Key - Unique key for the assessment type
// assignmentKey - The security governance assignment key - the assessment key of the required governance assignment
// options - GovernanceAssignmentsClientDeleteOptions contains the optional parameters for the GovernanceAssignmentsClient.Delete
// method.
func (client *GovernanceAssignmentsClient) Delete(ctx context.Context, scope string, assessmentName string, assignmentKey string, options *GovernanceAssignmentsClientDeleteOptions) (GovernanceAssignmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, assessmentName, assignmentKey, options)
	if err != nil {
		return GovernanceAssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GovernanceAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return GovernanceAssignmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return GovernanceAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GovernanceAssignmentsClient) deleteCreateRequest(ctx context.Context, scope string, assessmentName string, assignmentKey string, options *GovernanceAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments/{assignmentKey}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assignmentKey == "" {
		return nil, errors.New("parameter assignmentKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentKey}", url.PathEscape(assignmentKey))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a specific governanceAssignment for the requested scope by AssignmentKey
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// assessmentName - The Assessment Key - Unique key for the assessment type
// assignmentKey - The security governance assignment key - the assessment key of the required governance assignment
// options - GovernanceAssignmentsClientGetOptions contains the optional parameters for the GovernanceAssignmentsClient.Get
// method.
func (client *GovernanceAssignmentsClient) Get(ctx context.Context, scope string, assessmentName string, assignmentKey string, options *GovernanceAssignmentsClientGetOptions) (GovernanceAssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, assessmentName, assignmentKey, options)
	if err != nil {
		return GovernanceAssignmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GovernanceAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GovernanceAssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GovernanceAssignmentsClient) getCreateRequest(ctx context.Context, scope string, assessmentName string, assignmentKey string, options *GovernanceAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments/{assignmentKey}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assignmentKey == "" {
		return nil, errors.New("parameter assignmentKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentKey}", url.PathEscape(assignmentKey))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GovernanceAssignmentsClient) getHandleResponse(resp *http.Response) (GovernanceAssignmentsClientGetResponse, error) {
	result := GovernanceAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceAssignment); err != nil {
		return GovernanceAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get security governanceAssignments on all your resources inside a scope
// Generated from API version 2022-01-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// assessmentName - The Assessment Key - Unique key for the assessment type
// options - GovernanceAssignmentsClientListOptions contains the optional parameters for the GovernanceAssignmentsClient.List
// method.
func (client *GovernanceAssignmentsClient) NewListPager(scope string, assessmentName string, options *GovernanceAssignmentsClientListOptions) *runtime.Pager[GovernanceAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GovernanceAssignmentsClientListResponse]{
		More: func(page GovernanceAssignmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GovernanceAssignmentsClientListResponse) (GovernanceAssignmentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, assessmentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GovernanceAssignmentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GovernanceAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GovernanceAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GovernanceAssignmentsClient) listCreateRequest(ctx context.Context, scope string, assessmentName string, options *GovernanceAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GovernanceAssignmentsClient) listHandleResponse(resp *http.Response) (GovernanceAssignmentsClientListResponse, error) {
	result := GovernanceAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceAssignmentsList); err != nil {
		return GovernanceAssignmentsClientListResponse{}, err
	}
	return result, nil
}