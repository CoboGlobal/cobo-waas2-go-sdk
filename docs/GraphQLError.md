# GraphQLError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The error message. | [optional] 
**Locations** | Pointer to [**[]GraphQLErrorLocationsInner**](GraphQLErrorLocationsInner.md) | The locations in the query where the error occurred. | [optional] 
**Path** | Pointer to **[]string** | The path in the response where the error occurred. | [optional] 

## Methods

### NewGraphQLError

`func NewGraphQLError() *GraphQLError`

NewGraphQLError instantiates a new GraphQLError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLErrorWithDefaults

`func NewGraphQLErrorWithDefaults() *GraphQLError`

NewGraphQLErrorWithDefaults instantiates a new GraphQLError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GraphQLError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GraphQLError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GraphQLError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GraphQLError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLocations

`func (o *GraphQLError) GetLocations() []GraphQLErrorLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GraphQLError) GetLocationsOk() (*[]GraphQLErrorLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GraphQLError) SetLocations(v []GraphQLErrorLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *GraphQLError) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetPath

`func (o *GraphQLError) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GraphQLError) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GraphQLError) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GraphQLError) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


