# TokenizationEvmContractCallParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TokenizationContractCallType**](TokenizationContractCallType.md) |  | [optional] 
**Calldata** | **string** | The data used to invoke a specific function or method within the specified contract at the destination address, with a maximum length of 65,000 characters.  | 
**Value** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 

## Methods

### NewTokenizationEvmContractCallParams

`func NewTokenizationEvmContractCallParams(calldata string, ) *TokenizationEvmContractCallParams`

NewTokenizationEvmContractCallParams instantiates a new TokenizationEvmContractCallParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationEvmContractCallParamsWithDefaults

`func NewTokenizationEvmContractCallParamsWithDefaults() *TokenizationEvmContractCallParams`

NewTokenizationEvmContractCallParamsWithDefaults instantiates a new TokenizationEvmContractCallParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TokenizationEvmContractCallParams) GetType() TokenizationContractCallType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizationEvmContractCallParams) GetTypeOk() (*TokenizationContractCallType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizationEvmContractCallParams) SetType(v TokenizationContractCallType)`

SetType sets Type field to given value.

### HasType

`func (o *TokenizationEvmContractCallParams) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCalldata

`func (o *TokenizationEvmContractCallParams) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *TokenizationEvmContractCallParams) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *TokenizationEvmContractCallParams) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.


### GetValue

`func (o *TokenizationEvmContractCallParams) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TokenizationEvmContractCallParams) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TokenizationEvmContractCallParams) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TokenizationEvmContractCallParams) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


