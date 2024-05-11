# UtxoFeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | ID of the fee token. Unique in all chains scope. | [optional] 
**FeeRate** | **string** | The fee rate, unit sat/vB. | 
**FeeAmount** | Pointer to **string** | The estimated fee amount in fee_coin. | [optional] 

## Methods

### NewUtxoFeeSlow

`func NewUtxoFeeSlow(feeRate string, ) *UtxoFeeSlow`

NewUtxoFeeSlow instantiates a new UtxoFeeSlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoFeeSlowWithDefaults

`func NewUtxoFeeSlowWithDefaults() *UtxoFeeSlow`

NewUtxoFeeSlowWithDefaults instantiates a new UtxoFeeSlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *UtxoFeeSlow) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *UtxoFeeSlow) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *UtxoFeeSlow) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *UtxoFeeSlow) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetFeeRate

`func (o *UtxoFeeSlow) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *UtxoFeeSlow) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *UtxoFeeSlow) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetFeeAmount

`func (o *UtxoFeeSlow) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *UtxoFeeSlow) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *UtxoFeeSlow) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *UtxoFeeSlow) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


