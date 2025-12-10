# UpdateMerchantByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The merchant name. | [optional] 
**DeveloperFeeRate** | Pointer to **string** | The fee rate applied when topping up the merchant account. Represented as a string percentage (e.g., \&quot;0.1\&quot; means 10%). | [optional] 
**SubscriptionDeveloperFeeRate** | Pointer to **string** | The fee rate applied when subscribe the merchant account. Represented as a string percentage (e.g., \&quot;0.1\&quot; means 10%). | [optional] 

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

### GetSubscriptionDeveloperFeeRate

`func (o *UpdateMerchantByIdRequest) GetSubscriptionDeveloperFeeRate() string`

GetSubscriptionDeveloperFeeRate returns the SubscriptionDeveloperFeeRate field if non-nil, zero value otherwise.

### GetSubscriptionDeveloperFeeRateOk

`func (o *UpdateMerchantByIdRequest) GetSubscriptionDeveloperFeeRateOk() (*string, bool)`

GetSubscriptionDeveloperFeeRateOk returns a tuple with the SubscriptionDeveloperFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDeveloperFeeRate

`func (o *UpdateMerchantByIdRequest) SetSubscriptionDeveloperFeeRate(v string)`

SetSubscriptionDeveloperFeeRate sets SubscriptionDeveloperFeeRate field to given value.

### HasSubscriptionDeveloperFeeRate

`func (o *UpdateMerchantByIdRequest) HasSubscriptionDeveloperFeeRate() bool`

HasSubscriptionDeveloperFeeRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


