# ComplianceKyaScreeningsUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**ScreeningId** | **string** | The unique system-generated identifier for this screening request (UUID format, fixed 36 characters). | 
**Address** | **string** | The screened blockchain address. | 
**ChainId** | **string** | The chain identifier. | 
**Status** | [**KyaScreeningStatus**](KyaScreeningStatus.md) |  | 
**CreatedTimestamp** | **int64** | The time when the screening request was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTimestamp** | **int64** | The time when the screening status was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewComplianceKyaScreeningsUpdateEventData

`func NewComplianceKyaScreeningsUpdateEventData(dataType string, screeningId string, address string, chainId string, status KyaScreeningStatus, createdTimestamp int64, updatedTimestamp int64, ) *ComplianceKyaScreeningsUpdateEventData`

NewComplianceKyaScreeningsUpdateEventData instantiates a new ComplianceKyaScreeningsUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceKyaScreeningsUpdateEventDataWithDefaults

`func NewComplianceKyaScreeningsUpdateEventDataWithDefaults() *ComplianceKyaScreeningsUpdateEventData`

NewComplianceKyaScreeningsUpdateEventDataWithDefaults instantiates a new ComplianceKyaScreeningsUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *ComplianceKyaScreeningsUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ComplianceKyaScreeningsUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ComplianceKyaScreeningsUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetScreeningId

`func (o *ComplianceKyaScreeningsUpdateEventData) GetScreeningId() string`

GetScreeningId returns the ScreeningId field if non-nil, zero value otherwise.

### GetScreeningIdOk

`func (o *ComplianceKyaScreeningsUpdateEventData) GetScreeningIdOk() (*string, bool)`

GetScreeningIdOk returns a tuple with the ScreeningId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningId

`func (o *ComplianceKyaScreeningsUpdateEventData) SetScreeningId(v string)`

SetScreeningId sets ScreeningId field to given value.


### GetAddress

`func (o *ComplianceKyaScreeningsUpdateEventData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ComplianceKyaScreeningsUpdateEventData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ComplianceKyaScreeningsUpdateEventData) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *ComplianceKyaScreeningsUpdateEventData) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ComplianceKyaScreeningsUpdateEventData) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ComplianceKyaScreeningsUpdateEventData) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetStatus

`func (o *ComplianceKyaScreeningsUpdateEventData) GetStatus() KyaScreeningStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComplianceKyaScreeningsUpdateEventData) GetStatusOk() (*KyaScreeningStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComplianceKyaScreeningsUpdateEventData) SetStatus(v KyaScreeningStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *ComplianceKyaScreeningsUpdateEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ComplianceKyaScreeningsUpdateEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ComplianceKyaScreeningsUpdateEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ComplianceKyaScreeningsUpdateEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ComplianceKyaScreeningsUpdateEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ComplianceKyaScreeningsUpdateEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


