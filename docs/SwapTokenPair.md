# SwapTokenPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayTokenId** | Pointer to **string** | The source token symbol. | [optional] 
**ReceiveTokenId** | Pointer to **string** | The target token symbol. | [optional] 

## Methods

### NewSwapTokenPair

`func NewSwapTokenPair() *SwapTokenPair`

NewSwapTokenPair instantiates a new SwapTokenPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapTokenPairWithDefaults

`func NewSwapTokenPairWithDefaults() *SwapTokenPair`

NewSwapTokenPairWithDefaults instantiates a new SwapTokenPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayTokenId

`func (o *SwapTokenPair) GetPayTokenId() string`

GetPayTokenId returns the PayTokenId field if non-nil, zero value otherwise.

### GetPayTokenIdOk

`func (o *SwapTokenPair) GetPayTokenIdOk() (*string, bool)`

GetPayTokenIdOk returns a tuple with the PayTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTokenId

`func (o *SwapTokenPair) SetPayTokenId(v string)`

SetPayTokenId sets PayTokenId field to given value.

### HasPayTokenId

`func (o *SwapTokenPair) HasPayTokenId() bool`

HasPayTokenId returns a boolean if a field has been set.

### GetReceiveTokenId

`func (o *SwapTokenPair) GetReceiveTokenId() string`

GetReceiveTokenId returns the ReceiveTokenId field if non-nil, zero value otherwise.

### GetReceiveTokenIdOk

`func (o *SwapTokenPair) GetReceiveTokenIdOk() (*string, bool)`

GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveTokenId

`func (o *SwapTokenPair) SetReceiveTokenId(v string)`

SetReceiveTokenId sets ReceiveTokenId field to given value.

### HasReceiveTokenId

`func (o *SwapTokenPair) HasReceiveTokenId() bool`

HasReceiveTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


