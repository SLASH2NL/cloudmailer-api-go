# \AutomationApi

All URIs are relative to *https://app.cloudmailer.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationIndex**](AutomationApi.md#AutomationIndex) | **Get** /automations | Get automations paginated
[**AutomationInternName**](AutomationApi.md#AutomationInternName) | **Get** /automations/{internName}/name | Show automation using the internal name
[**AutomationSend**](AutomationApi.md#AutomationSend) | **Post** /automations/{internName}/send | Send an automation
[**AutomationTrackId**](AutomationApi.md#AutomationTrackId) | **Get** /automations/track/{trackId}/id | Track automation by ID
[**AutomationTrackRef**](AutomationApi.md#AutomationTrackRef) | **Get** /automations/track/{trackRef}/ref | Track automation by reference



## AutomationIndex

> AutomationIndexResponse AutomationIndex(ctx, optional)

Get automations paginated

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutomationIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AutomationIndexOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number | 
 **search** | **optional.String**| Search string | 

### Return type

[**AutomationIndexResponse**](AutomationIndexResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationInternName

> AutomationShowNameResponse AutomationInternName(ctx, internName)

Show automation using the internal name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internName** | **string**| Internal name of the automation | 

### Return type

[**AutomationShowNameResponse**](AutomationShowNameResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationSend

> AutomationSendResponse AutomationSend(ctx, internName, body)

Send an automation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internName** | **string**| Internal name of the automation | 
**body** | [**AutomationSendRequest**](AutomationSendRequest.md)|  | 

### Return type

[**AutomationSendResponse**](AutomationSendResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationTrackId

> AutomationTrackIdResponse AutomationTrackId(ctx, trackId)

Track automation by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackId** | **string**| Track ID of the automation | 

### Return type

[**AutomationTrackIdResponse**](AutomationTrackIdResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationTrackRef

> AutomationTrackRefResponse AutomationTrackRef(ctx, trackRef)

Track automation by reference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackRef** | **string**| Track reference of the automation | 

### Return type

[**AutomationTrackRefResponse**](AutomationTrackRefResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

