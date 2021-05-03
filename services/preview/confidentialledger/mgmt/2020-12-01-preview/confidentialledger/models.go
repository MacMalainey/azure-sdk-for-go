package confidentialledger

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/confidentialledger/mgmt/2020-12-01-preview/confidentialledger"

// AADBasedSecurityPrincipal AAD based security principal with associated Ledger RoleName
type AADBasedSecurityPrincipal struct {
	// PrincipalID - UUID/GUID based Principal Id of the Security Principal
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - UUID/GUID based Tenant Id of the Security Principal
	TenantID *string `json:"tenantId,omitempty"`
	// LedgerRoleName - Possible values include: 'LedgerRoleNameReader', 'LedgerRoleNameContributor', 'LedgerRoleNameAdministrator'
	LedgerRoleName LedgerRoleName `json:"ledgerRoleName,omitempty"`
}

// CertBasedSecurityPrincipal cert based security principal with Ledger RoleName
type CertBasedSecurityPrincipal struct {
	// Cert - Base64 encoded public key of the user cert (.pem or .cer)
	Cert *string `json:"cert,omitempty"`
	// LedgerRoleName - Possible values include: 'LedgerRoleNameReader', 'LedgerRoleNameContributor', 'LedgerRoleNameAdministrator'
	LedgerRoleName LedgerRoleName `json:"ledgerRoleName,omitempty"`
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// ErrorDetail the error detail.
type ErrorDetail struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorDetail `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
}

// ErrorResponse common error response for all Azure Resource Manager APIs to return error details for
// failed operations. (This also follows the OData error response format.).
type ErrorResponse struct {
	// Error - The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// LedgerCreateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type LedgerCreateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(LedgerClient) (Model, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *LedgerCreateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for LedgerCreateFuture.Result.
func (future *LedgerCreateFuture) result(client LedgerClient) (mVar Model, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confidentialledger.LedgerCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		mVar.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("confidentialledger.LedgerCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if mVar.Response.Response, err = future.GetResult(sender); err == nil && mVar.Response.Response.StatusCode != http.StatusNoContent {
		mVar, err = client.CreateResponder(mVar.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "confidentialledger.LedgerCreateFuture", "Result", mVar.Response.Response, "Failure responding to request")
		}
	}
	return
}

// LedgerDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type LedgerDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(LedgerClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *LedgerDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for LedgerDeleteFuture.Result.
func (future *LedgerDeleteFuture) result(client LedgerClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confidentialledger.LedgerDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		ar.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("confidentialledger.LedgerDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// LedgerProperties additional Confidential Ledger properties.
type LedgerProperties struct {
	// LedgerName - READ-ONLY; Unique name for the Confidential Ledger.
	LedgerName *string `json:"ledgerName,omitempty"`
	// LedgerURI - READ-ONLY; Endpoint for calling Ledger Service.
	LedgerURI *string `json:"ledgerUri,omitempty"`
	// IdentityServiceURI - READ-ONLY; Endpoint for accessing network identity.
	IdentityServiceURI *string `json:"identityServiceUri,omitempty"`
	// LedgerInternalNamespace - READ-ONLY; Internal namespace for the Ledger
	LedgerInternalNamespace *string `json:"ledgerInternalNamespace,omitempty"`
	// LedgerStorageAccount - Name of the Blob Storage Account for saving ledger files
	LedgerStorageAccount *string `json:"ledgerStorageAccount,omitempty"`
	// LedgerType - Type of Confidential Ledger. Possible values include: 'LedgerTypeUnknown', 'LedgerTypePublic', 'LedgerTypePrivate'
	LedgerType LedgerType `json:"ledgerType,omitempty"`
	// ProvisioningState - READ-ONLY; Provisioning state of Ledger Resource. Possible values include: 'ProvisioningStateUnknown', 'ProvisioningStateSucceeded', 'ProvisioningStateFailed', 'ProvisioningStateCanceled', 'ProvisioningStateCreating', 'ProvisioningStateDeleting', 'ProvisioningStateUpdating'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// AadBasedSecurityPrincipals - Array of all AAD based Security Principals.
	AadBasedSecurityPrincipals *[]AADBasedSecurityPrincipal `json:"aadBasedSecurityPrincipals,omitempty"`
	// CertBasedSecurityPrincipals - Array of all cert based Security Principals.
	CertBasedSecurityPrincipals *[]CertBasedSecurityPrincipal `json:"certBasedSecurityPrincipals,omitempty"`
}

// MarshalJSON is the custom marshaler for LedgerProperties.
func (lp LedgerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if lp.LedgerStorageAccount != nil {
		objectMap["ledgerStorageAccount"] = lp.LedgerStorageAccount
	}
	if lp.LedgerType != "" {
		objectMap["ledgerType"] = lp.LedgerType
	}
	if lp.AadBasedSecurityPrincipals != nil {
		objectMap["aadBasedSecurityPrincipals"] = lp.AadBasedSecurityPrincipals
	}
	if lp.CertBasedSecurityPrincipals != nil {
		objectMap["certBasedSecurityPrincipals"] = lp.CertBasedSecurityPrincipals
	}
	return json.Marshal(objectMap)
}

// LedgerUpdateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type LedgerUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(LedgerClient) (Model, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *LedgerUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for LedgerUpdateFuture.Result.
func (future *LedgerUpdateFuture) result(client LedgerClient) (mVar Model, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confidentialledger.LedgerUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		mVar.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("confidentialledger.LedgerUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if mVar.Response.Response, err = future.GetResult(sender); err == nil && mVar.Response.Response.StatusCode != http.StatusNoContent {
		mVar, err = client.UpdateResponder(mVar.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "confidentialledger.LedgerUpdateFuture", "Result", mVar.Response.Response, "Failure responding to request")
		}
	}
	return
}

// List object that includes an array of Confidential Ledgers and a possible link for next set.
type List struct {
	autorest.Response `json:"-"`
	// Value - List of Confidential Ledgers
	Value *[]Model `json:"value,omitempty"`
	// NextLink - The URL the client should use to fetch the next page (per server side paging).
	NextLink *string `json:"nextLink,omitempty"`
}

// ListIterator provides access to a complete listing of Model values.
type ListIterator struct {
	i    int
	page ListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListIterator) Response() List {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListIterator) Value() Model {
	if !iter.page.NotDone() {
		return Model{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListIterator type.
func NewListIterator(page ListPage) ListIterator {
	return ListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (l List) IsEmpty() bool {
	return l.Value == nil || len(*l.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (l List) hasNextLink() bool {
	return l.NextLink != nil && len(*l.NextLink) != 0
}

// listPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (l List) listPreparer(ctx context.Context) (*http.Request, error) {
	if !l.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(l.NextLink)))
}

// ListPage contains a page of Model values.
type ListPage struct {
	fn func(context.Context, List) (List, error)
	l  List
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.l)
		if err != nil {
			return err
		}
		page.l = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListPage) NotDone() bool {
	return !page.l.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListPage) Response() List {
	return page.l
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListPage) Values() []Model {
	if page.l.IsEmpty() {
		return nil
	}
	return *page.l.Value
}

// Creates a new instance of the ListPage type.
func NewListPage(cur List, getNextPage func(context.Context, List) (List, error)) ListPage {
	return ListPage{
		fn: getNextPage,
		l:  cur,
	}
}

// Location location of the ARM Resource
type Location struct {
	// Location - The Azure location where the Confidential Ledger is running.
	Location *string `json:"location,omitempty"`
}

// Model confidential Ledger. Contains the properties of Confidential Ledger Resource.
type Model struct {
	autorest.Response `json:"-"`
	// Name - READ-ONLY; Name of the Resource.
	Name *string `json:"name,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// Location - The Azure location where the Confidential Ledger is running.
	Location *string `json:"location,omitempty"`
	// Tags - Additional tags for Confidential Ledger
	Tags map[string]*string `json:"tags"`
	// Properties - Properties of Confidential Ledger Resource.
	Properties *LedgerProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Model.
func (mVar Model) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mVar.Location != nil {
		objectMap["location"] = mVar.Location
	}
	if mVar.Tags != nil {
		objectMap["tags"] = mVar.Tags
	}
	if mVar.Properties != nil {
		objectMap["properties"] = mVar.Properties
	}
	return json.Marshal(objectMap)
}

