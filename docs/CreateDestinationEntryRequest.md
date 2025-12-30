# CreateDestinationEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** | The destination ID. | 
**WalletAddresses** | Pointer to [**[]CreateWalletAddress**](CreateWalletAddress.md) | The wallet addresses of the destination. | [optional] 
**BankAccounts** | Pointer to [**[]CreateDestinationBankAccount**](CreateDestinationBankAccount.md) | The bank accounts of the destination. | [optional] 

## Methods

### NewCreateDestinationEntryRequest

`func NewCreateDestinationEntryRequest(destinationId string, ) *CreateDestinationEntryRequest`

NewCreateDestinationEntryRequest instantiates a new CreateDestinationEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationEntryRequestWithDefaults

`func NewCreateDestinationEntryRequestWithDefaults() *CreateDestinationEntryRequest`

NewCreateDestinationEntryRequestWithDefaults instantiates a new CreateDestinationEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *CreateDestinationEntryRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *CreateDestinationEntryRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *CreateDestinationEntryRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetWalletAddresses

`func (o *CreateDestinationEntryRequest) GetWalletAddresses() []CreateWalletAddress`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *CreateDestinationEntryRequest) GetWalletAddressesOk() (*[]CreateWalletAddress, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *CreateDestinationEntryRequest) SetWalletAddresses(v []CreateWalletAddress)`

SetWalletAddresses sets WalletAddresses field to given value.

### HasWalletAddresses

`func (o *CreateDestinationEntryRequest) HasWalletAddresses() bool`

HasWalletAddresses returns a boolean if a field has been set.

### GetBankAccounts

`func (o *CreateDestinationEntryRequest) GetBankAccounts() []CreateDestinationBankAccount`

GetBankAccounts returns the BankAccounts field if non-nil, zero value otherwise.

### GetBankAccountsOk

`func (o *CreateDestinationEntryRequest) GetBankAccountsOk() (*[]CreateDestinationBankAccount, bool)`

GetBankAccountsOk returns a tuple with the BankAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccounts

`func (o *CreateDestinationEntryRequest) SetBankAccounts(v []CreateDestinationBankAccount)`

SetBankAccounts sets BankAccounts field to given value.

### HasBankAccounts

`func (o *CreateDestinationEntryRequest) HasBankAccounts() bool`

HasBankAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


