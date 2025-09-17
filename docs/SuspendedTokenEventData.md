# SuspendedTokenEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. | 
**TokenIds** | **string** | A list of token IDs, separated by comma. | 
**OperationType** | [**SuspendedTokenOperationType**](SuspendedTokenOperationType.md) |  | 

## Methods

### NewSuspendedTokenEventData

`func NewSuspendedTokenEventData(dataType string, tokenIds string, operationType SuspendedTokenOperationType, ) *SuspendedTokenEventData`

NewSuspendedTokenEventData instantiates a new SuspendedTokenEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuspendedTokenEventDataWithDefaults

`func NewSuspendedTokenEventDataWithDefaults() *SuspendedTokenEventData`

NewSuspendedTokenEventDataWithDefaults instantiates a new SuspendedTokenEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *SuspendedTokenEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *SuspendedTokenEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *SuspendedTokenEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetTokenIds

`func (o *SuspendedTokenEventData) GetTokenIds() string`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *SuspendedTokenEventData) GetTokenIdsOk() (*string, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *SuspendedTokenEventData) SetTokenIds(v string)`

SetTokenIds sets TokenIds field to given value.


### GetOperationType

`func (o *SuspendedTokenEventData) GetOperationType() SuspendedTokenOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *SuspendedTokenEventData) GetOperationTypeOk() (*SuspendedTokenOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *SuspendedTokenEventData) SetOperationType(v SuspendedTokenOperationType)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


