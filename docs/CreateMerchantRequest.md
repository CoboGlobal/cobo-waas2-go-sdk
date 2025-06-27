# CreateMerchantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The merchant name. | 
**WalletId** | **string** | The ID of the wallet linked to the merchant. | 
**DeveloperFeeRate** | Pointer to **string** | The fee rate applied when topping up the merchant account. Represented as a string percentage (e.g., \&quot;0.1\&quot; means 10%). | [optional] 

## Methods

### NewCreateMerchantRequest

`func NewCreateMerchantRequest(name string, walletId string, ) *CreateMerchantRequest`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


