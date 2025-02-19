# CreateAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
**Count** | **int32** | The number of addresses to create. This property will be ignored if you are generating tweaked Taproot addresses. | [default to 1]
**TaprootScriptTreeHashes** | Pointer to **[]string** | A list of script tree hashes used to generate a tweaked Taproot address. This property is required only if you want to generate tweaked Taproot addresses. | [optional] 
**TaprootInternalAddress** | Pointer to **string** | The original Taproot address to be tweaked. This property is required only if you want to generate tweaked Taproot addresses. | [optional] 
**Encoding** | Pointer to [**AddressEncoding**](AddressEncoding.md) |  | [optional] 

## Methods

### NewCreateAddressRequest

`func NewCreateAddressRequest(chainId string, count int32, ) *CreateAddressRequest`

NewCreateAddressRequest instantiates a new CreateAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAddressRequestWithDefaults

`func NewCreateAddressRequestWithDefaults() *CreateAddressRequest`

NewCreateAddressRequestWithDefaults instantiates a new CreateAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *CreateAddressRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateAddressRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateAddressRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetCount

`func (o *CreateAddressRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateAddressRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateAddressRequest) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTaprootScriptTreeHashes

`func (o *CreateAddressRequest) GetTaprootScriptTreeHashes() []string`

GetTaprootScriptTreeHashes returns the TaprootScriptTreeHashes field if non-nil, zero value otherwise.

### GetTaprootScriptTreeHashesOk

`func (o *CreateAddressRequest) GetTaprootScriptTreeHashesOk() (*[]string, bool)`

GetTaprootScriptTreeHashesOk returns a tuple with the TaprootScriptTreeHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaprootScriptTreeHashes

`func (o *CreateAddressRequest) SetTaprootScriptTreeHashes(v []string)`

SetTaprootScriptTreeHashes sets TaprootScriptTreeHashes field to given value.

### HasTaprootScriptTreeHashes

`func (o *CreateAddressRequest) HasTaprootScriptTreeHashes() bool`

HasTaprootScriptTreeHashes returns a boolean if a field has been set.

### GetTaprootInternalAddress

`func (o *CreateAddressRequest) GetTaprootInternalAddress() string`

GetTaprootInternalAddress returns the TaprootInternalAddress field if non-nil, zero value otherwise.

### GetTaprootInternalAddressOk

`func (o *CreateAddressRequest) GetTaprootInternalAddressOk() (*string, bool)`

GetTaprootInternalAddressOk returns a tuple with the TaprootInternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaprootInternalAddress

`func (o *CreateAddressRequest) SetTaprootInternalAddress(v string)`

SetTaprootInternalAddress sets TaprootInternalAddress field to given value.

### HasTaprootInternalAddress

`func (o *CreateAddressRequest) HasTaprootInternalAddress() bool`

HasTaprootInternalAddress returns a boolean if a field has been set.

### GetEncoding

`func (o *CreateAddressRequest) GetEncoding() AddressEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *CreateAddressRequest) GetEncodingOk() (*AddressEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *CreateAddressRequest) SetEncoding(v AddressEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *CreateAddressRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


