# AddWalletAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID. | 
**Count** | **int32** | The number of addresses to create. | [default to 1]
**Encoding** | Pointer to [**AddressEncoding**](AddressEncoding.md) |  | [optional] 

## Methods

### NewAddWalletAddressRequest

`func NewAddWalletAddressRequest(tokenId string, count int32, ) *AddWalletAddressRequest`

NewAddWalletAddressRequest instantiates a new AddWalletAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWalletAddressRequestWithDefaults

`func NewAddWalletAddressRequestWithDefaults() *AddWalletAddressRequest`

NewAddWalletAddressRequestWithDefaults instantiates a new AddWalletAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *AddWalletAddressRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AddWalletAddressRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AddWalletAddressRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetCount

`func (o *AddWalletAddressRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AddWalletAddressRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AddWalletAddressRequest) SetCount(v int32)`

SetCount sets Count field to given value.


### GetEncoding

`func (o *AddWalletAddressRequest) GetEncoding() AddressEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AddWalletAddressRequest) GetEncodingOk() (*AddressEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AddWalletAddressRequest) SetEncoding(v AddressEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *AddWalletAddressRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


