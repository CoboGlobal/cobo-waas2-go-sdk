# TokenizationMintEstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Mints** | [**[]TokenizationMintTokenParamsMintsInner**](TokenizationMintTokenParamsMintsInner.md) | Details for each token mint, including amount and address to mint to. | 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**TokenId** | **string** | The ID of the token. | 

## Methods

### NewTokenizationMintEstimateFeeParams

`func NewTokenizationMintEstimateFeeParams(source TokenizationTokenOperationSource, mints []TokenizationMintTokenParamsMintsInner, operationType TokenizationOperationType, tokenId string, ) *TokenizationMintEstimateFeeParams`

NewTokenizationMintEstimateFeeParams instantiates a new TokenizationMintEstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationMintEstimateFeeParamsWithDefaults

`func NewTokenizationMintEstimateFeeParamsWithDefaults() *TokenizationMintEstimateFeeParams`

NewTokenizationMintEstimateFeeParamsWithDefaults instantiates a new TokenizationMintEstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationMintEstimateFeeParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationMintEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationMintEstimateFeeParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetMints

`func (o *TokenizationMintEstimateFeeParams) GetMints() []TokenizationMintTokenParamsMintsInner`

GetMints returns the Mints field if non-nil, zero value otherwise.

### GetMintsOk

`func (o *TokenizationMintEstimateFeeParams) GetMintsOk() (*[]TokenizationMintTokenParamsMintsInner, bool)`

GetMintsOk returns a tuple with the Mints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMints

`func (o *TokenizationMintEstimateFeeParams) SetMints(v []TokenizationMintTokenParamsMintsInner)`

SetMints sets Mints field to given value.


### GetOperationType

`func (o *TokenizationMintEstimateFeeParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationMintEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationMintEstimateFeeParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.


### GetTokenId

`func (o *TokenizationMintEstimateFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationMintEstimateFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationMintEstimateFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


