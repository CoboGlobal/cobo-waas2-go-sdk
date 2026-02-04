# CreateTopUpAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | The merchant ID. If not provided, the default merchant created during organization initialization will be used. | [optional] 
**TokenId** | **string** | The token ID, which identifies the cryptocurrency.  | 
**CustomPayerIds** | **[]string** | A list of unique custom payer IDs required to create top-up addresses.  | 

## Methods

### NewCreateTopUpAddresses

`func NewCreateTopUpAddresses(tokenId string, customPayerIds []string, ) *CreateTopUpAddresses`

NewCreateTopUpAddresses instantiates a new CreateTopUpAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTopUpAddressesWithDefaults

`func NewCreateTopUpAddressesWithDefaults() *CreateTopUpAddresses`

NewCreateTopUpAddressesWithDefaults instantiates a new CreateTopUpAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *CreateTopUpAddresses) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateTopUpAddresses) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateTopUpAddresses) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *CreateTopUpAddresses) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *CreateTopUpAddresses) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreateTopUpAddresses) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreateTopUpAddresses) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetCustomPayerIds

`func (o *CreateTopUpAddresses) GetCustomPayerIds() []string`

GetCustomPayerIds returns the CustomPayerIds field if non-nil, zero value otherwise.

### GetCustomPayerIdsOk

`func (o *CreateTopUpAddresses) GetCustomPayerIdsOk() (*[]string, bool)`

GetCustomPayerIdsOk returns a tuple with the CustomPayerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayerIds

`func (o *CreateTopUpAddresses) SetCustomPayerIds(v []string)`

SetCustomPayerIds sets CustomPayerIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


