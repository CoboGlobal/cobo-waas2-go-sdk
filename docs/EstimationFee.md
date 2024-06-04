# EstimationFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**Slow** | Pointer to [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | [optional] 
**Standard** | [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | 
**Fast** | Pointer to [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | [optional] 
**FeeTokenId** | Pointer to **string** | The token ID of the transaction fee. | [optional] 
**MaxFeeAmount** | Pointer to **string** | The maximum fee amount in fee_coin. | [optional] 

## Methods

### NewEstimationFee

`func NewEstimationFee(feeType FeeType, standard UtxoFeeSlow, ) *EstimationFee`

NewEstimationFee instantiates a new EstimationFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimationFeeWithDefaults

`func NewEstimationFeeWithDefaults() *EstimationFee`

NewEstimationFeeWithDefaults instantiates a new EstimationFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EstimationFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimationFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimationFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetSlow

`func (o *EstimationFee) GetSlow() UtxoFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EstimationFee) GetSlowOk() (*UtxoFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EstimationFee) SetSlow(v UtxoFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EstimationFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetStandard

`func (o *EstimationFee) GetStandard() UtxoFeeSlow`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *EstimationFee) GetStandardOk() (*UtxoFeeSlow, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *EstimationFee) SetStandard(v UtxoFeeSlow)`

SetStandard sets Standard field to given value.


### GetFast

`func (o *EstimationFee) GetFast() UtxoFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EstimationFee) GetFastOk() (*UtxoFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EstimationFee) SetFast(v UtxoFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EstimationFee) HasFast() bool`

HasFast returns a boolean if a field has been set.

### GetFeeTokenId

`func (o *EstimationFee) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *EstimationFee) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *EstimationFee) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *EstimationFee) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetMaxFeeAmount

`func (o *EstimationFee) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *EstimationFee) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *EstimationFee) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.

### HasMaxFeeAmount

`func (o *EstimationFee) HasMaxFeeAmount() bool`

HasMaxFeeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


