# PaymentAccountBalanceUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBankWithdrawal&#x60;: The payment bank withdrawal event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;PaymentBulkSendItem&#x60;: The payment bulk send item event data. - &#x60;PaymentAccountBalanceUpdate&#x60;: The Payments account balance updated event data, including account information and balance change details. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. - &#x60;Organization&#x60;: The organization event data. - &#x60;FiatTransaction&#x60;: The fiat transaction event data. | 
**SourceAccount** | **string** | The source account of the balance change. This field uses the same semantics as &#x60;source_account&#x60; in [List balance changes](https://www.cobo.com/developers/v2/api-references/payment/list-balance-changes). - When the account is a merchant account, this is the merchant ID (merchant code), which you can retrieve by calling [List all merchants](https://www.cobo.com/developers/v2/api-references/payment/list-all-merchants). - When the account is the developer account, use &#x60;developer&#x60;.  | 
**SourceId** | **string** | The source ID of the balance change. | 
**SourceType** | [**PaymentBalanceChangeSourceType**](PaymentBalanceChangeSourceType.md) |  | 
**TokenId** | **string** | The token ID of the balance change. | 
**Amount** | **string** | The balance change amount, truncated to two decimal places and represented as a numeric string. | 
**AmountRaw** | **string** | The balance change amount in the token&#39;s decimal precision, represented as a numeric string. | 
**BalanceBefore** | **string** | The account balance before the balance change, truncated to two decimal places and represented as a numeric string. | 
**BalanceBeforeRaw** | **string** | The account balance before the balance change in the token&#39;s decimal precision, represented as a numeric string. | 
**BalanceAfter** | **string** | The account balance after the balance change, truncated to two decimal places and represented as a numeric string. | 
**BalanceAfterRaw** | **string** | The account balance after the balance change in the token&#39;s decimal precision, represented as a numeric string. | 
**FlowDirection** | [**PaymentBalanceFlowDirection**](PaymentBalanceFlowDirection.md) |  | 
**UpdateTime** | **int64** | The time when the balance was updated, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentAccountBalanceUpdateEventData

`func NewPaymentAccountBalanceUpdateEventData(dataType string, sourceAccount string, sourceId string, sourceType PaymentBalanceChangeSourceType, tokenId string, amount string, amountRaw string, balanceBefore string, balanceBeforeRaw string, balanceAfter string, balanceAfterRaw string, flowDirection PaymentBalanceFlowDirection, updateTime int64, ) *PaymentAccountBalanceUpdateEventData`

NewPaymentAccountBalanceUpdateEventData instantiates a new PaymentAccountBalanceUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAccountBalanceUpdateEventDataWithDefaults

`func NewPaymentAccountBalanceUpdateEventDataWithDefaults() *PaymentAccountBalanceUpdateEventData`

NewPaymentAccountBalanceUpdateEventDataWithDefaults instantiates a new PaymentAccountBalanceUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentAccountBalanceUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentAccountBalanceUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentAccountBalanceUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetSourceAccount

`func (o *PaymentAccountBalanceUpdateEventData) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PaymentAccountBalanceUpdateEventData) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PaymentAccountBalanceUpdateEventData) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetSourceId

`func (o *PaymentAccountBalanceUpdateEventData) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PaymentAccountBalanceUpdateEventData) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PaymentAccountBalanceUpdateEventData) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceType

`func (o *PaymentAccountBalanceUpdateEventData) GetSourceType() PaymentBalanceChangeSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PaymentAccountBalanceUpdateEventData) GetSourceTypeOk() (*PaymentBalanceChangeSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PaymentAccountBalanceUpdateEventData) SetSourceType(v PaymentBalanceChangeSourceType)`

SetSourceType sets SourceType field to given value.


### GetTokenId

`func (o *PaymentAccountBalanceUpdateEventData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentAccountBalanceUpdateEventData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentAccountBalanceUpdateEventData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *PaymentAccountBalanceUpdateEventData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAccountBalanceUpdateEventData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAccountBalanceUpdateEventData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAmountRaw

`func (o *PaymentAccountBalanceUpdateEventData) GetAmountRaw() string`

GetAmountRaw returns the AmountRaw field if non-nil, zero value otherwise.

### GetAmountRawOk

`func (o *PaymentAccountBalanceUpdateEventData) GetAmountRawOk() (*string, bool)`

GetAmountRawOk returns a tuple with the AmountRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRaw

`func (o *PaymentAccountBalanceUpdateEventData) SetAmountRaw(v string)`

SetAmountRaw sets AmountRaw field to given value.


### GetBalanceBefore

`func (o *PaymentAccountBalanceUpdateEventData) GetBalanceBefore() string`

GetBalanceBefore returns the BalanceBefore field if non-nil, zero value otherwise.

### GetBalanceBeforeOk

`func (o *PaymentAccountBalanceUpdateEventData) GetBalanceBeforeOk() (*string, bool)`

GetBalanceBeforeOk returns a tuple with the BalanceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBefore

`func (o *PaymentAccountBalanceUpdateEventData) SetBalanceBefore(v string)`

SetBalanceBefore sets BalanceBefore field to given value.


### GetBalanceBeforeRaw

`func (o *PaymentAccountBalanceUpdateEventData) GetBalanceBeforeRaw() string`

GetBalanceBeforeRaw returns the BalanceBeforeRaw field if non-nil, zero value otherwise.

### GetBalanceBeforeRawOk

`func (o *PaymentAccountBalanceUpdateEventData) GetBalanceBeforeRawOk() (*string, bool)`

GetBalanceBeforeRawOk returns a tuple with the BalanceBeforeRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBeforeRaw

`func (o *PaymentAccountBalanceUpdateEventData) SetBalanceBeforeRaw(v string)`

SetBalanceBeforeRaw sets BalanceBeforeRaw field to given value.


### GetBalanceAfter

`func (o *PaymentAccountBalanceUpdateEventData) GetBalanceAfter() string`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *PaymentAccountBalanceUpdateEventData) GetBalanceAfterOk() (*string, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *PaymentAccountBalanceUpdateEventData) SetBalanceAfter(v string)`

SetBalanceAfter sets BalanceAfter field to given value.


### GetBalanceAfterRaw

`func (o *PaymentAccountBalanceUpdateEventData) GetBalanceAfterRaw() string`

GetBalanceAfterRaw returns the BalanceAfterRaw field if non-nil, zero value otherwise.

### GetBalanceAfterRawOk

`func (o *PaymentAccountBalanceUpdateEventData) GetBalanceAfterRawOk() (*string, bool)`

GetBalanceAfterRawOk returns a tuple with the BalanceAfterRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfterRaw

`func (o *PaymentAccountBalanceUpdateEventData) SetBalanceAfterRaw(v string)`

SetBalanceAfterRaw sets BalanceAfterRaw field to given value.


### GetFlowDirection

`func (o *PaymentAccountBalanceUpdateEventData) GetFlowDirection() PaymentBalanceFlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *PaymentAccountBalanceUpdateEventData) GetFlowDirectionOk() (*PaymentBalanceFlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *PaymentAccountBalanceUpdateEventData) SetFlowDirection(v PaymentBalanceFlowDirection)`

SetFlowDirection sets FlowDirection field to given value.


### GetUpdateTime

`func (o *PaymentAccountBalanceUpdateEventData) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PaymentAccountBalanceUpdateEventData) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PaymentAccountBalanceUpdateEventData) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


