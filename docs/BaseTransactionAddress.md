# BaseTransactionAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TransactionAddressType**](TransactionAddressType.md) |  | 
**Address** | Pointer to **string** | Address | [optional] 
**Memo** | Pointer to **string** | Address memo | [optional] 

## Methods

### NewBaseTransactionAddress

`func NewBaseTransactionAddress(type_ TransactionAddressType, ) *BaseTransactionAddress`

NewBaseTransactionAddress instantiates a new BaseTransactionAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTransactionAddressWithDefaults

`func NewBaseTransactionAddressWithDefaults() *BaseTransactionAddress`

NewBaseTransactionAddressWithDefaults instantiates a new BaseTransactionAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BaseTransactionAddress) GetType() TransactionAddressType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseTransactionAddress) GetTypeOk() (*TransactionAddressType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseTransactionAddress) SetType(v TransactionAddressType)`

SetType sets Type field to given value.


### GetAddress

`func (o *BaseTransactionAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BaseTransactionAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BaseTransactionAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BaseTransactionAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *BaseTransactionAddress) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *BaseTransactionAddress) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *BaseTransactionAddress) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *BaseTransactionAddress) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


