# BabylonStakingRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The registration ID, a unique identifier for tracking the staking registration. | [optional] 
**StakingId** | Pointer to **string** | The ID of the Phase-1 BTC staking position. | [optional] 
**BabylonAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**BtcAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**Status** | Pointer to [**BabylonRegistrationRequestStatus**](BabylonRegistrationRequestStatus.md) |  | [optional] 
**StakedAmount** | Pointer to **string** | The amount of BTC that is staked. | [optional] 
**ErrorMessage** | Pointer to **string** | The error message if the Babylon Phase-2 registration request failed. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The time when the registration was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time when the registration was updated, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewBabylonStakingRegistration

`func NewBabylonStakingRegistration() *BabylonStakingRegistration`

NewBabylonStakingRegistration instantiates a new BabylonStakingRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBabylonStakingRegistrationWithDefaults

`func NewBabylonStakingRegistrationWithDefaults() *BabylonStakingRegistration`

NewBabylonStakingRegistrationWithDefaults instantiates a new BabylonStakingRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BabylonStakingRegistration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BabylonStakingRegistration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BabylonStakingRegistration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BabylonStakingRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStakingId

`func (o *BabylonStakingRegistration) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *BabylonStakingRegistration) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *BabylonStakingRegistration) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.

### HasStakingId

`func (o *BabylonStakingRegistration) HasStakingId() bool`

HasStakingId returns a boolean if a field has been set.

### GetBabylonAddress

`func (o *BabylonStakingRegistration) GetBabylonAddress() StakingSource`

GetBabylonAddress returns the BabylonAddress field if non-nil, zero value otherwise.

### GetBabylonAddressOk

`func (o *BabylonStakingRegistration) GetBabylonAddressOk() (*StakingSource, bool)`

GetBabylonAddressOk returns a tuple with the BabylonAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonAddress

`func (o *BabylonStakingRegistration) SetBabylonAddress(v StakingSource)`

SetBabylonAddress sets BabylonAddress field to given value.

### HasBabylonAddress

`func (o *BabylonStakingRegistration) HasBabylonAddress() bool`

HasBabylonAddress returns a boolean if a field has been set.

### GetBtcAddress

`func (o *BabylonStakingRegistration) GetBtcAddress() StakingSource`

GetBtcAddress returns the BtcAddress field if non-nil, zero value otherwise.

### GetBtcAddressOk

`func (o *BabylonStakingRegistration) GetBtcAddressOk() (*StakingSource, bool)`

GetBtcAddressOk returns a tuple with the BtcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtcAddress

`func (o *BabylonStakingRegistration) SetBtcAddress(v StakingSource)`

SetBtcAddress sets BtcAddress field to given value.

### HasBtcAddress

`func (o *BabylonStakingRegistration) HasBtcAddress() bool`

HasBtcAddress returns a boolean if a field has been set.

### GetStatus

`func (o *BabylonStakingRegistration) GetStatus() BabylonRegistrationRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BabylonStakingRegistration) GetStatusOk() (*BabylonRegistrationRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BabylonStakingRegistration) SetStatus(v BabylonRegistrationRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BabylonStakingRegistration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStakedAmount

`func (o *BabylonStakingRegistration) GetStakedAmount() string`

GetStakedAmount returns the StakedAmount field if non-nil, zero value otherwise.

### GetStakedAmountOk

`func (o *BabylonStakingRegistration) GetStakedAmountOk() (*string, bool)`

GetStakedAmountOk returns a tuple with the StakedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakedAmount

`func (o *BabylonStakingRegistration) SetStakedAmount(v string)`

SetStakedAmount sets StakedAmount field to given value.

### HasStakedAmount

`func (o *BabylonStakingRegistration) HasStakedAmount() bool`

HasStakedAmount returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BabylonStakingRegistration) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BabylonStakingRegistration) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BabylonStakingRegistration) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BabylonStakingRegistration) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *BabylonStakingRegistration) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *BabylonStakingRegistration) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *BabylonStakingRegistration) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *BabylonStakingRegistration) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *BabylonStakingRegistration) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *BabylonStakingRegistration) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *BabylonStakingRegistration) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *BabylonStakingRegistration) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


