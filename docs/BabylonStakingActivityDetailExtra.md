# BabylonStakingActivityDetailExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**FinalityProviderPublicKey** | Pointer to **string** | The public key of the finality provider. | [optional] 
**StakeBlockTime** | Pointer to **int64** | The number of blocks that need to be processed before the locked tokens are unlocked and become accessible. | [optional] 
**AutoBroadcast** | Pointer to **bool** | Whether to automatically broadcast the transaction.  - &#x60;true&#x60;: Automatically broadcast the transaction. - &#x60;false&#x60;: The transaction will not be submitted to the blockchain automatically. You can call [Broadcast signed transactions](https://www.cobo.com/developers/v2/api-references/transactions/broadcast-signed-transactions) to broadcast the transaction to the blockchain, or retrieve the signed raw transaction data &#x60;raw_tx&#x60; by calling [Get transaction information](https://www.cobo.com/developers/v2/api-references/transactions/get-transaction-information) and broadcast it yourself.  | [optional] 
**ParamVersion** | Pointer to **int64** | The version of babylon global parameters. | [optional] 
**WithdrawFromType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**SlashFromType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**StakeAmount** | Pointer to **string** | The origin staking amount. | [optional] 

## Methods

### NewBabylonStakingActivityDetailExtra

`func NewBabylonStakingActivityDetailExtra(poolType StakingPoolType, ) *BabylonStakingActivityDetailExtra`

NewBabylonStakingActivityDetailExtra instantiates a new BabylonStakingActivityDetailExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBabylonStakingActivityDetailExtraWithDefaults

`func NewBabylonStakingActivityDetailExtraWithDefaults() *BabylonStakingActivityDetailExtra`

NewBabylonStakingActivityDetailExtraWithDefaults instantiates a new BabylonStakingActivityDetailExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *BabylonStakingActivityDetailExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *BabylonStakingActivityDetailExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *BabylonStakingActivityDetailExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetFinalityProviderPublicKey

`func (o *BabylonStakingActivityDetailExtra) GetFinalityProviderPublicKey() string`

GetFinalityProviderPublicKey returns the FinalityProviderPublicKey field if non-nil, zero value otherwise.

### GetFinalityProviderPublicKeyOk

`func (o *BabylonStakingActivityDetailExtra) GetFinalityProviderPublicKeyOk() (*string, bool)`

GetFinalityProviderPublicKeyOk returns a tuple with the FinalityProviderPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalityProviderPublicKey

`func (o *BabylonStakingActivityDetailExtra) SetFinalityProviderPublicKey(v string)`

SetFinalityProviderPublicKey sets FinalityProviderPublicKey field to given value.

### HasFinalityProviderPublicKey

`func (o *BabylonStakingActivityDetailExtra) HasFinalityProviderPublicKey() bool`

HasFinalityProviderPublicKey returns a boolean if a field has been set.

### GetStakeBlockTime

`func (o *BabylonStakingActivityDetailExtra) GetStakeBlockTime() int64`

GetStakeBlockTime returns the StakeBlockTime field if non-nil, zero value otherwise.

### GetStakeBlockTimeOk

`func (o *BabylonStakingActivityDetailExtra) GetStakeBlockTimeOk() (*int64, bool)`

GetStakeBlockTimeOk returns a tuple with the StakeBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeBlockTime

`func (o *BabylonStakingActivityDetailExtra) SetStakeBlockTime(v int64)`

SetStakeBlockTime sets StakeBlockTime field to given value.

### HasStakeBlockTime

`func (o *BabylonStakingActivityDetailExtra) HasStakeBlockTime() bool`

HasStakeBlockTime returns a boolean if a field has been set.

### GetAutoBroadcast

`func (o *BabylonStakingActivityDetailExtra) GetAutoBroadcast() bool`

GetAutoBroadcast returns the AutoBroadcast field if non-nil, zero value otherwise.

### GetAutoBroadcastOk

`func (o *BabylonStakingActivityDetailExtra) GetAutoBroadcastOk() (*bool, bool)`

GetAutoBroadcastOk returns a tuple with the AutoBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBroadcast

`func (o *BabylonStakingActivityDetailExtra) SetAutoBroadcast(v bool)`

SetAutoBroadcast sets AutoBroadcast field to given value.

### HasAutoBroadcast

`func (o *BabylonStakingActivityDetailExtra) HasAutoBroadcast() bool`

HasAutoBroadcast returns a boolean if a field has been set.

### GetParamVersion

`func (o *BabylonStakingActivityDetailExtra) GetParamVersion() int64`

GetParamVersion returns the ParamVersion field if non-nil, zero value otherwise.

### GetParamVersionOk

`func (o *BabylonStakingActivityDetailExtra) GetParamVersionOk() (*int64, bool)`

GetParamVersionOk returns a tuple with the ParamVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamVersion

`func (o *BabylonStakingActivityDetailExtra) SetParamVersion(v int64)`

SetParamVersion sets ParamVersion field to given value.

### HasParamVersion

`func (o *BabylonStakingActivityDetailExtra) HasParamVersion() bool`

HasParamVersion returns a boolean if a field has been set.

### GetWithdrawFromType

`func (o *BabylonStakingActivityDetailExtra) GetWithdrawFromType() ActivityType`

GetWithdrawFromType returns the WithdrawFromType field if non-nil, zero value otherwise.

### GetWithdrawFromTypeOk

`func (o *BabylonStakingActivityDetailExtra) GetWithdrawFromTypeOk() (*ActivityType, bool)`

GetWithdrawFromTypeOk returns a tuple with the WithdrawFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawFromType

`func (o *BabylonStakingActivityDetailExtra) SetWithdrawFromType(v ActivityType)`

SetWithdrawFromType sets WithdrawFromType field to given value.

### HasWithdrawFromType

`func (o *BabylonStakingActivityDetailExtra) HasWithdrawFromType() bool`

HasWithdrawFromType returns a boolean if a field has been set.

### GetSlashFromType

`func (o *BabylonStakingActivityDetailExtra) GetSlashFromType() ActivityType`

GetSlashFromType returns the SlashFromType field if non-nil, zero value otherwise.

### GetSlashFromTypeOk

`func (o *BabylonStakingActivityDetailExtra) GetSlashFromTypeOk() (*ActivityType, bool)`

GetSlashFromTypeOk returns a tuple with the SlashFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashFromType

`func (o *BabylonStakingActivityDetailExtra) SetSlashFromType(v ActivityType)`

SetSlashFromType sets SlashFromType field to given value.

### HasSlashFromType

`func (o *BabylonStakingActivityDetailExtra) HasSlashFromType() bool`

HasSlashFromType returns a boolean if a field has been set.

### GetStakeAmount

`func (o *BabylonStakingActivityDetailExtra) GetStakeAmount() string`

GetStakeAmount returns the StakeAmount field if non-nil, zero value otherwise.

### GetStakeAmountOk

`func (o *BabylonStakingActivityDetailExtra) GetStakeAmountOk() (*string, bool)`

GetStakeAmountOk returns a tuple with the StakeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAmount

`func (o *BabylonStakingActivityDetailExtra) SetStakeAmount(v string)`

SetStakeAmount sets StakeAmount field to given value.

### HasStakeAmount

`func (o *BabylonStakingActivityDetailExtra) HasStakeAmount() bool`

HasStakeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


