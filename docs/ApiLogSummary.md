# ApiLogSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogId** | Pointer to **string** | A unique identifier for the API log, used for tracking. | [optional] 
**ApiMethod** | **string** | The HTTP method used for the API request. | 
**ApiEndpoint** | **string** | The endpoint of the API request. | 
**RequestTimestamp** | **int64** | The time when the API request was created, in Unix timestamp format, measured in milliseconds. | 
**StatusCode** | **int32** | The HTTP status code returned by the API request. | 

## Methods

### NewApiLogSummary

`func NewApiLogSummary(apiMethod string, apiEndpoint string, requestTimestamp int64, statusCode int32, ) *ApiLogSummary`

NewApiLogSummary instantiates a new ApiLogSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLogSummaryWithDefaults

`func NewApiLogSummaryWithDefaults() *ApiLogSummary`

NewApiLogSummaryWithDefaults instantiates a new ApiLogSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogId

`func (o *ApiLogSummary) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *ApiLogSummary) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *ApiLogSummary) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *ApiLogSummary) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### GetApiMethod

`func (o *ApiLogSummary) GetApiMethod() string`

GetApiMethod returns the ApiMethod field if non-nil, zero value otherwise.

### GetApiMethodOk

`func (o *ApiLogSummary) GetApiMethodOk() (*string, bool)`

GetApiMethodOk returns a tuple with the ApiMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMethod

`func (o *ApiLogSummary) SetApiMethod(v string)`

SetApiMethod sets ApiMethod field to given value.


### GetApiEndpoint

`func (o *ApiLogSummary) GetApiEndpoint() string`

GetApiEndpoint returns the ApiEndpoint field if non-nil, zero value otherwise.

### GetApiEndpointOk

`func (o *ApiLogSummary) GetApiEndpointOk() (*string, bool)`

GetApiEndpointOk returns a tuple with the ApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpoint

`func (o *ApiLogSummary) SetApiEndpoint(v string)`

SetApiEndpoint sets ApiEndpoint field to given value.


### GetRequestTimestamp

`func (o *ApiLogSummary) GetRequestTimestamp() int64`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *ApiLogSummary) GetRequestTimestampOk() (*int64, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *ApiLogSummary) SetRequestTimestamp(v int64)`

SetRequestTimestamp sets RequestTimestamp field to given value.


### GetStatusCode

`func (o *ApiLogSummary) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApiLogSummary) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApiLogSummary) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


