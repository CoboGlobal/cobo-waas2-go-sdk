# CreateDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationName** | **string** | The destination name. | 
**DestinationType** | [**DestinationType**](DestinationType.md) |  | 
**WalletAddresses** | Pointer to [**[]CreateWalletAddress**](CreateWalletAddress.md) | The wallet addresses of the destination. | [optional] 
**BankAccounts** | Pointer to [**[]CreateDestinationBankAccount**](CreateDestinationBankAccount.md) | The bank accounts of the destination. | [optional] 
**MerchantId** | Pointer to **string** | The ID of the merchant linked to the destination. | [optional] 
**Country** | Pointer to **string** | The country of the destination, in ISO 3166-1 alpha-3 format. | [optional] 
**Email** | Pointer to **string** | The email of the destination. | [optional] 
**ContactAddress** | Pointer to **string** | The contact address of the destination. | [optional] 

## Methods

### NewCreateDestinationRequest

`func NewCreateDestinationRequest(destinationName string, destinationType DestinationType, ) *CreateDestinationRequest`

NewCreateDestinationRequest instantiates a new CreateDestinationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationRequestWithDefaults

`func NewCreateDestinationRequestWithDefaults() *CreateDestinationRequest`

NewCreateDestinationRequestWithDefaults instantiates a new CreateDestinationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationName

`func (o *CreateDestinationRequest) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *CreateDestinationRequest) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *CreateDestinationRequest) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.


### GetDestinationType

`func (o *CreateDestinationRequest) GetDestinationType() DestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *CreateDestinationRequest) GetDestinationTypeOk() (*DestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *CreateDestinationRequest) SetDestinationType(v DestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetWalletAddresses

`func (o *CreateDestinationRequest) GetWalletAddresses() []CreateWalletAddress`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *CreateDestinationRequest) GetWalletAddressesOk() (*[]CreateWalletAddress, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *CreateDestinationRequest) SetWalletAddresses(v []CreateWalletAddress)`

SetWalletAddresses sets WalletAddresses field to given value.

### HasWalletAddresses

`func (o *CreateDestinationRequest) HasWalletAddresses() bool`

HasWalletAddresses returns a boolean if a field has been set.

### GetBankAccounts

`func (o *CreateDestinationRequest) GetBankAccounts() []CreateDestinationBankAccount`

GetBankAccounts returns the BankAccounts field if non-nil, zero value otherwise.

### GetBankAccountsOk

`func (o *CreateDestinationRequest) GetBankAccountsOk() (*[]CreateDestinationBankAccount, bool)`

GetBankAccountsOk returns a tuple with the BankAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccounts

`func (o *CreateDestinationRequest) SetBankAccounts(v []CreateDestinationBankAccount)`

SetBankAccounts sets BankAccounts field to given value.

### HasBankAccounts

`func (o *CreateDestinationRequest) HasBankAccounts() bool`

HasBankAccounts returns a boolean if a field has been set.

### GetMerchantId

`func (o *CreateDestinationRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateDestinationRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateDestinationRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *CreateDestinationRequest) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetCountry

`func (o *CreateDestinationRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateDestinationRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateDestinationRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateDestinationRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *CreateDestinationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateDestinationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateDestinationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateDestinationRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactAddress

`func (o *CreateDestinationRequest) GetContactAddress() string`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *CreateDestinationRequest) GetContactAddressOk() (*string, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *CreateDestinationRequest) SetContactAddress(v string)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *CreateDestinationRequest) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


