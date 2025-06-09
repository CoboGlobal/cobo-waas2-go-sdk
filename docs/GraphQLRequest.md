# GraphQLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The GraphQL query or mutation as a string. | 
**Variables** | Pointer to **map[string]interface{}** | Variables to use in the GraphQL operation. Can include dynamic values required for the query or mutation. | [optional] 
**OperationName** | Pointer to **string** | The operation name in case of multiple operations in the same query or mutation. | [optional] 

## Methods

### NewGraphQLRequest

`func NewGraphQLRequest(query string, ) *GraphQLRequest`

NewGraphQLRequest instantiates a new GraphQLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLRequestWithDefaults

`func NewGraphQLRequestWithDefaults() *GraphQLRequest`

NewGraphQLRequestWithDefaults instantiates a new GraphQLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *GraphQLRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphQLRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphQLRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetVariables

`func (o *GraphQLRequest) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *GraphQLRequest) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *GraphQLRequest) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *GraphQLRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetOperationName

`func (o *GraphQLRequest) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *GraphQLRequest) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *GraphQLRequest) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *GraphQLRequest) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


