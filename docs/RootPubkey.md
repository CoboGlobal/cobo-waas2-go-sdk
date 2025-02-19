# RootPubkey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootPubkey** | Pointer to **string** | The vault&#39;s [root extended public key](https://www.cobo.com/developers/v2/guides/mpc-wallets/get-started-ocw#root-extended-public-keys). | [optional] 
**Curve** | Pointer to [**CurveType**](CurveType.md) |  | [optional] 

## Methods

### NewRootPubkey

`func NewRootPubkey() *RootPubkey`

NewRootPubkey instantiates a new RootPubkey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootPubkeyWithDefaults

`func NewRootPubkeyWithDefaults() *RootPubkey`

NewRootPubkeyWithDefaults instantiates a new RootPubkey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootPubkey

`func (o *RootPubkey) GetRootPubkey() string`

GetRootPubkey returns the RootPubkey field if non-nil, zero value otherwise.

### GetRootPubkeyOk

`func (o *RootPubkey) GetRootPubkeyOk() (*string, bool)`

GetRootPubkeyOk returns a tuple with the RootPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPubkey

`func (o *RootPubkey) SetRootPubkey(v string)`

SetRootPubkey sets RootPubkey field to given value.

### HasRootPubkey

`func (o *RootPubkey) HasRootPubkey() bool`

HasRootPubkey returns a boolean if a field has been set.

### GetCurve

`func (o *RootPubkey) GetCurve() CurveType`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *RootPubkey) GetCurveOk() (*CurveType, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *RootPubkey) SetCurve(v CurveType)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *RootPubkey) HasCurve() bool`

HasCurve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


