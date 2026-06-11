# FiatTransactionEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. - &#x60;Organization&#x60;: The organization event data. - &#x60;FiatTransaction&#x60;: The fiat transaction event data. | 
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

### NewFiatTransactionEventData

`func NewFiatTransactionEventData(dataType string, transactionId string, transactionType FeeStationFiatTransactionType, amount string, fiatCurrency string, status string, ) *FiatTransactionEventData`

NewFiatTransactionEventData instantiates a new FiatTransactionEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatTransactionEventDataWithDefaults

`func NewFiatTransactionEventDataWithDefaults() *FiatTransactionEventData`

NewFiatTransactionEventDataWithDefaults instantiates a new FiatTransactionEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *FiatTransactionEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *FiatTransactionEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *FiatTransactionEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetTransactionId

`func (o *FiatTransactionEventData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *FiatTransactionEventData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *FiatTransactionEventData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetMainTransactionId

`func (o *FiatTransactionEventData) GetMainTransactionId() string`

GetMainTransactionId returns the MainTransactionId field if non-nil, zero value otherwise.

### GetMainTransactionIdOk

`func (o *FiatTransactionEventData) GetMainTransactionIdOk() (*string, bool)`

GetMainTransactionIdOk returns a tuple with the MainTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTransactionId

`func (o *FiatTransactionEventData) SetMainTransactionId(v string)`

SetMainTransactionId sets MainTransactionId field to given value.

### HasMainTransactionId

`func (o *FiatTransactionEventData) HasMainTransactionId() bool`

HasMainTransactionId returns a boolean if a field has been set.

### GetTransactionType

`func (o *FiatTransactionEventData) GetTransactionType() FeeStationFiatTransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *FiatTransactionEventData) GetTransactionTypeOk() (*FeeStationFiatTransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *FiatTransactionEventData) SetTransactionType(v FeeStationFiatTransactionType)`

SetTransactionType sets TransactionType field to given value.


### GetAmount

`func (o *FiatTransactionEventData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FiatTransactionEventData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FiatTransactionEventData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFiatCurrency

`func (o *FiatTransactionEventData) GetFiatCurrency() string`

GetFiatCurrency returns the FiatCurrency field if non-nil, zero value otherwise.

### GetFiatCurrencyOk

`func (o *FiatTransactionEventData) GetFiatCurrencyOk() (*string, bool)`

GetFiatCurrencyOk returns a tuple with the FiatCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatCurrency

`func (o *FiatTransactionEventData) SetFiatCurrency(v string)`

SetFiatCurrency sets FiatCurrency field to given value.


### GetStatus

`func (o *FiatTransactionEventData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FiatTransactionEventData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FiatTransactionEventData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCoboCategory

`func (o *FiatTransactionEventData) GetCoboCategory() []string`

GetCoboCategory returns the CoboCategory field if non-nil, zero value otherwise.

### GetCoboCategoryOk

`func (o *FiatTransactionEventData) GetCoboCategoryOk() (*[]string, bool)`

GetCoboCategoryOk returns a tuple with the CoboCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboCategory

`func (o *FiatTransactionEventData) SetCoboCategory(v []string)`

SetCoboCategory sets CoboCategory field to given value.

### HasCoboCategory

`func (o *FiatTransactionEventData) HasCoboCategory() bool`

HasCoboCategory returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *FiatTransactionEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FiatTransactionEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FiatTransactionEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *FiatTransactionEventData) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetModifiedTimestamp

`func (o *FiatTransactionEventData) GetModifiedTimestamp() int64`

GetModifiedTimestamp returns the ModifiedTimestamp field if non-nil, zero value otherwise.

### GetModifiedTimestampOk

`func (o *FiatTransactionEventData) GetModifiedTimestampOk() (*int64, bool)`

GetModifiedTimestampOk returns a tuple with the ModifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimestamp

`func (o *FiatTransactionEventData) SetModifiedTimestamp(v int64)`

SetModifiedTimestamp sets ModifiedTimestamp field to given value.

### HasModifiedTimestamp

`func (o *FiatTransactionEventData) HasModifiedTimestamp() bool`

HasModifiedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


