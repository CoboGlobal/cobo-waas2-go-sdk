# Counterparty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartyId** | Pointer to **string** | The counterparty ID. | [optional] 
**CounterpartyType** | [**CounterpartyType**](CounterpartyType.md) |  | 
**CounterpartyName** | **string** | The counterparty name. | 
**Country** | Pointer to **string** | The country of the counterparty, in ISO 3166-1 alpha-3 format. | [optional] 
**Email** | Pointer to **string** | The email of the counterparty. | [optional] 
**ContactAddress** | Pointer to **string** | The contact address of the counterparty. | [optional] 
**CreatedTimestamp** | **int32** | The created time of the counterparty, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int32** | The updated time of the counterparty, represented as a UNIX timestamp in seconds. | 

## Methods

### NewCounterparty

`func NewCounterparty(counterpartyType CounterpartyType, counterpartyName string, createdTimestamp int32, updatedTimestamp int32, ) *Counterparty`

NewCounterparty instantiates a new Counterparty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterpartyWithDefaults

`func NewCounterpartyWithDefaults() *Counterparty`

NewCounterpartyWithDefaults instantiates a new Counterparty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartyId

`func (o *Counterparty) GetCounterpartyId() string`

GetCounterpartyId returns the CounterpartyId field if non-nil, zero value otherwise.

### GetCounterpartyIdOk

`func (o *Counterparty) GetCounterpartyIdOk() (*string, bool)`

GetCounterpartyIdOk returns a tuple with the CounterpartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyId

`func (o *Counterparty) SetCounterpartyId(v string)`

SetCounterpartyId sets CounterpartyId field to given value.

### HasCounterpartyId

`func (o *Counterparty) HasCounterpartyId() bool`

HasCounterpartyId returns a boolean if a field has been set.

### GetCounterpartyType

`func (o *Counterparty) GetCounterpartyType() CounterpartyType`

GetCounterpartyType returns the CounterpartyType field if non-nil, zero value otherwise.

### GetCounterpartyTypeOk

`func (o *Counterparty) GetCounterpartyTypeOk() (*CounterpartyType, bool)`

GetCounterpartyTypeOk returns a tuple with the CounterpartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyType

`func (o *Counterparty) SetCounterpartyType(v CounterpartyType)`

SetCounterpartyType sets CounterpartyType field to given value.


### GetCounterpartyName

`func (o *Counterparty) GetCounterpartyName() string`

GetCounterpartyName returns the CounterpartyName field if non-nil, zero value otherwise.

### GetCounterpartyNameOk

`func (o *Counterparty) GetCounterpartyNameOk() (*string, bool)`

GetCounterpartyNameOk returns a tuple with the CounterpartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyName

`func (o *Counterparty) SetCounterpartyName(v string)`

SetCounterpartyName sets CounterpartyName field to given value.


### GetCountry

`func (o *Counterparty) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Counterparty) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Counterparty) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Counterparty) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *Counterparty) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Counterparty) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Counterparty) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Counterparty) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactAddress

`func (o *Counterparty) GetContactAddress() string`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *Counterparty) GetContactAddressOk() (*string, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *Counterparty) SetContactAddress(v string)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *Counterparty) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Counterparty) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Counterparty) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Counterparty) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *Counterparty) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Counterparty) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Counterparty) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


