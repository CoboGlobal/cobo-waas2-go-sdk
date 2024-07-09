# EigenlayerValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IconUrl** | **string** | The URL of the validator&#39;s icon. | 
**Name** | **string** | The name of validator. | 
**Priority** | Pointer to **float32** | The priority of validator. | [optional] 
**Address** | **string** | The address of validator. | 
**CommissionRate** | Pointer to **float32** | The commission of validator. | [optional] 

## Methods

### NewEigenlayerValidator

`func NewEigenlayerValidator(iconUrl string, name string, address string, ) *EigenlayerValidator`

NewEigenlayerValidator instantiates a new EigenlayerValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEigenlayerValidatorWithDefaults

`func NewEigenlayerValidatorWithDefaults() *EigenlayerValidator`

NewEigenlayerValidatorWithDefaults instantiates a new EigenlayerValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIconUrl

`func (o *EigenlayerValidator) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *EigenlayerValidator) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *EigenlayerValidator) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetName

`func (o *EigenlayerValidator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EigenlayerValidator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EigenlayerValidator) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *EigenlayerValidator) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EigenlayerValidator) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EigenlayerValidator) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EigenlayerValidator) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAddress

`func (o *EigenlayerValidator) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EigenlayerValidator) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EigenlayerValidator) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCommissionRate

`func (o *EigenlayerValidator) GetCommissionRate() float32`

GetCommissionRate returns the CommissionRate field if non-nil, zero value otherwise.

### GetCommissionRateOk

`func (o *EigenlayerValidator) GetCommissionRateOk() (*float32, bool)`

GetCommissionRateOk returns a tuple with the CommissionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionRate

`func (o *EigenlayerValidator) SetCommissionRate(v float32)`

SetCommissionRate sets CommissionRate field to given value.

### HasCommissionRate

`func (o *EigenlayerValidator) HasCommissionRate() bool`

HasCommissionRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


