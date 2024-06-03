# UtxoTransactionFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | The token ID of the transaction fee. Unique in all chains scope. | [optional] 
**FeeRate** | **string** | The fee rate, in sats/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain. | 
**MaxFeeAmount** | Pointer to **string** | The max fee amount in fee_coin. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]

## Methods

### NewUtxoTransactionFee

`func NewUtxoTransactionFee(feeRate string, feeType FeeType, ) *UtxoTransactionFee`

NewUtxoTransactionFee instantiates a new UtxoTransactionFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoTransactionFeeWithDefaults

`func NewUtxoTransactionFeeWithDefaults() *UtxoTransactionFee`

NewUtxoTransactionFeeWithDefaults instantiates a new UtxoTransactionFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *UtxoTransactionFee) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *UtxoTransactionFee) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *UtxoTransactionFee) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *UtxoTransactionFee) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetFeeRate

`func (o *UtxoTransactionFee) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *UtxoTransactionFee) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *UtxoTransactionFee) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetMaxFeeAmount

`func (o *UtxoTransactionFee) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *UtxoTransactionFee) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *UtxoTransactionFee) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.

### HasMaxFeeAmount

`func (o *UtxoTransactionFee) HasMaxFeeAmount() bool`

HasMaxFeeAmount returns a boolean if a field has been set.

### GetFeeType

`func (o *UtxoTransactionFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *UtxoTransactionFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *UtxoTransactionFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


