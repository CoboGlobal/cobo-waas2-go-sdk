# AppWorkflowField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The workflow field name. | 
**ValueType** | [**PolicyFieldValueType**](PolicyFieldValueType.md) |  | 
**Value** | **string** | The workflow field value. | 

## Methods

### NewAppWorkflowField

`func NewAppWorkflowField(field string, valueType PolicyFieldValueType, value string, ) *AppWorkflowField`

NewAppWorkflowField instantiates a new AppWorkflowField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWorkflowFieldWithDefaults

`func NewAppWorkflowFieldWithDefaults() *AppWorkflowField`

NewAppWorkflowFieldWithDefaults instantiates a new AppWorkflowField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AppWorkflowField) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AppWorkflowField) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AppWorkflowField) SetField(v string)`

SetField sets Field field to given value.


### GetValueType

`func (o *AppWorkflowField) GetValueType() PolicyFieldValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *AppWorkflowField) GetValueTypeOk() (*PolicyFieldValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *AppWorkflowField) SetValueType(v PolicyFieldValueType)`

SetValueType sets ValueType field to given value.


### GetValue

`func (o *AppWorkflowField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppWorkflowField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppWorkflowField) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


