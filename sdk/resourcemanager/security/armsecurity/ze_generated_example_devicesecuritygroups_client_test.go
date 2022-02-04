//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/ListDeviceSecurityGroups_example.json
func ExampleDeviceSecurityGroupsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewDeviceSecurityGroupsClient(cred, nil)
	pager := client.List("<resource-id>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/GetDeviceSecurityGroups_example.json
func ExampleDeviceSecurityGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewDeviceSecurityGroupsClient(cred, nil)
	res, err := client.Get(ctx,
		"<resource-id>",
		"<device-security-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DeviceSecurityGroupsClientGetResult)
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/PutDeviceSecurityGroups_example.json
func ExampleDeviceSecurityGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewDeviceSecurityGroupsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-id>",
		"<device-security-group-name>",
		armsecurity.DeviceSecurityGroup{
			Properties: &armsecurity.DeviceSecurityGroupProperties{
				TimeWindowRules: []armsecurity.TimeWindowCustomAlertRuleClassification{
					&armsecurity.ActiveConnectionsNotInAllowedRange{
						IsEnabled:      to.BoolPtr(true),
						RuleType:       to.StringPtr("<rule-type>"),
						MaxThreshold:   to.Int32Ptr(30),
						MinThreshold:   to.Int32Ptr(0),
						TimeWindowSize: to.StringPtr("<time-window-size>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DeviceSecurityGroupsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/DeleteDeviceSecurityGroups_example.json
func ExampleDeviceSecurityGroupsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewDeviceSecurityGroupsClient(cred, nil)
	_, err = client.Delete(ctx,
		"<resource-id>",
		"<device-security-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}