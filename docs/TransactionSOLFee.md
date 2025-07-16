# TransactionSOLFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseFee** | Pointer to **string** | A fixed fee charged per signature. The default is 5,000 lamports per signature. | [optional] 
**RentAmount** | Pointer to **string** | The rent fee charged by the network to store non–rent-exempt accounts on-chain. It is deducted periodically until the account maintains the minimum balance required for rent exemption. | [optional] 
**ComputeUnitPrice** | Pointer to **string** | The price paid per compute unit. This value determines the priority fee for the transaction, allowing you to increase inclusion probability in congested conditions. | [optional] 
**ComputeUnitLimit** | Pointer to **string** | The maximum number of compute units your transaction is allowed to consume. It sets an upper bound on computational resource usage to prevent overload. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | Pointer to **string** | The token used to pay the transaction fee. | [optional] 
**FeeUsed** | Pointer to **string** | The actually charged transaction fee. | [optional] 
**EstimatedFeeUsed** | Pointer to **string** | The estimated transaction fee. | [optional] 

## Methods

### NewTransactionSOLFee

`func NewTransactionSOLFee(feeType FeeType, ) *TransactionSOLFee`

NewTransactionSOLFee instantiates a new TransactionSOLFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSOLFeeWithDefaults

`func NewTransactionSOLFeeWithDefaults() *TransactionSOLFee`

NewTransactionSOLFeeWithDefaults instantiates a new TransactionSOLFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseFee

`func (o *TransactionSOLFee) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *TransactionSOLFee) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *TransactionSOLFee) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.

### HasBaseFee

`func (o *TransactionSOLFee) HasBaseFee() bool`

HasBaseFee returns a boolean if a field has been set.

### GetRentAmount

`func (o *TransactionSOLFee) GetRentAmount() string`

GetRentAmount returns the RentAmount field if non-nil, zero value otherwise.

### GetRentAmountOk

`func (o *TransactionSOLFee) GetRentAmountOk() (*string, bool)`

GetRentAmountOk returns a tuple with the RentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentAmount

`func (o *TransactionSOLFee) SetRentAmount(v string)`

SetRentAmount sets RentAmount field to given value.

### HasRentAmount

`func (o *TransactionSOLFee) HasRentAmount() bool`

HasRentAmount returns a boolean if a field has been set.

### GetComputeUnitPrice

`func (o *TransactionSOLFee) GetComputeUnitPrice() string`

GetComputeUnitPrice returns the ComputeUnitPrice field if non-nil, zero value otherwise.

### GetComputeUnitPriceOk

`func (o *TransactionSOLFee) GetComputeUnitPriceOk() (*string, bool)`

GetComputeUnitPriceOk returns a tuple with the ComputeUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitPrice

`func (o *TransactionSOLFee) SetComputeUnitPrice(v string)`

SetComputeUnitPrice sets ComputeUnitPrice field to given value.

### HasComputeUnitPrice

`func (o *TransactionSOLFee) HasComputeUnitPrice() bool`

HasComputeUnitPrice returns a boolean if a field has been set.

### GetComputeUnitLimit

`func (o *TransactionSOLFee) GetComputeUnitLimit() string`

GetComputeUnitLimit returns the ComputeUnitLimit field if non-nil, zero value otherwise.

### GetComputeUnitLimitOk

`func (o *TransactionSOLFee) GetComputeUnitLimitOk() (*string, bool)`

GetComputeUnitLimitOk returns a tuple with the ComputeUnitLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitLimit

`func (o *TransactionSOLFee) SetComputeUnitLimit(v string)`

SetComputeUnitLimit sets ComputeUnitLimit field to given value.

### HasComputeUnitLimit

`func (o *TransactionSOLFee) HasComputeUnitLimit() bool`

HasComputeUnitLimit returns a boolean if a field has been set.

### GetFeeType

`func (o *TransactionSOLFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionSOLFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionSOLFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *TransactionSOLFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionSOLFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionSOLFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TransactionSOLFee) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetFeeUsed

`func (o *TransactionSOLFee) GetFeeUsed() string`

GetFeeUsed returns the FeeUsed field if non-nil, zero value otherwise.

### GetFeeUsedOk

`func (o *TransactionSOLFee) GetFeeUsedOk() (*string, bool)`

GetFeeUsedOk returns a tuple with the FeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeUsed

`func (o *TransactionSOLFee) SetFeeUsed(v string)`

SetFeeUsed sets FeeUsed field to given value.

### HasFeeUsed

`func (o *TransactionSOLFee) HasFeeUsed() bool`

HasFeeUsed returns a boolean if a field has been set.

### GetEstimatedFeeUsed

`func (o *TransactionSOLFee) GetEstimatedFeeUsed() string`

GetEstimatedFeeUsed returns the EstimatedFeeUsed field if non-nil, zero value otherwise.

### GetEstimatedFeeUsedOk

`func (o *TransactionSOLFee) GetEstimatedFeeUsedOk() (*string, bool)`

GetEstimatedFeeUsedOk returns a tuple with the EstimatedFeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFeeUsed

`func (o *TransactionSOLFee) SetEstimatedFeeUsed(v string)`

SetEstimatedFeeUsed sets EstimatedFeeUsed field to given value.

### HasEstimatedFeeUsed

`func (o *TransactionSOLFee) HasEstimatedFeeUsed() bool`

HasEstimatedFeeUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


