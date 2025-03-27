# SafeTxDecodedDataParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the parameter. | [optional] 
**Type** | Pointer to **string** | The data type of the parameter. | [optional] 
**Value** | Pointer to **string** | The value of the parameter. | [optional] 
**ValueDecoded** | Pointer to [**[]SafeTxSubTransaction**](SafeTxSubTransaction.md) | The decoded value of the parameter (if applicable). | [optional] 

## Methods

### NewSafeTxDecodedDataParameters

`func NewSafeTxDecodedDataParameters() *SafeTxDecodedDataParameters`

NewSafeTxDecodedDataParameters instantiates a new SafeTxDecodedDataParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeTxDecodedDataParametersWithDefaults

`func NewSafeTxDecodedDataParametersWithDefaults() *SafeTxDecodedDataParameters`

NewSafeTxDecodedDataParametersWithDefaults instantiates a new SafeTxDecodedDataParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SafeTxDecodedDataParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SafeTxDecodedDataParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SafeTxDecodedDataParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SafeTxDecodedDataParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SafeTxDecodedDataParameters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SafeTxDecodedDataParameters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SafeTxDecodedDataParameters) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SafeTxDecodedDataParameters) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *SafeTxDecodedDataParameters) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SafeTxDecodedDataParameters) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SafeTxDecodedDataParameters) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SafeTxDecodedDataParameters) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueDecoded

`func (o *SafeTxDecodedDataParameters) GetValueDecoded() []SafeTxSubTransaction`

GetValueDecoded returns the ValueDecoded field if non-nil, zero value otherwise.

### GetValueDecodedOk

`func (o *SafeTxDecodedDataParameters) GetValueDecodedOk() (*[]SafeTxSubTransaction, bool)`

GetValueDecodedOk returns a tuple with the ValueDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDecoded

`func (o *SafeTxDecodedDataParameters) SetValueDecoded(v []SafeTxSubTransaction)`

SetValueDecoded sets ValueDecoded field to given value.

### HasValueDecoded

`func (o *SafeTxDecodedDataParameters) HasValueDecoded() bool`

HasValueDecoded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


