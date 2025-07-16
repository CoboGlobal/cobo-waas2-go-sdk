# TransactionRequestUtxoFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeRate** | Pointer to **string** | The fee rate in sat/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token used to pay the transaction fee. | 
**MaxFeeAmount** | Pointer to **string** | The maximum fee that you are willing to pay for the transaction. Provide the value without applying precision. The transaction will fail if the transaction fee exceeds the maximum fee. | [optional] 

## Methods

### NewTransactionRequestUtxoFee

`func NewTransactionRequestUtxoFee(feeType FeeType, tokenId string, ) *TransactionRequestUtxoFee`

NewTransactionRequestUtxoFee instantiates a new TransactionRequestUtxoFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestUtxoFeeWithDefaults

`func NewTransactionRequestUtxoFeeWithDefaults() *TransactionRequestUtxoFee`

NewTransactionRequestUtxoFeeWithDefaults instantiates a new TransactionRequestUtxoFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeRate

`func (o *TransactionRequestUtxoFee) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *TransactionRequestUtxoFee) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *TransactionRequestUtxoFee) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *TransactionRequestUtxoFee) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetFeeType

`func (o *TransactionRequestUtxoFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionRequestUtxoFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionRequestUtxoFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *TransactionRequestUtxoFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionRequestUtxoFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionRequestUtxoFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetMaxFeeAmount

`func (o *TransactionRequestUtxoFee) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *TransactionRequestUtxoFee) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *TransactionRequestUtxoFee) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.

### HasMaxFeeAmount

`func (o *TransactionRequestUtxoFee) HasMaxFeeAmount() bool`

HasMaxFeeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


