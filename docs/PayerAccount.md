# PayerAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. | 
**PayerId** | **string** | A unique identifier assigned by Cobo to track and identify individual payers. | 
**DeveloperFeeRate** | **string** | The developer fee rate applied to the top-up transactions made by the payer. Expressed as a decimal string where \&quot;0.1\&quot; represents 10%. | 
**CreatedTimestamp** | Pointer to **int32** | The creation time of the payer, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The last update time of the payer, represented as a UNIX timestamp in seconds. | [optional] 
**Accounts** | Pointer to [**[]Account**](Account.md) | An array of accounts associated with this payer. | [optional] 

## Methods

### NewPayerAccount

`func NewPayerAccount(merchantId string, payerId string, developerFeeRate string, ) *PayerAccount`

NewPayerAccount instantiates a new PayerAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayerAccountWithDefaults

`func NewPayerAccountWithDefaults() *PayerAccount`

NewPayerAccountWithDefaults instantiates a new PayerAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *PayerAccount) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PayerAccount) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PayerAccount) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetPayerId

`func (o *PayerAccount) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *PayerAccount) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *PayerAccount) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetDeveloperFeeRate

`func (o *PayerAccount) GetDeveloperFeeRate() string`

GetDeveloperFeeRate returns the DeveloperFeeRate field if non-nil, zero value otherwise.

### GetDeveloperFeeRateOk

`func (o *PayerAccount) GetDeveloperFeeRateOk() (*string, bool)`

GetDeveloperFeeRateOk returns a tuple with the DeveloperFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperFeeRate

`func (o *PayerAccount) SetDeveloperFeeRate(v string)`

SetDeveloperFeeRate sets DeveloperFeeRate field to given value.


### GetCreatedTimestamp

`func (o *PayerAccount) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PayerAccount) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PayerAccount) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PayerAccount) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *PayerAccount) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PayerAccount) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PayerAccount) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PayerAccount) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetAccounts

`func (o *PayerAccount) GetAccounts() []Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PayerAccount) GetAccountsOk() (*[]Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PayerAccount) SetAccounts(v []Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PayerAccount) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


