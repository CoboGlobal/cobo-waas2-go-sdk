# RootPubkey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to **string** | The vault&#39;s [root extended public key](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/tss-node-deployment#tss-node-on-cobo-portal-and-mpc-root-extended-public-key). | [optional] 
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

### GetPubkey

`func (o *RootPubkey) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *RootPubkey) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *RootPubkey) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *RootPubkey) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

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


