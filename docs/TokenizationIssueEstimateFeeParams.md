# TokenizationIssueEstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain ID where the token will be issued. | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**TokenParams** | [**TokenizationIssueTokenParamsTokenParams**](TokenizationIssueTokenParamsTokenParams.md) |  | 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 

## Methods

### NewTokenizationIssueEstimateFeeParams

`func NewTokenizationIssueEstimateFeeParams(chainId string, source TokenizationTokenOperationSource, tokenParams TokenizationIssueTokenParamsTokenParams, operationType TokenizationOperationType, ) *TokenizationIssueEstimateFeeParams`

NewTokenizationIssueEstimateFeeParams instantiates a new TokenizationIssueEstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationIssueEstimateFeeParamsWithDefaults

`func NewTokenizationIssueEstimateFeeParamsWithDefaults() *TokenizationIssueEstimateFeeParams`

NewTokenizationIssueEstimateFeeParamsWithDefaults instantiates a new TokenizationIssueEstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *TokenizationIssueEstimateFeeParams) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenizationIssueEstimateFeeParams) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenizationIssueEstimateFeeParams) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *TokenizationIssueEstimateFeeParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationIssueEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationIssueEstimateFeeParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetTokenParams

`func (o *TokenizationIssueEstimateFeeParams) GetTokenParams() TokenizationIssueTokenParamsTokenParams`

GetTokenParams returns the TokenParams field if non-nil, zero value otherwise.

### GetTokenParamsOk

`func (o *TokenizationIssueEstimateFeeParams) GetTokenParamsOk() (*TokenizationIssueTokenParamsTokenParams, bool)`

GetTokenParamsOk returns a tuple with the TokenParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParams

`func (o *TokenizationIssueEstimateFeeParams) SetTokenParams(v TokenizationIssueTokenParamsTokenParams)`

SetTokenParams sets TokenParams field to given value.


### GetOperationType

`func (o *TokenizationIssueEstimateFeeParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationIssueEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationIssueEstimateFeeParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


