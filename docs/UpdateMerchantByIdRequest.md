# UpdateMerchantByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The merchant name. | [optional] 
**DeveloperFeeRate** | Pointer to **string** | The developer fee rate applied to this merchant. Expressed as a string in decimal format where \&quot;0.1\&quot; represents 10%. This fee is deducted from the payment amount and only applies to top-up transactions. If you are a merchant (directly serving the payer), you do not need to configure the developer fee rate. | [optional] 

## Methods

### NewUpdateMerchantByIdRequest

`func NewUpdateMerchantByIdRequest() *UpdateMerchantByIdRequest`

NewUpdateMerchantByIdRequest instantiates a new UpdateMerchantByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMerchantByIdRequestWithDefaults

`func NewUpdateMerchantByIdRequestWithDefaults() *UpdateMerchantByIdRequest`

NewUpdateMerchantByIdRequestWithDefaults instantiates a new UpdateMerchantByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMerchantByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMerchantByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMerchantByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMerchantByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeveloperFeeRate

`func (o *UpdateMerchantByIdRequest) GetDeveloperFeeRate() string`

GetDeveloperFeeRate returns the DeveloperFeeRate field if non-nil, zero value otherwise.

### GetDeveloperFeeRateOk

`func (o *UpdateMerchantByIdRequest) GetDeveloperFeeRateOk() (*string, bool)`

GetDeveloperFeeRateOk returns a tuple with the DeveloperFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperFeeRate

`func (o *UpdateMerchantByIdRequest) SetDeveloperFeeRate(v string)`

SetDeveloperFeeRate sets DeveloperFeeRate field to given value.

### HasDeveloperFeeRate

`func (o *UpdateMerchantByIdRequest) HasDeveloperFeeRate() bool`

HasDeveloperFeeRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


