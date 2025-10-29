# CreateOrderLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessInfo** | [**OrderLinkBusinessInfo**](OrderLinkBusinessInfo.md) |  | 
**DisplayInfo** | Pointer to [**LinkDisplayInfo**](LinkDisplayInfo.md) |  | [optional] 

## Methods

### NewCreateOrderLinkRequest

`func NewCreateOrderLinkRequest(businessInfo OrderLinkBusinessInfo, ) *CreateOrderLinkRequest`

NewCreateOrderLinkRequest instantiates a new CreateOrderLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderLinkRequestWithDefaults

`func NewCreateOrderLinkRequestWithDefaults() *CreateOrderLinkRequest`

NewCreateOrderLinkRequestWithDefaults instantiates a new CreateOrderLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessInfo

`func (o *CreateOrderLinkRequest) GetBusinessInfo() OrderLinkBusinessInfo`

GetBusinessInfo returns the BusinessInfo field if non-nil, zero value otherwise.

### GetBusinessInfoOk

`func (o *CreateOrderLinkRequest) GetBusinessInfoOk() (*OrderLinkBusinessInfo, bool)`

GetBusinessInfoOk returns a tuple with the BusinessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessInfo

`func (o *CreateOrderLinkRequest) SetBusinessInfo(v OrderLinkBusinessInfo)`

SetBusinessInfo sets BusinessInfo field to given value.


### GetDisplayInfo

`func (o *CreateOrderLinkRequest) GetDisplayInfo() LinkDisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *CreateOrderLinkRequest) GetDisplayInfoOk() (*LinkDisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *CreateOrderLinkRequest) SetDisplayInfo(v LinkDisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *CreateOrderLinkRequest) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


