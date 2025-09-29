# TransactionCoreStakeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraType** | [**TransactionExtraType**](TransactionExtraType.md) |  | 
**StakeAmount** | Pointer to **string** | The origin staking amount. | [optional] 
**Timelock** | Pointer to **int32** | The Unix timestamp (in seconds) when the staking position will be unlocked and available for withdrawal. | [optional] 
**ChangeAddress** | Pointer to **string** | The change address on the Bitcoin chain. If not provided, the source wallet&#39;s address will be used as the change address. | [optional] 
**ValidatorAddress** | Pointer to **string** | The validator&#39;s EVM address. | [optional] 
**RewardAddress** | Pointer to **string** | The EVM address to receive staking rewards. | [optional] 

## Methods

### NewTransactionCoreStakeInfo

`func NewTransactionCoreStakeInfo(extraType TransactionExtraType, ) *TransactionCoreStakeInfo`

NewTransactionCoreStakeInfo instantiates a new TransactionCoreStakeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCoreStakeInfoWithDefaults

`func NewTransactionCoreStakeInfoWithDefaults() *TransactionCoreStakeInfo`

NewTransactionCoreStakeInfoWithDefaults instantiates a new TransactionCoreStakeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraType

`func (o *TransactionCoreStakeInfo) GetExtraType() TransactionExtraType`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *TransactionCoreStakeInfo) GetExtraTypeOk() (*TransactionExtraType, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *TransactionCoreStakeInfo) SetExtraType(v TransactionExtraType)`

SetExtraType sets ExtraType field to given value.


### GetStakeAmount

`func (o *TransactionCoreStakeInfo) GetStakeAmount() string`

GetStakeAmount returns the StakeAmount field if non-nil, zero value otherwise.

### GetStakeAmountOk

`func (o *TransactionCoreStakeInfo) GetStakeAmountOk() (*string, bool)`

GetStakeAmountOk returns a tuple with the StakeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAmount

`func (o *TransactionCoreStakeInfo) SetStakeAmount(v string)`

SetStakeAmount sets StakeAmount field to given value.

### HasStakeAmount

`func (o *TransactionCoreStakeInfo) HasStakeAmount() bool`

HasStakeAmount returns a boolean if a field has been set.

### GetTimelock

`func (o *TransactionCoreStakeInfo) GetTimelock() int32`

GetTimelock returns the Timelock field if non-nil, zero value otherwise.

### GetTimelockOk

`func (o *TransactionCoreStakeInfo) GetTimelockOk() (*int32, bool)`

GetTimelockOk returns a tuple with the Timelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelock

`func (o *TransactionCoreStakeInfo) SetTimelock(v int32)`

SetTimelock sets Timelock field to given value.

### HasTimelock

`func (o *TransactionCoreStakeInfo) HasTimelock() bool`

HasTimelock returns a boolean if a field has been set.

### GetChangeAddress

`func (o *TransactionCoreStakeInfo) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *TransactionCoreStakeInfo) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *TransactionCoreStakeInfo) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *TransactionCoreStakeInfo) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.

### GetValidatorAddress

`func (o *TransactionCoreStakeInfo) GetValidatorAddress() string`

GetValidatorAddress returns the ValidatorAddress field if non-nil, zero value otherwise.

### GetValidatorAddressOk

`func (o *TransactionCoreStakeInfo) GetValidatorAddressOk() (*string, bool)`

GetValidatorAddressOk returns a tuple with the ValidatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorAddress

`func (o *TransactionCoreStakeInfo) SetValidatorAddress(v string)`

SetValidatorAddress sets ValidatorAddress field to given value.

### HasValidatorAddress

`func (o *TransactionCoreStakeInfo) HasValidatorAddress() bool`

HasValidatorAddress returns a boolean if a field has been set.

### GetRewardAddress

`func (o *TransactionCoreStakeInfo) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *TransactionCoreStakeInfo) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *TransactionCoreStakeInfo) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.

### HasRewardAddress

`func (o *TransactionCoreStakeInfo) HasRewardAddress() bool`

HasRewardAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


