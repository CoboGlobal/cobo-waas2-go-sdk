# TransactionBabylonTxParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraType** | [**TransactionExtraType**](TransactionExtraType.md) |  | 
**StakeAmount** | Pointer to **string** | The original staking amount. | [optional] 
**FinalityProviderPublicKey** | Pointer to **string** | The public key of the finality provider. | [optional] 
**StakeBlockTime** | Pointer to **int64** | The number of blocks that need to be processed before the locked tokens are unlocked and become accessible. | [optional] 
**ParamVersion** | Pointer to **int64** | The version of Babylon global parameters. | [optional] 
**WithdrawFromType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**SlashFromType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 

## Methods

### NewTransactionBabylonTxParameters

`func NewTransactionBabylonTxParameters(extraType TransactionExtraType, ) *TransactionBabylonTxParameters`

NewTransactionBabylonTxParameters instantiates a new TransactionBabylonTxParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionBabylonTxParametersWithDefaults

`func NewTransactionBabylonTxParametersWithDefaults() *TransactionBabylonTxParameters`

NewTransactionBabylonTxParametersWithDefaults instantiates a new TransactionBabylonTxParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraType

`func (o *TransactionBabylonTxParameters) GetExtraType() TransactionExtraType`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *TransactionBabylonTxParameters) GetExtraTypeOk() (*TransactionExtraType, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *TransactionBabylonTxParameters) SetExtraType(v TransactionExtraType)`

SetExtraType sets ExtraType field to given value.


### GetStakeAmount

`func (o *TransactionBabylonTxParameters) GetStakeAmount() string`

GetStakeAmount returns the StakeAmount field if non-nil, zero value otherwise.

### GetStakeAmountOk

`func (o *TransactionBabylonTxParameters) GetStakeAmountOk() (*string, bool)`

GetStakeAmountOk returns a tuple with the StakeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAmount

`func (o *TransactionBabylonTxParameters) SetStakeAmount(v string)`

SetStakeAmount sets StakeAmount field to given value.

### HasStakeAmount

`func (o *TransactionBabylonTxParameters) HasStakeAmount() bool`

HasStakeAmount returns a boolean if a field has been set.

### GetFinalityProviderPublicKey

`func (o *TransactionBabylonTxParameters) GetFinalityProviderPublicKey() string`

GetFinalityProviderPublicKey returns the FinalityProviderPublicKey field if non-nil, zero value otherwise.

### GetFinalityProviderPublicKeyOk

`func (o *TransactionBabylonTxParameters) GetFinalityProviderPublicKeyOk() (*string, bool)`

GetFinalityProviderPublicKeyOk returns a tuple with the FinalityProviderPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalityProviderPublicKey

`func (o *TransactionBabylonTxParameters) SetFinalityProviderPublicKey(v string)`

SetFinalityProviderPublicKey sets FinalityProviderPublicKey field to given value.

### HasFinalityProviderPublicKey

`func (o *TransactionBabylonTxParameters) HasFinalityProviderPublicKey() bool`

HasFinalityProviderPublicKey returns a boolean if a field has been set.

### GetStakeBlockTime

`func (o *TransactionBabylonTxParameters) GetStakeBlockTime() int64`

GetStakeBlockTime returns the StakeBlockTime field if non-nil, zero value otherwise.

### GetStakeBlockTimeOk

`func (o *TransactionBabylonTxParameters) GetStakeBlockTimeOk() (*int64, bool)`

GetStakeBlockTimeOk returns a tuple with the StakeBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeBlockTime

`func (o *TransactionBabylonTxParameters) SetStakeBlockTime(v int64)`

SetStakeBlockTime sets StakeBlockTime field to given value.

### HasStakeBlockTime

`func (o *TransactionBabylonTxParameters) HasStakeBlockTime() bool`

HasStakeBlockTime returns a boolean if a field has been set.

### GetParamVersion

`func (o *TransactionBabylonTxParameters) GetParamVersion() int64`

GetParamVersion returns the ParamVersion field if non-nil, zero value otherwise.

### GetParamVersionOk

`func (o *TransactionBabylonTxParameters) GetParamVersionOk() (*int64, bool)`

GetParamVersionOk returns a tuple with the ParamVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamVersion

`func (o *TransactionBabylonTxParameters) SetParamVersion(v int64)`

SetParamVersion sets ParamVersion field to given value.

### HasParamVersion

`func (o *TransactionBabylonTxParameters) HasParamVersion() bool`

HasParamVersion returns a boolean if a field has been set.

### GetWithdrawFromType

`func (o *TransactionBabylonTxParameters) GetWithdrawFromType() ActivityType`

GetWithdrawFromType returns the WithdrawFromType field if non-nil, zero value otherwise.

### GetWithdrawFromTypeOk

`func (o *TransactionBabylonTxParameters) GetWithdrawFromTypeOk() (*ActivityType, bool)`

GetWithdrawFromTypeOk returns a tuple with the WithdrawFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawFromType

`func (o *TransactionBabylonTxParameters) SetWithdrawFromType(v ActivityType)`

SetWithdrawFromType sets WithdrawFromType field to given value.

### HasWithdrawFromType

`func (o *TransactionBabylonTxParameters) HasWithdrawFromType() bool`

HasWithdrawFromType returns a boolean if a field has been set.

### GetSlashFromType

`func (o *TransactionBabylonTxParameters) GetSlashFromType() ActivityType`

GetSlashFromType returns the SlashFromType field if non-nil, zero value otherwise.

### GetSlashFromTypeOk

`func (o *TransactionBabylonTxParameters) GetSlashFromTypeOk() (*ActivityType, bool)`

GetSlashFromTypeOk returns a tuple with the SlashFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashFromType

`func (o *TransactionBabylonTxParameters) SetSlashFromType(v ActivityType)`

SetSlashFromType sets SlashFromType field to given value.

### HasSlashFromType

`func (o *TransactionBabylonTxParameters) HasSlashFromType() bool`

HasSlashFromType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


