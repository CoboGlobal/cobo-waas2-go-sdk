# TransactionDetails

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
**Fee** | Pointer to [**TransactionTransferFee**](TransactionTransferFee.md) |  | [optional] 
**Initiator** | Pointer to **string** | Transaction initiator | [optional] 
**InitiatorType** | [**TransactionInitiatorType**](TransactionInitiatorType.md) |  | 
**ConfirmedNum** | Pointer to **float32** | Transaction confirmed number | [optional] 
**ConfirmingThreshold** | Pointer to **int32** | Number of confirmations required for a transaction, such as 15 for ETH chain. | [optional] 
**BlockNumber** | Pointer to **int64** | The block number. | [optional] 
**BlockTime** | Pointer to **int64** | The time when the block was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**BlockHash** | Pointer to **string** | The block hash. | [optional] 
**Nonce** | Pointer to **int32** | Transaction nonce | [optional] 
**TransactionHash** | Pointer to **string** | The transaction hash. | [optional] 
**Replacement** | Pointer to [**TransactionReplacement**](TransactionReplacement.md) |  | [optional] 
**Category** | Pointer to **[]string** | A custom transaction category for you to identify your transfers more easily. | [optional] 
**Description** | Pointer to **string** | The description for your transaction. | [optional] 
**ForceInternal** | Pointer to **bool** | Whether the transfer request must be executed as an off-chain transfer.    - &#x60;true&#x60;: The transfer request must be executed an off-chain transfer.    - &#x60;false&#x60;: The transfer may not be executed as an off-chain transfer.  | [optional] 
**ForceExternal** | Pointer to **bool** | Whether the transfer must be executed as an on-chain transfer.   - &#x60;true&#x60;: The transfer must be executed as an on-chain transfer.   - &#x60;false&#x60;: The transfer may not be executed as an on-chain transfer.  | [optional] 
**IsLoop** | Pointer to **bool** | Whether the transfer is a Loop transfer. For more information about Loop, see [the website](https://loop.top/).  - &#x60;true&#x60;: The transfer is a Loop transfer. - &#x60;false&#x60;: The transfer is not a Loop transfer.  | [optional] 
**CreatedTime** | **float32** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTime** | **float32** | The time when the transaction was updated, in Unix timestamp format, measured in milliseconds. | 
**Approvers** | Pointer to [**[]TransactionApprover**](TransactionApprover.md) |  | [optional] 
**Signers** | Pointer to [**[]TransactionSigner**](TransactionSigner.md) |  | [optional] 
**ReplacedBy** | Pointer to **string** | Replace by transaction hash | [optional] 
**FueledBy** | Pointer to **string** | Fueled by address | [optional] 
**TokenApproval** | Pointer to [**TransactionTokeApproval**](TransactionTokeApproval.md) |  | [optional] 
**Message** | Pointer to **string** | Transaction raw message | [optional] 
**Algorithm** | Pointer to **string** | Transaction message signing algorithm | [optional] 
**Timeline** | Pointer to [**[]TransactionTimeline**](TransactionTimeline.md) |  | [optional] 

## Methods

### NewTransactionDetails

`func NewTransactionDetails(transactionId string, status TransactionStatus, source TransactionSource, destination TransactionDestination, initiatorType TransactionInitiatorType, createdTime float32, updatedTime float32, ) *TransactionDetails`

NewTransactionDetails instantiates a new TransactionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDetailsWithDefaults

`func NewTransactionDetailsWithDefaults() *TransactionDetails`

NewTransactionDetailsWithDefaults instantiates a new TransactionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *TransactionDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetCoboId

`func (o *TransactionDetails) GetCoboId() string`

GetCoboId returns the CoboId field if non-nil, zero value otherwise.

### GetCoboIdOk

`func (o *TransactionDetails) GetCoboIdOk() (*string, bool)`

GetCoboIdOk returns a tuple with the CoboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboId

`func (o *TransactionDetails) SetCoboId(v string)`

SetCoboId sets CoboId field to given value.

### HasCoboId

`func (o *TransactionDetails) HasCoboId() bool`

HasCoboId returns a boolean if a field has been set.

### GetRequestId

`func (o *TransactionDetails) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionDetails) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionDetails) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TransactionDetails) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetType

`func (o *TransactionDetails) GetType() TransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionDetails) GetTypeOk() (*TransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionDetails) SetType(v TransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionDetails) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionDetails) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionDetails) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetSubStatus

