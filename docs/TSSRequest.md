# TSSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The TSS request ID. | [optional] 
**Type** | Pointer to [**TSSRequestType**](TSSRequestType.md) |  | [optional] 
**Status** | Pointer to [**TSSRequestStatus**](TSSRequestStatus.md) |  | [optional] 

## Methods

### NewTSSRequest

`func NewTSSRequest() *TSSRequest`

NewTSSRequest instantiates a new TSSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSRequestWithDefaults

`func NewTSSRequestWithDefaults() *TSSRequest`

NewTSSRequestWithDefaults instantiates a new TSSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TSSRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TSSRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TSSRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TSSRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TSSRequest) GetType() TSSRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TSSRequest) GetTypeOk() (*TSSRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TSSRequest) SetType(v TSSRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *TSSRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *TSSRequest) GetStatus() TSSRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TSSRequest) GetStatusOk() (*TSSRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TSSRequest) SetStatus(v TSSRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TSSRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


