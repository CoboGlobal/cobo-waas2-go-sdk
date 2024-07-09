# EstimateFeeTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
**RequestType** | **string** | The request type. Possible values include:   - &#x60;Transfer&#x60;: A request to transfer tokens.   - &#x60;ContractCall&#x60;: A request to interact with a smart contract.   - &#x60;MessageSign&#x60;: A request to sign a message.  | 
**Source** | [**TransferSource**](TransferSource.md) |  | 
**TokenId** | **string** | The token ID of the transaction fee. You can retrieve token IDs by using the [Get fee rates](/api-references/v2/transactions/get-fee-rates) operation. | 
**Destination** | [**TransferDestination**](TransferDestination.md) |  | 

## Methods

### NewEstimateFeeTransfer

`func NewEstimateFeeTransfer(requestId string, requestType string, source TransferSource, tokenId string, destination TransferDestination, ) *EstimateFeeTransfer`

NewEstimateFeeTransfer instantiates a new EstimateFeeTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateFeeTransferWithDefaults

`func NewEstimateFeeTransferWithDefaults() *EstimateFeeTransfer`

NewEstimateFeeTransferWithDefaults instantiates a new EstimateFeeTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EstimateFeeTransfer) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateFeeTransfer) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateFeeTransfer) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequestType

`func (o *EstimateFeeTransfer) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EstimateFeeTransfer) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EstimateFeeTransfer) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetSource

`func (o *EstimateFeeTransfer) GetSource() TransferSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateFeeTransfer) GetSourceOk() (*TransferSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateFeeTransfer) SetSource(v TransferSource)`

SetSource sets Source field to given value.


### GetTokenId

`func (o *EstimateFeeTransfer) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimateFeeTransfer) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimateFeeTransfer) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDestination

`func (o *EstimateFeeTransfer) GetDestination() TransferDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *EstimateFeeTransfer) GetDestinationOk() (*TransferDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *EstimateFeeTransfer) SetDestination(v TransferDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


