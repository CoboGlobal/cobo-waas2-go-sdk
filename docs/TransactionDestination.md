# TransactionDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**AccountOutput** | Pointer to [**TransactionTransferToAddressDestinationAccountOutput**](TransactionTransferToAddressDestinationAccountOutput.md) |  | [optional] 
**UtxoOutputs** | Pointer to [**[]TransactionTransferToAddressDestinationUtxoOutputsInner**](TransactionTransferToAddressDestinationUtxoOutputsInner.md) |  | [optional] 
**ChangeAddress** | Pointer to **string** | The address used to receive the remaining funds or change from the transaction. | [optional] 
**ForceInternal** | Pointer to **bool** | Whether the transaction request must be executed as a Loop transfer. For more information about Loop, see [Loop&#39;s website](https://loop.top/).   - &#x60;true&#x60;: The transaction request must be executed as a Loop transfer.   - &#x60;false&#x60;: The transaction request may not be executed as a Loop transfer.  | [optional] 
**ForceExternal** | Pointer to **bool** | Whether the transaction request must not be executed as a Loop transfer. For more information about Loop, see [Loop&#39;s website](https://loop.top/).   - &#x60;true&#x60;: The transaction request must not be executed as a Loop transfer.   - &#x60;false&#x60;: The transaction request can be executed as a Loop transfer.  | [optional] 
**WalletId** | **string** | The wallet ID. | 
**TradingAccountType** | Pointer to **string** | The trading account type. | [optional] 
**ExchangeId** | Pointer to [**ExchangeId**](ExchangeId.md) |  | [optional] 
**Amount** | **string** | The transfer amount. For example, if you trade 1.5 BTC, then the value is &#x60;1.5&#x60;.  | 
**Address** | **string** | The destination address. | 
**Value** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Calldata** | **string** | The data that is used to invoke a specific function or method within the specified contract at the destination address.  | 
**Message** | **string** | The raw data of the message to be signed, encoded in Base64 format. | 
**StructuredData** | **map[string]interface{}** | The structured data to be signed, formatted as a JSON object according to the EIP-712 standard. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Memo** | Pointer to **string** | The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON. | [optional] 

## Methods

### NewTransactionDestination

`func NewTransactionDestination(destinationType TransactionDestinationType, walletId string, amount string, address string, calldata string, message string, structuredData map[string]interface{}, walletType WalletType, walletSubtype WalletSubtype, ) *TransactionDestination`

NewTransactionDestination instantiates a new TransactionDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDestinationWithDefaults

`func NewTransactionDestinationWithDefaults() *TransactionDestination`

NewTransactionDestinationWithDefaults instantiates a new TransactionDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetAccountOutput

`func (o *TransactionDestination) GetAccountOutput() TransactionTransferToAddressDestinationAccountOutput`

GetAccountOutput returns the AccountOutput field if non-nil, zero value otherwise.

### GetAccountOutputOk

`func (o *TransactionDestination) GetAccountOutputOk() (*TransactionTransferToAddressDestinationAccountOutput, bool)`

GetAccountOutputOk returns a tuple with the AccountOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOutput

`func (o *TransactionDestination) SetAccountOutput(v TransactionTransferToAddressDestinationAccountOutput)`

SetAccountOutput sets AccountOutput field to given value.

### HasAccountOutput

`func (o *TransactionDestination) HasAccountOutput() bool`

HasAccountOutput returns a boolean if a field has been set.

### GetUtxoOutputs

`func (o *TransactionDestination) GetUtxoOutputs() []TransactionTransferToAddressDestinationUtxoOutputsInner`

GetUtxoOutputs returns the UtxoOutputs field if non-nil, zero value otherwise.

### GetUtxoOutputsOk

`func (o *TransactionDestination) GetUtxoOutputsOk() (*[]TransactionTransferToAddressDestinationUtxoOutputsInner, bool)`

GetUtxoOutputsOk returns a tuple with the UtxoOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoOutputs

`func (o *TransactionDestination) SetUtxoOutputs(v []TransactionTransferToAddressDestinationUtxoOutputsInner)`

SetUtxoOutputs sets UtxoOutputs field to given value.

### HasUtxoOutputs

`func (o *TransactionDestination) HasUtxoOutputs() bool`

HasUtxoOutputs returns a boolean if a field has been set.

### GetChangeAddress

`func (o *TransactionDestination) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *TransactionDestination) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *TransactionDestination) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *TransactionDestination) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.

### GetForceInternal

`func (o *TransactionDestination) GetForceInternal() bool`

GetForceInternal returns the ForceInternal field if non-nil, zero value otherwise.

### GetForceInternalOk

`func (o *TransactionDestination) GetForceInternalOk() (*bool, bool)`

GetForceInternalOk returns a tuple with the ForceInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInternal

`func (o *TransactionDestination) SetForceInternal(v bool)`

SetForceInternal sets ForceInternal field to given value.

### HasForceInternal

`func (o *TransactionDestination) HasForceInternal() bool`

HasForceInternal returns a boolean if a field has been set.

### GetForceExternal

`func (o *TransactionDestination) GetForceExternal() bool`

GetForceExternal returns the ForceExternal field if non-nil, zero value otherwise.

### GetForceExternalOk

`func (o *TransactionDestination) GetForceExternalOk() (*bool, bool)`

GetForceExternalOk returns a tuple with the ForceExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExternal

`func (o *TransactionDestination) SetForceExternal(v bool)`

SetForceExternal sets ForceExternal field to given value.

### HasForceExternal

`func (o *TransactionDestination) HasForceExternal() bool`

HasForceExternal returns a boolean if a field has been set.

### GetWalletId

`func (o *TransactionDestination) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionDestination) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionDestination) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTradingAccountType

`func (o *TransactionDestination) GetTradingAccountType() string`

GetTradingAccountType returns the TradingAccountType field if non-nil, zero value otherwise.

### GetTradingAccountTypeOk

`func (o *TransactionDestination) GetTradingAccountTypeOk() (*string, bool)`

GetTradingAccountTypeOk returns a tuple with the TradingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountType

`func (o *TransactionDestination) SetTradingAccountType(v string)`

SetTradingAccountType sets TradingAccountType field to given value.

### HasTradingAccountType

`func (o *TransactionDestination) HasTradingAccountType() bool`

HasTradingAccountType returns a boolean if a field has been set.

### GetExchangeId

`func (o *TransactionDestination) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *TransactionDestination) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *TransactionDestination) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *TransactionDestination) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionDestination) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionDestination) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionDestination) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAddress

`func (o *TransactionDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionDestination) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetValue

`func (o *TransactionDestination) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransactionDestination) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransactionDestination) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TransactionDestination) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *TransactionDestination) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *TransactionDestination) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *TransactionDestination) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.


### GetMessage

`func (o *TransactionDestination) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionDestination) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionDestination) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStructuredData

`func (o *TransactionDestination) GetStructuredData() map[string]interface{}`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *TransactionDestination) GetStructuredDataOk() (*map[string]interface{}, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *TransactionDestination) SetStructuredData(v map[string]interface{})`

SetStructuredData sets StructuredData field to given value.


### GetWalletType

`func (o *TransactionDestination) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *TransactionDestination) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *TransactionDestination) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *TransactionDestination) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *TransactionDestination) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *TransactionDestination) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetMemo

`func (o *TransactionDestination) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionDestination) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionDestination) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionDestination) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


