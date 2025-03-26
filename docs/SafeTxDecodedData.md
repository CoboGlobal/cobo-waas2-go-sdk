# SafeTxDecodedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Name of the decoded method | [optional] 
**Parameters** | Pointer to [**[]SafeTxDecodedDataParameters**](SafeTxDecodedDataParameters.md) | List of method parameters | [optional] 

## Methods

### NewSafeTxDecodedData

`func NewSafeTxDecodedData() *SafeTxDecodedData`

NewSafeTxDecodedData instantiates a new SafeTxDecodedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeTxDecodedDataWithDefaults

`func NewSafeTxDecodedDataWithDefaults() *SafeTxDecodedData`

NewSafeTxDecodedDataWithDefaults instantiates a new SafeTxDecodedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *SafeTxDecodedData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SafeTxDecodedData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SafeTxDecodedData) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SafeTxDecodedData) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetParameters

`func (o *SafeTxDecodedData) GetParameters() []SafeTxDecodedDataParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SafeTxDecodedData) GetParametersOk() (*[]SafeTxDecodedDataParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SafeTxDecodedData) SetParameters(v []SafeTxDecodedDataParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *SafeTxDecodedData) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


