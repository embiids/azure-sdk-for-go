//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/managedHsmCreateKey.json
func ExampleManagedHsmKeysClient_CreateIfNotExist() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedHsmKeysClient().CreateIfNotExist(ctx, "sample-group", "sample-managedhsm-name", "sample-key-name", armkeyvault.ManagedHsmKeyCreateParameters{
		Properties: &armkeyvault.ManagedHsmKeyProperties{
			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedHsmKey = armkeyvault.ManagedHsmKey{
	// 	Name: to.Ptr("sample-key-name"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name"),
	// 	Properties: &armkeyvault.ManagedHsmKeyProperties{
	// 		Attributes: &armkeyvault.ManagedHsmKeyAttributes{
	// 			Created: to.Ptr[int64](1598533051),
	// 			Enabled: to.Ptr(true),
	// 			RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
	// 			Updated: to.Ptr[int64](1598533051),
	// 		},
	// 		KeyOps: []*armkeyvault.JSONWebKeyOperation{
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationEncrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationDecrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationSign),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationVerify),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationWrapKey),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationUnwrapKey)},
	// 			KeySize: to.Ptr[int32](2048),
	// 			KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name"),
	// 			KeyURIWithVersion: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name/fd618d9519b74f9aae94ade66b876acc"),
	// 			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/managedHsmGetKey.json
func ExampleManagedHsmKeysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedHsmKeysClient().Get(ctx, "sample-group", "sample-managedhsm-name", "sample-key-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedHsmKey = armkeyvault.ManagedHsmKey{
	// 	Name: to.Ptr("sample-key-name"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name"),
	// 	Properties: &armkeyvault.ManagedHsmKeyProperties{
	// 		Attributes: &armkeyvault.ManagedHsmKeyAttributes{
	// 			Created: to.Ptr[int64](1598533051),
	// 			Enabled: to.Ptr(true),
	// 			RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
	// 			Updated: to.Ptr[int64](1598533051),
	// 		},
	// 		KeyOps: []*armkeyvault.JSONWebKeyOperation{
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationEncrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationDecrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationSign),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationVerify),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationWrapKey),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationUnwrapKey)},
	// 			KeySize: to.Ptr[int32](2048),
	// 			KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name"),
	// 			KeyURIWithVersion: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name/fd618d9519b74f9aae94ade66b876acc"),
	// 			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/managedHsmListKeys.json
func ExampleManagedHsmKeysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedHsmKeysClient().NewListPager("sample-group", "sample-managedhsm-name", nil)
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
		// page.ManagedHsmKeyListResult = armkeyvault.ManagedHsmKeyListResult{
		// 	Value: []*armkeyvault.ManagedHsmKey{
		// 		{
		// 			Name: to.Ptr("sample-key-name-1"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name-1"),
		// 			Properties: &armkeyvault.ManagedHsmKeyProperties{
		// 				Attributes: &armkeyvault.ManagedHsmKeyAttributes{
		// 					Created: to.Ptr[int64](1596493796),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1596493796),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name-1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample-key-name-2"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name-2"),
		// 			Properties: &armkeyvault.ManagedHsmKeyProperties{
		// 				Attributes: &armkeyvault.ManagedHsmKeyAttributes{
		// 					Created: to.Ptr[int64](1596493797),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1596493797),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name-2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/managedHsmGetKeyVersion.json
func ExampleManagedHsmKeysClient_GetVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedHsmKeysClient().GetVersion(ctx, "sample-group", "sample-managedhsm-name", "sample-key-name", "fd618d9519b74f9aae94ade66b876acc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedHsmKey = armkeyvault.ManagedHsmKey{
	// 	Name: to.Ptr("fd618d9519b74f9aae94ade66b876acc"),
	// 	Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys/versions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name/versions/fd618d9519b74f9aae94ade66b876acc"),
	// 	Properties: &armkeyvault.ManagedHsmKeyProperties{
	// 		Attributes: &armkeyvault.ManagedHsmKeyAttributes{
	// 			Created: to.Ptr[int64](1598533051),
	// 			Enabled: to.Ptr(true),
	// 			RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
	// 			Updated: to.Ptr[int64](1598533051),
	// 		},
	// 		KeyOps: []*armkeyvault.JSONWebKeyOperation{
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationEncrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationDecrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationSign),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationVerify),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationWrapKey),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationUnwrapKey)},
	// 			KeySize: to.Ptr[int32](2048),
	// 			KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name"),
	// 			KeyURIWithVersion: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name/fd618d9519b74f9aae94ade66b876acc"),
	// 			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/managedHsmListKeyVersions.json
func ExampleManagedHsmKeysClient_NewListVersionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedHsmKeysClient().NewListVersionsPager("sample-group", "sample-managedhsm-name", "sample-key-name", nil)
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
		// page.ManagedHsmKeyListResult = armkeyvault.ManagedHsmKeyListResult{
		// 	Value: []*armkeyvault.ManagedHsmKey{
		// 		{
		// 			Name: to.Ptr("c2296aa24acf4daf86942bff5aca73dd"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name/versions/c2296aa24acf4daf86942bff5aca73dd"),
		// 			Properties: &armkeyvault.ManagedHsmKeyProperties{
		// 				Attributes: &armkeyvault.ManagedHsmKeyAttributes{
		// 					Created: to.Ptr[int64](1598641074),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1598641074),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name"),
		// 				KeyURIWithVersion: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name/c2296aa24acf4daf86942bff5aca73dd"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("d5a04667b6f44b0ca62825f5eae93da6"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name/versions/d5a04667b6f44b0ca62825f5eae93da6"),
		// 			Properties: &armkeyvault.ManagedHsmKeyProperties{
		// 				Attributes: &armkeyvault.ManagedHsmKeyAttributes{
		// 					Created: to.Ptr[int64](1598641295),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1598641295),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name"),
		// 				KeyURIWithVersion: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name/d5a04667b6f44b0ca62825f5eae93da6"),
		// 			},
		// 	}},
		// }
	}
}
