# ForcedSweep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForcedSweepId** | **string** | The force sweep ID. | 
**RequestId** | **string** | The request ID provided by you when creating the force sweep request. | 
**WalletId** | Pointer to **string** | The wallet ID to force sweep, which is the unique identifier of a wallet. | [optional] 
**TokenId** | Pointer to **string** | The token ID to force sweep, which is the unique identifier of a token. | [optional] 
**Amount** | Pointer to **string** | The amount of needing force sweep. | [optional] 
**Status** | [**ForcedSweepStatus**](ForcedSweepStatus.md) |  | 
**CreatedTimestamp** | Pointer to **int32** | The created time of the force sweep request, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the force sweep request, represented as a UNIX timestamp in seconds. | [optional] 

## Methods

### NewForcedSweep

`func NewForcedSweep(forcedSweepId string, requestId string, status ForcedSweepStatus, ) *ForcedSweep`

NewForcedSweep instantiates a new ForcedSweep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForcedSweepWithDefaults

`func NewForcedSweepWithDefaults() *ForcedSweep`

NewForcedSweepWithDefaults instantiates a new ForcedSweep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForcedSweepId

`func (o *ForcedSweep) GetForcedSweepId() string`

GetForcedSweepId returns the ForcedSweepId field if non-nil, zero value otherwise.

### GetForcedSweepIdOk

`func (o *ForcedSweep) GetForcedSweepIdOk() (*string, bool)`

GetForcedSweepIdOk returns a tuple with the ForcedSweepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSweepId

`func (o *ForcedSweep) SetForcedSweepId(v string)`

SetForcedSweepId sets ForcedSweepId field to given value.


### GetRequestId

`func (o *ForcedSweep) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ForcedSweep) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ForcedSweep) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetWalletId

`func (o *ForcedSweep) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ForcedSweep) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ForcedSweep) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *ForcedSweep) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetTokenId

`func (o *ForcedSweep) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ForcedSweep) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ForcedSweep) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ForcedSweep) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetAmount

`func (o *ForcedSweep) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ForcedSweep) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ForcedSweep) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ForcedSweep) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetStatus

`func (o *ForcedSweep) GetStatus() ForcedSweepStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ForcedSweep) GetStatusOk() (*ForcedSweepStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ForcedSweep) SetStatus(v ForcedSweepStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *ForcedSweep) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ForcedSweep) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ForcedSweep) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ForcedSweep) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ForcedSweep) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ForcedSweep) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ForcedSweep) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *ForcedSweep) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


