# TransactionTokeApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID. Unique in all chains scope. | 
**ChainId** | **string** | The blockchain on which the token operates. | 
**Symbol** | Pointer to **string** | The token symbol. | [optional] 
**Description** | Pointer to **string** | The token description. | [optional] 
**IconUrl** | Pointer to **string** | The URL of the token icon. | [optional] 
**TokenAddress** | Pointer to **string** | The token address, if applicable. | [optional] 
**AssetId** | Pointer to **string** | The asset ID, which is used to group the balances of the correponding tokens. For example, if you have $1,000 worth of ETH_USDT and $2,000 worth of TRON_USDT, the balance of your USDT assets will be $3,000. | [optional] 
**Amount** | Pointer to **float32** | Transaction value (Note that this is an absolute value. If you trade 1.5 BTC, then the value is 1.5)  | [optional] 
**Spender** | Pointer to **string** | Spender address | [optional] 

## Methods

### NewTransactionTokeApproval

`func NewTransactionTokeApproval(tokenId string, chainId string, ) *TransactionTokeApproval`

NewTransactionTokeApproval instantiates a new TransactionTokeApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTokeApprovalWithDefaults

`func NewTransactionTokeApprovalWithDefaults() *TransactionTokeApproval`

NewTransactionTokeApprovalWithDefaults instantiates a new TransactionTokeApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TransactionTokeApproval) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionTokeApproval) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionTokeApproval) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *TransactionTokeApproval) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TransactionTokeApproval) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TransactionTokeApproval) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSymbol

`func (o *TransactionTokeApproval) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransactionTokeApproval) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransactionTokeApproval) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TransactionTokeApproval) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionTokeApproval) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionTokeApproval) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionTokeApproval) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionTokeApproval) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *TransactionTokeApproval) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *TransactionTokeApproval) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *TransactionTokeApproval) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *TransactionTokeApproval) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetTokenAddress

`func (o *TransactionTokeApproval) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TransactionTokeApproval) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TransactionTokeApproval) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TransactionTokeApproval) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetAssetId

`func (o *TransactionTokeApproval) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransactionTokeApproval) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransactionTokeApproval) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *TransactionTokeApproval) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionTokeApproval) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionTokeApproval) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionTokeApproval) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionTokeApproval) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSpender

`func (o *TransactionTokeApproval) GetSpender() string`

GetSpender returns the Spender field if non-nil, zero value otherwise.

### GetSpenderOk

`func (o *TransactionTokeApproval) GetSpenderOk() (*string, bool)`

GetSpenderOk returns a tuple with the Spender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpender

`func (o *TransactionTokeApproval) SetSpender(v string)`

SetSpender sets Spender field to given value.

### HasSpender

`func (o *TransactionTokeApproval) HasSpender() bool`

HasSpender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


