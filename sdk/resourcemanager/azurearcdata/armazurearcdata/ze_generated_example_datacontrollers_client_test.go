//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/ListSubscriptionDataController.json
func ExampleDataControllersClient_ListInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	pager := client.ListInSubscription(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DataControllerResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/ListByResourceGroupDataController.json
func ExampleDataControllersClient_ListInGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	pager := client.ListInGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DataControllerResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/CreateOrUpdateDataController.json
func ExampleDataControllersClient_BeginPutDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPutDataController(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
		armazurearcdata.DataControllerResource{
			TrackedResource: armazurearcdata.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"mytag": to.StringPtr("myval"),
				},
			},
			ExtendedLocation: &armazurearcdata.ExtendedLocation{
				Name: to.StringPtr("<name>"),
				Type: armazurearcdata.ExtendedLocationTypesCustomLocation.ToPtr(),
			},
			Properties: &armazurearcdata.DataControllerProperties{
				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
					Password: to.StringPtr("<password>"),
					Username: to.StringPtr("<username>"),
				},
				ClusterID:      to.StringPtr("<cluster-id>"),
				ExtensionID:    to.StringPtr("<extension-id>"),
				Infrastructure: armazurearcdata.InfrastructureOnpremises.ToPtr(),
				LogAnalyticsWorkspaceConfig: &armazurearcdata.LogAnalyticsWorkspaceConfig{
					PrimaryKey:  to.StringPtr("<primary-key>"),
					WorkspaceID: to.StringPtr("<workspace-id>"),
				},
				LogsDashboardCredential: &armazurearcdata.BasicLoginInformation{
					Password: to.StringPtr("<password>"),
					Username: to.StringPtr("<username>"),
				},
				MetricsDashboardCredential: &armazurearcdata.BasicLoginInformation{
					Password: to.StringPtr("<password>"),
					Username: to.StringPtr("<username>"),
				},
				OnPremiseProperty: &armazurearcdata.OnPremiseProperty{
					ID:               to.StringPtr("<id>"),
					PublicSigningKey: to.StringPtr("<public-signing-key>"),
				},
				UploadServicePrincipal: &armazurearcdata.UploadServicePrincipal{
					Authority:    to.StringPtr("<authority>"),
					ClientID:     to.StringPtr("<client-id>"),
					ClientSecret: to.StringPtr("<client-secret>"),
					TenantID:     to.StringPtr("<tenant-id>"),
				},
				UploadWatermark: &armazurearcdata.UploadWatermark{
					Logs:    to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
					Metrics: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
					Usages:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t }()),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataControllerResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/DeleteDataController.json
func ExampleDataControllersClient_BeginDeleteDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeleteDataController(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/GetDataController.json
func ExampleDataControllersClient_GetDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	res, err := client.GetDataController(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataControllerResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/UpdateDataController.json
func ExampleDataControllersClient_PatchDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	res, err := client.PatchDataController(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
		armazurearcdata.DataControllerUpdate{
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataControllerResource.ID: %s\n", *res.ID)
}
