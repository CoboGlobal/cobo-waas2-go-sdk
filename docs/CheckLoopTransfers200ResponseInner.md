# CheckLoopTransfers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The wallet address. | [optional] 
**IsLoop** | Pointer to **bool** | Whether the transaction from the given source wallet to the destination address can be executed as a Loop transfer.  - &#x60;true&#x60;: The transaction can be executed as a Loop transfer. - &#x60;false&#x60;: The transaction cannot be executed as a Loop transfer.  | [optional] 

## Methods

### NewCheckLoopTransfers200ResponseInner

`func NewCheckLoopTransfers200ResponseInner() *CheckLoopTransfers200ResponseInner`

NewCheckLoopTransfers200ResponseInner instantiates a new CheckLoopTransfers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckLoopTransfers200ResponseInnerWithDefaults

`func NewCheckLoopTransfers200ResponseInnerWithDefaults() *CheckLoopTransfers200ResponseInner`

NewCheckLoopTransfers200ResponseInnerWithDefaults instantiates a new CheckLoopTransfers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CheckLoopTransfers200ResponseInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CheckLoopTransfers200ResponseInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CheckLoopTransfers200ResponseInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CheckLoopTransfers200ResponseInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIsLoop

`func (o *CheckLoopTransfers200ResponseInner) GetIsLoop() bool`

GetIsLoop returns the IsLoop field if non-nil, zero value otherwise.

### GetIsLoopOk

`func (o *CheckLoopTransfers200ResponseInner) GetIsLoopOk() (*bool, bool)`

GetIsLoopOk returns a tuple with the IsLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoop

`func (o *CheckLoopTransfers200ResponseInner) SetIsLoop(v bool)`

SetIsLoop sets IsLoop field to given value.

### HasIsLoop

`func (o *CheckLoopTransfers200ResponseInner) HasIsLoop() bool`

HasIsLoop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


