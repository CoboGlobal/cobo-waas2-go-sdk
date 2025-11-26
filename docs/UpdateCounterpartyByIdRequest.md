# UpdateCounterpartyByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartyName** | **string** | The counterparty name. | 
**CounterpartyType** | [**CounterpartyType**](CounterpartyType.md) |  | 
**Country** | Pointer to **string** | The country of the counterparty, in ISO 3166-1 alpha-3 format. | [optional] 
**Email** | Pointer to **string** | The email of the counterparty. | [optional] 
**ContactAddress** | Pointer to **string** | The contact address of the counterparty. | [optional] 

## Methods

### NewUpdateCounterpartyByIdRequest

`func NewUpdateCounterpartyByIdRequest(counterpartyName string, counterpartyType CounterpartyType, ) *UpdateCounterpartyByIdRequest`

NewUpdateCounterpartyByIdRequest instantiates a new UpdateCounterpartyByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCounterpartyByIdRequestWithDefaults

`func NewUpdateCounterpartyByIdRequestWithDefaults() *UpdateCounterpartyByIdRequest`

NewUpdateCounterpartyByIdRequestWithDefaults instantiates a new UpdateCounterpartyByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartyName

`func (o *UpdateCounterpartyByIdRequest) GetCounterpartyName() string`

GetCounterpartyName returns the CounterpartyName field if non-nil, zero value otherwise.

### GetCounterpartyNameOk

`func (o *UpdateCounterpartyByIdRequest) GetCounterpartyNameOk() (*string, bool)`

GetCounterpartyNameOk returns a tuple with the CounterpartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyName

`func (o *UpdateCounterpartyByIdRequest) SetCounterpartyName(v string)`

SetCounterpartyName sets CounterpartyName field to given value.


### GetCounterpartyType

`func (o *UpdateCounterpartyByIdRequest) GetCounterpartyType() CounterpartyType`

GetCounterpartyType returns the CounterpartyType field if non-nil, zero value otherwise.

### GetCounterpartyTypeOk

`func (o *UpdateCounterpartyByIdRequest) GetCounterpartyTypeOk() (*CounterpartyType, bool)`

GetCounterpartyTypeOk returns a tuple with the CounterpartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyType

`func (o *UpdateCounterpartyByIdRequest) SetCounterpartyType(v CounterpartyType)`

SetCounterpartyType sets CounterpartyType field to given value.


### GetCountry

`func (o *UpdateCounterpartyByIdRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateCounterpartyByIdRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateCounterpartyByIdRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateCounterpartyByIdRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateCounterpartyByIdRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateCounterpartyByIdRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateCounterpartyByIdRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateCounterpartyByIdRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactAddress

`func (o *UpdateCounterpartyByIdRequest) GetContactAddress() string`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *UpdateCounterpartyByIdRequest) GetContactAddressOk() (*string, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *UpdateCounterpartyByIdRequest) SetContactAddress(v string)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *UpdateCounterpartyByIdRequest) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


