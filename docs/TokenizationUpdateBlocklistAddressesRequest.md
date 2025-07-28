# TokenizationUpdateBlocklistAddressesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**TokenizationUpdateAddressAction**](TokenizationUpdateAddressAction.md) |  | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Addresses** | [**[]TokenizationUpdateBlocklistAddressesParamsAddressesInner**](TokenizationUpdateBlocklistAddressesParamsAddressesInner.md) | A list of addresses to manage. For &#39;add&#39; operations, notes can be provided. For &#39;remove&#39; operations, notes are ignored. | 
**AppInitiator** | Pointer to **string** | The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | [optional] 

## Methods

### NewTokenizationUpdateBlocklistAddressesRequest

`func NewTokenizationUpdateBlocklistAddressesRequest(action TokenizationUpdateAddressAction, source TokenizationTokenOperationSource, addresses []TokenizationUpdateBlocklistAddressesParamsAddressesInner, fee TransactionRequestFee, ) *TokenizationUpdateBlocklistAddressesRequest`

NewTokenizationUpdateBlocklistAddressesRequest instantiates a new TokenizationUpdateBlocklistAddressesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUpdateBlocklistAddressesRequestWithDefaults

`func NewTokenizationUpdateBlocklistAddressesRequestWithDefaults() *TokenizationUpdateBlocklistAddressesRequest`

NewTokenizationUpdateBlocklistAddressesRequestWithDefaults instantiates a new TokenizationUpdateBlocklistAddressesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetAction() TokenizationUpdateAddressAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetActionOk() (*TokenizationUpdateAddressAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationUpdateBlocklistAddressesRequest) SetAction(v TokenizationUpdateAddressAction)`

SetAction sets Action field to given value.


### GetSource

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationUpdateBlocklistAddressesRequest) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetAddresses

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetAddresses() []TokenizationUpdateBlocklistAddressesParamsAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetAddressesOk() (*[]TokenizationUpdateBlocklistAddressesParamsAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TokenizationUpdateBlocklistAddressesRequest) SetAddresses(v []TokenizationUpdateBlocklistAddressesParamsAddressesInner)`

SetAddresses sets Addresses field to given value.


### GetAppInitiator

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationUpdateBlocklistAddressesRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationUpdateBlocklistAddressesRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.

### GetFee

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TokenizationUpdateBlocklistAddressesRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetRequestId

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TokenizationUpdateBlocklistAddressesRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TokenizationUpdateBlocklistAddressesRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TokenizationUpdateBlocklistAddressesRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


