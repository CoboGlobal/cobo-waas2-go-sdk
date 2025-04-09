# ActivityExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**FinalityProviderPublicKey** | Pointer to **string** | The public key of the finality provider. | [optional] 
**StakeBlockTime** | Pointer to **int64** | The number of blocks that need to be processed before the locked tokens are unlocked and become accessible. | [optional] 
**AutoBroadcast** | Pointer to **bool** | Whether to automatically broadcast the transaction.  - &#x60;true&#x60;: Automatically broadcast the transaction. - &#x60;false&#x60;: The transaction will not be submitted to the blockchain automatically. You can call [Broadcast signed transactions](https://www.cobo.com/developers/v2/api-references/transactions/broadcast-signed-transactions) to broadcast the transaction to the blockchain, or retrieve the signed raw transaction data &#x60;raw_tx&#x60; by calling [Get transaction information](https://www.cobo.com/developers/v2/api-references/transactions/get-transaction-information) and broadcast it yourself.  | [optional] 
**ProviderName** | Pointer to **string** | The name of the provider. | [optional] 
**ValidatorPubkeys** | Pointer to **[]string** | A list of public keys associated with the Ethereum validators for this unstaking operation. | [optional] 
**Timelock** | Pointer to **int32** | The Unix timestamp (in seconds) when the staking position will be unlocked and available for withdrawal. | [optional] 
**ChangeAddress** | Pointer to **string** | The change address on the Bitcoin chain. If not provided, the source wallet&#39;s address will be used as the change address. | [optional] 
**ValidatorAddress** | Pointer to **string** | The validator&#39;s EVM address. | [optional] 
**RewardAddress** | Pointer to **string** | The EVM address to receive staking rewards. | [optional] 

## Methods

### NewActivityExtra

`func NewActivityExtra(poolType StakingPoolType, ) *ActivityExtra`

NewActivityExtra instantiates a new ActivityExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityExtraWithDefaults

`func NewActivityExtraWithDefaults() *ActivityExtra`

NewActivityExtraWithDefaults instantiates a new ActivityExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *ActivityExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *ActivityExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *ActivityExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetFinalityProviderPublicKey

`func (o *ActivityExtra) GetFinalityProviderPublicKey() string`

GetFinalityProviderPublicKey returns the FinalityProviderPublicKey field if non-nil, zero value otherwise.

### GetFinalityProviderPublicKeyOk

`func (o *ActivityExtra) GetFinalityProviderPublicKeyOk() (*string, bool)`

GetFinalityProviderPublicKeyOk returns a tuple with the FinalityProviderPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalityProviderPublicKey

`func (o *ActivityExtra) SetFinalityProviderPublicKey(v string)`

SetFinalityProviderPublicKey sets FinalityProviderPublicKey field to given value.

### HasFinalityProviderPublicKey

`func (o *ActivityExtra) HasFinalityProviderPublicKey() bool`

HasFinalityProviderPublicKey returns a boolean if a field has been set.

### GetStakeBlockTime

`func (o *ActivityExtra) GetStakeBlockTime() int64`

GetStakeBlockTime returns the StakeBlockTime field if non-nil, zero value otherwise.

### GetStakeBlockTimeOk

`func (o *ActivityExtra) GetStakeBlockTimeOk() (*int64, bool)`

GetStakeBlockTimeOk returns a tuple with the StakeBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeBlockTime

`func (o *ActivityExtra) SetStakeBlockTime(v int64)`

SetStakeBlockTime sets StakeBlockTime field to given value.

### HasStakeBlockTime

`func (o *ActivityExtra) HasStakeBlockTime() bool`

HasStakeBlockTime returns a boolean if a field has been set.

### GetAutoBroadcast

`func (o *ActivityExtra) GetAutoBroadcast() bool`

GetAutoBroadcast returns the AutoBroadcast field if non-nil, zero value otherwise.

### GetAutoBroadcastOk

`func (o *ActivityExtra) GetAutoBroadcastOk() (*bool, bool)`

GetAutoBroadcastOk returns a tuple with the AutoBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBroadcast

`func (o *ActivityExtra) SetAutoBroadcast(v bool)`

SetAutoBroadcast sets AutoBroadcast field to given value.

### HasAutoBroadcast

`func (o *ActivityExtra) HasAutoBroadcast() bool`

HasAutoBroadcast returns a boolean if a field has been set.

### GetProviderName

`func (o *ActivityExtra) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ActivityExtra) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ActivityExtra) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ActivityExtra) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetValidatorPubkeys

`func (o *ActivityExtra) GetValidatorPubkeys() []string`

GetValidatorPubkeys returns the ValidatorPubkeys field if non-nil, zero value otherwise.

### GetValidatorPubkeysOk

`func (o *ActivityExtra) GetValidatorPubkeysOk() (*[]string, bool)`

GetValidatorPubkeysOk returns a tuple with the ValidatorPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorPubkeys

`func (o *ActivityExtra) SetValidatorPubkeys(v []string)`

SetValidatorPubkeys sets ValidatorPubkeys field to given value.

### HasValidatorPubkeys

`func (o *ActivityExtra) HasValidatorPubkeys() bool`

HasValidatorPubkeys returns a boolean if a field has been set.

### GetTimelock

`func (o *ActivityExtra) GetTimelock() int32`

GetTimelock returns the Timelock field if non-nil, zero value otherwise.

### GetTimelockOk

`func (o *ActivityExtra) GetTimelockOk() (*int32, bool)`

GetTimelockOk returns a tuple with the Timelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelock

`func (o *ActivityExtra) SetTimelock(v int32)`

SetTimelock sets Timelock field to given value.

### HasTimelock

`func (o *ActivityExtra) HasTimelock() bool`

HasTimelock returns a boolean if a field has been set.

### GetChangeAddress

`func (o *ActivityExtra) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *ActivityExtra) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *ActivityExtra) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *ActivityExtra) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.

### GetValidatorAddress

`func (o *ActivityExtra) GetValidatorAddress() string`

GetValidatorAddress returns the ValidatorAddress field if non-nil, zero value otherwise.

### GetValidatorAddressOk

`func (o *ActivityExtra) GetValidatorAddressOk() (*string, bool)`

GetValidatorAddressOk returns a tuple with the ValidatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorAddress

`func (o *ActivityExtra) SetValidatorAddress(v string)`

SetValidatorAddress sets ValidatorAddress field to given value.

### HasValidatorAddress

`func (o *ActivityExtra) HasValidatorAddress() bool`

HasValidatorAddress returns a boolean if a field has been set.

### GetRewardAddress

`func (o *ActivityExtra) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *ActivityExtra) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *ActivityExtra) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.

### HasRewardAddress

`func (o *ActivityExtra) HasRewardAddress() bool`

HasRewardAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


