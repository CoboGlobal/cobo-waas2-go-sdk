# StellarContractCallTrustLineParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractType** | [**StellarContractCallContractType**](StellarContractCallContractType.md) |  | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. | 
**OperationType** | [**StellarContractCallTrustLineOperationType**](StellarContractCallTrustLineOperationType.md) |  | 

## Methods

### NewStellarContractCallTrustLineParam

`func NewStellarContractCallTrustLineParam(contractType StellarContractCallContractType, tokenId string, operationType StellarContractCallTrustLineOperationType, ) *StellarContractCallTrustLineParam`

NewStellarContractCallTrustLineParam instantiates a new StellarContractCallTrustLineParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStellarContractCallTrustLineParamWithDefaults

`func NewStellarContractCallTrustLineParamWithDefaults() *StellarContractCallTrustLineParam`

NewStellarContractCallTrustLineParamWithDefaults instantiates a new StellarContractCallTrustLineParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractType

`func (o *StellarContractCallTrustLineParam) GetContractType() StellarContractCallContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *StellarContractCallTrustLineParam) GetContractTypeOk() (*StellarContractCallContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *StellarContractCallTrustLineParam) SetContractType(v StellarContractCallContractType)`

SetContractType sets ContractType field to given value.


### GetTokenId

`func (o *StellarContractCallTrustLineParam) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *StellarContractCallTrustLineParam) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *StellarContractCallTrustLineParam) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetOperationType

`func (o *StellarContractCallTrustLineParam) GetOperationType() StellarContractCallTrustLineOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *StellarContractCallTrustLineParam) GetOperationTypeOk() (*StellarContractCallTrustLineOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *StellarContractCallTrustLineParam) SetOperationType(v StellarContractCallTrustLineOperationType)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


