# TokenizationUpdateAllowlistAddressesEstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**TokenizationUpdateAddressAction**](TokenizationUpdateAddressAction.md) |  | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Addresses** | [**[]TokenizationUpdateAllowlistAddressesParamsAddressesInner**](TokenizationUpdateAllowlistAddressesParamsAddressesInner.md) | A list of addresses to manage. For &#39;add&#39; operations, notes can be provided. For &#39;remove&#39; operations, notes are ignored. | 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**TokenId** | **string** | The ID of the token. | 

## Methods

### NewTokenizationUpdateAllowlistAddressesEstimateFeeParams

`func NewTokenizationUpdateAllowlistAddressesEstimateFeeParams(action TokenizationUpdateAddressAction, source TokenizationTokenOperationSource, addresses []TokenizationUpdateAllowlistAddressesParamsAddressesInner, operationType TokenizationOperationType, tokenId string, ) *TokenizationUpdateAllowlistAddressesEstimateFeeParams`

NewTokenizationUpdateAllowlistAddressesEstimateFeeParams instantiates a new TokenizationUpdateAllowlistAddressesEstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUpdateAllowlistAddressesEstimateFeeParamsWithDefaults

`func NewTokenizationUpdateAllowlistAddressesEstimateFeeParamsWithDefaults() *TokenizationUpdateAllowlistAddressesEstimateFeeParams`

NewTokenizationUpdateAllowlistAddressesEstimateFeeParamsWithDefaults instantiates a new TokenizationUpdateAllowlistAddressesEstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetAction() TokenizationUpdateAddressAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetActionOk() (*TokenizationUpdateAddressAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) SetAction(v TokenizationUpdateAddressAction)`

SetAction sets Action field to given value.


### GetSource

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetAddresses

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetAddresses() []TokenizationUpdateAllowlistAddressesParamsAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetAddressesOk() (*[]TokenizationUpdateAllowlistAddressesParamsAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) SetAddresses(v []TokenizationUpdateAllowlistAddressesParamsAddressesInner)`

SetAddresses sets Addresses field to given value.


### GetOperationType

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.


### GetTokenId

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationUpdateAllowlistAddressesEstimateFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


