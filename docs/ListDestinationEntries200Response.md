# ListDestinationEntries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletAddresses** | Pointer to [**[]DestinationWalletAddressDetail**](DestinationWalletAddressDetail.md) |  | [optional] 
**BankAccounts** | Pointer to [**[]DestinationBankAccountDetail**](DestinationBankAccountDetail.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListDestinationEntries200Response

`func NewListDestinationEntries200Response() *ListDestinationEntries200Response`

NewListDestinationEntries200Response instantiates a new ListDestinationEntries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDestinationEntries200ResponseWithDefaults

`func NewListDestinationEntries200ResponseWithDefaults() *ListDestinationEntries200Response`

NewListDestinationEntries200ResponseWithDefaults instantiates a new ListDestinationEntries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletAddresses

`func (o *ListDestinationEntries200Response) GetWalletAddresses() []DestinationWalletAddressDetail`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *ListDestinationEntries200Response) GetWalletAddressesOk() (*[]DestinationWalletAddressDetail, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *ListDestinationEntries200Response) SetWalletAddresses(v []DestinationWalletAddressDetail)`

SetWalletAddresses sets WalletAddresses field to given value.

### HasWalletAddresses

`func (o *ListDestinationEntries200Response) HasWalletAddresses() bool`

HasWalletAddresses returns a boolean if a field has been set.

### GetBankAccounts

`func (o *ListDestinationEntries200Response) GetBankAccounts() []DestinationBankAccountDetail`

GetBankAccounts returns the BankAccounts field if non-nil, zero value otherwise.

### GetBankAccountsOk

`func (o *ListDestinationEntries200Response) GetBankAccountsOk() (*[]DestinationBankAccountDetail, bool)`

GetBankAccountsOk returns a tuple with the BankAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccounts

`func (o *ListDestinationEntries200Response) SetBankAccounts(v []DestinationBankAccountDetail)`

SetBankAccounts sets BankAccounts field to given value.

### HasBankAccounts

`func (o *ListDestinationEntries200Response) HasBankAccounts() bool`

HasBankAccounts returns a boolean if a field has been set.

### GetPagination

`func (o *ListDestinationEntries200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDestinationEntries200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDestinationEntries200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListDestinationEntries200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


