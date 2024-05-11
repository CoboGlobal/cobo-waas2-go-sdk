# UtxoFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**Slow** | Pointer to [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | [optional] 
**Standard** | [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | 
**Fast** | Pointer to [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | [optional] 

## Methods

### NewUtxoFee

`func NewUtxoFee(feeType FeeType, standard UtxoFeeSlow, ) *UtxoFee`

NewUtxoFee instantiates a new UtxoFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoFeeWithDefaults

`func NewUtxoFeeWithDefaults() *UtxoFee`

NewUtxoFeeWithDefaults instantiates a new UtxoFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *UtxoFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *UtxoFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *UtxoFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetSlow

`func (o *UtxoFee) GetSlow() UtxoFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *UtxoFee) GetSlowOk() (*UtxoFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *UtxoFee) SetSlow(v UtxoFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *UtxoFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetStandard

`func (o *UtxoFee) GetStandard() UtxoFeeSlow`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *UtxoFee) GetStandardOk() (*UtxoFeeSlow, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *UtxoFee) SetStandard(v UtxoFeeSlow)`

SetStandard sets Standard field to given value.


### GetFast

`func (o *UtxoFee) GetFast() UtxoFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *UtxoFee) GetFastOk() (*UtxoFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *UtxoFee) SetFast(v UtxoFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *UtxoFee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


