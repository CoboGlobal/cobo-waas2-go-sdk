# PaymentAddressUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**CustomPayerId** | **string** | A unique identifier assigned by the developer to track and identify individual payers in their system. | 
**PayerId** | **string** | A unique identifier assigned by Cobo to track and identify individual payers. | 
**Chain** | **string** | The chain ID. | 
**PreviousAddress** | **string** | The previous top-up address that was assigned to the payer. | 
**UpdatedAddress** | **string** | The new top-up address that has been assigned to the payer. | 

## Methods

### NewPaymentAddressUpdateEventData

`func NewPaymentAddressUpdateEventData(dataType string, customPayerId string, payerId string, chain string, previousAddress string, updatedAddress string, ) *PaymentAddressUpdateEventData`

NewPaymentAddressUpdateEventData instantiates a new PaymentAddressUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAddressUpdateEventDataWithDefaults

`func NewPaymentAddressUpdateEventDataWithDefaults() *PaymentAddressUpdateEventData`

NewPaymentAddressUpdateEventDataWithDefaults instantiates a new PaymentAddressUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentAddressUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentAddressUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentAddressUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetCustomPayerId

`func (o *PaymentAddressUpdateEventData) GetCustomPayerId() string`

GetCustomPayerId returns the CustomPayerId field if non-nil, zero value otherwise.

### GetCustomPayerIdOk

`func (o *PaymentAddressUpdateEventData) GetCustomPayerIdOk() (*string, bool)`

GetCustomPayerIdOk returns a tuple with the CustomPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayerId

`func (o *PaymentAddressUpdateEventData) SetCustomPayerId(v string)`

SetCustomPayerId sets CustomPayerId field to given value.


### GetPayerId

`func (o *PaymentAddressUpdateEventData) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *PaymentAddressUpdateEventData) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *PaymentAddressUpdateEventData) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetChain

`func (o *PaymentAddressUpdateEventData) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *PaymentAddressUpdateEventData) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *PaymentAddressUpdateEventData) SetChain(v string)`

SetChain sets Chain field to given value.


### GetPreviousAddress

`func (o *PaymentAddressUpdateEventData) GetPreviousAddress() string`

GetPreviousAddress returns the PreviousAddress field if non-nil, zero value otherwise.

### GetPreviousAddressOk

`func (o *PaymentAddressUpdateEventData) GetPreviousAddressOk() (*string, bool)`

GetPreviousAddressOk returns a tuple with the PreviousAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAddress

`func (o *PaymentAddressUpdateEventData) SetPreviousAddress(v string)`

SetPreviousAddress sets PreviousAddress field to given value.


### GetUpdatedAddress

`func (o *PaymentAddressUpdateEventData) GetUpdatedAddress() string`

GetUpdatedAddress returns the UpdatedAddress field if non-nil, zero value otherwise.

### GetUpdatedAddressOk

`func (o *PaymentAddressUpdateEventData) GetUpdatedAddressOk() (*string, bool)`

GetUpdatedAddressOk returns a tuple with the UpdatedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAddress

`func (o *PaymentAddressUpdateEventData) SetUpdatedAddress(v string)`

SetUpdatedAddress sets UpdatedAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


