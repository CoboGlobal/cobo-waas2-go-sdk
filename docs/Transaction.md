# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction ID. | 
**CoboId** | Pointer to **string** | The Cobo ID, which can be used to track a transaction. | [optional] 
**RequestId** | Pointer to **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | [optional] 
**Type** | Pointer to [**TransactionType**](TransactionType.md) |  | [optional] 
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | 
**SubStatus** | Pointer to [**TransactionSubStatus**](TransactionSubStatus.md) |  | [optional] 
**FailedReason** | Pointer to **string** | The reason why the transaction failed. | [optional] 
**ChainId** | Pointer to **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | [optional] 
**Source** | [**TransactionSource**](TransactionSource.md) |  | 
**Destination** | [**TransactionDestination**](TransactionDestination.md) |  | 
**Fee** | Pointer to [**TransactionFee**](TransactionFee.md) |  | [optional] 
**Initiator** | Pointer to **string** | The transaction initiator. | [optional] 
**InitiatorType** | [**TransactionInitiatorType**](TransactionInitiatorType.md) |  | 
**ConfirmedNum** | Pointer to **float32** | The number of confirmations this transaction has received. | [optional] 
**ConfirmingThreshold** | Pointer to **int32** | The minimum number of confirmations required to deem a transaction secure. The common threshold is 6 for a Bitcoin transaction. | [optional] 
**BlockNumber** | Pointer to **int64** | The block number. | [optional] 
**BlockTime** | Pointer to **int64** | The time when the block was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**BlockHash** | Pointer to **string** | The block hash. | [optional] 
**Nonce** | Pointer to **int32** | The transaction nonce. | [optional] 
**TransactionHash** | Pointer to **string** | The transaction hash. | [optional] 
**Replacement** | Pointer to [**TransactionReplacement**](TransactionReplacement.md) |  | [optional] 
**Category** | Pointer to **[]string** | A custom transaction category for you to identify your transfers more easily. | [optional] 
**Description** | Pointer to **string** | The description for your transaction. | [optional] 
**ForceInternal** | Pointer to **bool** | Whether the transfer request must be executed as an off-chain transfer.    - &#x60;true&#x60;: The transfer request must be executed an off-chain transfer.    - &#x60;false&#x60;: The transfer may not be executed as an off-chain transfer.  | [optional] 
**ForceExternal** | Pointer to **bool** | Whether the transfer must be executed as an on-chain transfer.   - &#x60;true&#x60;: The transfer must be executed as an on-chain transfer.   - &#x60;false&#x60;: The transfer may not be executed as an on-chain transfer.  | [optional] 
**IsLoop** | Pointer to **bool** | Whether the transfer is a Loop transfer. For more information about Loop, see [the website](https://loop.top/).  - &#x60;true&#x60;: The transfer is a Loop transfer. - &#x60;false&#x60;: The transfer is not a Loop transfer.  | [optional] 
**CreatedTime** | **float32** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTime** | **float32** | The time when the transaction was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewTransaction

`func NewTransaction(transactionId string, status TransactionStatus, source TransactionSource, destination TransactionDestination, initiatorType TransactionInitiatorType, createdTime float32, updatedTime float32, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *Transaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetCoboId

`func (o *Transaction) GetCoboId() string`

GetCoboId returns the CoboId field if non-nil, zero value otherwise.

### GetCoboIdOk

`func (o *Transaction) GetCoboIdOk() (*string, bool)`

GetCoboIdOk returns a tuple with the CoboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboId

`func (o *Transaction) SetCoboId(v string)`

SetCoboId sets CoboId field to given value.

### HasCoboId

`func (o *Transaction) HasCoboId() bool`

HasCoboId returns a boolean if a field has been set.

### GetRequestId

`func (o *Transaction) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Transaction) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Transaction) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Transaction) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetType

`func (o *Transaction) GetType() TransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*TransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v TransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetSubStatus

`func (o *Transaction) GetSubStatus() TransactionSubStatus`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *Transaction) GetSubStatusOk() (*TransactionSubStatus, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *Transaction) SetSubStatus(v TransactionSubStatus)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *Transaction) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetFailedReason

`func (o *Transaction) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *Transaction) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *Transaction) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *Transaction) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetChainId

`func (o *Transaction) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Transaction) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Transaction) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Transaction) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetSource

`func (o *Transaction) GetSource() TransactionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Transaction) GetSourceOk() (*TransactionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Transaction) SetSource(v TransactionSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *Transaction) GetDestination() TransactionDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Transaction) GetDestinationOk() (*TransactionDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Transaction) SetDestination(v TransactionDestination)`

SetDestination sets Destination field to given value.


### GetFee

`func (o *Transaction) GetFee() TransactionFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Transaction) GetFeeOk() (*TransactionFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Transaction) SetFee(v TransactionFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *Transaction) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetInitiator

