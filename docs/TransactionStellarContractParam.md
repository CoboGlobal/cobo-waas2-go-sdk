# TransactionStellarContractParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractType** | [**TransactionStellarContractType**](TransactionStellarContractType.md) |  | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. | 
**OperationType** | [**TransactionStellarTrustLineOperationType**](TransactionStellarTrustLineOperationType.md) |  | 

## Methods

### NewTransactionStellarContractParam

`func NewTransactionStellarContractParam(contractType TransactionStellarContractType, tokenId string, operationType TransactionStellarTrustLineOperationType, ) *TransactionStellarContractParam`

NewTransactionStellarContractParam instantiates a new TransactionStellarContractParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStellarContractParamWithDefaults

`func NewTransactionStellarContractParamWithDefaults() *TransactionStellarContractParam`

NewTransactionStellarContractParamWithDefaults instantiates a new TransactionStellarContractParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractType

`func (o *TransactionStellarContractParam) GetContractType() TransactionStellarContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *TransactionStellarContractParam) GetContractTypeOk() (*TransactionStellarContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *TransactionStellarContractParam) SetContractType(v TransactionStellarContractType)`

SetContractType sets ContractType field to given value.


### GetTokenId

`func (o *TransactionStellarContractParam) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionStellarContractParam) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionStellarContractParam) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetOperationType

`func (o *TransactionStellarContractParam) GetOperationType() TransactionStellarTrustLineOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TransactionStellarContractParam) GetOperationTypeOk() (*TransactionStellarTrustLineOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TransactionStellarContractParam) SetOperationType(v TransactionStellarTrustLineOperationType)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


