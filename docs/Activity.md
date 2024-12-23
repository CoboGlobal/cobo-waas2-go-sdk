# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The activity ID. | [optional] 
**Initiator** | Pointer to **string** | The initiator of the activity. | [optional] 
**InitiatorType** | Pointer to [**TransactionInitiatorType**](TransactionInitiatorType.md) |  | [optional] 
**Type** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**WalletId** | Pointer to **string** | The staker&#39;s wallet ID. | [optional] 
**Address** | Pointer to **string** | The staker&#39;s wallet address. | [optional] 
**PoolId** | [**StakingPoolId**](StakingPoolId.md) |  | 
**TokenId** | **string** | The token ID. | 
**StakingId** | Pointer to **string** | The ID of the corresponding staking position. | [optional] 
**RequestId** | Pointer to **string** | The request ID that is used to track a request. | [optional] 
**Amount** | **string** | The staking amount. | 
**TransactionIds** | Pointer to **[]string** | The IDs of the corresponding transactions of the activity. | [optional] 
**Timeline** | Pointer to [**[]ActivityTimeline**](ActivityTimeline.md) | The timeline of the activity. | [optional] 
**Fee** | Pointer to [**TransactionRequestFee**](TransactionRequestFee.md) |  | [optional] 
**Status** | [**ActivityStatus**](ActivityStatus.md) |  | 
**Extra** | Pointer to [**ActivityExtra**](ActivityExtra.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The time when the activity was created. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time when the activity was last updated. | [optional] 

## Methods

### NewActivity

`func NewActivity(poolId StakingPoolId, tokenId string, amount string, status ActivityStatus, ) *Activity`

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

### GetInitiatorType

`func (o *Activity) GetInitiatorType() TransactionInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *Activity) GetInitiatorTypeOk() (*TransactionInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *Activity) SetInitiatorType(v TransactionInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *Activity) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

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

`func (o *Activity) GetPoolId() StakingPoolId`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Activity) GetPoolIdOk() (*StakingPoolId, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *Activity) SetPoolId(v StakingPoolId)`

SetPoolId sets PoolId field to given value.


### GetTokenId

`func (o *Activity) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Activity) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Activity) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


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

### GetRequestId

`func (o *Activity) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Activity) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Activity) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Activity) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

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


### GetTransactionIds

`func (o *Activity) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *Activity) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *Activity) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.

### HasTransactionIds

`func (o *Activity) HasTransactionIds() bool`

HasTransactionIds returns a boolean if a field has been set.

### GetTimeline

`func (o *Activity) GetTimeline() []ActivityTimeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *Activity) GetTimelineOk() (*[]ActivityTimeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *Activity) SetTimeline(v []ActivityTimeline)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *Activity) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetFee

`func (o *Activity) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Activity) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Activity) SetFee(v TransactionRequestFee)`

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


### GetExtra

`func (o *Activity) GetExtra() ActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Activity) GetExtraOk() (*ActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Activity) SetExtra(v ActivityExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Activity) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Activity) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Activity) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Activity) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Activity) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Activity) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Activity) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Activity) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *Activity) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