`func (o *Transaction) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *Transaction) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *Transaction) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *Transaction) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetInitiatorType

`func (o *Transaction) GetInitiatorType() TransactionInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *Transaction) GetInitiatorTypeOk() (*TransactionInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *Transaction) SetInitiatorType(v TransactionInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.


### GetConfirmedNum

`func (o *Transaction) GetConfirmedNum() float32`

GetConfirmedNum returns the ConfirmedNum field if non-nil, zero value otherwise.

### GetConfirmedNumOk

`func (o *Transaction) GetConfirmedNumOk() (*float32, bool)`

GetConfirmedNumOk returns a tuple with the ConfirmedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedNum

`func (o *Transaction) SetConfirmedNum(v float32)`

SetConfirmedNum sets ConfirmedNum field to given value.

### HasConfirmedNum

`func (o *Transaction) HasConfirmedNum() bool`

HasConfirmedNum returns a boolean if a field has been set.

### GetConfirmingThreshold

`func (o *Transaction) GetConfirmingThreshold() int32`

GetConfirmingThreshold returns the ConfirmingThreshold field if non-nil, zero value otherwise.

### GetConfirmingThresholdOk

`func (o *Transaction) GetConfirmingThresholdOk() (*int32, bool)`

GetConfirmingThresholdOk returns a tuple with the ConfirmingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmingThreshold

`func (o *Transaction) SetConfirmingThreshold(v int32)`

SetConfirmingThreshold sets ConfirmingThreshold field to given value.

### HasConfirmingThreshold

`func (o *Transaction) HasConfirmingThreshold() bool`

HasConfirmingThreshold returns a boolean if a field has been set.

### GetBlockNumber

`func (o *Transaction) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *Transaction) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *Transaction) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *Transaction) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetBlockTime

`func (o *Transaction) GetBlockTime() int64`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *Transaction) GetBlockTimeOk() (*int64, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *Transaction) SetBlockTime(v int64)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *Transaction) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetBlockHash

`func (o *Transaction) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *Transaction) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *Transaction) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *Transaction) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetNonce

`func (o *Transaction) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Transaction) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Transaction) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Transaction) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetTransactionHash

`func (o *Transaction) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Transaction) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Transaction) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *Transaction) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetReplacement

`func (o *Transaction) GetReplacement() TransactionReplacement`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *Transaction) GetReplacementOk() (*TransactionReplacement, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *Transaction) SetReplacement(v TransactionReplacement)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *Transaction) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetCategory

`func (o *Transaction) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transaction) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transaction) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Transaction) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceInternal

`func (o *Transaction) GetForceInternal() bool`

GetForceInternal returns the ForceInternal field if non-nil, zero value otherwise.

### GetForceInternalOk

`func (o *Transaction) GetForceInternalOk() (*bool, bool)`

GetForceInternalOk returns a tuple with the ForceInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInternal

`func (o *Transaction) SetForceInternal(v bool)`

SetForceInternal sets ForceInternal field to given value.

### HasForceInternal

`func (o *Transaction) HasForceInternal() bool`

HasForceInternal returns a boolean if a field has been set.

### GetForceExternal

`func (o *Transaction) GetForceExternal() bool`

GetForceExternal returns the ForceExternal field if non-nil, zero value otherwise.

### GetForceExternalOk

`func (o *Transaction) GetForceExternalOk() (*bool, bool)`

GetForceExternalOk returns a tuple with the ForceExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExternal

`func (o *Transaction) SetForceExternal(v bool)`

SetForceExternal sets ForceExternal field to given value.

### HasForceExternal

`func (o *Transaction) HasForceExternal() bool`

HasForceExternal returns a boolean if a field has been set.

### GetIsLoop

`func (o *Transaction) GetIsLoop() bool`

GetIsLoop returns the IsLoop field if non-nil, zero value otherwise.

### GetIsLoopOk

`func (o *Transaction) GetIsLoopOk() (*bool, bool)`

GetIsLoopOk returns a tuple with the IsLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoop

`func (o *Transaction) SetIsLoop(v bool)`

SetIsLoop sets IsLoop field to given value.

### HasIsLoop

`func (o *Transaction) HasIsLoop() bool`

HasIsLoop returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Transaction) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Transaction) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Transaction) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetUpdatedTime

`func (o *Transaction) GetUpdatedTime() float32`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Transaction) GetUpdatedTimeOk() (*float32, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Transaction) SetUpdatedTime(v float32)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


