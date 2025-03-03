# BabylonAirdropRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The registration ID, a unique identifier for tracking the airdrop registration. | [optional] 
**Status** | Pointer to [**BabylonRegistrationRequestStatus**](BabylonRegistrationRequestStatus.md) |  | [optional] 
**BtcAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**BabylonAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**AirdropAmount** | Pointer to **string** | The actual airdrop amount allocated for this BTC address. | [optional] 
**ErrorMessage** | Pointer to **string** | The detailed error message if the registration failed. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The time when the registration was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time when the registration was updated, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewBabylonAirdropRegistration

`func NewBabylonAirdropRegistration() *BabylonAirdropRegistration`

NewBabylonAirdropRegistration instantiates a new BabylonAirdropRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBabylonAirdropRegistrationWithDefaults

`func NewBabylonAirdropRegistrationWithDefaults() *BabylonAirdropRegistration`

NewBabylonAirdropRegistrationWithDefaults instantiates a new BabylonAirdropRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BabylonAirdropRegistration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BabylonAirdropRegistration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BabylonAirdropRegistration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BabylonAirdropRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BabylonAirdropRegistration) GetStatus() BabylonRegistrationRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BabylonAirdropRegistration) GetStatusOk() (*BabylonRegistrationRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BabylonAirdropRegistration) SetStatus(v BabylonRegistrationRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BabylonAirdropRegistration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBtcAddress

`func (o *BabylonAirdropRegistration) GetBtcAddress() StakingSource`

GetBtcAddress returns the BtcAddress field if non-nil, zero value otherwise.

### GetBtcAddressOk

`func (o *BabylonAirdropRegistration) GetBtcAddressOk() (*StakingSource, bool)`

GetBtcAddressOk returns a tuple with the BtcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtcAddress

`func (o *BabylonAirdropRegistration) SetBtcAddress(v StakingSource)`

SetBtcAddress sets BtcAddress field to given value.

### HasBtcAddress

`func (o *BabylonAirdropRegistration) HasBtcAddress() bool`

HasBtcAddress returns a boolean if a field has been set.

### GetBabylonAddress

`func (o *BabylonAirdropRegistration) GetBabylonAddress() StakingSource`

GetBabylonAddress returns the BabylonAddress field if non-nil, zero value otherwise.

### GetBabylonAddressOk

`func (o *BabylonAirdropRegistration) GetBabylonAddressOk() (*StakingSource, bool)`

GetBabylonAddressOk returns a tuple with the BabylonAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonAddress

`func (o *BabylonAirdropRegistration) SetBabylonAddress(v StakingSource)`

SetBabylonAddress sets BabylonAddress field to given value.

### HasBabylonAddress

`func (o *BabylonAirdropRegistration) HasBabylonAddress() bool`

HasBabylonAddress returns a boolean if a field has been set.

### GetAirdropAmount

`func (o *BabylonAirdropRegistration) GetAirdropAmount() string`

GetAirdropAmount returns the AirdropAmount field if non-nil, zero value otherwise.

### GetAirdropAmountOk

`func (o *BabylonAirdropRegistration) GetAirdropAmountOk() (*string, bool)`

GetAirdropAmountOk returns a tuple with the AirdropAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirdropAmount

`func (o *BabylonAirdropRegistration) SetAirdropAmount(v string)`

SetAirdropAmount sets AirdropAmount field to given value.

### HasAirdropAmount

`func (o *BabylonAirdropRegistration) HasAirdropAmount() bool`

HasAirdropAmount returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BabylonAirdropRegistration) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BabylonAirdropRegistration) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BabylonAirdropRegistration) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BabylonAirdropRegistration) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *BabylonAirdropRegistration) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *BabylonAirdropRegistration) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *BabylonAirdropRegistration) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *BabylonAirdropRegistration) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *BabylonAirdropRegistration) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *BabylonAirdropRegistration) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *BabylonAirdropRegistration) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *BabylonAirdropRegistration) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


