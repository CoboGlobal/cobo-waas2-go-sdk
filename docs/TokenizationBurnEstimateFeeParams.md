# TokenizationBurnEstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Burns** | [**[]TokenizationBurnTokenParamsBurnsInner**](TokenizationBurnTokenParamsBurnsInner.md) | Details for each token burn, including amount and address to burn from. | 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**TokenId** | **string** | The ID of the token. | 
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | [optional] 

## Methods

### NewTokenizationBurnEstimateFeeParams

`func NewTokenizationBurnEstimateFeeParams(source TokenizationTokenOperationSource, burns []TokenizationBurnTokenParamsBurnsInner, operationType TokenizationOperationType, tokenId string, ) *TokenizationBurnEstimateFeeParams`

NewTokenizationBurnEstimateFeeParams instantiates a new TokenizationBurnEstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationBurnEstimateFeeParamsWithDefaults

`func NewTokenizationBurnEstimateFeeParamsWithDefaults() *TokenizationBurnEstimateFeeParams`

NewTokenizationBurnEstimateFeeParamsWithDefaults instantiates a new TokenizationBurnEstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationBurnEstimateFeeParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationBurnEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationBurnEstimateFeeParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetBurns

`func (o *TokenizationBurnEstimateFeeParams) GetBurns() []TokenizationBurnTokenParamsBurnsInner`

GetBurns returns the Burns field if non-nil, zero value otherwise.

### GetBurnsOk

`func (o *TokenizationBurnEstimateFeeParams) GetBurnsOk() (*[]TokenizationBurnTokenParamsBurnsInner, bool)`

GetBurnsOk returns a tuple with the Burns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurns

`func (o *TokenizationBurnEstimateFeeParams) SetBurns(v []TokenizationBurnTokenParamsBurnsInner)`

SetBurns sets Burns field to given value.


### GetOperationType

`func (o *TokenizationBurnEstimateFeeParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationBurnEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationBurnEstimateFeeParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.


### GetTokenId

`func (o *TokenizationBurnEstimateFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationBurnEstimateFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationBurnEstimateFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetRequestId

`func (o *TokenizationBurnEstimateFeeParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TokenizationBurnEstimateFeeParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TokenizationBurnEstimateFeeParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TokenizationBurnEstimateFeeParams) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


