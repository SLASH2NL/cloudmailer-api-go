# \FlowApi

All URIs are relative to *https://app.cloudmailer.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowConfigTrigger**](FlowApi.md#FlowConfigTrigger) | **Post** /flow-configs/{id}/trigger | Trigger a specific flow config



## FlowConfigTrigger

> FlowConfigTriggerResponse FlowConfigTrigger(ctx, id, body)

Trigger a specific flow config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| ID of the flow config | 
**body** | [**FlowConfigTriggerRequest**](FlowConfigTriggerRequest.md)|  | 

### Return type

[**FlowConfigTriggerResponse**](FlowConfigTriggerResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

