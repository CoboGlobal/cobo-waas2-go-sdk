# EvmEip1559FeePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | The token ID of the transaction fee. Unique in all chains scope. | [optional] 
**MaxPriorityFee** | **string** | The max priority fee, in gwei. The max priority fee represents the highest amount of miner tips you are willing to pay for your transaction. | 
**BaseFee** | **string** | The base fee of chain. | 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]

## Methods

### NewEvmEip1559FeePrice

`func NewEvmEip1559FeePrice(maxPriorityFee string, baseFee string, feeType FeeType, ) *EvmEip1559FeePrice`

NewEvmEip1559FeePrice instantiates a new EvmEip1559FeePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEip1559FeePriceWithDefaults

`func NewEvmEip1559FeePriceWithDefaults() *EvmEip1559FeePrice`

NewEvmEip1559FeePriceWithDefaults instantiates a new EvmEip1559FeePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *EvmEip1559FeePrice) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *EvmEip1559FeePrice) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *EvmEip1559FeePrice) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *EvmEip1559FeePrice) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetMaxPriorityFee

`func (o *EvmEip1559FeePrice) GetMaxPriorityFee() string`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *EvmEip1559FeePrice) GetMaxPriorityFeeOk() (*string, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *EvmEip1559FeePrice) SetMaxPriorityFee(v string)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *EvmEip1559FeePrice) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *EvmEip1559FeePrice) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *EvmEip1559FeePrice) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.


### GetFeeType

`func (o *EvmEip1559FeePrice) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmEip1559FeePrice) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmEip1559FeePrice) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


