# PolicyCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The field name. | 
**ValueType** | [**PolicyFieldValueType**](PolicyFieldValueType.md) |  | 
**Value** | **string** | The field value. | 
**Operator** | [**PolicyFieldOperator**](PolicyFieldOperator.md) |  | 

## Methods

### NewPolicyCondition

`func NewPolicyCondition(field string, valueType PolicyFieldValueType, value string, operator PolicyFieldOperator, ) *PolicyCondition`

NewPolicyCondition instantiates a new PolicyCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConditionWithDefaults

`func NewPolicyConditionWithDefaults() *PolicyCondition`

NewPolicyConditionWithDefaults instantiates a new PolicyCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *PolicyCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *PolicyCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *PolicyCondition) SetField(v string)`

SetField sets Field field to given value.


### GetValueType

`func (o *PolicyCondition) GetValueType() PolicyFieldValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *PolicyCondition) GetValueTypeOk() (*PolicyFieldValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *PolicyCondition) SetValueType(v PolicyFieldValueType)`

SetValueType sets ValueType field to given value.


### GetValue

`func (o *PolicyCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PolicyCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PolicyCondition) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *PolicyCondition) GetOperator() PolicyFieldOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PolicyCondition) GetOperatorOk() (*PolicyFieldOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PolicyCondition) SetOperator(v PolicyFieldOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