`func (o *TransactionDetails) GetSubStatus() TransactionSubStatus`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *TransactionDetails) GetSubStatusOk() (*TransactionSubStatus, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *TransactionDetails) SetSubStatus(v TransactionSubStatus)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *TransactionDetails) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetFailedReason

`func (o *TransactionDetails) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *TransactionDetails) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *TransactionDetails) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *TransactionDetails) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetChainId

`func (o *TransactionDetails) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TransactionDetails) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TransactionDetails) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *TransactionDetails) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetSource

`func (o *TransactionDetails) GetSource() TransactionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TransactionDetails) GetSourceOk() (*TransactionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TransactionDetails) SetSource(v TransactionSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *TransactionDetails) GetDestination() TransactionDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TransactionDetails) GetDestinationOk() (*TransactionDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TransactionDetails) SetDestination(v TransactionDestination)`

SetDestination sets Destination field to given value.


### GetFee

`func (o *TransactionDetails) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransactionDetails) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransactionDetails) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *TransactionDetails) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetInitiator

`func (o *TransactionDetails) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *TransactionDetails) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *TransactionDetails) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *TransactionDetails) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetInitiatorType

`func (o *TransactionDetails) GetInitiatorType() TransactionInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *TransactionDetails) GetInitiatorTypeOk() (*TransactionInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *TransactionDetails) SetInitiatorType(v TransactionInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.


### GetConfirmedNum

`func (o *TransactionDetails) GetConfirmedNum() float32`

GetConfirmedNum returns the ConfirmedNum field if non-nil, zero value otherwise.

### GetConfirmedNumOk

`func (o *TransactionDetails) GetConfirmedNumOk() (*float32, bool)`

GetConfirmedNumOk returns a tuple with the ConfirmedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedNum

`func (o *TransactionDetails) SetConfirmedNum(v float32)`

SetConfirmedNum sets ConfirmedNum field to given value.

### HasConfirmedNum

`func (o *TransactionDetails) HasConfirmedNum() bool`

HasConfirmedNum returns a boolean if a field has been set.

### GetConfirmingThreshold

`func (o *TransactionDetails) GetConfirmingThreshold() int32`

GetConfirmingThreshold returns the ConfirmingThreshold field if non-nil, zero value otherwise.

### GetConfirmingThresholdOk

`func (o *TransactionDetails) GetConfirmingThresholdOk() (*int32, bool)`

GetConfirmingThresholdOk returns a tuple with the ConfirmingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmingThreshold

`func (o *TransactionDetails) SetConfirmingThreshold(v int32)`

SetConfirmingThreshold sets ConfirmingThreshold field to given value.

### HasConfirmingThreshold

`func (o *TransactionDetails) HasConfirmingThreshold() bool`

HasConfirmingThreshold returns a boolean if a field has been set.

### GetBlockNumber

`func (o *TransactionDetails) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *TransactionDetails) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *TransactionDetails) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *TransactionDetails) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetBlockTime

`func (o *TransactionDetails) GetBlockTime() int64`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *TransactionDetails) GetBlockTimeOk() (*int64, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *TransactionDetails) SetBlockTime(v int64)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *TransactionDetails) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetBlockHash

`func (o *TransactionDetails) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionDetails) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionDetails) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *TransactionDetails) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetNonce

`func (o *TransactionDetails) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransactionDetails) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransactionDetails) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *TransactionDetails) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetTransactionHash

`func (o *TransactionDetails) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *TransactionDetails) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *TransactionDetails) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *TransactionDetails) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetReplacement

`func (o *TransactionDetails) GetReplacement() TransactionReplacement`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *TransactionDetails) GetReplacementOk() (*TransactionReplacement, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *TransactionDetails) SetReplacement(v TransactionReplacement)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *TransactionDetails) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetCategory

`func (o *TransactionDetails) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransactionDetails) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransactionDetails) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TransactionDetails) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceInternal

`func (o *TransactionDetails) GetForceInternal() bool`

GetForceInternal returns the ForceInternal field if non-nil, zero value otherwise.

### GetForceInternalOk

`func (o *TransactionDetails) GetForceInternalOk() (*bool, bool)`

GetForceInternalOk returns a tuple with the ForceInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInternal

`func (o *TransactionDetails) SetForceInternal(v bool)`

SetForceInternal sets ForceInternal field to given value.

### HasForceInternal

`func (o *TransactionDetails) HasForceInternal() bool`

HasForceInternal returns a boolean if a field has been set.

### GetForceExternal

