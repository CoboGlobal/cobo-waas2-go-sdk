# CallbackMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The callback message ID. | 
**CreatedTimestamp** | **int64** | The time when the callback message was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTimestamp** | **int64** | The time when the callback message was updated, in Unix timestamp format, measured in milliseconds. | 
**RequestId** | **string** | The request ID of the transaction. | 
**TransactionId** | **string** | The transaction ID. | 
**WalletId** | Pointer to **string** | The wallet ID. | [optional] 
**Url** | **string** | The callback endpoint URL. | 
**Data** | [**Transaction**](Transaction.md) |  | 
**Status** | **string** | The callback message status. Possible values include &#x60;Denied&#x60;, &#x60;Approved&#x60;, and &#x60;Failed&#x60;.  | 
**Result** | Pointer to **string** | The callback message result. Possible values include &#x60;ok&#x60; and &#x60;deny&#x60;.  | [optional] 

## Methods

### NewCallbackMessage

`func NewCallbackMessage(id string, createdTimestamp int64, updatedTimestamp int64, requestId string, transactionId string, url string, data Transaction, status string, ) *CallbackMessage`

NewCallbackMessage instantiates a new CallbackMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbackMessageWithDefaults

`func NewCallbackMessageWithDefaults() *CallbackMessage`

NewCallbackMessageWithDefaults instantiates a new CallbackMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CallbackMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallbackMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallbackMessage) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedTimestamp

`func (o *CallbackMessage) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *CallbackMessage) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *CallbackMessage) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *CallbackMessage) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *CallbackMessage) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *CallbackMessage) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRequestId

`func (o *CallbackMessage) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CallbackMessage) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CallbackMessage) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTransactionId

`func (o *CallbackMessage) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CallbackMessage) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CallbackMessage) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetWalletId

`func (o *CallbackMessage) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CallbackMessage) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CallbackMessage) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *CallbackMessage) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetUrl

`func (o *CallbackMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CallbackMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CallbackMessage) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetData

`func (o *CallbackMessage) GetData() Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallbackMessage) GetDataOk() (*Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallbackMessage) SetData(v Transaction)`

SetData sets Data field to given value.


### GetStatus

`func (o *CallbackMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallbackMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallbackMessage) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResult

`func (o *CallbackMessage) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CallbackMessage) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CallbackMessage) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CallbackMessage) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


