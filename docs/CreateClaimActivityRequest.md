# CreateClaimActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**StakingId** | **string** | The ID of the staking position. You can retrieve a list of staking positions by calling [List staking positions](/v2/api-references/stakings/list-staking-positions). | 
**Fee** | Pointer to [**TransactionRequestFee**](TransactionRequestFee.md) |  | [optional] 
**AppInitiator** | Pointer to **string** | The initiator of the staking activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 

## Methods

### NewCreateClaimActivityRequest

`func NewCreateClaimActivityRequest(stakingId string, ) *CreateClaimActivityRequest`

NewCreateClaimActivityRequest instantiates a new CreateClaimActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClaimActivityRequestWithDefaults

`func NewCreateClaimActivityRequestWithDefaults() *CreateClaimActivityRequest`

NewCreateClaimActivityRequestWithDefaults instantiates a new CreateClaimActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateClaimActivityRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateClaimActivityRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateClaimActivityRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateClaimActivityRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStakingId

`func (o *CreateClaimActivityRequest) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *CreateClaimActivityRequest) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *CreateClaimActivityRequest) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetFee

`func (o *CreateClaimActivityRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateClaimActivityRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateClaimActivityRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CreateClaimActivityRequest) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetAppInitiator

`func (o *CreateClaimActivityRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *CreateClaimActivityRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *CreateClaimActivityRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *CreateClaimActivityRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