`func (o *TransactionDetails) GetForceExternal() bool`

GetForceExternal returns the ForceExternal field if non-nil, zero value otherwise.

### GetForceExternalOk

`func (o *TransactionDetails) GetForceExternalOk() (*bool, bool)`

GetForceExternalOk returns a tuple with the ForceExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExternal

`func (o *TransactionDetails) SetForceExternal(v bool)`

SetForceExternal sets ForceExternal field to given value.

### HasForceExternal

`func (o *TransactionDetails) HasForceExternal() bool`

HasForceExternal returns a boolean if a field has been set.

### GetIsLoop

`func (o *TransactionDetails) GetIsLoop() bool`

GetIsLoop returns the IsLoop field if non-nil, zero value otherwise.

### GetIsLoopOk

`func (o *TransactionDetails) GetIsLoopOk() (*bool, bool)`

GetIsLoopOk returns a tuple with the IsLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoop

`func (o *TransactionDetails) SetIsLoop(v bool)`

SetIsLoop sets IsLoop field to given value.

### HasIsLoop

`func (o *TransactionDetails) HasIsLoop() bool`

HasIsLoop returns a boolean if a field has been set.

### GetCreatedTime

`func (o *TransactionDetails) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *TransactionDetails) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *TransactionDetails) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetUpdatedTime

`func (o *TransactionDetails) GetUpdatedTime() float32`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *TransactionDetails) GetUpdatedTimeOk() (*float32, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *TransactionDetails) SetUpdatedTime(v float32)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetApprovers

`func (o *TransactionDetails) GetApprovers() []TransactionApprover`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *TransactionDetails) GetApproversOk() (*[]TransactionApprover, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *TransactionDetails) SetApprovers(v []TransactionApprover)`

SetApprovers sets Approvers field to given value.

### HasApprovers

`func (o *TransactionDetails) HasApprovers() bool`

HasApprovers returns a boolean if a field has been set.

### GetSigners

`func (o *TransactionDetails) GetSigners() []TransactionSigner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *TransactionDetails) GetSignersOk() (*[]TransactionSigner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *TransactionDetails) SetSigners(v []TransactionSigner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *TransactionDetails) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### GetReplacedBy

`func (o *TransactionDetails) GetReplacedBy() string`

GetReplacedBy returns the ReplacedBy field if non-nil, zero value otherwise.

### GetReplacedByOk

`func (o *TransactionDetails) GetReplacedByOk() (*string, bool)`

GetReplacedByOk returns a tuple with the ReplacedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedBy

`func (o *TransactionDetails) SetReplacedBy(v string)`

SetReplacedBy sets ReplacedBy field to given value.

### HasReplacedBy

`func (o *TransactionDetails) HasReplacedBy() bool`

HasReplacedBy returns a boolean if a field has been set.

### GetFueledBy

`func (o *TransactionDetails) GetFueledBy() string`

GetFueledBy returns the FueledBy field if non-nil, zero value otherwise.

### GetFueledByOk

`func (o *TransactionDetails) GetFueledByOk() (*string, bool)`

GetFueledByOk returns a tuple with the FueledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFueledBy

`func (o *TransactionDetails) SetFueledBy(v string)`

SetFueledBy sets FueledBy field to given value.

### HasFueledBy

`func (o *TransactionDetails) HasFueledBy() bool`

HasFueledBy returns a boolean if a field has been set.

### GetTokenApproval

`func (o *TransactionDetails) GetTokenApproval() TransactionTokeApproval`

GetTokenApproval returns the TokenApproval field if non-nil, zero value otherwise.

### GetTokenApprovalOk

`func (o *TransactionDetails) GetTokenApprovalOk() (*TransactionTokeApproval, bool)`

GetTokenApprovalOk returns a tuple with the TokenApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenApproval

`func (o *TransactionDetails) SetTokenApproval(v TransactionTokeApproval)`

SetTokenApproval sets TokenApproval field to given value.

### HasTokenApproval

`func (o *TransactionDetails) HasTokenApproval() bool`

HasTokenApproval returns a boolean if a field has been set.

### GetMessage

`func (o *TransactionDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransactionDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAlgorithm

`func (o *TransactionDetails) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TransactionDetails) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TransactionDetails) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *TransactionDetails) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetTimeline

`func (o *TransactionDetails) GetTimeline() []TransactionTimeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *TransactionDetails) GetTimelineOk() (*[]TransactionTimeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *TransactionDetails) SetTimeline(v []TransactionTimeline)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *TransactionDetails) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


