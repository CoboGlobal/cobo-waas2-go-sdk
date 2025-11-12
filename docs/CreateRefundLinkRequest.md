# CreateRefundLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessInfo** | [**RefundLinkBusinessInfo**](RefundLinkBusinessInfo.md) |  | 
**DisplayInfo** | Pointer to [**LinkDisplayInfo**](LinkDisplayInfo.md) |  | [optional] 

## Methods

### NewCreateRefundLinkRequest

`func NewCreateRefundLinkRequest(businessInfo RefundLinkBusinessInfo, ) *CreateRefundLinkRequest`

NewCreateRefundLinkRequest instantiates a new CreateRefundLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRefundLinkRequestWithDefaults

`func NewCreateRefundLinkRequestWithDefaults() *CreateRefundLinkRequest`

NewCreateRefundLinkRequestWithDefaults instantiates a new CreateRefundLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessInfo

`func (o *CreateRefundLinkRequest) GetBusinessInfo() RefundLinkBusinessInfo`

GetBusinessInfo returns the BusinessInfo field if non-nil, zero value otherwise.

### GetBusinessInfoOk

`func (o *CreateRefundLinkRequest) GetBusinessInfoOk() (*RefundLinkBusinessInfo, bool)`

GetBusinessInfoOk returns a tuple with the BusinessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessInfo

`func (o *CreateRefundLinkRequest) SetBusinessInfo(v RefundLinkBusinessInfo)`

SetBusinessInfo sets BusinessInfo field to given value.


### GetDisplayInfo

`func (o *CreateRefundLinkRequest) GetDisplayInfo() LinkDisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *CreateRefundLinkRequest) GetDisplayInfoOk() (*LinkDisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *CreateRefundLinkRequest) SetDisplayInfo(v LinkDisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *CreateRefundLinkRequest) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


