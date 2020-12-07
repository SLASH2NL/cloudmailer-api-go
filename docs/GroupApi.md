# \GroupApi

All URIs are relative to *https://app.cloudmailer.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupImportJson**](GroupApi.md#GroupImportJson) | **Post** /groups/{id}/import | Import JSON into a group
[**GroupIndex**](GroupApi.md#GroupIndex) | **Get** /groups | Get groups paginated
[**GroupIndexMember**](GroupApi.md#GroupIndexMember) | **Get** /groups/{id}/members | Get members in group paginated



## GroupImportJson

> GroupImportJson(ctx, id, body)

Import JSON into a group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| ID of the group | 
**body** | [**GroupImportRequest**](GroupImportRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupIndex

> GroupIndexResponse GroupIndex(ctx, optional)

Get groups paginated

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupIndexOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number | 
 **search** | **optional.String**| Search string | 

### Return type

[**GroupIndexResponse**](GroupIndexResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupIndexMember

> GroupMemberIndexResponse GroupIndexMember(ctx, id, optional)

Get members in group paginated

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| ID of the group | 
 **optional** | ***GroupIndexMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupIndexMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number | 
 **search** | **optional.String**| Search string | 

### Return type

[**GroupMemberIndexResponse**](GroupMemberIndexResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

