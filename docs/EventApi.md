# \EventApi

All URIs are relative to *https://app.cloudmailer.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventTriggerGeneric**](EventApi.md#EventTriggerGeneric) | **Post** /events/generic/trigger | Trigger a generic event
[**EventTriggerOrderCreated**](EventApi.md#EventTriggerOrderCreated) | **Post** /events/order-created/trigger | Trigger the order created event
[**EventTriggerOrderPaid**](EventApi.md#EventTriggerOrderPaid) | **Post** /events/order-paid/trigger | Trigger the order paid event



## EventTriggerGeneric

> EventTriggerResponse EventTriggerGeneric(ctx, body)

Trigger a generic event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventGenericRequest**](EventGenericRequest.md)|  | 

### Return type

[**EventTriggerResponse**](EventTriggerResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventTriggerOrderCreated

> EventTriggerResponse EventTriggerOrderCreated(ctx, body)

Trigger the order created event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventOrderRequest**](EventOrderRequest.md)|  | 

### Return type

[**EventTriggerResponse**](EventTriggerResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventTriggerOrderPaid

> EventTriggerResponse EventTriggerOrderPaid(ctx, body)

Trigger the order paid event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EventOrderRequest**](EventOrderRequest.md)|  | 

### Return type

[**EventTriggerResponse**](EventTriggerResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

