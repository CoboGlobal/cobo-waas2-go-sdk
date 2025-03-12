# ChainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. | 
**Symbol** | Pointer to **string** | The chain symbol, which is the abbreviated name of a chain. | [optional] 
**IconUrl** | Pointer to **string** | The URL of the chain icon. | [optional] 
**ExplorerTxUrl** | Pointer to **string** | The transaction URL pattern on the blockchain explorer. You can use it to concatenate the transaction URLs. | [optional] 
**ExplorerAddressUrl** | Pointer to **string** | The address URL pattern on the blockchain explorer. You can use it to concatenate the address URLs. | [optional] 
**RequireMemo** | Pointer to **bool** | Whether the chain requires a memo. | [optional] 
**ConfirmingThreshold** | Pointer to **int32** | The number of confirmations required for an on-chain transaction, such as 64 for Ethereum. | [optional] 

## Methods

### NewChainInfo

`func NewChainInfo(chainId string, ) *ChainInfo`

NewChainInfo instantiates a new ChainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainInfoWithDefaults

`func NewChainInfoWithDefaults() *ChainInfo`

NewChainInfoWithDefaults instantiates a new ChainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *ChainInfo) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ChainInfo) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ChainInfo) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSymbol

`func (o *ChainInfo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ChainInfo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ChainInfo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ChainInfo) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetIconUrl

`func (o *ChainInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ChainInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ChainInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ChainInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetExplorerTxUrl

`func (o *ChainInfo) GetExplorerTxUrl() string`

GetExplorerTxUrl returns the ExplorerTxUrl field if non-nil, zero value otherwise.

### GetExplorerTxUrlOk

`func (o *ChainInfo) GetExplorerTxUrlOk() (*string, bool)`

GetExplorerTxUrlOk returns a tuple with the ExplorerTxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerTxUrl

`func (o *ChainInfo) SetExplorerTxUrl(v string)`

SetExplorerTxUrl sets ExplorerTxUrl field to given value.

### HasExplorerTxUrl

`func (o *ChainInfo) HasExplorerTxUrl() bool`

HasExplorerTxUrl returns a boolean if a field has been set.

### GetExplorerAddressUrl

`func (o *ChainInfo) GetExplorerAddressUrl() string`

GetExplorerAddressUrl returns the ExplorerAddressUrl field if non-nil, zero value otherwise.

### GetExplorerAddressUrlOk

`func (o *ChainInfo) GetExplorerAddressUrlOk() (*string, bool)`

GetExplorerAddressUrlOk returns a tuple with the ExplorerAddressUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerAddressUrl

`func (o *ChainInfo) SetExplorerAddressUrl(v string)`

SetExplorerAddressUrl sets ExplorerAddressUrl field to given value.

### HasExplorerAddressUrl

`func (o *ChainInfo) HasExplorerAddressUrl() bool`

HasExplorerAddressUrl returns a boolean if a field has been set.

### GetRequireMemo

`func (o *ChainInfo) GetRequireMemo() bool`

GetRequireMemo returns the RequireMemo field if non-nil, zero value otherwise.

### GetRequireMemoOk

`func (o *ChainInfo) GetRequireMemoOk() (*bool, bool)`

GetRequireMemoOk returns a tuple with the RequireMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMemo

`func (o *ChainInfo) SetRequireMemo(v bool)`

SetRequireMemo sets RequireMemo field to given value.

### HasRequireMemo

`func (o *ChainInfo) HasRequireMemo() bool`

HasRequireMemo returns a boolean if a field has been set.

### GetConfirmingThreshold

`func (o *ChainInfo) GetConfirmingThreshold() int32`

GetConfirmingThreshold returns the ConfirmingThreshold field if non-nil, zero value otherwise.

### GetConfirmingThresholdOk

`func (o *ChainInfo) GetConfirmingThresholdOk() (*int32, bool)`

GetConfirmingThresholdOk returns a tuple with the ConfirmingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmingThreshold

`func (o *ChainInfo) SetConfirmingThreshold(v int32)`

SetConfirmingThreshold sets ConfirmingThreshold field to given value.

### HasConfirmingThreshold

`func (o *ChainInfo) HasConfirmingThreshold() bool`

HasConfirmingThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


