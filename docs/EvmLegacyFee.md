# EvmLegacyFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**Slow** | Pointer to [**EvmLegacyFeeSlow**](EvmLegacyFeeSlow.md) |  | [optional] 
**Standard** | [**EvmLegacyFeeSlow**](EvmLegacyFeeSlow.md) |  | 
**Fast** | Pointer to [**EvmLegacyFeeSlow**](EvmLegacyFeeSlow.md) |  | [optional] 

## Methods

### NewEvmLegacyFee

`func NewEvmLegacyFee(feeType FeeType, standard EvmLegacyFeeSlow, ) *EvmLegacyFee`

NewEvmLegacyFee instantiates a new EvmLegacyFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmLegacyFeeWithDefaults

`func NewEvmLegacyFeeWithDefaults() *EvmLegacyFee`

NewEvmLegacyFeeWithDefaults instantiates a new EvmLegacyFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EvmLegacyFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmLegacyFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmLegacyFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetSlow

`func (o *EvmLegacyFee) GetSlow() EvmLegacyFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EvmLegacyFee) GetSlowOk() (*EvmLegacyFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EvmLegacyFee) SetSlow(v EvmLegacyFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EvmLegacyFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetStandard

`func (o *EvmLegacyFee) GetStandard() EvmLegacyFeeSlow`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *EvmLegacyFee) GetStandardOk() (*EvmLegacyFeeSlow, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *EvmLegacyFee) SetStandard(v EvmLegacyFeeSlow)`

SetStandard sets Standard field to given value.


### GetFast

`func (o *EvmLegacyFee) GetFast() EvmLegacyFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EvmLegacyFee) GetFastOk() (*EvmLegacyFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EvmLegacyFee) SetFast(v EvmLegacyFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EvmLegacyFee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


