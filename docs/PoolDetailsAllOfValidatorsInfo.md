# PoolDetailsAllOfValidatorsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IconUrl** | **string** | The URL of the validator&#39;s icon. | 
**Name** | **string** | The name of validator. | 
**Priority** | Pointer to **int32** | The priority of validator. | [optional] 
**PublicKey** | **string** | The public key of validator. | 
**CommissionRate** | **float32** | The commission of validator. | 
**SupportedPosChains** | **[]string** | The list of supported pos chains. | 
**Address** | **string** | The address of validator. | 

## Methods

### NewPoolDetailsAllOfValidatorsInfo

`func NewPoolDetailsAllOfValidatorsInfo(iconUrl string, name string, publicKey string, commissionRate float32, supportedPosChains []string, address string, ) *PoolDetailsAllOfValidatorsInfo`

NewPoolDetailsAllOfValidatorsInfo instantiates a new PoolDetailsAllOfValidatorsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDetailsAllOfValidatorsInfoWithDefaults

`func NewPoolDetailsAllOfValidatorsInfoWithDefaults() *PoolDetailsAllOfValidatorsInfo`

NewPoolDetailsAllOfValidatorsInfoWithDefaults instantiates a new PoolDetailsAllOfValidatorsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIconUrl

`func (o *PoolDetailsAllOfValidatorsInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *PoolDetailsAllOfValidatorsInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *PoolDetailsAllOfValidatorsInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetName

`func (o *PoolDetailsAllOfValidatorsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolDetailsAllOfValidatorsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolDetailsAllOfValidatorsInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *PoolDetailsAllOfValidatorsInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PoolDetailsAllOfValidatorsInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PoolDetailsAllOfValidatorsInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PoolDetailsAllOfValidatorsInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublicKey

`func (o *PoolDetailsAllOfValidatorsInfo) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PoolDetailsAllOfValidatorsInfo) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PoolDetailsAllOfValidatorsInfo) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetCommissionRate

`func (o *PoolDetailsAllOfValidatorsInfo) GetCommissionRate() float32`

GetCommissionRate returns the CommissionRate field if non-nil, zero value otherwise.

### GetCommissionRateOk

`func (o *PoolDetailsAllOfValidatorsInfo) GetCommissionRateOk() (*float32, bool)`

GetCommissionRateOk returns a tuple with the CommissionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionRate

`func (o *PoolDetailsAllOfValidatorsInfo) SetCommissionRate(v float32)`

SetCommissionRate sets CommissionRate field to given value.


### GetSupportedPosChains

`func (o *PoolDetailsAllOfValidatorsInfo) GetSupportedPosChains() []string`

GetSupportedPosChains returns the SupportedPosChains field if non-nil, zero value otherwise.

### GetSupportedPosChainsOk

`func (o *PoolDetailsAllOfValidatorsInfo) GetSupportedPosChainsOk() (*[]string, bool)`

GetSupportedPosChainsOk returns a tuple with the SupportedPosChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPosChains

`func (o *PoolDetailsAllOfValidatorsInfo) SetSupportedPosChains(v []string)`

SetSupportedPosChains sets SupportedPosChains field to given value.


### GetAddress

`func (o *PoolDetailsAllOfValidatorsInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PoolDetailsAllOfValidatorsInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PoolDetailsAllOfValidatorsInfo) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


