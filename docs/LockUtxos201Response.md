# LockUtxos201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executed** | Pointer to **bool** | Whether the locking or unlocking operation has been successfully executed. - &#x60;true&#x60;: The operation has been successfully executed. - &#x60;false&#x60;: The operation has not been executed.  | [optional] 

## Methods

### NewLockUtxos201Response

`func NewLockUtxos201Response() *LockUtxos201Response`

NewLockUtxos201Response instantiates a new LockUtxos201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockUtxos201ResponseWithDefaults

`func NewLockUtxos201ResponseWithDefaults() *LockUtxos201Response`

NewLockUtxos201ResponseWithDefaults instantiates a new LockUtxos201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecuted

`func (o *LockUtxos201Response) GetExecuted() bool`

GetExecuted returns the Executed field if non-nil, zero value otherwise.

### GetExecutedOk

`func (o *LockUtxos201Response) GetExecutedOk() (*bool, bool)`

GetExecutedOk returns a tuple with the Executed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuted

`func (o *LockUtxos201Response) SetExecuted(v bool)`

SetExecuted sets Executed field to given value.

### HasExecuted

`func (o *LockUtxos201Response) HasExecuted() bool`

HasExecuted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


