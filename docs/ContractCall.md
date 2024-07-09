# ContractCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 
**Source** | [**ContractCallSource**](ContractCallSource.md) |  | 
**Destination** | [**EstimateFeeContractCallDestination**](EstimateFeeContractCallDestination.md) |  | 
**Fee** | Pointer to [**TransactionTransferFee**](TransactionTransferFee.md) |  | [optional] 

## Methods

### NewContractCall

`func NewContractCall(requestId string, chainId string, source ContractCallSource, destination EstimateFeeContractCallDestination, ) *ContractCall`

NewContractCall instantiates a new ContractCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallWithDefaults

`func NewContractCallWithDefaults() *ContractCall`

NewContractCallWithDefaults instantiates a new ContractCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ContractCall) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ContractCall) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ContractCall) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetChainId

`func (o *ContractCall) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ContractCall) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ContractCall) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *ContractCall) GetSource() ContractCallSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContractCall) GetSourceOk() (*ContractCallSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContractCall) SetSource(v ContractCallSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *ContractCall) GetDestination() EstimateFeeContractCallDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ContractCall) GetDestinationOk() (*EstimateFeeContractCallDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ContractCall) SetDestination(v EstimateFeeContractCallDestination)`

SetDestination sets Destination field to given value.


### GetFee

`func (o *ContractCall) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ContractCall) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ContractCall) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *ContractCall) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


