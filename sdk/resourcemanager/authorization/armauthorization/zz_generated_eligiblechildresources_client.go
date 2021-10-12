//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// EligibleChildResourcesClient contains the methods for the EligibleChildResources group.
// Don't use this type directly, use NewEligibleChildResourcesClient() instead.
type EligibleChildResourcesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewEligibleChildResourcesClient creates a new instance of EligibleChildResourcesClient with the specified values.
func NewEligibleChildResourcesClient(con *arm.Connection) *EligibleChildResourcesClient {
	return &EligibleChildResourcesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Get - Get the child resources of a resource on which user has eligible access
// If the operation fails it returns the *CloudError error type.
func (client *EligibleChildResourcesClient) Get(scope string, options *EligibleChildResourcesGetOptions) *EligibleChildResourcesGetPager {
	return &EligibleChildResourcesGetPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp EligibleChildResourcesGetResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EligibleChildResourcesListResult.NextLink)
		},
	}
}

// getCreateRequest creates the Get request.
func (client *EligibleChildResourcesClient) getCreateRequest(ctx context.Context, scope string, options *EligibleChildResourcesGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/eligibleChildResources"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EligibleChildResourcesClient) getHandleResponse(resp *http.Response) (EligibleChildResourcesGetResponse, error) {
	result := EligibleChildResourcesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EligibleChildResourcesListResult); err != nil {
		return EligibleChildResourcesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *EligibleChildResourcesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
