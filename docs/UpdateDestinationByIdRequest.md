# UpdateDestinationByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationName** | **string** | The destination name. | 
**DestinationType** | [**DestinationType**](DestinationType.md) |  | 
**MerchantId** | Pointer to **string** | The ID of the merchant linked to the destination. | [optional] 
**Country** | Pointer to **string** | The country of the destination, in ISO 3166-1 alpha-3 format. | [optional] 
**Email** | Pointer to **string** | The email of the destination. | [optional] 
**ContactAddress** | Pointer to **string** | The contact address of the destination. | [optional] 

## Methods

### NewUpdateDestinationByIdRequest

`func NewUpdateDestinationByIdRequest(destinationName string, destinationType DestinationType, ) *UpdateDestinationByIdRequest`

NewUpdateDestinationByIdRequest instantiates a new UpdateDestinationByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDestinationByIdRequestWithDefaults

`func NewUpdateDestinationByIdRequestWithDefaults() *UpdateDestinationByIdRequest`

NewUpdateDestinationByIdRequestWithDefaults instantiates a new UpdateDestinationByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationName

`func (o *UpdateDestinationByIdRequest) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *UpdateDestinationByIdRequest) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *UpdateDestinationByIdRequest) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.


### GetDestinationType

`func (o *UpdateDestinationByIdRequest) GetDestinationType() DestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *UpdateDestinationByIdRequest) GetDestinationTypeOk() (*DestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *UpdateDestinationByIdRequest) SetDestinationType(v DestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMerchantId

`func (o *UpdateDestinationByIdRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UpdateDestinationByIdRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UpdateDestinationByIdRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UpdateDestinationByIdRequest) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetCountry

`func (o *UpdateDestinationByIdRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateDestinationByIdRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateDestinationByIdRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateDestinationByIdRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateDestinationByIdRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateDestinationByIdRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateDestinationByIdRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateDestinationByIdRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactAddress

`func (o *UpdateDestinationByIdRequest) GetContactAddress() string`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *UpdateDestinationByIdRequest) GetContactAddressOk() (*string, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *UpdateDestinationByIdRequest) SetContactAddress(v string)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *UpdateDestinationByIdRequest) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


