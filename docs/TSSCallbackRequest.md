# TSSCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | he unique ID of the callback request. | [optional] 
**RequestType** | Pointer to [**TSSCallbackRequestType**](TSSCallbackRequestType.md) |  | [optional] 
**RequestDetail** | Pointer to **string** | Details specific to the request type. The structure varies depending on the request type.  The content is a JSON-serialized string. | [optional] 
**ExtraInfo** | Pointer to **string** | Additional contextual information.  The structure varies depending on the request type. The content is a JSON-serialized string. | [optional] 

## Methods

### NewTSSCallbackRequest

`func NewTSSCallbackRequest() *TSSCallbackRequest`

NewTSSCallbackRequest instantiates a new TSSCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSCallbackRequestWithDefaults

`func NewTSSCallbackRequestWithDefaults() *TSSCallbackRequest`

NewTSSCallbackRequestWithDefaults instantiates a new TSSCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TSSCallbackRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TSSCallbackRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TSSCallbackRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TSSCallbackRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestType

`func (o *TSSCallbackRequest) GetRequestType() TSSCallbackRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TSSCallbackRequest) GetRequestTypeOk() (*TSSCallbackRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TSSCallbackRequest) SetRequestType(v TSSCallbackRequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TSSCallbackRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRequestDetail

`func (o *TSSCallbackRequest) GetRequestDetail() string`

GetRequestDetail returns the RequestDetail field if non-nil, zero value otherwise.

### GetRequestDetailOk

`func (o *TSSCallbackRequest) GetRequestDetailOk() (*string, bool)`

GetRequestDetailOk returns a tuple with the RequestDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetail

`func (o *TSSCallbackRequest) SetRequestDetail(v string)`

SetRequestDetail sets RequestDetail field to given value.

### HasRequestDetail

`func (o *TSSCallbackRequest) HasRequestDetail() bool`

HasRequestDetail returns a boolean if a field has been set.

### GetExtraInfo

`func (o *TSSCallbackRequest) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *TSSCallbackRequest) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *TSSCallbackRequest) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *TSSCallbackRequest) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


