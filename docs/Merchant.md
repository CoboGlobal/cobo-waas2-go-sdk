# Merchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. | 
**Name** | **string** | The merchant name. | 
**WalletId** | **string** | The ID of the linked wallet. | 
**DeveloperFeeRate** | Pointer to **string** | Developer fee rate for this token. For example, 0.01 represents a 1% fee.  | [optional] 
**WalletSetup** | Pointer to [**WalletSetup**](WalletSetup.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The created time of the merchant, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the merchant, represented as a UNIX timestamp in seconds. | [optional] 

## Methods

### NewMerchant

`func NewMerchant(merchantId string, name string, walletId string, ) *Merchant`

NewMerchant instantiates a new Merchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWithDefaults

`func NewMerchantWithDefaults() *Merchant`

NewMerchantWithDefaults instantiates a new Merchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *Merchant) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *Merchant) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *Merchant) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetName

`func (o *Merchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Merchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Merchant) SetName(v string)`

SetName sets Name field to given value.


### GetWalletId

`func (o *Merchant) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *Merchant) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *Merchant) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetDeveloperFeeRate

`func (o *Merchant) GetDeveloperFeeRate() string`

GetDeveloperFeeRate returns the DeveloperFeeRate field if non-nil, zero value otherwise.

### GetDeveloperFeeRateOk

`func (o *Merchant) GetDeveloperFeeRateOk() (*string, bool)`

GetDeveloperFeeRateOk returns a tuple with the DeveloperFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperFeeRate

`func (o *Merchant) SetDeveloperFeeRate(v string)`

SetDeveloperFeeRate sets DeveloperFeeRate field to given value.

### HasDeveloperFeeRate

`func (o *Merchant) HasDeveloperFeeRate() bool`

HasDeveloperFeeRate returns a boolean if a field has been set.

### GetWalletSetup

`func (o *Merchant) GetWalletSetup() WalletSetup`

GetWalletSetup returns the WalletSetup field if non-nil, zero value otherwise.

### GetWalletSetupOk

`func (o *Merchant) GetWalletSetupOk() (*WalletSetup, bool)`

GetWalletSetupOk returns a tuple with the WalletSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSetup

`func (o *Merchant) SetWalletSetup(v WalletSetup)`

SetWalletSetup sets WalletSetup field to given value.

### HasWalletSetup

`func (o *Merchant) HasWalletSetup() bool`

HasWalletSetup returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Merchant) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Merchant) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Merchant) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Merchant) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Merchant) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Merchant) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Merchant) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *Merchant) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


