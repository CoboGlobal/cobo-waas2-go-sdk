# GenerateWalletAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
**Count** | **int32** | The number of addresses to create. | [default to 1]
**Encoding** | Pointer to [**AddressEncoding**](AddressEncoding.md) |  | [optional] 

## Methods

### NewGenerateWalletAddressRequest

`func NewGenerateWalletAddressRequest(tokenId string, count int32, ) *GenerateWalletAddressRequest`

NewGenerateWalletAddressRequest instantiates a new GenerateWalletAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateWalletAddressRequestWithDefaults

`func NewGenerateWalletAddressRequestWithDefaults() *GenerateWalletAddressRequest`

NewGenerateWalletAddressRequestWithDefaults instantiates a new GenerateWalletAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *GenerateWalletAddressRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GenerateWalletAddressRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GenerateWalletAddressRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetCount

`func (o *GenerateWalletAddressRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GenerateWalletAddressRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GenerateWalletAddressRequest) SetCount(v int32)`

SetCount sets Count field to given value.


### GetEncoding

`func (o *GenerateWalletAddressRequest) GetEncoding() AddressEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *GenerateWalletAddressRequest) GetEncodingOk() (*AddressEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *GenerateWalletAddressRequest) SetEncoding(v AddressEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *GenerateWalletAddressRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


