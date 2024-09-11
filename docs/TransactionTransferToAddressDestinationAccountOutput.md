# TransactionTransferToAddressDestinationAccountOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The destination address. | [optional] 
**Memo** | Pointer to **string** | The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON. | [optional] 
**Amount** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 BTC, then the value is &#x60;1.5&#x60;.  | [optional] 

## Methods

### NewTransactionTransferToAddressDestinationAccountOutput

`func NewTransactionTransferToAddressDestinationAccountOutput() *TransactionTransferToAddressDestinationAccountOutput`

NewTransactionTransferToAddressDestinationAccountOutput instantiates a new TransactionTransferToAddressDestinationAccountOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTransferToAddressDestinationAccountOutputWithDefaults

`func NewTransactionTransferToAddressDestinationAccountOutputWithDefaults() *TransactionTransferToAddressDestinationAccountOutput`

NewTransactionTransferToAddressDestinationAccountOutputWithDefaults instantiates a new TransactionTransferToAddressDestinationAccountOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransactionTransferToAddressDestinationAccountOutput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionTransferToAddressDestinationAccountOutput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionTransferToAddressDestinationAccountOutput) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionTransferToAddressDestinationAccountOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *TransactionTransferToAddressDestinationAccountOutput) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionTransferToAddressDestinationAccountOutput) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionTransferToAddressDestinationAccountOutput) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionTransferToAddressDestinationAccountOutput) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionTransferToAddressDestinationAccountOutput) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionTransferToAddressDestinationAccountOutput) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionTransferToAddressDestinationAccountOutput) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionTransferToAddressDestinationAccountOutput) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


