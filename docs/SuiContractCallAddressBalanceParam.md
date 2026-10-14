# SuiContractCallAddressBalanceParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractType** | [**SuiContractCallContractType**](SuiContractCallContractType.md) |  | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. | 
**OperationType** | [**SuiContractCallAddressBalanceOperationType**](SuiContractCallAddressBalanceOperationType.md) |  | 
**Txid** | Pointer to **string** | The transaction hash of the deposit to return. Required when &#x60;operation_type&#x60; is &#x60;AddressBalanceRefund&#x60;, and must be omitted otherwise. The destination address is resolved from this deposit and cannot be specified.  | [optional] 

## Methods

### NewSuiContractCallAddressBalanceParam

`func NewSuiContractCallAddressBalanceParam(contractType SuiContractCallContractType, tokenId string, operationType SuiContractCallAddressBalanceOperationType, ) *SuiContractCallAddressBalanceParam`

NewSuiContractCallAddressBalanceParam instantiates a new SuiContractCallAddressBalanceParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiContractCallAddressBalanceParamWithDefaults

`func NewSuiContractCallAddressBalanceParamWithDefaults() *SuiContractCallAddressBalanceParam`

NewSuiContractCallAddressBalanceParamWithDefaults instantiates a new SuiContractCallAddressBalanceParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractType

`func (o *SuiContractCallAddressBalanceParam) GetContractType() SuiContractCallContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *SuiContractCallAddressBalanceParam) GetContractTypeOk() (*SuiContractCallContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *SuiContractCallAddressBalanceParam) SetContractType(v SuiContractCallContractType)`

SetContractType sets ContractType field to given value.


### GetTokenId

`func (o *SuiContractCallAddressBalanceParam) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SuiContractCallAddressBalanceParam) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SuiContractCallAddressBalanceParam) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetOperationType

`func (o *SuiContractCallAddressBalanceParam) GetOperationType() SuiContractCallAddressBalanceOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *SuiContractCallAddressBalanceParam) GetOperationTypeOk() (*SuiContractCallAddressBalanceOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *SuiContractCallAddressBalanceParam) SetOperationType(v SuiContractCallAddressBalanceOperationType)`

SetOperationType sets OperationType field to given value.


### GetTxid

`func (o *SuiContractCallAddressBalanceParam) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *SuiContractCallAddressBalanceParam) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *SuiContractCallAddressBalanceParam) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *SuiContractCallAddressBalanceParam) HasTxid() bool`

HasTxid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


