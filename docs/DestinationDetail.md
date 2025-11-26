# DestinationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** | The destination ID. | 
**DestinationType** | [**DestinationType**](DestinationType.md) |  | 
**DestinationName** | **string** | The destination name. | 
**Country** | Pointer to **string** | The country of the destination, in ISO 3166-1 alpha-3 format. | [optional] 
**Email** | Pointer to **string** | The email of the destination. | [optional] 
**ContactAddress** | Pointer to **string** | The contact address of the destination. | [optional] 
**WalletAddresses** | Pointer to [**[]WalletAddress**](WalletAddress.md) | The wallet addresses of the destination. | [optional] 
**BankAccounts** | Pointer to [**[]DestinationBankAccount**](DestinationBankAccount.md) | The bank accounts of the destination. | [optional] 
**MerchantId** | Pointer to **string** | The ID of the merchant linked to the destination. | [optional] 
**CreatedTimestamp** | **int32** | The created time of the destination, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int32** | The updated time of the destination, represented as a UNIX timestamp in seconds. | 

## Methods

### NewDestinationDetail

`func NewDestinationDetail(destinationId string, destinationType DestinationType, destinationName string, createdTimestamp int32, updatedTimestamp int32, ) *DestinationDetail`

NewDestinationDetail instantiates a new DestinationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationDetailWithDefaults

`func NewDestinationDetailWithDefaults() *DestinationDetail`

NewDestinationDetailWithDefaults instantiates a new DestinationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *DestinationDetail) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DestinationDetail) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DestinationDetail) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetDestinationType

`func (o *DestinationDetail) GetDestinationType() DestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *DestinationDetail) GetDestinationTypeOk() (*DestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *DestinationDetail) SetDestinationType(v DestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetDestinationName

`func (o *DestinationDetail) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *DestinationDetail) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *DestinationDetail) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.


### GetCountry

`func (o *DestinationDetail) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DestinationDetail) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DestinationDetail) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DestinationDetail) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *DestinationDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DestinationDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DestinationDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DestinationDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactAddress

`func (o *DestinationDetail) GetContactAddress() string`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *DestinationDetail) GetContactAddressOk() (*string, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *DestinationDetail) SetContactAddress(v string)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *DestinationDetail) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.

### GetWalletAddresses

`func (o *DestinationDetail) GetWalletAddresses() []WalletAddress`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *DestinationDetail) GetWalletAddressesOk() (*[]WalletAddress, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *DestinationDetail) SetWalletAddresses(v []WalletAddress)`

SetWalletAddresses sets WalletAddresses field to given value.

### HasWalletAddresses

`func (o *DestinationDetail) HasWalletAddresses() bool`

HasWalletAddresses returns a boolean if a field has been set.

### GetBankAccounts

`func (o *DestinationDetail) GetBankAccounts() []DestinationBankAccount`

GetBankAccounts returns the BankAccounts field if non-nil, zero value otherwise.

### GetBankAccountsOk

`func (o *DestinationDetail) GetBankAccountsOk() (*[]DestinationBankAccount, bool)`

GetBankAccountsOk returns a tuple with the BankAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccounts

`func (o *DestinationDetail) SetBankAccounts(v []DestinationBankAccount)`

SetBankAccounts sets BankAccounts field to given value.

### HasBankAccounts

`func (o *DestinationDetail) HasBankAccounts() bool`

HasBankAccounts returns a boolean if a field has been set.

### GetMerchantId

`func (o *DestinationDetail) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *DestinationDetail) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *DestinationDetail) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *DestinationDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *DestinationDetail) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DestinationDetail) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DestinationDetail) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *DestinationDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DestinationDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DestinationDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


