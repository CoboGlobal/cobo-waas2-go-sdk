# WalletInfoEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. | 
**Wallet** | Pointer to [**WalletInfo**](WalletInfo.md) |  | [optional] 

## Methods

### NewWalletInfoEventData

`func NewWalletInfoEventData(dataType string, ) *WalletInfoEventData`

NewWalletInfoEventData instantiates a new WalletInfoEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletInfoEventDataWithDefaults

`func NewWalletInfoEventDataWithDefaults() *WalletInfoEventData`

NewWalletInfoEventDataWithDefaults instantiates a new WalletInfoEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *WalletInfoEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *WalletInfoEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *WalletInfoEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetWallet

`func (o *WalletInfoEventData) GetWallet() WalletInfo`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *WalletInfoEventData) GetWalletOk() (*WalletInfo, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *WalletInfoEventData) SetWallet(v WalletInfo)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *WalletInfoEventData) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


