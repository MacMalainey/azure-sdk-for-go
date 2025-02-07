//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	moduleVersion = "v1.3.0-beta.1"
)

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

type RecordType string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeCAA   RecordType = "CAA"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeDS    RecordType = "DS"
	RecordTypeMX    RecordType = "MX"
	RecordTypeNAPTR RecordType = "NAPTR"
	RecordTypeNS    RecordType = "NS"
	RecordTypePTR   RecordType = "PTR"
	RecordTypeSOA   RecordType = "SOA"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTLSA  RecordType = "TLSA"
	RecordTypeTXT   RecordType = "TXT"
)

// PossibleRecordTypeValues returns the possible values for the RecordType const type.
func PossibleRecordTypeValues() []RecordType {
	return []RecordType{
		RecordTypeA,
		RecordTypeAAAA,
		RecordTypeCAA,
		RecordTypeCNAME,
		RecordTypeDS,
		RecordTypeMX,
		RecordTypeNAPTR,
		RecordTypeNS,
		RecordTypePTR,
		RecordTypeSOA,
		RecordTypeSRV,
		RecordTypeTLSA,
		RecordTypeTXT,
	}
}

// ZoneType - The type of this DNS zone (Public or Private).
type ZoneType string

const (
	ZoneTypePrivate ZoneType = "Private"
	ZoneTypePublic  ZoneType = "Public"
)

// PossibleZoneTypeValues returns the possible values for the ZoneType const type.
func PossibleZoneTypeValues() []ZoneType {
	return []ZoneType{
		ZoneTypePrivate,
		ZoneTypePublic,
	}
}
