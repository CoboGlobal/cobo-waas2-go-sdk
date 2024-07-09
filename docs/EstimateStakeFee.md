# EstimateStakeFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**WalletId** | **string** | The id of the wallet to stake. | 
**Address** | **string** | The staker wallet address. | 
**PoolId** | **string** | The id of the staking pool | 
**Amount** | **string** | The amount to stake | 
**Fee** | [**TransactionTransferFee**](TransactionTransferFee.md) |  | 
**Extra** | [**CreateStakeActivityExtra**](CreateStakeActivityExtra.md) |  | 

## Methods

### NewEstimateStakeFee

`func NewEstimateStakeFee(walletId string, address string, poolId string, amount string, fee TransactionTransferFee, extra CreateStakeActivityExtra, ) *EstimateStakeFee`

NewEstimateStakeFee instantiates a new EstimateStakeFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateStakeFeeWithDefaults

`func NewEstimateStakeFeeWithDefaults() *EstimateStakeFee`

NewEstimateStakeFeeWithDefaults instantiates a new EstimateStakeFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *EstimateStakeFee) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *EstimateStakeFee) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *EstimateStakeFee) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *EstimateStakeFee) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetWalletId

`func (o *EstimateStakeFee) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *EstimateStakeFee) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *EstimateStakeFee) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *EstimateStakeFee) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EstimateStakeFee) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EstimateStakeFee) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPoolId

`func (o *EstimateStakeFee) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *EstimateStakeFee) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *EstimateStakeFee) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAmount

`func (o *EstimateStakeFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EstimateStakeFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EstimateStakeFee) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *EstimateStakeFee) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EstimateStakeFee) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EstimateStakeFee) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.


### GetExtra

`func (o *EstimateStakeFee) GetExtra() CreateStakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *EstimateStakeFee) GetExtraOk() (*CreateStakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *EstimateStakeFee) SetExtra(v CreateStakeActivityExtra)`

SetExtra sets Extra field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


