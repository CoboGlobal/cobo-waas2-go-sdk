# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique id of the activity. | [optional] 
**Initiator** | Pointer to **string** | The initiator of the activity. | [optional] 
**Type** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**WalletId** | Pointer to **string** | The id of the wallet to stake. | [optional] 
**Address** | Pointer to **string** | The staker wallet address. | [optional] 
**PoolId** | **string** | The id of the staking pool. | 
**StakingId** | Pointer to **string** | The id of the related staking. | [optional] 
**Amount** | **string** | The amount of the activity. | 
**TxIds** | Pointer to **[]string** | The related txs of the activity. | [optional] 
**Fee** | Pointer to [**TransactionTransferFee**](TransactionTransferFee.md) |  | [optional] 
**Status** | [**ActivityStatus**](ActivityStatus.md) |  | 
**CreatedTime** | Pointer to **int64** | The time when the activity was created. | [optional] 
**UpdatedTime** | Pointer to **int64** | The time when the activity was last updated. | [optional] 

## Methods

### NewActivity

`func NewActivity(poolId string, amount string, status ActivityStatus, ) *Activity`

NewActivity instantiates a new Activity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityWithDefaults

`func NewActivityWithDefaults() *Activity`

NewActivityWithDefaults instantiates a new Activity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Activity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Activity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Activity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Activity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitiator

`func (o *Activity) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *Activity) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *Activity) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *Activity) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetType

`func (o *Activity) GetType() ActivityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Activity) GetTypeOk() (*ActivityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Activity) SetType(v ActivityType)`

SetType sets Type field to given value.

### HasType

`func (o *Activity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWalletId

`func (o *Activity) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *Activity) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *Activity) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *Activity) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetAddress

`func (o *Activity) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Activity) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Activity) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Activity) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPoolId

`func (o *Activity) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Activity) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *Activity) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetStakingId

`func (o *Activity) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *Activity) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *Activity) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.

### HasStakingId

`func (o *Activity) HasStakingId() bool`

HasStakingId returns a boolean if a field has been set.

### GetAmount

`func (o *Activity) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Activity) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Activity) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTxIds

`func (o *Activity) GetTxIds() []string`

GetTxIds returns the TxIds field if non-nil, zero value otherwise.

### GetTxIdsOk

`func (o *Activity) GetTxIdsOk() (*[]string, bool)`

GetTxIdsOk returns a tuple with the TxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIds

`func (o *Activity) SetTxIds(v []string)`

SetTxIds sets TxIds field to given value.

### HasTxIds

`func (o *Activity) HasTxIds() bool`

HasTxIds returns a boolean if a field has been set.

### GetFee

`func (o *Activity) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Activity) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Activity) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *Activity) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetStatus

`func (o *Activity) GetStatus() ActivityStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Activity) GetStatusOk() (*ActivityStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Activity) SetStatus(v ActivityStatus)`

SetStatus sets Status field to given value.


### GetCreatedTime

`func (o *Activity) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Activity) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Activity) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Activity) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Activity) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Activity) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Activity) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Activity) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


