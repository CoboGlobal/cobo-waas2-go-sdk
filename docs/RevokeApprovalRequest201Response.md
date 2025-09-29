# RevokeApprovalRequest201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalId** | **string** |  | 
**Status** | Pointer to [**ApprovalStatus**](ApprovalStatus.md) |  | [optional] 

## Methods

### NewRevokeApprovalRequest201Response

`func NewRevokeApprovalRequest201Response(approvalId string, ) *RevokeApprovalRequest201Response`

NewRevokeApprovalRequest201Response instantiates a new RevokeApprovalRequest201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeApprovalRequest201ResponseWithDefaults

`func NewRevokeApprovalRequest201ResponseWithDefaults() *RevokeApprovalRequest201Response`

NewRevokeApprovalRequest201ResponseWithDefaults instantiates a new RevokeApprovalRequest201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalId

`func (o *RevokeApprovalRequest201Response) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *RevokeApprovalRequest201Response) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *RevokeApprovalRequest201Response) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.


### GetStatus

`func (o *RevokeApprovalRequest201Response) GetStatus() ApprovalStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RevokeApprovalRequest201Response) GetStatusOk() (*ApprovalStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RevokeApprovalRequest201Response) SetStatus(v ApprovalStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RevokeApprovalRequest201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


