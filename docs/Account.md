# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The ID of the cryptocurrency. | 
**Address** | **string** | The token receiving address of the account. | 
**MerchantBalance** | **string** | The merchant account balance, as a decimal string. | 
**PspBalance** | **string** | The PSP account balance, as a decimal string. | 
**CreatedTimestamp** | Pointer to **int64** | The time when the account was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time when the account was updated, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewAccount

`func NewAccount(tokenId string, address string, merchantBalance string, pspBalance string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *Account) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Account) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Account) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAddress

`func (o *Account) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Account) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Account) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMerchantBalance

`func (o *Account) GetMerchantBalance() string`

GetMerchantBalance returns the MerchantBalance field if non-nil, zero value otherwise.

### GetMerchantBalanceOk

`func (o *Account) GetMerchantBalanceOk() (*string, bool)`

GetMerchantBalanceOk returns a tuple with the MerchantBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantBalance

`func (o *Account) SetMerchantBalance(v string)`

SetMerchantBalance sets MerchantBalance field to given value.


### GetPspBalance

`func (o *Account) GetPspBalance() string`

GetPspBalance returns the PspBalance field if non-nil, zero value otherwise.

### GetPspBalanceOk

`func (o *Account) GetPspBalanceOk() (*string, bool)`

GetPspBalanceOk returns a tuple with the PspBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspBalance

`func (o *Account) SetPspBalance(v string)`

SetPspBalance sets PspBalance field to given value.


### GetCreatedTimestamp

`func (o *Account) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Account) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Account) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Account) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Account) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Account) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Account) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *Account) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


