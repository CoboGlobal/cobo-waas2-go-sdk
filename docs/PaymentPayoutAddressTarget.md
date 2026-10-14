# PaymentPayoutAddressTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetType** | **string** | The type of the payout target. | 
**Address** | **string** | The recipient address. | 
**ChainId** | **string** | The chain ID of the recipient address. | 

## Methods

### NewPaymentPayoutAddressTarget

`func NewPaymentPayoutAddressTarget(targetType string, address string, chainId string, ) *PaymentPayoutAddressTarget`

NewPaymentPayoutAddressTarget instantiates a new PaymentPayoutAddressTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayoutAddressTargetWithDefaults

`func NewPaymentPayoutAddressTargetWithDefaults() *PaymentPayoutAddressTarget`

NewPaymentPayoutAddressTargetWithDefaults instantiates a new PaymentPayoutAddressTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetType

`func (o *PaymentPayoutAddressTarget) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *PaymentPayoutAddressTarget) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *PaymentPayoutAddressTarget) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetAddress

`func (o *PaymentPayoutAddressTarget) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentPayoutAddressTarget) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentPayoutAddressTarget) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *PaymentPayoutAddressTarget) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PaymentPayoutAddressTarget) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PaymentPayoutAddressTarget) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


