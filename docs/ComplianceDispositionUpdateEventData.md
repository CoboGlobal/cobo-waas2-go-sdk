# ComplianceDispositionUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**TransactionId** | **string** | The transaction ID. | 
**DispositionType** | [**DispositionType**](DispositionType.md) |  | 
**DispositionStatus** | [**DispositionStatus**](DispositionStatus.md) |  | 
**DestinationAddress** | Pointer to **string** | The blockchain address to receive the refunded/isolated funds. | [optional] 
**DispositionAmount** | Pointer to **string** | The amount to be refunded/isolated from the original transaction, specified as a numeric string. This value cannot exceed the total amount of the original transaction.  | [optional] 
**UpdatedTimestamp** | **int64** | The time when the disposition was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewComplianceDispositionUpdateEventData

`func NewComplianceDispositionUpdateEventData(dataType string, transactionId string, dispositionType DispositionType, dispositionStatus DispositionStatus, updatedTimestamp int64, ) *ComplianceDispositionUpdateEventData`

NewComplianceDispositionUpdateEventData instantiates a new ComplianceDispositionUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceDispositionUpdateEventDataWithDefaults

`func NewComplianceDispositionUpdateEventDataWithDefaults() *ComplianceDispositionUpdateEventData`

NewComplianceDispositionUpdateEventDataWithDefaults instantiates a new ComplianceDispositionUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *ComplianceDispositionUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ComplianceDispositionUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ComplianceDispositionUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetTransactionId

`func (o *ComplianceDispositionUpdateEventData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ComplianceDispositionUpdateEventData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ComplianceDispositionUpdateEventData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetDispositionType

`func (o *ComplianceDispositionUpdateEventData) GetDispositionType() DispositionType`

GetDispositionType returns the DispositionType field if non-nil, zero value otherwise.

### GetDispositionTypeOk

`func (o *ComplianceDispositionUpdateEventData) GetDispositionTypeOk() (*DispositionType, bool)`

GetDispositionTypeOk returns a tuple with the DispositionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionType

`func (o *ComplianceDispositionUpdateEventData) SetDispositionType(v DispositionType)`

SetDispositionType sets DispositionType field to given value.


### GetDispositionStatus

`func (o *ComplianceDispositionUpdateEventData) GetDispositionStatus() DispositionStatus`

GetDispositionStatus returns the DispositionStatus field if non-nil, zero value otherwise.

### GetDispositionStatusOk

`func (o *ComplianceDispositionUpdateEventData) GetDispositionStatusOk() (*DispositionStatus, bool)`

GetDispositionStatusOk returns a tuple with the DispositionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionStatus

`func (o *ComplianceDispositionUpdateEventData) SetDispositionStatus(v DispositionStatus)`

SetDispositionStatus sets DispositionStatus field to given value.


### GetDestinationAddress

`func (o *ComplianceDispositionUpdateEventData) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *ComplianceDispositionUpdateEventData) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *ComplianceDispositionUpdateEventData) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *ComplianceDispositionUpdateEventData) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDispositionAmount

`func (o *ComplianceDispositionUpdateEventData) GetDispositionAmount() string`

GetDispositionAmount returns the DispositionAmount field if non-nil, zero value otherwise.

### GetDispositionAmountOk

`func (o *ComplianceDispositionUpdateEventData) GetDispositionAmountOk() (*string, bool)`

GetDispositionAmountOk returns a tuple with the DispositionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionAmount

`func (o *ComplianceDispositionUpdateEventData) SetDispositionAmount(v string)`

SetDispositionAmount sets DispositionAmount field to given value.

### HasDispositionAmount

`func (o *ComplianceDispositionUpdateEventData) HasDispositionAmount() bool`

HasDispositionAmount returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ComplianceDispositionUpdateEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ComplianceDispositionUpdateEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ComplianceDispositionUpdateEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


