# BankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountId** | **string** | The bank account ID. | 
**Info** | **map[string]interface{}** | JSON-formatted bank account details. | 
**CreatedTimestamp** | Pointer to **int32** | The creation time of the bank account, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The last update time of the bank account, represented as a UNIX timestamp in seconds. | [optional] 

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


### GetCreatedTimestamp

`func (o *BankAccount) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *BankAccount) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *BankAccount) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *BankAccount) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *BankAccount) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *BankAccount) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *BankAccount) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *BankAccount) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


