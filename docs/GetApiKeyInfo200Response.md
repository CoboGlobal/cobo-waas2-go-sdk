# GetApiKeyInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The API key name. | 
**CurveType** | **string** | The curve type used for the API key, which determines the cryptographic algorithm for key generation and signing. Possible values include: - &#x60;ED25519&#x60;: Ed25519 - &#x60;SECP256K1&#x60;: Secp256k1  | 
**Key** | **string** | The API key value. | 
**CallbackUrl** | Pointer to **string** | The URL of the callback endpoint that receives callback messages triggered by this API key. | [optional] 
**ValidIps** | Pointer to **[]string** | (Applicable to permanent API keys only) The list of IP addresses that are allowed to use this API key. | [optional] 
**CreatedTimestamp** | **int64** | The time when the API key was registered, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTimestamp** | **int64** | The time when the API key information was last updated, in Unix timestamp format, measured in milliseconds. | 
**ExpiredTimestamp** | Pointer to **int64** | The time when the API key expires, in Unix timestamp format, measured in milliseconds. For permanent API keys, this property value is &#x60;null&#x60;. | [optional] 
**RoleScopes** | Pointer to [**[]RoleScopes**](RoleScopes.md) | The list of user roles and wallet scopes associated with the API key. | [optional] 

## Methods

### NewGetApiKeyInfo200Response

`func NewGetApiKeyInfo200Response(name string, curveType string, key string, createdTimestamp int64, updatedTimestamp int64, ) *GetApiKeyInfo200Response`

NewGetApiKeyInfo200Response instantiates a new GetApiKeyInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiKeyInfo200ResponseWithDefaults

`func NewGetApiKeyInfo200ResponseWithDefaults() *GetApiKeyInfo200Response`

NewGetApiKeyInfo200ResponseWithDefaults instantiates a new GetApiKeyInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetApiKeyInfo200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApiKeyInfo200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApiKeyInfo200Response) SetName(v string)`

SetName sets Name field to given value.


### GetCurveType

`func (o *GetApiKeyInfo200Response) GetCurveType() string`

GetCurveType returns the CurveType field if non-nil, zero value otherwise.

### GetCurveTypeOk

`func (o *GetApiKeyInfo200Response) GetCurveTypeOk() (*string, bool)`

GetCurveTypeOk returns a tuple with the CurveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurveType

`func (o *GetApiKeyInfo200Response) SetCurveType(v string)`

SetCurveType sets CurveType field to given value.


### GetKey

`func (o *GetApiKeyInfo200Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetApiKeyInfo200Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetApiKeyInfo200Response) SetKey(v string)`

SetKey sets Key field to given value.


### GetCallbackUrl

`func (o *GetApiKeyInfo200Response) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *GetApiKeyInfo200Response) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *GetApiKeyInfo200Response) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *GetApiKeyInfo200Response) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetValidIps

`func (o *GetApiKeyInfo200Response) GetValidIps() []string`

GetValidIps returns the ValidIps field if non-nil, zero value otherwise.

### GetValidIpsOk

`func (o *GetApiKeyInfo200Response) GetValidIpsOk() (*[]string, bool)`

GetValidIpsOk returns a tuple with the ValidIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidIps

`func (o *GetApiKeyInfo200Response) SetValidIps(v []string)`

SetValidIps sets ValidIps field to given value.

### HasValidIps

`func (o *GetApiKeyInfo200Response) HasValidIps() bool`

HasValidIps returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *GetApiKeyInfo200Response) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *GetApiKeyInfo200Response) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *GetApiKeyInfo200Response) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *GetApiKeyInfo200Response) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *GetApiKeyInfo200Response) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *GetApiKeyInfo200Response) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetExpiredTimestamp

`func (o *GetApiKeyInfo200Response) GetExpiredTimestamp() int64`

GetExpiredTimestamp returns the ExpiredTimestamp field if non-nil, zero value otherwise.

### GetExpiredTimestampOk

`func (o *GetApiKeyInfo200Response) GetExpiredTimestampOk() (*int64, bool)`

GetExpiredTimestampOk returns a tuple with the ExpiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimestamp

`func (o *GetApiKeyInfo200Response) SetExpiredTimestamp(v int64)`

SetExpiredTimestamp sets ExpiredTimestamp field to given value.

### HasExpiredTimestamp

`func (o *GetApiKeyInfo200Response) HasExpiredTimestamp() bool`

HasExpiredTimestamp returns a boolean if a field has been set.

### GetRoleScopes

`func (o *GetApiKeyInfo200Response) GetRoleScopes() []RoleScopes`

GetRoleScopes returns the RoleScopes field if non-nil, zero value otherwise.

### GetRoleScopesOk

`func (o *GetApiKeyInfo200Response) GetRoleScopesOk() (*[]RoleScopes, bool)`

GetRoleScopesOk returns a tuple with the RoleScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleScopes

`func (o *GetApiKeyInfo200Response) SetRoleScopes(v []RoleScopes)`

SetRoleScopes sets RoleScopes field to given value.

### HasRoleScopes

`func (o *GetApiKeyInfo200Response) HasRoleScopes() bool`

HasRoleScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


