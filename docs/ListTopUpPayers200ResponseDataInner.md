# ListTopUpPayers200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. | 
**PayerId** | **string** | A unique identifier assigned by Cobo to track and identify individual payers. | 
**DeveloperFeeRate** | **string** | The developer fee rate applied to top-up transactions made by this payer. Expressed as a decimal string where \&quot;0.1\&quot; represents 10%. | 
**CreatedTimestamp** | Pointer to **int32** | The creation time of the payer, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The last update time of the payer, represented as a UNIX timestamp in seconds. | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of top-up transactions associated with this payer. | [optional] 

## Methods

### NewListTopUpPayers200ResponseDataInner

`func NewListTopUpPayers200ResponseDataInner(merchantId string, payerId string, developerFeeRate string, ) *ListTopUpPayers200ResponseDataInner`

NewListTopUpPayers200ResponseDataInner instantiates a new ListTopUpPayers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTopUpPayers200ResponseDataInnerWithDefaults

`func NewListTopUpPayers200ResponseDataInnerWithDefaults() *ListTopUpPayers200ResponseDataInner`

NewListTopUpPayers200ResponseDataInnerWithDefaults instantiates a new ListTopUpPayers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *ListTopUpPayers200ResponseDataInner) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *ListTopUpPayers200ResponseDataInner) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *ListTopUpPayers200ResponseDataInner) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetPayerId

`func (o *ListTopUpPayers200ResponseDataInner) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *ListTopUpPayers200ResponseDataInner) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *ListTopUpPayers200ResponseDataInner) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetDeveloperFeeRate

`func (o *ListTopUpPayers200ResponseDataInner) GetDeveloperFeeRate() string`

GetDeveloperFeeRate returns the DeveloperFeeRate field if non-nil, zero value otherwise.

### GetDeveloperFeeRateOk

`func (o *ListTopUpPayers200ResponseDataInner) GetDeveloperFeeRateOk() (*string, bool)`

GetDeveloperFeeRateOk returns a tuple with the DeveloperFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperFeeRate

`func (o *ListTopUpPayers200ResponseDataInner) SetDeveloperFeeRate(v string)`

SetDeveloperFeeRate sets DeveloperFeeRate field to given value.


### GetCreatedTimestamp

`func (o *ListTopUpPayers200ResponseDataInner) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ListTopUpPayers200ResponseDataInner) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ListTopUpPayers200ResponseDataInner) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ListTopUpPayers200ResponseDataInner) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ListTopUpPayers200ResponseDataInner) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ListTopUpPayers200ResponseDataInner) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ListTopUpPayers200ResponseDataInner) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *ListTopUpPayers200ResponseDataInner) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetTransactions

`func (o *ListTopUpPayers200ResponseDataInner) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ListTopUpPayers200ResponseDataInner) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ListTopUpPayers200ResponseDataInner) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ListTopUpPayers200ResponseDataInner) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


