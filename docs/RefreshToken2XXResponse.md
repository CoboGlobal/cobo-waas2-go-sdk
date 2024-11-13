# RefreshToken2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | The new Org Access Token. | [optional] 
**TokenType** | Pointer to **string** | The type of the tokens, which is Bearer. | [optional] 
**Scope** | Pointer to **string** | The scope of the Org Access Token to limit the app&#39;s access to the organization&#39;s resources. **Note**: Currently this property value is empty. The scope of the Org Access Token is based on the permissions granted when the app user installs the app.  | [optional] 
**ExpiresIn** | Pointer to **int32** | The time in seconds in which the new Org Access Token expires. | [optional] 
**RefreshToken** | Pointer to **string** | The Refresh Token, used to obtain another Org Access Token when the new Org Access Token expires. The expiration time for Refresh Tokens is currently set to 30 days and is subject to change. | [optional] 

## Methods

### NewRefreshToken2XXResponse

`func NewRefreshToken2XXResponse() *RefreshToken2XXResponse`

NewRefreshToken2XXResponse instantiates a new RefreshToken2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshToken2XXResponseWithDefaults

`func NewRefreshToken2XXResponseWithDefaults() *RefreshToken2XXResponse`

NewRefreshToken2XXResponseWithDefaults instantiates a new RefreshToken2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *RefreshToken2XXResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RefreshToken2XXResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RefreshToken2XXResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *RefreshToken2XXResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetTokenType

`func (o *RefreshToken2XXResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *RefreshToken2XXResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *RefreshToken2XXResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *RefreshToken2XXResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetScope

`func (o *RefreshToken2XXResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RefreshToken2XXResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RefreshToken2XXResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RefreshToken2XXResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExpiresIn

`func (o *RefreshToken2XXResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *RefreshToken2XXResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *RefreshToken2XXResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *RefreshToken2XXResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *RefreshToken2XXResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RefreshToken2XXResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RefreshToken2XXResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *RefreshToken2XXResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


