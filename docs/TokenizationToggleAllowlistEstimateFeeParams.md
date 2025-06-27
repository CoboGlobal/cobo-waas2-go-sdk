# TokenizationToggleAllowlistEstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Activation** | **bool** | Whether to activate the allowlist feature for the token. | 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**TokenId** | **string** | The ID of the token. | 

## Methods

### NewTokenizationToggleAllowlistEstimateFeeParams

`func NewTokenizationToggleAllowlistEstimateFeeParams(source TokenizationTokenOperationSource, activation bool, operationType TokenizationOperationType, tokenId string, ) *TokenizationToggleAllowlistEstimateFeeParams`

NewTokenizationToggleAllowlistEstimateFeeParams instantiates a new TokenizationToggleAllowlistEstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationToggleAllowlistEstimateFeeParamsWithDefaults

`func NewTokenizationToggleAllowlistEstimateFeeParamsWithDefaults() *TokenizationToggleAllowlistEstimateFeeParams`

NewTokenizationToggleAllowlistEstimateFeeParamsWithDefaults instantiates a new TokenizationToggleAllowlistEstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationToggleAllowlistEstimateFeeParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationToggleAllowlistEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationToggleAllowlistEstimateFeeParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetActivation

`func (o *TokenizationToggleAllowlistEstimateFeeParams) GetActivation() bool`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *TokenizationToggleAllowlistEstimateFeeParams) GetActivationOk() (*bool, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *TokenizationToggleAllowlistEstimateFeeParams) SetActivation(v bool)`

SetActivation sets Activation field to given value.


### GetOperationType

`func (o *TokenizationToggleAllowlistEstimateFeeParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationToggleAllowlistEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationToggleAllowlistEstimateFeeParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.


### GetTokenId

`func (o *TokenizationToggleAllowlistEstimateFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationToggleAllowlistEstimateFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationToggleAllowlistEstimateFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


