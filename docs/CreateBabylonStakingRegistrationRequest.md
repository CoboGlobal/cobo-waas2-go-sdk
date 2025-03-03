# CreateBabylonStakingRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StakingId** | Pointer to **string** | The ID of the Phase-1 BTC staking position. | [optional] 
**BabylonAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 

## Methods

### NewCreateBabylonStakingRegistrationRequest

`func NewCreateBabylonStakingRegistrationRequest() *CreateBabylonStakingRegistrationRequest`

NewCreateBabylonStakingRegistrationRequest instantiates a new CreateBabylonStakingRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBabylonStakingRegistrationRequestWithDefaults

`func NewCreateBabylonStakingRegistrationRequestWithDefaults() *CreateBabylonStakingRegistrationRequest`

NewCreateBabylonStakingRegistrationRequestWithDefaults instantiates a new CreateBabylonStakingRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStakingId

`func (o *CreateBabylonStakingRegistrationRequest) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *CreateBabylonStakingRegistrationRequest) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *CreateBabylonStakingRegistrationRequest) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.

### HasStakingId

`func (o *CreateBabylonStakingRegistrationRequest) HasStakingId() bool`

HasStakingId returns a boolean if a field has been set.

### GetBabylonAddress

`func (o *CreateBabylonStakingRegistrationRequest) GetBabylonAddress() StakingSource`

GetBabylonAddress returns the BabylonAddress field if non-nil, zero value otherwise.

### GetBabylonAddressOk

`func (o *CreateBabylonStakingRegistrationRequest) GetBabylonAddressOk() (*StakingSource, bool)`

GetBabylonAddressOk returns a tuple with the BabylonAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonAddress

`func (o *CreateBabylonStakingRegistrationRequest) SetBabylonAddress(v StakingSource)`

SetBabylonAddress sets BabylonAddress field to given value.

### HasBabylonAddress

`func (o *CreateBabylonStakingRegistrationRequest) HasBabylonAddress() bool`

HasBabylonAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


