# TokenizationUpdateBlocklistAddressesEstimateFeeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**TokenizationUpdateAddressAction**](TokenizationUpdateAddressAction.md) |  | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Addresses** | [**[]TokenizationUpdateBlocklistAddressesParamsAddressesInner**](TokenizationUpdateBlocklistAddressesParamsAddressesInner.md) | A list of addresses to manage. For &#39;add&#39; operations, notes can be provided. For &#39;remove&#39; operations, notes are ignored. | 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**TokenId** | **string** | The ID of the token. | 

## Methods

### NewTokenizationUpdateBlocklistAddressesEstimateFeeParams

`func NewTokenizationUpdateBlocklistAddressesEstimateFeeParams(action TokenizationUpdateAddressAction, source TokenizationTokenOperationSource, addresses []TokenizationUpdateBlocklistAddressesParamsAddressesInner, operationType TokenizationOperationType, tokenId string, ) *TokenizationUpdateBlocklistAddressesEstimateFeeParams`

NewTokenizationUpdateBlocklistAddressesEstimateFeeParams instantiates a new TokenizationUpdateBlocklistAddressesEstimateFeeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUpdateBlocklistAddressesEstimateFeeParamsWithDefaults

`func NewTokenizationUpdateBlocklistAddressesEstimateFeeParamsWithDefaults() *TokenizationUpdateBlocklistAddressesEstimateFeeParams`

NewTokenizationUpdateBlocklistAddressesEstimateFeeParamsWithDefaults instantiates a new TokenizationUpdateBlocklistAddressesEstimateFeeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetAction() TokenizationUpdateAddressAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetActionOk() (*TokenizationUpdateAddressAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) SetAction(v TokenizationUpdateAddressAction)`

SetAction sets Action field to given value.


### GetSource

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetAddresses

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetAddresses() []TokenizationUpdateBlocklistAddressesParamsAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetAddressesOk() (*[]TokenizationUpdateBlocklistAddressesParamsAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) SetAddresses(v []TokenizationUpdateBlocklistAddressesParamsAddressesInner)`

SetAddresses sets Addresses field to given value.


### GetOperationType

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.


### GetTokenId

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationUpdateBlocklistAddressesEstimateFeeParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


