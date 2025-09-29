# EstimatedSOLFeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeUnitPrice** | **string** | The cost per compute unit. Transactions consume computational resources measured in compute units, and this price helps determine the cost of executing transactions, especially complex ones involving smart contracts. | 
**ComputeUnitLimit** | **string** | The maximum number of compute units allowed for a transaction. This limits the resources any single transaction can consume, preventing excessive resource usage that could impact network performance negatively. | 
**BaseFee** | **string** | The fundamental fee required for each transaction. It is charged to prevent spam transactions and network congestion, ensuring that only meaningful transactions consume network resources. | 
**RentAmount** | Pointer to **string** | The fee charged as rent for maintaining the state of accounts on the blockchain. This rent ensures accounts are stored on-chain over the long term and that there&#39;s sufficient balance to sustain the account state. | [optional] 

## Methods

### NewEstimatedSOLFeeSlow

`func NewEstimatedSOLFeeSlow(computeUnitPrice string, computeUnitLimit string, baseFee string, ) *EstimatedSOLFeeSlow`

NewEstimatedSOLFeeSlow instantiates a new EstimatedSOLFeeSlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedSOLFeeSlowWithDefaults

`func NewEstimatedSOLFeeSlowWithDefaults() *EstimatedSOLFeeSlow`

NewEstimatedSOLFeeSlowWithDefaults instantiates a new EstimatedSOLFeeSlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeUnitPrice

`func (o *EstimatedSOLFeeSlow) GetComputeUnitPrice() string`

GetComputeUnitPrice returns the ComputeUnitPrice field if non-nil, zero value otherwise.

### GetComputeUnitPriceOk

`func (o *EstimatedSOLFeeSlow) GetComputeUnitPriceOk() (*string, bool)`

GetComputeUnitPriceOk returns a tuple with the ComputeUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitPrice

`func (o *EstimatedSOLFeeSlow) SetComputeUnitPrice(v string)`

SetComputeUnitPrice sets ComputeUnitPrice field to given value.


### GetComputeUnitLimit

`func (o *EstimatedSOLFeeSlow) GetComputeUnitLimit() string`

GetComputeUnitLimit returns the ComputeUnitLimit field if non-nil, zero value otherwise.

### GetComputeUnitLimitOk

`func (o *EstimatedSOLFeeSlow) GetComputeUnitLimitOk() (*string, bool)`

GetComputeUnitLimitOk returns a tuple with the ComputeUnitLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitLimit

`func (o *EstimatedSOLFeeSlow) SetComputeUnitLimit(v string)`

SetComputeUnitLimit sets ComputeUnitLimit field to given value.


### GetBaseFee

`func (o *EstimatedSOLFeeSlow) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *EstimatedSOLFeeSlow) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *EstimatedSOLFeeSlow) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.


### GetRentAmount

`func (o *EstimatedSOLFeeSlow) GetRentAmount() string`

GetRentAmount returns the RentAmount field if non-nil, zero value otherwise.

### GetRentAmountOk

`func (o *EstimatedSOLFeeSlow) GetRentAmountOk() (*string, bool)`

GetRentAmountOk returns a tuple with the RentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentAmount

`func (o *EstimatedSOLFeeSlow) SetRentAmount(v string)`

SetRentAmount sets RentAmount field to given value.

### HasRentAmount

`func (o *EstimatedSOLFeeSlow) HasRentAmount() bool`

HasRentAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


