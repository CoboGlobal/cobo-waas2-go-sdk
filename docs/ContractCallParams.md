# ContractCallParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
**Source** | [**ContractCallSource**](ContractCallSource.md) |  | 
**Destination** | [**ContractCallDestination**](ContractCallDestination.md) |  | 
**Description** | Pointer to **string** | The description of the contract call transaction. | [optional] 
**CategoryNames** | Pointer to **[]string** | The custom category for you to identify your transactions. | [optional] 
**Fee** | Pointer to [**TransactionRequestFee**](TransactionRequestFee.md) |  | [optional] 
**TransactionProcessType** | Pointer to **string** | Transaction processing type. Possible values are: - &#x60;AutoProcess&#x60; (default): After the transaction is constructed, it will be automatically signed and broadcast.   - &#x60;BuildOnly&#x60;: Set to this value if you want to build the transaction first without automatically signing and broadcasting it. You can manually call the [Sign and broadcast transaction](https://www.cobo.com/developers/v2/api-references/transactions/sign-and-broadcast-transaction) operation to complete the signing and broadcasting process.  | [optional] 
**AutoFuel** | Pointer to [**AutoFuelType**](AutoFuelType.md) |  | [optional] 

## Methods

### NewContractCallParams

`func NewContractCallParams(requestId string, chainId string, source ContractCallSource, destination ContractCallDestination, ) *ContractCallParams`

NewContractCallParams instantiates a new ContractCallParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallParamsWithDefaults

`func NewContractCallParamsWithDefaults() *ContractCallParams`

NewContractCallParamsWithDefaults instantiates a new ContractCallParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ContractCallParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ContractCallParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ContractCallParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetChainId

`func (o *ContractCallParams) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ContractCallParams) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ContractCallParams) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *ContractCallParams) GetSource() ContractCallSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContractCallParams) GetSourceOk() (*ContractCallSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContractCallParams) SetSource(v ContractCallSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *ContractCallParams) GetDestination() ContractCallDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ContractCallParams) GetDestinationOk() (*ContractCallDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ContractCallParams) SetDestination(v ContractCallDestination)`

SetDestination sets Destination field to given value.


### GetDescription

`func (o *ContractCallParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContractCallParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContractCallParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContractCallParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategoryNames

`func (o *ContractCallParams) GetCategoryNames() []string`

GetCategoryNames returns the CategoryNames field if non-nil, zero value otherwise.

### GetCategoryNamesOk

`func (o *ContractCallParams) GetCategoryNamesOk() (*[]string, bool)`

GetCategoryNamesOk returns a tuple with the CategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNames

`func (o *ContractCallParams) SetCategoryNames(v []string)`

SetCategoryNames sets CategoryNames field to given value.

### HasCategoryNames

`func (o *ContractCallParams) HasCategoryNames() bool`

HasCategoryNames returns a boolean if a field has been set.

### GetFee

`func (o *ContractCallParams) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ContractCallParams) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ContractCallParams) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *ContractCallParams) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetTransactionProcessType

`func (o *ContractCallParams) GetTransactionProcessType() string`

GetTransactionProcessType returns the TransactionProcessType field if non-nil, zero value otherwise.

### GetTransactionProcessTypeOk

`func (o *ContractCallParams) GetTransactionProcessTypeOk() (*string, bool)`

GetTransactionProcessTypeOk returns a tuple with the TransactionProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionProcessType

`func (o *ContractCallParams) SetTransactionProcessType(v string)`

SetTransactionProcessType sets TransactionProcessType field to given value.

### HasTransactionProcessType

`func (o *ContractCallParams) HasTransactionProcessType() bool`

HasTransactionProcessType returns a boolean if a field has been set.

### GetAutoFuel

`func (o *ContractCallParams) GetAutoFuel() AutoFuelType`

GetAutoFuel returns the AutoFuel field if non-nil, zero value otherwise.

### GetAutoFuelOk

`func (o *ContractCallParams) GetAutoFuelOk() (*AutoFuelType, bool)`

GetAutoFuelOk returns a tuple with the AutoFuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFuel

`func (o *ContractCallParams) SetAutoFuel(v AutoFuelType)`

SetAutoFuel sets AutoFuel field to given value.

### HasAutoFuel

`func (o *ContractCallParams) HasAutoFuel() bool`

HasAutoFuel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


