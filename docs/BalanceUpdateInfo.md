# BalanceUpdateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
**Address** | **string** | The wallet address. | 
**WalletUuid** | **string** | The wallet ID. | 
**UpdatedTimestamp** | **int64** | The time when the balance updated, in Unix timestamp format, measured in milliseconds.  | 
**Balance** | [**Balance**](Balance.md) |  | 

## Methods

### NewBalanceUpdateInfo

`func NewBalanceUpdateInfo(tokenId string, address string, walletUuid string, updatedTimestamp int64, balance Balance, ) *BalanceUpdateInfo`

NewBalanceUpdateInfo instantiates a new BalanceUpdateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceUpdateInfoWithDefaults

`func NewBalanceUpdateInfoWithDefaults() *BalanceUpdateInfo`

NewBalanceUpdateInfoWithDefaults instantiates a new BalanceUpdateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *BalanceUpdateInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BalanceUpdateInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BalanceUpdateInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAddress

`func (o *BalanceUpdateInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BalanceUpdateInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BalanceUpdateInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetWalletUuid

`func (o *BalanceUpdateInfo) GetWalletUuid() string`

GetWalletUuid returns the WalletUuid field if non-nil, zero value otherwise.

### GetWalletUuidOk

`func (o *BalanceUpdateInfo) GetWalletUuidOk() (*string, bool)`

GetWalletUuidOk returns a tuple with the WalletUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletUuid

`func (o *BalanceUpdateInfo) SetWalletUuid(v string)`

SetWalletUuid sets WalletUuid field to given value.


### GetUpdatedTimestamp

`func (o *BalanceUpdateInfo) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *BalanceUpdateInfo) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *BalanceUpdateInfo) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetBalance

`func (o *BalanceUpdateInfo) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceUpdateInfo) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceUpdateInfo) SetBalance(v Balance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


