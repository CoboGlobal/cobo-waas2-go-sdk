# TokenizationContractCallEstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | [optional] 
**Data** | Pointer to [**TokenizationContractCallParamsData**](TokenizationContractCallParamsData.md) |  | [optional] 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**TokenId** | **string** | The ID of the token. | 

## Methods

### NewTokenizationContractCallEstimateFeeParams

`func NewTokenizationContractCallEstimateFeeParams(operationType TokenizationOperationType, tokenId string, ) *TokenizationContractCallEstimateFeeParams`

NewTokenizationContractCallEstimateFeeParams instantiates a new TokenizationContractCallEstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationContractCallEstimateFeeParamsWithDefaults

`func NewTokenizationContractCallEstimateFeeParamsWithDefaults() *TokenizationContractCallEstimateFeeParams`

NewTokenizationContractCallEstimateFeeParamsWithDefaults instantiates a new TokenizationContractCallEstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationContractCallEstimateFeeParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationContractCallEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationContractCallEstimateFeeParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *TokenizationContractCallEstimateFeeParams) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetData

`func (o *TokenizationContractCallEstimateFeeParams) GetData() TokenizationContractCallParamsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationContractCallEstimateFeeParams) GetDataOk() (*TokenizationContractCallParamsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationContractCallEstimateFeeParams) SetData(v TokenizationContractCallParamsData)`

SetData sets Data field to given value.

### HasData

`func (o *TokenizationContractCallEstimateFeeParams) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOperationType

`func (o *TokenizationContractCallEstimateFeeParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationContractCallEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationContractCallEstimateFeeParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.


### GetTokenId

`func (o *TokenizationContractCallEstimateFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationContractCallEstimateFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationContractCallEstimateFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