// Resource an Azure resource.
type Resource struct {
	// Name - READ-ONLY; Name of the Resource.
	Name *string `json:"name,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty"`
}

// ResourceProviderOperationDefinition describes the Resource Provider Operation.
type ResourceProviderOperationDefinition struct {
	// Name - Resource provider operation name.
	Name *string `json:"name,omitempty"`
	// IsDataAction - Indicates whether the operation is data action or not.
	IsDataAction *bool `json:"isDataAction,omitempty"`
	// Display - Details about the operations
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`
}

// ResourceProviderOperationDisplay describes the properties of the Operation.
type ResourceProviderOperationDisplay struct {
	// Provider - Name of the resource provider.
	Provider *string `json:"provider,omitempty"`
	// Resource - Name of the resource type.
	Resource *string `json:"resource,omitempty"`
	// Operation - Name of the resource provider operation.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the resource provider operation.
	Description *string `json:"description,omitempty"`
}

// ResourceProviderOperationList list containing this Resource Provider's available operations.
type ResourceProviderOperationList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Resource provider operations list.
	Value *[]ResourceProviderOperationDefinition `json:"value,omitempty"`
	// NextLink - READ-ONLY; The URI that can be used to request the next page for list of Azure operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceProviderOperationListIterator provides access to a complete listing of
// ResourceProviderOperationDefinition values.
type ResourceProviderOperationListIterator struct {
	i    int
	page ResourceProviderOperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceProviderOperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ResourceProviderOperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceProviderOperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceProviderOperationListIterator) Response() ResourceProviderOperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceProviderOperationListIterator) Value() ResourceProviderOperationDefinition {
	if !iter.page.NotDone() {
		return ResourceProviderOperationDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceProviderOperationListIterator type.
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return ResourceProviderOperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rpol ResourceProviderOperationList) IsEmpty() bool {
	return rpol.Value == nil || len(*rpol.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rpol ResourceProviderOperationList) hasNextLink() bool {
	return rpol.NextLink != nil && len(*rpol.NextLink) != 0
}

// resourceProviderOperationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rpol ResourceProviderOperationList) resourceProviderOperationListPreparer(ctx context.Context) (*http.Request, error) {
	if !rpol.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rpol.NextLink)))
}

// ResourceProviderOperationListPage contains a page of ResourceProviderOperationDefinition values.
type ResourceProviderOperationListPage struct {
	fn   func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)
	rpol ResourceProviderOperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceProviderOperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rpol)
		if err != nil {
			return err
		}
		page.rpol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceProviderOperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceProviderOperationListPage) NotDone() bool {
	return !page.rpol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceProviderOperationListPage) Response() ResourceProviderOperationList {
	return page.rpol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceProviderOperationListPage) Values() []ResourceProviderOperationDefinition {
	if page.rpol.IsEmpty() {
		return nil
	}
	return *page.rpol.Value
}

// Creates a new instance of the ResourceProviderOperationListPage type.
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return ResourceProviderOperationListPage{
		fn:   getNextPage,
		rpol: cur,
	}
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'CreatedByTypeUser', 'CreatedByTypeApplication', 'CreatedByTypeManagedIdentity', 'CreatedByTypeKey'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'CreatedByTypeUser', 'CreatedByTypeApplication', 'CreatedByTypeManagedIdentity', 'CreatedByTypeKey'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource last modification (UTC)
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// Tags tags for Confidential Ledger Resource
type Tags struct {
	// Tags - Additional tags for Confidential Ledger
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Tags.
func (t Tags) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if t.Tags != nil {
		objectMap["tags"] = t.Tags
	}
	return json.Marshal(objectMap)
}
