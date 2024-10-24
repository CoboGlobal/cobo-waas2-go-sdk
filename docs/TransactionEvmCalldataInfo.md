# TransactionEvmCalldataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **string** | The ID of the chain on which the smart contract is deployed. | [optional] 
**Address** | Pointer to **string** | The address of the smart contract. | [optional] 
**Name** | Pointer to **string** | The name of the smart contract. | [optional] 
**ImplAddress** | Pointer to **string** | The address of the implementation smart contract. This property is applicable only when the specified smart contract is a proxy contract. | [optional] 
**ImplName** | Pointer to **string** | The name of the implementation smart contract. This property is applicable only when the specified smart contract is a proxy contract. | [optional] 
**Proxy** | Pointer to **bool** | Whether the specified smart contract address is a proxy contract. - &#x60;true&#x60;: The specified smart contract address is a proxy contract. - &#x60;false&#x60;: The specified smart contract address is not a proxy contract.  | [optional] 
**Method** | Pointer to [**TransactionEvmContractMethod**](TransactionEvmContractMethod.md) |  | [optional] 
**Params** | Pointer to **string** | The parameters of the contract method are represented as a JSON array of arrays. Each element in the outer array is itself an array containing three elements that provide detailed information about a specific parameter: - Parameter name: The unique identifier of the parameter, such as &#x60;kind&#x60;, &#x60;swaps&#x60;, and &#x60;to&#x60;. - Parameter type: The Solidity data type of the parameter, such as &#x60;uint8&#x60;, &#x60;tuple[]&#x60;, &#x60;address[]&#x60;, and &#x60;int256[]&#x60;. - Parameter value: The actual value of the parameter. If the parameter type is a basic type such as &#x60;uint256&#x60; or &#x60;address&#x60;, this value is a single element. If the parameter type is a complex type such as &#x60;tuple[]&#x60; or &#x60;address[]&#x60;, the value is a nested array, with each inner array containing parameter names, types, and values.  | [optional] 

## Methods

### NewTransactionEvmCalldataInfo

`func NewTransactionEvmCalldataInfo() *TransactionEvmCalldataInfo`

NewTransactionEvmCalldataInfo instantiates a new TransactionEvmCalldataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEvmCalldataInfoWithDefaults

`func NewTransactionEvmCalldataInfoWithDefaults() *TransactionEvmCalldataInfo`

NewTransactionEvmCalldataInfoWithDefaults instantiates a new TransactionEvmCalldataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *TransactionEvmCalldataInfo) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TransactionEvmCalldataInfo) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TransactionEvmCalldataInfo) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *TransactionEvmCalldataInfo) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetAddress

`func (o *TransactionEvmCalldataInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionEvmCalldataInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionEvmCalldataInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionEvmCalldataInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *TransactionEvmCalldataInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionEvmCalldataInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionEvmCalldataInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionEvmCalldataInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImplAddress

`func (o *TransactionEvmCalldataInfo) GetImplAddress() string`

GetImplAddress returns the ImplAddress field if non-nil, zero value otherwise.

### GetImplAddressOk

`func (o *TransactionEvmCalldataInfo) GetImplAddressOk() (*string, bool)`

GetImplAddressOk returns a tuple with the ImplAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplAddress

`func (o *TransactionEvmCalldataInfo) SetImplAddress(v string)`

SetImplAddress sets ImplAddress field to given value.

### HasImplAddress

`func (o *TransactionEvmCalldataInfo) HasImplAddress() bool`

HasImplAddress returns a boolean if a field has been set.

### GetImplName

`func (o *TransactionEvmCalldataInfo) GetImplName() string`

GetImplName returns the ImplName field if non-nil, zero value otherwise.

### GetImplNameOk

`func (o *TransactionEvmCalldataInfo) GetImplNameOk() (*string, bool)`

GetImplNameOk returns a tuple with the ImplName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplName

`func (o *TransactionEvmCalldataInfo) SetImplName(v string)`

SetImplName sets ImplName field to given value.

### HasImplName

`func (o *TransactionEvmCalldataInfo) HasImplName() bool`

HasImplName returns a boolean if a field has been set.

### GetProxy

`func (o *TransactionEvmCalldataInfo) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *TransactionEvmCalldataInfo) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *TransactionEvmCalldataInfo) SetProxy(v bool)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *TransactionEvmCalldataInfo) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetMethod

`func (o *TransactionEvmCalldataInfo) GetMethod() TransactionEvmContractMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionEvmCalldataInfo) GetMethodOk() (*TransactionEvmContractMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionEvmCalldataInfo) SetMethod(v TransactionEvmContractMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TransactionEvmCalldataInfo) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetParams

`func (o *TransactionEvmCalldataInfo) GetParams() string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TransactionEvmCalldataInfo) GetParamsOk() (*string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TransactionEvmCalldataInfo) SetParams(v string)`

SetParams sets Params field to given value.

### HasParams

`func (o *TransactionEvmCalldataInfo) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


