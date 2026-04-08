# FeeStationFiatTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction ID. | 
**MainTransactionId** | Pointer to **string** | The UUID of the parent (main) transaction that this record is associated with. Set only when the current record is a gas/fee transaction initiated by FeeStation; omit for main transactions. | [optional] 
**TransactionType** | [**FeeStationFiatTransactionType**](FeeStationFiatTransactionType.md) |  | 
**Amount** | **string** | The transaction amount. | 
**FiatCurrency** | **string** | The fiat currency of the transaction. Possible values include:   - &#x60;USD&#x60;: US Dollar.  | 
**Status** | **string** | The status of the fiat transaction. Possible values include:   - &#x60;Created&#x60;: The transaction has been created.   - &#x60;Succeeded&#x60;: The transaction has been completed successfully.  | 
**CoboCategory** | Pointer to **[]string** | The Cobo category of the transaction. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**ModifiedTimestamp** | Pointer to **int64** | The time when the transaction was last modified, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewFeeStationFiatTransaction

`func NewFeeStationFiatTransaction(transactionId string, transactionType FeeStationFiatTransactionType, amount string, fiatCurrency string, status string, ) *FeeStationFiatTransaction`

NewFeeStationFiatTransaction instantiates a new FeeStationFiatTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeStationFiatTransactionWithDefaults

`func NewFeeStationFiatTransactionWithDefaults() *FeeStationFiatTransaction`

NewFeeStationFiatTransactionWithDefaults instantiates a new FeeStationFiatTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *FeeStationFiatTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *FeeStationFiatTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *FeeStationFiatTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetMainTransactionId

`func (o *FeeStationFiatTransaction) GetMainTransactionId() string`

GetMainTransactionId returns the MainTransactionId field if non-nil, zero value otherwise.

### GetMainTransactionIdOk

`func (o *FeeStationFiatTransaction) GetMainTransactionIdOk() (*string, bool)`

GetMainTransactionIdOk returns a tuple with the MainTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTransactionId

`func (o *FeeStationFiatTransaction) SetMainTransactionId(v string)`

SetMainTransactionId sets MainTransactionId field to given value.

### HasMainTransactionId

`func (o *FeeStationFiatTransaction) HasMainTransactionId() bool`

HasMainTransactionId returns a boolean if a field has been set.

### GetTransactionType

`func (o *FeeStationFiatTransaction) GetTransactionType() FeeStationFiatTransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *FeeStationFiatTransaction) GetTransactionTypeOk() (*FeeStationFiatTransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *FeeStationFiatTransaction) SetTransactionType(v FeeStationFiatTransactionType)`

SetTransactionType sets TransactionType field to given value.


### GetAmount

`func (o *FeeStationFiatTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeStationFiatTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeStationFiatTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFiatCurrency

`func (o *FeeStationFiatTransaction) GetFiatCurrency() string`

GetFiatCurrency returns the FiatCurrency field if non-nil, zero value otherwise.

### GetFiatCurrencyOk

`func (o *FeeStationFiatTransaction) GetFiatCurrencyOk() (*string, bool)`

GetFiatCurrencyOk returns a tuple with the FiatCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatCurrency

`func (o *FeeStationFiatTransaction) SetFiatCurrency(v string)`

SetFiatCurrency sets FiatCurrency field to given value.


### GetStatus

`func (o *FeeStationFiatTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeeStationFiatTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeeStationFiatTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCoboCategory

`func (o *FeeStationFiatTransaction) GetCoboCategory() []string`

GetCoboCategory returns the CoboCategory field if non-nil, zero value otherwise.

### GetCoboCategoryOk

`func (o *FeeStationFiatTransaction) GetCoboCategoryOk() (*[]string, bool)`

GetCoboCategoryOk returns a tuple with the CoboCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboCategory

`func (o *FeeStationFiatTransaction) SetCoboCategory(v []string)`

SetCoboCategory sets CoboCategory field to given value.

### HasCoboCategory

`func (o *FeeStationFiatTransaction) HasCoboCategory() bool`

HasCoboCategory returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *FeeStationFiatTransaction) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FeeStationFiatTransaction) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FeeStationFiatTransaction) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *FeeStationFiatTransaction) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetModifiedTimestamp

`func (o *FeeStationFiatTransaction) GetModifiedTimestamp() int64`

GetModifiedTimestamp returns the ModifiedTimestamp field if non-nil, zero value otherwise.

### GetModifiedTimestampOk

`func (o *FeeStationFiatTransaction) GetModifiedTimestampOk() (*int64, bool)`

GetModifiedTimestampOk returns a tuple with the ModifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimestamp

`func (o *FeeStationFiatTransaction) SetModifiedTimestamp(v int64)`

SetModifiedTimestamp sets ModifiedTimestamp field to given value.

### HasModifiedTimestamp

`func (o *FeeStationFiatTransaction) HasModifiedTimestamp() bool`

HasModifiedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


