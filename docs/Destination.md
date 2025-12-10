# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | Pointer to **string** | The destination ID. | [optional] 
**DestinationType** | [**DestinationType**](DestinationType.md) |  | 
**DestinationName** | **string** | The destination name. | 
**Country** | Pointer to **string** | The country of the destination, in ISO 3166-1 alpha-3 format. | [optional] 
**Email** | Pointer to **string** | The email of the destination. | [optional] 
**ContactAddress** | Pointer to **string** | The contact address of the destination. | [optional] 
**MerchantId** | Pointer to **string** | The ID of the merchant linked to the destination. | [optional] 
**CreatedTimestamp** | **int32** | The created time of the destination, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int32** | The updated time of the destination, represented as a UNIX timestamp in seconds. | 

## Methods

### NewDestination

`func NewDestination(destinationType DestinationType, destinationName string, createdTimestamp int32, updatedTimestamp int32, ) *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *Destination) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *Destination) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *Destination) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *Destination) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetDestinationType

`func (o *Destination) GetDestinationType() DestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *Destination) GetDestinationTypeOk() (*DestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *Destination) SetDestinationType(v DestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetDestinationName

`func (o *Destination) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *Destination) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *Destination) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.


### GetCountry

`func (o *Destination) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Destination) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Destination) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Destination) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *Destination) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Destination) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Destination) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Destination) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactAddress

`func (o *Destination) GetContactAddress() string`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *Destination) GetContactAddressOk() (*string, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *Destination) SetContactAddress(v string)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *Destination) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.

### GetMerchantId

`func (o *Destination) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *Destination) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *Destination) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *Destination) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Destination) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Destination) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Destination) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *Destination) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Destination) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Destination) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


