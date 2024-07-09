# EstimateFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
**RequestType** | **string** | The request type. Possible values include:   - &#x60;Transfer&#x60;: A request to transfer tokens.   - &#x60;ContractCall&#x60;: A request to interact with a smart contract.   - &#x60;MessageSign&#x60;: A request to sign a message.  | 
**Source** | [**ContractCallSource**](ContractCallSource.md) |  | 
**TokenId** | **string** | The token ID of the transaction fee. You can retrieve token IDs by using the [Get fee rates](/api-references/v2/transactions/get-fee-rates) operation. | 
**Destination** | [**EstimateFeeContractCallDestination**](EstimateFeeContractCallDestination.md) |  | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 

## Methods

### NewEstimateFee

`func NewEstimateFee(requestId string, requestType string, source ContractCallSource, tokenId string, destination EstimateFeeContractCallDestination, chainId string, ) *EstimateFee`

NewEstimateFee instantiates a new EstimateFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateFeeWithDefaults

`func NewEstimateFeeWithDefaults() *EstimateFee`

NewEstimateFeeWithDefaults instantiates a new EstimateFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EstimateFee) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateFee) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateFee) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequestType

`func (o *EstimateFee) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EstimateFee) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EstimateFee) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetSource

`func (o *EstimateFee) GetSource() ContractCallSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateFee) GetSourceOk() (*ContractCallSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateFee) SetSource(v ContractCallSource)`

SetSource sets Source field to given value.


### GetTokenId

`func (o *EstimateFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimateFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimateFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDestination

`func (o *EstimateFee) GetDestination() EstimateFeeContractCallDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *EstimateFee) GetDestinationOk() (*EstimateFeeContractCallDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *EstimateFee) SetDestination(v EstimateFeeContractCallDestination)`

SetDestination sets Destination field to given value.


### GetChainId

`func (o *EstimateFee) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *EstimateFee) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *EstimateFee) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


