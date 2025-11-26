# CreateCounterpartyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartyName** | **string** | The counterparty name. | 
**CounterpartyType** | [**CounterpartyType**](CounterpartyType.md) |  | 
**WalletAddresses** | [**[]CreateWalletAddress**](CreateWalletAddress.md) | The wallet addresses of the counterparty. | 
**Country** | Pointer to **string** | The country of the counterparty, in ISO 3166-1 alpha-3 format. | [optional] 
**Email** | Pointer to **string** | The email of the counterparty. | [optional] 
**ContactAddress** | Pointer to **string** | The contact address of the counterparty. | [optional] 

## Methods

### NewCreateCounterpartyRequest

`func NewCreateCounterpartyRequest(counterpartyName string, counterpartyType CounterpartyType, walletAddresses []CreateWalletAddress, ) *CreateCounterpartyRequest`

NewCreateCounterpartyRequest instantiates a new CreateCounterpartyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCounterpartyRequestWithDefaults

`func NewCreateCounterpartyRequestWithDefaults() *CreateCounterpartyRequest`

NewCreateCounterpartyRequestWithDefaults instantiates a new CreateCounterpartyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartyName

`func (o *CreateCounterpartyRequest) GetCounterpartyName() string`

GetCounterpartyName returns the CounterpartyName field if non-nil, zero value otherwise.

### GetCounterpartyNameOk

`func (o *CreateCounterpartyRequest) GetCounterpartyNameOk() (*string, bool)`

GetCounterpartyNameOk returns a tuple with the CounterpartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyName

`func (o *CreateCounterpartyRequest) SetCounterpartyName(v string)`

SetCounterpartyName sets CounterpartyName field to given value.


### GetCounterpartyType

`func (o *CreateCounterpartyRequest) GetCounterpartyType() CounterpartyType`

GetCounterpartyType returns the CounterpartyType field if non-nil, zero value otherwise.

### GetCounterpartyTypeOk

`func (o *CreateCounterpartyRequest) GetCounterpartyTypeOk() (*CounterpartyType, bool)`

GetCounterpartyTypeOk returns a tuple with the CounterpartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyType

`func (o *CreateCounterpartyRequest) SetCounterpartyType(v CounterpartyType)`

SetCounterpartyType sets CounterpartyType field to given value.


### GetWalletAddresses

`func (o *CreateCounterpartyRequest) GetWalletAddresses() []CreateWalletAddress`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *CreateCounterpartyRequest) GetWalletAddressesOk() (*[]CreateWalletAddress, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *CreateCounterpartyRequest) SetWalletAddresses(v []CreateWalletAddress)`

SetWalletAddresses sets WalletAddresses field to given value.


### GetCountry

`func (o *CreateCounterpartyRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateCounterpartyRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateCounterpartyRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateCounterpartyRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *CreateCounterpartyRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateCounterpartyRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateCounterpartyRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateCounterpartyRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactAddress

`func (o *CreateCounterpartyRequest) GetContactAddress() string`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *CreateCounterpartyRequest) GetContactAddressOk() (*string, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *CreateCounterpartyRequest) SetContactAddress(v string)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *CreateCounterpartyRequest) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


