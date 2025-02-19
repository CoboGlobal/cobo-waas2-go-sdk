# QueryApprovalStatement200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | approval statement id. | [optional] 
**UserId** | Pointer to **string** | approver user id. | [optional] 
**Pubkey** | Pointer to **string** | approver user pubkey. | [optional] 
**Status** | Pointer to [**ApprovalStatementStatus**](ApprovalStatementStatus.md) |  | [optional] 

## Methods

### NewQueryApprovalStatement200Response

`func NewQueryApprovalStatement200Response() *QueryApprovalStatement200Response`

NewQueryApprovalStatement200Response instantiates a new QueryApprovalStatement200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryApprovalStatement200ResponseWithDefaults

`func NewQueryApprovalStatement200ResponseWithDefaults() *QueryApprovalStatement200Response`

NewQueryApprovalStatement200ResponseWithDefaults instantiates a new QueryApprovalStatement200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueryApprovalStatement200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryApprovalStatement200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryApprovalStatement200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueryApprovalStatement200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *QueryApprovalStatement200Response) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *QueryApprovalStatement200Response) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *QueryApprovalStatement200Response) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *QueryApprovalStatement200Response) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPubkey

`func (o *QueryApprovalStatement200Response) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *QueryApprovalStatement200Response) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *QueryApprovalStatement200Response) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *QueryApprovalStatement200Response) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetStatus

`func (o *QueryApprovalStatement200Response) GetStatus() ApprovalStatementStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryApprovalStatement200Response) GetStatusOk() (*ApprovalStatementStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryApprovalStatement200Response) SetStatus(v ApprovalStatementStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryApprovalStatement200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


