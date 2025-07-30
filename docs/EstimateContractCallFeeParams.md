# EstimateContractCallFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. It is recommended to use the same request ID as the transaction for which you want to estimate the transaction fee. | [optional] 
**RequestType** | [**EstimateFeeRequestType**](EstimateFeeRequestType.md) |  | 
**ChainId** | **string** | The chain ID of the chain on which the smart contract is issued. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
**Source** | [**ContractCallSource**](ContractCallSource.md) |  | 
**Destination** | Pointer to [**ContractCallDestination**](ContractCallDestination.md) |  | [optional] 
**FeeType** | Pointer to [**FeeType**](FeeType.md) |  | [optional] [default to FEETYPE_EVM_EIP_1559]
**ReplacedTransactionId** | Pointer to **string** | The ID of the transaction that this transaction replaced. | [optional] 

## Methods

### NewEstimateContractCallFeeParams

`func NewEstimateContractCallFeeParams(requestType EstimateFeeRequestType, chainId string, source ContractCallSource, ) *EstimateContractCallFeeParams`

NewEstimateContractCallFeeParams instantiates a new EstimateContractCallFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateContractCallFeeParamsWithDefaults

`func NewEstimateContractCallFeeParamsWithDefaults() *EstimateContractCallFeeParams`

NewEstimateContractCallFeeParamsWithDefaults instantiates a new EstimateContractCallFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EstimateContractCallFeeParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateContractCallFeeParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateContractCallFeeParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EstimateContractCallFeeParams) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestType

`func (o *EstimateContractCallFeeParams) GetRequestType() EstimateFeeRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EstimateContractCallFeeParams) GetRequestTypeOk() (*EstimateFeeRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EstimateContractCallFeeParams) SetRequestType(v EstimateFeeRequestType)`

SetRequestType sets RequestType field to given value.


### GetChainId

`func (o *EstimateContractCallFeeParams) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *EstimateContractCallFeeParams) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *EstimateContractCallFeeParams) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *EstimateContractCallFeeParams) GetSource() ContractCallSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateContractCallFeeParams) GetSourceOk() (*ContractCallSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateContractCallFeeParams) SetSource(v ContractCallSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *EstimateContractCallFeeParams) GetDestination() ContractCallDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *EstimateContractCallFeeParams) GetDestinationOk() (*ContractCallDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *EstimateContractCallFeeParams) SetDestination(v ContractCallDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *EstimateContractCallFeeParams) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetFeeType

`func (o *EstimateContractCallFeeParams) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimateContractCallFeeParams) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimateContractCallFeeParams) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *EstimateContractCallFeeParams) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetReplacedTransactionId

`func (o *EstimateContractCallFeeParams) GetReplacedTransactionId() string`

GetReplacedTransactionId returns the ReplacedTransactionId field if non-nil, zero value otherwise.

### GetReplacedTransactionIdOk

`func (o *EstimateContractCallFeeParams) GetReplacedTransactionIdOk() (*string, bool)`

GetReplacedTransactionIdOk returns a tuple with the ReplacedTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedTransactionId

`func (o *EstimateContractCallFeeParams) SetReplacedTransactionId(v string)`

SetReplacedTransactionId sets ReplacedTransactionId field to given value.

### HasReplacedTransactionId

`func (o *EstimateContractCallFeeParams) HasReplacedTransactionId() bool`

HasReplacedTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


