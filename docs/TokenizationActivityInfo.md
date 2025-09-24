# TokenizationActivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **string** | The ID of the activity. | 
**TokenId** | **string** | The token ID. | 
**Type** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**Status** | [**TokenizationActivityStatus**](TokenizationActivityStatus.md) |  | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Initiator** | **string** | The initiator of the activity. | 
**InitiatorType** | [**TransactionInitiatorType**](TransactionInitiatorType.md) |  | 
**TransactionIds** | **[]string** | The IDs of the corresponding transactions of the activity. | 
**CreatedTimestamp** | **int64** | The creation timestamp of the activity in milliseconds since the Unix epoch. | 
**UpdatedTimestamp** | **int64** | The last update timestamp of the activity in milliseconds since the Unix epoch. | 

## Methods

### NewTokenizationActivityInfo

`func NewTokenizationActivityInfo(activityId string, tokenId string, type_ TokenizationOperationType, status TokenizationActivityStatus, source TokenizationTokenOperationSource, initiator string, initiatorType TransactionInitiatorType, transactionIds []string, createdTimestamp int64, updatedTimestamp int64, ) *TokenizationActivityInfo`

NewTokenizationActivityInfo instantiates a new TokenizationActivityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationActivityInfoWithDefaults

`func NewTokenizationActivityInfoWithDefaults() *TokenizationActivityInfo`

NewTokenizationActivityInfoWithDefaults instantiates a new TokenizationActivityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *TokenizationActivityInfo) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *TokenizationActivityInfo) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *TokenizationActivityInfo) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetTokenId

`func (o *TokenizationActivityInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationActivityInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationActivityInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetType

`func (o *TokenizationActivityInfo) GetType() TokenizationOperationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizationActivityInfo) GetTypeOk() (*TokenizationOperationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizationActivityInfo) SetType(v TokenizationOperationType)`

SetType sets Type field to given value.


### GetStatus

`func (o *TokenizationActivityInfo) GetStatus() TokenizationActivityStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenizationActivityInfo) GetStatusOk() (*TokenizationActivityStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenizationActivityInfo) SetStatus(v TokenizationActivityStatus)`

SetStatus sets Status field to given value.


### GetSource

`func (o *TokenizationActivityInfo) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationActivityInfo) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationActivityInfo) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetInitiator

`func (o *TokenizationActivityInfo) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *TokenizationActivityInfo) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *TokenizationActivityInfo) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.


### GetInitiatorType

`func (o *TokenizationActivityInfo) GetInitiatorType() TransactionInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *TokenizationActivityInfo) GetInitiatorTypeOk() (*TransactionInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *TokenizationActivityInfo) SetInitiatorType(v TransactionInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.


### GetTransactionIds

`func (o *TokenizationActivityInfo) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *TokenizationActivityInfo) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *TokenizationActivityInfo) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.


### GetCreatedTimestamp

`func (o *TokenizationActivityInfo) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TokenizationActivityInfo) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TokenizationActivityInfo) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *TokenizationActivityInfo) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *TokenizationActivityInfo) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *TokenizationActivityInfo) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


