# TSSSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bip32Path** | Pointer to **string** | The BIP32 path. | [optional] 
**MsgHash** | Pointer to **string** | The message hash. | [optional] 
**Tweak** | Pointer to **string** | The tweak. | [optional] 
**Signature** | Pointer to **string** | The signature. | [optional] 
**SignatureRecovery** | Pointer to **string** | The signature recovery. | [optional] 

## Methods

### NewTSSSignature

`func NewTSSSignature() *TSSSignature`

NewTSSSignature instantiates a new TSSSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSSignatureWithDefaults

`func NewTSSSignatureWithDefaults() *TSSSignature`

NewTSSSignatureWithDefaults instantiates a new TSSSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBip32Path

`func (o *TSSSignature) GetBip32Path() string`

GetBip32Path returns the Bip32Path field if non-nil, zero value otherwise.

### GetBip32PathOk

`func (o *TSSSignature) GetBip32PathOk() (*string, bool)`

GetBip32PathOk returns a tuple with the Bip32Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBip32Path

`func (o *TSSSignature) SetBip32Path(v string)`

SetBip32Path sets Bip32Path field to given value.

### HasBip32Path

`func (o *TSSSignature) HasBip32Path() bool`

HasBip32Path returns a boolean if a field has been set.

### GetMsgHash

`func (o *TSSSignature) GetMsgHash() string`

GetMsgHash returns the MsgHash field if non-nil, zero value otherwise.

### GetMsgHashOk

`func (o *TSSSignature) GetMsgHashOk() (*string, bool)`

GetMsgHashOk returns a tuple with the MsgHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgHash

`func (o *TSSSignature) SetMsgHash(v string)`

SetMsgHash sets MsgHash field to given value.

### HasMsgHash

`func (o *TSSSignature) HasMsgHash() bool`

HasMsgHash returns a boolean if a field has been set.

### GetTweak

`func (o *TSSSignature) GetTweak() string`

GetTweak returns the Tweak field if non-nil, zero value otherwise.

### GetTweakOk

`func (o *TSSSignature) GetTweakOk() (*string, bool)`

GetTweakOk returns a tuple with the Tweak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweak

`func (o *TSSSignature) SetTweak(v string)`

SetTweak sets Tweak field to given value.

### HasTweak

`func (o *TSSSignature) HasTweak() bool`

HasTweak returns a boolean if a field has been set.

### GetSignature

`func (o *TSSSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TSSSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TSSSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *TSSSignature) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignatureRecovery

`func (o *TSSSignature) GetSignatureRecovery() string`

GetSignatureRecovery returns the SignatureRecovery field if non-nil, zero value otherwise.

### GetSignatureRecoveryOk

`func (o *TSSSignature) GetSignatureRecoveryOk() (*string, bool)`

GetSignatureRecoveryOk returns a tuple with the SignatureRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRecovery

`func (o *TSSSignature) SetSignatureRecovery(v string)`

SetSignatureRecovery sets SignatureRecovery field to given value.

### HasSignatureRecovery

`func (o *TSSSignature) HasSignatureRecovery() bool`

HasSignatureRecovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


