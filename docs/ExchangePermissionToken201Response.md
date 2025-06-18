# ExchangePermissionToken201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | The new Permission Access Token. | [optional] 
**TokenType** | Pointer to **string** | The type of the tokens, which is Bearer. | [optional] 
**ExpiresIn** | Pointer to **int32** | The time in seconds in which the new Permission Access Token expires. | [optional] 
**RefreshToken** | Pointer to **string** | The Refresh Token, used to obtain another Org Access Token when the new Permission Access Token expires. The expiration time for Refresh Tokens is currently set to 7 days and is subject to change. | [optional] 

## Methods

### NewExchangePermissionToken201Response

`func NewExchangePermissionToken201Response() *ExchangePermissionToken201Response`

NewExchangePermissionToken201Response instantiates a new ExchangePermissionToken201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangePermissionToken201ResponseWithDefaults

`func NewExchangePermissionToken201ResponseWithDefaults() *ExchangePermissionToken201Response`

NewExchangePermissionToken201ResponseWithDefaults instantiates a new ExchangePermissionToken201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *ExchangePermissionToken201Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ExchangePermissionToken201Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ExchangePermissionToken201Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ExchangePermissionToken201Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetTokenType

`func (o *ExchangePermissionToken201Response) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ExchangePermissionToken201Response) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ExchangePermissionToken201Response) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *ExchangePermissionToken201Response) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetExpiresIn

`func (o *ExchangePermissionToken201Response) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ExchangePermissionToken201Response) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ExchangePermissionToken201Response) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ExchangePermissionToken201Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *ExchangePermissionToken201Response) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ExchangePermissionToken201Response) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ExchangePermissionToken201Response) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ExchangePermissionToken201Response) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


