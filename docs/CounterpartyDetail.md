# CounterpartyDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartyId** | Pointer to **string** | The counterparty ID. | [optional] 
**CounterpartyType** | [**CounterpartyType**](CounterpartyType.md) |  | 
**CounterpartyName** | **string** | The counterparty name. | 
**Country** | Pointer to **string** | The country of the counterparty, in ISO 3166-1 alpha-3 format. | [optional] 
**Email** | Pointer to **string** | The email of the counterparty. | [optional] 
**ContactAddress** | Pointer to **string** | The contact address of the counterparty. | [optional] 
**WalletAddresses** | [**[]WalletAddress**](WalletAddress.md) | The wallet addresses of the counterparty. | 
**CreatedTimestamp** | **int32** | The created time of the counterparty, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int32** | The updated time of the counterparty, represented as a UNIX timestamp in seconds. | 

## Methods

### NewCounterpartyDetail

`func NewCounterpartyDetail(counterpartyType CounterpartyType, counterpartyName string, walletAddresses []WalletAddress, createdTimestamp int32, updatedTimestamp int32, ) *CounterpartyDetail`

NewCounterpartyDetail instantiates a new CounterpartyDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterpartyDetailWithDefaults

`func NewCounterpartyDetailWithDefaults() *CounterpartyDetail`

NewCounterpartyDetailWithDefaults instantiates a new CounterpartyDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartyId

`func (o *CounterpartyDetail) GetCounterpartyId() string`

GetCounterpartyId returns the CounterpartyId field if non-nil, zero value otherwise.

### GetCounterpartyIdOk

`func (o *CounterpartyDetail) GetCounterpartyIdOk() (*string, bool)`

GetCounterpartyIdOk returns a tuple with the CounterpartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyId

`func (o *CounterpartyDetail) SetCounterpartyId(v string)`

SetCounterpartyId sets CounterpartyId field to given value.

### HasCounterpartyId

`func (o *CounterpartyDetail) HasCounterpartyId() bool`

HasCounterpartyId returns a boolean if a field has been set.

### GetCounterpartyType

`func (o *CounterpartyDetail) GetCounterpartyType() CounterpartyType`

GetCounterpartyType returns the CounterpartyType field if non-nil, zero value otherwise.

### GetCounterpartyTypeOk

`func (o *CounterpartyDetail) GetCounterpartyTypeOk() (*CounterpartyType, bool)`

GetCounterpartyTypeOk returns a tuple with the CounterpartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyType

`func (o *CounterpartyDetail) SetCounterpartyType(v CounterpartyType)`

SetCounterpartyType sets CounterpartyType field to given value.


### GetCounterpartyName

`func (o *CounterpartyDetail) GetCounterpartyName() string`

GetCounterpartyName returns the CounterpartyName field if non-nil, zero value otherwise.

### GetCounterpartyNameOk

`func (o *CounterpartyDetail) GetCounterpartyNameOk() (*string, bool)`

GetCounterpartyNameOk returns a tuple with the CounterpartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyName

`func (o *CounterpartyDetail) SetCounterpartyName(v string)`

SetCounterpartyName sets CounterpartyName field to given value.


### GetCountry

`func (o *CounterpartyDetail) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CounterpartyDetail) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CounterpartyDetail) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CounterpartyDetail) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *CounterpartyDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CounterpartyDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CounterpartyDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CounterpartyDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactAddress

`func (o *CounterpartyDetail) GetContactAddress() string`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *CounterpartyDetail) GetContactAddressOk() (*string, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *CounterpartyDetail) SetContactAddress(v string)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *CounterpartyDetail) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.

### GetWalletAddresses

`func (o *CounterpartyDetail) GetWalletAddresses() []WalletAddress`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *CounterpartyDetail) GetWalletAddressesOk() (*[]WalletAddress, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *CounterpartyDetail) SetWalletAddresses(v []WalletAddress)`

SetWalletAddresses sets WalletAddresses field to given value.


### GetCreatedTimestamp

`func (o *CounterpartyDetail) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *CounterpartyDetail) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *CounterpartyDetail) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *CounterpartyDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *CounterpartyDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *CounterpartyDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


