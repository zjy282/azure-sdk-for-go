//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ReplicationsClient contains the methods for the Replications group.
// Don't use this type directly, use NewReplicationsClient() instead.
type ReplicationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewReplicationsClient creates a new instance of ReplicationsClient with the specified values.
func NewReplicationsClient(con *arm.Connection, subscriptionID string) *ReplicationsClient {
	return &ReplicationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Creates a replication for a container registry with the specified parameters.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsBeginCreateOptions) (ReplicationsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, registryName, replicationName, replication, options)
	if err != nil {
		return ReplicationsCreatePollerResponse{}, err
	}
	result := ReplicationsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationsClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return ReplicationsCreatePollerResponse{}, err
	}
	result.Poller = &ReplicationsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a replication for a container registry with the specified parameters.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) create(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, replicationName, replication, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, replication)
}

// createHandleError handles the Create error response.
func (client *ReplicationsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes a replication from a container registry.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsBeginDeleteOptions) (ReplicationsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, replicationName, options)
	if err != nil {
		return ReplicationsDeletePollerResponse{}, err
	}
	result := ReplicationsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ReplicationsDeletePollerResponse{}, err
	}
	result.Poller = &ReplicationsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a replication from a container registry.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, replicationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ReplicationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the properties of the specified replication.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) Get(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsGetOptions) (ReplicationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, replicationName, options)
	if err != nil {
		return ReplicationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationsClient) getHandleResponse(resp *http.Response) (ReplicationsGetResponse, error) {
	result := ReplicationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Replication); err != nil {
		return ReplicationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ReplicationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Lists all the replications for the specified container registry.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) List(resourceGroupName string, registryName string, options *ReplicationsListOptions) *ReplicationsListPager {
	return &ReplicationsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		advancer: func(ctx context.Context, resp ReplicationsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReplicationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ReplicationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ReplicationsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationsClient) listHandleResponse(resp *http.Response) (ReplicationsListResponse, error) {
	result := ReplicationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationListResult); err != nil {
		return ReplicationsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ReplicationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - Updates a replication for a container registry with the specified parameters.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsBeginUpdateOptions) (ReplicationsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, registryName, replicationName, replicationUpdateParameters, options)
	if err != nil {
		return ReplicationsUpdatePollerResponse{}, err
	}
	result := ReplicationsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ReplicationsUpdatePollerResponse{}, err
	}
	result.Poller = &ReplicationsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a replication for a container registry with the specified parameters.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) update(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, replicationName, replicationUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ReplicationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, replicationUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *ReplicationsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
