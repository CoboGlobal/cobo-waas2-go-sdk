# SwapActivityApprovers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The approver name of the swap activity.  | 
**Status** | [**SwapApproversStatus**](SwapApproversStatus.md) |  | 

## Methods

### NewSwapActivityApprovers

`func NewSwapActivityApprovers(name string, status SwapApproversStatus, ) *SwapActivityApprovers`

NewSwapActivityApprovers instantiates a new SwapActivityApprovers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapActivityApproversWithDefaults

`func NewSwapActivityApproversWithDefaults() *SwapActivityApprovers`

NewSwapActivityApproversWithDefaults instantiates a new SwapActivityApprovers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SwapActivityApprovers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwapActivityApprovers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwapActivityApprovers) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *SwapActivityApprovers) GetStatus() SwapApproversStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwapActivityApprovers) GetStatusOk() (*SwapApproversStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwapActivityApprovers) SetStatus(v SwapApproversStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


