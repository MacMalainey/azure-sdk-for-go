//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleByName.json
func ExampleRoleEligibilitySchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleEligibilitySchedulesClient().Get(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "b1477448-2cc6-4ceb-93b4-54a202a89413", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleEligibilitySchedule = armauthorization.RoleEligibilitySchedule{
	// 	Name: to.Ptr("b1477448-2cc6-4ceb-93b4-54a202a89413"),
	// 	Type: to.Ptr("Microsoft.Authorization/RoleEligibilitySchedules"),
	// 	ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilitySchedules/b1477448-2cc6-4ceb-93b4-54a202a89413"),
	// 	Properties: &armauthorization.RoleEligibilityScheduleProperties{
	// 		Condition: to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
	// 		ConditionVersion: to.Ptr("1.0"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:33:06.300Z"); return t}()),
	// 		EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-09T21:32:28.490Z"); return t}()),
	// 		ExpandedProperties: &armauthorization.ExpandedProperties{
	// 			Principal: &armauthorization.ExpandedPropertiesPrincipal{
	// 				Type: to.Ptr("User"),
	// 				DisplayName: to.Ptr("User Account"),
	// 				Email: to.Ptr("user@my-tenant.com"),
	// 				ID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 			},
	// 			RoleDefinition: &armauthorization.ExpandedPropertiesRoleDefinition{
	// 				Type: to.Ptr("BuiltInRole"),
	// 				DisplayName: to.Ptr("Contributor"),
	// 				ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
	// 			},
	// 			Scope: &armauthorization.ExpandedPropertiesScope{
	// 				Type: to.Ptr("subscription"),
	// 				DisplayName: to.Ptr("Pay-As-You-Go"),
	// 				ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 			},
	// 		},
	// 		MemberType: to.Ptr(armauthorization.MemberTypeDirect),
	// 		PrincipalID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c"),
	// 		RoleEligibilityScheduleRequestID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilityScheduleRequests/64caffb6-55c0-4deb-a585-68e948ea1ad6"),
	// 		Scope: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:33:14.557Z"); return t}()),
	// 		Status: to.Ptr(armauthorization.StatusProvisioned),
	// 		UpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T22:27:00.513Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilitySchedulesByScope.json
func ExampleRoleEligibilitySchedulesClient_NewListForScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleEligibilitySchedulesClient().NewListForScopePager("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", &armauthorization.RoleEligibilitySchedulesClientListForScopeOptions{Filter: to.Ptr("assignedTo('a3bb8764-cb92-4276-9d2a-ca1e895e55ea')")})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RoleEligibilityScheduleListResult = armauthorization.RoleEligibilityScheduleListResult{
		// 	Value: []*armauthorization.RoleEligibilitySchedule{
		// 		{
		// 			Name: to.Ptr("b1477448-2cc6-4ceb-93b4-54a202a89413"),
		// 			Type: to.Ptr("Microsoft.Authorization/RoleEligibilitySchedules"),
		// 			ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilitySchedules/b1477448-2cc6-4ceb-93b4-54a202a89413"),
		// 			Properties: &armauthorization.RoleEligibilityScheduleProperties{
		// 				Condition: to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
		// 				ConditionVersion: to.Ptr("1.0"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:33:06.300Z"); return t}()),
		// 				EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-09T21:32:28.490Z"); return t}()),
		// 				ExpandedProperties: &armauthorization.ExpandedProperties{
		// 					Principal: &armauthorization.ExpandedPropertiesPrincipal{
		// 						Type: to.Ptr("User"),
		// 						DisplayName: to.Ptr("User Account"),
		// 						Email: to.Ptr("user@my-tenant.com"),
		// 						ID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
		// 					},
		// 					RoleDefinition: &armauthorization.ExpandedPropertiesRoleDefinition{
		// 						Type: to.Ptr("BuiltInRole"),
		// 						DisplayName: to.Ptr("Contributor"),
		// 						ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
		// 					},
		// 					Scope: &armauthorization.ExpandedPropertiesScope{
		// 						Type: to.Ptr("subscription"),
		// 						DisplayName: to.Ptr("Pay-As-You-Go"),
		// 						ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
		// 					},
		// 				},
		// 				MemberType: to.Ptr(armauthorization.MemberTypeDirect),
		// 				PrincipalID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c"),
		// 				RoleEligibilityScheduleRequestID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilityScheduleRequests/64caffb6-55c0-4deb-a585-68e948ea1ad6"),
		// 				Scope: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
		// 				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:33:14.557Z"); return t}()),
		// 				Status: to.Ptr(armauthorization.StatusProvisioned),
		// 				UpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T22:27:00.513Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
