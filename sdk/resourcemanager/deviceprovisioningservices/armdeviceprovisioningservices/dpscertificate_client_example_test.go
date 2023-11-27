//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGetCertificate.json
func ExampleDpsCertificateClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDpsCertificateClient().Get(ctx, "cert", "myResourceGroup", "myFirstProvisioningService", &armdeviceprovisioningservices.DpsCertificateClientGetOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateResponse = armdeviceprovisioningservices.CertificateResponse{
	// 	Name: to.Ptr("cert"),
	// 	Type: to.Ptr("Microsoft.Devices/ProvisioningServices/Certificates"),
	// 	Etag: to.Ptr("AAAAAAExpNs="),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/andbuc-hub/certificates/cert"),
	// 	Properties: &armdeviceprovisioningservices.CertificateProperties{
	// 		Certificate: []byte("######################################"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		IsVerified: to.Ptr(false),
	// 		Subject: to.Ptr("CN=testdevice1"),
	// 		Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSCertificateCreateOrUpdate.json
func ExampleDpsCertificateClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDpsCertificateClient().CreateOrUpdate(ctx, "myResourceGroup", "myFirstProvisioningService", "cert", armdeviceprovisioningservices.CertificateResponse{
		Properties: &armdeviceprovisioningservices.CertificateProperties{
			Certificate: []byte("############################################"),
		},
	}, &armdeviceprovisioningservices.DpsCertificateClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateResponse = armdeviceprovisioningservices.CertificateResponse{
	// 	Name: to.Ptr("cert"),
	// 	Type: to.Ptr("Microsoft.Devices/ProvisioningServices/Certificates"),
	// 	Etag: to.Ptr("AAAAAAExpNs="),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServives/myFirstProvisioningService/certificates/cert"),
	// 	Properties: &armdeviceprovisioningservices.CertificateProperties{
	// 		Certificate: []byte("############################################"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		IsVerified: to.Ptr(false),
	// 		Subject: to.Ptr("CN=testdevice1"),
	// 		Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSDeleteCertificate.json
func ExampleDpsCertificateClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDpsCertificateClient().Delete(ctx, "myResourceGroup", "AAAAAAAADGk=", "myFirstProvisioningService", "cert", &armdeviceprovisioningservices.DpsCertificateClientDeleteOptions{CertificateName1: nil,
		CertificateIsVerified:    nil,
		CertificatePurpose:       nil,
		CertificateCreated:       nil,
		CertificateLastUpdated:   nil,
		CertificateHasPrivateKey: nil,
		CertificateNonce:         nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGetCertificates.json
func ExampleDpsCertificateClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDpsCertificateClient().List(ctx, "myResourceGroup", "myFirstProvisioningService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateListDescription = armdeviceprovisioningservices.CertificateListDescription{
	// 	Value: []*armdeviceprovisioningservices.CertificateResponse{
	// 		{
	// 			Name: to.Ptr("cert"),
	// 			Type: to.Ptr("Microsoft.Devices/ProvisioningServices/Certificates"),
	// 			Etag: to.Ptr("AAAAAAExpNs="),
	// 			ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/andbuc-hub/certificates/cert"),
	// 			Properties: &armdeviceprovisioningservices.CertificateProperties{
	// 				Certificate: []byte("############################################"),
	// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 				Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2039-12-31T23:59:59.000Z"); return t}()),
	// 				IsVerified: to.Ptr(false),
	// 				Subject: to.Ptr("CN=testdevice1"),
	// 				Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 				Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGenerateVerificationCode.json
func ExampleDpsCertificateClient_GenerateVerificationCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDpsCertificateClient().GenerateVerificationCode(ctx, "cert", "AAAAAAAADGk=", "myResourceGroup", "myFirstProvisioningService", &armdeviceprovisioningservices.DpsCertificateClientGenerateVerificationCodeOptions{CertificateName1: nil,
		CertificateIsVerified:    nil,
		CertificatePurpose:       nil,
		CertificateCreated:       nil,
		CertificateLastUpdated:   nil,
		CertificateHasPrivateKey: nil,
		CertificateNonce:         nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VerificationCodeResponse = armdeviceprovisioningservices.VerificationCodeResponse{
	// 	Name: to.Ptr("cert"),
	// 	Properties: &armdeviceprovisioningservices.VerificationCodeResponseProperties{
	// 		Certificate: []byte("###########################"),
	// 		Created: to.Ptr("Thu, 12 Oct 2017 19:23:50 GMT"),
	// 		Expiry: to.Ptr("Sat, 31 Dec 2039 23:59:59 GMT"),
	// 		IsVerified: to.Ptr(false),
	// 		Subject: to.Ptr("CN=andbucdevice1"),
	// 		Thumbprint: to.Ptr("##############################"),
	// 		Updated: to.Ptr("Thu, 12 Oct 2017 19:26:56 GMT"),
	// 		VerificationCode: to.Ptr("##################################"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSVerifyCertificate.json
func ExampleDpsCertificateClient_VerifyCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDpsCertificateClient().VerifyCertificate(ctx, "cert", "AAAAAAAADGk=", "myResourceGroup", "myFirstProvisioningService", armdeviceprovisioningservices.VerificationCodeRequest{
		Certificate: to.Ptr("#####################################"),
	}, &armdeviceprovisioningservices.DpsCertificateClientVerifyCertificateOptions{CertificateName1: nil,
		CertificateIsVerified:    nil,
		CertificatePurpose:       nil,
		CertificateCreated:       nil,
		CertificateLastUpdated:   nil,
		CertificateHasPrivateKey: nil,
		CertificateNonce:         nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateResponse = armdeviceprovisioningservices.CertificateResponse{
	// 	Name: to.Ptr("cert"),
	// 	Type: to.Ptr("Microsoft.Devices/ProvisioningServices/Certificates"),
	// 	Etag: to.Ptr("AAAAAAExpTQ="),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService/certificates/cert"),
	// 	Properties: &armdeviceprovisioningservices.CertificateProperties{
	// 		Certificate: []byte("#####################################"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		IsVerified: to.Ptr(true),
	// 		Subject: to.Ptr("CN=andbucdevice1"),
	// 		Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:26:56.000Z"); return t}()),
	// 	},
	// }
}
