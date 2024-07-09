# CreateTransferTransaction201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
**TransactionId** | **string** | The transaction ID. | 
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | 

## Methods

### NewCreateTransferTransaction201Response

`func NewCreateTransferTransaction201Response(requestId string, transactionId string, status TransactionStatus, ) *CreateTransferTransaction201Response`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


