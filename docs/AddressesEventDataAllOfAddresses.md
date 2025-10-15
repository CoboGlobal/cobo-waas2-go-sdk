# AddressesEventDataAllOfAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
**Memo** | Pointer to **string** | The memo code. | [optional] 
**Path** | Pointer to **string** | The derivation path of the address. This property applies to MPC Wallets only. To learn the meaning of each level in the path, see [Path levels](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki#path-levels). | [optional] 
**Encoding** | Pointer to [**AddressEncoding**](AddressEncoding.md) |  | [optional] 
**Pubkey** | Pointer to **string** | The public key of the address. This property applies to MPC Wallets only. | [optional] 
**XOnlyPubkey** | Pointer to **string** | The 32-byte x-only public key in hexadecimal format after tweaking. | [optional] 
**RootPubkey** | Pointer to **string** | The root public key of the address. This property applies to MPC Wallets only. | [optional] 
**TaprootScriptTreeHash** | Pointer to **string** | The information about the new address. | [optional] 
**TaprootInternalAddress** | Pointer to **string** | The Taproot address before tweaking. | [optional] 
**StellarTrustedTokenIds** | Pointer to **[]string** | The list of token IDs for which this address has already established trustlines on the Stellar network. | [optional] 
**WalletId** | **string** | The wallet ID. | 

## Methods

### NewAddressesEventDataAllOfAddresses

`func NewAddressesEventDataAllOfAddresses(address string, chainId string, walletId string, ) *AddressesEventDataAllOfAddresses`

NewAddressesEventDataAllOfAddresses instantiates a new AddressesEventDataAllOfAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressesEventDataAllOfAddressesWithDefaults

`func NewAddressesEventDataAllOfAddressesWithDefaults() *AddressesEventDataAllOfAddresses`

NewAddressesEventDataAllOfAddressesWithDefaults instantiates a new AddressesEventDataAllOfAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressesEventDataAllOfAddresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressesEventDataAllOfAddresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressesEventDataAllOfAddresses) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *AddressesEventDataAllOfAddresses) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *AddressesEventDataAllOfAddresses) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *AddressesEventDataAllOfAddresses) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetMemo

`func (o *AddressesEventDataAllOfAddresses) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AddressesEventDataAllOfAddresses) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AddressesEventDataAllOfAddresses) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *AddressesEventDataAllOfAddresses) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPath

`func (o *AddressesEventDataAllOfAddresses) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AddressesEventDataAllOfAddresses) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AddressesEventDataAllOfAddresses) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AddressesEventDataAllOfAddresses) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetEncoding

`func (o *AddressesEventDataAllOfAddresses) GetEncoding() AddressEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AddressesEventDataAllOfAddresses) GetEncodingOk() (*AddressEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AddressesEventDataAllOfAddresses) SetEncoding(v AddressEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *AddressesEventDataAllOfAddresses) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPubkey

`func (o *AddressesEventDataAllOfAddresses) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *AddressesEventDataAllOfAddresses) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *AddressesEventDataAllOfAddresses) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *AddressesEventDataAllOfAddresses) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetXOnlyPubkey

`func (o *AddressesEventDataAllOfAddresses) GetXOnlyPubkey() string`

GetXOnlyPubkey returns the XOnlyPubkey field if non-nil, zero value otherwise.

### GetXOnlyPubkeyOk

`func (o *AddressesEventDataAllOfAddresses) GetXOnlyPubkeyOk() (*string, bool)`

GetXOnlyPubkeyOk returns a tuple with the XOnlyPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXOnlyPubkey

`func (o *AddressesEventDataAllOfAddresses) SetXOnlyPubkey(v string)`

SetXOnlyPubkey sets XOnlyPubkey field to given value.

### HasXOnlyPubkey

`func (o *AddressesEventDataAllOfAddresses) HasXOnlyPubkey() bool`

HasXOnlyPubkey returns a boolean if a field has been set.

### GetRootPubkey

`func (o *AddressesEventDataAllOfAddresses) GetRootPubkey() string`

GetRootPubkey returns the RootPubkey field if non-nil, zero value otherwise.

### GetRootPubkeyOk

`func (o *AddressesEventDataAllOfAddresses) GetRootPubkeyOk() (*string, bool)`

GetRootPubkeyOk returns a tuple with the RootPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPubkey

`func (o *AddressesEventDataAllOfAddresses) SetRootPubkey(v string)`

SetRootPubkey sets RootPubkey field to given value.

### HasRootPubkey

`func (o *AddressesEventDataAllOfAddresses) HasRootPubkey() bool`

HasRootPubkey returns a boolean if a field has been set.

### GetTaprootScriptTreeHash

`func (o *AddressesEventDataAllOfAddresses) GetTaprootScriptTreeHash() string`

GetTaprootScriptTreeHash returns the TaprootScriptTreeHash field if non-nil, zero value otherwise.

### GetTaprootScriptTreeHashOk

`func (o *AddressesEventDataAllOfAddresses) GetTaprootScriptTreeHashOk() (*string, bool)`

GetTaprootScriptTreeHashOk returns a tuple with the TaprootScriptTreeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaprootScriptTreeHash

`func (o *AddressesEventDataAllOfAddresses) SetTaprootScriptTreeHash(v string)`

SetTaprootScriptTreeHash sets TaprootScriptTreeHash field to given value.

### HasTaprootScriptTreeHash

`func (o *AddressesEventDataAllOfAddresses) HasTaprootScriptTreeHash() bool`

HasTaprootScriptTreeHash returns a boolean if a field has been set.

### GetTaprootInternalAddress

`func (o *AddressesEventDataAllOfAddresses) GetTaprootInternalAddress() string`

GetTaprootInternalAddress returns the TaprootInternalAddress field if non-nil, zero value otherwise.

### GetTaprootInternalAddressOk

`func (o *AddressesEventDataAllOfAddresses) GetTaprootInternalAddressOk() (*string, bool)`

GetTaprootInternalAddressOk returns a tuple with the TaprootInternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaprootInternalAddress

`func (o *AddressesEventDataAllOfAddresses) SetTaprootInternalAddress(v string)`

SetTaprootInternalAddress sets TaprootInternalAddress field to given value.

### HasTaprootInternalAddress

`func (o *AddressesEventDataAllOfAddresses) HasTaprootInternalAddress() bool`

HasTaprootInternalAddress returns a boolean if a field has been set.

### GetStellarTrustedTokenIds

`func (o *AddressesEventDataAllOfAddresses) GetStellarTrustedTokenIds() []string`

GetStellarTrustedTokenIds returns the StellarTrustedTokenIds field if non-nil, zero value otherwise.

### GetStellarTrustedTokenIdsOk

`func (o *AddressesEventDataAllOfAddresses) GetStellarTrustedTokenIdsOk() (*[]string, bool)`

GetStellarTrustedTokenIdsOk returns a tuple with the StellarTrustedTokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStellarTrustedTokenIds

`func (o *AddressesEventDataAllOfAddresses) SetStellarTrustedTokenIds(v []string)`

SetStellarTrustedTokenIds sets StellarTrustedTokenIds field to given value.

### HasStellarTrustedTokenIds

`func (o *AddressesEventDataAllOfAddresses) HasStellarTrustedTokenIds() bool`

HasStellarTrustedTokenIds returns a boolean if a field has been set.

### GetWalletId

`func (o *AddressesEventDataAllOfAddresses) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AddressesEventDataAllOfAddresses) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AddressesEventDataAllOfAddresses) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


