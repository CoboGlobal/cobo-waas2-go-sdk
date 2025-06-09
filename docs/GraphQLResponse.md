# GraphQLResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | The response data for the GraphQL operation. | [optional] 
**Errors** | Pointer to [**[]GraphQLError**](GraphQLError.md) | Any errors that occurred during the GraphQL operation. | [optional] 

## Methods

### NewGraphQLResponse

`func NewGraphQLResponse() *GraphQLResponse`

NewGraphQLResponse instantiates a new GraphQLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLResponseWithDefaults

`func NewGraphQLResponseWithDefaults() *GraphQLResponse`

NewGraphQLResponseWithDefaults instantiates a new GraphQLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GraphQLResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GraphQLResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GraphQLResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GraphQLResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *GraphQLResponse) GetErrors() []GraphQLError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GraphQLResponse) GetErrorsOk() (*[]GraphQLError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GraphQLResponse) SetErrors(v []GraphQLError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GraphQLResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


