# CreateStakeActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**Source** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**PoolId** | [**StakingPoolId**](StakingPoolId.md) |  | 
**Amount** | **string** | The amount to stake. | 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**Extra** | [**CreateStakeActivityExtra**](CreateStakeActivityExtra.md) |  | 
**AppInitiator** | Pointer to **string** | The initiator of the staking activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 

## Methods

### NewCreateStakeActivityRequest

`func NewCreateStakeActivityRequest(poolId StakingPoolId, amount string, fee TransactionRequestFee, extra CreateStakeActivityExtra, ) *CreateStakeActivityRequest`

NewCreateStakeActivityRequest instantiates a new CreateStakeActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStakeActivityRequestWithDefaults

`func NewCreateStakeActivityRequestWithDefaults() *CreateStakeActivityRequest`

NewCreateStakeActivityRequestWithDefaults instantiates a new CreateStakeActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateStakeActivityRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateStakeActivityRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateStakeActivityRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateStakeActivityRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSource

`func (o *CreateStakeActivityRequest) GetSource() StakingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateStakeActivityRequest) GetSourceOk() (*StakingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateStakeActivityRequest) SetSource(v StakingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateStakeActivityRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPoolId

`func (o *CreateStakeActivityRequest) GetPoolId() StakingPoolId`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateStakeActivityRequest) GetPoolIdOk() (*StakingPoolId, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateStakeActivityRequest) SetPoolId(v StakingPoolId)`

SetPoolId sets PoolId field to given value.


### GetAmount

`func (o *CreateStakeActivityRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateStakeActivityRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateStakeActivityRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *CreateStakeActivityRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateStakeActivityRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateStakeActivityRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetExtra

`func (o *CreateStakeActivityRequest) GetExtra() CreateStakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CreateStakeActivityRequest) GetExtraOk() (*CreateStakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CreateStakeActivityRequest) SetExtra(v CreateStakeActivityExtra)`

SetExtra sets Extra field to given value.


### GetAppInitiator

`func (o *CreateStakeActivityRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *CreateStakeActivityRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *CreateStakeActivityRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *CreateStakeActivityRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


