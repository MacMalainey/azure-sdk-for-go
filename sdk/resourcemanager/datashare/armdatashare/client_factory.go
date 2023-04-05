//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatashare

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewAccountsClient() *AccountsClient {
	subClient, _ := NewAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConsumerInvitationsClient() *ConsumerInvitationsClient {
	subClient, _ := NewConsumerInvitationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataSetsClient() *DataSetsClient {
	subClient, _ := NewDataSetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataSetMappingsClient() *DataSetMappingsClient {
	subClient, _ := NewDataSetMappingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEmailRegistrationsClient() *EmailRegistrationsClient {
	subClient, _ := NewEmailRegistrationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInvitationsClient() *InvitationsClient {
	subClient, _ := NewInvitationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSharesClient() *SharesClient {
	subClient, _ := NewSharesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProviderShareSubscriptionsClient() *ProviderShareSubscriptionsClient {
	subClient, _ := NewProviderShareSubscriptionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewShareSubscriptionsClient() *ShareSubscriptionsClient {
	subClient, _ := NewShareSubscriptionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConsumerSourceDataSetsClient() *ConsumerSourceDataSetsClient {
	subClient, _ := NewConsumerSourceDataSetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSynchronizationSettingsClient() *SynchronizationSettingsClient {
	subClient, _ := NewSynchronizationSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTriggersClient() *TriggersClient {
	subClient, _ := NewTriggersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}