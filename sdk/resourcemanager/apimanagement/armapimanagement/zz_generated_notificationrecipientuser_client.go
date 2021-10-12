//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// NotificationRecipientUserClient contains the methods for the NotificationRecipientUser group.
// Don't use this type directly, use NewNotificationRecipientUserClient() instead.
type NotificationRecipientUserClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewNotificationRecipientUserClient creates a new instance of NotificationRecipientUserClient with the specified values.
func NewNotificationRecipientUserClient(con *arm.Connection, subscriptionID string) *NotificationRecipientUserClient {
	return &NotificationRecipientUserClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CheckEntityExists - Determine if the Notification Recipient User is subscribed to the notification.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotificationRecipientUserClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserCheckEntityExistsOptions) (NotificationRecipientUserCheckEntityExistsResponse, error) {
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, notificationName, userID, options)
	if err != nil {
		return NotificationRecipientUserCheckEntityExistsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationRecipientUserCheckEntityExistsResponse{}, err
	}
	result := NotificationRecipientUserCheckEntityExistsResponse{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *NotificationRecipientUserClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// CreateOrUpdate - Adds the API Management User to the list of Recipients for the Notification.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotificationRecipientUserClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserCreateOrUpdateOptions) (NotificationRecipientUserCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, notificationName, userID, options)
	if err != nil {
		return NotificationRecipientUserCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationRecipientUserCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return NotificationRecipientUserCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotificationRecipientUserClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NotificationRecipientUserClient) createOrUpdateHandleResponse(resp *http.Response) (NotificationRecipientUserCreateOrUpdateResponse, error) {
	result := NotificationRecipientUserCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientUserContract); err != nil {
		return NotificationRecipientUserCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NotificationRecipientUserClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Removes the API Management user from the list of Notification.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotificationRecipientUserClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserDeleteOptions) (NotificationRecipientUserDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, notificationName, userID, options)
	if err != nil {
		return NotificationRecipientUserDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationRecipientUserDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return NotificationRecipientUserDeleteResponse{}, client.deleteHandleError(resp)
	}
	return NotificationRecipientUserDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotificationRecipientUserClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *NotificationRecipientUserClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByNotification - Gets the list of the Notification Recipient User subscribed to the notification.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotificationRecipientUserClient) ListByNotification(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationRecipientUserListByNotificationOptions) (NotificationRecipientUserListByNotificationResponse, error) {
	req, err := client.listByNotificationCreateRequest(ctx, resourceGroupName, serviceName, notificationName, options)
	if err != nil {
		return NotificationRecipientUserListByNotificationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationRecipientUserListByNotificationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotificationRecipientUserListByNotificationResponse{}, client.listByNotificationHandleError(resp)
	}
	return client.listByNotificationHandleResponse(resp)
}

// listByNotificationCreateRequest creates the ListByNotification request.
func (client *NotificationRecipientUserClient) listByNotificationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationRecipientUserListByNotificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByNotificationHandleResponse handles the ListByNotification response.
func (client *NotificationRecipientUserClient) listByNotificationHandleResponse(resp *http.Response) (NotificationRecipientUserListByNotificationResponse, error) {
	result := NotificationRecipientUserListByNotificationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientUserCollection); err != nil {
		return NotificationRecipientUserListByNotificationResponse{}, err
	}
	return result, nil
}

// listByNotificationHandleError handles the ListByNotification error response.
func (client *NotificationRecipientUserClient) listByNotificationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
