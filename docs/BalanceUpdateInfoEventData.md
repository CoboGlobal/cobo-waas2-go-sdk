# BalanceUpdateInfoEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
**Address** | **string** | The wallet address. | 
**WalletUuid** | **string** | The wallet ID. | 
**UpdatedTimestamp** | **int64** | The time when the balance updated, in Unix timestamp format, measured in milliseconds.  | 
**Balance** | [**Balance**](Balance.md) |  | 

## Methods

### NewBalanceUpdateInfoEventData

`func NewBalanceUpdateInfoEventData(dataType string, tokenId string, address string, walletUuid string, updatedTimestamp int64, balance Balance, ) *BalanceUpdateInfoEventData`

NewBalanceUpdateInfoEventData instantiates a new BalanceUpdateInfoEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceUpdateInfoEventDataWithDefaults

`func NewBalanceUpdateInfoEventDataWithDefaults() *BalanceUpdateInfoEventData`

NewBalanceUpdateInfoEventDataWithDefaults instantiates a new BalanceUpdateInfoEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *BalanceUpdateInfoEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *BalanceUpdateInfoEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *BalanceUpdateInfoEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetTokenId

`func (o *BalanceUpdateInfoEventData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BalanceUpdateInfoEventData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BalanceUpdateInfoEventData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAddress

`func (o *BalanceUpdateInfoEventData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BalanceUpdateInfoEventData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BalanceUpdateInfoEventData) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetWalletUuid

`func (o *BalanceUpdateInfoEventData) GetWalletUuid() string`

GetWalletUuid returns the WalletUuid field if non-nil, zero value otherwise.

### GetWalletUuidOk

`func (o *BalanceUpdateInfoEventData) GetWalletUuidOk() (*string, bool)`

GetWalletUuidOk returns a tuple with the WalletUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletUuid

`func (o *BalanceUpdateInfoEventData) SetWalletUuid(v string)`

SetWalletUuid sets WalletUuid field to given value.


### GetUpdatedTimestamp

`func (o *BalanceUpdateInfoEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *BalanceUpdateInfoEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *BalanceUpdateInfoEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetBalance

`func (o *BalanceUpdateInfoEventData) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceUpdateInfoEventData) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceUpdateInfoEventData) SetBalance(v Balance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


