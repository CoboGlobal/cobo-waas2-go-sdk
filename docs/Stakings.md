# Stakings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the stake. | 
**WalletId** | **string** | The unique wallet id. | 
**Address** | **string** | The staker wallet address. | 
**Amounts** | [**[]AmountDetailsInner**](AmountDetailsInner.md) | The staking amount details. | 
**Initiator** | Pointer to **string** | The initiator of the stake. | [optional] 
**UnlockTime** | Pointer to **int64** | The unlock time. | [optional] 
**UnlockBlockHeight** | Pointer to **int64** | The unlock block height. | [optional] 
**PoolId** | **string** | The unique pool id. | 
**TokenId** | **string** | The token id. | 
**PosChain** | Pointer to **string** | The pos chain of the stake. | [optional] 
**RewardsInfo** | Pointer to **map[string]interface{}** | The rewards info of the stake. | [optional] 
**CreatedTime** | **int64** | The time when the stake was created. | 
**UpdatedTime** | **int64** | The time when the stake was last updated. | 
**ValidatorInfo** | [**StakingsValidatorInfo**](StakingsValidatorInfo.md) |  | 

## Methods

### NewStakings

`func NewStakings(id string, walletId string, address string, amounts []AmountDetailsInner, poolId string, tokenId string, createdTime int64, updatedTime int64, validatorInfo StakingsValidatorInfo, ) *Stakings`

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


### GetInitiator

`func (o *Stakings) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *Stakings) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *Stakings) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *Stakings) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetUnlockTime

`func (o *Stakings) GetUnlockTime() int64`

GetUnlockTime returns the UnlockTime field if non-nil, zero value otherwise.

### GetUnlockTimeOk

`func (o *Stakings) GetUnlockTimeOk() (*int64, bool)`

GetUnlockTimeOk returns a tuple with the UnlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockTime

`func (o *Stakings) SetUnlockTime(v int64)`

SetUnlockTime sets UnlockTime field to given value.

### HasUnlockTime

`func (o *Stakings) HasUnlockTime() bool`

HasUnlockTime returns a boolean if a field has been set.

### GetUnlockBlockHeight

`func (o *Stakings) GetUnlockBlockHeight() int64`

GetUnlockBlockHeight returns the UnlockBlockHeight field if non-nil, zero value otherwise.

### GetUnlockBlockHeightOk

`func (o *Stakings) GetUnlockBlockHeightOk() (*int64, bool)`

GetUnlockBlockHeightOk returns a tuple with the UnlockBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockBlockHeight

`func (o *Stakings) SetUnlockBlockHeight(v int64)`

SetUnlockBlockHeight sets UnlockBlockHeight field to given value.

### HasUnlockBlockHeight

`func (o *Stakings) HasUnlockBlockHeight() bool`

HasUnlockBlockHeight returns a boolean if a field has been set.

### GetPoolId

`func (o *Stakings) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Stakings) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *Stakings) SetPoolId(v string)`

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


### GetPosChain

`func (o *Stakings) GetPosChain() string`

GetPosChain returns the PosChain field if non-nil, zero value otherwise.

### GetPosChainOk

`func (o *Stakings) GetPosChainOk() (*string, bool)`

GetPosChainOk returns a tuple with the PosChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosChain

`func (o *Stakings) SetPosChain(v string)`

SetPosChain sets PosChain field to given value.

### HasPosChain

`func (o *Stakings) HasPosChain() bool`

HasPosChain returns a boolean if a field has been set.

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

### GetCreatedTime

`func (o *Stakings) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Stakings) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Stakings) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.


### GetUpdatedTime

`func (o *Stakings) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Stakings) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Stakings) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetValidatorInfo

`func (o *Stakings) GetValidatorInfo() StakingsValidatorInfo`

GetValidatorInfo returns the ValidatorInfo field if non-nil, zero value otherwise.

### GetValidatorInfoOk

`func (o *Stakings) GetValidatorInfoOk() (*StakingsValidatorInfo, bool)`

GetValidatorInfoOk returns a tuple with the ValidatorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorInfo

`func (o *Stakings) SetValidatorInfo(v StakingsValidatorInfo)`

SetValidatorInfo sets ValidatorInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


