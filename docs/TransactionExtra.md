# TransactionExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraType** | [**TransactionExtraType**](TransactionExtraType.md) |  | 
**BabylonAddressInfo** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**BtcAddressInfo** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**StakeAmount** | Pointer to **string** | The origin staking amount. | [optional] 
**FinalityProviderPublicKey** | Pointer to **string** | The public key of the finality provider. | [optional] 
**StakeBlockTime** | Pointer to **int64** | The number of blocks that need to be processed before the locked tokens are unlocked and become accessible. | [optional] 
**ParamVersion** | Pointer to **int64** | The version of babylon global parameters. | [optional] 
**WithdrawFromType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**SlashFromType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**Timelock** | Pointer to **int32** | The Unix timestamp (in seconds) when the staking position will be unlocked and available for withdrawal. | [optional] 
**ChangeAddress** | Pointer to **string** | The change address on the Bitcoin chain. If not provided, the source wallet&#39;s address will be used as the change address. | [optional] 
**ValidatorAddress** | Pointer to **string** | The validator&#39;s EVM address. | [optional] 
**RewardAddress** | Pointer to **string** | The EVM address to receive staking rewards. | [optional] 

## Methods

### NewTransactionExtra

`func NewTransactionExtra(extraType TransactionExtraType, ) *TransactionExtra`

NewTransactionExtra instantiates a new TransactionExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionExtraWithDefaults

`func NewTransactionExtraWithDefaults() *TransactionExtra`

NewTransactionExtraWithDefaults instantiates a new TransactionExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraType

