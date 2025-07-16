# TokenizationUpdateAllowlistAddressesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**TokenizationUpdateAddressAction**](TokenizationUpdateAddressAction.md) |  | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Addresses** | [**[]TokenizationUpdateAllowlistAddressesParamsAddressesInner**](TokenizationUpdateAllowlistAddressesParamsAddressesInner.md) | A list of addresses to manage. For &#39;add&#39; operations, notes can be provided. For &#39;remove&#39; operations, notes are ignored. | 

## Methods

### NewTokenizationUpdateAllowlistAddressesParams

`func NewTokenizationUpdateAllowlistAddressesParams(action TokenizationUpdateAddressAction, source TokenizationTokenOperationSource, addresses []TokenizationUpdateAllowlistAddressesParamsAddressesInner, ) *TokenizationUpdateAllowlistAddressesParams`

NewTokenizationUpdateAllowlistAddressesParams instantiates a new TokenizationUpdateAllowlistAddressesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUpdateAllowlistAddressesParamsWithDefaults

`func NewTokenizationUpdateAllowlistAddressesParamsWithDefaults() *TokenizationUpdateAllowlistAddressesParams`

NewTokenizationUpdateAllowlistAddressesParamsWithDefaults instantiates a new TokenizationUpdateAllowlistAddressesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TokenizationUpdateAllowlistAddressesParams) GetAction() TokenizationUpdateAddressAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationUpdateAllowlistAddressesParams) GetActionOk() (*TokenizationUpdateAddressAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationUpdateAllowlistAddressesParams) SetAction(v TokenizationUpdateAddressAction)`

SetAction sets Action field to given value.


### GetSource

`func (o *TokenizationUpdateAllowlistAddressesParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationUpdateAllowlistAddressesParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationUpdateAllowlistAddressesParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetAddresses

`func (o *TokenizationUpdateAllowlistAddressesParams) GetAddresses() []TokenizationUpdateAllowlistAddressesParamsAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TokenizationUpdateAllowlistAddressesParams) GetAddressesOk() (*[]TokenizationUpdateAllowlistAddressesParamsAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TokenizationUpdateAllowlistAddressesParams) SetAddresses(v []TokenizationUpdateAllowlistAddressesParamsAddressesInner)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


