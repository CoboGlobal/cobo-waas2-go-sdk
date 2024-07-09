# GetStakingEstimationFeeRequest

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
**StakingId** | **string** | The id of the related staking. | 

## Methods

### NewGetStakingEstimationFeeRequest

`func NewGetStakingEstimationFeeRequest(walletId string, address string, poolId string, amount string, fee TransactionTransferFee, extra CreateStakeActivityExtra, stakingId string, ) *GetStakingEstimationFeeRequest`

NewGetStakingEstimationFeeRequest instantiates a new GetStakingEstimationFeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStakingEstimationFeeRequestWithDefaults

`func NewGetStakingEstimationFeeRequestWithDefaults() *GetStakingEstimationFeeRequest`

NewGetStakingEstimationFeeRequestWithDefaults instantiates a new GetStakingEstimationFeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *GetStakingEstimationFeeRequest) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *GetStakingEstimationFeeRequest) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *GetStakingEstimationFeeRequest) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *GetStakingEstimationFeeRequest) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetWalletId

`func (o *GetStakingEstimationFeeRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *GetStakingEstimationFeeRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *GetStakingEstimationFeeRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *GetStakingEstimationFeeRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetStakingEstimationFeeRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetStakingEstimationFeeRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPoolId

`func (o *GetStakingEstimationFeeRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GetStakingEstimationFeeRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GetStakingEstimationFeeRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAmount

`func (o *GetStakingEstimationFeeRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetStakingEstimationFeeRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetStakingEstimationFeeRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *GetStakingEstimationFeeRequest) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetStakingEstimationFeeRequest) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetStakingEstimationFeeRequest) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.


### GetExtra

`func (o *GetStakingEstimationFeeRequest) GetExtra() CreateStakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *GetStakingEstimationFeeRequest) GetExtraOk() (*CreateStakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *GetStakingEstimationFeeRequest) SetExtra(v CreateStakeActivityExtra)`

SetExtra sets Extra field to given value.


### GetStakingId

`func (o *GetStakingEstimationFeeRequest) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *GetStakingEstimationFeeRequest) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *GetStakingEstimationFeeRequest) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


