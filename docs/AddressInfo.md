# AddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | **string** | The address ID. | 
**Address** | **string** | The wallet address. | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
**Memo** | Pointer to **string** | The memo code. | [optional] 
**Path** | Pointer to **string** | The derivation path of the address. | [optional] 
**Encoding** | Pointer to [**AddressEncoding**](AddressEncoding.md) |  | [optional] 
**Pubkey** | Pointer to **string** | The public key of the address. | [optional] 

## Methods

### NewAddressInfo

`func NewAddressInfo(addressId string, address string, tokenId string, ) *AddressInfo`

NewAddressInfo instantiates a new AddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressInfoWithDefaults

`func NewAddressInfoWithDefaults() *AddressInfo`

NewAddressInfoWithDefaults instantiates a new AddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *AddressInfo) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *AddressInfo) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *AddressInfo) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.


### GetAddress

`func (o *AddressInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTokenId

`func (o *AddressInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AddressInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AddressInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetMemo

`func (o *AddressInfo) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AddressInfo) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AddressInfo) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *AddressInfo) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPath

`func (o *AddressInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AddressInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AddressInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AddressInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetEncoding

`func (o *AddressInfo) GetEncoding() AddressEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AddressInfo) GetEncodingOk() (*AddressEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AddressInfo) SetEncoding(v AddressEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *AddressInfo) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPubkey

`func (o *AddressInfo) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *AddressInfo) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *AddressInfo) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *AddressInfo) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


