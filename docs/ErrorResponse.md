# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the API operation was successful. Always false for errors. | [default to false]
**ErrorCode** | **int32** | A machine-readable error code.&#x60; | 
**ErrorDescription** | **string** | A human-readable error description for users. | 
**ErrorId** | **string** | A unique ID for the error log, mainly used for debugging. | 

## Methods

### NewErrorResponse

`func NewErrorResponse(success bool, errorCode int32, errorDescription string, errorId string, ) *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ErrorResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ErrorResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ErrorResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrorCode

`func (o *ErrorResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorDescription

`func (o *ErrorResponse) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ErrorResponse) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ErrorResponse) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.


### GetErrorId

`func (o *ErrorResponse) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ErrorResponse) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ErrorResponse) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


