# TokenizationHoldingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The unique identifier of the wallet holding the token. | 
**WalletName** | Pointer to **string** | The name of the wallet. | [optional] 
**Address** | **string** | The address holding the token. | 
**Balance** | **string** | The token balance held by this address. | 
**AddressLabel** | Pointer to **string** | The label of the address. | [optional] 

## Methods

### NewTokenizationHoldingInfo

`func NewTokenizationHoldingInfo(walletId string, address string, balance string, ) *TokenizationHoldingInfo`

NewTokenizationHoldingInfo instantiates a new TokenizationHoldingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationHoldingInfoWithDefaults

`func NewTokenizationHoldingInfoWithDefaults() *TokenizationHoldingInfo`

NewTokenizationHoldingInfoWithDefaults instantiates a new TokenizationHoldingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *TokenizationHoldingInfo) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TokenizationHoldingInfo) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TokenizationHoldingInfo) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletName

`func (o *TokenizationHoldingInfo) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *TokenizationHoldingInfo) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *TokenizationHoldingInfo) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.

### HasWalletName

`func (o *TokenizationHoldingInfo) HasWalletName() bool`

HasWalletName returns a boolean if a field has been set.

### GetAddress

`func (o *TokenizationHoldingInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationHoldingInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationHoldingInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *TokenizationHoldingInfo) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TokenizationHoldingInfo) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TokenizationHoldingInfo) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetAddressLabel

`func (o *TokenizationHoldingInfo) GetAddressLabel() string`

GetAddressLabel returns the AddressLabel field if non-nil, zero value otherwise.

### GetAddressLabelOk

`func (o *TokenizationHoldingInfo) GetAddressLabelOk() (*string, bool)`

GetAddressLabelOk returns a tuple with the AddressLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLabel

`func (o *TokenizationHoldingInfo) SetAddressLabel(v string)`

SetAddressLabel sets AddressLabel field to given value.

### HasAddressLabel

`func (o *TokenizationHoldingInfo) HasAddressLabel() bool`

HasAddressLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


