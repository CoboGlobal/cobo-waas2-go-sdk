# SwapActivitySigners

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signer** | Pointer to **string** | The signer name. | [optional] 
**Status** | Pointer to [**SwapSingingStatus**](SwapSingingStatus.md) |  | [optional] 
**FailedReason** | Pointer to **string** | Failed reason of signing process. | [optional] 

## Methods

### NewSwapActivitySigners

`func NewSwapActivitySigners() *SwapActivitySigners`

NewSwapActivitySigners instantiates a new SwapActivitySigners object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapActivitySignersWithDefaults

`func NewSwapActivitySignersWithDefaults() *SwapActivitySigners`

NewSwapActivitySignersWithDefaults instantiates a new SwapActivitySigners object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigner

`func (o *SwapActivitySigners) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *SwapActivitySigners) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *SwapActivitySigners) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *SwapActivitySigners) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetStatus

`func (o *SwapActivitySigners) GetStatus() SwapSingingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwapActivitySigners) GetStatusOk() (*SwapSingingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwapActivitySigners) SetStatus(v SwapSingingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SwapActivitySigners) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailedReason

`func (o *SwapActivitySigners) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *SwapActivitySigners) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *SwapActivitySigners) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *SwapActivitySigners) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


