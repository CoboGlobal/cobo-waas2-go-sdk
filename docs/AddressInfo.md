# AddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | **string** |  | 
**AddressStr** | **string** | Then blockchain address | 
**TokenId** | **string** | ID of the token. Unique in all chains scope. | 
**Memo** | Pointer to **string** | From address memo | [optional] 
**IsInternal** | Pointer to **bool** | Ture if the address is in same fund pool, False otherwise | [optional] [default to false]

## Methods

### NewAddressInfo

`func NewAddressInfo(addressId string, addressStr string, tokenId string, ) *AddressInfo`

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


### GetAddressStr

`func (o *AddressInfo) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *AddressInfo) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *AddressInfo) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.


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

### GetIsInternal

`func (o *AddressInfo) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *AddressInfo) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *AddressInfo) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.

### HasIsInternal

`func (o *AddressInfo) HasIsInternal() bool`

HasIsInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


