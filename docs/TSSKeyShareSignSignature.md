# TSSKeyShareSignSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The TSS key share group ID. | [optional] 
**SignedMsg** | Pointer to **string** | The hexadecimal encoded signed message. | [optional] 
**MsgHash** | Pointer to **string** | The message hash. | [optional] 
**Signature** | Pointer to **string** | The signature. | [optional] 

## Methods

### NewTSSKeyShareSignSignature

`func NewTSSKeyShareSignSignature() *TSSKeyShareSignSignature`

NewTSSKeyShareSignSignature instantiates a new TSSKeyShareSignSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyShareSignSignatureWithDefaults

`func NewTSSKeyShareSignSignatureWithDefaults() *TSSKeyShareSignSignature`

NewTSSKeyShareSignSignatureWithDefaults instantiates a new TSSKeyShareSignSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *TSSKeyShareSignSignature) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TSSKeyShareSignSignature) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TSSKeyShareSignSignature) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *TSSKeyShareSignSignature) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSignedMsg

`func (o *TSSKeyShareSignSignature) GetSignedMsg() string`

GetSignedMsg returns the SignedMsg field if non-nil, zero value otherwise.

### GetSignedMsgOk

`func (o *TSSKeyShareSignSignature) GetSignedMsgOk() (*string, bool)`

GetSignedMsgOk returns a tuple with the SignedMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMsg

`func (o *TSSKeyShareSignSignature) SetSignedMsg(v string)`

SetSignedMsg sets SignedMsg field to given value.

### HasSignedMsg

`func (o *TSSKeyShareSignSignature) HasSignedMsg() bool`

HasSignedMsg returns a boolean if a field has been set.

### GetMsgHash

`func (o *TSSKeyShareSignSignature) GetMsgHash() string`

GetMsgHash returns the MsgHash field if non-nil, zero value otherwise.

### GetMsgHashOk

`func (o *TSSKeyShareSignSignature) GetMsgHashOk() (*string, bool)`

GetMsgHashOk returns a tuple with the MsgHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgHash

`func (o *TSSKeyShareSignSignature) SetMsgHash(v string)`

SetMsgHash sets MsgHash field to given value.

### HasMsgHash

`func (o *TSSKeyShareSignSignature) HasMsgHash() bool`

HasMsgHash returns a boolean if a field has been set.

### GetSignature

`func (o *TSSKeyShareSignSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TSSKeyShareSignSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TSSKeyShareSignSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *TSSKeyShareSignSignature) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


