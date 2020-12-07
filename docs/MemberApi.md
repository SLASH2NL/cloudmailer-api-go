# \MemberApi

All URIs are relative to *https://app.cloudmailer.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberDelete**](MemberApi.md#MemberDelete) | **Delete** /members/{id} | Delete a member by ID
[**MemberDeleteEmail**](MemberApi.md#MemberDeleteEmail) | **Delete** /members/{email}/email | Delete a member by email
[**MemberShow**](MemberApi.md#MemberShow) | **Get** /members/{id} | Get a member by ID
[**MemberShowEmail**](MemberApi.md#MemberShowEmail) | **Get** /members/{email}/email | Get a member by email
[**MemberStoreUnvalidated**](MemberApi.md#MemberStoreUnvalidated) | **Post** /members/unvalidated | Subscribe and send validation email
[**MemberStoreValidated**](MemberApi.md#MemberStoreValidated) | **Post** /members/validated | Subscribe a member without a validation email
[**MemberUnsubscribe**](MemberApi.md#MemberUnsubscribe) | **Post** /members/unsubscribe | Unsubscribe a member from a client or groups
[**MemberUpdate**](MemberApi.md#MemberUpdate) | **Put** /members/{id} | Update a member by ID
[**MemberUpdateEmail**](MemberApi.md#MemberUpdateEmail) | **Put** /members/{email}/email | Update a member by email



## MemberDelete

> MessageResponse MemberDelete(ctx, id)

Delete a member by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| ID of the member | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberDeleteEmail

> MessageResponse MemberDeleteEmail(ctx, email)

Delete a member by email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| Email address of the member | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberShow

> MemberResponse MemberShow(ctx, id)

Get a member by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| ID of the member | 

### Return type

[**MemberResponse**](MemberResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberShowEmail

> MemberResponse MemberShowEmail(ctx, email)

Get a member by email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| Email address of the member | 

### Return type

[**MemberResponse**](MemberResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberStoreUnvalidated

> MemberStoreResponse MemberStoreUnvalidated(ctx, body)

Subscribe and send validation email

The meta_fields, groups and automation items are optional. Automation defines the mail to be sent with the link to validate the member.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MemberStoreRequest**](MemberStoreRequest.md)|  | 

### Return type

[**MemberStoreResponse**](MemberStoreResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberStoreValidated

> MemberStoreResponse MemberStoreValidated(ctx, body)

Subscribe a member without a validation email

The meta_fields and groups items are optional. The automation item is irrelevant. The member will be validated straight away and no mail will be sent.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MemberStoreRequest**](MemberStoreRequest.md)|  | 

### Return type

[**MemberStoreResponse**](MemberStoreResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberUnsubscribe

> MemberUnsubscribeResponse MemberUnsubscribe(ctx, body)

Unsubscribe a member from a client or groups

The groups item is optional. Without any group the member will be blacklisted for the client. This means the member is still visible inside the group, but will not receive any campaigns.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MemberUnsubscribeRequest**](MemberUnsubscribeRequest.md)|  | 

### Return type

[**MemberUnsubscribeResponse**](MemberUnsubscribeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberUpdate

> MemberUpdateResponse MemberUpdate(ctx, id, body)

Update a member by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| ID of the member | 
**body** | [**MemberUpdateRequest**](MemberUpdateRequest.md)|  | 

### Return type

[**MemberUpdateResponse**](MemberUpdateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberUpdateEmail

> MemberUpdateResponse MemberUpdateEmail(ctx, email, body)

Update a member by email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| Email address of the member | 
**body** | [**MemberUpdateRequest**](MemberUpdateRequest.md)|  | 

### Return type

[**MemberUpdateResponse**](MemberUpdateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

