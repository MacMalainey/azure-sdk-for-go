//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasCheckNameAvailability.json
func ExampleDisasterRecoveryConfigsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().CheckNameAvailability(ctx, "exampleResourceGroup", "sdk-Namespace-9080", armeventhub.CheckNameAvailabilityParameter{
		Name: to.Ptr("sdk-DisasterRecovery-9474"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResult = armeventhub.CheckNameAvailabilityResult{
	// 	Message: to.Ptr(""),
	// 	NameAvailable: to.Ptr(true),
	// 	Reason: to.Ptr(armeventhub.UnavailableReasonNone),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasList.json
func ExampleDisasterRecoveryConfigsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDisasterRecoveryConfigsClient().NewListPager("exampleResourceGroup", "sdk-Namespace-8859", nil)
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
		// page.ArmDisasterRecoveryListResult = armeventhub.ArmDisasterRecoveryListResult{
		// 	Value: []*armeventhub.ArmDisasterRecovery{
		// 		{
		// 			Name: to.Ptr("sdk-DisasterRecovery-3814"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/DisasterRecoveryConfig"),
		// 			ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-8859/disasterRecoveryConfig/sdk-DisasterRecovery-3814"),
		// 			Properties: &armeventhub.ArmDisasterRecoveryProperties{
		// 				PartnerNamespace: to.Ptr("sdk-Namespace-37"),
		// 				ProvisioningState: to.Ptr(armeventhub.ProvisioningStateDRSucceeded),
		// 				Role: to.Ptr(armeventhub.RoleDisasterRecoveryPrimary),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasCreate.json
func ExampleDisasterRecoveryConfigsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", armeventhub.ArmDisasterRecovery{
		Properties: &armeventhub.ArmDisasterRecoveryProperties{
			PartnerNamespace: to.Ptr("sdk-Namespace-37"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArmDisasterRecovery = armeventhub.ArmDisasterRecovery{
	// 	Name: to.Ptr("sdk-DisasterRecovery-3814"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/DisasterRecoveryConfig"),
	// 	ID: to.Ptr("/subscriptions/exampleResourceGroup/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-8859/disasterRecoveryConfig/sdk-DisasterRecovery-3814"),
	// 	Properties: &armeventhub.ArmDisasterRecoveryProperties{
	// 		PartnerNamespace: to.Ptr("sdk-Namespace-37"),
	// 		ProvisioningState: to.Ptr(armeventhub.ProvisioningStateDRSucceeded),
	// 		Role: to.Ptr(armeventhub.RoleDisasterRecoveryPrimary),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasDelete.json
func ExampleDisasterRecoveryConfigsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDisasterRecoveryConfigsClient().Delete(ctx, "exampleResourceGroup", "sdk-Namespace-5849", "sdk-DisasterRecovery-3814", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasGet.json
func ExampleDisasterRecoveryConfigsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().Get(ctx, "exampleResourceGroup", "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArmDisasterRecovery = armeventhub.ArmDisasterRecovery{
	// 	Name: to.Ptr("sdk-DisasterRecovery-3814"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/DisasterRecoveryConfig"),
	// 	ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-37/disasterRecoveryConfig/sdk-DisasterRecovery-3814"),
	// 	Properties: &armeventhub.ArmDisasterRecoveryProperties{
	// 		PartnerNamespace: to.Ptr("sdk-Namespace-8859"),
	// 		PendingReplicationOperationsCount: to.Ptr[int64](0),
	// 		ProvisioningState: to.Ptr(armeventhub.ProvisioningStateDRSucceeded),
	// 		Role: to.Ptr(armeventhub.RoleDisasterRecoverySecondary),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasBreakPairing.json
func ExampleDisasterRecoveryConfigsClient_BreakPairing() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDisasterRecoveryConfigsClient().BreakPairing(ctx, "exampleResourceGroup", "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasFailOver.json
func ExampleDisasterRecoveryConfigsClient_FailOver() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDisasterRecoveryConfigsClient().FailOver(ctx, "exampleResourceGroup", "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleListAll.json
func ExampleDisasterRecoveryConfigsClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDisasterRecoveryConfigsClient().NewListAuthorizationRulesPager("exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4047", nil)
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
		// page.AuthorizationRuleListResult = armeventhub.AuthorizationRuleListResult{
		// 	Value: []*armeventhub.AuthorizationRule{
		// 		{
		// 			Name: to.Ptr("RootManageSharedAccessKey"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/RootManageSharedAccessKey"),
		// 			Properties: &armeventhub.AuthorizationRuleProperties{
		// 				Rights: []*armeventhub.AccessRights{
		// 					to.Ptr(armeventhub.AccessRightsListen),
		// 					to.Ptr(armeventhub.AccessRightsManage),
		// 					to.Ptr(armeventhub.AccessRightsSend)},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sdk-Authrules-1067"),
		// 				Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
		// 				ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-1067"),
		// 				Properties: &armeventhub.AuthorizationRuleProperties{
		// 					Rights: []*armeventhub.AccessRights{
		// 						to.Ptr(armeventhub.AccessRightsListen),
		// 						to.Ptr(armeventhub.AccessRightsSend)},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("sdk-Authrules-1684"),
		// 					Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
		// 					ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-1684"),
		// 					Properties: &armeventhub.AuthorizationRuleProperties{
		// 						Rights: []*armeventhub.AccessRights{
		// 							to.Ptr(armeventhub.AccessRightsListen),
		// 							to.Ptr(armeventhub.AccessRightsSend)},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("sdk-Authrules-4879"),
		// 						Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
		// 						ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-4879"),
		// 						Properties: &armeventhub.AuthorizationRuleProperties{
		// 							Rights: []*armeventhub.AccessRights{
		// 								to.Ptr(armeventhub.AccessRightsListen),
		// 								to.Ptr(armeventhub.AccessRightsSend)},
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleGet.json
func ExampleDisasterRecoveryConfigsClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().GetAuthorizationRule(ctx, "exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4879", "sdk-Authrules-4879", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationRule = armeventhub.AuthorizationRule{
	// 	Name: to.Ptr("sdk-Authrules-4879"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-4879"),
	// 	Properties: &armeventhub.AuthorizationRuleProperties{
	// 		Rights: []*armeventhub.AccessRights{
	// 			to.Ptr(armeventhub.AccessRightsListen),
	// 			to.Ptr(armeventhub.AccessRightsSend)},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleListKey.json
func ExampleDisasterRecoveryConfigsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().ListKeys(ctx, "exampleResourceGroup", "sdk-Namespace-2702", "sdk-DisasterRecovery-4047", "sdk-Authrules-1746", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armeventhub.AccessKeys{
	// 	AliasPrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-disasterrecovery-4047.servicebus.windows-int.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=############################################"),
	// 	AliasSecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-disasterrecovery-4047.servicebus.windows-int.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=############################################"),
	// 	KeyName: to.Ptr("sdk-Authrules-1746"),
	// 	PrimaryKey: to.Ptr("############################################"),
	// 	SecondaryKey: to.Ptr("############################################"),
	// }
}
