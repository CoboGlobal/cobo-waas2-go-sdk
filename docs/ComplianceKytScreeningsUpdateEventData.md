# ComplianceKytScreeningsUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**TransactionId** | **string** | The transaction ID. | 
**TransactionType** | [**KytScreeningsTransactionType**](KytScreeningsTransactionType.md) |  | 
**ReviewStatus** | [**ReviewStatusType**](ReviewStatusType.md) |  | 
**FundsStatus** | [**FundsStatusType**](FundsStatusType.md) |  | 
**UpdatedTimestamp** | **int64** | The time when the KYT screening information was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewComplianceKytScreeningsUpdateEventData

`func NewComplianceKytScreeningsUpdateEventData(dataType string, transactionId string, transactionType KytScreeningsTransactionType, reviewStatus ReviewStatusType, fundsStatus FundsStatusType, updatedTimestamp int64, ) *ComplianceKytScreeningsUpdateEventData`

NewComplianceKytScreeningsUpdateEventData instantiates a new ComplianceKytScreeningsUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceKytScreeningsUpdateEventDataWithDefaults

`func NewComplianceKytScreeningsUpdateEventDataWithDefaults() *ComplianceKytScreeningsUpdateEventData`

NewComplianceKytScreeningsUpdateEventDataWithDefaults instantiates a new ComplianceKytScreeningsUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *ComplianceKytScreeningsUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ComplianceKytScreeningsUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ComplianceKytScreeningsUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetTransactionId

`func (o *ComplianceKytScreeningsUpdateEventData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ComplianceKytScreeningsUpdateEventData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ComplianceKytScreeningsUpdateEventData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionType

`func (o *ComplianceKytScreeningsUpdateEventData) GetTransactionType() KytScreeningsTransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ComplianceKytScreeningsUpdateEventData) GetTransactionTypeOk() (*KytScreeningsTransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ComplianceKytScreeningsUpdateEventData) SetTransactionType(v KytScreeningsTransactionType)`

SetTransactionType sets TransactionType field to given value.


### GetReviewStatus

`func (o *ComplianceKytScreeningsUpdateEventData) GetReviewStatus() ReviewStatusType`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *ComplianceKytScreeningsUpdateEventData) GetReviewStatusOk() (*ReviewStatusType, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *ComplianceKytScreeningsUpdateEventData) SetReviewStatus(v ReviewStatusType)`

SetReviewStatus sets ReviewStatus field to given value.


### GetFundsStatus

`func (o *ComplianceKytScreeningsUpdateEventData) GetFundsStatus() FundsStatusType`

GetFundsStatus returns the FundsStatus field if non-nil, zero value otherwise.

### GetFundsStatusOk

`func (o *ComplianceKytScreeningsUpdateEventData) GetFundsStatusOk() (*FundsStatusType, bool)`

GetFundsStatusOk returns a tuple with the FundsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsStatus

`func (o *ComplianceKytScreeningsUpdateEventData) SetFundsStatus(v FundsStatusType)`

SetFundsStatus sets FundsStatus field to given value.


### GetUpdatedTimestamp

`func (o *ComplianceKytScreeningsUpdateEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ComplianceKytScreeningsUpdateEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ComplianceKytScreeningsUpdateEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


