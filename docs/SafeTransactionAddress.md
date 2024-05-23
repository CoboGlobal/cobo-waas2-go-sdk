# SafeTransactionAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TransactionAddressType**](TransactionAddressType.md) |  | 
**Address** | Pointer to **string** | Address | [optional] 
**Memo** | Pointer to **string** | Address memo | [optional] 
**Delegate** | [**SafeTransactionAddressAllOfDelegate**](SafeTransactionAddressAllOfDelegate.md) |  | 

## Methods

### NewSafeTransactionAddress

`func NewSafeTransactionAddress(type_ TransactionAddressType, delegate SafeTransactionAddressAllOfDelegate, ) *SafeTransactionAddress`

NewSafeTransactionAddress instantiates a new SafeTransactionAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeTransactionAddressWithDefaults

`func NewSafeTransactionAddressWithDefaults() *SafeTransactionAddress`

NewSafeTransactionAddressWithDefaults instantiates a new SafeTransactionAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SafeTransactionAddress) GetType() TransactionAddressType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SafeTransactionAddress) GetTypeOk() (*TransactionAddressType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SafeTransactionAddress) SetType(v TransactionAddressType)`

SetType sets Type field to given value.


### GetAddress

`func (o *SafeTransactionAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SafeTransactionAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SafeTransactionAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SafeTransactionAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *SafeTransactionAddress) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *SafeTransactionAddress) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *SafeTransactionAddress) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *SafeTransactionAddress) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetDelegate

`func (o *SafeTransactionAddress) GetDelegate() SafeTransactionAddressAllOfDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *SafeTransactionAddress) GetDelegateOk() (*SafeTransactionAddressAllOfDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *SafeTransactionAddress) SetDelegate(v SafeTransactionAddressAllOfDelegate)`

SetDelegate sets Delegate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


