# UpdateDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedDestinationName** | **string** | The updated destination name. | 
**UpdatedDestinationType** | [**DestinationType**](DestinationType.md) |  | 
**UpdatedMerchantId** | Pointer to **string** | The updated ID of the merchant linked to the destination. | [optional] 
**UpdatedCountry** | Pointer to **string** | The updated country of the destination, in ISO 3166-1 alpha-3 format. | [optional] 
**UpdatedEmail** | Pointer to **string** | The updated email of the destination. | [optional] 
**UpdatedContactAddress** | Pointer to **string** | The updated contact address of the destination. | [optional] 

## Methods

### NewUpdateDestinationRequest

`func NewUpdateDestinationRequest(updatedDestinationName string, updatedDestinationType DestinationType, ) *UpdateDestinationRequest`

NewUpdateDestinationRequest instantiates a new UpdateDestinationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDestinationRequestWithDefaults

`func NewUpdateDestinationRequestWithDefaults() *UpdateDestinationRequest`

NewUpdateDestinationRequestWithDefaults instantiates a new UpdateDestinationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedDestinationName

`func (o *UpdateDestinationRequest) GetUpdatedDestinationName() string`

GetUpdatedDestinationName returns the UpdatedDestinationName field if non-nil, zero value otherwise.

### GetUpdatedDestinationNameOk

`func (o *UpdateDestinationRequest) GetUpdatedDestinationNameOk() (*string, bool)`

GetUpdatedDestinationNameOk returns a tuple with the UpdatedDestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDestinationName

`func (o *UpdateDestinationRequest) SetUpdatedDestinationName(v string)`

SetUpdatedDestinationName sets UpdatedDestinationName field to given value.


### GetUpdatedDestinationType

`func (o *UpdateDestinationRequest) GetUpdatedDestinationType() DestinationType`

GetUpdatedDestinationType returns the UpdatedDestinationType field if non-nil, zero value otherwise.

### GetUpdatedDestinationTypeOk

`func (o *UpdateDestinationRequest) GetUpdatedDestinationTypeOk() (*DestinationType, bool)`

GetUpdatedDestinationTypeOk returns a tuple with the UpdatedDestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDestinationType

`func (o *UpdateDestinationRequest) SetUpdatedDestinationType(v DestinationType)`

SetUpdatedDestinationType sets UpdatedDestinationType field to given value.


### GetUpdatedMerchantId

`func (o *UpdateDestinationRequest) GetUpdatedMerchantId() string`

GetUpdatedMerchantId returns the UpdatedMerchantId field if non-nil, zero value otherwise.

### GetUpdatedMerchantIdOk

`func (o *UpdateDestinationRequest) GetUpdatedMerchantIdOk() (*string, bool)`

GetUpdatedMerchantIdOk returns a tuple with the UpdatedMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedMerchantId

`func (o *UpdateDestinationRequest) SetUpdatedMerchantId(v string)`

SetUpdatedMerchantId sets UpdatedMerchantId field to given value.

### HasUpdatedMerchantId

`func (o *UpdateDestinationRequest) HasUpdatedMerchantId() bool`

HasUpdatedMerchantId returns a boolean if a field has been set.

### GetUpdatedCountry

`func (o *UpdateDestinationRequest) GetUpdatedCountry() string`

GetUpdatedCountry returns the UpdatedCountry field if non-nil, zero value otherwise.

### GetUpdatedCountryOk

`func (o *UpdateDestinationRequest) GetUpdatedCountryOk() (*string, bool)`

GetUpdatedCountryOk returns a tuple with the UpdatedCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedCountry

`func (o *UpdateDestinationRequest) SetUpdatedCountry(v string)`

SetUpdatedCountry sets UpdatedCountry field to given value.

### HasUpdatedCountry

`func (o *UpdateDestinationRequest) HasUpdatedCountry() bool`

HasUpdatedCountry returns a boolean if a field has been set.

### GetUpdatedEmail

`func (o *UpdateDestinationRequest) GetUpdatedEmail() string`

GetUpdatedEmail returns the UpdatedEmail field if non-nil, zero value otherwise.

### GetUpdatedEmailOk

`func (o *UpdateDestinationRequest) GetUpdatedEmailOk() (*string, bool)`

GetUpdatedEmailOk returns a tuple with the UpdatedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEmail

`func (o *UpdateDestinationRequest) SetUpdatedEmail(v string)`

SetUpdatedEmail sets UpdatedEmail field to given value.

### HasUpdatedEmail

`func (o *UpdateDestinationRequest) HasUpdatedEmail() bool`

HasUpdatedEmail returns a boolean if a field has been set.

### GetUpdatedContactAddress

`func (o *UpdateDestinationRequest) GetUpdatedContactAddress() string`

GetUpdatedContactAddress returns the UpdatedContactAddress field if non-nil, zero value otherwise.

### GetUpdatedContactAddressOk

`func (o *UpdateDestinationRequest) GetUpdatedContactAddressOk() (*string, bool)`

GetUpdatedContactAddressOk returns a tuple with the UpdatedContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedContactAddress

`func (o *UpdateDestinationRequest) SetUpdatedContactAddress(v string)`

SetUpdatedContactAddress sets UpdatedContactAddress field to given value.

### HasUpdatedContactAddress

`func (o *UpdateDestinationRequest) HasUpdatedContactAddress() bool`

HasUpdatedContactAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


