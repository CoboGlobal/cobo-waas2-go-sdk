# UpdateKeyShareHolderGroupByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateKeyShareHolderGroupAction** | [**UpdateGroupAction**](UpdateGroupAction.md) |  | 
**OriginalMainGroupHandling** | Pointer to [**OriginalMainGroupHandling**](OriginalMainGroupHandling.md) |  | [optional] [default to ORIGINALMAINGROUPHANDLING_INVALIDATE]

## Methods

### NewUpdateKeyShareHolderGroupByIdRequest

`func NewUpdateKeyShareHolderGroupByIdRequest(updateKeyShareHolderGroupAction UpdateGroupAction, ) *UpdateKeyShareHolderGroupByIdRequest`

NewUpdateKeyShareHolderGroupByIdRequest instantiates a new UpdateKeyShareHolderGroupByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKeyShareHolderGroupByIdRequestWithDefaults

`func NewUpdateKeyShareHolderGroupByIdRequestWithDefaults() *UpdateKeyShareHolderGroupByIdRequest`

NewUpdateKeyShareHolderGroupByIdRequestWithDefaults instantiates a new UpdateKeyShareHolderGroupByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateKeyShareHolderGroupAction

`func (o *UpdateKeyShareHolderGroupByIdRequest) GetUpdateKeyShareHolderGroupAction() UpdateGroupAction`

GetUpdateKeyShareHolderGroupAction returns the UpdateKeyShareHolderGroupAction field if non-nil, zero value otherwise.

### GetUpdateKeyShareHolderGroupActionOk

`func (o *UpdateKeyShareHolderGroupByIdRequest) GetUpdateKeyShareHolderGroupActionOk() (*UpdateGroupAction, bool)`

GetUpdateKeyShareHolderGroupActionOk returns a tuple with the UpdateKeyShareHolderGroupAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateKeyShareHolderGroupAction

`func (o *UpdateKeyShareHolderGroupByIdRequest) SetUpdateKeyShareHolderGroupAction(v UpdateGroupAction)`

SetUpdateKeyShareHolderGroupAction sets UpdateKeyShareHolderGroupAction field to given value.


### GetOriginalMainGroupHandling

`func (o *UpdateKeyShareHolderGroupByIdRequest) GetOriginalMainGroupHandling() OriginalMainGroupHandling`

GetOriginalMainGroupHandling returns the OriginalMainGroupHandling field if non-nil, zero value otherwise.

### GetOriginalMainGroupHandlingOk

`func (o *UpdateKeyShareHolderGroupByIdRequest) GetOriginalMainGroupHandlingOk() (*OriginalMainGroupHandling, bool)`

GetOriginalMainGroupHandlingOk returns a tuple with the OriginalMainGroupHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMainGroupHandling

`func (o *UpdateKeyShareHolderGroupByIdRequest) SetOriginalMainGroupHandling(v OriginalMainGroupHandling)`

SetOriginalMainGroupHandling sets OriginalMainGroupHandling field to given value.

### HasOriginalMainGroupHandling

`func (o *UpdateKeyShareHolderGroupByIdRequest) HasOriginalMainGroupHandling() bool`

HasOriginalMainGroupHandling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


