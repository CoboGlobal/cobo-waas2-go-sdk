# SignMessageSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | Pointer to **string** | Unique id of the wallet to sign message. | [optional] 
**AddressStr** | Pointer to **string** | From address | [optional] 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewSignMessageSource

`func NewSignMessageSource() *SignMessageSource`

NewSignMessageSource instantiates a new SignMessageSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignMessageSourceWithDefaults

`func NewSignMessageSourceWithDefaults() *SignMessageSource`

NewSignMessageSourceWithDefaults instantiates a new SignMessageSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *SignMessageSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SignMessageSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SignMessageSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *SignMessageSource) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetAddressStr

`func (o *SignMessageSource) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *SignMessageSource) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *SignMessageSource) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.

### HasAddressStr

`func (o *SignMessageSource) HasAddressStr() bool`

HasAddressStr returns a boolean if a field has been set.

### GetMpcUsedKeyGroup

`func (o *SignMessageSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *SignMessageSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *SignMessageSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *SignMessageSource) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


