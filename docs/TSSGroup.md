# TSSGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The TSS key share group ID. | [optional] 
**CanonicalGroupId** | Pointer to **string** | The canonical group ID. | [optional] 
**ProtocolGroupId** | Pointer to **string** | The protocol group ID. | [optional] 
**ProtocolType** | Pointer to **string** | The protocol type. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The group creation timestamp, in Unix timestamp format, measured in milliseconds. | [optional] 
**Type** | Pointer to [**TSSGroupType**](TSSGroupType.md) |  | [optional] 
**RootExtendedPublicKey** | Pointer to **string** | The root extended public key. | [optional] 
**Chaincode** | Pointer to **string** | The chaincode. | [optional] 
**Curve** | Pointer to [**TSSCurveType**](TSSCurveType.md) |  | [optional] 
**Threshold** | Pointer to **int32** | The threshold. | [optional] 
**Participants** | Pointer to [**[]TSSParticipant**](TSSParticipant.md) |  | [optional] 

## Methods

### NewTSSGroup

`func NewTSSGroup() *TSSGroup`

NewTSSGroup instantiates a new TSSGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSGroupWithDefaults

`func NewTSSGroupWithDefaults() *TSSGroup`

NewTSSGroupWithDefaults instantiates a new TSSGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TSSGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TSSGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TSSGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TSSGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCanonicalGroupId

`func (o *TSSGroup) GetCanonicalGroupId() string`

GetCanonicalGroupId returns the CanonicalGroupId field if non-nil, zero value otherwise.

### GetCanonicalGroupIdOk

`func (o *TSSGroup) GetCanonicalGroupIdOk() (*string, bool)`

GetCanonicalGroupIdOk returns a tuple with the CanonicalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalGroupId

`func (o *TSSGroup) SetCanonicalGroupId(v string)`

SetCanonicalGroupId sets CanonicalGroupId field to given value.

### HasCanonicalGroupId

`func (o *TSSGroup) HasCanonicalGroupId() bool`

HasCanonicalGroupId returns a boolean if a field has been set.

### GetProtocolGroupId

`func (o *TSSGroup) GetProtocolGroupId() string`

GetProtocolGroupId returns the ProtocolGroupId field if non-nil, zero value otherwise.

### GetProtocolGroupIdOk

`func (o *TSSGroup) GetProtocolGroupIdOk() (*string, bool)`

GetProtocolGroupIdOk returns a tuple with the ProtocolGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolGroupId

`func (o *TSSGroup) SetProtocolGroupId(v string)`

SetProtocolGroupId sets ProtocolGroupId field to given value.

### HasProtocolGroupId

`func (o *TSSGroup) HasProtocolGroupId() bool`

HasProtocolGroupId returns a boolean if a field has been set.

### GetProtocolType

`func (o *TSSGroup) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *TSSGroup) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *TSSGroup) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *TSSGroup) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *TSSGroup) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TSSGroup) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TSSGroup) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TSSGroup) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetType

`func (o *TSSGroup) GetType() TSSGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TSSGroup) GetTypeOk() (*TSSGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TSSGroup) SetType(v TSSGroupType)`

SetType sets Type field to given value.

### HasType

`func (o *TSSGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRootExtendedPublicKey

`func (o *TSSGroup) GetRootExtendedPublicKey() string`

GetRootExtendedPublicKey returns the RootExtendedPublicKey field if non-nil, zero value otherwise.

### GetRootExtendedPublicKeyOk

`func (o *TSSGroup) GetRootExtendedPublicKeyOk() (*string, bool)`

GetRootExtendedPublicKeyOk returns a tuple with the RootExtendedPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootExtendedPublicKey

`func (o *TSSGroup) SetRootExtendedPublicKey(v string)`

SetRootExtendedPublicKey sets RootExtendedPublicKey field to given value.

### HasRootExtendedPublicKey

`func (o *TSSGroup) HasRootExtendedPublicKey() bool`

HasRootExtendedPublicKey returns a boolean if a field has been set.

### GetChaincode

`func (o *TSSGroup) GetChaincode() string`

GetChaincode returns the Chaincode field if non-nil, zero value otherwise.

### GetChaincodeOk

`func (o *TSSGroup) GetChaincodeOk() (*string, bool)`

GetChaincodeOk returns a tuple with the Chaincode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChaincode

`func (o *TSSGroup) SetChaincode(v string)`

SetChaincode sets Chaincode field to given value.

### HasChaincode

`func (o *TSSGroup) HasChaincode() bool`

HasChaincode returns a boolean if a field has been set.

### GetCurve

`func (o *TSSGroup) GetCurve() TSSCurveType`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *TSSGroup) GetCurveOk() (*TSSCurveType, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *TSSGroup) SetCurve(v TSSCurveType)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *TSSGroup) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### GetThreshold

`func (o *TSSGroup) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TSSGroup) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TSSGroup) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *TSSGroup) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetParticipants

`func (o *TSSGroup) GetParticipants() []TSSParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *TSSGroup) GetParticipantsOk() (*[]TSSParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *TSSGroup) SetParticipants(v []TSSParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *TSSGroup) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


