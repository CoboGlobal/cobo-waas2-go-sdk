# EstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. It is recommended to use the same request ID as the transaction for which you want to estimate the transaction fee. | 
**RequestType** | [**EstimateFeeRequestType**](EstimateFeeRequestType.md) |  | 
**Source** | [**ContractCallSource**](ContractCallSource.md) |  | 
**TokenId** | **string** | The token ID of the transferred token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens). | 
**Destination** | [**ContractCallDestination**](ContractCallDestination.md) |  | 
**FeeType** | Pointer to [**FeeType**](FeeType.md) |  | [optional] [default to FEETYPE_EVM_EIP_1559]
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](/v2/api-references/wallets/list-enabled-chains). | 

## Methods

### NewEstimateFeeParams

`func NewEstimateFeeParams(requestId string, requestType EstimateFeeRequestType, source ContractCallSource, tokenId string, destination ContractCallDestination, chainId string, ) *EstimateFeeParams`

NewEstimateFeeParams instantiates a new EstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateFeeParamsWithDefaults

`func NewEstimateFeeParamsWithDefaults() *EstimateFeeParams`

NewEstimateFeeParamsWithDefaults instantiates a new EstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EstimateFeeParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateFeeParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateFeeParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequestType

`func (o *EstimateFeeParams) GetRequestType() EstimateFeeRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EstimateFeeParams) GetRequestTypeOk() (*EstimateFeeRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EstimateFeeParams) SetRequestType(v EstimateFeeRequestType)`

SetRequestType sets RequestType field to given value.


### GetSource

`func (o *EstimateFeeParams) GetSource() ContractCallSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateFeeParams) GetSourceOk() (*ContractCallSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateFeeParams) SetSource(v ContractCallSource)`

SetSource sets Source field to given value.


### GetTokenId

`func (o *EstimateFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimateFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimateFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDestination

`func (o *EstimateFeeParams) GetDestination() ContractCallDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *EstimateFeeParams) GetDestinationOk() (*ContractCallDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *EstimateFeeParams) SetDestination(v ContractCallDestination)`

SetDestination sets Destination field to given value.


### GetFeeType

`func (o *EstimateFeeParams) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimateFeeParams) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimateFeeParams) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *EstimateFeeParams) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetChainId

`func (o *EstimateFeeParams) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *EstimateFeeParams) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *EstimateFeeParams) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


