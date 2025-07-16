# TokenizationIssueTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain ID where the token will be issued. | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**TokenParams** | [**TokenizationIssueTokenParamsTokenParams**](TokenizationIssueTokenParamsTokenParams.md) |  | 

## Methods

### NewTokenizationIssueTokenParams

`func NewTokenizationIssueTokenParams(chainId string, source TokenizationTokenOperationSource, tokenParams TokenizationIssueTokenParamsTokenParams, ) *TokenizationIssueTokenParams`

NewTokenizationIssueTokenParams instantiates a new TokenizationIssueTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationIssueTokenParamsWithDefaults

`func NewTokenizationIssueTokenParamsWithDefaults() *TokenizationIssueTokenParams`

NewTokenizationIssueTokenParamsWithDefaults instantiates a new TokenizationIssueTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *TokenizationIssueTokenParams) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenizationIssueTokenParams) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenizationIssueTokenParams) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *TokenizationIssueTokenParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationIssueTokenParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationIssueTokenParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetTokenParams

`func (o *TokenizationIssueTokenParams) GetTokenParams() TokenizationIssueTokenParamsTokenParams`

GetTokenParams returns the TokenParams field if non-nil, zero value otherwise.

### GetTokenParamsOk

`func (o *TokenizationIssueTokenParams) GetTokenParamsOk() (*TokenizationIssueTokenParamsTokenParams, bool)`

GetTokenParamsOk returns a tuple with the TokenParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParams

`func (o *TokenizationIssueTokenParams) SetTokenParams(v TokenizationIssueTokenParamsTokenParams)`

SetTokenParams sets TokenParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


