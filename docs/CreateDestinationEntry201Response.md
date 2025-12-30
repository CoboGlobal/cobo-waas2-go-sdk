# CreateDestinationEntry201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** | The destination ID. | 
**WalletAddresses** | Pointer to [**[]WalletAddress**](WalletAddress.md) | The wallet addresses of the destination. | [optional] 
**BankAccounts** | Pointer to [**[]DestinationBankAccount**](DestinationBankAccount.md) | The bank accounts of the destination. | [optional] 

## Methods

### NewCreateDestinationEntry201Response

`func NewCreateDestinationEntry201Response(destinationId string, ) *CreateDestinationEntry201Response`

NewCreateDestinationEntry201Response instantiates a new CreateDestinationEntry201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationEntry201ResponseWithDefaults

`func NewCreateDestinationEntry201ResponseWithDefaults() *CreateDestinationEntry201Response`

NewCreateDestinationEntry201ResponseWithDefaults instantiates a new CreateDestinationEntry201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *CreateDestinationEntry201Response) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *CreateDestinationEntry201Response) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *CreateDestinationEntry201Response) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetWalletAddresses

`func (o *CreateDestinationEntry201Response) GetWalletAddresses() []WalletAddress`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *CreateDestinationEntry201Response) GetWalletAddressesOk() (*[]WalletAddress, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *CreateDestinationEntry201Response) SetWalletAddresses(v []WalletAddress)`

SetWalletAddresses sets WalletAddresses field to given value.

### HasWalletAddresses

`func (o *CreateDestinationEntry201Response) HasWalletAddresses() bool`

HasWalletAddresses returns a boolean if a field has been set.

### GetBankAccounts

`func (o *CreateDestinationEntry201Response) GetBankAccounts() []DestinationBankAccount`

GetBankAccounts returns the BankAccounts field if non-nil, zero value otherwise.

### GetBankAccountsOk

`func (o *CreateDestinationEntry201Response) GetBankAccountsOk() (*[]DestinationBankAccount, bool)`

GetBankAccountsOk returns a tuple with the BankAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccounts

`func (o *CreateDestinationEntry201Response) SetBankAccounts(v []DestinationBankAccount)`

SetBankAccounts sets BankAccounts field to given value.

### HasBankAccounts

`func (o *CreateDestinationEntry201Response) HasBankAccounts() bool`

HasBankAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


