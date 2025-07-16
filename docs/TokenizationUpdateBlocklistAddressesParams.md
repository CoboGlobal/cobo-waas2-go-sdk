# TokenizationUpdateBlocklistAddressesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**TokenizationUpdateAddressAction**](TokenizationUpdateAddressAction.md) |  | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Addresses** | [**[]TokenizationUpdateBlocklistAddressesParamsAddressesInner**](TokenizationUpdateBlocklistAddressesParamsAddressesInner.md) | A list of addresses to manage. For &#39;add&#39; operations, notes can be provided. For &#39;remove&#39; operations, notes are ignored. | 

## Methods

### NewTokenizationUpdateBlocklistAddressesParams

`func NewTokenizationUpdateBlocklistAddressesParams(action TokenizationUpdateAddressAction, source TokenizationTokenOperationSource, addresses []TokenizationUpdateBlocklistAddressesParamsAddressesInner, ) *TokenizationUpdateBlocklistAddressesParams`

NewTokenizationUpdateBlocklistAddressesParams instantiates a new TokenizationUpdateBlocklistAddressesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUpdateBlocklistAddressesParamsWithDefaults

`func NewTokenizationUpdateBlocklistAddressesParamsWithDefaults() *TokenizationUpdateBlocklistAddressesParams`

NewTokenizationUpdateBlocklistAddressesParamsWithDefaults instantiates a new TokenizationUpdateBlocklistAddressesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TokenizationUpdateBlocklistAddressesParams) GetAction() TokenizationUpdateAddressAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationUpdateBlocklistAddressesParams) GetActionOk() (*TokenizationUpdateAddressAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationUpdateBlocklistAddressesParams) SetAction(v TokenizationUpdateAddressAction)`

SetAction sets Action field to given value.


### GetSource

`func (o *TokenizationUpdateBlocklistAddressesParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationUpdateBlocklistAddressesParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationUpdateBlocklistAddressesParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetAddresses

`func (o *TokenizationUpdateBlocklistAddressesParams) GetAddresses() []TokenizationUpdateBlocklistAddressesParamsAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TokenizationUpdateBlocklistAddressesParams) GetAddressesOk() (*[]TokenizationUpdateBlocklistAddressesParamsAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TokenizationUpdateBlocklistAddressesParams) SetAddresses(v []TokenizationUpdateBlocklistAddressesParamsAddressesInner)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


