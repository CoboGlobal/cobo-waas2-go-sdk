# ApiLogDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogId** | **string** | A unique identifier for the API log, used for tracking. | 
**ApiMethod** | **string** | The HTTP method used for the API request. | 
**ApiEndpoint** | **string** | The endpoint of the API request. | 
**StatusCode** | **int32** | The HTTP status code returned by the API request. | 
**IpAddress** | **string** | The client&#39;s IP address that made the API request. | 
**RequestTimestamp** | **int64** | The time when the API request was created, in Unix timestamp format, measured in milliseconds. | 
**ApiKey** | **string** | The API key used to call the API. For more details, refer to [API key](https://www.cobo.com/developers/v2/guides/overview/cobo-auth#api-key). | 
**ResponseBody** | **string** | The response body of the API request. | 
**QueryParams** | **string** | The query parameters of the API request. | 
**RequestBody** | **string** | The request body of the API request. | 

## Methods

### NewApiLogDetails

`func NewApiLogDetails(logId string, apiMethod string, apiEndpoint string, statusCode int32, ipAddress string, requestTimestamp int64, apiKey string, responseBody string, queryParams string, requestBody string, ) *ApiLogDetails`

NewApiLogDetails instantiates a new ApiLogDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLogDetailsWithDefaults

`func NewApiLogDetailsWithDefaults() *ApiLogDetails`

NewApiLogDetailsWithDefaults instantiates a new ApiLogDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogId

`func (o *ApiLogDetails) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *ApiLogDetails) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *ApiLogDetails) SetLogId(v string)`

SetLogId sets LogId field to given value.


### GetApiMethod

`func (o *ApiLogDetails) GetApiMethod() string`

GetApiMethod returns the ApiMethod field if non-nil, zero value otherwise.

### GetApiMethodOk

`func (o *ApiLogDetails) GetApiMethodOk() (*string, bool)`

GetApiMethodOk returns a tuple with the ApiMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMethod

`func (o *ApiLogDetails) SetApiMethod(v string)`

SetApiMethod sets ApiMethod field to given value.


### GetApiEndpoint

`func (o *ApiLogDetails) GetApiEndpoint() string`

GetApiEndpoint returns the ApiEndpoint field if non-nil, zero value otherwise.

### GetApiEndpointOk

`func (o *ApiLogDetails) GetApiEndpointOk() (*string, bool)`

GetApiEndpointOk returns a tuple with the ApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpoint

`func (o *ApiLogDetails) SetApiEndpoint(v string)`

SetApiEndpoint sets ApiEndpoint field to given value.


### GetStatusCode

`func (o *ApiLogDetails) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApiLogDetails) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApiLogDetails) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetIpAddress

`func (o *ApiLogDetails) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApiLogDetails) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApiLogDetails) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetRequestTimestamp

`func (o *ApiLogDetails) GetRequestTimestamp() int64`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *ApiLogDetails) GetRequestTimestampOk() (*int64, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *ApiLogDetails) SetRequestTimestamp(v int64)`

SetRequestTimestamp sets RequestTimestamp field to given value.


### GetApiKey

`func (o *ApiLogDetails) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiLogDetails) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiLogDetails) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetResponseBody

`func (o *ApiLogDetails) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *ApiLogDetails) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *ApiLogDetails) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.


### GetQueryParams

`func (o *ApiLogDetails) GetQueryParams() string`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *ApiLogDetails) GetQueryParamsOk() (*string, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *ApiLogDetails) SetQueryParams(v string)`

SetQueryParams sets QueryParams field to given value.


### GetRequestBody

`func (o *ApiLogDetails) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *ApiLogDetails) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *ApiLogDetails) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


