# GraphQLErrorLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to **int32** | The line number in the query where the error occurred. | [optional] 
**Column** | Pointer to **int32** | The column number in the query where the error occurred. | [optional] 

## Methods

### NewGraphQLErrorLocationsInner

`func NewGraphQLErrorLocationsInner() *GraphQLErrorLocationsInner`

NewGraphQLErrorLocationsInner instantiates a new GraphQLErrorLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLErrorLocationsInnerWithDefaults

`func NewGraphQLErrorLocationsInnerWithDefaults() *GraphQLErrorLocationsInner`

NewGraphQLErrorLocationsInnerWithDefaults instantiates a new GraphQLErrorLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *GraphQLErrorLocationsInner) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *GraphQLErrorLocationsInner) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *GraphQLErrorLocationsInner) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *GraphQLErrorLocationsInner) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetColumn

`func (o *GraphQLErrorLocationsInner) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *GraphQLErrorLocationsInner) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *GraphQLErrorLocationsInner) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *GraphQLErrorLocationsInner) HasColumn() bool`

HasColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


