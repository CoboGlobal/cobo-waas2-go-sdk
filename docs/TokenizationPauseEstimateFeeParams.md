# TokenizationPauseEstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**TokenId** | **string** | The ID of the token. | 
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | [optional] 

## Methods

### NewTokenizationPauseEstimateFeeParams

`func NewTokenizationPauseEstimateFeeParams(source TokenizationTokenOperationSource, operationType TokenizationOperationType, tokenId string, ) *TokenizationPauseEstimateFeeParams`

NewTokenizationPauseEstimateFeeParams instantiates a new TokenizationPauseEstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationPauseEstimateFeeParamsWithDefaults

`func NewTokenizationPauseEstimateFeeParamsWithDefaults() *TokenizationPauseEstimateFeeParams`

NewTokenizationPauseEstimateFeeParamsWithDefaults instantiates a new TokenizationPauseEstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationPauseEstimateFeeParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationPauseEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationPauseEstimateFeeParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetOperationType

`func (o *TokenizationPauseEstimateFeeParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationPauseEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationPauseEstimateFeeParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.


### GetTokenId

`func (o *TokenizationPauseEstimateFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationPauseEstimateFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationPauseEstimateFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetRequestId

`func (o *TokenizationPauseEstimateFeeParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TokenizationPauseEstimateFeeParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TokenizationPauseEstimateFeeParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TokenizationPauseEstimateFeeParams) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


