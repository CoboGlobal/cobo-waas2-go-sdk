# TokenizationUnarchiveTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInitiator** | Pointer to **string** | The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 

## Methods

### NewTokenizationUnarchiveTokenRequest

`func NewTokenizationUnarchiveTokenRequest() *TokenizationUnarchiveTokenRequest`

NewTokenizationUnarchiveTokenRequest instantiates a new TokenizationUnarchiveTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUnarchiveTokenRequestWithDefaults

`func NewTokenizationUnarchiveTokenRequestWithDefaults() *TokenizationUnarchiveTokenRequest`

NewTokenizationUnarchiveTokenRequestWithDefaults instantiates a new TokenizationUnarchiveTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInitiator

`func (o *TokenizationUnarchiveTokenRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationUnarchiveTokenRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationUnarchiveTokenRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationUnarchiveTokenRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


