# EvmEip1559FeeBasePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | ID of the fee token. Unique in all chains scope. | [optional] 
**MaxFee** | **string** | The highest Gas price paid for the transfer, unit GWei. | 
**MaxPriorityFee** | **int32** | The maximum Gas price paid to miners, the higher it is, the faster it is likely to be packaged into the block, unit GWei. | 
**BaseFee** | **int32** | The Base Fee of chain. | 

## Methods

### NewEvmEip1559FeeBasePrice

`func NewEvmEip1559FeeBasePrice(maxFee string, maxPriorityFee int32, baseFee int32, ) *EvmEip1559FeeBasePrice`

NewEvmEip1559FeeBasePrice instantiates a new EvmEip1559FeeBasePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEip1559FeeBasePriceWithDefaults

`func NewEvmEip1559FeeBasePriceWithDefaults() *EvmEip1559FeeBasePrice`

NewEvmEip1559FeeBasePriceWithDefaults instantiates a new EvmEip1559FeeBasePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *EvmEip1559FeeBasePrice) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *EvmEip1559FeeBasePrice) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *EvmEip1559FeeBasePrice) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *EvmEip1559FeeBasePrice) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetMaxFee

`func (o *EvmEip1559FeeBasePrice) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *EvmEip1559FeeBasePrice) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *EvmEip1559FeeBasePrice) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.


### GetMaxPriorityFee

`func (o *EvmEip1559FeeBasePrice) GetMaxPriorityFee() int32`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *EvmEip1559FeeBasePrice) GetMaxPriorityFeeOk() (*int32, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *EvmEip1559FeeBasePrice) SetMaxPriorityFee(v int32)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *EvmEip1559FeeBasePrice) GetBaseFee() int32`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *EvmEip1559FeeBasePrice) GetBaseFeeOk() (*int32, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *EvmEip1559FeeBasePrice) SetBaseFee(v int32)`

SetBaseFee sets BaseFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


