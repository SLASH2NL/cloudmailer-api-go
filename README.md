# Go API client for cloudmailer

Reference for the public Cloudmailer API calls.

Generated packages use semantic versioning since 2.0.0. The endpoints (v1) itself should always be backwards compatible.

The API uses rate limiting based on requests per minute. Every API response contains the following two headers: 'X-RateLimit-Limit', 'X-RateLimit-Remaining'. `X-RateLimit-Limit` indicates the request limit per minute. `X-RateLimit-Remaining` indicates the amount of available requests.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 3.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./cloudmailer"
```

## Documentation for API Endpoints

All URIs are relative to *https://app.cloudmailer.nl/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AutomationApi* | [**AutomationIndex**](docs/AutomationApi.md#automationindex) | **Get** /automations | Get automations paginated
*AutomationApi* | [**AutomationInternName**](docs/AutomationApi.md#automationinternname) | **Get** /automations/{internName}/name | Show automation using the internal name
*AutomationApi* | [**AutomationSend**](docs/AutomationApi.md#automationsend) | **Post** /automations/{internName}/send | Send an automation
*AutomationApi* | [**AutomationTrackId**](docs/AutomationApi.md#automationtrackid) | **Get** /automations/track/{trackId}/id | Track automation by ID
*AutomationApi* | [**AutomationTrackRef**](docs/AutomationApi.md#automationtrackref) | **Get** /automations/track/{trackRef}/ref | Track automation by reference
*EventApi* | [**EventTriggerGeneric**](docs/EventApi.md#eventtriggergeneric) | **Post** /events/generic/trigger | Trigger a generic event
*EventApi* | [**EventTriggerOrderCreated**](docs/EventApi.md#eventtriggerordercreated) | **Post** /events/order-created/trigger | Trigger the order created event
*EventApi* | [**EventTriggerOrderPaid**](docs/EventApi.md#eventtriggerorderpaid) | **Post** /events/order-paid/trigger | Trigger the order paid event
*FlowApi* | [**FlowConfigTrigger**](docs/FlowApi.md#flowconfigtrigger) | **Post** /flow-configs/{id}/trigger | Trigger a specific flow config
*GroupApi* | [**GroupImportJson**](docs/GroupApi.md#groupimportjson) | **Post** /groups/{id}/import | Import JSON into a group
*GroupApi* | [**GroupIndex**](docs/GroupApi.md#groupindex) | **Get** /groups | Get groups paginated
*GroupApi* | [**GroupIndexMember**](docs/GroupApi.md#groupindexmember) | **Get** /groups/{id}/members | Get members in group paginated
*MemberApi* | [**MemberDelete**](docs/MemberApi.md#memberdelete) | **Delete** /members/{id} | Delete a member by ID
*MemberApi* | [**MemberDeleteEmail**](docs/MemberApi.md#memberdeleteemail) | **Delete** /members/{email}/email | Delete a member by email
*MemberApi* | [**MemberShow**](docs/MemberApi.md#membershow) | **Get** /members/{id} | Get a member by ID
*MemberApi* | [**MemberShowEmail**](docs/MemberApi.md#membershowemail) | **Get** /members/{email}/email | Get a member by email
*MemberApi* | [**MemberStoreUnvalidated**](docs/MemberApi.md#memberstoreunvalidated) | **Post** /members/unvalidated | Subscribe and send validation email
*MemberApi* | [**MemberStoreValidated**](docs/MemberApi.md#memberstorevalidated) | **Post** /members/validated | Subscribe a member without a validation email
*MemberApi* | [**MemberUnsubscribe**](docs/MemberApi.md#memberunsubscribe) | **Post** /members/unsubscribe | Unsubscribe a member from a client or groups
*MemberApi* | [**MemberUpdate**](docs/MemberApi.md#memberupdate) | **Put** /members/{id} | Update a member by ID
*MemberApi* | [**MemberUpdateEmail**](docs/MemberApi.md#memberupdateemail) | **Put** /members/{email}/email | Update a member by email


## Documentation For Models

 - [AutomationIndexResponse](docs/AutomationIndexResponse.md)
 - [AutomationIndexResponseData](docs/AutomationIndexResponseData.md)
 - [AutomationIndexResponseTemplate](docs/AutomationIndexResponseTemplate.md)
 - [AutomationSendRequest](docs/AutomationSendRequest.md)
 - [AutomationSendRequestAttachments](docs/AutomationSendRequestAttachments.md)
 - [AutomationSendRequestRecipients](docs/AutomationSendRequestRecipients.md)
 - [AutomationSendResponse](docs/AutomationSendResponse.md)
 - [AutomationSendResponseData](docs/AutomationSendResponseData.md)
 - [AutomationShowNameResponse](docs/AutomationShowNameResponse.md)
 - [AutomationShowNameResponseData](docs/AutomationShowNameResponseData.md)
 - [AutomationTrackIdResponse](docs/AutomationTrackIdResponse.md)
 - [AutomationTrackIdResponseData](docs/AutomationTrackIdResponseData.md)
 - [AutomationTrackIdResponseDataBounces](docs/AutomationTrackIdResponseDataBounces.md)
 - [AutomationTrackIdResponseDataOpens](docs/AutomationTrackIdResponseDataOpens.md)
 - [AutomationTrackRefResponse](docs/AutomationTrackRefResponse.md)
 - [DefaultValidationErrorResponse](docs/DefaultValidationErrorResponse.md)
 - [EventGenericRequest](docs/EventGenericRequest.md)
 - [EventOrderRequest](docs/EventOrderRequest.md)
 - [EventOrderRequestPayload](docs/EventOrderRequestPayload.md)
 - [EventOrderTrackAndTraceRequest](docs/EventOrderTrackAndTraceRequest.md)
 - [EventOrderTrackAndTraceRequestPayload](docs/EventOrderTrackAndTraceRequestPayload.md)
 - [EventTriggerResponse](docs/EventTriggerResponse.md)
 - [FlowConfigRequestDynamicGroup](docs/FlowConfigRequestDynamicGroup.md)
 - [FlowConfigRequestGroup](docs/FlowConfigRequestGroup.md)
 - [FlowConfigTriggerRequest](docs/FlowConfigTriggerRequest.md)
 - [FlowConfigTriggerResponse](docs/FlowConfigTriggerResponse.md)
 - [GroupImportRequest](docs/GroupImportRequest.md)
 - [GroupImportRequestMembers](docs/GroupImportRequestMembers.md)
 - [GroupImportRequestMetaFields](docs/GroupImportRequestMetaFields.md)
 - [GroupIndexResponse](docs/GroupIndexResponse.md)
 - [GroupIndexResponseData](docs/GroupIndexResponseData.md)
 - [GroupMemberIndexResponse](docs/GroupMemberIndexResponse.md)
 - [GroupMemberIndexResponseData](docs/GroupMemberIndexResponseData.md)
 - [GroupMemberIndexResponseMetaFields](docs/GroupMemberIndexResponseMetaFields.md)
 - [MemberRequestGroup](docs/MemberRequestGroup.md)
 - [MemberRequestMetaField](docs/MemberRequestMetaField.md)
 - [MemberResponse](docs/MemberResponse.md)
 - [MemberResponseData](docs/MemberResponseData.md)
 - [MemberStoreRequest](docs/MemberStoreRequest.md)
 - [MemberStoreRequestAutomation](docs/MemberStoreRequestAutomation.md)
 - [MemberStoreResponse](docs/MemberStoreResponse.md)
 - [MemberUnsubscribeRequest](docs/MemberUnsubscribeRequest.md)
 - [MemberUnsubscribeResponse](docs/MemberUnsubscribeResponse.md)
 - [MemberUpdateRequest](docs/MemberUpdateRequest.md)
 - [MemberUpdateResponse](docs/MemberUpdateResponse.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [ReplacementAddress](docs/ReplacementAddress.md)
 - [ReplacementCoupon](docs/ReplacementCoupon.md)
 - [ReplacementImage](docs/ReplacementImage.md)
 - [ReplacementOrder](docs/ReplacementOrder.md)
 - [ReplacementOrderLine](docs/ReplacementOrderLine.md)
 - [ResponseAutomation](docs/ResponseAutomation.md)
 - [ResponseAutomationReplacements](docs/ResponseAutomationReplacements.md)
 - [ResponseBounce](docs/ResponseBounce.md)
 - [ResponseCampaign](docs/ResponseCampaign.md)
 - [ResponseClick](docs/ResponseClick.md)
 - [ResponseClient](docs/ResponseClient.md)
 - [ResponseGroup](docs/ResponseGroup.md)
 - [ResponseMail](docs/ResponseMail.md)
 - [ResponseMember](docs/ResponseMember.md)
 - [ResponseMetaField](docs/ResponseMetaField.md)
 - [ResponseOpen](docs/ResponseOpen.md)
 - [ResponseRecipient](docs/ResponseRecipient.md)
 - [ResponseReplacement](docs/ResponseReplacement.md)
 - [ResponseTemplate](docs/ResponseTemplate.md)


## Documentation For Authorization



## BearerAuth

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```



## Author

support@slash2.nl

