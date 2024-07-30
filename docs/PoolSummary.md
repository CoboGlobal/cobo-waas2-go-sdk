# PoolSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique protocol id. | 
**ChainId** | **string** | The unique chain id. | 
**Protocol** | **string** | The name of the protocol. | 
**ProtocolIconUrl** | **string** | The URL of the protocol&#39;s icon. | 
**SupportedWalletTypes** | [**[]WalletType**](WalletType.md) | The list of available wallet types. | 
**SupportedWalletSubtypes** | [**[]WalletSubtype**](WalletSubtype.md) | The list of available wallet types. | 
**TokenId** | **string** | The unique token id. | 
**EstApr** | **float32** | The estimated APR. | 

## Methods

### NewPoolSummary

`func NewPoolSummary(id string, chainId string, protocol string, protocolIconUrl string, supportedWalletTypes []WalletType, supportedWalletSubtypes []WalletSubtype, tokenId string, estApr float32, ) *PoolSummary`

NewPoolSummary instantiates a new PoolSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolSummaryWithDefaults

`func NewPoolSummaryWithDefaults() *PoolSummary`

NewPoolSummaryWithDefaults instantiates a new PoolSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolSummary) SetId(v string)`

SetId sets Id field to given value.


### GetChainId

`func (o *PoolSummary) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PoolSummary) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PoolSummary) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetProtocol

`func (o *PoolSummary) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PoolSummary) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PoolSummary) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetProtocolIconUrl

`func (o *PoolSummary) GetProtocolIconUrl() string`

GetProtocolIconUrl returns the ProtocolIconUrl field if non-nil, zero value otherwise.

### GetProtocolIconUrlOk

`func (o *PoolSummary) GetProtocolIconUrlOk() (*string, bool)`

GetProtocolIconUrlOk returns a tuple with the ProtocolIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolIconUrl

`func (o *PoolSummary) SetProtocolIconUrl(v string)`

SetProtocolIconUrl sets ProtocolIconUrl field to given value.


### GetSupportedWalletTypes

`func (o *PoolSummary) GetSupportedWalletTypes() []WalletType`

GetSupportedWalletTypes returns the SupportedWalletTypes field if non-nil, zero value otherwise.

### GetSupportedWalletTypesOk

`func (o *PoolSummary) GetSupportedWalletTypesOk() (*[]WalletType, bool)`

GetSupportedWalletTypesOk returns a tuple with the SupportedWalletTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedWalletTypes

`func (o *PoolSummary) SetSupportedWalletTypes(v []WalletType)`

SetSupportedWalletTypes sets SupportedWalletTypes field to given value.


### GetSupportedWalletSubtypes

`func (o *PoolSummary) GetSupportedWalletSubtypes() []WalletSubtype`

GetSupportedWalletSubtypes returns the SupportedWalletSubtypes field if non-nil, zero value otherwise.

### GetSupportedWalletSubtypesOk

`func (o *PoolSummary) GetSupportedWalletSubtypesOk() (*[]WalletSubtype, bool)`

GetSupportedWalletSubtypesOk returns a tuple with the SupportedWalletSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedWalletSubtypes

`func (o *PoolSummary) SetSupportedWalletSubtypes(v []WalletSubtype)`

SetSupportedWalletSubtypes sets SupportedWalletSubtypes field to given value.


### GetTokenId

`func (o *PoolSummary) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PoolSummary) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PoolSummary) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetEstApr

`func (o *PoolSummary) GetEstApr() float32`

GetEstApr returns the EstApr field if non-nil, zero value otherwise.

### GetEstAprOk

`func (o *PoolSummary) GetEstAprOk() (*float32, bool)`

GetEstAprOk returns a tuple with the EstApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstApr

`func (o *PoolSummary) SetEstApr(v float32)`

SetEstApr sets EstApr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


