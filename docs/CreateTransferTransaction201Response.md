# CreateTransferTransaction201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | Unique id of the request. | [optional] 
**TransactionId** | Pointer to **string** | The transaction id of of the request. | [optional] 
**Status** | Pointer to [**TransactionStatus**](TransactionStatus.md) |  | [optional] 

## Methods

### NewCreateTransferTransaction201Response

`func NewCreateTransferTransaction201Response() *CreateTransferTransaction201Response`

NewCreateTransferTransaction201Response instantiates a new CreateTransferTransaction201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferTransaction201ResponseWithDefaults

`func NewCreateTransferTransaction201ResponseWithDefaults() *CreateTransferTransaction201Response`

NewCreateTransferTransaction201ResponseWithDefaults instantiates a new CreateTransferTransaction201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateTransferTransaction201Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateTransferTransaction201Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateTransferTransaction201Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateTransferTransaction201Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTransactionId

`func (o *CreateTransferTransaction201Response) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CreateTransferTransaction201Response) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CreateTransferTransaction201Response) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CreateTransferTransaction201Response) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateTransferTransaction201Response) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTransferTransaction201Response) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTransferTransaction201Response) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTransferTransaction201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


