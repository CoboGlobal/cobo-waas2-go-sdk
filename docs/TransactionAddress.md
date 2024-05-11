# TransactionAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address | [optional] 
**Memo** | Pointer to **string** | Address memo | [optional] 
**Label** | Pointer to **string** | Address label | [optional] 

## Methods

### NewTransactionAddress

`func NewTransactionAddress() *TransactionAddress`

NewTransactionAddress instantiates a new TransactionAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAddressWithDefaults

`func NewTransactionAddressWithDefaults() *TransactionAddress`

NewTransactionAddressWithDefaults instantiates a new TransactionAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransactionAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *TransactionAddress) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionAddress) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionAddress) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionAddress) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetLabel

`func (o *TransactionAddress) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TransactionAddress) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TransactionAddress) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TransactionAddress) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


