# WebhookEventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The event log ID. | 
**CreatedTimestamp** | **int32** | The time when the log was created, in Unix timestamp format, measured in milliseconds. | 
**RequestHeaders** | **map[string]interface{}** | The request headers of the webhook event. | 
**RequestBody** | [**WebhookEvent**](WebhookEvent.md) |  | 
**ResponseBody** | Pointer to **map[string]interface{}** | The response body of the webhook event. | [optional] 
**ResponseStatusCode** | Pointer to **int32** | The response status code of the webhook event. | [optional] 
**ResponseTime** | Pointer to **int32** | The response time of the webhook event, in milliseconds. | [optional] 
**Success** | **bool** | Whether the webhook event has been successfully delivered. | 
**FailureReason** | Pointer to **string** | The reason why the webhook event fails to be delivered. | [optional] 

## Methods

### NewWebhookEventLog

`func NewWebhookEventLog(id string, createdTimestamp int32, requestHeaders map[string]interface{}, requestBody WebhookEvent, success bool, ) *WebhookEventLog`

NewWebhookEventLog instantiates a new WebhookEventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventLogWithDefaults

`func NewWebhookEventLogWithDefaults() *WebhookEventLog`

NewWebhookEventLogWithDefaults instantiates a new WebhookEventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEventLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEventLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEventLog) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedTimestamp

`func (o *WebhookEventLog) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *WebhookEventLog) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *WebhookEventLog) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetRequestHeaders

`func (o *WebhookEventLog) GetRequestHeaders() map[string]interface{}`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *WebhookEventLog) GetRequestHeadersOk() (*map[string]interface{}, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *WebhookEventLog) SetRequestHeaders(v map[string]interface{})`

SetRequestHeaders sets RequestHeaders field to given value.


### GetRequestBody

`func (o *WebhookEventLog) GetRequestBody() WebhookEvent`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *WebhookEventLog) GetRequestBodyOk() (*WebhookEvent, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *WebhookEventLog) SetRequestBody(v WebhookEvent)`

SetRequestBody sets RequestBody field to given value.


### GetResponseBody

`func (o *WebhookEventLog) GetResponseBody() map[string]interface{}`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *WebhookEventLog) GetResponseBodyOk() (*map[string]interface{}, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *WebhookEventLog) SetResponseBody(v map[string]interface{})`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *WebhookEventLog) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *WebhookEventLog) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *WebhookEventLog) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *WebhookEventLog) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *WebhookEventLog) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetResponseTime

`func (o *WebhookEventLog) GetResponseTime() int32`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *WebhookEventLog) GetResponseTimeOk() (*int32, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *WebhookEventLog) SetResponseTime(v int32)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *WebhookEventLog) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetSuccess

`func (o *WebhookEventLog) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WebhookEventLog) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WebhookEventLog) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetFailureReason

`func (o *WebhookEventLog) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *WebhookEventLog) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *WebhookEventLog) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *WebhookEventLog) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


