//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

// CertificatePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetCertificateProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CertificateProperties, *ContentCertificateProperties, *KeyVaultCertificateProperties
type CertificatePropertiesClassification interface {
	// GetCertificateProperties returns the CertificateProperties content of the underlying type.
	GetCertificateProperties() *CertificateProperties
}

// UploadedUserSourceInfoClassification provides polymorphic access to related types.
// Call the interface's GetUploadedUserSourceInfo() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *JarUploadedUserSourceInfo, *NetCoreZipUploadedUserSourceInfo, *SourceUploadedUserSourceInfo, *UploadedUserSourceInfo
type UploadedUserSourceInfoClassification interface {
	UserSourceInfoClassification
	// GetUploadedUserSourceInfo returns the UploadedUserSourceInfo content of the underlying type.
	GetUploadedUserSourceInfo() *UploadedUserSourceInfo
}

// UserSourceInfoClassification provides polymorphic access to related types.
// Call the interface's GetUserSourceInfo() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BuildResultUserSourceInfo, *JarUploadedUserSourceInfo, *NetCoreZipUploadedUserSourceInfo, *SourceUploadedUserSourceInfo,
// - *UploadedUserSourceInfo, *UserSourceInfo
type UserSourceInfoClassification interface {
	// GetUserSourceInfo returns the UserSourceInfo content of the underlying type.
	GetUserSourceInfo() *UserSourceInfo
}
