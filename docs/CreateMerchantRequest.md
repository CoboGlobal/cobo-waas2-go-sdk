# CreateMerchantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The merchant name. | 
**WalletId** | Pointer to **string** | This field has been deprecated. | [optional] 
**DeveloperFeeRate** | Pointer to **string** | The developer fee rate applied to this merchant. Must be a valid float between 0 and 1 (inclusive), with up to 4 decimal places. For more information on developer fee rate, please refer to [Funds allocation and balances](https://www.cobo.com/developers/v2/payments/amounts-and-balances). | [optional] 
**WalletSetup** | Pointer to [**WalletSetup**](WalletSetup.md) |  | [optional] 

## Methods

### NewCreateMerchantRequest

`func NewCreateMerchantRequest(name string, ) *CreateMerchantRequest`

NewCreateMerchantRequest instantiates a new CreateMerchantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMerchantRequestWithDefaults

`func NewCreateMerchantRequestWithDefaults() *CreateMerchantRequest`

NewCreateMerchantRequestWithDefaults instantiates a new CreateMerchantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMerchantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMerchantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMerchantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetWalletId

`func (o *CreateMerchantRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateMerchantRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateMerchantRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *CreateMerchantRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetDeveloperFeeRate

`func (o *CreateMerchantRequest) GetDeveloperFeeRate() string`

GetDeveloperFeeRate returns the DeveloperFeeRate field if non-nil, zero value otherwise.

### GetDeveloperFeeRateOk

`func (o *CreateMerchantRequest) GetDeveloperFeeRateOk() (*string, bool)`

GetDeveloperFeeRateOk returns a tuple with the DeveloperFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperFeeRate

`func (o *CreateMerchantRequest) SetDeveloperFeeRate(v string)`

SetDeveloperFeeRate sets DeveloperFeeRate field to given value.

### HasDeveloperFeeRate

`func (o *CreateMerchantRequest) HasDeveloperFeeRate() bool`

HasDeveloperFeeRate returns a boolean if a field has been set.

### GetWalletSetup

`func (o *CreateMerchantRequest) GetWalletSetup() WalletSetup`

GetWalletSetup returns the WalletSetup field if non-nil, zero value otherwise.

### GetWalletSetupOk

`func (o *CreateMerchantRequest) GetWalletSetupOk() (*WalletSetup, bool)`

GetWalletSetupOk returns a tuple with the WalletSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSetup

`func (o *CreateMerchantRequest) SetWalletSetup(v WalletSetup)`

SetWalletSetup sets WalletSetup field to given value.

### HasWalletSetup

`func (o *CreateMerchantRequest) HasWalletSetup() bool`

HasWalletSetup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


