# BankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountId** | **string** | The bank account ID. | 
**Info** | **map[string]interface{}** | JSON-formatted bank account details. | 

## Methods

### NewBankAccount

`func NewBankAccount(bankAccountId string, info map[string]interface{}, ) *BankAccount`

NewBankAccount instantiates a new BankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountWithDefaults

`func NewBankAccountWithDefaults() *BankAccount`

NewBankAccountWithDefaults instantiates a new BankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountId

`func (o *BankAccount) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankAccount) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankAccount) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.


### GetInfo

`func (o *BankAccount) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BankAccount) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BankAccount) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


