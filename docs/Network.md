# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The network name. | 
**IconUrl** | Pointer to **string** | URL of the icon image. | [optional] 

## Methods

### NewNetwork

`func NewNetwork(name string, ) *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Network) SetName(v string)`

SetName sets Name field to given value.


### GetIconUrl

`func (o *Network) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Network) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Network) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Network) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


