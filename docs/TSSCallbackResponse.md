# TSSCallbackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The unique ID of the callback request. | [optional] 
**Status** | Pointer to **int32** | The response status code. 0 indicates success.  Any other value indicates that an error occurred while processing the request in the callback server. | [optional] 
**Action** | Pointer to [**TSSCallbackActionType**](TSSCallbackActionType.md) |  | [optional] 
**Error** | Pointer to **string** | The error message. - When status is not 0, Contains internal error messages from the callback server. - When status is 0 and action is REJECT, Contains the specific reason for the rejection. | [optional] 

## Methods

### NewTSSCallbackResponse

`func NewTSSCallbackResponse() *TSSCallbackResponse`

NewTSSCallbackResponse instantiates a new TSSCallbackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSCallbackResponseWithDefaults

`func NewTSSCallbackResponseWithDefaults() *TSSCallbackResponse`

NewTSSCallbackResponseWithDefaults instantiates a new TSSCallbackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TSSCallbackResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TSSCallbackResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TSSCallbackResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TSSCallbackResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *TSSCallbackResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TSSCallbackResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TSSCallbackResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TSSCallbackResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAction

`func (o *TSSCallbackResponse) GetAction() TSSCallbackActionType`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TSSCallbackResponse) GetActionOk() (*TSSCallbackActionType, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TSSCallbackResponse) SetAction(v TSSCallbackActionType)`

SetAction sets Action field to given value.

### HasAction

`func (o *TSSCallbackResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetError

`func (o *TSSCallbackResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TSSCallbackResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TSSCallbackResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TSSCallbackResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


