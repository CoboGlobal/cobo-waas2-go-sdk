# CheckAddressesValidity200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The wallet address. | 
**Validity** | **bool** | Whether the address is valid.  - &#x60;true&#x60;: The address is valid.  - &#x60;false&#x60;: The address is invalid.  | 

## Methods

### NewCheckAddressesValidity200ResponseInner

`func NewCheckAddressesValidity200ResponseInner(address string, validity bool, ) *CheckAddressesValidity200ResponseInner`

NewCheckAddressesValidity200ResponseInner instantiates a new CheckAddressesValidity200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAddressesValidity200ResponseInnerWithDefaults

`func NewCheckAddressesValidity200ResponseInnerWithDefaults() *CheckAddressesValidity200ResponseInner`

NewCheckAddressesValidity200ResponseInnerWithDefaults instantiates a new CheckAddressesValidity200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CheckAddressesValidity200ResponseInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CheckAddressesValidity200ResponseInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CheckAddressesValidity200ResponseInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetValidity

`func (o *CheckAddressesValidity200ResponseInner) GetValidity() bool`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *CheckAddressesValidity200ResponseInner) GetValidityOk() (*bool, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *CheckAddressesValidity200ResponseInner) SetValidity(v bool)`

SetValidity sets Validity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


