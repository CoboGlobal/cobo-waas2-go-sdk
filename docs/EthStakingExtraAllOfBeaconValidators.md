# EthStakingExtraAllOfBeaconValidators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to **string** | The public key of the validator. | [optional] 
**Status** | Pointer to [**AmountStatus**](AmountStatus.md) |  | [optional] 
**Apy** | Pointer to **float32** | The annual percentage yield (APY) of the validator. | [optional] 
**StakedAmount** | Pointer to **string** | The staked amount. | [optional] 
**RewardsReceived** | Pointer to **string** | The rewards received. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time when the validator&#39;s status was last updated, in Unix timestamp format and measured in milliseconds. | [optional] 

## Methods

### NewEthStakingExtraAllOfBeaconValidators

`func NewEthStakingExtraAllOfBeaconValidators() *EthStakingExtraAllOfBeaconValidators`

NewEthStakingExtraAllOfBeaconValidators instantiates a new EthStakingExtraAllOfBeaconValidators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthStakingExtraAllOfBeaconValidatorsWithDefaults

`func NewEthStakingExtraAllOfBeaconValidatorsWithDefaults() *EthStakingExtraAllOfBeaconValidators`

NewEthStakingExtraAllOfBeaconValidatorsWithDefaults instantiates a new EthStakingExtraAllOfBeaconValidators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *EthStakingExtraAllOfBeaconValidators) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *EthStakingExtraAllOfBeaconValidators) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *EthStakingExtraAllOfBeaconValidators) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *EthStakingExtraAllOfBeaconValidators) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetStatus

`func (o *EthStakingExtraAllOfBeaconValidators) GetStatus() AmountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EthStakingExtraAllOfBeaconValidators) GetStatusOk() (*AmountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EthStakingExtraAllOfBeaconValidators) SetStatus(v AmountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EthStakingExtraAllOfBeaconValidators) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApy

`func (o *EthStakingExtraAllOfBeaconValidators) GetApy() float32`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *EthStakingExtraAllOfBeaconValidators) GetApyOk() (*float32, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *EthStakingExtraAllOfBeaconValidators) SetApy(v float32)`

SetApy sets Apy field to given value.

### HasApy

`func (o *EthStakingExtraAllOfBeaconValidators) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetStakedAmount

`func (o *EthStakingExtraAllOfBeaconValidators) GetStakedAmount() string`

GetStakedAmount returns the StakedAmount field if non-nil, zero value otherwise.

### GetStakedAmountOk

`func (o *EthStakingExtraAllOfBeaconValidators) GetStakedAmountOk() (*string, bool)`

GetStakedAmountOk returns a tuple with the StakedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakedAmount

`func (o *EthStakingExtraAllOfBeaconValidators) SetStakedAmount(v string)`

SetStakedAmount sets StakedAmount field to given value.

### HasStakedAmount

`func (o *EthStakingExtraAllOfBeaconValidators) HasStakedAmount() bool`

HasStakedAmount returns a boolean if a field has been set.

### GetRewardsReceived

`func (o *EthStakingExtraAllOfBeaconValidators) GetRewardsReceived() string`

GetRewardsReceived returns the RewardsReceived field if non-nil, zero value otherwise.

### GetRewardsReceivedOk

`func (o *EthStakingExtraAllOfBeaconValidators) GetRewardsReceivedOk() (*string, bool)`

GetRewardsReceivedOk returns a tuple with the RewardsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsReceived

`func (o *EthStakingExtraAllOfBeaconValidators) SetRewardsReceived(v string)`

SetRewardsReceived sets RewardsReceived field to given value.

### HasRewardsReceived

`func (o *EthStakingExtraAllOfBeaconValidators) HasRewardsReceived() bool`

HasRewardsReceived returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *EthStakingExtraAllOfBeaconValidators) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *EthStakingExtraAllOfBeaconValidators) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *EthStakingExtraAllOfBeaconValidators) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *EthStakingExtraAllOfBeaconValidators) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


