# TokensEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data. | 
**Tokens** | [**[]TokenInfo**](TokenInfo.md) | The enabled tokens. | 

## Methods

### NewTokensEventData

`func NewTokensEventData(dataType string, tokens []TokenInfo, ) *TokensEventData`

NewTokensEventData instantiates a new TokensEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensEventDataWithDefaults

`func NewTokensEventDataWithDefaults() *TokensEventData`

NewTokensEventDataWithDefaults instantiates a new TokensEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *TokensEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TokensEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TokensEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetTokens

`func (o *TokensEventData) GetTokens() []TokenInfo`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TokensEventData) GetTokensOk() (*[]TokenInfo, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TokensEventData) SetTokens(v []TokenInfo)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


