# TokenBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | ID of the token. Unique in all chains scope. | 
**Balance** | [**TokenBalanceBalance**](TokenBalanceBalance.md) |  | 

## Methods

### NewTokenBalance

`func NewTokenBalance(tokenId string, balance TokenBalanceBalance, ) *TokenBalance`

NewTokenBalance instantiates a new TokenBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBalanceWithDefaults

`func NewTokenBalanceWithDefaults() *TokenBalance`

NewTokenBalanceWithDefaults instantiates a new TokenBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenBalance) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenBalance) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenBalance) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetBalance

`func (o *TokenBalance) GetBalance() TokenBalanceBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TokenBalance) GetBalanceOk() (*TokenBalanceBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TokenBalance) SetBalance(v TokenBalanceBalance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


