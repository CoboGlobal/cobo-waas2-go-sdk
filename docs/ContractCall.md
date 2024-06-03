# ContractCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique id of the request. | 
**RequestType** | **string** |  | 
**ChainId** | **string** | The blockchain on which the token operates. | 
**Source** | [**ContractCallSource**](ContractCallSource.md) |  | 
**Destination** | Pointer to [**ContractCallDestination**](ContractCallDestination.md) |  | [optional] 
**Fee** | Pointer to [**TransactionFee**](TransactionFee.md) |  | [optional] 

## Methods

### NewContractCall

`func NewContractCall(requestId string, requestType string, chainId string, source ContractCallSource, ) *ContractCall`

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


### GetRequestType

`func (o *ContractCall) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ContractCall) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ContractCall) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


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

`func (o *ContractCall) GetDestination() ContractCallDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ContractCall) GetDestinationOk() (*ContractCallDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ContractCall) SetDestination(v ContractCallDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ContractCall) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetFee

`func (o *ContractCall) GetFee() TransactionFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ContractCall) GetFeeOk() (*TransactionFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ContractCall) SetFee(v TransactionFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *ContractCall) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


