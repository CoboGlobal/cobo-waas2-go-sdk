# ListTokenBalancesForFeeStation200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
**ChainId** | Pointer to **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | [optional] 
**Balance** | [**Balance**](Balance.md) |  | 

## Methods

### NewListTokenBalancesForFeeStation200ResponseDataInner

`func NewListTokenBalancesForFeeStation200ResponseDataInner(tokenId string, balance Balance, ) *ListTokenBalancesForFeeStation200ResponseDataInner`

NewListTokenBalancesForFeeStation200ResponseDataInner instantiates a new ListTokenBalancesForFeeStation200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTokenBalancesForFeeStation200ResponseDataInnerWithDefaults

`func NewListTokenBalancesForFeeStation200ResponseDataInnerWithDefaults() *ListTokenBalancesForFeeStation200ResponseDataInner`

NewListTokenBalancesForFeeStation200ResponseDataInnerWithDefaults instantiates a new ListTokenBalancesForFeeStation200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetBalance

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ListTokenBalancesForFeeStation200ResponseDataInner) SetBalance(v Balance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


