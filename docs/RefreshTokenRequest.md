# RefreshTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The client ID, a unique identifier to distinguish Cobo Portal Apps. You can get the client ID by retrieving the manifest file after publishing the app. | [optional] 
**GrantType** | Pointer to **string** | The OAuth grant type. Set the value as &#x60;refresh_token&#x60;. | [optional] 
**RefreshToken** | Pointer to **string** | The Refresh Token of the current Org Access Token. | [optional] 

## Methods

### NewRefreshTokenRequest

`func NewRefreshTokenRequest() *RefreshTokenRequest`

NewRefreshTokenRequest instantiates a new RefreshTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenRequestWithDefaults

`func NewRefreshTokenRequestWithDefaults() *RefreshTokenRequest`

NewRefreshTokenRequestWithDefaults instantiates a new RefreshTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *RefreshTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RefreshTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RefreshTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *RefreshTokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetGrantType

`func (o *RefreshTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *RefreshTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *RefreshTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *RefreshTokenRequest) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetRefreshToken

`func (o *RefreshTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RefreshTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RefreshTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *RefreshTokenRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


