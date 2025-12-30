# GetDestinationEntry200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletAddress** | Pointer to [**WalletAddress**](WalletAddress.md) |  | [optional] 
**BankAccount** | Pointer to [**DestinationBankAccountDetail**](DestinationBankAccountDetail.md) |  | [optional] 

## Methods

### NewGetDestinationEntry200Response

`func NewGetDestinationEntry200Response() *GetDestinationEntry200Response`

NewGetDestinationEntry200Response instantiates a new GetDestinationEntry200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDestinationEntry200ResponseWithDefaults

`func NewGetDestinationEntry200ResponseWithDefaults() *GetDestinationEntry200Response`

NewGetDestinationEntry200ResponseWithDefaults instantiates a new GetDestinationEntry200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletAddress

`func (o *GetDestinationEntry200Response) GetWalletAddress() WalletAddress`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *GetDestinationEntry200Response) GetWalletAddressOk() (*WalletAddress, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *GetDestinationEntry200Response) SetWalletAddress(v WalletAddress)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *GetDestinationEntry200Response) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetBankAccount

`func (o *GetDestinationEntry200Response) GetBankAccount() DestinationBankAccountDetail`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *GetDestinationEntry200Response) GetBankAccountOk() (*DestinationBankAccountDetail, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *GetDestinationEntry200Response) SetBankAccount(v DestinationBankAccountDetail)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *GetDestinationEntry200Response) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


