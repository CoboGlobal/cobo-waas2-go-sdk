# AddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](/v2/api-references/wallets/list-enabled-chains). | 
**Memo** | Pointer to **string** | The memo code. | [optional] 
**Path** | Pointer to **string** | The derivation path of the address. This property applies to MPC Wallets only. To learn the meaning of each level in the path, see [Path levels](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki#path-levels). | [optional] 
**Encoding** | Pointer to [**AddressEncoding**](AddressEncoding.md) |  | [optional] 
**Pubkey** | Pointer to **string** | The public key of the address. This property applies to MPC Wallets only. | [optional] 
**XOnlyPubkey** | Pointer to **string** | The 32-byte x-only public key in hexadecimal format after tweaking. | [optional] 
**TaprootScriptTreeHash** | Pointer to **string** | The information about the new address. | [optional] 
**TaprootInternalAddress** | Pointer to **string** | The Taproot address before tweaking. | [optional] 

## Methods

### NewAddressInfo

`func NewAddressInfo(address string, chainId string, ) *AddressInfo`

NewAddressInfo instantiates a new AddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressInfoWithDefaults

`func NewAddressInfoWithDefaults() *AddressInfo`

NewAddressInfoWithDefaults instantiates a new AddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetChainId

`func (o *AddressInfo) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *AddressInfo) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *AddressInfo) SetChainId(v string)`

SetChainId sets ChainId field to given value.


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

### GetXOnlyPubkey

`func (o *AddressInfo) GetXOnlyPubkey() string`

GetXOnlyPubkey returns the XOnlyPubkey field if non-nil, zero value otherwise.

### GetXOnlyPubkeyOk

`func (o *AddressInfo) GetXOnlyPubkeyOk() (*string, bool)`

GetXOnlyPubkeyOk returns a tuple with the XOnlyPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXOnlyPubkey

`func (o *AddressInfo) SetXOnlyPubkey(v string)`

SetXOnlyPubkey sets XOnlyPubkey field to given value.

### HasXOnlyPubkey

`func (o *AddressInfo) HasXOnlyPubkey() bool`

HasXOnlyPubkey returns a boolean if a field has been set.

### GetTaprootScriptTreeHash

`func (o *AddressInfo) GetTaprootScriptTreeHash() string`

GetTaprootScriptTreeHash returns the TaprootScriptTreeHash field if non-nil, zero value otherwise.

### GetTaprootScriptTreeHashOk

`func (o *AddressInfo) GetTaprootScriptTreeHashOk() (*string, bool)`

GetTaprootScriptTreeHashOk returns a tuple with the TaprootScriptTreeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaprootScriptTreeHash

`func (o *AddressInfo) SetTaprootScriptTreeHash(v string)`

SetTaprootScriptTreeHash sets TaprootScriptTreeHash field to given value.

### HasTaprootScriptTreeHash

`func (o *AddressInfo) HasTaprootScriptTreeHash() bool`

HasTaprootScriptTreeHash returns a boolean if a field has been set.

### GetTaprootInternalAddress

`func (o *AddressInfo) GetTaprootInternalAddress() string`

GetTaprootInternalAddress returns the TaprootInternalAddress field if non-nil, zero value otherwise.

### GetTaprootInternalAddressOk

`func (o *AddressInfo) GetTaprootInternalAddressOk() (*string, bool)`

GetTaprootInternalAddressOk returns a tuple with the TaprootInternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaprootInternalAddress

`func (o *AddressInfo) SetTaprootInternalAddress(v string)`

SetTaprootInternalAddress sets TaprootInternalAddress field to given value.

### HasTaprootInternalAddress

`func (o *AddressInfo) HasTaprootInternalAddress() bool`

HasTaprootInternalAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


