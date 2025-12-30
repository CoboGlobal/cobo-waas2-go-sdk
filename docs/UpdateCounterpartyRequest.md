# UpdateCounterpartyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedCounterpartyName** | **string** | The updated counterparty name. | 
**UpdatedCounterpartyType** | [**CounterpartyType**](CounterpartyType.md) |  | 
**UpdatedCountry** | Pointer to **string** | The updated country of the counterparty, in ISO 3166-1 alpha-3 format. | [optional] 
**UpdatedEmail** | Pointer to **string** | The updated email of the counterparty. | [optional] 
**UpdatedContactAddress** | Pointer to **string** | The updated contact address of the counterparty. | [optional] 

## Methods

### NewUpdateCounterpartyRequest

`func NewUpdateCounterpartyRequest(updatedCounterpartyName string, updatedCounterpartyType CounterpartyType, ) *UpdateCounterpartyRequest`

NewUpdateCounterpartyRequest instantiates a new UpdateCounterpartyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCounterpartyRequestWithDefaults

`func NewUpdateCounterpartyRequestWithDefaults() *UpdateCounterpartyRequest`

NewUpdateCounterpartyRequestWithDefaults instantiates a new UpdateCounterpartyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedCounterpartyName

`func (o *UpdateCounterpartyRequest) GetUpdatedCounterpartyName() string`

GetUpdatedCounterpartyName returns the UpdatedCounterpartyName field if non-nil, zero value otherwise.

### GetUpdatedCounterpartyNameOk

`func (o *UpdateCounterpartyRequest) GetUpdatedCounterpartyNameOk() (*string, bool)`

GetUpdatedCounterpartyNameOk returns a tuple with the UpdatedCounterpartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedCounterpartyName

`func (o *UpdateCounterpartyRequest) SetUpdatedCounterpartyName(v string)`

SetUpdatedCounterpartyName sets UpdatedCounterpartyName field to given value.


### GetUpdatedCounterpartyType

`func (o *UpdateCounterpartyRequest) GetUpdatedCounterpartyType() CounterpartyType`

GetUpdatedCounterpartyType returns the UpdatedCounterpartyType field if non-nil, zero value otherwise.

### GetUpdatedCounterpartyTypeOk

`func (o *UpdateCounterpartyRequest) GetUpdatedCounterpartyTypeOk() (*CounterpartyType, bool)`

GetUpdatedCounterpartyTypeOk returns a tuple with the UpdatedCounterpartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedCounterpartyType

`func (o *UpdateCounterpartyRequest) SetUpdatedCounterpartyType(v CounterpartyType)`

SetUpdatedCounterpartyType sets UpdatedCounterpartyType field to given value.


### GetUpdatedCountry

`func (o *UpdateCounterpartyRequest) GetUpdatedCountry() string`

GetUpdatedCountry returns the UpdatedCountry field if non-nil, zero value otherwise.

### GetUpdatedCountryOk

`func (o *UpdateCounterpartyRequest) GetUpdatedCountryOk() (*string, bool)`

GetUpdatedCountryOk returns a tuple with the UpdatedCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedCountry

`func (o *UpdateCounterpartyRequest) SetUpdatedCountry(v string)`

SetUpdatedCountry sets UpdatedCountry field to given value.

### HasUpdatedCountry

`func (o *UpdateCounterpartyRequest) HasUpdatedCountry() bool`

HasUpdatedCountry returns a boolean if a field has been set.

### GetUpdatedEmail

`func (o *UpdateCounterpartyRequest) GetUpdatedEmail() string`

GetUpdatedEmail returns the UpdatedEmail field if non-nil, zero value otherwise.

### GetUpdatedEmailOk

`func (o *UpdateCounterpartyRequest) GetUpdatedEmailOk() (*string, bool)`

GetUpdatedEmailOk returns a tuple with the UpdatedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEmail

`func (o *UpdateCounterpartyRequest) SetUpdatedEmail(v string)`

SetUpdatedEmail sets UpdatedEmail field to given value.

### HasUpdatedEmail

`func (o *UpdateCounterpartyRequest) HasUpdatedEmail() bool`

HasUpdatedEmail returns a boolean if a field has been set.

### GetUpdatedContactAddress

`func (o *UpdateCounterpartyRequest) GetUpdatedContactAddress() string`

GetUpdatedContactAddress returns the UpdatedContactAddress field if non-nil, zero value otherwise.

### GetUpdatedContactAddressOk

`func (o *UpdateCounterpartyRequest) GetUpdatedContactAddressOk() (*string, bool)`

GetUpdatedContactAddressOk returns a tuple with the UpdatedContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedContactAddress

`func (o *UpdateCounterpartyRequest) SetUpdatedContactAddress(v string)`

SetUpdatedContactAddress sets UpdatedContactAddress field to given value.

### HasUpdatedContactAddress

`func (o *UpdateCounterpartyRequest) HasUpdatedContactAddress() bool`

HasUpdatedContactAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


