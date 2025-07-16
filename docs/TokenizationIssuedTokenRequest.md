# TokenizationIssuedTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain ID where the token will be issued. | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**TokenParams** | [**TokenizationIssueTokenParamsTokenParams**](TokenizationIssueTokenParamsTokenParams.md) |  | 
**AppInitiator** | Pointer to **string** | The address of the app initiator.  | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewTokenizationIssuedTokenRequest

`func NewTokenizationIssuedTokenRequest(chainId string, source TokenizationTokenOperationSource, tokenParams TokenizationIssueTokenParamsTokenParams, fee TransactionRequestFee, ) *TokenizationIssuedTokenRequest`

NewTokenizationIssuedTokenRequest instantiates a new TokenizationIssuedTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationIssuedTokenRequestWithDefaults

`func NewTokenizationIssuedTokenRequestWithDefaults() *TokenizationIssuedTokenRequest`

NewTokenizationIssuedTokenRequestWithDefaults instantiates a new TokenizationIssuedTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *TokenizationIssuedTokenRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenizationIssuedTokenRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenizationIssuedTokenRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *TokenizationIssuedTokenRequest) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationIssuedTokenRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationIssuedTokenRequest) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetTokenParams

`func (o *TokenizationIssuedTokenRequest) GetTokenParams() TokenizationIssueTokenParamsTokenParams`

GetTokenParams returns the TokenParams field if non-nil, zero value otherwise.

### GetTokenParamsOk

`func (o *TokenizationIssuedTokenRequest) GetTokenParamsOk() (*TokenizationIssueTokenParamsTokenParams, bool)`

GetTokenParamsOk returns a tuple with the TokenParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParams

`func (o *TokenizationIssuedTokenRequest) SetTokenParams(v TokenizationIssueTokenParamsTokenParams)`

SetTokenParams sets TokenParams field to given value.


### GetAppInitiator

`func (o *TokenizationIssuedTokenRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationIssuedTokenRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationIssuedTokenRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationIssuedTokenRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.

### GetFee

`func (o *TokenizationIssuedTokenRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TokenizationIssuedTokenRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TokenizationIssuedTokenRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


