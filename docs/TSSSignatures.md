# TSSSignatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signatures** | Pointer to [**[]TSSSignature**](TSSSignature.md) |  | [optional] 
**SignatureType** | Pointer to [**TSSSignatureType**](TSSSignatureType.md) |  | [optional] 
**TssProtocol** | Pointer to [**TSSProtocol**](TSSProtocol.md) |  | [optional] 

## Methods

### NewTSSSignatures

`func NewTSSSignatures() *TSSSignatures`

NewTSSSignatures instantiates a new TSSSignatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSSignaturesWithDefaults

`func NewTSSSignaturesWithDefaults() *TSSSignatures`

NewTSSSignaturesWithDefaults instantiates a new TSSSignatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatures

`func (o *TSSSignatures) GetSignatures() []TSSSignature`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *TSSSignatures) GetSignaturesOk() (*[]TSSSignature, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *TSSSignatures) SetSignatures(v []TSSSignature)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *TSSSignatures) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.

### GetSignatureType

`func (o *TSSSignatures) GetSignatureType() TSSSignatureType`

GetSignatureType returns the SignatureType field if non-nil, zero value otherwise.

### GetSignatureTypeOk

`func (o *TSSSignatures) GetSignatureTypeOk() (*TSSSignatureType, bool)`

GetSignatureTypeOk returns a tuple with the SignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureType

`func (o *TSSSignatures) SetSignatureType(v TSSSignatureType)`

SetSignatureType sets SignatureType field to given value.

### HasSignatureType

`func (o *TSSSignatures) HasSignatureType() bool`

HasSignatureType returns a boolean if a field has been set.

### GetTssProtocol

`func (o *TSSSignatures) GetTssProtocol() TSSProtocol`

GetTssProtocol returns the TssProtocol field if non-nil, zero value otherwise.

### GetTssProtocolOk

`func (o *TSSSignatures) GetTssProtocolOk() (*TSSProtocol, bool)`

GetTssProtocolOk returns a tuple with the TssProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssProtocol

`func (o *TSSSignatures) SetTssProtocol(v TSSProtocol)`

SetTssProtocol sets TssProtocol field to given value.

### HasTssProtocol

`func (o *TSSSignatures) HasTssProtocol() bool`

HasTssProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


