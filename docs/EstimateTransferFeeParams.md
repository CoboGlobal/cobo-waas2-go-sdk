# EstimateTransferFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. It is recommended to use the same request ID as the transaction for which you want to estimate the transaction fee. | [optional] 
**RequestType** | [**EstimateFeeRequestType**](EstimateFeeRequestType.md) |  | 
**Source** | [**TransferSource**](TransferSource.md) |  | 
**TokenId** | **string** | The token ID of the transferred token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
**Destination** | Pointer to [**TransferDestination**](TransferDestination.md) |  | [optional] 
**FeeType** | Pointer to [**FeeType**](FeeType.md) |  | [optional] [default to FEETYPE_EVM_EIP_1559]
**ReplacedTransactionId** | Pointer to **string** | The ID of the transaction that this transaction replaced. | [optional] 

## Methods

### NewEstimateTransferFeeParams

`func NewEstimateTransferFeeParams(requestType EstimateFeeRequestType, source TransferSource, tokenId string, ) *EstimateTransferFeeParams`

NewEstimateTransferFeeParams instantiates a new EstimateTransferFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateTransferFeeParamsWithDefaults

`func NewEstimateTransferFeeParamsWithDefaults() *EstimateTransferFeeParams`

NewEstimateTransferFeeParamsWithDefaults instantiates a new EstimateTransferFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EstimateTransferFeeParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateTransferFeeParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateTransferFeeParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EstimateTransferFeeParams) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestType

`func (o *EstimateTransferFeeParams) GetRequestType() EstimateFeeRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EstimateTransferFeeParams) GetRequestTypeOk() (*EstimateFeeRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EstimateTransferFeeParams) SetRequestType(v EstimateFeeRequestType)`

SetRequestType sets RequestType field to given value.


### GetSource

`func (o *EstimateTransferFeeParams) GetSource() TransferSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateTransferFeeParams) GetSourceOk() (*TransferSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateTransferFeeParams) SetSource(v TransferSource)`

SetSource sets Source field to given value.


### GetTokenId

`func (o *EstimateTransferFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimateTransferFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimateTransferFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDestination

`func (o *EstimateTransferFeeParams) GetDestination() TransferDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *EstimateTransferFeeParams) GetDestinationOk() (*TransferDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *EstimateTransferFeeParams) SetDestination(v TransferDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *EstimateTransferFeeParams) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetFeeType

`func (o *EstimateTransferFeeParams) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimateTransferFeeParams) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimateTransferFeeParams) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *EstimateTransferFeeParams) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetReplacedTransactionId

`func (o *EstimateTransferFeeParams) GetReplacedTransactionId() string`

GetReplacedTransactionId returns the ReplacedTransactionId field if non-nil, zero value otherwise.

### GetReplacedTransactionIdOk

`func (o *EstimateTransferFeeParams) GetReplacedTransactionIdOk() (*string, bool)`

GetReplacedTransactionIdOk returns a tuple with the ReplacedTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedTransactionId

`func (o *EstimateTransferFeeParams) SetReplacedTransactionId(v string)`

SetReplacedTransactionId sets ReplacedTransactionId field to given value.

### HasReplacedTransactionId

`func (o *EstimateTransferFeeParams) HasReplacedTransactionId() bool`

HasReplacedTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