`func (o *TransactionExtra) GetExtraType() TransactionExtraType`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *TransactionExtra) GetExtraTypeOk() (*TransactionExtraType, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *TransactionExtra) SetExtraType(v TransactionExtraType)`

SetExtraType sets ExtraType field to given value.


### GetBabylonAddressInfo

`func (o *TransactionExtra) GetBabylonAddressInfo() AddressInfo`

GetBabylonAddressInfo returns the BabylonAddressInfo field if non-nil, zero value otherwise.

### GetBabylonAddressInfoOk

`func (o *TransactionExtra) GetBabylonAddressInfoOk() (*AddressInfo, bool)`

GetBabylonAddressInfoOk returns a tuple with the BabylonAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonAddressInfo

`func (o *TransactionExtra) SetBabylonAddressInfo(v AddressInfo)`

SetBabylonAddressInfo sets BabylonAddressInfo field to given value.

### HasBabylonAddressInfo

`func (o *TransactionExtra) HasBabylonAddressInfo() bool`

HasBabylonAddressInfo returns a boolean if a field has been set.

### GetBtcAddressInfo

`func (o *TransactionExtra) GetBtcAddressInfo() AddressInfo`

GetBtcAddressInfo returns the BtcAddressInfo field if non-nil, zero value otherwise.

### GetBtcAddressInfoOk

`func (o *TransactionExtra) GetBtcAddressInfoOk() (*AddressInfo, bool)`

GetBtcAddressInfoOk returns a tuple with the BtcAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtcAddressInfo

`func (o *TransactionExtra) SetBtcAddressInfo(v AddressInfo)`

SetBtcAddressInfo sets BtcAddressInfo field to given value.

### HasBtcAddressInfo

`func (o *TransactionExtra) HasBtcAddressInfo() bool`

HasBtcAddressInfo returns a boolean if a field has been set.

### GetStakeAmount

`func (o *TransactionExtra) GetStakeAmount() string`

GetStakeAmount returns the StakeAmount field if non-nil, zero value otherwise.

### GetStakeAmountOk

`func (o *TransactionExtra) GetStakeAmountOk() (*string, bool)`

GetStakeAmountOk returns a tuple with the StakeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAmount

`func (o *TransactionExtra) SetStakeAmount(v string)`

SetStakeAmount sets StakeAmount field to given value.

### HasStakeAmount

`func (o *TransactionExtra) HasStakeAmount() bool`

HasStakeAmount returns a boolean if a field has been set.

### GetFinalityProviderPublicKey

`func (o *TransactionExtra) GetFinalityProviderPublicKey() string`

GetFinalityProviderPublicKey returns the FinalityProviderPublicKey field if non-nil, zero value otherwise.

### GetFinalityProviderPublicKeyOk

`func (o *TransactionExtra) GetFinalityProviderPublicKeyOk() (*string, bool)`

GetFinalityProviderPublicKeyOk returns a tuple with the FinalityProviderPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalityProviderPublicKey

`func (o *TransactionExtra) SetFinalityProviderPublicKey(v string)`

SetFinalityProviderPublicKey sets FinalityProviderPublicKey field to given value.

### HasFinalityProviderPublicKey

`func (o *TransactionExtra) HasFinalityProviderPublicKey() bool`

HasFinalityProviderPublicKey returns a boolean if a field has been set.

### GetStakeBlockTime

`func (o *TransactionExtra) GetStakeBlockTime() int64`

GetStakeBlockTime returns the StakeBlockTime field if non-nil, zero value otherwise.

### GetStakeBlockTimeOk

`func (o *TransactionExtra) GetStakeBlockTimeOk() (*int64, bool)`

GetStakeBlockTimeOk returns a tuple with the StakeBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeBlockTime

`func (o *TransactionExtra) SetStakeBlockTime(v int64)`

SetStakeBlockTime sets StakeBlockTime field to given value.

### HasStakeBlockTime

`func (o *TransactionExtra) HasStakeBlockTime() bool`

HasStakeBlockTime returns a boolean if a field has been set.

### GetParamVersion

`func (o *TransactionExtra) GetParamVersion() int64`

GetParamVersion returns the ParamVersion field if non-nil, zero value otherwise.

### GetParamVersionOk

`func (o *TransactionExtra) GetParamVersionOk() (*int64, bool)`

GetParamVersionOk returns a tuple with the ParamVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamVersion

`func (o *TransactionExtra) SetParamVersion(v int64)`

SetParamVersion sets ParamVersion field to given value.

### HasParamVersion

`func (o *TransactionExtra) HasParamVersion() bool`

HasParamVersion returns a boolean if a field has been set.

### GetWithdrawFromType

`func (o *TransactionExtra) GetWithdrawFromType() ActivityType`

GetWithdrawFromType returns the WithdrawFromType field if non-nil, zero value otherwise.

### GetWithdrawFromTypeOk

`func (o *TransactionExtra) GetWithdrawFromTypeOk() (*ActivityType, bool)`

GetWithdrawFromTypeOk returns a tuple with the WithdrawFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawFromType

`func (o *TransactionExtra) SetWithdrawFromType(v ActivityType)`

SetWithdrawFromType sets WithdrawFromType field to given value.

### HasWithdrawFromType

`func (o *TransactionExtra) HasWithdrawFromType() bool`

HasWithdrawFromType returns a boolean if a field has been set.

### GetSlashFromType

`func (o *TransactionExtra) GetSlashFromType() ActivityType`

GetSlashFromType returns the SlashFromType field if non-nil, zero value otherwise.

### GetSlashFromTypeOk

`func (o *TransactionExtra) GetSlashFromTypeOk() (*ActivityType, bool)`

GetSlashFromTypeOk returns a tuple with the SlashFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashFromType

`func (o *TransactionExtra) SetSlashFromType(v ActivityType)`

SetSlashFromType sets SlashFromType field to given value.

### HasSlashFromType

`func (o *TransactionExtra) HasSlashFromType() bool`

HasSlashFromType returns a boolean if a field has been set.

### GetTimelock

`func (o *TransactionExtra) GetTimelock() int32`

GetTimelock returns the Timelock field if non-nil, zero value otherwise.

### GetTimelockOk

`func (o *TransactionExtra) GetTimelockOk() (*int32, bool)`

GetTimelockOk returns a tuple with the Timelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelock

`func (o *TransactionExtra) SetTimelock(v int32)`

SetTimelock sets Timelock field to given value.

### HasTimelock

`func (o *TransactionExtra) HasTimelock() bool`

HasTimelock returns a boolean if a field has been set.

### GetChangeAddress

`func (o *TransactionExtra) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *TransactionExtra) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *TransactionExtra) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *TransactionExtra) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.

### GetValidatorAddress

`func (o *TransactionExtra) GetValidatorAddress() string`

GetValidatorAddress returns the ValidatorAddress field if non-nil, zero value otherwise.

### GetValidatorAddressOk

`func (o *TransactionExtra) GetValidatorAddressOk() (*string, bool)`

GetValidatorAddressOk returns a tuple with the ValidatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorAddress

`func (o *TransactionExtra) SetValidatorAddress(v string)`

SetValidatorAddress sets ValidatorAddress field to given value.

### HasValidatorAddress

`func (o *TransactionExtra) HasValidatorAddress() bool`

HasValidatorAddress returns a boolean if a field has been set.

### GetRewardAddress

`func (o *TransactionExtra) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *TransactionExtra) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *TransactionExtra) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.

### HasRewardAddress

`func (o *TransactionExtra) HasRewardAddress() bool`

HasRewardAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


