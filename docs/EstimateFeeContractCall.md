# EstimateFeeContractCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
**RequestType** | **string** | The request type. Possible values include:   - &#x60;Transfer&#x60;: A request to transfer tokens.   - &#x60;ContractCall&#x60;: A request to interact with a smart contract.   - &#x60;MessageSign&#x60;: A request to sign a message.  | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 
**Source** | [**ContractCallSource**](ContractCallSource.md) |  | 
**Destination** | [**EstimateFeeContractCallDestination**](EstimateFeeContractCallDestination.md) |  | 

## Methods

### NewEstimateFeeContractCall

`func NewEstimateFeeContractCall(requestId string, requestType string, chainId string, source ContractCallSource, destination EstimateFeeContractCallDestination, ) *EstimateFeeContractCall`

NewEstimateFeeContractCall instantiates a new EstimateFeeContractCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateFeeContractCallWithDefaults

`func NewEstimateFeeContractCallWithDefaults() *EstimateFeeContractCall`

NewEstimateFeeContractCallWithDefaults instantiates a new EstimateFeeContractCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EstimateFeeContractCall) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateFeeContractCall) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateFeeContractCall) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequestType

`func (o *EstimateFeeContractCall) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EstimateFeeContractCall) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EstimateFeeContractCall) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetChainId

`func (o *EstimateFeeContractCall) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *EstimateFeeContractCall) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *EstimateFeeContractCall) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *EstimateFeeContractCall) GetSource() ContractCallSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateFeeContractCall) GetSourceOk() (*ContractCallSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateFeeContractCall) SetSource(v ContractCallSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *EstimateFeeContractCall) GetDestination() EstimateFeeContractCallDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *EstimateFeeContractCall) GetDestinationOk() (*EstimateFeeContractCallDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *EstimateFeeContractCall) SetDestination(v EstimateFeeContractCallDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


