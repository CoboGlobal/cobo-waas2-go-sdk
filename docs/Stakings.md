# Stakings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the staking position. | 
**WalletId** | **string** | The staker&#39;s wallet ID. | 
**Address** | **string** | The staker&#39;s wallet address. | 
**Amounts** | [**[]AmountDetailsInner**](AmountDetailsInner.md) | The details about the staking amount. | 
**PoolId** | [**StakingPoolId**](StakingPoolId.md) |  | 
**TokenId** | **string** | The token ID. | 
**RewardsInfo** | Pointer to **map[string]interface{}** | The information about the staking rewards. | [optional] 
**CreatedTimestamp** | **int64** | The time when the staking position was created. | 
**UpdatedTimestamp** | **int64** | The time when the staking position was last updated. | 
**ValidatorInfo** | [**BabylonValidator**](BabylonValidator.md) |  | 
**Extra** | Pointer to [**StakingsExtra**](StakingsExtra.md) |  | [optional] 

## Methods

### NewStakings

`func NewStakings(id string, walletId string, address string, amounts []AmountDetailsInner, poolId StakingPoolId, tokenId string, createdTimestamp int64, updatedTimestamp int64, validatorInfo BabylonValidator, ) *Stakings`

NewStakings instantiates a new Stakings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakingsWithDefaults

`func NewStakingsWithDefaults() *Stakings`

NewStakingsWithDefaults instantiates a new Stakings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Stakings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Stakings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Stakings) SetId(v string)`

SetId sets Id field to given value.


### GetWalletId

`func (o *Stakings) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *Stakings) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *Stakings) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *Stakings) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Stakings) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Stakings) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmounts

`func (o *Stakings) GetAmounts() []AmountDetailsInner`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *Stakings) GetAmountsOk() (*[]AmountDetailsInner, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *Stakings) SetAmounts(v []AmountDetailsInner)`

SetAmounts sets Amounts field to given value.


### GetPoolId

`func (o *Stakings) GetPoolId() StakingPoolId`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Stakings) GetPoolIdOk() (*StakingPoolId, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *Stakings) SetPoolId(v StakingPoolId)`

SetPoolId sets PoolId field to given value.


### GetTokenId

`func (o *Stakings) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Stakings) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Stakings) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetRewardsInfo

`func (o *Stakings) GetRewardsInfo() map[string]interface{}`

GetRewardsInfo returns the RewardsInfo field if non-nil, zero value otherwise.

### GetRewardsInfoOk

`func (o *Stakings) GetRewardsInfoOk() (*map[string]interface{}, bool)`

GetRewardsInfoOk returns a tuple with the RewardsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsInfo

`func (o *Stakings) SetRewardsInfo(v map[string]interface{})`

SetRewardsInfo sets RewardsInfo field to given value.

### HasRewardsInfo

`func (o *Stakings) HasRewardsInfo() bool`

HasRewardsInfo returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Stakings) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Stakings) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Stakings) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *Stakings) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Stakings) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Stakings) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetValidatorInfo

`func (o *Stakings) GetValidatorInfo() BabylonValidator`

GetValidatorInfo returns the ValidatorInfo field if non-nil, zero value otherwise.

### GetValidatorInfoOk

`func (o *Stakings) GetValidatorInfoOk() (*BabylonValidator, bool)`

GetValidatorInfoOk returns a tuple with the ValidatorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorInfo

`func (o *Stakings) SetValidatorInfo(v BabylonValidator)`

SetValidatorInfo sets ValidatorInfo field to given value.


### GetExtra

`func (o *Stakings) GetExtra() StakingsExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Stakings) GetExtraOk() (*StakingsExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Stakings) SetExtra(v StakingsExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Stakings) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


