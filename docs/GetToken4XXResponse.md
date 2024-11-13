# GetToken4XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | The error name. | 
**ErrorDescription** | **string** | The error description. | 

## Methods

### NewGetToken4XXResponse

`func NewGetToken4XXResponse(error_ string, errorDescription string, ) *GetToken4XXResponse`

NewGetToken4XXResponse instantiates a new GetToken4XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetToken4XXResponseWithDefaults

`func NewGetToken4XXResponseWithDefaults() *GetToken4XXResponse`

NewGetToken4XXResponseWithDefaults instantiates a new GetToken4XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetToken4XXResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetToken4XXResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetToken4XXResponse) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *GetToken4XXResponse) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *GetToken4XXResponse) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *GetToken4XXResponse) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


